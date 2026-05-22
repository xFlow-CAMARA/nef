# Runbook — Rotate `FIELD_KEY`

`FIELD_KEY` (or `FIELD_KEY_PASSPHRASE`) protects PEM keys + Keycloak client
secrets at rest in `invokers.secrets.*`. Rotate it on a schedule, after a
suspected leak, or as part of any incident response.

## Pre-flight

- All read paths that touch `.secrets` use `decrypt()`, which **raises** on
  failure (no silent fallback). If you rotate without migrating, every read
  starts returning HTTP 503 — visible, recoverable.
- The migration script needs both keys available simultaneously. Decide
  *before* you start whether you're keeping the old key around (rollback)
  or destroying it.

## Steps

1. **Pick the new key.** Either a fresh 44-char URL-safe Fernet key
   (`python -c "from cryptography.fernet import Fernet; print(Fernet.generate_key().decode())"`)
   or a strong passphrase.

2. **Stop the invoker-onboarding service.** This avoids new writes while the
   migration runs.
   ```bash
   docker stop invoker-onboarding
   ```

3. **Take a Mongo backup.** Standard `mongodump` of the `camara` database is
   enough; the script is idempotent but irrecoverable encryption mistakes
   beat any "should be safe" claim.

4. **Run the migration.** From inside the project tree:
   ```bash
   docker run --rm --network xflow \
     -e MONGODB_URI=mongodb://camara-mongodb:27017/camara \
     -e OLD_FIELD_KEY_PASSPHRASE='current-passphrase' \
     -e NEW_FIELD_KEY_PASSPHRASE='new-passphrase' \
     -v $(pwd)/scripts:/work \
     -w /work python:3.11-slim \
     bash -c "pip install -q cryptography pymongo && python rotate-field-key.py"
   ```
   Expected output ends with `Done. total=N  migrated=N  already-new=0  failed=0`.
   Any `failed>0` means at least one doc didn't decrypt under either key —
   investigate before continuing.

5. **Switch the service to the new key** and restart.
   ```bash
   # If you use FIELD_KEY_PASSPHRASE in docker-compose / .env, update it here.
   docker start invoker-onboarding
   ```

6. **Smoke test.** Hit `GET /admin/invokers/{any-approved-id}/credentials`
   from the dashboard. A 200 with the right secret means rotation worked.
   A 503 means at least one doc still uses the old key — re-run step 4 with
   the right OLD value.

7. **Destroy the old key.** Wipe `OLD_FIELD_KEY_PASSPHRASE` from secrets
   management. Past mongodumps remain decryptable under the old key — apply
   the same retention policy you do to other crypto material.

## Rollback

If something goes wrong mid-step:
- Re-run the migration with OLD and NEW swapped — the script flips encrypted
  fields back, since it always checks decryptability under both keys before
  writing.
- If the service is already started under the new key but is failing reads,
  restart it under the old key while you investigate; the data is still on
  disk in old-key form until the migration writes the new ciphertext.

## What the script does NOT migrate

- Pure-plaintext fields (`invoker_id`, `invoker_name`, audit log rows) — they
  contain no sensitive material.
- Keycloak's own client secrets in Keycloak's database. Those are independent
  of `FIELD_KEY`. To rotate them, use the admin "Rotate Secret" action per
  invoker.
