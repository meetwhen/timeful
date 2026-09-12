# PostgreSQL Data Boundaries

## Scope

PostgreSQL is the only store for accounts, profiles, events of every kind, responses, attendees, signup data, availability groups, folders, folder membership, calendar integrations, OTP challenges, and historical daily user logs.
This document is the durable backend contract for the retained record kinds: calendar integrations, OTP challenges, historical daily user logs, and reporting reads.
It fixes the single-authoritative-store boundary, the single account identifier rule, the fresh-identity rules, the OTP one-way-hash rule, and the provider-credential encryption boundary.
The observable API behavior of PostgreSQL-owned events remains governed by the [PostgreSQL Anonymous Event Compatibility Contract](postgres-anonymous-event-compatibility.md).
The durable decisions are recorded as [ADR-018](../../docs/design/architecture/adr/ADR-018.md), [ADR-020](../../docs/design/architecture/adr/ADR-020.md), and [ADR-021](../../docs/design/architecture/adr/ADR-021.md).

## Authoritative Store Ownership

Every record has exactly one authoritative store, and PostgreSQL is that store for every record kind.
A record is authoritative in the store that accepts its reads and writes, and no other store may be consulted for that record or re-authorize it.
No permanent dual write exists, so a record is authored in exactly one store.

| Record kind                                                                                                                   | Authoritative store      | Notes                                                                                       |
| ----------------------------------------------------------------------------------------------------------------------------- | ------------------------ | ------------------------------------------------------------------------------------------- |
| [Calendar Connection](../../docs/terminology/glossary.md#calendar-connection), provider credentials, and calendar preferences | PostgreSQL               | The whole integration aggregate lives together per account.                                 |
| Sub-calendar and its enabled state                                                                                            | PostgreSQL               | Child of its [Calendar Connection](../../docs/terminology/glossary.md#calendar-connection). |
| OTP challenge                                                                                                                 | PostgreSQL               | Ephemeral; one active challenge per email.                                                  |
| Historical daily user log and its account membership                                                                          | PostgreSQL               | Membership references an authoritative account.                                             |
| Event-creator analytics                                                                                                       | PostgreSQL event storage | Reads `postgres_events.creator_posthog_id`; no event row is copied again.                   |
| Active-user and signed-up-user reporting                                                                                      | PostgreSQL               | Reads `daily_user_logs`, `daily_user_log_members`, and `accounts`.                          |
| Friend request                                                                                                                | Retired                  | The dormant collection and its accessors are removed, and no table replaces it.             |
| Core account, event, response, attendee, signup, group, and folder data                                                       | PostgreSQL               | Governed by the event compatibility contract and the PostgreSQL schema.                     |

No record kind has a second read path, and no compatibility redirect serves a record from another store.

## Fresh Identity And Single Authority

### Accounts

The platform identity is the account's sole identifier: `platform_identities.id` is a native UUIDv7, and the legacy `external_user_id` column is removed.
Every account reference is a native `uuid` column referencing `platform_identities(id)`, and the sign-in session carries the canonical lowercase hyphenated UUID string.
Account deletion records that uuid in `account_deletion_tombstones` without a foreign key, so the tombstone survives the platform identity's deletion.
The tombstone is an audit record rather than a runtime gate, because deletion removes the `platform_identities` row that account resolution keys on and a later sign-in mints a fresh uuidv7 identity, so a deleted identity can never be adopted again.
Legacy tombstones that name a pre-cutover 24-hex identity are dropped by migration `20260912120000` because no live platform identity exists to protect; export them first with `server/scripts/20260912_export_legacy_deletion_tombstones/` when deletion audit retention is required.
The all-zero UUID is the wire representation of an absent account identity, such as a guest response's `userId` or an unowned event's `ownerId`, and it is never a stored platform identity.

### Calendar Accounts And Sub-Calendars

Each [Calendar Connection](../../docs/terminology/glossary.md#calendar-connection) receives a fresh `calendar_accounts.id` and is owned by exactly one `platform_identities` row.
Its runtime key remains the `email_calendarType` map key with the same normalization and legacy-suffix behavior, so the existing key semantics are preserved even though the identity is new.
Sub-calendars receive fresh identities scoped to their **Calendar Connection**, and the legacy account-map key is not preserved as a permanent lookup.
The connection, its encrypted credentials, its sub-calendars, and the calendar preferences are read and written only in PostgreSQL.

### OTP Challenges

Each OTP challenge receives a fresh identity and is keyed by email.
A challenge grants no account authority: a successful verification resolves an authoritative PostgreSQL account exactly as before.
At most one challenge is active per email, and sending a new challenge replaces any existing challenge for that email.
The one-time code is stored as a salted one-way hash, never in plaintext, and verification compares hashes.
The ten-minute expiry and the five-attempt lockout are preserved.

### Daily Logs

Each daily user log receives a fresh identity keyed by its account-local date, and each membership references an authoritative account.
Membership stays idempotent per account per day and preserves first-seen order.
The log and its memberships are read and written only in PostgreSQL.

## Credential Encryption Boundary

Provider credentials are encrypted at rest with authenticated AES-256-GCM and an envelope that carries a format version.

- Algorithm: AES-256 in Galois/Counter Mode with a 128-bit authentication tag.
- Key source: the `ENCRYPTION_KEY` environment variable, read as exactly 32 raw bytes; startup validation refuses a missing or wrong-length key rather than deriving one, and the key is never logged.
- Nonce: a fresh 96-bit random nonce per encryption, never reused for a key.
- Envelope: the versioned string `v1:` followed by the base64 encoding of `nonce || ciphertext || tag`, so a future key or algorithm change is distinguishable without ambiguity.
- Encrypted fields: the OAuth2 access token, the OAuth2 refresh token, the Apple app password, and the ICS feed URL.
- Plaintext fields: access-token expiry, granted OAuth2 scope, email, picture, enabled state, sub-calendar name and enabled state, primary calendar preference, token origin, and calendar options; none of these is a provider secret.
- OTP one-time codes are not provider credentials and are not stored with this envelope; they are stored as salted one-way hashes so the code is not readable at rest.

Decryption failures are errors, never empty credentials.
A failed decrypt of an access token fails the token refresh for that connection and is reported; it never submits an empty credential.
A failed decrypt of a refresh token, Apple password, or ICS feed URL surfaces as an error to the caller and does not partially load a connection.

Credentials are never returned through the API.
No calendar route may expose a token, password, or feed URL in a response.

## Reporting Reads

Event-creator analytics aggregate `postgres_events.creator_posthog_id`, counting each migrated or new event exactly once.
Active-user reporting reads `daily_user_logs` and `daily_user_log_members`, and signed-up-user reporting reads `accounts`.
Reporting reads only PostgreSQL records and copies no reporting data to a second store.
