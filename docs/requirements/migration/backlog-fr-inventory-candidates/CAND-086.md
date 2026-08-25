---
id: CAND-086
verdict: proposed-requirement
requirement_type: FR
related_requirements: [FR-009, FR-014]
confidence: inferred
---

# CAND-086

## Source

> - [x] Disabled padding cells - "Unavailable, outside the event dates in the event timezone"

## Candidate behavior

Disabled padding cells expose the text `Unavailable, outside the event dates in the event timezone`.

## Applicability

Actor: event visitor; Location: timed grid; Event kind: timed; Interaction mode: inspecting a disabled padding cell; Viewport: unspecified; State: outside event dates; Exclusions: non-padding disabled cells not established.

## Classification

candidate FR

## Existing Requirements and Confidence

Overlap: accepted FR-009 and FR-014 define disabled padding cells but not this text. Confidence: inferred.

## Disposition

Review as a state-description requirement.

## Open Questions

Is the quoted text visible, a tooltip, or an accessible name?
