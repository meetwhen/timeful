---
id: CAND-119
verdict: covered
related_requirements: [FR-012]
confidence: confirmed
---

# CAND-119

## Source

> - [x] Given I'm on the event page and there's a scheduled event, when I click Reschedule event, the Schedule button is active

[Source lines 453-453](../../../../backlog/backlog.md#L453-L453)

## Candidate behavior

When an event is already scheduled, entering Reschedule event leaves Schedule active.

## Applicability

Actor: event visitor; Location: event page scheduling controls; Event kind: timed; Interaction mode: reschedule; Viewport: unspecified; State: scheduled event exists; Exclusions: unscheduled event not established.

## Classification

existing requirement

## Existing Requirements and Confidence

Overlap: accepted FR-012 requires that a Scheduled Event Time can be replaced. Confidence: confirmed.

## Disposition

Map to FR-012; retain the active-button presentation only if separately needed.

## Open Questions

Does active mean enabled, selected, or immediately actionable without changing slots?
