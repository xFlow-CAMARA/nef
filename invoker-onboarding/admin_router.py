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
import secrets
from typing import Literal

import pymongo
from fastapi import APIRouter, Depends, Header, HTTPException, Query
from pydantic import BaseModel

from acl_bridge import create_acl_entries, remove_acl_entries
from db import (
    ADMIN_PROJECTION,
    encrypt,
)
from db import (
    audit as _audit_log,
)
from db import (
    audit_logs as _audit,
)
from db import (
    invokers as _col,
)
from db import (
    now as _now,
)
from keycloak_bridge import create_keycloak_client, delete_keycloak_client, rotate_keycloak_client_secret

log = logging.getLogger("invoker-onboarding.admin")


def _require_admin_key(x_admin_api_key: str | None = Header(None)):
    """Header-based gate. The expected key is read on every call so rotating
    env at runtime (or monkeypatching in tests) takes effect immediately.
    If INVOKER_ADMIN_API_KEY is unset, routes are open (dev mode).
    """
    expected = os.getenv("INVOKER_ADMIN_API_KEY", "")
    if not expected:
        return
    if not x_admin_api_key or not secrets.compare_digest(x_admin_api_key, expected):
        raise HTTPException(status_code=401, detail="Admin API key missing or invalid")


router = APIRouter(dependencies=[Depends(_require_admin_key)])


# ── Schemas ────────────────────────────────────────────────────────────────────

class ApproveRequest(BaseModel):
    scopes_approved: list[str]          # subset of CAMARA API names to grant
    approved_by:     str                # admin username/email — injected by dashboard
    note:            str | None = None


class RejectRequest(BaseModel):
    rejection_reason: str
    rejected_by:      str               # injected by dashboard


class RevokeRequest(BaseModel):
    reason:     str
    revoked_by: str                     # injected by dashboard


class RotateRequest(BaseModel):
    reason:     str | None = None
    rotated_by: str                     # injected by dashboard


ApprovalStatus = Literal[
    "pending", "approved", "rejected", "suspended",
    "approving", "rotating",     # transient
]


class InvokerSummary(BaseModel):
    invoker_id:      str
    invoker_name:    str
    approval_status: ApprovalStatus
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

class InvokerList(BaseModel):
    items:    list[InvokerSummary]
    has_more: bool
    skip:     int
    limit:    int


@router.get("/invokers", response_model=InvokerList)
def list_invokers(
    status: str | None = Query(None, description="Filter by approval_status"),
    limit:  int        = Query(50, ge=1, le=500),
    skip:   int        = Query(0,  ge=0),
):
    """List registered invokers, newest first. Paginated to keep responses
    bounded for installations with many invokers."""
    query = {"approval_status": status} if status else {}
    # Fetch one extra to detect a next page without a separate count query.
    docs = list(
        _col.find(query, ADMIN_PROJECTION)
            .sort("submitted_at", -1)
            .skip(skip)
            .limit(limit + 1)
    )
    has_more = len(docs) > limit
    items = [_doc_to_summary(d) for d in docs[:limit]]
    return InvokerList(items=items, has_more=has_more, skip=skip, limit=limit)


@router.get("/invokers/{invoker_id}", response_model=InvokerDetail)
def get_invoker_detail(invoker_id: str):
    """Full invoker detail including audit history."""
    doc = _col.find_one({"invoker_id": invoker_id}, ADMIN_PROJECTION)
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

    Three external mutations happen here:
      1. Create the Keycloak client (network)
      2. Publish CAPIF ACL events to Redis      (network)
      3. Persist approval state in Mongo

    Each can fail independently. If step 3 fails after 1/2 succeeded, we
    end up with an orphan Keycloak client + dangling ACL entries that the
    next approve attempt will collide with. Compensate explicitly so the
    system either fully commits or fully rolls back.
    """
    if not req.scopes_approved:
        raise HTTPException(422, "At least one scope must be approved")

    # Atomic claim: transition pending → approving in one Mongo operation so
    # two concurrent approvals can't both proceed (would otherwise create
    # duplicate Keycloak clients and orphan one). Anyone else trying to
    # approve at the same time sees a 409.
    doc = _col.find_one_and_update(
        {"invoker_id": invoker_id, "approval_status": "pending"},
        {"$set": {"approval_status": "approving"}},
        return_document=pymongo.ReturnDocument.BEFORE,
    )
    if doc is None:
        existing = _col.find_one({"invoker_id": invoker_id}, {"approval_status": 1})
        if existing is None:
            raise HTTPException(404, f"Invoker {invoker_id} not found")
        status_now = existing["approval_status"]
        if status_now == "approving":
            raise HTTPException(409, "Invoker is currently being approved by another request")
        if status_now == "approved":
            raise HTTPException(409, "Invoker is already approved")
        if status_now == "rejected":
            raise HTTPException(409, "Rejected invokers cannot be approved — use a new registration")
        raise HTTPException(409, f"Invoker not in pending state (current: {status_now})")

    # Step 1: Keycloak. Failure here means rolling back our transient state.
    try:
        kc = create_keycloak_client(
            invoker_id=invoker_id,
            invoker_name=doc["invoker_name"],
            approved_scopes=req.scopes_approved,
        )
    except Exception as e:
        log.error("Keycloak client creation failed for %s: %s", invoker_id, e)
        # Revert the transient 'approving' state back to 'pending' so the
        # admin can retry without hitting the 409 above.
        _col.update_one(
            {"invoker_id": invoker_id, "approval_status": "approving"},
            {"$set": {"approval_status": "pending"}},
        )
        raise HTTPException(502, "Keycloak client creation failed") from None

    # Step 2: ACL events (best-effort by design — we still proceed even if
    # some scopes aren't published, but we track which were).
    acl_published = create_acl_entries(invoker_id, req.scopes_approved)
    if len(acl_published) < len(req.scopes_approved):
        log.warning(
            "ACL entries created for %d/%d scopes for invoker %s",
            len(acl_published), len(req.scopes_approved), invoker_id,
        )

    # Step 3: Mongo. If this fails we MUST undo the Keycloak client (else
    # the next retry collides with 409 client-exists), and try to take
    # back the ACL entries we already announced.
    now = _now()
    try:
        _col.update_one(
            {"invoker_id": invoker_id},
            {"$set": {
                "approval_status":         "approved",
                "scopes_approved":         req.scopes_approved,
                "approved_at":             now,
                "approved_by":             req.approved_by,
                "rejection_reason":        None,
                "keycloak_client_id":      kc["client_id"],
                "secrets.keycloak_secret": encrypt(kc["client_secret"]),
                "internal.keycloak_uuid":  kc["keycloak_uuid"],
                "internal.acl_published":  acl_published,
            }},
        )
    except Exception as e:
        log.error("Mongo update failed for %s after external provisioning: %s — compensating", invoker_id, e)
        try:
            delete_keycloak_client(kc["keycloak_uuid"])
        except Exception as ke:
            log.error("Keycloak rollback failed for %s: %s — manual cleanup needed", invoker_id, ke)
        if acl_published:
            try:
                remove_acl_entries(invoker_id)
            except Exception as ae:
                log.error("ACL rollback failed for %s: %s", invoker_id, ae)
        # Revert transient state to pending so the admin can retry.
        _col.update_one(
            {"invoker_id": invoker_id, "approval_status": "approving"},
            {"$set": {"approval_status": "pending"}},
        )
        raise HTTPException(503, "Approval failed mid-flight; please retry") from None
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

    # The secret is intentionally NOT returned here. Admins (or the
    # developer) fetch it from GET /invokers/{id}/credentials, which
    # records an audit row for every reveal so each disclosure leaves
    # a trail. This keeps the credential out of the approve response
    # body (and out of any access log that captures it).
    return {
        "invoker_id":         invoker_id,
        "approval_status":    "approved",
        "scopes_approved":    req.scopes_approved,
        "keycloak_client_id": kc["client_id"],
        "message": "Invoker approved. Fetch credentials via GET /invokers/{id}/credentials.",
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

    kc_uuid = (doc.get("internal") or {}).get("keycloak_uuid")
    if kc_uuid:
        delete_keycloak_client(kc_uuid)

    # Remove CAPIF ACL entries so the invoker can't call NEF-direct paths
    remove_acl_entries(invoker_id)

    _col.update_one(
        {"invoker_id": invoker_id},
        {"$set": {
            "approval_status":         "suspended",
            "keycloak_client_id":      None,
            "secrets.keycloak_secret": None,
            "internal.keycloak_uuid":  None,
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


@router.post("/invokers/{invoker_id}/rotate-secret", status_code=200)
def rotate_invoker_secret(invoker_id: str, req: RotateRequest):
    """Generate a new Keycloak client secret for an approved invoker.

    Used when the developer has lost their secret. The OLD secret stops
    working immediately at Keycloak. The NEW secret is encrypted into the
    invoker doc and returned ONCE in the response — fetch via /credentials
    afterwards (also audited).
    """
    # Atomic claim against concurrent rotates. Transition approved → rotating
    # in a single Mongo op; if two admins click Rotate at the same time, the
    # second sees 409.
    doc = _col.find_one_and_update(
        {"invoker_id": invoker_id, "approval_status": "approved"},
        {"$set": {"approval_status": "rotating"}},
        return_document=pymongo.ReturnDocument.BEFORE,
    )
    if doc is None:
        existing = _col.find_one({"invoker_id": invoker_id}, {"approval_status": 1})
        if existing is None:
            raise HTTPException(404, f"Invoker {invoker_id} not found")
        if existing["approval_status"] == "rotating":
            raise HTTPException(409, "Invoker is currently being rotated by another request")
        raise HTTPException(409, "Only approved invokers have a secret to rotate")

    kc_uuid = (doc.get("internal") or {}).get("keycloak_uuid")
    if not kc_uuid:
        # Restore state so the admin doesn't get stuck in 'rotating'.
        _col.update_one(
            {"invoker_id": invoker_id, "approval_status": "rotating"},
            {"$set": {"approval_status": "approved"}},
        )
        raise HTTPException(500, "No Keycloak UUID stored — rotation impossible")

    try:
        new_secret = rotate_keycloak_client_secret(kc_uuid)
    except Exception as e:
        log.error("Keycloak rotate failed for %s: %s", invoker_id, e)
        _col.update_one(
            {"invoker_id": invoker_id, "approval_status": "rotating"},
            {"$set": {"approval_status": "approved"}},
        )
        raise HTTPException(502, "Keycloak rotation failed") from None

    # If we can't persist the new secret to Mongo, Keycloak holds the new
    # value but /credentials would hand the developer the OLD one. Try ONCE
    # more to put the doc in a consistent state (re-rotate + re-store). If
    # that fails too, mark the doc as inconsistent and audit loudly so an
    # operator can fix it by hand via Keycloak admin.
    try:
        _col.update_one(
            {"invoker_id": invoker_id},
            {"$set": {"secrets.keycloak_secret": encrypt(new_secret)}},
        )
    except Exception as e1:
        log.error("Mongo write failed after Keycloak rotate for %s: %s — re-rotating", invoker_id, e1)
        try:
            new_secret = rotate_keycloak_client_secret(kc_uuid)
            _col.update_one(
                {"invoker_id": invoker_id},
                {"$set": {"secrets.keycloak_secret": encrypt(new_secret)}},
            )
        except Exception as e2:
            log.error("Re-rotation also failed for %s: %s — marking inconsistent", invoker_id, e2)
            try:
                _col.update_one(
                    {"invoker_id": invoker_id},
                    {"$set": {"internal.rotation_inconsistent": True}},
                )
            except Exception:
                pass
            _audit_log(
                "secret_rotation_inconsistent", invoker_id,
                actor=req.rotated_by,
                detail={"reason": req.reason, "error": str(e2)[:200]},
            )
            raise HTTPException(
                503,
                "Rotation failed — Keycloak and Mongo are out of sync; "
                "regenerate the secret manually via Keycloak admin and update Mongo",
            ) from None

    # Release the transient 'rotating' state.
    _col.update_one(
        {"invoker_id": invoker_id, "approval_status": "rotating"},
        {"$set": {"approval_status": "approved"}},
    )
    _audit_log(
        "secret_rotated", invoker_id,
        actor=req.rotated_by,
        detail={"reason": req.reason},
    )
    log.info("Rotated Keycloak secret for %s by %s", invoker_id, req.rotated_by)

    return {
        "invoker_id":        invoker_id,
        "keycloak_secret":   new_secret,
        "message": "Secret rotated. The previous secret no longer works.",
    }


@router.get("/audit", response_model=list[AuditEntry])
def get_audit_log(
    invoker_id: str | None = Query(None),
    action:     str | None = Query(None),
    limit:      int        = Query(100, ge=1, le=500),
    skip:       int        = Query(0,   ge=0),
):
    """Paginated audit log of all governance actions, newest first."""
    query: dict = {}
    if invoker_id:
        query["invoker_id"] = invoker_id
    if action:
        query["action"] = action

    entries = list(
        _audit.find(query, {"_id": 0})
              .sort("timestamp", -1)
              .skip(skip)
              .limit(limit)
    )
    for e in entries:
        e["timestamp"] = e["timestamp"].isoformat()
    return entries
