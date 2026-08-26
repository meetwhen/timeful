---
id: CAND-018
verdict: proposed-requirement
requirement_type: FR
related_requirements:
  - FR-049
confidence: inferred
---

# CAND-018

## Source

> - [x] mobile version - switching between 3 days and 7 days doesn't work
>   - Works now

## Candidate behavior

On mobile, an event visitor can switch the event-page view between the selected 3-day and 7-day ranges.

## Applicability

Actor: mobile event visitor.
Location: event page.
Event kind: unconfirmed.
Interaction mode: view-range switching.
Viewport: mobile.
State: 3-day or 7-day selection.
Exclusions: desktop.

## Classification

candidate FR

## Existing Requirements and Confidence

FR-049 requires the number-of-days control but does not define range switching.
Confidence: inferred.

## Disposition

Retain as a candidate FR; FR-049 does not define switching between selected ranges.

## Open Questions

What control, event kinds, and persistence behavior define the range switch?
