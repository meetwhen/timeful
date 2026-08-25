---
id: CAND-089
verdict: proposed-requirement
requirement_type: FR
related_requirements: [FR-020]
confidence: inferred
---

# CAND-089

## Source

> - [x] on mobile, tooltip should be below the Responses offcanvas panel

## Candidate behavior

On mobile, the grid tooltip is layered below the Responses offcanvas panel.

## Applicability

Actor: event visitor; Location: mobile event page; Event kind: timed; Interaction mode: viewing a selected slot; Viewport: mobile; State: Responses offcanvas open; Exclusions: no panel or desktop not established.

## Classification

candidate FR

## Existing Requirements and Confidence

Overlap: accepted FR-020 requires a selected slot to remain selected while scrolling, but does not set panel layering. Confidence: inferred.

## Disposition

Review as a mobile overlay-layering behavior.

## Open Questions

Should the tooltip remain selected but be visually obscured, or be repositioned?
