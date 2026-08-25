---
id: CAND-198
verdict: proposed-requirement
requirement_type: FR
related_requirements:
  - FR-011
  - FR-048
confidence: inferred
---

# CAND-198: Mobile Availability Show-All-Hours Placement

### Source

> - [x] On mobile, when adding availability, Show all hours should be an option at its normal place - under other options between the event description and the grid

[Source lines 576-576](../../../../backlog/backlog.md#L576)

### Candidate behavior

On mobile while adding availability, `Show all hours` appears under other options between the event description and the grid.

### Applicability

- Actor: event visitor
- Location: mobile event page
- Event kind: timed
- Interaction mode: adding availability
- Viewport: mobile
- State: other options and event description visible
- Exclusions: editing availability, desktop layout, and dates-only pages

### Classification

candidate FR

### Existing Requirements and Confidence

Existing requirements: [FR-011](../../functional/fr/FR-011.md) covers the expansion behavior; [FR-048](../../functional/fr/FR-048.md) is proposed and covers centering in a no-response state, not this placement.

Confidence: inferred

### Disposition

Consolidate with mobile timed-grid controls if the “normal place” is formally defined.

### Open Questions

- Does “adding availability” include editing an existing availability response?
