# ADR-005: Frontend Scheduling Time Semantics

Date: 2026-05-06

Status:

- Accepted

## Context

Scheduling code must not silently substitute datetime, sentinel, or inferred
week behavior for explicit domain semantics.

## Decision

The frontend keeps scheduling time semantics explicit:

- Civil-date concepts use `Temporal.PlainDate` when timezone is not part of
  the invariant.
- End-of-day remains an explicit boundary concept rather than a fake `24:00`
  time value.
- Working-hours and overlap construction use normalized Temporal values with
  the selected schedule timezone or fixed offset.
- Weekly helper call sites state whether they use seed-week or rendered-week
  semantics.

## Consequences

Scheduling-domain code makes its civil-date, end-of-day, timezone, and weekly
interpretation explicit at the boundary where it is modeled.
