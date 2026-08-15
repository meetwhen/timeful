# PostgreSQL Anonymous Event Compatibility Contract

## Scope

These tables own only new anonymous timed and dates-only polls. MongoDB remains
authoritative for legacy events, authenticated creation, groups, signup forms,
folders, account adoption, and dashboard event loading. `postgres_events` and
`postgres_event_responses` are not HTTP DTOs and must not use BSON types.

`postgres_events.id` and `postgres_event_responses.id` are internal UUIDv7
identities. API handlers expose only `short_id`, an eight-character Crockford
Base32 identifier. PostgreSQL's UUID primary key remains internal.

## Explicit Columns

Event columns hold identifiers, soft-delete state, name, type, response count,
schedule version, creator PostHog ID, and timestamps. Response columns hold
the event relation, account-or-guest kind, response-map lookup identity, guest
ownership credentials, and timestamps. Mongo user IDs are external strings;
there is no PostgreSQL user table or cross-database foreign key.

`canonical_guest_name` is produced by `respondents.NormalizeGuestName` in Go.
PostgreSQL must not reimplement guest-name normalization. The partial indexes
preserve one response per account, token-backed guest ID, or canonical guest
name within an event.

## JSONB Payloads

`postgres_events.payload` holds all remaining event state, including:

- Description and nullable/default settings.
- Dates, active slots, timezone, slot generation, and timed recurrence.
- Legacy schedule fields retained for compatibility.
- Selected-schedule snapshot, remindees, attendee-compatible fields, and
  unsupported-feature fields accepted by existing request decoding.

`postgres_event_responses.payload` holds display name, email, availability,
if-needed availability, manual availability, and calendar-related fields.

JSON arrays preserve their existing behavior: date and recurrence arrays keep
input order and duplicates; active slots are normalized by route validation;
response availability keeps first-seen order after deduplication. Instants are
normalized to millisecond precision before writing JSONB and before API output.

## Compatibility Rules

The repository must distinguish absent fields, JSON null, empty arrays/maps,
and zero scalar values. In particular, an omitted description preserves the
existing value, an explicit empty description persists, and a timed edit with
an explicit empty `activeSlots` retains the existing slots to match Mongo BSON
`omitempty` behavior. Public schedule save/replace/clear remains supported.

`guest_edit_token` intentionally stores the raw token in phase one. Open,
tokenless mutation must return the stored credential, so a hash-only design
would change current behavior. Token hashing is deferred until the API and
frontend ownership contract change together.

## Transactions

Response create, update, and delete lock the event row and update the response
row plus `num_responses` in one transaction. Guest rename, policy changes, and
legacy-to-token transitions are also transactional. Unique-index conflicts
must map to the existing duplicate-name route error. Event edit and selected
schedule replace/clear write the event aggregate atomically.

Transactions deliberately prevent duplicate response races and response-count
drift; reproducing those internal Mongo failure modes is not required for API
compatibility.
