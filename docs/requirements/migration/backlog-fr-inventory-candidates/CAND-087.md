---
id: CAND-087
verdict: proposed-requirement
requirement_type: FR
related_requirements: [FR-043]
confidence: inferred
---

# CAND-087

## Source

> - [x] In desktop app, in firefox, the selected timeslot must follow the mouse when it moves inside the grid, no matter clicks.

## Candidate behavior

In desktop Firefox, moving the pointer within the grid updates the selected timeslot regardless of prior clicks.

## Applicability

Actor: event visitor; Location: desktop timed grid; Event kind: timed; Interaction mode: pointer hover; Viewport: desktop; State: Firefox; Exclusions: other browsers not established.

## Classification

candidate FR

## Existing Requirements and Confidence

Overlap: proposed FR-043 concerns tooltip placement at a hovered slot but does not require the selected slot to follow pointer movement. Confidence: inferred.

## Disposition

Review as a browser-scoped selection behavior.

## Open Questions

Is Firefox a support constraint or only the reproduction environment?
