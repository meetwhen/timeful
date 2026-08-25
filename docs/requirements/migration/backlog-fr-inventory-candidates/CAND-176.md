---
id: CAND-176
verdict: proposed-requirement
requirement_type: FR
related_requirements: [FR-018]
confidence: inferred
---

# CAND-176: Dates-Only Edit-Event Entry

### Source

> - [x] On dates-only event page, when I click the "Edit event" button, a form for editing the event opens

[Source lines 546-546](../../../../backlog/backlog.md#L546)

### Candidate behavior

Selecting `Edit event` on a dates-only event page opens an event-editing form.

### Applicability

- Actor: unspecified event visitor
- Location: dates-only event page
- Event kind: dates-only
- Interaction mode: click Edit event
- Viewport: unspecified
- State: Edit event action available
- Exclusions: timed events

### Classification

candidate FR

### Existing Requirements and Confidence

Existing requirements: proposed [FR-018](../../functional/fr/FR-018.md) covers authorization to edit event settings, not this entry point.

Confidence: inferred

### Disposition

Consolidate with event-settings navigation if that navigation is durable.

### Open Questions

Is this entry point owner-only, and what should non-owners see?
