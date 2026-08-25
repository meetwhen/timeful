---
id: CAND-168
verdict: excluded
related_requirements: []
confidence: confirmed
---

# CAND-168: Postgres Anonymous-Event Storage

### Source

> - [x] Switch to Postgres for new anon events

[Source lines 536-536](../../../../backlog/backlog.md#L536)

### Candidate behavior

No new requirement behavior asserted; the source selects storage implementation for new anonymous events.

### Applicability

- Actor: system implementation team
- Location: anonymous-event creation
- Event kind: unspecified
- Interaction mode: not applicable
- Viewport: not applicable
- State: new anonymous events
- Exclusions: identifier routing and existing-event migration reporting

### Classification

implementation detail

### Existing Requirements and Confidence

Existing requirements: None identified.

Confidence: confirmed

### Disposition

Retain as migration provenance; do not create an FR or QR.
