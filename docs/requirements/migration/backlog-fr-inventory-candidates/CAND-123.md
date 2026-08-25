---
id: CAND-123
verdict: proposed-requirement
requirement_type: FR
related_requirements: [FR-009, FR-014]
confidence: inferred
---

# CAND-123

## Source

> - [x] For dates-only events, the color of disabled dates must be dark-grey like for disabled padding cells in timed event

[Source lines 459-459](../../../../backlog/backlog.md#L459-L459)

## Candidate behavior

Dates-only calendars render disabled dates dark-grey like timed-event disabled padding cells.

## Applicability

Actor: event visitor; Location: dates-only event calendar; Event kind: dates-only; Interaction mode: viewing; Viewport: unspecified; State: disabled date; Exclusions: enabled dates.

## Classification

candidate FR

## Existing Requirements and Confidence

Overlap: accepted FR-009 and FR-014 establish dark-grey timed-grid states, not dates-only calendar styling. Confidence: inferred.

## Disposition

Review as cross-event visual consistency.

## Open Questions

Does dark-grey require an exact color and accessibility contrast threshold?
