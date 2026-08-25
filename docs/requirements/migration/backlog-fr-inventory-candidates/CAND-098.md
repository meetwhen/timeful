---
id: CAND-098
verdict: covered
related_requirements: [FR-004]
confidence: confirmed
---

# CAND-098

## Source

> - [x] Given on the edit availability page, when no timeslot is marked as available/if needed, then the Save button should be disabled

## Candidate behavior

No new requirement behavior asserted; accepted FR-004 already requires a non-empty Availability Response before saving.

## Applicability

Actor: availability editor; Location: edit availability page; Event kind: timed; Interaction mode: editing; Viewport: unspecified; State: no Available or If needed slot; Exclusions: add availability not explicit.

## Classification

existing requirement

## Existing Requirements and Confidence

Overlap: accepted FR-004 directly covers the save precondition. Confidence: confirmed.

## Disposition

Map to FR-004; retain the disabled-button presentation only if separately needed.
