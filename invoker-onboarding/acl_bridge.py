"""
CAPIF ACL bridge — publishes Redis events to create/remove access-control
policy entries in the running OpenCAPIF ACL service.

The ACL service listens on the 'acls-messages' Redis pub/sub channel and
processes messages of the form:
  create-acl:{invoker_id}:{service_id}:{aef_id}
  remove-acl:{invoker_id}:{service_id}:{aef_id}

To map approved scope names (e.g. "quality-on-demand") to CAPIF service_id /
aef_id pairs we call the capif-service /catalog endpoint, which returns the
published API catalogue.
"""

import logging
import os

import httpx
import redis

log = logging.getLogger("invoker-onboarding.acl")

CAPIF_REDIS_URL  = os.getenv("CAPIF_REDIS_URL", "redis://services-redis-1:6379/0")
CAPIF_SERVICE_URL = os.getenv("CAPIF_SERVICE_URL", "http://capif-service:8080")

_redis: redis.Redis | None = None


def _get_redis() -> redis.Redis:
    global _redis
    if _redis is None:
        _redis = redis.from_url(CAPIF_REDIS_URL, decode_responses=True)
    return _redis


def _catalog_map() -> dict[str, dict]:
    """
    Return a dict mapping api_name → {service_id, aef_id} from the
    capif-service catalog endpoint.  Returns empty dict on failure.
    """
    try:
        r = httpx.get(f"{CAPIF_SERVICE_URL}/catalog", timeout=5.0)
        r.raise_for_status()
        entries = r.json()
        return {
            e["api_name"]: {"service_id": e["service_id"], "aef_id": e["aef_id"]}
            for e in entries
            if "api_name" in e and "service_id" in e and "aef_id" in e
        }
    except Exception as e:
        log.warning("Could not fetch CAPIF catalog: %s", e)
        return {}


def create_acl_entries(invoker_id: str, approved_scopes: list[str]) -> list[str]:
    """
    Publish create-acl events for every approved scope.

    Returns the list of scope names for which an ACL event was published.
    """
    catalog = _catalog_map()
    if not catalog:
        log.warning("Empty catalog — no ACL entries will be created for %s", invoker_id)
        return []

    r = _get_redis()
    published = []
    for scope in approved_scopes:
        entry = catalog.get(scope)
        if not entry:
            log.warning("Scope %s not found in CAPIF catalog — skipping ACL", scope)
            continue
        message = f"create-acl:{invoker_id}:{entry['service_id']}:{entry['aef_id']}"
        try:
            r.publish("acls-messages", message)
            log.info("Published ACL event: %s", message)
            published.append(scope)
        except Exception as e:
            log.error("Redis publish failed for %s/%s: %s", invoker_id, scope, e)

    return published


def remove_acl_entries(invoker_id: str) -> None:
    """
    Publish a remove-acl event to remove ALL ACL entries for an invoker.
    Uses the single-id form that triggers remove_invoker_acl() in the ACL service.
    """
    try:
        r = _get_redis()
        message = f"remove-acl:{invoker_id}"
        r.publish("acls-messages", message)
        log.info("Published ACL removal event for invoker %s", invoker_id)
    except Exception as e:
        log.error("Redis publish failed for ACL removal of %s: %s", invoker_id, e)
