---
id: CAND-197
verdict: proposed-requirement
requirement_type: FR
related_requirements:
  - FR-020
  - FR-043
  - FR-096
confidence: inferred
---

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

candidate FR

### Existing Requirements and Confidence

Existing requirements: [FR-020](../../functional/fr/FR-020.md) requires a visible mobile selected-slot tooltip adjacent to a visible selected slot, and [FR-043](../../functional/fr/FR-043.md) requires placement at the interaction location; neither expressly requires screen containment.

Confidence: inferred

### Disposition

Promoted to [FR-096](../../functional/fr/FR-096.md) covering every mobile event-page tooltip.

### Open Questions

- Does this apply to every tooltip or only selected-slot tooltips?
