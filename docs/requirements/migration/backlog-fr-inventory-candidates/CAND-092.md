---
id: CAND-092
verdict: needs-decision
related_requirements: [FR-035]
confidence: needs-product-decision
---

# CAND-092

## Source

<!-- prettier-ignore -->
> - [x] In the desktop version, the alignment of rows in add/edit availability should be similar to the event page.
>
>       - event title - Cancel, Save
>       - Edit event - Show all hours
>       - Add description - shouldn't be visible, actually

## Candidate behavior

On desktop availability editing, rows align with the event page and do not show Add description.

## Applicability

Actor: event visitor; Location: desktop availability editing; Event kind: unspecified; Interaction mode: add or edit availability; Viewport: desktop; State: editing; Exclusions: mobile not established.

## Classification

needs product decision

## Existing Requirements and Confidence

Overlap: proposed FR-035 permits an event description during event creation but does not establish the availability-editing control.
Confidence: needs product decision.

## Disposition

Needs product decision on the alignment and Add description control meanings before migration.

## Open Questions

Are the two listed row pairings mandatory, and does `Add description` refer to an availability or event control?
