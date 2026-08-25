# CAND-197: Mobile Tooltip Screen Containment

### Source

> - [x] On mobile, on the event page, tooltip shall be fully on-screen

[Source lines 575-575](../../../../backlog/backlog.md#L575)

### Candidate behavior

On mobile event pages, a tooltip remains fully within the visible screen.

### Applicability

- Actor: event visitor
- Location: event page tooltip
- Event kind: unspecified
- Interaction mode: selection or hover-equivalent interaction
- Viewport: mobile
- State: tooltip visible
- Exclusions: desktop placement and tooltip content

### Classification

duplicate or refinement

### Existing Requirements and Confidence

Existing requirements: [FR-020](../../functional/fr/FR-020.md) already requires a visible mobile selected-slot tooltip adjacent to a visible selected slot; it does not expressly require screen containment.

Confidence: inferred

### Disposition

Refine FR-020 if containment is confirmed; do not create a separate requirement.

### Open Questions

- Does this apply to every tooltip or only selected-slot tooltips?
