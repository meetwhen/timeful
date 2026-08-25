# CAND-079

## Source

> - [x] on mobile, the tooltip should be under top navbar when the grid gets scrolled up

## Candidate behavior

No new requirement behavior asserted; this is a placement refinement not established by the accepted adjacency rule.

## Applicability

Actor: mobile event visitor. Location: timed grid tooltip. Event kind: timed. Interaction mode: scrolling. Viewport: mobile. State: grid scrolled upward. Exclusions: unscrolled grid and no selected slot.

## Classification

needs product decision

## Existing Requirements and Confidence

FR-020 requires adjacency when the selected slot is visible, not a navbar relationship. Confidence: confirmed.

## Disposition

Do not migrate pending placement policy.

## Open Questions

If the selected slot is under the navbar, should the tooltip be hidden, clipped, or repositioned?
