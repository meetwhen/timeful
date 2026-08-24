# Event-Page Interaction Review Inventory

This is a durable batch artifact and is non-normative. `CAND-*` IDs are temporary review IDs; this batch will be consolidated by [../backlog-fr-inventory.md](../backlog-fr-inventory.md).

## CAND-081

#### Source

> - [x] "Best times" toggle should appear when there is at least one response, not more than one

[Source lines 381-381](../../../../backlog/backlog.md#L381-L381)

#### Candidate behavior

With at least one response, the event page shows the `Best times` toggle.

#### Applicability

Actor: event visitor; Location: event page; Event kind: unspecified; Interaction mode: viewing; Viewport: unspecified; State: at least one response; Exclusions: none established.

#### Classification

candidate FR

#### Existing Requirements and Confidence

Overlap: none in accepted FR/QR records. Confidence: inferred.

#### Disposition

Review as a candidate functional requirement.

#### Open Questions

Does `Best times` apply to dates-only events, where the corresponding control may be `Show best days`?

## CAND-082

#### Source

> - [x] When Best times and More options are both visible, they should be side by side

[Source lines 382-382](../../../../backlog/backlog.md#L382-L382)

#### Candidate behavior

When both controls are visible, the event page places `Best times` and `More options` side by side.

#### Applicability

Actor: event visitor; Location: event page controls; Event kind: unspecified; Interaction mode: viewing; Viewport: unspecified; State: both controls visible; Exclusions: none established.

#### Classification

candidate FR

#### Existing Requirements and Confidence

Overlap: accepted FR-049 covers mobile control rows but not this arrangement generally. Confidence: inferred.

#### Disposition

Review as a refinement of control-layout behavior.

#### Open Questions

Is the intended viewport mobile only, and does `Best times` include `Show best days`?

## CAND-083

#### Source

> - [x] On mobile, on the event page, when swiping inside the grid, the grid should scroll just like when swiping outside

[Source lines 383-383](../../../../backlog/backlog.md#L383-L383)

#### Candidate behavior

On mobile, a swipe inside the grid scrolls the event page as a swipe outside the grid does.

#### Applicability

Actor: event visitor; Location: mobile event-page grid; Event kind: unspecified; Interaction mode: touch swipe; Viewport: mobile; State: viewing; Exclusions: none established.

#### Classification

candidate QR

#### Existing Requirements and Confidence

Overlap: none in accepted FR/QR records. Confidence: inferred.

#### Disposition

Review as an interaction-operability quality outcome.

#### Open Questions

Does the gesture ever select or edit a slot, and what scroll target is intended?

## CAND-084

#### Source

> - [x] "select in Add/Edit availability" -> "change in ..."

[Source lines 384-384](../../../../backlog/backlog.md#L384-L384)

#### Candidate behavior

Availability-editing copy uses `change in ...` instead of `select in Add/Edit availability`.

#### Applicability

Actor: event visitor; Location: availability-editing copy; Event kind: unspecified; Interaction mode: viewing; Viewport: unspecified; State: unspecified; Exclusions: none established.

#### Classification

candidate FR

#### Existing Requirements and Confidence

Overlap: none in accepted FR/QR records. Confidence: needs product decision.

#### Disposition

Hold for exact replacement copy and location.

#### Open Questions

What is the complete replacement phrase, and which label or instruction contains it?

## CAND-085

#### Source

> - [x] On mobile, on the event page, when not interacting with the grid, make it scrollable with finger on the grid

[Source lines 385-385](../../../../backlog/backlog.md#L385-L385)

#### Candidate behavior

On mobile, touching the grid while not interacting with it allows event-page scrolling.

#### Applicability

Actor: event visitor; Location: mobile event-page grid; Event kind: unspecified; Interaction mode: touch scroll; Viewport: mobile; State: not interacting with grid; Exclusions: grid interactions not established.

#### Classification

duplicate or refinement

#### Existing Requirements and Confidence

Overlap: CAND-083 is the same touch-scroll concern; no accepted FR/QR overlap. Confidence: inferred.

#### Disposition

Consolidate with CAND-083.

#### Open Questions

Is this intentionally distinct from swiping inside the grid?

## CAND-086

#### Source

> - [x] Disabled padding cells - "Unavailable, outside the event dates in the event timezone"

[Source lines 386-386](../../../../backlog/backlog.md#L386-L386)

#### Candidate behavior

Disabled padding cells expose the text `Unavailable, outside the event dates in the event timezone`.

#### Applicability

Actor: event visitor; Location: timed grid; Event kind: timed; Interaction mode: inspecting a disabled padding cell; Viewport: unspecified; State: outside event dates; Exclusions: non-padding disabled cells not established.

#### Classification

candidate FR

#### Existing Requirements and Confidence

Overlap: none in accepted FR/QR records. Confidence: inferred.

#### Disposition

Review as a state-description requirement.

#### Open Questions

Is the quoted text visible, a tooltip, or an accessible name?

## CAND-087

#### Source

> - [x] In desktop app, in firefox, the selected timeslot must follow the mouse when it moves inside the grid, no matter clicks.

[Source lines 387-387](../../../../backlog/backlog.md#L387-L387)

#### Candidate behavior

In desktop Firefox, moving the pointer within the grid updates the selected timeslot regardless of prior clicks.

#### Applicability

Actor: event visitor; Location: desktop timed grid; Event kind: timed; Interaction mode: pointer hover; Viewport: desktop; State: Firefox; Exclusions: other browsers not established.

#### Classification

candidate FR

#### Existing Requirements and Confidence

Overlap: proposed FR-043 places the tooltip at a hovered slot; no accepted FR/QR overlap. Confidence: inferred.

#### Disposition

Review as a browser-scoped selection behavior.

#### Open Questions

Is Firefox a support constraint or only the reproduction environment?

## CAND-088

#### Source

> - [x] doesn't show time on hover after clicking and moving the cursor - <http://127.0.0.1:4173/e/dEeaF>
>   - Can't reproduce

[Source lines 388-389](../../../../backlog/backlog.md#L388-L389)

#### Candidate behavior

No new requirement behavior asserted; the source is an unreproduced hover-time defect report.

#### Applicability

Actor: event visitor; Location: event-page grid; Event kind: unspecified; Interaction mode: click then pointer movement; Viewport: unspecified; State: issue not reproducible; Exclusions: none established.

#### Classification

bug or investigation

#### Existing Requirements and Confidence

Overlap: proposed FR-043 concerns hover tooltip placement; no accepted FR/QR overlap. Confidence: needs product decision.

#### Disposition

Do not migrate unless the defect is reproduced and intended behavior is confirmed.

#### Open Questions

What event kind, browser, and exact expected time display were involved?

## CAND-089

#### Source

> - [x] on mobile, tooltip should be below the Responses offcanvas panel

[Source lines 390-390](../../../../backlog/backlog.md#L390-L390)

#### Candidate behavior

On mobile, the grid tooltip is layered below the Responses offcanvas panel.

#### Applicability

Actor: event visitor; Location: mobile event page; Event kind: timed; Interaction mode: viewing a selected slot; Viewport: mobile; State: Responses offcanvas open; Exclusions: no panel or desktop not established.

#### Classification

candidate FR

#### Existing Requirements and Confidence

Overlap: accepted FR-020 requires a selected slot to remain selected while scrolling, but does not set panel layering. Confidence: inferred.

#### Disposition

Review as a mobile overlay-layering behavior.

#### Open Questions

Should the tooltip remain selected but be visually obscured, or be repositioned?

## CAND-090

#### Source

> - [x] get rid of the comparator, leave just inspect and update docs

[Source lines 391-391](../../../../backlog/backlog.md#L391-L391)

#### Candidate behavior

No new requirement behavior asserted; this is an implementation and documentation instruction.

#### Applicability

Actor: developer; Location: unspecified comparator; Event kind: not applicable; Interaction mode: implementation; Viewport: not applicable; State: unspecified; Exclusions: product behavior.

#### Classification

implementation detail

#### Existing Requirements and Confidence

Overlap: none in accepted FR/QR records. Confidence: inferred.

#### Disposition

Exclude from requirement migration.

#### Open Questions

Which comparator and documentation were meant?

## CAND-091

#### Source

> - [x] Given I'm on the event page, when I move the mouse cursor out of the grid, the Responses must show just the number of responses and not show who's available

[Source lines 392-392](../../../../backlog/backlog.md#L392-L392)

#### Candidate behavior

When the pointer leaves the grid, Responses shows only the response count and not individual availability.

#### Applicability

Actor: event visitor; Location: event page Responses; Event kind: timed; Interaction mode: pointer leaves grid; Viewport: desktop; State: grid not hovered; Exclusions: selected slots not established.

#### Classification

candidate FR

#### Existing Requirements and Confidence

Overlap: none in accepted FR/QR records. Confidence: inferred.

#### Disposition

Review as response-summary behavior.

#### Open Questions

Does this apply after clicking a slot or on touch devices?

## CAND-092

#### Source

> - [x] In the desktop version, the alignment of rows in add/edit availability should be similar to the event page.
>
>       - event title - Cancel, Save
>       - Edit event - Show all hours
>       - Add description - shouldn't be visible, actually

[Source lines 393-397](../../../../backlog/backlog.md#L393-L397)

#### Candidate behavior

On desktop availability editing, rows align with the event page and do not show Add description.

#### Applicability

Actor: event visitor; Location: desktop availability editing; Event kind: unspecified; Interaction mode: add or edit availability; Viewport: desktop; State: editing; Exclusions: mobile not established.

#### Classification

candidate FR

#### Existing Requirements and Confidence

Overlap: none in accepted FR/QR records. Confidence: inferred.

#### Disposition

Split alignment and visibility concerns if promoted.

#### Open Questions

Are the two listed row pairings mandatory, and does `Add description` refer to an availability or event control?

## CAND-093

#### Source

> - [x] make scheduled event color #76AFF2 so that it's visible on the green background

[Source lines 398-398](../../../../backlog/backlog.md#L398-L398)

#### Candidate behavior

A scheduled event is rendered in `#76AFF2` on the green background.

#### Applicability

Actor: event visitor; Location: event-page grid; Event kind: timed; Interaction mode: viewing; Viewport: unspecified; State: scheduled event visible; Exclusions: other backgrounds not established.

#### Classification

candidate FR

#### Existing Requirements and Confidence

Overlap: accepted FR-012 defines scheduled-event time, not its visual treatment. Confidence: inferred.

#### Disposition

Review as a visual-state requirement.

#### Open Questions

Is the exact color a durable product constraint or an implementation token?

## CAND-094

#### Source

> - [x] Given on the event page, when scrolled down, then clicked Edit event, then the navbar (timeful, create an event, etc.) must not move higher

[Source lines 399-399](../../../../backlog/backlog.md#L399-L399)

#### Candidate behavior

Opening Edit event from a scrolled event page does not move the navbar higher.

#### Applicability

Actor: event visitor; Location: event page and navbar; Event kind: unspecified; Interaction mode: open Edit event; Viewport: unspecified; State: page scrolled down; Exclusions: other dialogs not established.

#### Classification

candidate FR

#### Existing Requirements and Confidence

Overlap: none in accepted FR/QR records. Confidence: inferred.

#### Disposition

Review as a scroll-position behavior.

#### Open Questions

What navbar position is the baseline, and does the requirement concern layout or scroll offset?

## CAND-095

#### Source

<!-- prettier-ignore -->
> - [x] On the event page, align:
>       - event title with Add availability/Edit availability
>       - "Edit event" and Copy link with Show best times and More options
>       - The event description with Schedule/Reschedule event

[Source lines 400-403](../../../../backlog/backlog.md#L400-L403)

#### Candidate behavior

The event page aligns each listed content pair on the same horizontal axis.

#### Applicability

Actor: event visitor; Location: event page; Event kind: unspecified; Interaction mode: viewing; Viewport: unspecified; State: listed controls visible; Exclusions: none established.

#### Classification

candidate FR

#### Existing Requirements and Confidence

Overlap: accepted FR-049 governs a subset of mobile control layout. Confidence: inferred.

#### Disposition

Review as page-layout behavior.

#### Open Questions

Which viewport and responsive breakpoints are intended?

## CAND-096

#### Source

> - [x] When changing availability, there should be these buttons:
>
>   - Cancel | Save
>   - Overlay availability | More options (Show all hours)
>   - Delete

[Source lines 404-408](../../../../backlog/backlog.md#L404-L408)

#### Candidate behavior

Availability editing exposes Cancel, Save, Overlay availability, More options with Show all hours, and Delete controls.

#### Applicability

Actor: availability editor; Location: availability editing; Event kind: unspecified; Interaction mode: changing availability; Viewport: unspecified; State: editing an existing or new response not established; Exclusions: none established.

#### Classification

candidate FR

#### Existing Requirements and Confidence

Overlap: accepted FR-003 covers deleting responses; accepted FR-004 constrains Save. Confidence: inferred.

#### Disposition

Review each control separately against established editing modes.

#### Open Questions

Which controls apply to Add availability versus Edit availability, and what does Overlay availability do?

## CAND-097

#### Source

> - [x] show all hours should show all hours, not trimmed. Currently, it trims wrong
>   - can't reproduce

[Source lines 409-410](../../../../backlog/backlog.md#L409-L410)

#### Candidate behavior

No new requirement behavior asserted; the intended outcome is stated but the reported trimming defect was not reproducible.

#### Applicability

Actor: event visitor; Location: event-page Show all hours; Event kind: timed; Interaction mode: activate control; Viewport: unspecified; State: defect not reproducible; Exclusions: none established.

#### Classification

bug or investigation

#### Existing Requirements and Confidence

Overlap: accepted FR-011 collapses inactive timed-grid runs but does not define this control. Confidence: needs product decision.

#### Disposition

Investigate before migration.

#### Open Questions

What hours were incorrectly trimmed, and is `Show all hours` expected to override every collapsed run?

## CAND-098

#### Source

> - [x] Given on the edit availability page, when no timeslot is marked as available/if needed, then the Save button should be disabled

[Source lines 411-411](../../../../backlog/backlog.md#L411-L411)

#### Candidate behavior

No new requirement behavior asserted; accepted FR-004 already requires a non-empty Availability Response before saving.

#### Applicability

Actor: availability editor; Location: edit availability page; Event kind: timed; Interaction mode: editing; Viewport: unspecified; State: no Available or If needed slot; Exclusions: add availability not explicit.

#### Classification

existing requirement

#### Existing Requirements and Confidence

Overlap: accepted FR-004 directly covers the save precondition. Confidence: confirmed.

#### Disposition

Map to FR-004; retain the disabled-button presentation only if separately needed.

## CAND-099

#### Source

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

[Source lines 412-425](../../../../backlog/backlog.md#L412-L425)

#### Candidate behavior

No new requirement behavior asserted; the source records a decision not to implement multiple scheduled events.

#### Applicability

Actor: event visitor; Location: Reschedule event; Event kind: timed; Interaction mode: slot selection; Viewport: unspecified; State: multiple scheduled events proposed; Exclusions: single scheduled event.

#### Classification

ADR or decision

#### Existing Requirements and Confidence

Overlap: accepted FR-012 allows one optional Scheduled Event Time; it does not authorize multiple events. Confidence: confirmed.

#### Disposition

Record as a non-normative product decision; do not derive multiple-event behavior.

## CAND-100

#### Source

> - [x] On desktop, the width of buttons when adding availability should be the same as on the event page

[Source lines 426-426](../../../../backlog/backlog.md#L426-L426)

#### Candidate behavior

On desktop, Add availability buttons match the width of their event-page counterparts.

#### Applicability

Actor: event visitor; Location: add availability and event page; Event kind: unspecified; Interaction mode: viewing; Viewport: desktop; State: adding availability; Exclusions: mobile not established.

#### Classification

candidate FR

#### Existing Requirements and Confidence

Overlap: none in accepted FR/QR records. Confidence: inferred.

#### Disposition

Review as responsive visual consistency.

#### Open Questions

Which buttons are counterparts, and does equal width mean exact pixels or aligned columns?

## CAND-101

#### Source

> - [x] On desktop, when rescheduling event, add availability, edit availability, edit event should be disabled

[Source lines 427-427](../../../../backlog/backlog.md#L427-L427)

#### Candidate behavior

While rescheduling an event on desktop, Add availability, Edit availability, and Edit event are disabled.

#### Applicability

Actor: event visitor; Location: desktop event page; Event kind: timed; Interaction mode: rescheduling; Viewport: desktop; State: rescheduling active; Exclusions: mobile not established.

#### Classification

candidate FR

#### Existing Requirements and Confidence

Overlap: accepted FR-012 defines scheduled-event persistence, not concurrent editing. Confidence: inferred.

#### Disposition

Review as mode-exclusivity behavior.

#### Open Questions

Are controls disabled, hidden, or blocked after activation, and should the same rule apply on mobile?

## CAND-102

#### Source

<!-- prettier-ignore -->
> - [x] On desktop, Shown in and timezone should be above Responses and scroll with it so that when all timeslots are shown
>      and one selects a timeslot and see the tooltip with the time, they can see which timezone that time belongs to

[Source lines 428-429](../../../../backlog/backlog.md#L428-L429)

#### Candidate behavior

On desktop, Shown in and timezone remain above Responses while scrolling so a selected timeslot's timezone is visible.

#### Applicability

Actor: event visitor; Location: desktop event-page Responses; Event kind: timed; Interaction mode: scroll and select slot; Viewport: desktop; State: all timeslots shown; Exclusions: mobile not established.

#### Classification

candidate FR

#### Existing Requirements and Confidence

Overlap: accepted FR-047 labels the Display Timezone control but does not set its position. Confidence: inferred.

#### Disposition

Review as sidebar-positioning behavior.

#### Open Questions

Does `Shown in and timezone` name one control or two, and should it remain sticky?

## CAND-103

#### Source

> - [x] Use monospace font for times
>   - [x] in the grid
>   - [x] in the tooltip

[Source lines 430-432](../../../../backlog/backlog.md#L430-L432)

#### Candidate behavior

No new requirement behavior asserted; the source prescribes a presentation implementation for grid and tooltip times.

#### Applicability

Actor: event visitor; Location: timed grid and tooltip; Event kind: timed; Interaction mode: viewing; Viewport: unspecified; State: time visible; Exclusions: other time displays not established.

#### Classification

implementation detail

#### Existing Requirements and Confidence

Overlap: accepted FR-024 defines which time format controls grid and tooltip times, not typography. Confidence: inferred.

#### Disposition

Exclude unless typography is confirmed as a durable accessibility or brand requirement.

#### Open Questions

Is monospace required for legibility, and which font family and fallback are intended?

## CAND-104

#### Source

> - [x] Don't show the cursor and tooltip in enabled, inactive or disabled cells

[Source lines 433-433](../../../../backlog/backlog.md#L433-L433)

#### Candidate behavior

The grid shows neither cursor nor tooltip for enabled-inactive or disabled cells.

#### Applicability

Actor: event visitor; Location: timed grid; Event kind: timed; Interaction mode: hover or selection; Viewport: unspecified; State: enabled-inactive or disabled cell; Exclusions: active cells.

#### Classification

candidate FR

#### Existing Requirements and Confidence

Overlap: accepted FR-014 derives grid cell states from event domains but does not prescribe cursor or tooltip visibility. Confidence: inferred.

#### Disposition

Review with CAND-112 as inactive-cell interaction behavior.

#### Open Questions

Does `cursor` mean a grid selection indicator, pointer cursor, or scheduled-event cursor?

## CAND-105

#### Source

> - [x] On mobile, given I'm on the event page and a timeslot is selected and the tooltip is visible, when I click the responses offcanvas, the selected timeslot and its tooltip must not disappear

[Source lines 434-434](../../../../backlog/backlog.md#L434-L434)

#### Candidate behavior

No new requirement behavior asserted; accepted FR-020 already preserves mobile selected Timed Slot state and tooltip visibility.

#### Applicability

Actor: event visitor; Location: mobile event page; Event kind: timed; Interaction mode: open Responses offcanvas; Viewport: mobile; State: a timeslot selected and tooltip visible; Exclusions: no selection.

#### Classification

existing requirement

#### Existing Requirements and Confidence

Overlap: accepted FR-020 directly covers preserving the selection and tooltip. Confidence: confirmed.

#### Disposition

Map to FR-020; assess CAND-089 separately for panel layering.

## CAND-106

#### Source

> - [x] On mobile, make the grid labels fit on the screen

[Source lines 435-435](../../../../backlog/backlog.md#L435-L435)

#### Candidate behavior

On mobile, grid labels fit within the screen.

#### Applicability

Actor: event visitor; Location: mobile grid; Event kind: unspecified; Interaction mode: viewing; Viewport: mobile; State: labels displayed; Exclusions: desktop.

#### Classification

candidate QR

#### Existing Requirements and Confidence

Overlap: none in accepted FR/QR records. Confidence: inferred.

#### Disposition

Review as a measurable responsive operability outcome.

#### Open Questions

Which labels, screen widths, zoom level, and clipping or wrapping criteria define fit?

## CAND-107

#### Source

> - [x] Given I'm on desktop and on the event page, when I hover over collapsed hours,
>   - the responses should show 0/N (behave similarly to disabled cells)
>   - and mark everyone unavailable
>   - and the selection at the highlight at the last hovered timeslot must be cleared

[Source lines 436-439](../../../../backlog/backlog.md#L436-L439)

#### Candidate behavior

Hovering collapsed hours on desktop shows `0/N`, marks all responses unavailable, and clears the prior highlight selection.

#### Applicability

Actor: event visitor; Location: desktop timed grid and Responses; Event kind: timed; Interaction mode: hover; Viewport: desktop; State: collapsed hours; Exclusions: active and disabled cells.

#### Classification

candidate FR

#### Existing Requirements and Confidence

Overlap: accepted FR-011 defines collapsed inactive runs, not their hover response state. Confidence: inferred.

#### Disposition

Review as collapsed-hour interaction behavior.

#### Open Questions

Does `0/N` include protected responses, and what visual mark means unavailable?

## CAND-108

#### Source

> - [x] Given the selected dates are non-consecutive and make the grid split into sub-grids (e.g. Aug 6 and 9), when I click (desktop, mobile) or hover (desktop) in the space between sub-grids, the highlight and tooltip at the last selected timeslot must be cleared like when clicking or hovering inactive timeslots

[Source lines 440-440](../../../../backlog/backlog.md#L440-L440)

#### Candidate behavior

Interacting in space between split sub-grids clears the prior slot highlight and tooltip.

#### Applicability

Actor: event visitor; Location: timed grid; Event kind: timed; Interaction mode: click on desktop or mobile, hover on desktop; Viewport: desktop or mobile; State: non-consecutive dates split grid; Exclusions: contiguous grids.

#### Classification

candidate FR

#### Existing Requirements and Confidence

Overlap: accepted FR-002 projects timed slots into date columns but does not define inter-grid-space interaction. Confidence: inferred.

#### Disposition

Review as non-cell interaction behavior.

#### Open Questions

Should touch movement in the gap also clear selection?

## CAND-109

#### Source

> - [x] On desktop, given I'm on the event page, when I hover an active timeslot or hover outside the grid and then hover over an inactive slot, then the Responses sidebar should show (0/N) and all responses crossed-out and the status square should have the corresponding color (light-grey for enabled, inactive etc.)

[Source lines 441-441](../../../../backlog/backlog.md#L441-L441)

#### Candidate behavior

Hovering an inactive slot after leaving the grid shows `0/N`, crossed-out responses, and the corresponding status-square color.

#### Applicability

Actor: event visitor; Location: desktop event-page grid and Responses; Event kind: timed; Interaction mode: pointer hover; Viewport: desktop; State: inactive slot after grid exit; Exclusions: active slot display except transition.

#### Classification

duplicate or refinement

#### Existing Requirements and Confidence

Overlap: CAND-107 specifies the collapsed-hours variant; no accepted FR/QR overlap. Confidence: inferred.

#### Disposition

Consolidate with inactive-cell response-display behavior.

#### Open Questions

Does this apply to disabled cells and every inactive status color?

## CAND-110

#### Source

> - [x] On mobile, buttons with arrows for switching pages when there are several days, buttons should be the same size

[Source lines 442-442](../../../../backlog/backlog.md#L442-L442)

#### Candidate behavior

On mobile multi-day pages, arrow buttons for page switching have equal size.

#### Applicability

Actor: event visitor; Location: mobile event page pagination; Event kind: unspecified; Interaction mode: viewing or switching pages; Viewport: mobile; State: several days; Exclusions: single-day pages and desktop.

#### Classification

candidate FR

#### Existing Requirements and Confidence

Overlap: none in accepted FR/QR records. Confidence: inferred.

#### Disposition

Review as responsive control consistency.

#### Open Questions

What dimensions and which arrows are in scope?

## CAND-111

#### Source

> - [x] On mobile, when I click a disabled timeslot, the selection should disappear

[Source lines 443-443](../../../../backlog/backlog.md#L443-L443)

#### Candidate behavior

On mobile, selecting a disabled timeslot clears the existing selection.

#### Applicability

Actor: event visitor; Location: mobile timed grid; Event kind: timed; Interaction mode: tap; Viewport: mobile; State: disabled timeslot; Exclusions: enabled-inactive cells not explicit.

#### Classification

candidate FR

#### Existing Requirements and Confidence

Overlap: CAND-104 excludes a cursor and tooltip in disabled cells; no accepted FR/QR overlap. Confidence: inferred.

#### Disposition

Review with CAND-104 and CAND-112.

#### Open Questions

Should tap show no replacement selection immediately, and does this include padding cells?

## CAND-112

#### Source

> - [x] On desktop, given I'm on the event page, when I hover or click inside the grid outside active cells, the highlight at the last timeslot and any tooltip should not be visible

[Source lines 444-444](../../../../backlog/backlog.md#L444-L444)

#### Candidate behavior

On desktop, hovering or clicking outside active grid cells clears the prior highlight and hides its tooltip.

#### Applicability

Actor: event visitor; Location: desktop timed grid; Event kind: timed; Interaction mode: hover or click; Viewport: desktop; State: outside active cells; Exclusions: active cells.

#### Classification

duplicate or refinement

#### Existing Requirements and Confidence

Overlap: CAND-104 and CAND-111 specify the same visibility and clearing rule by cell state; no accepted FR/QR overlap. Confidence: inferred.

#### Disposition

Consolidate into one cross-viewport inactive-cell requirement.

#### Open Questions

Does `outside active cells` include inter-grid space and collapsed hours?

## CAND-113

#### Source

> - [x] Create a glossary and oblige agents to read it in AGENTS.md

[Source lines 445-445](../../../../backlog/backlog.md#L445-L445)

#### Candidate behavior

No new requirement behavior asserted; the source is documentation governance.

#### Applicability

Actor: agent; Location: repository documentation; Event kind: not applicable; Interaction mode: authoring; Viewport: not applicable; State: requirement work; Exclusions: runtime product behavior.

#### Classification

ADR or decision

#### Existing Requirements and Confidence

Overlap: the requirements README already directs controlled terminology use; no accepted FR/QR behavior overlap. Confidence: confirmed.

#### Disposition

Exclude from product requirement consolidation.

## CAND-114

#### Source

> - [x] In responses,
>   - when I hover or click in a grid, show the square for the status of a person at the corresponding timeslot (available, if needed, unavailable, etc.) instead of a profile image;
>   - when I hover or click outside of the grid, replace the status square with a profile icon like it's now

[Source lines 446-448](../../../../backlog/backlog.md#L446-L448)

#### Candidate behavior

Responses show a per-slot status square for grid interaction and a profile icon outside the grid.

#### Applicability

Actor: event visitor; Location: Responses; Event kind: timed; Interaction mode: hover or click inside or outside grid; Viewport: unspecified; State: corresponding slot selected or no grid interaction; Exclusions: dates-only events not established.

#### Classification

candidate FR

#### Existing Requirements and Confidence

Overlap: accepted FR-014 derives timed-grid cell states but does not prescribe response icons. Confidence: inferred.

#### Disposition

Review as response-status presentation behavior.

#### Open Questions

Which statuses require squares, and what restores the profile icon after leaving the grid?

## CAND-115

#### Source

> - [x] Add "Disabled, collapsed" legend item

[Source lines 449-449](../../../../backlog/backlog.md#L449-L449)

#### Candidate behavior

The timed-event legend includes a `Disabled, collapsed` item.

#### Applicability

Actor: event visitor; Location: timed-event legend; Event kind: timed range; Interaction mode: viewing; Viewport: unspecified; State: legend shown; Exclusions: non-range timed events not established.

#### Classification

candidate FR

#### Existing Requirements and Confidence

Overlap: accepted FR-011 defines collapsed runs but not legend disclosure. Confidence: inferred.

#### Disposition

Review with CAND-116 as one legend item.

#### Open Questions

Does the item appear only when collapsed cells exist?

## CAND-116

#### Source

> - [x] dashed outline in the legend bullet for Disabled, collapsed

[Source lines 450-450](../../../../backlog/backlog.md#L450-L450)

#### Candidate behavior

No new requirement behavior asserted; this is a visual refinement of CAND-115's legend item.

#### Applicability

Actor: event visitor; Location: timed-event legend; Event kind: timed range; Interaction mode: viewing; Viewport: unspecified; State: Disabled, collapsed item displayed; Exclusions: other legend items.

#### Classification

duplicate or refinement

#### Existing Requirements and Confidence

Overlap: CAND-115 is the parent legend behavior; no accepted FR/QR overlap. Confidence: inferred.

#### Disposition

Consolidate with CAND-115 if the visual distinction is retained.

#### Open Questions

Is a dashed outline required for interpretation or merely the chosen styling?

## CAND-117

#### Source

> - [x] On desktop, in the new event form and on the new event page, when I scroll the time zone menu, the width is always 520 and doesn't change based on the content length.

[Source lines 451-451](../../../../backlog/backlog.md#L451-L451)

#### Candidate behavior

On desktop, the time-zone menu in the new-event form and page remains 520 wide while scrolling.

#### Applicability

Actor: event creator; Location: desktop new-event form and page time-zone menu; Event kind: unspecified; Interaction mode: scroll menu; Viewport: desktop; State: menu open; Exclusions: existing-event pages and mobile.

#### Classification

candidate FR

#### Existing Requirements and Confidence

Overlap: accepted FR-021 initializes new-event timezone but not menu dimensions. Confidence: inferred.

#### Disposition

Review as a fixed visual-control outcome.

#### Open Questions

Are units pixels, and does this apply to every time-zone selector?

## CAND-118

#### Source

> - [x] Schedule event on Google Calendar should happen in the display timezone

[Source lines 452-452](../../../../backlog/backlog.md#L452-L452)

#### Candidate behavior

Scheduling an event on Google Calendar uses the display timezone.

#### Applicability

Actor: event visitor; Location: Google Calendar scheduling; Event kind: timed; Interaction mode: schedule event; Viewport: unspecified; State: display timezone set; Exclusions: dates-only events not established.

#### Classification

candidate FR

#### Existing Requirements and Confidence

Overlap: accepted FR-024 keeps display time format independent; no accepted requirement assigns Calendar timezone. Confidence: inferred.

#### Disposition

Review as external-calendar integration behavior.

#### Open Questions

Does the Calendar event preserve instants while expressing its timezone as display timezone?

## CAND-119

#### Source

> - [x] Given I'm on the event page and there's a scheduled event, when I click Reschedule event, the Schedule button is active

[Source lines 453-453](../../../../backlog/backlog.md#L453-L453)

#### Candidate behavior

When an event is already scheduled, entering Reschedule event leaves Schedule active.

#### Applicability

Actor: event visitor; Location: event page scheduling controls; Event kind: timed; Interaction mode: reschedule; Viewport: unspecified; State: scheduled event exists; Exclusions: unscheduled event not established.

#### Classification

candidate FR

#### Existing Requirements and Confidence

Overlap: accepted FR-012 permits replacing a Scheduled Event Time but does not prescribe button state. Confidence: inferred.

#### Disposition

Review as an interaction refinement of FR-012.

#### Open Questions

Does active mean enabled, selected, or immediately actionable without changing slots?

## CAND-120

#### Source

> - [x] On the Event not found page, when using tab-navigation, the rectangle should coincide with the Back to home button

[Source lines 454-454](../../../../backlog/backlog.md#L454-L454)

#### Candidate behavior

Keyboard focus on Back to home is visually aligned with the button bounds.

#### Applicability

Actor: keyboard user; Location: Event not found page; Event kind: not applicable; Interaction mode: tab navigation; Viewport: unspecified; State: Back to home focused; Exclusions: other controls not established.

#### Classification

candidate QR

#### Existing Requirements and Confidence

Overlap: proposed QR-008 covers WCAG conformance; no accepted FR/QR overlap. Confidence: inferred.

#### Disposition

Review as an accessibility quality scenario.

#### Open Questions

What focus-indicator contrast, thickness, and offset satisfy the intended criterion?

## CAND-121

#### Source

> - [x] On desktop, on a dates only event page, in the second row, on the right:
>   - When there are no responses, show "Start on Monday" option in the second row
>   - When there are responses, show "Show best days" option and "More options" (which includes "Start on Monday" and "Hide if needed")

[Source lines 455-457](../../../../backlog/backlog.md#L455-L457)

#### Candidate behavior

On desktop dates-only pages, the second row shows Start on Monday with no responses, or Show best days and More options with responses.

#### Applicability

Actor: event visitor; Location: desktop dates-only event page second row; Event kind: dates-only; Interaction mode: viewing; Viewport: desktop; State: no responses or one or more responses; Exclusions: timed events and mobile.

#### Classification

candidate FR

#### Existing Requirements and Confidence

Overlap: accepted FR-049 organizes mobile controls, not this desktop dates-only conditional display. Confidence: inferred.

#### Disposition

Review as conditional controls behavior.

#### Open Questions

Should More options always contain both listed options, and what does Hide if needed affect?

## CAND-122

#### Source

> - [x] For dates-only events, the combination of days of week and dates in the calendar must match the reality

[Source lines 458-458](../../../../backlog/backlog.md#L458-L458)

#### Candidate behavior

Dates-only calendars display each date under its actual day of week.

#### Applicability

Actor: event visitor; Location: dates-only event calendar; Event kind: dates-only; Interaction mode: viewing; Viewport: unspecified; State: calendar rendered; Exclusions: timed grids.

#### Classification

candidate FR

#### Existing Requirements and Confidence

Overlap: none in accepted FR/QR records. Confidence: inferred.

#### Disposition

Review as calendar correctness behavior.

#### Open Questions

Which timezone determines a date near a date boundary?

## CAND-123

#### Source

> - [x] For dates-only events, the color of disabled dates must be dark-grey like for disabled padding cells in timed event

[Source lines 459-459](../../../../backlog/backlog.md#L459-L459)

#### Candidate behavior

Dates-only calendars render disabled dates dark-grey like timed-event disabled padding cells.

#### Applicability

Actor: event visitor; Location: dates-only event calendar; Event kind: dates-only; Interaction mode: viewing; Viewport: unspecified; State: disabled date; Exclusions: enabled dates.

#### Classification

candidate FR

#### Existing Requirements and Confidence

Overlap: none in accepted FR/QR records. Confidence: inferred.

#### Disposition

Review as cross-event visual consistency.

#### Open Questions

Does dark-grey require an exact color and accessibility contrast threshold?

## CAND-124

#### Source

> - [x] On the Event not found page, the "Back to home" button must have a black shadow (like on the home page), not greenish glow

[Source lines 460-460](../../../../backlog/backlog.md#L460-L460)

#### Candidate behavior

The Event not found Back to home button uses the home page's black shadow rather than a greenish glow.

#### Applicability

Actor: event visitor; Location: Event not found page; Event kind: not applicable; Interaction mode: viewing; Viewport: unspecified; State: button visible; Exclusions: other buttons.

#### Classification

candidate FR

#### Existing Requirements and Confidence

Overlap: accepted FR-040 centers the not-found page, not button styling. Confidence: inferred.

#### Disposition

Review as page-specific visual behavior.

#### Open Questions

Is matching the home-page token sufficient, or is black shadow an exact visual constraint?

## CAND-125

#### Source

> - [x] "If needed" - show in the legend to explain the status color in responses

[Source lines 461-461](../../../../backlog/backlog.md#L461-L461)

#### Candidate behavior

The Responses legend shows `If needed` to explain its status color.

#### Applicability

Actor: event visitor; Location: Responses legend; Event kind: unspecified; Interaction mode: viewing; Viewport: unspecified; State: legend visible; Exclusions: no legend.

#### Classification

candidate FR

#### Existing Requirements and Confidence

Overlap: accepted FR-006 keeps availability states mutually exclusive but does not require this legend item. Confidence: inferred.

#### Disposition

Review as status explanation behavior.

#### Open Questions

Does the legend apply to timed and dates-only response presentations?

## CAND-126

#### Source

> - [x] On dates-only event page, the Responses top of the text must horizontally coincide with the top edge of the grid

[Source lines 462-462](../../../../backlog/backlog.md#L462-L462)

#### Candidate behavior

On dates-only event pages, the top of Responses text horizontally aligns with the grid's top edge.

#### Applicability

Actor: event visitor; Location: dates-only event page; Event kind: dates-only; Interaction mode: viewing; Viewport: unspecified; State: Responses visible; Exclusions: timed events.

#### Classification

candidate FR

#### Existing Requirements and Confidence

Overlap: none in accepted FR/QR records. Confidence: inferred.

#### Disposition

Review as layout behavior.

#### Open Questions

Does `horizontally coincide` mean equal vertical position or shared horizontal coordinate?

## CAND-127

#### Source

<!-- prettier-ignore -->
> - [x] Keep yellow status for "if needed" responses but don't:
>   - highlight such responses text
>   - show asterisk ("\*") near such responses
>   - show "* if needed" under responses

[Source lines 463-466](../../../../backlog/backlog.md#L463-L466)

#### Candidate behavior

If needed responses retain yellow status but omit text highlight, an asterisk, and the `* if needed` note.

#### Applicability

Actor: event visitor; Location: Responses; Event kind: unspecified; Interaction mode: viewing; Viewport: unspecified; State: If needed response; Exclusions: other response statuses.

#### Classification

candidate FR

#### Existing Requirements and Confidence

Overlap: accepted FR-006 defines availability-state exclusivity, not its presentation. Confidence: inferred.

#### Disposition

Review as one response-status presentation requirement.

#### Open Questions

Is yellow defined by a shared status-color token and is it accessible in every response view?

## CAND-128

#### Source

> - [x] On the read-only timed event page, when I switch the timezone, the grid shouldn't collapse

[Source lines 467-467](../../../../backlog/backlog.md#L467-L467)

#### Candidate behavior

Changing timezone on a read-only timed-event page does not collapse the grid.

#### Applicability

Actor: read-only event visitor; Location: timed-event page; Event kind: timed; Interaction mode: change timezone; Viewport: unspecified; State: read-only; Exclusions: editable pages.

#### Classification

candidate FR

#### Existing Requirements and Confidence

Overlap: accepted FR-013 preserves timed-slot instants across display-timezone changes; it does not specify grid collapse. Confidence: inferred.

#### Disposition

Review as display-timezone interaction behavior.

#### Open Questions

Which collapse state must be preserved, and does this apply after every timezone change?

## CAND-129

#### Source

> - [x] In timed range events, the Legend shall show the light-gray "Disabled, change in Edit event"

[Source lines 468-468](../../../../backlog/backlog.md#L468-L468)

#### Candidate behavior

Timed range event legends show a light-gray `Disabled, change in Edit event` item.

#### Applicability

Actor: event visitor; Location: timed range event legend; Event kind: timed range; Interaction mode: viewing; Viewport: unspecified; State: legend displayed; Exclusions: custom-domain and dates-only events not established.

#### Classification

candidate FR

#### Existing Requirements and Confidence

Overlap: accepted FR-074 derives the timed-event enabled domain but does not require this legend guidance. Confidence: inferred.

#### Disposition

Review alongside CAND-115 for a coherent disabled-state legend.

#### Open Questions

Is `change in Edit event` available to non-owners, and how does this item differ from Disabled, collapsed?

## CAND-130

#### Source

> - [x] Address linter warnings

[Source lines 469-469](../../../../backlog/backlog.md#L469-L469)

#### Candidate behavior

No new requirement behavior asserted; this is a code-quality maintenance instruction.

#### Applicability

Actor: developer; Location: codebase; Event kind: not applicable; Interaction mode: development; Viewport: not applicable; State: linter warnings present; Exclusions: runtime behavior.

#### Classification

implementation detail

#### Existing Requirements and Confidence

Overlap: none in accepted FR/QR records. Confidence: confirmed.

#### Disposition

Exclude from requirement migration.

## CAND-131

#### Source

> - [x] In "New event" form, when "Dates only" is selected, in "Advanced options", there should be no "Time increment"

[Source lines 470-470](../../../../backlog/backlog.md#L470-L470)

#### Candidate behavior

When Dates only is selected in New event, Advanced options omits Time increment.

#### Applicability

Actor: event creator; Location: New event Advanced options; Event kind: dates-only; Interaction mode: choose event kind; Viewport: unspecified; State: Dates only selected; Exclusions: Dates and times.

#### Classification

candidate FR

#### Existing Requirements and Confidence

Overlap: accepted FR-026 defines Dates only as an event kind but not this option visibility. Confidence: inferred.

#### Disposition

Review as event-kind-specific form behavior.

#### Open Questions

Should Time increment be hidden, disabled, or absent from submitted data?

## CAND-132

#### Source

> - [x] On the dates-only event page, when I click the "Edit event" button, then the edit form opens

[Source lines 471-471](../../../../backlog/backlog.md#L471-L471)

#### Candidate behavior

Selecting Edit event on a dates-only event page opens the edit form.

#### Applicability

Actor: unspecified event visitor; Location: dates-only event page; Event kind: dates-only; Interaction mode: click Edit event; Viewport: unspecified; State: edit control available; Exclusions: timed events.

#### Classification

candidate FR

#### Existing Requirements and Confidence

Overlap: accepted FR-018 restricts settings editing to the Event Owner; no accepted requirement establishes this entry point. Confidence: inferred.

#### Disposition

Review with owner-permission applicability.

#### Open Questions

Is the actor the Event Owner, and is the opened form editable or read-only?

## CAND-133

#### Source

> - [x] On dates-only event page, there must be space between the grid and Responses sidebar

[Source lines 472-472](../../../../backlog/backlog.md#L472-L472)

#### Candidate behavior

Dates-only event pages provide visible space between the grid and Responses sidebar.

#### Applicability

Actor: event visitor; Location: dates-only event page; Event kind: dates-only; Interaction mode: viewing; Viewport: unspecified; State: sidebar present; Exclusions: timed events.

#### Classification

candidate FR

#### Existing Requirements and Confidence

Overlap: none in accepted FR/QR records. Confidence: inferred.

#### Disposition

Review as layout behavior.

#### Open Questions

What minimum spacing and responsive behavior are intended?

## CAND-134

#### Source

> - [x] On the dates-only event page, show the Schedule button even if there are no responses, like on the timed event page.

[Source lines 473-473](../../../../backlog/backlog.md#L473-L473)

#### Candidate behavior

Dates-only event pages show Schedule even when there are no responses.

#### Applicability

Actor: event visitor; Location: dates-only event page; Event kind: dates-only; Interaction mode: viewing; Viewport: unspecified; State: no responses; Exclusions: timed events used only as comparison.

#### Classification

candidate FR

#### Existing Requirements and Confidence

Overlap: accepted FR-012 permits a dates-only Scheduled Event Time but does not require this button's visibility. Confidence: inferred.

#### Disposition

Review as scheduling-entry-point behavior.

#### Open Questions

Who may activate Schedule and is it enabled with no responses?

## CAND-135

#### Source

> - [x] The width of the description field must be the same as on the timed event page

[Source lines 474-474](../../../../backlog/backlog.md#L474-L474)

#### Candidate behavior

The description field matches the width used on the timed event page.

#### Applicability

Actor: event visitor; Location: dates-only event page description field; Event kind: dates-only; Interaction mode: viewing or editing not established; Viewport: unspecified; State: field visible; Exclusions: timed event used as reference.

#### Classification

candidate FR

#### Existing Requirements and Confidence

Overlap: accepted FR-019 preserves multiline event descriptions, not field width. Confidence: inferred.

#### Disposition

Review as cross-page visual consistency.

#### Open Questions

Is the source page dates-only, and does width mean the input, display region, or container?

## CAND-136

#### Source

> - [x] On the dates-only event page, "Start on Monday" must be centered vertically with "Add availability"

[Source lines 475-475](../../../../backlog/backlog.md#L475-L475)

#### Candidate behavior

On dates-only event pages, Start on Monday is vertically centered with Add availability.

#### Applicability

Actor: event visitor; Location: dates-only event-page controls; Event kind: dates-only; Interaction mode: viewing; Viewport: unspecified; State: both controls visible; Exclusions: timed events.

#### Classification

candidate FR

#### Existing Requirements and Confidence

Overlap: CAND-121 covers conditional Start on Monday visibility; no accepted FR/QR overlap. Confidence: inferred.

#### Disposition

Review with CAND-121 as a control-layout refinement.

#### Open Questions

Which baseline defines vertical centering at responsive breakpoints?

## CAND-137

#### Source

> - [x] When VITE_ENABLE_SIGN_IN is not false, on the landing page, there should be the Sign in button.
>       When I click that button, "Sign in" page opens with options like Google Calendar / Email.

[Source lines 476-477](../../../../backlog/backlog.md#L476-L477)

#### Candidate behavior

When sign-in is enabled, the landing page shows Sign in, which opens a page with Google Calendar and Email options.

#### Applicability

Actor: landing-page visitor; Location: landing page and Sign in page; Event kind: not applicable; Interaction mode: select Sign in; Viewport: unspecified; State: `VITE_ENABLE_SIGN_IN` is not false; Exclusions: sign-in disabled.

#### Classification

candidate FR

#### Existing Requirements and Confidence

Overlap: proposed FR-007 gates sign-in entry points when disabled; accepted FR-031 sends email magic links, but neither establishes this enabled-state entry point. Confidence: inferred.

#### Disposition

Review as sign-in entry and provider-choice behavior.

#### Open Questions

Are Google Calendar and Email the complete option set, and is this condition valid for every deployment?

## CAND-138

#### Source

> - [x] On each page, Sign in must be the left-most button

[Source lines 478-478](../../../../backlog/backlog.md#L478-L478)

#### Candidate behavior

When present, Sign in is the left-most button on each page.

#### Applicability

Actor: visitor; Location: every page header or action area; Event kind: not applicable; Interaction mode: viewing; Viewport: unspecified; State: Sign in available; Exclusions: pages without Sign in not established.

#### Classification

candidate FR

#### Existing Requirements and Confidence

Overlap: proposed FR-007 governs disabled-state visibility, not ordering. Confidence: inferred.

#### Disposition

Review as global navigation layout behavior.

#### Open Questions

Does left-most mean visual order in RTL layouts and on mobile menus?

## CAND-139

#### Source

> - [x] Given I sign in, when I enter an unregistered email and click Continue with email:
>   - The input field is highlighted red
>   - The error appears like on accounts.google.com: (red alert icon) "Couldn’t find this account. Create account"

[Source lines 479-481](../../../../backlog/backlog.md#L479-L481)

#### Candidate behavior

An unregistered email submission highlights the input red and shows `Couldn’t find this account. Create account` with a red alert icon.

#### Applicability

Actor: sign-in visitor; Location: Sign in page email flow; Event kind: not applicable; Interaction mode: submit unregistered email; Viewport: unspecified; State: email has no account; Exclusions: registered email and other providers.

#### Classification

candidate FR

#### Existing Requirements and Confidence

Overlap: accepted FR-023 displays email registration status during sign-in; this source supplies its unregistered-email presentation. Confidence: inferred.

#### Disposition

Review as a presentation refinement of FR-023.

#### Open Questions

Is the exact accounts.google.com-like visual treatment a durable requirement, and does Create account initiate registration?
