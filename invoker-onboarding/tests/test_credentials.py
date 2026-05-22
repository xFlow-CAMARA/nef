"""GET /invokers/{id}/credentials — reveal flow + previous_reveal field."""

from datetime import UTC, datetime

from fastapi.testclient import TestClient

import db
from main import app

client = TestClient(app)


def _seed_approved(invoker_id: str = "INV-cred-test") -> None:
    db.invokers.insert_one({
        "invoker_id":         invoker_id,
        "invoker_name":       "test-app",
        "approval_status":    "approved",
        "submitted_at":       datetime.now(UTC),
        "submitted_by":       {"email": "dev@example.com"},
        "scopes_approved":    ["sim-swap"],
        "keycloak_client_id": invoker_id,
        "secrets":  {"keycloak_secret": db.encrypt("THE-SECRET")},
        "internal": {"client_id": invoker_id, "keycloak_uuid": "uuid-1"},
    })


def test_first_reveal_has_no_previous():
    """The first time anyone calls /credentials there's no prior reveal."""
    _seed_approved("INV-first")
    r = client.get("/invokers/INV-first/credentials", headers={"X-Actor": "alice@x"})
    assert r.status_code == 200
    body = r.json()
    assert body["keycloak_secret"] == "THE-SECRET"
    assert body["previous_reveal"] is None


def test_second_reveal_returns_first_in_previous():
    """The previous_reveal field reflects the LAST reveal — not this call's."""
    _seed_approved("INV-second")
    # First reveal
    client.get("/invokers/INV-second/credentials", headers={"X-Actor": "alice@x"})
    # Second reveal — should reference Alice's reveal in previous_reveal
    r = client.get("/invokers/INV-second/credentials", headers={"X-Actor": "bob@x"})
    assert r.status_code == 200
    prev = r.json()["previous_reveal"]
    assert prev is not None
    assert prev["actor"] == "alice@x"
    assert prev["at"] is not None


def test_reveal_writes_audit_row():
    _seed_approved("INV-aud")
    client.get("/invokers/INV-aud/credentials", headers={"X-Actor": "ops@x"})
    rows = list(db.audit_logs.find({"invoker_id": "INV-aud", "action": "credentials_revealed"}))
    assert len(rows) == 1
    assert rows[0]["actor"] == "ops@x"


def test_reveal_rejects_non_approved():
    _seed_approved("INV-rej")
    db.invokers.update_one({"invoker_id": "INV-rej"}, {"$set": {"approval_status": "suspended"}})
    r = client.get("/invokers/INV-rej/credentials", headers={"X-Actor": "x"})
    assert r.status_code == 403
