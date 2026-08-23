# ADR-008: Frontend Event Ownership Semantics

Date: 2026-05-18

Status:

- Accepted

## Context

Migrated frontend code must interpret legacy and anonymous event ownership
consistently without spreading sentinel checks through views and composables.

## Decision

The frontend keeps one shared event-ownership interpretation:

- `event.ownerId` remains the event ownership field.
- `guestUserId` is the canonical migrated representation of an
  anonymous-created event; missing or empty `ownerId` has the same fallback
  interpretation at the shared helper boundary.
- Migrated runtime code does not reintroduce legacy `ownerId == 0` checks.
- Shared ownership logic derives semantic ownership and permission values from
  raw event data rather than exposing raw sentinel checks to views.
- Metadata-edit and availability-edit permission values remain distinct.

## Related Requirements

Event-settings authority and non-owner visibility are defined by [FR-018](../../requirements/functional/fr/FR-018.md)
and [FR-039](../../requirements/functional/fr/FR-039.md). Anonymous
browser-local owners can associate event management with an authenticated
account as defined by [FR-063](../../requirements/functional/fr/FR-063.md).

Protected event and response mutations are governed by [QR-005](../../requirements/quality/qr/QR-005.md)
and [QR-006](../../requirements/quality/qr/QR-006.md).

## Quality Attributes

- Security: integrity and authenticity.
- Maintainability: modularity.

## Consequences

Frontend ownership checks share one interpretation while product authorization
behavior remains defined by functional and quality requirements.
