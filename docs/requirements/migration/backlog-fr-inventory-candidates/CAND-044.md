---
id: CAND-044
verdict: covered
related_requirements:
  - FR-002
confidence: confirmed
---

# CAND-044

## Source

> - [x] Create event with 9 - 17, timezone +9 (<http://127.0.0.1:4173/e/ee4Cb>), june 11 and 12.
>       Expected: june 10 and 11 are shown on the event page
>       Actual: june 11, 12, 13 are shown there

## Candidate behavior

The event page should render the dates to which timed slots project in the selected display timezone.

## Applicability

Actor: event visitor. Location: event-page timed grid. Event kind: timed. Interaction mode: viewing. Viewport: any. State: timezone projection crosses dates. Exclusions: changing event timezone.

## Classification

existing requirement

## Existing Requirements and Confidence

FR-002 requires projected-date columns. Confidence: confirmed.

## Disposition

Map to FR-002; retain as regression provenance.

## Open Questions

None.
