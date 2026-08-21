# Functional Requirements

## FR-004

If the response can be edited, there's a pencil on the right. Otherwise, a lock.

## FR-005

Button styles on the event page follow Material Design:

- `Add availability`: filled primary, subtle shadow if needed for separation, no persistent glow
- `Edit availability`: filled primary, calmer than add, no glow
- `Edit event` and `Copy link`: outlined, no shadow

## FR-006

Setting "Shown in" shouldn't affect the initial event time zone.

## FR-007

When scheduling an event, it can't be empty.

## FR-009

In the new event form, when I click a button near the month or the year, the form should stay in place and not scroll to the top.

## FR-010

In the new event form, when the event name isn't set and when the user changes the month, the form scrolls to the top adn requires the event name to remind the user to set it.

## FR-012

When there are responses but no responses to edit, the user should see disabled Edit availability button.

## FR-014

Dates picked in the date picker shall be the source of truth for enabled time slots.

## FR-016

When scheduling an event, the tooltip with the info about the time slot should follow the mouse cursor and not be above the slot where scheduling the event started

## FR-019

Collapsed hours rectangle height should be the same as half-hour line.

## FR-020

Each full-hour line in the grid should have a label on the left, including top of the collapsed hours rectangle.

## FR-022

Time-range picker labels shall represent the end-of-day boundary clearly:

- In 24-hour mode, labels shall be zero-padded from `00:00` through `23:00`, followed by `24:00`.
- In 12-hour mode, the end-of-day option shall be `12 AM`.
- Selecting the end-of-day boundary shall render as `00:00` in the next date column of the grid.
