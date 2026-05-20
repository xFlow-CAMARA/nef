"""Smoke-test the admin API-key gate using a stripped-down FastAPI app
with an in-memory mongomock so we don't need a running Mongo."""

import os
import sys
import pathlib
import importlib

sys.path.insert(0, str(pathlib.Path(__file__).resolve().parent.parent))

os.environ["APP_ENV"]                 = "dev"
os.environ["FIELD_KEY_PASSPHRASE"]    = "unit-test-key"
os.environ["INVOKER_ADMIN_API_KEY"]   = "test-secret"

import mongomock                          # noqa: E402
import db as db_mod                       # noqa: E402

# Swap the real MongoClient connection for an in-memory one *before*
# anything reads the collection handles.
_fake = mongomock.MongoClient()
db_mod._db        = _fake.get_database("camara")
db_mod.invokers   = db_mod._db["invokers"]
db_mod.audit_logs = db_mod._db["audit_logs"]

# Re-import admin_router so it picks up the swapped collections
import admin_router                       # noqa: E402
importlib.reload(admin_router)

from fastapi import FastAPI               # noqa: E402
from fastapi.testclient import TestClient # noqa: E402

app = FastAPI()
app.include_router(admin_router.router, prefix="/admin")
client = TestClient(app)


def test_unauthenticated_request_is_rejected():
    r = client.get("/admin/invokers")
    assert r.status_code == 401


def test_wrong_key_is_rejected():
    r = client.get("/admin/invokers", headers={"X-Admin-Api-Key": "nope"})
    assert r.status_code == 401


def test_correct_key_succeeds():
    r = client.get("/admin/invokers", headers={"X-Admin-Api-Key": "test-secret"})
    assert r.status_code == 200
    assert r.json() == []      # empty mongo
