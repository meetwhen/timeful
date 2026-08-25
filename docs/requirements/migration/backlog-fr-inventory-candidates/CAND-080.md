---
id: CAND-080
verdict: proposed-requirement
requirement_type: FR
related_requirements:
  - FR-048
confidence: inferred
---

# CAND-080

## Source

<!-- prettier-ignore -->
> - [x] On the event page without responses, there should be only Add availability and Show all hours, not very wide Add availability and More options
>
>       "Show all hours" should be under "Add availability", as before.
>       The toggle and "Show all hours" text should be centered vertically within their box

## Candidate behavior

On a timed event page with no responses, the event page should show Add availability and Show all hours with Show all hours below Add availability.

## Applicability

Actor: event visitor. Location: event page. Event kind: timed. Interaction mode: viewing. Viewport: unconfirmed. State: no responses. Exclusions: responses present.

## Classification

candidate FR

## Existing Requirements and Confidence

FR-048 covers centering the no-response mobile Show all hours control, but no canonical requirement covers the control set or ordering. Confidence: inferred.

## Disposition

Candidate for a canonical FR defining no-response event-page controls and their arrangement.

## Open Questions

None.
