# Design Records

This directory is the canonical record of durable Timeful design decisions.
Design records describe architectural choices that constrain implementation,
while requirements specify outcomes.

## Layout

- `architecture/README.md` contains architecture-decision-record authoring
  guidance.
- `architecture/adr/` contains architecture decision records (ADRs).
- Future design-record categories belong in their own subdirectory and are
  indexed here.

## Related Artifacts

- Functional requirements (FRs) state required system behavior.
- Quality requirements (QRs) state required quality attributes or runtime
  constraints.
- Specifications (SPECs) describe a feature or domain in broader context and
  may link to multiple requirements.
- Architecture decision records (ADRs) record architectural decisions.
- ADRs may link to the FRs and QRs that they enable, constrain, or satisfy.
  FRs and QRs remain self-contained specifications and must not cite ADRs as
  normative dependencies. Requirement provenance is recorded separately from
  the requirement statement.
- Backlog tasks track work to investigate, implement, or verify requirements.

## Architecture Records

| ID | Scope | Title |
| --- | --- | --- |
| [ADR-001](architecture/adr/001-frontend-boundary-models-and-canonical-internal-shapes.md) | Frontend | Frontend Boundary Models and Canonical Internal Shapes |
| [ADR-002](architecture/adr/002-frontend-timezone-decoding-and-fixed-offset-boundaries.md) | Frontend | Frontend Timezone Decoding and Fixed-Offset Boundaries |
| [ADR-003](architecture/adr/003-frontend-freemium-operational-gating.md) | Frontend | Frontend Freemium Operational Gating |
| [ADR-004](architecture/adr/004-frontend-temporal-runtime-model.md) | Frontend | Frontend Temporal Runtime Model |
| [ADR-005](architecture/adr/005-frontend-civil-date-and-end-of-day-model.md) | Frontend | Frontend Civil-Date and End-of-Day Model |
| [ADR-006](architecture/adr/006-frontend-temporal-collection-semantics.md) | Frontend | Frontend Temporal Collection Semantics |
| [ADR-007](architecture/adr/007-frontend-semantic-styling-tokens.md) | Frontend | Frontend Semantic Styling Tokens |
| [ADR-008](architecture/adr/008-frontend-event-ownership-semantics.md) | Frontend | Frontend Event Ownership Semantics |
| [ADR-009](architecture/adr/009-frontend-guest-response-ownership-semantics.md) | Frontend | Frontend Guest Response Ownership Semantics |
