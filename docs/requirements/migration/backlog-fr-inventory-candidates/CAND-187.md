---
id: CAND-187
verdict: proposed-requirement
requirement_type: FR
related_requirements:
  - FR-011
confidence: inferred
---

# CAND-187: Collapsed-Strip Hover Treatment

### Source

> - [x] On the event page, when I hover over collapsed hours strip, neither the strip nor any cell is highlighted like active cells are highlighted on hover.

[Source lines 562-562](../../../../backlog/backlog.md#L562)

### Candidate behavior

Hovering a collapsed-hours strip does not apply the active-cell highlight to the strip or any grid cell.

### Applicability

- Actor: event visitor
- Location: event page grid
- Event kind: timed
- Interaction mode: viewing
- Viewport: hover-capable
- State: collapsed hours visible
- Exclusions: active-cell hover behavior

### Classification

candidate FR

### Existing Requirements and Confidence

Existing requirements: [FR-011](../../functional/fr/FR-011.md) establishes collapsed timed-grid runs but does not specify hover treatment.

Confidence: inferred

### Disposition

Consolidate as a grid interaction refinement if the treatment is intentionally observable.

### Open Questions

- Should touch interaction have an equivalent non-highlight rule?
