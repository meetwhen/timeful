---
id: CAND-076
verdict: covered
related_requirements:
  - FR-020
confidence: confirmed
---

# CAND-076

## Source

> - [x] When a timeslot is selected, it's position is saved.
>       When the page is reloaded, the selection is rendered at the timeslot.
>       However, there's no tooltip near the selection.

## Candidate behavior

No new requirement behavior asserted; this is a persisted-selection tooltip regression, and persistence itself is not established.

## Applicability

Actor: mobile event visitor. Location: timed grid tooltip. Event kind: timed. Interaction mode: reload. Viewport: mobile. State: previously selected slot restored and visible. Exclusions: selection persistence policy.

## Classification

existing requirement

## Existing Requirements and Confidence

FR-020 requires an adjacent tooltip for a visible selected slot, but not selection persistence. Confidence: confirmed.

## Disposition

Map the missing-tooltip behavior to FR-020; retain the persistence observation as provenance.

## Open Questions

None.
