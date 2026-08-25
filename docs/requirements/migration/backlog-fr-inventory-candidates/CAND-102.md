# CAND-102

## Source

<!-- prettier-ignore -->
> - [x] On desktop, Shown in and timezone should be above Responses and scroll with it so that when all timeslots are shown
>      and one selects a timeslot and see the tooltip with the time, they can see which timezone that time belongs to

## Candidate behavior

On desktop, Shown in and timezone remain above Responses while scrolling so a selected timeslot's timezone is visible.

## Applicability

Actor: event visitor; Location: desktop event-page Responses; Event kind: timed; Interaction mode: scroll and select slot; Viewport: desktop; State: all timeslots shown; Exclusions: mobile not established.

## Classification

candidate FR

## Existing Requirements and Confidence

Overlap: accepted FR-047 labels the Display Timezone control but does not set its position. Confidence: inferred.

## Disposition

Review as sidebar-positioning behavior.

## Open Questions

Does `Shown in and timezone` name one control or two, and should it remain sticky?
