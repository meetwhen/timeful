# CAND-070

## Source

> - [x] on mobile and desktop, there should be no delay between selecting the time slot and seeing the tooltip with time and date

## Candidate behavior

No new requirement behavior asserted; this requests a latency target without a measurable threshold.

## Applicability

Actor: event visitor. Location: timed grid tooltip. Event kind: timed. Interaction mode: slot selection. Viewport: mobile and desktop. State: slot selected. Exclusions: no selected slot.

## Classification

needs product decision

## Existing Requirements and Confidence

FR-020 requires the mobile tooltip when the selected slot is visible; QR-009 is the relevant performance family. Confidence: needs product decision.

## Disposition

Hold for a measurable responsiveness scenario.

## Open Questions

What maximum selection-to-tooltip latency is acceptable, and which devices are in scope?
