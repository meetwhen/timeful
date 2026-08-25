---
id: CAND-103
verdict: excluded
related_requirements: [FR-024]
confidence: confirmed
---

# CAND-103

## Source

> - [x] Use monospace font for times
>   - [x] in the grid
>   - [x] in the tooltip

## Candidate behavior

No new requirement behavior asserted; the source prescribes a presentation implementation for grid and tooltip times.

## Applicability

Actor: event visitor; Location: timed grid and tooltip; Event kind: timed; Interaction mode: viewing; Viewport: unspecified; State: time visible; Exclusions: other time displays not established.

## Classification

implementation detail

## Existing Requirements and Confidence

Overlap: accepted FR-024 defines which time format controls grid and tooltip times, not typography. Confidence: confirmed.

## Disposition

Exclude unless typography is confirmed as a durable accessibility or brand requirement.

## Open Questions

Is monospace required for legibility, and which font family and fallback are intended?
