"""
Admin router — operator approval workflow for CAPIF invokers.

Mounted at /admin by main.py. In production, protect these endpoints
with an admin API key or mTLS — they are NOT authenticated here to
keep the demo self-contained.

  GET  /admin/invokers                    list all (filter by ?status=)
  GET  /admin/invokers/{id}               invoker detail + audit history
  POST /admin/invokers/{id}/approve       approve + set scopes → Keycloak client created
  POST /admin/invokers/{id}/reject        reject with reason
  POST /admin/invokers/{id}/revoke        suspend an approved invoker
  GET  /admin/audit                       paginated audit log
"""

import logging
import os
import secrets as secrets_mod

from fastapi import APIRouter, Depends, Header, HTTPException, Query
from pydantic import BaseModel

from keycloak_bridge import create_keycloak_client, delete_keycloak_client
from acl_bridge import create_acl_entries, remove_acl_entries
from db import (
    invokers as _col,
    audit_logs as _audit,
    audit as _audit_log,
    now as _now,
    SAFE_PROJECTION,
    encrypt,
)

log = logging.getLogger("invoker-onboarding.admin")

# When INVOKER_ADMIN_API_KEY is set, all /admin/* endpoints require
# `X-Admin-Api-Key: <key>` header matching. When unset (dev mode), routes
# are open — the dashboard middleware is then the only gate.
_ADMIN_API_KEY = os.getenv("INVOKER_ADMIN_API_KEY", "")


def _require_admin_key(x_admin_api_key: str = Header(None)):
    if not _ADMIN_API_KEY:
        return                       # dev mode — no enforcement
    if not x_admin_api_key or not secrets_mod.compare_digest(x_admin_api_key, _ADMIN_API_KEY):
        raise HTTPException(status_code=401, detail="Admin API key missing or invalid")


router = APIRouter(dependencies=[Depends(_require_admin_key)])


# ── Schemas ────────────────────────────────────────────────────────────────────

class ApproveRequest(BaseModel):
    scopes_approved: list[str]          # subset of CAMARA API names to grant
    approved_by:     str = "admin"      # admin username or email
    note:            str | None = None


class RejectRequest(BaseModel):
    rejection_reason: str
    rejected_by:      str = "admin"


class RevokeRequest(BaseModel):
    reason:     str
    revoked_by: str = "admin"


class InvokerSummary(BaseModel):
    invoker_id:      str
    invoker_name:    str
    approval_status: str
    submitted_at:    str
    contact_email:   str | None = None
    company:         str | None = None
    requested_apis:  list[str] = []
    scopes_approved: list[str] = []


class InvokerDetail(InvokerSummary):
    description:       str | None = None
    use_case:          str | None = None
    notification_url:  str | None = None
    approved_at:       str | None = None
    approved_by:       str | None = None
    rejection_reason:  str | None = None
    keycloak_client_id: str | None = None
    audit_history:     list[dict] = []


class AuditEntry(BaseModel):
    action:     str
    invoker_id: str
    actor:      str
    timestamp:  str
    detail:     dict


# ── Helpers ────────────────────────────────────────────────────────────────────

def _doc_to_summary(doc: dict) -> InvokerSummary:
    return InvokerSummary(
        invoker_id=doc["invoker_id"],
        invoker_name=doc["invoker_name"],
        approval_status=doc["approval_status"],
        submitted_at=doc["submitted_at"].isoformat(),
        contact_email=doc.get("submitted_by", {}).get("email"),
        company=doc.get("submitted_by", {}).get("company"),
        requested_apis=doc.get("requested_apis", []),
        scopes_approved=doc.get("scopes_approved", []),
    )


# ── Routes ─────────────────────────────────────────────────────────────────────

@router.get("/invokers", response_model=list[InvokerSummary])
def list_invokers(status: str | None = Query(None, description="Filter by approval_status")):
    """List all registered invokers, optionally filtered by status."""
    query = {"approval_status": status} if status else {}
    docs = list(_col.find(query, SAFE_PROJECTION).sort("submitted_at", -1))
    return [_doc_to_summary(d) for d in docs]


@router.get("/invokers/{invoker_id}", response_model=InvokerDetail)
def get_invoker_detail(invoker_id: str):
    """Full invoker detail including audit history."""
    doc = _col.find_one({"invoker_id": invoker_id}, SAFE_PROJECTION)
    if not doc:
        raise HTTPException(404, f"Invoker {invoker_id} not found")

    audit_entries = list(
        _audit.find({"invoker_id": invoker_id}, {"_id": 0}).sort("timestamp", -1).limit(50)
    )
    for e in audit_entries:
        e["timestamp"] = e["timestamp"].isoformat()

    return InvokerDetail(
        invoker_id=doc["invoker_id"],
        invoker_name=doc["invoker_name"],
        approval_status=doc["approval_status"],
        submitted_at=doc["submitted_at"].isoformat(),
        contact_email=doc.get("submitted_by", {}).get("email"),
        company=doc.get("submitted_by", {}).get("company"),
        use_case=doc.get("submitted_by", {}).get("use_case"),
        description=doc.get("description"),
        notification_url=doc.get("notification_url"),
        requested_apis=doc.get("requested_apis", []),
        scopes_approved=doc.get("scopes_approved", []),
        approved_at=doc["approved_at"].isoformat() if doc.get("approved_at") else None,
        approved_by=doc.get("approved_by"),
        rejection_reason=doc.get("rejection_reason"),
        keycloak_client_id=doc.get("keycloak_client_id"),
        audit_history=audit_entries,
    )


@router.post("/invokers/{invoker_id}/approve", status_code=200)
def approve_invoker(invoker_id: str, req: ApproveRequest):
    """
    Approve an invoker registration.

    1. Sets approval_status → 'approved'
    2. Creates a Keycloak client scoped to req.scopes_approved
    3. Stores the Keycloak client_id for the invoker to use
    """
    doc = _col.find_one({"invoker_id": invoker_id})
    if not doc:
        raise HTTPException(404, f"Invoker {invoker_id} not found")

    if doc["approval_status"] == "approved":
        raise HTTPException(409, "Invoker is already approved")
    if doc["approval_status"] == "rejected":
        raise HTTPException(409, "Rejected invokers cannot be approved — use a new registration")

    if not req.scopes_approved:
        raise HTTPException(422, "At least one scope must be approved")

    # Create Keycloak client with approved scopes
    try:
        kc = create_keycloak_client(
            invoker_id=invoker_id,
            invoker_name=doc["invoker_name"],
            approved_scopes=req.scopes_approved,
        )
    except Exception as e:
        log.error("Keycloak client creation failed for %s: %s", invoker_id, e)
        raise HTTPException(502, f"Keycloak client creation failed: {e}")

    now = _now()
    # Create CAPIF ACL entries for approved scopes (best-effort)
    acl_published = create_acl_entries(invoker_id, req.scopes_approved)
    if len(acl_published) < len(req.scopes_approved):
        log.warning(
            "ACL entries created for %d/%d scopes for invoker %s",
            len(acl_published), len(req.scopes_approved), invoker_id,
        )

    _col.update_one(
        {"invoker_id": invoker_id},
        {"$set": {
            "approval_status":      "approved",
            "scopes_approved":      req.scopes_approved,
            "approved_at":          now,
            "approved_by":          req.approved_by,
            "rejection_reason":     None,
            "keycloak_client_id":   kc["client_id"],
            "secrets.keycloak_uuid":   kc["keycloak_uuid"],
            "secrets.keycloak_secret": encrypt(kc["client_secret"]),
            "secrets.acl_published":   acl_published,
        }},
    )
    _audit_log(
        "approved", invoker_id,
        actor=req.approved_by,
        detail={
            "scopes_approved":  req.scopes_approved,
            "acl_published":    acl_published,
            "keycloak_client":  kc["client_id"],
            "note":             req.note,
        },
    )
    log.info("Invoker %s approved by %s  scopes=%s  acl=%s", invoker_id, req.approved_by, req.scopes_approved, acl_published)

    return {
        "invoker_id":         invoker_id,
        "approval_status":    "approved",
        "scopes_approved":    req.scopes_approved,
        "keycloak_client_id": kc["client_id"],
        "keycloak_secret":    kc["client_secret"],
        "message": (
            f"Invoker approved. Share keycloak_client_id and keycloak_secret "
            f"with the developer to obtain tokens via Keycloak."
        ),
    }


@router.post("/invokers/{invoker_id}/reject", status_code=200)
def reject_invoker(invoker_id: str, req: RejectRequest):
    """Reject a pending invoker registration."""
    doc = _col.find_one({"invoker_id": invoker_id})
    if not doc:
        raise HTTPException(404, f"Invoker {invoker_id} not found")

    if doc["approval_status"] not in ("pending",):
        raise HTTPException(409, f"Only pending invokers can be rejected (current: {doc['approval_status']})")

    _col.update_one(
        {"invoker_id": invoker_id},
        {"$set": {
            "approval_status":  "rejected",
            "rejection_reason": req.rejection_reason,
            "rejected_by":      req.rejected_by,
            "rejected_at":      _now(),
        }},
    )
    _audit_log(
        "rejected", invoker_id,
        actor=req.rejected_by,
        detail={"reason": req.rejection_reason},
    )
    log.info("Invoker %s rejected by %s: %s", invoker_id, req.rejected_by, req.rejection_reason)

    return {
        "invoker_id":      invoker_id,
        "approval_status": "rejected",
        "rejection_reason": req.rejection_reason,
    }


@router.post("/invokers/{invoker_id}/revoke", status_code=200)
def revoke_invoker(invoker_id: str, req: RevokeRequest):
    """
    Suspend an approved invoker.

    Deletes the Keycloak client so its tokens immediately stop working at Kong.
    The CAPIF registration is left intact; it can be re-approved without re-onboarding.
    """
    doc = _col.find_one({"invoker_id": invoker_id})
    if not doc:
        raise HTTPException(404, f"Invoker {invoker_id} not found")

    if doc["approval_status"] != "approved":
        raise HTTPException(409, f"Only approved invokers can be revoked (current: {doc['approval_status']})")

    kc_uuid = (doc.get("secrets") or {}).get("keycloak_uuid")
    if kc_uuid:
        delete_keycloak_client(kc_uuid)

    # Remove CAPIF ACL entries so the invoker can't call NEF-direct paths
    remove_acl_entries(invoker_id)

    _col.update_one(
        {"invoker_id": invoker_id},
        {"$set": {
            "approval_status":         "suspended",
            "keycloak_client_id":      None,
            "secrets.keycloak_uuid":   None,
            "secrets.keycloak_secret": None,
            "suspended_at":       _now(),
            "suspended_by":       req.revoked_by,
        }},
    )
    _audit_log(
        "revoked", invoker_id,
        actor=req.revoked_by,
        detail={"reason": req.reason},
    )
    log.info("Invoker %s suspended by %s", invoker_id, req.revoked_by)

    return {
        "invoker_id":      invoker_id,
        "approval_status": "suspended",
        "message": "Invoker suspended. Keycloak client deleted — active tokens will be rejected by Kong.",
    }


@router.get("/audit", response_model=list[AuditEntry])
def get_audit_log(
    invoker_id: str | None = Query(None),
    action:     str | None = Query(None),
    limit:      int        = Query(100, le=500),
):
    """Paginated audit log of all governance actions."""
    query = {}
    if invoker_id:
        query["invoker_id"] = invoker_id
    if action:
        query["action"] = action

    entries = list(_audit.find(query, {"_id": 0}).sort("timestamp", -1).limit(limit))
    for e in entries:
        e["timestamp"] = e["timestamp"].isoformat()
    return entries
