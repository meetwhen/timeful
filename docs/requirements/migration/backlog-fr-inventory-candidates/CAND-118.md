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

Overlap: accepted FR-024 keeps display time format independent; no accepted requirement assigns Calendar timezone. Confidence: inferred.

## Disposition

Review as external-calendar integration behavior.

## Open Questions

Does the Calendar event preserve instants while expressing its timezone as display timezone?
