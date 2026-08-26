---
id: CAND-129
verdict: proposed-requirement
requirement_type: FR
related_requirements: [FR-009, FR-014]
confidence: inferred
---

# CAND-129

## Source

> - [x] In timed range events, the Legend shall show the light-gray "Disabled, change in Edit event"

[Source lines 468-468](../../../../backlog/backlog.md#L468-L468)

## Candidate behavior

Timed range event legends show a light-gray `Disabled, change in Edit event` item.

## Applicability

Actor: event visitor; Location: timed range event legend; Event kind: timed range; Interaction mode: viewing; Viewport: unspecified; State: legend displayed; Exclusions: custom-domain and dates-only events not established.

## Classification

candidate FR

## Existing Requirements and Confidence

Overlap: accepted FR-009 and FR-014 define timed-grid states, but neither requires this legend guidance.
Confidence: inferred.

## Disposition

Review alongside CAND-115 for a coherent disabled-state legend.

## Open Questions

Is `change in Edit event` available to non-owners, and how does this item differ from Disabled, collapsed?
