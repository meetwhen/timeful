---
id: CAND-101
verdict: proposed-requirement
requirement_type: FR
related_requirements: [FR-016]
confidence: inferred
---

# CAND-101

## Source

> - [x] On desktop, when rescheduling event, add availability, edit availability, edit event should be disabled

## Candidate behavior

While rescheduling an event on desktop, Add availability, Edit availability, and Edit event are disabled.

## Applicability

Actor: event visitor; Location: desktop event page; Event kind: timed; Interaction mode: rescheduling; Viewport: desktop; State: rescheduling active; Exclusions: mobile not established.

## Classification

candidate FR

## Existing Requirements and Confidence

Overlap: proposed FR-016 limits rescheduling to Active Slots but does not govern concurrent controls. Confidence: inferred.

## Disposition

Review as mode-exclusivity behavior.

## Open Questions

Are controls disabled, hidden, or blocked after activation, and should the same rule apply on mobile?
