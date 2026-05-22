#!/usr/bin/env python3
"""Re-encrypt every encrypted field in the invokers collection from an OLD
FIELD_KEY (or passphrase) to a NEW one. Run with the service STOPPED to
avoid races against live writes.

Usage:
    OLD_FIELD_KEY_PASSPHRASE=old-key \
    NEW_FIELD_KEY_PASSPHRASE=new-key \
    MONGODB_URI=mongodb://camara-mongodb:27017/camara \
    python3 scripts/rotate-field-key.py [--apply]

Defaults to dry-run mode (prints what would change, writes nothing).
Pass --apply to actually commit the migration. Writes happen in batches
of 500 via Mongo bulk_write. A `field_key_rotated_started` and
`field_key_rotated_finished` row is inserted into audit_logs so the
event is visible to anyone tailing the audit table.
"""

import argparse
import base64
import hashlib
import os
import sys
from datetime import UTC, datetime

from cryptography.fernet import Fernet, InvalidToken
from pymongo import MongoClient, UpdateOne


def _build_fernet(passphrase_var: str, raw_var: str) -> Fernet:
    raw = os.getenv(raw_var)
    if raw:
        return Fernet(raw.encode())
    passphrase = os.getenv(passphrase_var)
    if not passphrase:
        sys.exit(f"Set ${raw_var} (raw 44-char Fernet key) or ${passphrase_var}")
    digest = hashlib.sha256(passphrase.encode()).digest()
    return Fernet(base64.urlsafe_b64encode(digest))


def _now():
    return datetime.now(UTC)


def _audit(audit_logs, action: str, detail: dict | None = None) -> None:
    """Stand-alone audit insert (the runtime db.py module isn't importable
    from this script — script runs in its own venv outside the service)."""
    audit_logs.insert_one({
        "action":     action,
        "invoker_id": "(field-key-rotation)",
        "actor":      os.getenv("USER") or "operator",
        "timestamp":  _now(),
        "detail":     detail or {},
    })


def main() -> None:
    ap = argparse.ArgumentParser(description=__doc__)
    ap.add_argument("--apply", action="store_true",
                    help="actually write the migration (default: dry-run, no writes)")
    ap.add_argument("--batch", type=int, default=500, help="bulk_write batch size")
    args = ap.parse_args()

    dry_run = not args.apply

    old = _build_fernet("OLD_FIELD_KEY_PASSPHRASE", "OLD_FIELD_KEY")
    new = _build_fernet("NEW_FIELD_KEY_PASSPHRASE", "NEW_FIELD_KEY")
    if old.encrypt(b"x") == new.encrypt(b"x"):
        sys.exit("OLD and NEW keys derive to the same Fernet — nothing to do")

    mongo_uri = os.getenv("MONGODB_URI", "mongodb://camara-mongodb:27017/camara")
    db = MongoClient(mongo_uri).get_default_database()
    invokers   = db["invokers"]
    audit_logs = db["audit_logs"]

    FIELDS = ["key_pem", "cert_pem", "client_secret", "keycloak_secret"]

    if not dry_run:
        _audit(audit_logs, "field_key_rotated_started", {"batch_size": args.batch})

    total = invokers.count_documents({})
    migrated = 0
    skipped  = 0
    failed   = 0
    pending: list[UpdateOne] = []

    print(f"[{'DRY-RUN' if dry_run else 'APPLY'}] scanning {total} invokers …\n")

    for doc in invokers.find({}):
        secrets = doc.get("secrets") or {}
        updates: dict[str, str] = {}
        for f in FIELDS:
            v = secrets.get(f)
            if not v:
                continue
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
            migrated += 1
            print(f"  ✓ {doc['invoker_id']}  ({len(updates)} field(s))")
            if not dry_run:
                pending.append(UpdateOne({"_id": doc["_id"]}, {"$set": updates}))
                if len(pending) >= args.batch:
                    invokers.bulk_write(pending, ordered=False)
                    pending.clear()
        else:
            skipped += 1

    if pending:
        invokers.bulk_write(pending, ordered=False)

    summary = {"total": total, "migrated": migrated, "already_new": skipped, "failed": failed}
    print(f"\n[{'DRY-RUN' if dry_run else 'APPLY'}] {summary}")

    if not dry_run:
        _audit(audit_logs, "field_key_rotated_finished", summary)

    if dry_run:
        print("\nNothing was written. Re-run with --apply to commit.")
    if failed:
        sys.exit(1)


if __name__ == "__main__":
    main()
