---
id: CAND-079
verdict: proposed-requirement
requirement_type: FR
related_requirements:
  - FR-020
confidence: inferred
---

# CAND-079

## Source

> - [x] on mobile, the tooltip should be under top navbar when the grid gets scrolled up

## Candidate behavior

On mobile, when the timed grid is scrolled upward, the selected Timed Slot tooltip should appear below the top navbar.

## Applicability

Actor: mobile event visitor.
Location: timed grid tooltip.
Event kind: timed.
Interaction mode: scrolling.
Viewport: mobile.
State: grid scrolled upward.
Exclusions: unscrolled grid and no selected slot.

## Classification

candidate FR

## Existing Requirements and Confidence

FR-020 requires adjacency when the selected slot is visible, not a navbar relationship.
Confidence: inferred.

## Disposition

Candidate for a canonical FR defining tooltip placement relative to the top navbar.

## Open Questions

None.
