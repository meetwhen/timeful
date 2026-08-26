---
id: CAND-063
verdict: covered
related_requirements:
  - FR-011
confidence: confirmed
---

# CAND-063

## Source

> - [x] don't uncollapse rows when scheduling

## Candidate behavior

No new requirement behavior asserted; accepted behavior already preserves collapsed bands during scheduling.

## Applicability

Actor: event visitor.
Location: timed grid.
Event kind: timed.
Interaction mode: scheduling.
Viewport: any.
State: collapsed inactive runs.
Exclusions: availability editing and specific-times setting.

## Classification

existing requirement

## Existing Requirements and Confidence

FR-011 explicitly says scheduling does not expand collapsed bands.
Confidence: confirmed.

## Disposition

Map to FR-011; no migration.

## Open Questions

None.
