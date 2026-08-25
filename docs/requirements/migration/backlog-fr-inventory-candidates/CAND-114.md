---
id: CAND-114
verdict: proposed-requirement
requirement_type: FR
related_requirements: [FR-045]
confidence: inferred
---

# CAND-114

## Source

> - [x] In responses,
>   - when I hover or click in a grid, show the square for the status of a person at the corresponding timeslot (available, if needed, unavailable, etc.) instead of a profile image;
>   - when I hover or click outside of the grid, replace the status square with a profile icon like it's now

[Source lines 446-448](../../../../backlog/backlog.md#L446-L448)

## Candidate behavior

Responses show a per-slot status square for grid interaction and a profile icon outside the grid.

## Applicability

Actor: event visitor; Location: Responses; Event kind: timed; Interaction mode: hover or click inside or outside grid; Viewport: unspecified; State: corresponding slot selected or no grid interaction; Exclusions: dates-only events not established.

## Classification

candidate FR

## Existing Requirements and Confidence

Overlap: proposed FR-045 covers unavailable response state but does not prescribe response icons. Confidence: inferred.

## Disposition

Review as response-status presentation behavior.

## Open Questions

Which statuses require squares, and what restores the profile icon after leaving the grid?
