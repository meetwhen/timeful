# Design Records

This directory is the canonical record of durable Timeful design decisions.
Design records describe architectural choices that constrain implementation,
while requirements specify outcomes.

## Record Types

- `adr/` contains architecture decision records (ADRs).
- Future design-record types belong in their own subdirectory and are indexed
  here.

## ADR Format

Each ADR has a permanent, zero-padded identifier and an ID-prefixed filename:

```text
adr/001-frontend-boundary-models-and-canonical-internal-shapes.md
```

IDs are assigned sequentially across the repository, not per component. An ADR
declares its scope in the title or context and includes context, decision,
consequences, quality attributes, and related requirements.

An ADR may link to the functional requirements (FRs) and quality requirements
(QRs) that it enables, constrains, or satisfies. FRs and QRs must remain
self-contained specifications and must not cite ADRs as normative dependencies.
Requirement provenance is recorded separately from the requirement statement.

## ADR Index

| ID | Scope | Title |
| --- | --- | --- |
| [ADR-001](adr/001-frontend-boundary-models-and-canonical-internal-shapes.md) | Frontend | Frontend Boundary Models and Canonical Internal Shapes |
| [ADR-002](adr/002-frontend-timezone-decoding-and-fixed-offset-boundaries.md) | Frontend | Frontend Timezone Decoding and Fixed-Offset Boundaries |
| [ADR-003](adr/003-frontend-freemium-operational-gating.md) | Frontend | Frontend Freemium Operational Gating |
| [ADR-004](adr/004-frontend-temporal-runtime-model.md) | Frontend | Frontend Temporal Runtime Model |
| [ADR-005](adr/005-frontend-civil-date-and-end-of-day-model.md) | Frontend | Frontend Civil-Date and End-of-Day Model |
| [ADR-006](adr/006-frontend-temporal-collection-semantics.md) | Frontend | Frontend Temporal Collection Semantics |
| [ADR-007](adr/007-frontend-semantic-styling-tokens.md) | Frontend | Frontend Semantic Styling Tokens |
| [ADR-008](adr/008-frontend-event-ownership-semantics.md) | Frontend | Frontend Event Ownership Semantics |
| [ADR-009](adr/009-frontend-guest-response-ownership-semantics.md) | Frontend | Frontend Guest Response Ownership Semantics |
