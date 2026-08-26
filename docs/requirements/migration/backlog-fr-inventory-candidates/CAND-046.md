---
id: CAND-046
verdict: covered
related_requirements:
  - FR-013
confidence: confirmed
---

# CAND-046

## Source

> - [x] Shown in shouldn't affect the time zone

## Candidate behavior

Changing the display timezone should not change the event timezone.

## Applicability

Actor: event visitor.
Location: event-page display controls.
Event kind: timed.
Interaction mode: display-timezone change.
Viewport: any.
State: timezone selected.
Exclusions: explicit event-settings timezone edits.

## Classification

existing requirement

## Existing Requirements and Confidence

FR-013 explicitly preserves the event timezone.
Confidence: confirmed.

## Disposition

Map to FR-013; no migration.

## Open Questions

None.
