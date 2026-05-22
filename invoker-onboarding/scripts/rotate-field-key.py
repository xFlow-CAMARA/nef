#!/usr/bin/env python3
"""Re-encrypt every encrypted field in the invokers collection from an OLD
FIELD_KEY (or passphrase) to a NEW one. Run with the service STOPPED to
avoid races against live writes.

Usage:
    OLD_FIELD_KEY_PASSPHRASE=old-key \
    NEW_FIELD_KEY_PASSPHRASE=new-key \
    MONGODB_URI=mongodb://camara-mongodb:27017/camara \
    python3 scripts/rotate-field-key.py

After it finishes, set FIELD_KEY_PASSPHRASE=<new-value> in your service env
and restart. The script is idempotent — re-running with the same OLD/NEW
keys is a no-op on documents already migrated.
"""

import base64
import hashlib
import os
import sys

from cryptography.fernet import Fernet, InvalidToken
from pymongo import MongoClient


def _build_fernet(passphrase_var: str, raw_var: str) -> Fernet:
    raw = os.getenv(raw_var)
    if raw:
        return Fernet(raw.encode())
    passphrase = os.getenv(passphrase_var)
    if not passphrase:
        sys.exit(f"Set ${raw_var} (raw 44-char Fernet key) or ${passphrase_var}")
    digest = hashlib.sha256(passphrase.encode()).digest()
    return Fernet(base64.urlsafe_b64encode(digest))


def main() -> None:
    old = _build_fernet("OLD_FIELD_KEY_PASSPHRASE", "OLD_FIELD_KEY")
    new = _build_fernet("NEW_FIELD_KEY_PASSPHRASE", "NEW_FIELD_KEY")
    if old.encrypt(b"x") == new.encrypt(b"x"):
        sys.exit("OLD and NEW keys derive to the same Fernet — nothing to do")

    mongo_uri = os.getenv("MONGODB_URI", "mongodb://camara-mongodb:27017/camara")
    db = MongoClient(mongo_uri).get_default_database()
    invokers = db["invokers"]

    # The encrypted fields under .secrets we know about.
    FIELDS = ["key_pem", "cert_pem", "client_secret", "keycloak_secret"]

    total = invokers.count_documents({})
    migrated = 0
    skipped  = 0
    failed   = 0

    for doc in invokers.find({}):
        secrets = doc.get("secrets") or {}
        updates: dict[str, str] = {}
        for f in FIELDS:
            v = secrets.get(f)
            if not v:
                continue
            # Try old → new. If it already decrypts with new, leave it alone.
            try:
                new.decrypt(v.encode())
                continue                    # already on new key
            except InvalidToken:
                pass
            try:
                plain = old.decrypt(v.encode())
            except InvalidToken:
                print(f"  ! could not decrypt secrets.{f} on {doc['invoker_id']} with EITHER key — skipping")
                failed += 1
                continue
            updates[f"secrets.{f}"] = new.encrypt(plain).decode()

        if updates:
            invokers.update_one({"_id": doc["_id"]}, {"$set": updates})
            migrated += 1
            print(f"  ✓ {doc['invoker_id']}  ({len(updates)} field(s))")
        else:
            skipped += 1

    print(f"\nDone. total={total}  migrated={migrated}  already-new={skipped}  failed={failed}")
    if failed:
        sys.exit(1)


if __name__ == "__main__":
    main()
