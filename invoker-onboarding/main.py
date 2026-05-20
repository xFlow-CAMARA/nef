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
import datetime
import secrets
import tempfile

import httpx
from fastapi import FastAPI, HTTPException
from fastapi.middleware.cors import CORSMiddleware
from pydantic import BaseModel
from typing import Dict, List, Optional
from pymongo import MongoClient
from cryptography import x509
from cryptography.x509.oid import NameOID
from cryptography.hazmat.primitives import hashes, serialization
from cryptography.hazmat.primitives.asymmetric import rsa

logging.basicConfig(level=logging.INFO)
log = logging.getLogger("invoker-onboarding")

CAPIF_CORE_URL     = os.getenv("CAPIF_CORE_URL",     "https://capifcore:443")
CAPIF_REGISTER_URL = os.getenv("CAPIF_REGISTER_URL", "https://register:8080")
CAPIF_USERNAME     = os.getenv("CAPIF_USERNAME",     "nef-xflow")
CAPIF_PASSWORD     = os.getenv("CAPIF_PASSWORD",     "xflow-nef-2026")
CAPIF_SERVICE_URL  = os.getenv("CAPIF_SERVICE_URL",  "http://capif-service:8080")
MONGODB_URI        = os.getenv("MONGODB_URI",        "mongodb://camara-mongodb:27017/camara")

_http = httpx.Client(verify=False, timeout=15.0)

# ── MongoDB ────────────────────────────────────────────────────────────────────
_mongo  = MongoClient(MONGODB_URI)
_db     = _mongo.get_default_database()
_col    = _db["invokers"]      # developer portal collection
_audit  = _db["audit_logs"]


def _now() -> datetime.datetime:
    return datetime.datetime.utcnow()


def _audit_log(action: str, invoker_id: str, actor: str = "system", detail: dict = None):
    _audit.insert_one({
        "action":     action,
        "invoker_id": invoker_id,
        "actor":      actor,
        "timestamp":  _now(),
        "detail":     detail or {},
    })


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
        log.error("CAPIF %s → %d: %s", path, r.status_code, r.text)
        raise HTTPException(status_code=r.status_code, detail=r.text)
    return r.json()


def _mtls_client(cert_pem: str, key_pem: str) -> httpx.Client:
    with tempfile.NamedTemporaryFile(suffix=".pem", delete=False, mode="w") as cf:
        cf.write(cert_pem)
        cert_path = cf.name
    with tempfile.NamedTemporaryFile(suffix=".key", delete=False, mode="w") as kf:
        kf.write(key_pem)
        key_path = kf.name
    return httpx.Client(cert=(cert_path, key_path), verify=False, timeout=15.0)


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
    invoker_name:     str
    description:      Optional[str] = ""
    notification_url: Optional[str] = "http://localhost/capif-callback"
    # Developer identity — stored for admin review
    contact_email:    Optional[str] = ""
    company:          Optional[str] = ""
    use_case:         Optional[str] = ""
    requested_apis:   Optional[List[str]] = []   # e.g. ["quality-on-demand", "location-retrieval"]


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
    invoker_id:      str
    invoker_name:    str
    approval_status: str
    submitted_at:    str
    approved_at:     Optional[str] = None
    rejection_reason: Optional[str] = None
    scopes_approved: Optional[List[str]] = None
    keycloak_client_id: Optional[str] = None


# ── Routes ─────────────────────────────────────────────────────────────────────

@app.get("/health")
def health():
    return {"status": "ok", "service": "invoker-onboarding", "version": "2.0.0"}


@app.get("/apis")
def discover_apis():
    try:
        r = _http.get(f"{CAPIF_SERVICE_URL}/catalog", timeout=10.0)
        r.raise_for_status()
        entries = r.json()
    except Exception as e:
        raise HTTPException(502, f"Cannot reach capif-service catalog: {e}")
    return {"apis": entries, "capif_core": CAPIF_CORE_URL}


@app.post("/invokers", response_model=OnboardResponse, status_code=201)
def submit_registration(req: OnboardRequest):
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

    # Register mTLS security context
    catalog_entries = []
    try:
        cat_r = _http.get(f"{CAPIF_SERVICE_URL}/catalog", timeout=5.0)
        if cat_r.status_code == 200:
            catalog_entries = cat_r.json()
    except Exception:
        pass

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
        mtls = _mtls_client(cert_pem, key_pem)
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

    # Persist to MongoDB — status starts as 'pending'
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
        # CAPIF mTLS credentials (needed for token requests)
        "_key_pem":          key_pem,
        "_cert_pem":         cert_pem,
        "_client_id":        client_id,
        "_client_secret":    client_secret,
    }
    _col.insert_one(doc)
    _audit_log("submitted", invoker_id, actor=req.contact_email or "anonymous",
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


@app.get("/invokers/{invoker_id}/credentials")
def get_invoker_credentials(invoker_id: str):
    """Return the Keycloak client_id + secret for an approved invoker.

    Used by the developer portal so developers don't need to save the secret themselves.
    """
    doc = _col.find_one({"invoker_id": invoker_id})
    if not doc:
        raise HTTPException(404, f"Invoker {invoker_id} not found")
    if doc.get("approval_status") != "approved":
        raise HTTPException(403, "Credentials only available for approved invokers")

    return {
        "keycloak_client_id": doc.get("keycloak_client_id"),
        "keycloak_secret":    doc.get("_keycloak_secret"),
        "scopes_approved":    doc.get("scopes_approved", []),
    }


@app.get("/invokers/by-email/{email}", response_model=List[InvokerStatus])
def list_invokers_by_email(email: str):
    """List all invokers submitted by a given developer email (oldest last)."""
    docs = list(
        _col.find(
            {"submitted_by.email": email},
            {"_key_pem": 0, "_cert_pem": 0, "_client_secret": 0, "_keycloak_secret": 0},
        ).sort("submitted_at", -1)
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
    doc = _col.find_one({"invoker_id": invoker_id}, {"_key_pem": 0, "_cert_pem": 0, "_client_secret": 0})
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


@app.post("/invokers/{invoker_id}/token", response_model=TokenResponse)
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

    mtls = _mtls_client(doc["_cert_pem"], doc["_key_pem"])
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


@app.delete("/invokers/{invoker_id}", status_code=204)
def offboard_invoker(invoker_id: str):
    """Remove an invoker from CAPIF and the portal."""
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
    _audit_log("offboarded", invoker_id)


# Mount admin router
from admin_router import router as admin_router
app.include_router(admin_router, prefix="/admin", tags=["Admin"])
