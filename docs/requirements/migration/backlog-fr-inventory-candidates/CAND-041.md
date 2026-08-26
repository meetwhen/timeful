---
id: CAND-041
verdict: excluded
related_requirements: []
confidence: confirmed
---

# CAND-041

## Source

> - [x] do we have recurring events?
>       Yes — it uses a custom TimedRecurrence model with two kinds: specific_dates (explicit date list) and weekly (day-of-week pattern).
>       No iCalendar RRULE support.

## Candidate behavior

No new requirement behavior asserted; the child reports an implementation model and exclusion.

## Applicability

Actor: maintainer.
Location: recurrence model.
Event kind: timed.
Interaction mode: configuration.
Viewport: any.
State: recurrence.
Exclusions: iCalendar RRULE support.

## Classification

implementation detail

## Existing Requirements and Confidence

None.
Confidence: confirmed.

## Disposition

Exclude from requirements; retain as implementation and exclusion provenance.

## Open Questions

Is recurrence a supported product behavior and what are its observable rules?
