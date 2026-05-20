"""
CAPIF Invoker Onboarding Service

Wraps the ETSI OpenCAPIF invoker management and security APIs into a simple
REST interface. Adds a developer-portal approval workflow on top of raw CAPIF:

  Developer flow
  ──────────────
  POST /invokers           — submit registration request (status: pending)
  GET  /invokers/{id}      — check own registration status
  POST /invokers/{id}/token — get CAPIF token (only if approved)
  GET  /apis               — discover published CAMARA APIs

  Admin flow
  ──────────
  see admin_router.py (mounted at /admin)

Approval state machine
  pending → approved  (admin approves, Keycloak client created)
  pending → rejected  (admin rejects)
  approved → suspended (admin revokes)

Environment variables
─────────────────────
CAPIF_CORE_URL      Base HTTPS URL of the CAPIF nginx proxy
CAPIF_REGISTER_URL  Base HTTPS URL of the register service
CAPIF_USERNAME      Register credentials for bootstrap token
CAPIF_PASSWORD      Register credentials for bootstrap token
CAPIF_SERVICE_URL   Internal URL of the NEF capif-service for catalog lookup
MONGODB_URI         MongoDB connection string (default: mongodb://camara-mongodb:27017/camara)
"""

import os
import logging
import secrets
import tempfile
from contextlib import contextmanager
from datetime import datetime, timezone

import httpx
from fastapi import Depends, FastAPI, Header, HTTPException
from fastapi.middleware.cors import CORSMiddleware
from pydantic import BaseModel, Field, StringConstraints
from typing import Annotated
from cryptography import x509
from cryptography.x509.oid import NameOID
from cryptography.hazmat.primitives import hashes, serialization
from cryptography.hazmat.primitives.asymmetric import rsa

from db import (
    invokers as _col,
    audit as _audit_log,
    now as _now,
    SAFE_PROJECTION,
    encrypt,
    decrypt,
)
from config import CAMARA_SCOPES, CamaraScope, required, assert_dev_is_local
from admin_router import router as admin_router

logging.basicConfig(level=logging.INFO)
log = logging.getLogger("invoker-onboarding")

# Fail fast if APP_ENV=dev is paired with non-local upstreams.
assert_dev_is_local()

CAPIF_CORE_URL     = os.getenv("CAPIF_CORE_URL",     "https://capifcore:443")
CAPIF_REGISTER_URL = os.getenv("CAPIF_REGISTER_URL", "https://register:8080")
CAPIF_USERNAME     = required("CAPIF_USERNAME", "nef-xflow")
CAPIF_PASSWORD     = required("CAPIF_PASSWORD", "xflow-nef-2026")
CAPIF_SERVICE_URL  = os.getenv("CAPIF_SERVICE_URL",  "http://capif-service:8080")

# TLS verification of outbound calls. Default to verifying; only disable
# explicitly via env when talking to OpenCAPIF's self-signed certs in dev.
_CA_BUNDLE = os.getenv("CAPIF_CA_BUNDLE")          # path to a CA file
_VERIFY    = os.getenv("CAPIF_TLS_VERIFY", "true").lower() not in ("0", "false", "no")
_http = httpx.Client(verify=_CA_BUNDLE if _CA_BUNDLE else _VERIFY, timeout=15.0)


# ── App ────────────────────────────────────────────────────────────────────────
app = FastAPI(
    title="CAPIF Invoker Onboarding",
    description="Register applications as CAPIF invokers with operator approval workflow",
    version="2.0.0",
)

app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],
    allow_methods=["*"],
    allow_headers=["*"],
)


@app.on_event("shutdown")
def _shutdown_clients() -> None:
    """Tidy up module-level singletons on graceful shutdown so upstreams
    don't see half-open connections."""
    try:
        _http.close()
    except Exception:
        pass
    # acl_bridge holds a pooled httpx client AND a lazy redis client.
    try:
        import acl_bridge
        acl_bridge._http.close()
        if acl_bridge._redis is not None:
            acl_bridge._redis.close()
            acl_bridge._redis = None
    except Exception:
        pass
    # keycloak_bridge's httpx client.
    try:
        import keycloak_bridge
        keycloak_bridge._http.close()
    except Exception:
        pass


def _require_dev_key(x_dev_api_key: str | None = Header(None)) -> None:
    """Optional gate for /invokers/* developer endpoints.
    When INVOKER_DEV_API_KEY is set, the dashboard (and only the dashboard)
    must include it. Unset means dev mode — open to anyone on the network.
    Read on every call so rotation/tests work without import-time gymnastics.
    """
    expected = os.getenv("INVOKER_DEV_API_KEY", "")
    if not expected:
        return
    if not x_dev_api_key or not secrets.compare_digest(x_dev_api_key, expected):
        raise HTTPException(status_code=401, detail="Developer API key missing or invalid")


def _require_admin_key(x_admin_api_key: str | None = Header(None)) -> None:
    """Same shape as the admin_router gate — used here for destructive ops
    on root-level routes (e.g. DELETE /invokers/{id})."""
    expected = os.getenv("INVOKER_ADMIN_API_KEY", "")
    if not expected:
        return
    if not x_admin_api_key or not secrets.compare_digest(x_admin_api_key, expected):
        raise HTTPException(status_code=401, detail="Admin API key missing or invalid")


# ── CAPIF helpers ──────────────────────────────────────────────────────────────

def _bootstrap_token() -> str:
    r = _http.get(
        f"{CAPIF_REGISTER_URL}/getauth",
        auth=(CAPIF_USERNAME, CAPIF_PASSWORD),
    )
    r.raise_for_status()
    return r.json()["access_token"]


def _capif_post(path: str, body: dict, token: str) -> dict:
    r = _http.post(
        f"{CAPIF_CORE_URL}/{path.lstrip('/')}",
        json=body,
        headers={"Authorization": f"Bearer {token}"},
    )
    if r.status_code not in (200, 201):
        # Log the full upstream response server-side; return a sanitized
        # message to the caller so internal hostnames / stack traces don't
        # flow to the browser.
        log.error("CAPIF %s → %d: %s", path, r.status_code, r.text[:500])
        raise HTTPException(
            status_code=502 if r.status_code >= 500 else r.status_code,
            detail=f"CAPIF rejected request ({r.status_code})",
        )
    return r.json()


@contextmanager
def _mtls_client(cert_pem: str, key_pem: str):
    """
    Yield an httpx.Client configured with the invoker's mTLS cert.
    Tempfiles holding the PEM bytes and the client itself are torn down
    when the context exits so we don't leak FDs or leave keys in /tmp.
    """
    cf = tempfile.NamedTemporaryFile(suffix=".pem", mode="w", delete=False)
    kf = tempfile.NamedTemporaryFile(suffix=".key", mode="w", delete=False)
    try:
        cf.write(cert_pem); cf.flush(); cf.close()
        kf.write(key_pem);  kf.flush(); kf.close()
        with httpx.Client(cert=(cf.name, kf.name), verify=False, timeout=15.0) as client:
            yield client
    finally:
        for p in (cf.name, kf.name):
            try:
                os.unlink(p)
            except OSError:
                pass


def _generate_key_and_csr(common_name: str):
    key = rsa.generate_private_key(public_exponent=65537, key_size=2048)
    csr = (
        x509.CertificateSigningRequestBuilder()
        .subject_name(x509.Name([x509.NameAttribute(NameOID.COMMON_NAME, common_name)]))
        .sign(key, hashes.SHA256())
    )
    key_pem = key.private_bytes(
        serialization.Encoding.PEM,
        serialization.PrivateFormat.TraditionalOpenSSL,
        serialization.NoEncryption(),
    ).decode()
    csr_pem = csr.public_bytes(serialization.Encoding.PEM).decode()
    return key_pem, csr_pem


# ── Schemas ────────────────────────────────────────────────────────────────────

class OnboardRequest(BaseModel):
    # Tight invoker_name regex: lower-case-safe characters only, 3–40 chars.
    # This becomes part of a Keycloak client_id (see submit_registration),
    # so it must be a value Keycloak will accept without further escaping.
    invoker_name: Annotated[
        str,
        StringConstraints(min_length=3, max_length=40, pattern=r"^[A-Za-z0-9][A-Za-z0-9 _-]+[A-Za-z0-9]$"),
    ]
    description:      str | None = None
    notification_url: str | None = "http://localhost/capif-callback"
    # Developer identity — stored for admin review
    contact_email:    str | None = None
    company:          str | None = None
    use_case:         str | None = None
    # Restricted to the canonical CAMARA scope list (config.CAMARA_SCOPES).
    # Unknown values now produce a 422 at the boundary instead of a silent
    # garbage-in/garbage-out admin review item.
    requested_apis:   list[CamaraScope] = Field(default_factory=list)


class OnboardResponse(BaseModel):
    invoker_id:      str
    approval_status: str
    message:         str


class TokenRequest(BaseModel):
    client_secret: str
    scope:         str


class TokenResponse(BaseModel):
    access_token: str
    token_type:   str
    expires_in:   int
    scope:        str


class InvokerStatus(BaseModel):
    invoker_id:         str
    invoker_name:       str
    approval_status:    str
    submitted_at:       str
    approved_at:        str | None = None
    rejection_reason:   str | None = None
    scopes_approved:    list[str] = []
    keycloak_client_id: str | None = None


# ── Routes ─────────────────────────────────────────────────────────────────────

@app.get("/health")
def health():
    return {"status": "ok", "service": "invoker-onboarding", "version": "2.0.0"}


@app.get("/scopes")
def list_scopes():
    """Return the canonical CAMARA scope list this deployment supports.
    Single source of truth for the dashboard's playground + try-route
    allowlist + admin approval modal."""
    return {"scopes": list(CAMARA_SCOPES)}


@app.get("/apis", dependencies=[Depends(_require_dev_key)])
def discover_apis():
    try:
        r = _http.get(f"{CAPIF_SERVICE_URL}/catalog", timeout=10.0)
        r.raise_for_status()
        entries = r.json()
    except Exception as e:
        raise HTTPException(502, f"Cannot reach capif-service catalog: {e}")
    return {"apis": entries, "capif_core": CAPIF_CORE_URL}


@app.post("/invokers", response_model=OnboardResponse, status_code=201,
          dependencies=[Depends(_require_dev_key)])
def submit_registration(req: OnboardRequest, x_actor: str | None = Header(None)):
    """
    Submit a new invoker registration request.

    The invoker is registered with CAPIF immediately (for cert issuance) but
    starts in 'pending' status. It cannot obtain tokens until an admin approves
    the request via POST /admin/invokers/{id}/approve.
    """
    client_id     = f"client-{req.invoker_name.lower().replace(' ', '-')}-{secrets.token_hex(4)}"
    client_secret = secrets.token_urlsafe(32)
    key_pem, csr_pem = _generate_key_and_csr(client_id)

    try:
        token = _bootstrap_token()
    except Exception as e:
        raise HTTPException(502, f"Cannot reach CAPIF register: {e}")

    onboard_body = {
        "apiInvokerName":        req.invoker_name,
        "notificationDestination": req.notification_url,
        "supportedFeatures":     "0",
        "onboardingInformation": {"apiInvokerPublicKey": csr_pem},
    }

    log.info("Registering invoker %s with CAPIF (pending approval)", req.invoker_name)
    result     = _capif_post("api-invoker-management/v1/onboardedInvokers", onboard_body, token)
    invoker_id = result.get("apiInvokerId", "")
    cert_pem   = result.get("onboardingInformation", {}).get("apiInvokerCertificate", "")

    if not cert_pem:
        raise HTTPException(500, "CAPIF did not return a signed certificate")

    # Register mTLS security context.
    # Catalog lookup is best-effort: if capif-service is unreachable we skip
    # the security context (the invoker can still register and be approved;
    # they just won't have a pre-built ACL until the catalog comes back).
    catalog_entries = []
    try:
        cat_r = _http.get(f"{CAPIF_SERVICE_URL}/catalog", timeout=5.0)
        if cat_r.status_code == 200:
            catalog_entries = cat_r.json()
        else:
            log.warning("CAPIF catalog returned %d; skipping security context", cat_r.status_code)
    except Exception as e:
        log.warning("CAPIF catalog unreachable (%s); skipping security context", e)

    if catalog_entries:
        security_info = [
            {
                "apiId":                entry["service_id"],
                "aefId":               entry["aef_id"],
                "prefSecurityMethods": ["PSK"],
                "authenticationInfo":  client_id,
                "authorizationInfo":   client_secret,
            }
            for entry in catalog_entries
        ]
        with _mtls_client(cert_pem, key_pem) as mtls:
            sec_r = mtls.put(
                f"{CAPIF_CORE_URL}/capif-security/v1/trustedInvokers/{invoker_id}",
                json={
                    "securityInfo":           security_info,
                    "notificationDestination": req.notification_url,
                    "supportedFeatures":      "0",
                },
            )
            if sec_r.status_code not in (200, 201):
                log.warning("Security context registration returned %d", sec_r.status_code)

    # Persist to MongoDB — status starts as 'pending'.
    # All sensitive material lives under .secrets so the SAFE_PROJECTION
    # (defined in db.py) keeps it out of every developer-facing response
    # by default — no per-call field list to forget.
    doc = {
        "invoker_id":      invoker_id,
        "invoker_name":    req.invoker_name,
        "description":     req.description,
        "notification_url": req.notification_url,
        "approval_status": "pending",
        "submitted_at":    _now(),
        "submitted_by": {
            "email":    req.contact_email,
            "company":  req.company,
            "use_case": req.use_case,
        },
        "requested_apis":    req.requested_apis or [],
        "scopes_approved":   [],
        "approved_at":       None,
        "approved_by":       None,
        "rejection_reason":  None,
        "keycloak_client_id": None,
        "secrets": {
            "key_pem":       encrypt(key_pem),
            "cert_pem":      encrypt(cert_pem),
            "client_secret": encrypt(client_secret),
        },
        "internal": {
            "client_id": client_id,
        },
    }
    _col.insert_one(doc)
    # Actor on the submit row comes from X-Actor (set by the trusted
    # dashboard proxy from the authenticated session). If absent (direct
    # FastAPI call), fall back to anonymous — never trust req.contact_email
    # since the developer could put anything in the body.
    _audit_log("submitted", invoker_id, actor=x_actor or "anonymous",
               detail={"requested_apis": req.requested_apis})

    log.info("Invoker %s registered — awaiting admin approval", invoker_id)

    return OnboardResponse(
        invoker_id=invoker_id,
        approval_status="pending",
        message=(
            f"Registration submitted for '{req.invoker_name}'. "
            "Awaiting operator approval. Check status with GET /invokers/{invoker_id}."
        ),
    )


@app.get("/invokers/{invoker_id}/credentials",
         dependencies=[Depends(_require_dev_key)])
def get_invoker_credentials(invoker_id: str, x_actor: str | None = Header(None)):
    """Return the Keycloak client_id + secret for an approved invoker.

    Every reveal is recorded in the audit log so disclosure is never silent.
    The dashboard injects the authenticated user as `X-Actor` (developer email
    or admin email). Unset → "unknown".
    """
    doc = _col.find_one({"invoker_id": invoker_id})
    if not doc:
        raise HTTPException(404, f"Invoker {invoker_id} not found")
    if doc.get("approval_status") != "approved":
        raise HTTPException(403, "Credentials only available for approved invokers")

    encrypted = (doc.get("secrets") or {}).get("keycloak_secret")
    secret = None
    if encrypted:
        try:
            secret = decrypt(encrypted)
        except Exception:
            log.error("Could not decrypt stored secret for %s — key rotated?", invoker_id)
            raise HTTPException(503, "Stored credentials unreadable — contact operator")

    _audit_log("credentials_revealed", invoker_id, actor=x_actor or "unknown")

    return {
        "keycloak_client_id": doc.get("keycloak_client_id"),
        "keycloak_secret":    secret,
        "scopes_approved":    doc.get("scopes_approved", []),
    }


@app.get("/invokers/by-email/{email}", response_model=list[InvokerStatus],
         dependencies=[Depends(_require_dev_key)])
def list_invokers_by_email(email: str):
    """List all invokers submitted by a given developer email (oldest last)."""
    docs = list(
        _col.find({"submitted_by.email": email}, SAFE_PROJECTION).sort("submitted_at", -1)
    )
    return [
        InvokerStatus(
            invoker_id=d["invoker_id"],
            invoker_name=d["invoker_name"],
            approval_status=d["approval_status"],
            submitted_at=d["submitted_at"].isoformat(),
            approved_at=d["approved_at"].isoformat() if d.get("approved_at") else None,
            rejection_reason=d.get("rejection_reason"),
            scopes_approved=d.get("scopes_approved", []),
            keycloak_client_id=d.get("keycloak_client_id"),
        )
        for d in docs
    ]


@app.get("/invokers/{invoker_id}", response_model=InvokerStatus)
def get_invoker_status(invoker_id: str):
    """Check the approval status of a registration request."""
    doc = _col.find_one({"invoker_id": invoker_id}, SAFE_PROJECTION)
    if not doc:
        raise HTTPException(404, f"Invoker {invoker_id} not found")

    return InvokerStatus(
        invoker_id=doc["invoker_id"],
        invoker_name=doc["invoker_name"],
        approval_status=doc["approval_status"],
        submitted_at=doc["submitted_at"].isoformat(),
        approved_at=doc["approved_at"].isoformat() if doc.get("approved_at") else None,
        rejection_reason=doc.get("rejection_reason"),
        scopes_approved=doc.get("scopes_approved", []),
        keycloak_client_id=doc.get("keycloak_client_id"),
    )


@app.post("/invokers/{invoker_id}/token", response_model=TokenResponse,
          dependencies=[Depends(_require_dev_key)])
def get_token(invoker_id: str, req: TokenRequest):
    """
    Exchange client credentials for a CAPIF-issued JWT.
    Only succeeds if the invoker has been approved by an operator.

    scope format: 3gpp#<AEF_ID>:<api_name>
    Use GET /apis to discover valid AEF IDs and scopes.
    """
    doc = _col.find_one({"invoker_id": invoker_id})
    if not doc:
        raise HTTPException(404, f"Invoker {invoker_id} not found")

    if doc["approval_status"] == "pending":
        raise HTTPException(403, "Registration is pending operator approval")
    if doc["approval_status"] == "rejected":
        raise HTTPException(403, f"Registration was rejected: {doc.get('rejection_reason', '')}")
    if doc["approval_status"] == "suspended":
        raise HTTPException(403, "Invoker access has been suspended by the operator")

    try:
        cert_pem = decrypt(doc["secrets"]["cert_pem"])
        key_pem  = decrypt(doc["secrets"]["key_pem"])
    except Exception:
        log.error("Could not decrypt stored mTLS material for %s — FIELD_KEY rotated?", invoker_id)
        raise HTTPException(503, "Stored credentials unreadable — contact operator")
    with _mtls_client(cert_pem, key_pem) as mtls:
        r = mtls.post(
            f"{CAPIF_CORE_URL}/capif-security/v1/securities/{invoker_id}/token",
            data={
                "grant_type":    "client_credentials",
                "client_id":     invoker_id,
                "client_secret": req.client_secret,
                "scope":         req.scope,
            },
            headers={"Content-Type": "application/x-www-form-urlencoded"},
        )

    if r.status_code != 200:
        log.error("Token request failed %d: %s", r.status_code, r.text)
        raise HTTPException(r.status_code, detail=r.text)

    data = r.json()
    return TokenResponse(
        access_token=data["access_token"],
        token_type=data.get("token_type", "Bearer"),
        expires_in=data.get("expires_in", 3600),
        scope=data.get("scope", req.scope),
    )


@app.delete("/invokers/{invoker_id}", status_code=204,
            dependencies=[Depends(_require_admin_key)])
def offboard_invoker(invoker_id: str, x_actor: str | None = Header(None)):
    """Remove an invoker from CAPIF and the portal. Admin key required."""
    doc = _col.find_one({"invoker_id": invoker_id})
    if not doc:
        raise HTTPException(404, f"Invoker {invoker_id} not found")

    try:
        token = _bootstrap_token()
    except Exception as e:
        raise HTTPException(502, f"Cannot reach CAPIF register: {e}")

    r = _http.delete(
        f"{CAPIF_CORE_URL}/api-invoker-management/v1/onboardedInvokers/{invoker_id}",
        headers={"Authorization": f"Bearer {token}"},
    )
    if r.status_code not in (200, 204):
        raise HTTPException(r.status_code, r.text)

    _col.delete_one({"invoker_id": invoker_id})
    _audit_log("offboarded", invoker_id, actor=x_actor or "admin")


app.include_router(admin_router, prefix="/admin", tags=["Admin"])
