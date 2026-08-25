---
id: CAND-097
verdict: covered
related_requirements: [FR-014]
confidence: confirmed
---

# CAND-097

## Source

> - [x] show all hours should show all hours, not trimmed. Currently, it trims wrong
>   - can't reproduce

## Candidate behavior

No new requirement behavior asserted; the intended outcome is stated but the reported trimming defect was not reproducible.

## Applicability

Actor: event visitor; Location: event-page Show all hours; Event kind: timed; Interaction mode: activate control; Viewport: unspecified; State: defect not reproducible; Exclusions: none established.

## Classification

existing requirement

## Existing Requirements and Confidence

Overlap: accepted FR-014 requires the full civil-day axis when Show all hours is enabled. Confidence: confirmed.

## Disposition

Map to FR-014; retain the unreproduced defect report as provenance.

## Open Questions

What hours were incorrectly trimmed, and is `Show all hours` expected to override every collapsed run?
