---
id: CAND-105
verdict: proposed-requirement
requirement_type: FR
related_requirements: [FR-020, FR-088]
confidence: inferred
---

# CAND-105

## Source

> - [x] On mobile, given I'm on the event page and a timeslot is selected and the tooltip is visible, when I click the responses offcanvas, the selected timeslot and its tooltip must not disappear

## Candidate behavior

On mobile, opening the Responses offcanvas preserves an already selected Timed Slot and its visible tooltip.

## Applicability

Actor: event visitor; Location: mobile event page; Event kind: timed; Interaction mode: open Responses offcanvas; Viewport: mobile; State: a timeslot selected and tooltip visible; Exclusions: no selection.

## Classification

candidate FR

## Existing Requirements and Confidence

Overlap: accepted FR-020 preserves selection while scrolling, not when opening Responses.
Confidence: inferred.

## Disposition

Promoted to [FR-088](../../functional/fr/FR-088.md), extending the FR-020 preservation case to opening Responses.
