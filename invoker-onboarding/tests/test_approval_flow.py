"""Approve / reject / revoke state machine + token gating + secret encryption.

External services (Keycloak admin API, Redis ACL pub/sub) are stubbed at the
module level so the tests run without any other process — pure pytest in CI.
"""

from datetime import UTC, datetime
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
        "submitted_at":    datetime.now(UTC),
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
    body = r.json()
    assert {row["invoker_id"] for row in body["items"]} == {"INV-list-1"}
    assert body["has_more"] is False


def test_audit_endpoint_returns_recent_events():
    _seed_pending("INV-audit")
    db.audit_logs.insert_one({
        "action":     "approved",
        "invoker_id": "INV-audit",
        "actor":      "tester",
        "timestamp":  datetime.now(UTC),
        "detail":     {"scopes_approved": ["sim-swap"]},
    })
    r = client.get("/admin/audit?invoker_id=INV-audit")
    assert r.status_code == 200
    assert r.json()[0]["action"] == "approved"


def test_approve_rolls_back_keycloak_when_mongo_update_fails(monkeypatch):
    """If the final Mongo write fails after Keycloak + ACL succeeded, the
    Keycloak client and ACL entries MUST be torn down — otherwise the next
    retry collides on 'client already exists' and we have orphan state."""
    _seed_pending("INV-rollback")

    kc_del = patch("admin_router.delete_keycloak_client").start()
    acl_rm = patch("admin_router.remove_acl_entries").start()
    patch("admin_router.create_keycloak_client", return_value={
        "client_id": "INV-rollback", "client_secret": "x", "keycloak_uuid": "uuid-rb",
    }).start()
    patch("admin_router.create_acl_entries", return_value=["sim-swap"]).start()

    # Force the Mongo update_one (the FINAL one — not the CAS claim) to raise.
    original_update_one = db.invokers.update_one
    call_count = {"n": 0}

    def flaky_update(*args, **kwargs):
        call_count["n"] += 1
        # The CAS claim (first update_one) happens via find_one_and_update,
        # so this only intercepts the final state-persist update_one.
        if call_count["n"] == 1:
            raise RuntimeError("simulated mongo outage")
        return original_update_one(*args, **kwargs)

    monkeypatch.setattr(db.invokers, "update_one", flaky_update)

    try:
        r = client.post("/admin/invokers/INV-rollback/approve",
                        json={"scopes_approved": ["sim-swap"], "approved_by": "tester"})
    finally:
        patch.stopall()

    assert r.status_code == 503
    kc_del.assert_called_once_with("uuid-rb")
    acl_rm.assert_called_once_with("INV-rollback")
    # Transient 'approving' state must be reverted to 'pending' so retry works.
    doc = db.invokers.find_one({"invoker_id": "INV-rollback"})
    assert doc["approval_status"] == "pending"


def test_rotate_secret_replaces_keycloak_secret_and_audits():
    """Rotating swaps the encrypted secret in Mongo and logs a 'secret_rotated'
    audit row. The OLD encrypted value must no longer be present."""
    _seed_pending("INV-rot")
    db.invokers.update_one(
        {"invoker_id": "INV-rot"},
        {"$set": {
            "approval_status":         "approved",
            "internal.keycloak_uuid":  "uuid-rot",
            "secrets.keycloak_secret": db.encrypt("OLD-SECRET"),
        }},
    )
    with patch("admin_router.rotate_keycloak_client_secret", return_value="NEW-SECRET"):
        r = client.post("/admin/invokers/INV-rot/rotate-secret",
                        json={"reason": "dev lost it", "rotated_by": "ops@x"})

    assert r.status_code == 200
    body = r.json()
    assert body["keycloak_secret"] == "NEW-SECRET"

    doc = db.invokers.find_one({"invoker_id": "INV-rot"})
    assert db.decrypt(doc["secrets"]["keycloak_secret"]) == "NEW-SECRET"

    audit = db.audit_logs.find_one({"invoker_id": "INV-rot", "action": "secret_rotated"})
    assert audit["actor"] == "ops@x"
    assert audit["detail"]["reason"] == "dev lost it"


def test_rotate_secret_only_works_on_approved():
    """Pending or suspended invokers don't have a secret to rotate."""
    _seed_pending("INV-rot-pending")
    r = client.post("/admin/invokers/INV-rot-pending/rotate-secret",
                    json={"rotated_by": "ops@x"})
    assert r.status_code == 409


def test_concurrent_approve_returns_409():
    """The CAS claim ensures only one approve attempt can proceed at a time."""
    _seed_pending("INV-conc")
    # Simulate a second admin already grabbed the claim.
    db.invokers.update_one({"invoker_id": "INV-conc"}, {"$set": {"approval_status": "approving"}})

    r = client.post("/admin/invokers/INV-conc/approve",
                    json={"scopes_approved": ["sim-swap"], "approved_by": "x"})
    assert r.status_code == 409
    assert "currently being approved" in r.json()["detail"]
