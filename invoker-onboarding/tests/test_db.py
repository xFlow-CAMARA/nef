"""Unit tests for db.py — encrypt/decrypt round-trip + strict failure mode."""

import pytest
from cryptography.fernet import InvalidToken

from db import encrypt, decrypt


def test_round_trip_string():
    plain = "hello-world-123"
    cipher = encrypt(plain)
    assert cipher is not None and cipher != plain
    assert decrypt(cipher) == plain


def test_encrypt_none():
    assert encrypt(None) is None
    assert decrypt(None) is None


def test_decrypt_invalid_raises_by_default():
    """A garbage value must NOT silently come back as plaintext —
    that's how stale credentials end up displayed to end users."""
    with pytest.raises(InvalidToken):
        decrypt("not-actually-encrypted")


def test_decrypt_allows_plaintext_only_when_opted_in():
    """For one-time migration off plaintext, callers may opt in."""
    assert decrypt("legacy-plaintext", allow_plaintext_fallback=True) == "legacy-plaintext"


def test_decrypt_fallback_still_rejects_real_ciphertext_corruption():
    """A corrupted Fernet token (gAAAAA…) must still raise even with the
    fallback enabled — it's clearly meant to be encrypted, just broken."""
    with pytest.raises(InvalidToken):
        decrypt("gAAAAAcorruptedtoken", allow_plaintext_fallback=True)


def test_long_value_round_trip():
    pem = "-----BEGIN RSA PRIVATE KEY-----\n" + ("x" * 2048) + "\n-----END RSA PRIVATE KEY-----"
    assert decrypt(encrypt(pem)) == pem
