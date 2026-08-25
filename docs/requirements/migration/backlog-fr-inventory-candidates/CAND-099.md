---
id: CAND-099
verdict: excluded
requirement_type: null
related_requirements: [FR-012]
confidence: confirmed
---

# CAND-099

## Source

<!-- prettier-ignore -->
> - [x] Allow multiple scheduled events.
>
>     Remove clear in Reschedule event because one can just wipe it by dragging the event cursor.
>
>     Should behave just like the availability cursor.
>
>     Given I clicked Reschedule event,
>     when I click a slot that belongs to the existing event, the slot gets cleared
>     and when I click a free slot, it becomes an event slot
>     and adjacent slots merge into a single event
>
>   - Won't implement because it makes the interface much more complicated.
>     If there's 0 slots selected, how to save?
>     If there's several events selected, need to warn that exactly one event is required to schedule on Google Calendar

## Candidate behavior

No new requirement behavior asserted; the source records a decision not to implement multiple scheduled events.

## Applicability

Actor: event visitor; Location: Reschedule event; Event kind: timed; Interaction mode: slot selection; Viewport: unspecified; State: multiple scheduled events proposed; Exclusions: single scheduled event.

## Classification

ADR or decision

## Existing Requirements and Confidence

Overlap: accepted FR-012 allows one optional Scheduled Event Time; it does not authorize multiple events. Confidence: confirmed.

## Disposition

Record as a non-normative product decision; do not derive multiple-event behavior.
