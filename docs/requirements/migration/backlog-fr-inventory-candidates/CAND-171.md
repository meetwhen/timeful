---
id: CAND-171
verdict: proposed-requirement
requirement_type: FR
related_requirements: [FR-011, FR-105]
confidence: inferred
---

# CAND-171: Timed More-Options Order

### Source

> - [x] On the timed event page, in More options, "Show all hours" must be above "Hide if needed times"

[Source lines 541-541](../../../../backlog/backlog.md#L541)

### Candidate behavior

On the timed event page, More options presents `Show all hours` above `Hide if needed times`.

### Applicability

- Actor: event visitor
- Location: timed event page, More options
- Event kind: timed
- Interaction mode: viewing
- Viewport: unspecified
- State: More options open
- Exclusions: dates-only event page

### Classification

candidate FR

### Existing Requirements and Confidence

Existing requirements: [FR-011](../../functional/fr/FR-011.md) establishes collapsed timed-grid behavior but not this control order.

Confidence: inferred

### Disposition

Promoted to [FR-105](../../functional/fr/FR-105.md); the owner fixed the order across viewports and states.

### Open Questions

- Is the required order intended for every viewport and control state?
