---
id: CAND-112
verdict: proposed-requirement
requirement_type: FR
related_requirements: []
confidence: inferred
---

# CAND-112

## Source

> - [x] On desktop, given I'm on the event page, when I hover or click inside the grid outside active cells, the highlight at the last timeslot and any tooltip should not be visible

[Source lines 444-444](../../../../backlog/backlog.md#L444-L444)

## Candidate behavior

On desktop, hovering or clicking outside active grid cells clears the prior highlight and hides its tooltip.

## Applicability

Actor: event visitor; Location: desktop timed grid; Event kind: timed; Interaction mode: hover or click; Viewport: desktop; State: outside active cells; Exclusions: active cells.

## Classification

candidate FR

## Existing Requirements and Confidence

None. Confidence: inferred.

## Disposition

Review as desktop inactive-cell interaction behavior.

## Open Questions

Does `outside active cells` include inter-grid space and collapsed hours?
