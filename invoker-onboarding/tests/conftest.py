"""Shared pytest fixtures.

Two responsibilities:
  1. Make the `invoker-onboarding/` package importable from anywhere.
  2. Swap the real Mongo client for `mongomock` once, before any test
     module imports `db` or `admin_router`.
"""

import os
import sys
import pathlib

# Make modules importable
sys.path.insert(0, str(pathlib.Path(__file__).resolve().parent.parent))

# Set a stable dev field key before any module reads it
os.environ.setdefault("APP_ENV", "dev")
os.environ.setdefault("FIELD_KEY_PASSPHRASE", "unit-test-key")

import mongomock                                # noqa: E402
import pytest                                   # noqa: E402

import db                                       # noqa: E402

# Replace the real MongoClient connection with an in-memory one
_fake_client = mongomock.MongoClient()
db._db        = _fake_client.get_database("camara")
db.invokers   = db._db["invokers"]
db.audit_logs = db._db["audit_logs"]


@pytest.fixture(autouse=True)
def _reset_mongo():
    """Empty collections before each test so order doesn't matter."""
    db.invokers.delete_many({})
    db.audit_logs.delete_many({})
    yield
