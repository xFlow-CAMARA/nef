"""Approve / reject / revoke state machine + token gating + secret encryption.

External services (Keycloak admin API, Redis ACL pub/sub) are stubbed at the
module level so the tests run without any other process — pure pytest in CI.
"""

from datetime import datetime, timezone
from unittest.mock import patch

from fastapi import FastAPI
from fastapi.testclient import TestClient

import db
from admin_router import router as admin_router

app = FastAPI()
app.include_router(admin_router, prefix="/admin")
client = TestClient(app)


def _seed_pending(invoker_id: str = "INV-test-1", email: str = "dev@example.com") -> None:
    db.invokers.insert_one({
        "invoker_id":      invoker_id,
        "invoker_name":    "test-app",
        "approval_status": "pending",
        "submitted_at":    datetime.now(timezone.utc),
        "submitted_by":    {"email": email, "company": "Acme", "use_case": "demo"},
        "requested_apis":  ["sim-swap", "quality-on-demand"],
        "scopes_approved": [],
        "secrets":  {"key_pem": "enc", "cert_pem": "enc", "client_secret": "enc"},
        "internal": {"client_id": invoker_id},
    })


def test_approve_moves_pending_to_approved_and_encrypts_secret():
    _seed_pending("INV-approve-test")
    with patch("admin_router.create_keycloak_client", return_value={
        "client_id":     "INV-approve-test",
        "client_secret": "supersecret",
        "keycloak_uuid": "kc-uuid-123",
    }), patch("admin_router.create_acl_entries", return_value=["sim-swap"]):
        r = client.post("/admin/invokers/INV-approve-test/approve", json={
            "scopes_approved": ["sim-swap"], "approved_by": "test-admin",
        })
    assert r.status_code == 200
    body = r.json()
    assert body["approval_status"]      == "approved"
    assert body["keycloak_client_id"]   == "INV-approve-test"

    doc = db.invokers.find_one({"invoker_id": "INV-approve-test"})
    assert doc["approval_status"] == "approved"
    assert doc["approved_by"]     == "test-admin"
    # Real secret encrypted at rest, not plaintext
    assert doc["secrets"]["keycloak_secret"].startswith("gAAAAA")
    # Operational metadata in the right subdocument
    assert doc["internal"]["keycloak_uuid"]  == "kc-uuid-123"
    assert doc["internal"]["acl_published"]  == ["sim-swap"]


def test_double_approve_returns_409():
    _seed_pending("INV-dup")
    db.invokers.update_one({"invoker_id": "INV-dup"}, {"$set": {"approval_status": "approved"}})
    r = client.post("/admin/invokers/INV-dup/approve",
                    json={"scopes_approved": ["sim-swap"], "approved_by": "x"})
    assert r.status_code == 409


def test_reject_only_works_on_pending():
    _seed_pending("INV-reject")
    r = client.post("/admin/invokers/INV-reject/reject",
                    json={"rejection_reason": "spam", "rejected_by": "x"})
    assert r.status_code == 200
    # Can't reject again — no longer pending
    r2 = client.post("/admin/invokers/INV-reject/reject",
                     json={"rejection_reason": "more spam", "rejected_by": "x"})
    assert r2.status_code == 409


def test_revoke_only_works_on_approved_and_clears_internal():
    _seed_pending("INV-revoke")
    db.invokers.update_one(
        {"invoker_id": "INV-revoke"},
        {"$set": {
            "approval_status":         "approved",
            "internal.keycloak_uuid":  "kc-rev",
            "secrets.keycloak_secret": db.encrypt("topsecret"),
        }},
    )
    with patch("admin_router.delete_keycloak_client") as kc_del, \
         patch("admin_router.remove_acl_entries") as acl_rm:
        r = client.post("/admin/invokers/INV-revoke/revoke",
                        json={"reason": "abuse", "revoked_by": "x"})
        kc_del.assert_called_once_with("kc-rev")
        acl_rm.assert_called_once_with("INV-revoke")

    assert r.status_code == 200
    doc = db.invokers.find_one({"invoker_id": "INV-revoke"})
    assert doc["approval_status"]              == "suspended"
    assert doc["secrets"]["keycloak_secret"]   is None
    assert doc["internal"]["keycloak_uuid"]    is None

    # Revoke on already-suspended → 409
    r2 = client.post("/admin/invokers/INV-revoke/revoke", json={"reason": "x", "revoked_by": "x"})
    assert r2.status_code == 409


def test_list_filter_by_status():
    _seed_pending("INV-list-1", email="a@x")
    _seed_pending("INV-list-2", email="b@x")
    db.invokers.update_one({"invoker_id": "INV-list-1"}, {"$set": {"approval_status": "approved"}})

    r = client.get("/admin/invokers?status=approved")
    assert r.status_code == 200
    rows = r.json()
    assert {row["invoker_id"] for row in rows} == {"INV-list-1"}


def test_audit_endpoint_returns_recent_events():
    _seed_pending("INV-audit")
    db.audit_logs.insert_one({
        "action":     "approved",
        "invoker_id": "INV-audit",
        "actor":      "tester",
        "timestamp":  datetime.now(timezone.utc),
        "detail":     {"scopes_approved": ["sim-swap"]},
    })
    r = client.get("/admin/audit?invoker_id=INV-audit")
    assert r.status_code == 200
    assert r.json()[0]["action"] == "approved"
