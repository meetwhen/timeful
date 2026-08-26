---
id: CAND-091
verdict: proposed-requirement
requirement_type: FR
related_requirements: [FR-045]
confidence: inferred
---

# CAND-091

## Source

> - [x] Given I'm on the event page, when I move the mouse cursor out of the grid, the Responses must show just the number of responses and not show who's available

## Candidate behavior

When the pointer leaves the grid, Responses shows only the response count and not individual availability.

## Applicability

Actor: event visitor; Location: event page Responses; Event kind: timed; Interaction mode: pointer leaves grid; Viewport: desktop; State: grid not hovered; Exclusions: selected slots not established.

## Classification

candidate FR

## Existing Requirements and Confidence

Overlap: proposed FR-045 defines Responses while hovering unavailable states but not when the pointer leaves the grid.
Confidence: inferred.

## Disposition

Review as response-summary behavior.

## Open Questions

Does this apply after clicking a slot or on touch devices?
