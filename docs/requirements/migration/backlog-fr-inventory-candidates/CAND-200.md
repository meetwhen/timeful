---
id: CAND-200
verdict: proposed-requirement
requirement_type: FR
related_requirements:
  - FR-044
confidence: inferred
---

# CAND-200: Show Response Editability

## Source

> If the response can be edited, there's a pencil on the right. Otherwise, a lock.

## Candidate behavior

On the event page, an editable [Availability Response](../../../terminology/glossary.md#availability-response) shows a pencil action on its right, while a non-editable response shows a lock.

## Applicability

Actor: event visitor. Location: event-page response list. Event kind: unspecified. Interaction mode: viewing responses. Viewport: unspecified. State: response is editable or non-editable by the current visitor. Exclusions: availability-editor controls and the reason a response is non-editable.

## Classification

candidate FR

## Existing Requirements and Confidence

[FR-044](../../functional/fr/FR-044.md) specifies an access icon for a protected response and its explanatory tooltip, but not the editable pencil, lock glyph, or right-side placement. Confidence: inferred.

## Disposition

Hold as an event-page response-list candidate related to FR-044.

## Open Questions

Does selecting the lock explain why the response is not editable?
