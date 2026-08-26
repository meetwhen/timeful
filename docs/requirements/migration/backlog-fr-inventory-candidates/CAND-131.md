---
id: CAND-131
verdict: proposed-requirement
requirement_type: FR
related_requirements: [FR-026]
confidence: inferred
---

# CAND-131

## Source

> - [x] In "New event" form, when "Dates only" is selected, in "Advanced options", there should be no "Time increment"

[Source lines 470-470](../../../../backlog/backlog.md#L470-L470)

## Candidate behavior

When Dates only is selected in New event, Advanced options omits Time increment.

## Applicability

Actor: event creator; Location: New event Advanced options; Event kind: dates-only; Interaction mode: choose event kind; Viewport: unspecified; State: Dates only selected; Exclusions: Dates and times.

## Classification

candidate FR

## Existing Requirements and Confidence

Overlap: accepted FR-026 defines Dates only as an event kind but not this option visibility.
Confidence: inferred.

## Disposition

Review as event-kind-specific form behavior.

## Open Questions

Should Time increment be hidden, disabled, or absent from submitted data?
