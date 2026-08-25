---
id: CAND-184
verdict: proposed-requirement
requirement_type: FR
related_requirements:
  - FR-018
confidence: inferred
---

# CAND-184: Edit-Event Button Placement

### Source

> - [x] On the event page, Edit event button shall be under the event title

[Source lines 559-559](../../../../backlog/backlog.md#L559)

### Candidate behavior

The event page places the `Edit event` button below the event title.

### Applicability

- Actor: event owner
- Location: event page header
- Event kind: unspecified
- Interaction mode: viewing
- Viewport: unspecified
- State: edit action visible
- Exclusions: edit authorization and non-owner presentation

### Classification

candidate FR

### Existing Requirements and Confidence

Existing requirements: [FR-018](../../functional/fr/FR-018.md) covers edit authorization, not button placement.

Confidence: inferred

### Disposition

Consolidate only if placement is durable across responsive layouts.

### Open Questions

- Does “under” apply to both desktop and mobile headers?
