# CAND-109

## Source

> - [x] On desktop, given I'm on the event page, when I hover an active timeslot or hover outside the grid and then hover over an inactive slot, then the Responses sidebar should show (0/N) and all responses crossed-out and the status square should have the corresponding color (light-grey for enabled, inactive etc.)

[Source lines 441-441](../../../../backlog/backlog.md#L441-L441)

## Candidate behavior

Hovering an inactive slot after leaving the grid shows `0/N`, crossed-out responses, and the corresponding status-square color.

## Applicability

Actor: event visitor; Location: desktop event-page grid and Responses; Event kind: timed; Interaction mode: pointer hover; Viewport: desktop; State: inactive slot after grid exit; Exclusions: active slot display except transition.

## Classification

duplicate or refinement

## Existing Requirements and Confidence

Overlap: CAND-107 specifies the collapsed-hours variant; no accepted FR/QR overlap. Confidence: inferred.

## Disposition

Consolidate with inactive-cell response-display behavior.

## Open Questions

Does this apply to disabled cells and every inactive status color?
