---
id: CAND-194
verdict: proposed-requirement
requirement_type: FR
related_requirements:
  - FR-036
confidence: inferred
---

# CAND-194: Timezone-Reset Icon Direction

### Source

> - [x] The icon to reset the timezone shall have a counter-clockwise arrow

[Source lines 572-572](../../../../backlog/backlog.md#L572)

### Candidate behavior

The timezone-reset icon uses a counter-clockwise arrow.

### Applicability

- Actor: event visitor
- Location: timezone reset control
- Event kind: timed
- Interaction mode: viewing
- Viewport: unspecified
- State: reset control visible
- Exclusions: reset behavior and accessible name

### Classification

candidate FR

### Existing Requirements and Confidence

Existing requirements: [FR-036](../../functional/fr/FR-036.md) specifies reset behavior but not the icon direction.

Confidence: inferred

### Disposition

Consolidate only if icon direction carries a product meaning rather than implementation styling.

### Open Questions

- Is an accessible text label required in addition to the icon?
