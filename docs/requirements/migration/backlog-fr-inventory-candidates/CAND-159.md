---
id: CAND-159
verdict: proposed-requirement
requirement_type: FR
related_requirements: [FR-074, FR-103]
confidence: inferred
---

# CAND-159

### Source

> - [x] In the new timed event form, Advanced options must show the Time increment.

[Source lines 527-527](../../../../backlog/backlog.md#L527-L527)

### Candidate behavior

Advanced options in the new timed-event form show the time increment.

### Applicability

Actor: event owner.
Location: new timed-event form.
Event kind: timed.
Interaction mode: creating an event.
Viewport: any.
State: advanced options visible.
Exclusions: editing existing events.

### Classification

candidate FR

### Existing Requirements and Confidence

FR-074 uses the configured Slot Increment but does not require its visibility in the new timed-event form.
Confidence: inferred.

### Disposition

Promoted to [FR-103](../../functional/fr/FR-103.md) as owner-editable under Advanced options; the increment persists through validated event-settings updates.

### Open Questions

Is the control editable there, and is it visible only in a particular domain mode?
