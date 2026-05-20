"""
Keycloak bridge — called when an admin approves a CAPIF invoker.

On approval:
  1. Obtain a Keycloak admin token (client_credentials on master realm)
  2. Create a confidential Keycloak client in the 'camara' realm
     scoped only to the operator-approved CAMARA APIs
  3. Return the generated client_id and client_secret

On revocation:
  1. Obtain admin token
  2. Delete the Keycloak client by its internal UUID
"""

import logging
import os

import httpx

from config import CAMARA_SCOPES, required

log = logging.getLogger("invoker-onboarding.keycloak")

KEYCLOAK_URL      = os.getenv("KEYCLOAK_URL",   "http://keycloak:8080")
KEYCLOAK_REALM    = os.getenv("KEYCLOAK_REALM", "camara")
KEYCLOAK_ADMIN_ID = os.getenv("KEYCLOAK_ADMIN_CLIENT_ID", "admin-cli")
KEYCLOAK_ADMIN_UN = required("KEYCLOAK_ADMIN_USERNAME", "admin")
KEYCLOAK_ADMIN_PW = required("KEYCLOAK_ADMIN_PASSWORD", "admin")

# Re-export the central scope list for backward compatibility — the
# canonical source is config.CAMARA_SCOPES.
ALL_CAMARA_SCOPES = list(CAMARA_SCOPES)

_http = httpx.Client(timeout=15.0)


def _admin_token() -> str:
    r = _http.post(
        f"{KEYCLOAK_URL}/realms/master/protocol/openid-connect/token",
        data={
            "grant_type": "password",
            "client_id":  KEYCLOAK_ADMIN_ID,
            "username":   KEYCLOAK_ADMIN_UN,
            "password":   KEYCLOAK_ADMIN_PW,
        },
    )
    r.raise_for_status()
    return r.json()["access_token"]


def _scope_uuid(admin_token: str, scope_name: str) -> str | None:
    """Look up the internal UUID of a client scope by name."""
    r = _http.get(
        f"{KEYCLOAK_URL}/admin/realms/{KEYCLOAK_REALM}/client-scopes",
        headers={"Authorization": f"Bearer {admin_token}"},
    )
    r.raise_for_status()
    for s in r.json():
        if s["name"] == scope_name:
            return s["id"]
    return None


def create_keycloak_client(
    invoker_id: str,
    invoker_name: str,
    approved_scopes: list[str],
) -> dict:
    """
    Create a Keycloak client for the approved invoker.

    Returns {"client_id": str, "client_secret": str, "keycloak_uuid": str}
    """
    admin_token = _admin_token()
    client_id   = invoker_id          # reuse CAPIF invoker_id as Keycloak client_id

    # Resolve scope UUIDs
    optional_scope_ids = []
    for scope in approved_scopes:
        if scope not in ALL_CAMARA_SCOPES:
            log.warning("Unknown scope %s — skipping", scope)
            continue
        uid = _scope_uuid(admin_token, scope)
        if uid:
            optional_scope_ids.append(uid)
        else:
            log.warning("Scope %s not found in Keycloak realm — skipping", scope)

    client_payload = {
        "clientId":                client_id,
        "name":                    invoker_name,
        "description":             f"Auto-created for CAPIF invoker {invoker_id}",
        "enabled":                 True,
        "protocol":                "openid-connect",
        "publicClient":            False,
        "serviceAccountsEnabled":  True,   # enables client_credentials grant
        "standardFlowEnabled":     False,
        "directAccessGrantsEnabled": False,
        "optionalClientScopes":    [],
        "defaultClientScopes":     ["email", "profile", "roles", "web-origins"],
    }

    headers = {"Authorization": f"Bearer {admin_token}"}

    # Create client
    r = _http.post(
        f"{KEYCLOAK_URL}/admin/realms/{KEYCLOAK_REALM}/clients",
        json=client_payload,
        headers=headers,
    )
    if r.status_code == 409:
        log.warning("Keycloak client %s already exists — fetching existing", client_id)
    elif r.status_code not in (200, 201):
        log.error("Failed to create Keycloak client: %d %s", r.status_code, r.text)
        raise RuntimeError(f"Keycloak client creation failed: {r.status_code} {r.text}")

    # Fetch the newly created client to get its UUID
    r2 = _http.get(
        f"{KEYCLOAK_URL}/admin/realms/{KEYCLOAK_REALM}/clients",
        params={"clientId": client_id},
        headers=headers,
    )
    r2.raise_for_status()
    clients = r2.json()
    if not clients:
        raise RuntimeError(f"Could not find Keycloak client {client_id} after creation")
    kc_uuid = clients[0]["id"]

    # Assign the approved scopes as optional client scopes
    for scope_id in optional_scope_ids:
        sr = _http.put(
            f"{KEYCLOAK_URL}/admin/realms/{KEYCLOAK_REALM}/clients/{kc_uuid}/optional-client-scopes/{scope_id}",
            headers=headers,
        )
        if sr.status_code not in (200, 204):
            log.warning("Could not assign scope %s to client %s: %d", scope_id, client_id, sr.status_code)

    # Retrieve client secret
    sec_r = _http.get(
        f"{KEYCLOAK_URL}/admin/realms/{KEYCLOAK_REALM}/clients/{kc_uuid}/client-secret",
        headers=headers,
    )
    sec_r.raise_for_status()
    client_secret = sec_r.json().get("value", "")

    log.info(
        "Keycloak client created: %s  scopes=%s",
        client_id, approved_scopes,
    )
    return {
        "client_id":      client_id,
        "client_secret":  client_secret,
        "keycloak_uuid":  kc_uuid,
    }


def delete_keycloak_client(keycloak_uuid: str) -> None:
    """Delete a Keycloak client by its internal UUID (on invoker revocation)."""
    try:
        admin_token = _admin_token()
        r = _http.delete(
            f"{KEYCLOAK_URL}/admin/realms/{KEYCLOAK_REALM}/clients/{keycloak_uuid}",
            headers={"Authorization": f"Bearer {admin_token}"},
        )
        if r.status_code not in (200, 204, 404):
            log.error("Failed to delete Keycloak client %s: %d", keycloak_uuid, r.status_code)
        else:
            log.info("Keycloak client %s deleted", keycloak_uuid)
    except Exception as e:
        log.error("Error deleting Keycloak client %s: %s", keycloak_uuid, e)


