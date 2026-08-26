---
id: CAND-205
title: New-Event Form Month-Change Defect
verdict: excluded
related_requirements: []
confidence: confirmed
---

# CAND-205: New-Event Form Month-Change Defect

## Source

> In the new event form, when the event name isn't set and when the user changes the month, the form scrolls to the top adn requires the event name to remind the user to set it.
>
> this is a bug, not an fr

## Candidate behavior

No new requirement behavior asserted; the source identifies a new-event form defect rather than a durable product requirement.

## Applicability

Actor: event creator.
Location: new-event form.
Event kind: unspecified.
Interaction mode: changing the month.
Viewport: unspecified.
State: event name is unset.
Exclusions: the general scroll-preservation behavior in CAND-204.

## Classification

bug or investigation

## Existing Requirements and Confidence

No accepted or proposed FR or QR establishes the reported event-name prompt behavior.
Confidence: confirmed.

## Disposition

Excluded; retain as defect provenance and do not create an FR or QR.
