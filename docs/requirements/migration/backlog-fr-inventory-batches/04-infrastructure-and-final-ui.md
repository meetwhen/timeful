# Backlog FR Inventory: Infrastructure And Final UI

This is a durable batch artifact for review, not a normative requirements record. `CAND-*` IDs are temporary and will be consolidated by [../backlog-fr-inventory.md](../backlog-fr-inventory.md).

## CAND-168: Postgres Anonymous-Event Storage

### Source

> - [x] Switch to Postgres for new anon events

[Source lines 536-536](../../../../backlog/backlog.md#L536)

### Candidate behavior

No new requirement behavior asserted; the source selects storage implementation for new anonymous events.

### Applicability

- Actor: system implementation team
- Location: anonymous-event creation
- Event kind: unspecified
- Interaction mode: not applicable
- Viewport: not applicable
- State: new anonymous events
- Exclusions: identifier routing and existing-event migration reporting

### Classification

implementation detail

### Existing Requirements and Confidence

Existing requirements: None identified.

Confidence: confirmed

### Disposition

Retain as migration provenance; do not create an FR or QR.

## CAND-169: Postgres Anonymous-Event Identifier Routing

### Source

> - [x] Make 8-character Crockford base32 ids point to events in Postgres

[Source lines 537-537](../../../../backlog/backlog.md#L537)

### Candidate behavior

No new requirement behavior asserted; the source selects identifier-routing implementation for Postgres events.

### Applicability

- Actor: system implementation team
- Location: anonymous-event lookup
- Event kind: unspecified
- Interaction mode: not applicable
- Viewport: not applicable
- State: lookup by 8-character Crockford base32 identifier
- Exclusions: anonymous-event storage selection and existing-event migration reporting

### Classification

implementation detail

### Existing Requirements and Confidence

Existing requirements: None identified. Overlap: CAND-168 records the separate storage decision; no requirement overlap identified.

Confidence: confirmed

### Disposition

Retain as migration provenance; do not create an FR or QR.

### Open Questions

- Does the Postgres routing apply only to new anonymous events, or to every 8-character Crockford base32 identifier?

## CAND-170: Postgres Migration Link Report

### Source

> - [x] Migrate to postgres and report new links:
>   - <http://127.0.0.1:4173/e/BFfB4> -> <http://127.0.0.1:4173/e/M7FVZFYP>
>   - <http://127.0.0.1:4173/e/6df78> -> <http://127.0.0.1:4173/e/C9ZC3WZS>

[Source lines 538-540](../../../../backlog/backlog.md#L538-L540)

### Candidate behavior

No new requirement behavior asserted; this is a one-off migration verification report with concrete local links.

### Applicability

- Actor: migration operator
- Location: migration execution
- Event kind: unspecified
- Interaction mode: not applicable
- Viewport: not applicable
- State: one-off completed migration
- Exclusions: ongoing event-link behavior

### Classification

implementation detail

### Existing Requirements and Confidence

Existing requirements: None identified.

Confidence: confirmed

### Disposition

Retain as migration evidence only; do not create an FR or QR.

## CAND-171: Timed More-Options Order

### Source

> - [x] On the timed event page, in More options, "Show all hours" must be above "Hide if needed times"

[Source lines 541-541](../../../../backlog/backlog.md#L541)

### Candidate behavior

On the timed event page, More options presents `Show all hours` above `Hide if needed times`.

### Applicability

- Actor: event visitor
- Location: timed event page, More options
- Event kind: timed
- Interaction mode: viewing
- Viewport: unspecified
- State: More options open
- Exclusions: dates-only event page

### Classification

candidate FR

### Existing Requirements and Confidence

Existing requirements: [FR-011](../../functional/fr/FR-011.md) establishes collapsed timed-grid behavior but not this control order.

Confidence: inferred

### Disposition

Consolidate only if control order is a durable product behavior rather than incidental layout.

### Open Questions

- Is the required order intended for every viewport and control state?

## CAND-172: Dates-Only More-Options Order

### Source

> - [x] On the dates-only event page, in More options, "Start on Monday" must be above "Hide if needed days"

[Source lines 542-542](../../../../backlog/backlog.md#L542)

### Candidate behavior

On the dates-only event page, More options presents `Start on Monday` above `Hide if needed days`.

### Applicability

- Actor: event visitor
- Location: dates-only event page, More options
- Event kind: dates-only
- Interaction mode: viewing
- Viewport: unspecified
- State: More options open
- Exclusions: timed event page

### Classification

candidate FR

### Existing Requirements and Confidence

Existing requirements: None identified.

Confidence: inferred

### Disposition

Consolidate only if this ordering is product-facing rather than incidental layout.

### Open Questions

- Does the order apply on mobile and desktop?

## CAND-173: Sidebar Spacing Consistency

### Source

> - [x] Spaces between elements in the sidebar must be the same (Timezone-Responses-Legend)

[Source lines 543-543](../../../../backlog/backlog.md#L543)

### Candidate behavior

The sidebar renders equal spacing between the timezone, responses, and legend elements.

### Applicability

- Actor: event visitor
- Location: event-page sidebar
- Event kind: unspecified
- Interaction mode: viewing
- Viewport: unspecified
- State: sidebar elements visible
- Exclusions: controls outside the named elements

### Classification

candidate FR

### Existing Requirements and Confidence

Existing requirements: None identified.

Confidence: inferred

### Disposition

Consider consolidation as a visual specification only if the named elements and measurement are confirmed.

### Open Questions

- What spacing value and tolerance define “the same”?
- Does this apply where the sidebar is stacked on mobile?

## CAND-174: E2E Test Caching

### Source

> - [x] Set up caching to speed up E2E tests

[Source lines 544-544](../../../../backlog/backlog.md#L544)

### Candidate behavior

No new requirement behavior asserted; this directs test-infrastructure optimization without a measurable target.

### Applicability

- Actor: test infrastructure
- Location: end-to-end test execution
- Event kind: not applicable
- Interaction mode: not applicable
- Viewport: not applicable
- State: test setup and execution
- Exclusions: runtime product performance

### Classification

implementation detail

### Existing Requirements and Confidence

Existing requirements: [QR-009](../../quality/qr/QR-009.md) concerns timed-event runtime response time, not test execution.

Confidence: confirmed

### Disposition

Retain as implementation provenance; no QR is asserted.

## CAND-175: Range-Event Collapsed Hours

### Source

> - [x] Collapsed hours should work in range events too, not only specific-times events

[Source lines 545-545](../../../../backlog/backlog.md#L545)

### Candidate behavior

No new requirement behavior asserted; the source broadens implementation coverage of existing collapsed-grid behavior.

### Applicability

- Actor: event visitor
- Location: timed event page grid
- Event kind: timed
- Interaction mode: viewing or scheduling
- Viewport: unspecified
- State: collapsed hours enabled
- Exclusions: availability editing and specific-times editing

### Classification

existing requirement

### Existing Requirements and Confidence

Existing requirements: [FR-011](../../functional/fr/FR-011.md) already requires collapsed inactive runs for timed-event viewing and scheduling.

Confidence: confirmed

### Disposition

Map to FR-011; no new FR.

## CAND-176: Dates-Only Edit-Event Entry

### Source

> - [x] On dates-only event page, when I click the "Edit event" button, a form for editing the event opens

[Source lines 546-546](../../../../backlog/backlog.md#L546)

### Candidate behavior

Selecting `Edit event` on a dates-only event page opens an event-editing form.

### Applicability

- Actor: event owner
- Location: dates-only event page
- Event kind: dates-only
- Interaction mode: event settings editing
- Viewport: unspecified
- State: Edit event action available
- Exclusions: non-owner access and form contents

### Classification

candidate FR

### Existing Requirements and Confidence

Existing requirements: [FR-018](../../functional/fr/FR-018.md) covers authorization to edit event settings, not this entry point.

Confidence: inferred

### Disposition

Consolidate with event-settings navigation if that navigation is durable.

### Open Questions

- Is this entry point owner-only, and what should non-owners see?

## CAND-177: Active-Slot Persistence Model

### Source

> - [x] Store only active timeslots and calculate all other types (enabled inactive and disabled) on the fly

[Source lines 547-547](../../../../backlog/backlog.md#L547)

### Candidate behavior

No new requirement behavior asserted; the source prescribes a persistence and derivation model rather than an observable outcome.

### Applicability

- Actor: system implementation
- Location: timed-event storage and grid derivation
- Event kind: timed
- Interaction mode: not applicable
- Viewport: not applicable
- State: persistence and rendering
- Exclusions: user-visible state semantics already specified elsewhere

### Classification

implementation detail

### Existing Requirements and Confidence

Existing requirements: [FR-009](../../functional/fr/FR-009.md) specifies visible grid states; [FR-074](../../functional/fr/FR-074.md) is proposed and specifies enabled-domain derivation.

Confidence: confirmed

### Disposition

Retain as design provenance; do not create an FR or QR from storage wording.

## CAND-178: Comparator E2E Test Import

### Source

> - [x] add Playwright e2e tests from the comparator to the main repository

[Source lines 548-548](../../../../backlog/backlog.md#L548)

### Candidate behavior

No new requirement behavior asserted; the source directs test migration into the main repository.

### Applicability

- Actor: development team
- Location: repository end-to-end test suite
- Event kind: not applicable
- Interaction mode: not applicable
- Viewport: not applicable
- State: test-suite maintenance
- Exclusions: comparator code import and product behavior

### Classification

implementation detail

### Existing Requirements and Confidence

Existing requirements: None identified. Overlap: CAND-179 records the separate retired comparator-code item; no requirement overlap identified.

Confidence: confirmed

### Disposition

Retain as historical implementation provenance; no FR or QR.

## CAND-179: Comparator Code Import

### Source

> - [x] commit comparator code to the main repository until it works as expected
>   - not included because we don't use a comparator anymore

[Source lines 549-550](../../../../backlog/backlog.md#L549-L550)

### Candidate behavior

No new requirement behavior asserted; the comparator was explicitly not included and the remaining work is a retired code-migration item.

### Applicability

- Actor: development team
- Location: repository codebase
- Event kind: not applicable
- Interaction mode: not applicable
- Viewport: not applicable
- State: retired comparator workflow
- Exclusions: comparator E2E test import and product behavior

### Classification

implementation detail

### Existing Requirements and Confidence

Existing requirements: None identified. Overlap: CAND-178 records the separate E2E-test import item; no requirement overlap identified.

Confidence: confirmed

### Disposition

Retain as historical implementation provenance; no FR or QR.

### Open Questions

- What current test coverage replaces the comparator tests that were to be imported?

## CAND-180: Development Node Version Consistency

### Source

> - [x] Use the same Node 26.5.0 for frontend in dockerfile and in dev

[Source lines 551-551](../../../../backlog/backlog.md#L551)

### Candidate behavior

No new requirement behavior asserted; the source selects a development-environment Node version.

### Applicability

- Actor: development team
- Location: frontend Dockerfile and local development
- Event kind: not applicable
- Interaction mode: not applicable
- Viewport: not applicable
- State: development configuration
- Exclusions: example environment configuration and deployed product behavior

### Classification

ADR or decision

### Existing Requirements and Confidence

Existing requirements: None identified. Overlap: CAND-181 records the separate example-environment configuration check; no requirement overlap identified.

Confidence: confirmed

### Disposition

Retain as configuration provenance; no FR or QR.

## CAND-181: Example Environment Configuration

### Source

> - [x] Check whether NODE_ENV and GIN_MODE are in the example .env files
>   - NODE_ENV isn't there because it's not used
>   - GIN_MODE is there

[Source lines 552-554](../../../../backlog/backlog.md#L552-L554)

### Candidate behavior

No new requirement behavior asserted; this is an example-configuration check and decision.

### Applicability

- Actor: development team
- Location: example environment files
- Event kind: not applicable
- Interaction mode: not applicable
- Viewport: not applicable
- State: development configuration
- Exclusions: Node version consistency and deployed product behavior

### Classification

ADR or decision

### Existing Requirements and Confidence

Existing requirements: None identified. Overlap: CAND-180 records the separate Node-version decision; no requirement overlap identified.

Confidence: confirmed

### Disposition

Retain as configuration provenance; no FR or QR.

### Open Questions

- Which example `.env` files are in scope for the `NODE_ENV` and `GIN_MODE` check?

## CAND-182: Staging Environment

### Source

> - [x] introduce staging environment

[Source lines 555-555](../../../../backlog/backlog.md#L555)

### Candidate behavior

No new requirement behavior asserted; the source names an environment but does not specify its purpose, parity, access, or lifecycle.

### Applicability

- Actor: operations team
- Location: deployment environments
- Event kind: not applicable
- Interaction mode: not applicable
- Viewport: not applicable
- State: deployment and release validation
- Exclusions: unspecified production behavior

### Classification

needs product decision

### Existing Requirements and Confidence

Existing requirements: None identified.

Confidence: needs product decision

### Disposition

Require an environment decision before considering a QR or ADR.

### Open Questions

- What release-validation purpose and production parity must staging provide?
- Who can deploy to and access staging?

## CAND-183: Edit-Availability Visibility

### Source

> - [x] On the event page:
>   - when there are no responses, the Edit availability shall be not visible.
>   - When there are responses, shall be visible

[Source lines 556-558](../../../../backlog/backlog.md#L556-L558)

### Candidate behavior

The event page hides `Edit availability` when there are no responses and shows it when responses exist.

### Applicability

- Actor: event visitor
- Location: event page
- Event kind: unspecified
- Interaction mode: viewing
- Viewport: unspecified
- State: zero responses or one or more responses
- Exclusions: add-availability visibility and authorization

### Classification

candidate FR

### Existing Requirements and Confidence

Existing requirements: [FR-065](../../functional/fr/FR-065.md) is proposed and concerns availability-editing response lists, not action visibility.

Confidence: inferred

### Disposition

Consolidate only after confirming whether “responses” means all event responses or editable responses.

### Open Questions

- Which response count controls the action for a visitor without edit authorization?

## CAND-184: Edit-Event Button Placement

### Source

> - [x] On the event page, Edit event button shall be under the event title

[Source lines 559-559](../../../../backlog/backlog.md#L559)

### Candidate behavior

The event page places the `Edit event` button below the event title.

### Applicability

- Actor: event owner
- Location: event page header
- Event kind: unspecified
- Interaction mode: viewing
- Viewport: unspecified
- State: edit action visible
- Exclusions: edit authorization and non-owner presentation

### Classification

candidate FR

### Existing Requirements and Confidence

Existing requirements: [FR-018](../../functional/fr/FR-018.md) covers edit authorization, not button placement.

Confidence: inferred

### Disposition

Consolidate only if placement is durable across responsive layouts.

### Open Questions

- Does “under” apply to both desktop and mobile headers?

## CAND-185: GitHub Repository Navbar Placement

### Source

> - [x] On the event page, Github repo button shall be in the navbar, right-most button

[Source lines 560-560](../../../../backlog/backlog.md#L560)

### Candidate behavior

The event-page navbar renders the GitHub repository button as its right-most button.

### Applicability

- Actor: event visitor
- Location: event-page navbar
- Event kind: unspecified
- Interaction mode: viewing
- Viewport: unspecified
- State: navbar visible
- Exclusions: other pages and navbar actions not named in the source

### Classification

candidate FR

### Existing Requirements and Confidence

Existing requirements: None identified.

Confidence: inferred

### Disposition

Consolidate only if external-repository navigation is a product commitment.

### Open Questions

- Is the button required for every deployment and viewport?

## CAND-186: README Icon Theme Readability

### Source

> - [x] In the README, the Timeful icon shall be readable both in night and light modes

[Source lines 561-561](../../../../backlog/backlog.md#L561)

### Candidate behavior

The README's Timeful icon remains readable in night and light modes.

### Applicability

- Actor: documentation reader
- Location: README rendering
- Event kind: not applicable
- Interaction mode: documentation viewing
- Viewport: unspecified
- State: night or light mode
- Exclusions: application UI icons

### Classification

candidate QR

### Existing Requirements and Confidence

Existing requirements: [QR-008](../../quality/qr/QR-008.md) is proposed and broadly concerns accessible coordination flows; no accepted QR covers README icon readability.

Confidence: inferred

### Disposition

Consolidate as documentation accessibility guidance only if a contrast criterion and rendering targets are confirmed.

### Open Questions

- Which README renderers and contrast threshold define “readable”?

## CAND-187: Collapsed-Strip Hover Treatment

### Source

> - [x] On the event page, when I hover over collapsed hours strip, neither the strip nor any cell is highlighted like active cells are highlighted on hover.

[Source lines 562-562](../../../../backlog/backlog.md#L562)

### Candidate behavior

Hovering a collapsed-hours strip does not apply the active-cell highlight to the strip or any grid cell.

### Applicability

- Actor: event visitor
- Location: event page grid
- Event kind: timed
- Interaction mode: viewing
- Viewport: hover-capable
- State: collapsed hours visible
- Exclusions: active-cell hover behavior

### Classification

candidate FR

### Existing Requirements and Confidence

Existing requirements: [FR-011](../../functional/fr/FR-011.md) establishes collapsed timed-grid runs but does not specify hover treatment.

Confidence: inferred

### Disposition

Consolidate as a grid interaction refinement if the treatment is intentionally observable.

### Open Questions

- Should touch interaction have an equivalent non-highlight rule?

## CAND-188: Scrolling No-Responses Position

### Source

> - [x] Bug: when scrolling the gri, the No responses yet changes the position
>   - Can't reproduce

[Source lines 563-564](../../../../backlog/backlog.md#L563-L564)

### Candidate behavior

No new requirement behavior asserted; the reported defect could not be reproduced and gives no expected stable position.

### Applicability

- Actor: event visitor
- Location: grid and no-responses message
- Event kind: unspecified
- Interaction mode: scrolling
- Viewport: unspecified
- State: no responses
- Exclusions: reproduced layout defects

### Classification

bug or investigation

### Existing Requirements and Confidence

Existing requirements: None identified.

Confidence: confirmed

### Disposition

Retain as non-reproduced investigation history; no FR or QR.

### Open Questions

- What grid position and scroll sequence originally exposed the issue?

## CAND-189: Desktop Sidebar Control Order

### Source

> - [x] On desktop, on the timed event page, time format and timezone shall be above Responses in the sidebar

[Source lines 565-565](../../../../backlog/backlog.md#L565)

### Candidate behavior

On desktop timed event pages, the time-format and timezone controls appear above Responses in the sidebar.

### Applicability

- Actor: event visitor
- Location: timed event page sidebar
- Event kind: timed
- Interaction mode: viewing
- Viewport: desktop
- State: sidebar visible
- Exclusions: mobile layout and dates-only pages

### Classification

candidate FR

### Existing Requirements and Confidence

Existing requirements: None identified; [FR-046](../../functional/fr/FR-046.md) and [FR-047](../../functional/fr/FR-047.md) are proposed labels, not ordering.

Confidence: inferred

### Disposition

Consolidate with timed-page sidebar layout if the desktop breakpoint is confirmed.

### Open Questions

- What viewport threshold defines desktop?

## CAND-190: Specific-Times Date Regression Report

### Source

> - [x] Create an event with specific times for dates Aug 30, 31, mark hours 0-4 for both dates, edit event, set dates for 28, 29, click next. See May 28, 30, 31 in specific times page, and May 30, 31 on the event page.
>   - Can't reproduce. I see Aug 28, Aug 29 on the event page

[Source lines 566-567](../../../../backlog/backlog.md#L566-L567)

### Candidate behavior

No new requirement behavior asserted; this is a non-reproduced date-editing regression report.

### Applicability

- Actor: event owner
- Location: specific-times editing and event page
- Event kind: timed
- Interaction mode: event settings editing
- Viewport: unspecified
- State: changed date selection
- Exclusions: confirmed date-domain behavior

### Classification

bug or investigation

### Existing Requirements and Confidence

Existing requirements: [FR-050](../../functional/fr/FR-050.md), [FR-051](../../functional/fr/FR-051.md), and [FR-052](../../functional/fr/FR-052.md) are proposed date-domain rules; none confirms this report.

Confidence: confirmed

### Disposition

Retain as investigation history; no FR or QR.

### Open Questions

- Which timezone, date picker state, and saved data are required to reproduce the reported May dates?

## CAND-191: Specific-Times Missing-Date Regression Report

### Source

> - [x] Create an event with two dates, mark timeslots for only one day in specific times, save, edit again and see only one day on the event page
>   - Can't reproduce. I see both days.

[Source lines 568-569](../../../../backlog/backlog.md#L568-L569)

### Candidate behavior

No new requirement behavior asserted; this is a non-reproduced saved-date regression report.

### Applicability

- Actor: event owner
- Location: specific-times editing and event page
- Event kind: timed
- Interaction mode: event settings editing
- Viewport: unspecified
- State: two selected dates with slots marked on one date
- Exclusions: confirmed date persistence behavior

### Classification

bug or investigation

### Existing Requirements and Confidence

Existing requirements: [FR-074](../../functional/fr/FR-074.md) is proposed and derives the enabled domain for each picked date; it does not confirm the report.

Confidence: confirmed

### Disposition

Retain as investigation history; no FR or QR.

### Open Questions

- What persisted data or transition would cause one selected date to disappear?

## CAND-192: Mobile Timezone-Control Width

### Source

> - [x] On mobile, on the timed event page, the timezone control keeps a fixed width (112px) so the reset button fits inside without resizing the control, with the time format and days-per-page buttons on either side

[Source lines 570-570](../../../../backlog/backlog.md#L570)

### Candidate behavior

On mobile timed event pages, the timezone control remains 112px wide so its reset button fits without resizing the control, between the time-format and days-per-page buttons.

### Applicability

- Actor: event visitor
- Location: timed event page controls
- Event kind: timed
- Interaction mode: viewing
- Viewport: mobile
- State: timezone reset button available
- Exclusions: desktop layout and dates-only pages

### Classification

candidate FR

### Existing Requirements and Confidence

Existing requirements: [FR-049](../../functional/fr/FR-049.md) is proposed and places the mobile controls in a row but does not specify width or order.

Confidence: inferred

### Disposition

Consolidate only if the 112px measurement is a durable compatibility constraint.

### Open Questions

- What mobile width range and localization widths must the fixed control support?

## CAND-193: OTP Send-Failure Report

### Source

> - [x] On the Create your account form, when I click Continue and the OTP can't be sent, a report about that is visible

[Source lines 571-571](../../../../backlog/backlog.md#L571)

### Candidate behavior

When sending an OTP fails after `Continue` on the Create your account form, a failure report is visible.

### Applicability

- Actor: account registrant
- Location: Create your account form
- Event kind: not applicable
- Interaction mode: registration
- Viewport: unspecified
- State: OTP send failure
- Exclusions: successful sends and failure-report wording

### Classification

candidate FR

### Existing Requirements and Confidence

Existing requirements: [FR-023](../../functional/fr/FR-023.md) is accepted and covers registration status during sign-in, but not an OTP send failure.

Confidence: inferred

### Disposition

Consolidate with registration feedback behavior if OTP is confirmed as the current registration mechanism.

### Open Questions

- What failures receive a report, and what information may it disclose?

## CAND-194: Timezone-Reset Icon Direction

### Source

> - [x] The icon to reset the timezone shall have a counter-clockwise arrow

[Source lines 572-572](../../../../backlog/backlog.md#L572)

### Candidate behavior

The timezone-reset icon uses a counter-clockwise arrow.

### Applicability

- Actor: event visitor
- Location: timezone reset control
- Event kind: timed
- Interaction mode: viewing
- Viewport: unspecified
- State: reset control visible
- Exclusions: reset behavior and accessible name

### Classification

candidate FR

### Existing Requirements and Confidence

Existing requirements: None identified.

Confidence: inferred

### Disposition

Consolidate only if icon direction carries a product meaning rather than implementation styling.

### Open Questions

- Is an accessible text label required in addition to the icon?

## CAND-195: Timed Sidebar Time-Control Row

### Source

> - [x] On the timed event page, timezone button and time format button shall together span the full row in the sidebar, the time format aligned right.

[Source lines 573-573](../../../../backlog/backlog.md#L573)

### Candidate behavior

On the timed event page, the timezone and time-format buttons together span the full sidebar row, with the time format aligned right.

### Applicability

- Actor: event visitor
- Location: timed event page sidebar
- Event kind: timed
- Interaction mode: viewing
- Viewport: unspecified
- State: time controls visible
- Exclusions: time-control order, dates-only pages, and reset-control behavior

### Classification

candidate FR

### Existing Requirements and Confidence

Existing requirements: [FR-046](../../functional/fr/FR-046.md) and [FR-047](../../functional/fr/FR-047.md) are proposed label requirements, not row layout. Overlap: CAND-196 specifies the separate time-control order; no accepted FR/QR overlap.

Confidence: inferred

### Disposition

Consolidate only after resolving the source's potentially ambiguous alignment wording.

### Open Questions

- Does “time format aligned right” mean its text, button, or allocated area?

## CAND-196: Timed Sidebar Time-Control Order

### Source

> - [x] On the timed event page, the time format shall be on the left from the time zone so that the timezone button can get a reset button or stay the same freely

[Source lines 574-574](../../../../backlog/backlog.md#L574)

### Candidate behavior

On the timed event page, the time-format button is left of the timezone button so the timezone button can accommodate a reset button or remain unchanged.

### Applicability

- Actor: event visitor
- Location: timed event page sidebar
- Event kind: timed
- Interaction mode: viewing
- Viewport: unspecified
- State: time controls visible
- Exclusions: row spanning and alignment, dates-only pages, and reset-control behavior

### Classification

candidate FR

### Existing Requirements and Confidence

Existing requirements: [FR-046](../../functional/fr/FR-046.md) and [FR-047](../../functional/fr/FR-047.md) are proposed label requirements, not control order. Overlap: CAND-195 specifies the separate shared-row and alignment behavior; no accepted FR/QR overlap.

Confidence: inferred

### Disposition

Consolidate with timed-page sidebar layout if the order is durable across responsive layouts.

### Open Questions

- Does this order apply when the timezone reset control is absent?

## CAND-197: Mobile Tooltip Screen Containment

### Source

> - [x] On mobile, on the event page, tooltip shall be fully on-screen

[Source lines 575-575](../../../../backlog/backlog.md#L575)

### Candidate behavior

On mobile event pages, a tooltip remains fully within the visible screen.

### Applicability

- Actor: event visitor
- Location: event page tooltip
- Event kind: unspecified
- Interaction mode: selection or hover-equivalent interaction
- Viewport: mobile
- State: tooltip visible
- Exclusions: desktop placement and tooltip content

### Classification

duplicate or refinement

### Existing Requirements and Confidence

Existing requirements: [FR-020](../../functional/fr/FR-020.md) already requires a visible mobile selected-slot tooltip adjacent to a visible selected slot; it does not expressly require screen containment.

Confidence: inferred

### Disposition

Refine FR-020 if containment is confirmed; do not create a separate requirement.

### Open Questions

- Does this apply to every tooltip or only selected-slot tooltips?

## CAND-198: Mobile Availability Show-All-Hours Placement

### Source

> - [x] On mobile, when adding availability, Show all hours should be an option at its normal place - under other options between the event description and the grid

[Source lines 576-576](../../../../backlog/backlog.md#L576)

### Candidate behavior

On mobile while adding availability, `Show all hours` appears under other options between the event description and the grid.

### Applicability

- Actor: event visitor
- Location: mobile event page
- Event kind: timed
- Interaction mode: adding availability
- Viewport: mobile
- State: other options and event description visible
- Exclusions: editing availability, desktop layout, and dates-only pages

### Classification

candidate FR

### Existing Requirements and Confidence

Existing requirements: [FR-011](../../functional/fr/FR-011.md) covers the expansion behavior; [FR-048](../../functional/fr/FR-048.md) is proposed and covers centering in a no-response state, not this placement.

Confidence: inferred

### Disposition

Consolidate with mobile timed-grid controls if the “normal place” is formally defined.

### Open Questions

- Does “adding availability” include editing an existing availability response?

## CAND-199: Mobile Availability Panel Details

### Source

> - [x] Bug:
>   - on mobile
>   - I add and save availability
>   - I click an active timeslot
>   - Responses panel with my response slides from the bottom
>   - I click Edit availability
>   - I click a grid timeslot
>   - Expected: on the panel that slides from the bottom, I see only Available/if needed
>   - Actual: on the panel, I see both Available/If needed and Responses

[Source lines 577-585](../../../../backlog/backlog.md#L577-L585)

### Candidate behavior

No new requirement behavior asserted; the source is a concrete defect report already covered by the mobile editing rule.

### Applicability

- Actor: event visitor
- Location: mobile event page selected-slot panel
- Event kind: timed
- Interaction mode: availability editing
- Viewport: mobile
- State: saved availability, selected active slot, then Edit availability
- Exclusions: read-mode response details

### Classification

existing requirement

### Existing Requirements and Confidence

Existing requirements: [FR-028](../../functional/fr/FR-028.md) requires availability controls and excludes response details from the mobile selected-slot panel while editing availability.

Confidence: confirmed

### Disposition

Map to FR-028 as a regression example; no new FR.
