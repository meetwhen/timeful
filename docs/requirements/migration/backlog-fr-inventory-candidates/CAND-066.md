---
id: CAND-066
verdict: proposed-requirement
requirement_type: FR
related_requirements:
  - FR-085
confidence: confirmed
---

# CAND-066

## Source

> - [x] When there's a scheduled event and I edit availability, I shold draw only with green cells, not blue cells (scheduled)

## Candidate behavior

During availability editing for a timed event with a scheduled event time, the grid should not render scheduled-event cells.

## Applicability

Actor: availability editor.
Location: availability editor grid.
Event kind: timed.
Interaction mode: editing availability.
Viewport: any.
State: scheduled event exists.
Exclusions: viewing or scheduling mode.

## Classification

candidate FR

## Existing Requirements and Confidence

None.
Confidence: confirmed.

## Disposition

Promoted to [FR-085](../../functional/fr/FR-085.md).

## Open Questions

None.
