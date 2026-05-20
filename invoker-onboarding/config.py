"""Shared configuration helpers.

Two principles enforced here:
  1. Every secret read is a function call — never a module-level constant.
     Lets workers pick up rotated values on restart, and lets tests use
     monkeypatch.setenv without needing importlib.reload.
  2. APP_ENV=dev is a positive assertion: the service refuses to start in
     dev mode if any upstream URL points at something that isn't localhost
     or a known Docker hostname. Stops dev defaults from accidentally
     decrypting production data.
"""

import logging
import os
from urllib.parse import urlparse

log = logging.getLogger("invoker-onboarding.config")

# Hostnames considered "local enough" for APP_ENV=dev. Anything else triggers
# refusal at startup.
_DEV_OK_HOSTS = {
    "localhost", "127.0.0.1", "::1",
    "host.docker.internal", "host-gateway",
}


def app_env() -> str:
    """Return the current APP_ENV (read at call time)."""
    return os.getenv("APP_ENV", "dev")


def required(name: str, dev_default: str) -> str:
    """Read an env var. Falls back to `dev_default` only when APP_ENV == 'dev'.
    Raises RuntimeError otherwise.
    """
    v = os.getenv(name)
    if v:
        return v
    if app_env() == "dev":
        log.warning("%s not set — using dev default", name)
        return dev_default
    raise RuntimeError(f"{name} env var is required when APP_ENV != 'dev'")


def _looks_local(url: str) -> bool:
    if not url:
        return True
    host = urlparse(url).hostname or ""
    if not host:
        return True                      # bare hostnames like 'redis' come back empty when no scheme
    if host in _DEV_OK_HOSTS:
        return True
    # Anything that looks like an internal Docker name (no dots, no public TLD)
    if "." not in host:
        return True
    return False


def assert_dev_is_local() -> None:
    """Refuse to run with APP_ENV=dev if any upstream URL looks remote."""
    if app_env() != "dev":
        return
    for var in ("MONGODB_URI", "KEYCLOAK_URL", "CAPIF_CORE_URL", "CAPIF_REGISTER_URL"):
        url = os.getenv(var, "")
        if not _looks_local(url):
            raise RuntimeError(
                f"APP_ENV=dev but {var}={url!r} appears to be a non-local host; "
                f"refusing to use dev fallbacks against possibly-real data"
            )
