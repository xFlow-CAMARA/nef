"""POST /invokers — Pydantic validation + happy-path persistence.

CAPIF (register + onboarding + security context) is fully mocked via respx
so the test is hermetic. Upstream URLs come from conftest.py — pinned to
test-only hostnames that httpx will not strip default ports from.
"""

import respx
from httpx import Response
from fastapi.testclient import TestClient

import db
from main import app

client = TestClient(app)


def test_rejects_invoker_name_with_special_chars():
    """The name flows into a Keycloak client_id — only safe chars allowed."""
    r = client.post("/invokers", json={
        "invoker_name":   "../etc/passwd",
        "requested_apis": ["sim-swap"],
    })
    assert r.status_code == 422       # Pydantic StringConstraints rejects


def test_rejects_unknown_camara_scope():
    """requested_apis is a Literal — typos rejected at the boundary."""
    r = client.post("/invokers", json={
        "invoker_name":   "valid-name",
        "requested_apis": ["super-secret-internal-api"],
    })
    assert r.status_code == 422


@respx.mock
def test_happy_path_persists_encrypted_secrets():
    # Bootstrap token
    respx.get("https://test-register/getauth").mock(
        return_value=Response(200, json={"access_token": "stub-token"}))
    # Invoker onboarding (note: httpx strips :443 from default-port URLs)
    respx.post("https://test-capif/api-invoker-management/v1/onboardedInvokers").mock(
        return_value=Response(201, json={
            "apiInvokerId": "INV-test-happy",
            "onboardingInformation": {"apiInvokerCertificate": "-----BEGIN CERT-----\nfake\n-----END CERT-----"},
        }))
    # capif-service catalog: empty (skips security context registration)
    respx.get("http://test-capif-service/catalog").mock(return_value=Response(200, json=[]))

    r = client.post(
        "/invokers",
        json={
            "invoker_name":   "happy-app",
            "contact_email":  "submitted@example.com",
            "requested_apis": ["sim-swap"],
        },
        headers={"X-Actor": "trusted-dashboard@example.com"},
    )
    assert r.status_code == 201
    body = r.json()
    assert body["invoker_id"] == "INV-test-happy"
    assert body["approval_status"] == "pending"

    doc = db.invokers.find_one({"invoker_id": "INV-test-happy"})
    assert doc is not None
    # Sensitive material lives in .secrets and is Fernet-encrypted
    assert doc["secrets"]["key_pem"].startswith("gAAAAA")
    assert doc["secrets"]["cert_pem"].startswith("gAAAAA")
    assert doc["secrets"]["client_secret"].startswith("gAAAAA")
    # Audit row uses X-Actor (trusted), not contact_email (user-supplied)
    audit = db.audit_logs.find_one({"invoker_id": "INV-test-happy", "action": "submitted"})
    assert audit["actor"] == "trusted-dashboard@example.com"


def test_audit_actor_falls_back_to_anonymous_without_x_actor():
    with respx.mock:
        respx.get("https://test-register/getauth").mock(
            return_value=Response(200, json={"access_token": "t"}))
        respx.post("https://test-capif/api-invoker-management/v1/onboardedInvokers").mock(
            return_value=Response(201, json={
                "apiInvokerId": "INV-anon",
                "onboardingInformation": {"apiInvokerCertificate": "-----BEGIN CERT-----\nfake\n-----END CERT-----"},
            }))
        respx.get("http://test-capif-service/catalog").mock(return_value=Response(200, json=[]))

        r = client.post("/invokers", json={
            "invoker_name":   "no-actor-app",
            "contact_email":  "spoofed@example.com",     # ignored on audit row
            "requested_apis": ["sim-swap"],
        })
        assert r.status_code == 201

    audit = db.audit_logs.find_one({"invoker_id": "INV-anon", "action": "submitted"})
    assert audit["actor"] == "anonymous"
