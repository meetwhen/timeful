# CAND-097

## Source

> - [x] show all hours should show all hours, not trimmed. Currently, it trims wrong
>   - can't reproduce

## Candidate behavior

No new requirement behavior asserted; the intended outcome is stated but the reported trimming defect was not reproducible.

## Applicability

Actor: event visitor; Location: event-page Show all hours; Event kind: timed; Interaction mode: activate control; Viewport: unspecified; State: defect not reproducible; Exclusions: none established.

## Classification

bug or investigation

## Existing Requirements and Confidence

Overlap: accepted FR-011 collapses inactive timed-grid runs but does not define this control. Confidence: needs product decision.

## Disposition

Investigate before migration.

## Open Questions

What hours were incorrectly trimmed, and is `Show all hours` expected to override every collapsed run?
