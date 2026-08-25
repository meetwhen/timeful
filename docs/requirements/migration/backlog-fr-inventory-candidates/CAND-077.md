---
id: CAND-077
verdict: covered
related_requirements:
  - FR-020
confidence: confirmed
---

# CAND-077

## Source

> - [x] on mobile, when long press inside the grid changes the selected timeslot, the tooltip must also appear near the selected timeslot
>   - currently, it stays near the previously selected timeslot

## Candidate behavior

No new requirement behavior asserted; accepted behavior already requires adjacency to the selected visible slot.

## Applicability

Actor: mobile event visitor. Location: timed grid tooltip. Event kind: timed. Interaction mode: long press. Viewport: mobile. State: selected slot changes and is visible. Exclusions: no selected slot.

## Classification

existing requirement

## Existing Requirements and Confidence

FR-020 requires the tooltip adjacent to the visible selected slot. Confidence: confirmed.

## Disposition

Treat as FR-020 regression.

## Open Questions

None.
