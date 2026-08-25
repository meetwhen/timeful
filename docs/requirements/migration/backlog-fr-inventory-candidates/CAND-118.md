---
id: CAND-118
verdict: proposed-requirement
requirement_type: FR
related_requirements: [FR-012, FR-013]
confidence: inferred
---

# CAND-118

## Source

> - [x] Schedule event on Google Calendar should happen in the display timezone

[Source lines 452-452](../../../../backlog/backlog.md#L452-L452)

## Candidate behavior

Scheduling an event on Google Calendar uses the display timezone.

## Applicability

Actor: event visitor; Location: Google Calendar scheduling; Event kind: timed; Interaction mode: schedule event; Viewport: unspecified; State: display timezone set; Exclusions: dates-only events not established.

## Classification

candidate FR

## Existing Requirements and Confidence

Overlap: accepted FR-012 defines Scheduled Event Time and accepted FR-013 preserves Timed Slot instants across display-timezone changes, but neither assigns a Calendar timezone. Confidence: inferred.

## Disposition

Review as external-calendar integration behavior.

## Open Questions

Does the Calendar event preserve instants while expressing its timezone as display timezone?
