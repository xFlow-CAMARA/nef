"""Shared pytest fixtures.

Mongo swap happens at *conftest import time* so any test-module-level
import of `db` / `admin_router` / `main` picks up the in-memory client
before binding their own references.
"""

import os
import pathlib
import sys

import mongomock
import pytest

# Make modules importable
sys.path.insert(0, str(pathlib.Path(__file__).resolve().parent.parent))

# Test env vars must be set before importing `db` (which reads FIELD_KEY at
# first encrypt/decrypt call via @lru_cache, but APP_ENV is read at module
# load). Use os.environ.setdefault so a CI runner that pre-sets these wins.
os.environ.setdefault("APP_ENV", "dev")
os.environ.setdefault("FIELD_KEY_PASSPHRASE", "unit-test-key")

# Pin upstream URLs to known values so respx mocks match regardless of where
# the tests run (local docker, GitHub Actions, etc.). Use forced assignment
# rather than setdefault — the docker image may bake in CAPIF_CORE_URL via
# docker run -e and we don't want that leaking into test expectations.
os.environ["CAPIF_CORE_URL"]     = "https://test-capif"
os.environ["CAPIF_REGISTER_URL"] = "https://test-register"
os.environ["CAPIF_SERVICE_URL"]  = "http://test-capif-service"

import db  # noqa: E402

# In-memory Mongo for the whole test session. Must happen BEFORE any test
# module imports admin_router / main (those bind `db.invokers` at import).
_fake = mongomock.MongoClient()
db._db        = _fake.get_database("camara")
db.invokers   = db._db["invokers"]
db.audit_logs = db._db["audit_logs"]


@pytest.fixture(autouse=True)
def _reset_mongo():
    """Empty collections before each test so order doesn't matter."""
    db.invokers.delete_many({})
    db.audit_logs.delete_many({})
    yield
