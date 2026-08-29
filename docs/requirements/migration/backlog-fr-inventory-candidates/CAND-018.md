---
id: CAND-018
verdict: covered
related_requirements:
  - FR-114
confidence: confirmed
---

# CAND-018

## Source

> - [x] mobile version - switching between 3 days and 7 days doesn't work
>   - Works now

## Candidate behavior

On mobile, an event visitor can switch the event-page view between the selected 3-day and 7-day ranges.

## Applicability

Actor: mobile event visitor.
Location: event page.
Event kind: [Timed Event](../../../terminology/glossary.md#timed-event).
Interaction mode: view-range switching.
Viewport: mobile.
State: 3-day or 7-day selection.
Exclusions: desktop and [Dates-Only Event](../../../terminology/glossary.md#dates-only-event) pages.

## Classification

candidate FR

## Existing Requirements and Confidence

FR-114 specifies the mobile `3 days`/`7 days` range switch, its visibility rule, and its persistence.
Confidence: confirmed.

## Disposition

Specified by FR-114; the switch is hidden when the [Timed Grid](../../../terminology/glossary.md#timed-grid) can display 3 or fewer day columns.

## Open Questions

None.
