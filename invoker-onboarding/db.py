"""Shared MongoDB connection + audit + field-level encryption helpers.

All sensitive fields live in `invokers[].secrets` and are individually
encrypted at rest with Fernet. SAFE_PROJECTION still hides them from
casual reads; encrypt/decrypt wrap explicit accesses.
"""

import base64
import hashlib
import logging
import os
from datetime import UTC, datetime
from functools import lru_cache

from cryptography.fernet import Fernet, InvalidToken
from pymongo import MongoClient

from config import app_env

log = logging.getLogger("invoker-onboarding.db")

MONGODB_URI = os.getenv("MONGODB_URI", "mongodb://camara-mongodb:27017/camara")


@lru_cache(maxsize=1)
def _fernet() -> Fernet:
    """
    Build a Fernet client. Cached so we don't re-derive on every call.

    Order of precedence:
      1. FIELD_KEY            — raw Fernet key (44 url-safe base64 chars)
      2. FIELD_KEY_PASSPHRASE — derived via SHA-256
      3. APP_ENV=dev only     — fixed insecure demo passphrase

    Production deployments MUST set one of the first two.
    """
    raw = os.getenv("FIELD_KEY")
    if raw:
        return Fernet(raw.encode())
    passphrase = os.getenv("FIELD_KEY_PASSPHRASE")
    if passphrase:
        digest = hashlib.sha256(passphrase.encode()).digest()
        return Fernet(base64.urlsafe_b64encode(digest))
    if app_env() == "dev":
        log.warning("FIELD_KEY unset — using insecure dev key derived from a fixed passphrase")
        digest = hashlib.sha256(b"xflow-camara-demo-field-key").digest()
        return Fernet(base64.urlsafe_b64encode(digest))
    raise RuntimeError("FIELD_KEY (or FIELD_KEY_PASSPHRASE) env var required when APP_ENV != 'dev'")


def encrypt(value: str | None) -> str | None:
    """Encrypt a UTF-8 string for storage. Returns None if input is None."""
    if value is None:
        return None
    return _fernet().encrypt(value.encode()).decode()


def decrypt(token: str | None, *, allow_plaintext_fallback: bool = False) -> str | None:
    """Decrypt a Fernet-encrypted string.

    Raises InvalidToken on failure unless `allow_plaintext_fallback=True`, in
    which case strings that don't look like Fernet tokens (no `gAAAAA` prefix)
    are returned as-is — *only* useful during a one-time migration off plaintext.
    """
    if not token:
        return token
    try:
        return _fernet().decrypt(token.encode()).decode()
    except (InvalidToken, ValueError):
        if allow_plaintext_fallback and not token.startswith("gAAAAA"):
            return token
        raise


_mongo  = MongoClient(MONGODB_URI)
_db     = _mongo.get_default_database()

invokers   = _db["invokers"]
audit_logs = _db["audit_logs"]

# Two-level visibility scheme:
#   .secrets   — real secrets, encrypted at rest, never returned by default
#   .internal  — operational metadata (uuids, flags) — useful to admins,
#                still hidden from developer-facing responses
SAFE_PROJECTION  = {"secrets": 0, "internal": 0}   # developer-facing reads
ADMIN_PROJECTION = {"secrets": 0}                  # admin-facing reads (internal visible)


def now() -> datetime:
    return datetime.now(UTC)


def audit(action: str, invoker_id: str, actor: str = "system", detail: dict | None = None) -> None:
    """Insert one immutable audit row."""
    audit_logs.insert_one({
        "action":     action,
        "invoker_id": invoker_id,
        "actor":      actor,
        "timestamp":  now(),
        "detail":     detail or {},
    })
