# CAND-059

## Source

<!-- prettier-ignore -->
> - [x] the day ends at 24
>    No, at `00:00` for consistency with 12AM in 12-hour format

## Candidate behavior

No new requirement behavior asserted; the child records time-label formatting choice.

## Applicability

Actor: event visitor. Location: timed-grid axis. Event kind: timed. Interaction mode: viewing. Viewport: any. State: day boundary. Exclusions: time-format persistence.

## Classification

needs product decision

## Existing Requirements and Confidence

FR-024 separates event and display time formats. Confidence: inferred.

## Disposition

Retain as presentation decision.

## Open Questions

Does the `00:00` decision apply in both display formats and editor forms?
