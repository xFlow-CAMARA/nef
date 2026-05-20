"""Unit tests for db.py — encrypt/decrypt round-trip and graceful fallback."""

import os
import sys
import pathlib

# Make the package importable when pytest is run from the repo root
sys.path.insert(0, str(pathlib.Path(__file__).resolve().parent.parent))

# Force a stable dev key so the round-trip is deterministic
os.environ.setdefault("APP_ENV", "dev")
os.environ.setdefault("FIELD_KEY_PASSPHRASE", "unit-test-key")

from db import encrypt, decrypt  # noqa: E402


def test_round_trip_string():
    plain = "hello-world-123"
    cipher = encrypt(plain)
    assert cipher is not None and cipher != plain
    assert decrypt(cipher) == plain


def test_encrypt_none():
    assert encrypt(None) is None
    assert decrypt(None) is None


def test_decrypt_passthrough_on_garbage():
    """Non-Fernet input should round-trip unchanged so legacy plaintext docs
    survive deployment of the encryption layer."""
    assert decrypt("not-actually-encrypted") == "not-actually-encrypted"


def test_long_value_round_trip():
    """Whole PEM-sized payloads round-trip correctly."""
    pem = "-----BEGIN RSA PRIVATE KEY-----\n" + ("x" * 2048) + "\n-----END RSA PRIVATE KEY-----"
    assert decrypt(encrypt(pem)) == pem
