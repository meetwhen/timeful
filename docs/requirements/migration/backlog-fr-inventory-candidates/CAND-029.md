---
id: CAND-029
verdict: covered
related_requirements: [FR-002, FR-013]
confidence: confirmed
---

# CAND-029

## Source

> - [x] create event with specific availability in +2, 0-4 (day 1), 0-4 (day 2).
>       When opened in 0:00, should see the previous date

## Candidate behavior

The event page should show the prior projected date when slots project there in the selected display timezone.

## Applicability

Actor: event visitor.
Location: event-page timed grid.
Event kind: timed.
Interaction mode: viewing.
Viewport: any.
State: timezone projection crosses midnight.
Exclusions: event-timezone mutation.

## Classification

existing requirement

## Existing Requirements and Confidence

FR-002 requires projected date columns; FR-013 preserves instants across display-timezone changes.
Confidence: confirmed.

## Disposition

Map to FR-002 and FR-013; no migration.

## Open Questions

None.
