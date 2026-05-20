"""Admin API-key gate behaviour.

Thanks to the call-time env read in `_require_admin_key`, we don't need
`importlib.reload`. `monkeypatch.setenv` is enough.
"""

import pytest
from fastapi import FastAPI
from fastapi.testclient import TestClient

from admin_router import router as admin_router

app = FastAPI()
app.include_router(admin_router, prefix="/admin")
client = TestClient(app)


def test_dev_mode_passes_through_without_key(monkeypatch):
    monkeypatch.delenv("INVOKER_ADMIN_API_KEY", raising=False)
    r = client.get("/admin/invokers")
    assert r.status_code == 200


def test_missing_key_is_rejected_when_enabled(monkeypatch):
    monkeypatch.setenv("INVOKER_ADMIN_API_KEY", "expected")
    r = client.get("/admin/invokers")
    assert r.status_code == 401


def test_wrong_key_is_rejected(monkeypatch):
    monkeypatch.setenv("INVOKER_ADMIN_API_KEY", "expected")
    r = client.get("/admin/invokers", headers={"X-Admin-Api-Key": "wrong"})
    assert r.status_code == 401


def test_correct_key_succeeds(monkeypatch):
    monkeypatch.setenv("INVOKER_ADMIN_API_KEY", "expected")
    r = client.get("/admin/invokers", headers={"X-Admin-Api-Key": "expected"})
    assert r.status_code == 200
    assert r.json() == []
