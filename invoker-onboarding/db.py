"""Shared MongoDB connection + audit + field-level encryption helpers.

All sensitive fields live in `invokers[].secrets` and are individually
encrypted at rest with Fernet. SAFE_PROJECTION still hides them from
casual reads; encrypt/decrypt wrap explicit accesses.
"""

import logging
import os
import base64
import hashlib
from datetime import datetime, timezone

from cryptography.fernet import Fernet, InvalidToken
from pymongo import MongoClient

log = logging.getLogger("invoker-onboarding.db")

APP_ENV = os.getenv("APP_ENV", "dev")
MONGODB_URI = os.getenv("MONGODB_URI", "mongodb://camara-mongodb:27017/camara")


def _build_fernet() -> Fernet:
    """
    Build a Fernet from FIELD_KEY (raw key) or FIELD_KEY_PASSPHRASE
    (derived via SHA-256 → URL-safe b64). In dev, an unset key derives
    a stable demo key so existing data survives container restarts.
    """
    raw = os.getenv("FIELD_KEY")
    if raw:
        return Fernet(raw.encode())
    passphrase = os.getenv("FIELD_KEY_PASSPHRASE")
    if passphrase:
        digest = hashlib.sha256(passphrase.encode()).digest()
        return Fernet(base64.urlsafe_b64encode(digest))
    if APP_ENV == "dev":
        log.warning("FIELD_KEY unset — using insecure dev key derived from a fixed passphrase")
        digest = hashlib.sha256(b"xflow-camara-demo-field-key").digest()
        return Fernet(base64.urlsafe_b64encode(digest))
    raise RuntimeError("FIELD_KEY (or FIELD_KEY_PASSPHRASE) env var required when APP_ENV != 'dev'")


_fernet = _build_fernet()


def encrypt(value: str | None) -> str | None:
    """Encrypt a UTF-8 string for storage. Returns None if input is None."""
    if value is None:
        return None
    return _fernet.encrypt(value.encode()).decode()


def decrypt(token: str | None) -> str | None:
    """Decrypt a stored string. Returns the raw string if it isn't a valid token
    (lets old plaintext records keep working after deployment)."""
    if not token:
        return token
    try:
        return _fernet.decrypt(token.encode()).decode()
    except (InvalidToken, ValueError):
        return token


_mongo  = MongoClient(MONGODB_URI)
_db     = _mongo.get_default_database()

invokers   = _db["invokers"]
audit_logs = _db["audit_logs"]

# Default projection — hides anything inside .secrets from public responses.
SAFE_PROJECTION = {"secrets": 0}


def now() -> datetime:
    return datetime.now(timezone.utc)


def audit(action: str, invoker_id: str, actor: str = "system", detail: dict | None = None) -> None:
    """Insert one immutable audit row."""
    audit_logs.insert_one({
        "action":     action,
        "invoker_id": invoker_id,
        "actor":      actor,
        "timestamp":  now(),
        "detail":     detail or {},
    })
