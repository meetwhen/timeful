---
id: CAND-074
verdict: covered
related_requirements:
  - FR-020
confidence: confirmed
---

# CAND-074

## Source

> - [x] on mobile, when scrolling, the tooltip with time should scroll together with selected cell, not stay frozen at the same height of the screen while the screen scrolls

## Candidate behavior

No new requirement behavior asserted; accepted behavior already specifies tooltip availability adjacent to a visible selected slot after scrolling.

## Applicability

Actor: mobile event visitor. Location: timed grid tooltip. Event kind: timed. Interaction mode: scrolling. Viewport: mobile. State: selected slot. Exclusions: no selected slot or selected slot not visible.

## Classification

existing requirement

## Existing Requirements and Confidence

FR-020 requires adjacency when the selected slot is visible. Confidence: confirmed.

## Disposition

Map to FR-020; no migration.

## Open Questions

None.
