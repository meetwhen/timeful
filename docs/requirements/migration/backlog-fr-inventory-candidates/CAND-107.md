---
id: CAND-107
verdict: proposed-requirement
requirement_type: FR
related_requirements: [FR-011, FR-045]
confidence: inferred
---

# CAND-107

## Source

> - [x] Given I'm on desktop and on the event page, when I hover over collapsed hours,
>   - the responses should show 0/N (behave similarly to disabled cells)
>   - and mark everyone unavailable
>   - and the selection at the highlight at the last hovered timeslot must be cleared

[Source lines 436-439](../../../../backlog/backlog.md#L436-L439)

## Candidate behavior

Hovering collapsed hours on desktop shows `0/N`, marks all responses unavailable, and clears the prior highlight selection.

## Applicability

Actor: event visitor; Location: desktop timed grid and Responses; Event kind: timed; Interaction mode: hover; Viewport: desktop; State: collapsed hours; Exclusions: active and disabled cells.

## Classification

candidate FR

## Existing Requirements and Confidence

Overlap: accepted FR-011 defines collapsed inactive runs, and proposed FR-045 covers unavailable response display; neither covers clearing the prior selection.
Confidence: inferred.

## Disposition

Review as collapsed-hour interaction behavior.

## Open Questions

Does `0/N` include protected responses, and what visual mark means unavailable?
