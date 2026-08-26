---
id: CAND-068
verdict: proposed-requirement
requirement_type: FR
related_requirements:
  - FR-012
  - FR-086
confidence: confirmed
---

# CAND-068

## Source

> - [x] set specific times - just select, don't draw scheduled event

## Candidate behavior

In specific-times mode, an event editor should select times without rendering a scheduled event.

## Applicability

Actor: event editor.
Location: specific-times grid.
Event kind: timed.
Interaction mode: setting specific times.
Viewport: any.
State: selection.
Exclusions: scheduling mode.

## Classification

candidate FR

## Existing Requirements and Confidence

FR-012 defines scheduled event time but not its rendering in specific-times mode.
Confidence: confirmed.

## Disposition

Promoted to [FR-086](../../functional/fr/FR-086.md), which extends FR-012 with specific-times rendering.

## Open Questions

None.
