---
id: CAND-067
verdict: proposed-requirement
requirement_type: FR
related_requirements:
  - FR-005
confidence: inferred
---

# CAND-067

## Source

> - [x] when editing availability, overlay availabilities should preserve the marked time slots and not shift

## Candidate behavior

The edited availability overlay should remain aligned with its marked timed slots during availability editing.

## Applicability

Actor: availability editor.
Location: availability editor grid.
Event kind: timed.
Interaction mode: editing availability.
Viewport: any.
State: overlay enabled.
Exclusions: non-overlay rendering.

## Classification

candidate FR

## Existing Requirements and Confidence

FR-005 requires the edited response overlay but does not require its marked slots to remain aligned.
Confidence: inferred.

## Disposition

Candidate for a canonical FR defining overlay alignment during availability editing.

## Open Questions

None.
