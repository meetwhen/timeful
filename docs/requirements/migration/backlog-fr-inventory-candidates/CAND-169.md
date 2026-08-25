# CAND-169: Postgres Anonymous-Event Identifier Routing

### Source

> - [x] Make 8-character Crockford base32 ids point to events in Postgres

[Source lines 537-537](../../../../backlog/backlog.md#L537)

### Candidate behavior

No new requirement behavior asserted; the source selects identifier-routing implementation for Postgres events.

### Applicability

- Actor: system implementation team
- Location: anonymous-event lookup
- Event kind: unspecified
- Interaction mode: not applicable
- Viewport: not applicable
- State: lookup by 8-character Crockford base32 identifier
- Exclusions: anonymous-event storage selection and existing-event migration reporting

### Classification

implementation detail

### Existing Requirements and Confidence

Existing requirements: None identified. Overlap: CAND-168 records the separate storage decision; no requirement overlap identified.

Confidence: confirmed

### Disposition

Retain as migration provenance; do not create an FR or QR.

### Open Questions

- Does the Postgres routing apply only to new anonymous events, or to every 8-character Crockford base32 identifier?
