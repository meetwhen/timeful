---
id: CAND-117
verdict: excluded
related_requirements: [FR-021]
confidence: confirmed
---

# CAND-117

## Source

> - [x] On desktop, in the new event form and on the new event page, when I scroll the time zone menu, the width is always 520 and doesn't change based on the content length.

[Source lines 451-451](../../../../backlog/backlog.md#L451-L451)

## Candidate behavior

On desktop, the time-zone menu in the new-event form and page remains 520 wide while scrolling.

## Applicability

Actor: event creator; Location: desktop new-event form and page time-zone menu; Event kind: unspecified; Interaction mode: scroll menu; Viewport: desktop; State: menu open; Exclusions: existing-event pages and mobile.

## Classification

implementation detail

## Existing Requirements and Confidence

Overlap: proposed FR-021 initializes new-event timezone but not menu dimensions. Confidence: confirmed.

## Disposition

Exclude; the fixed menu dimension is an incidental visual implementation.

## Open Questions

Are units pixels, and does this apply to every time-zone selector?
