# Backlog FR Inventory: Layout And Timed Grid

This durable batch artifact is non-normative. `CAND-*` IDs are temporary review identifiers and will be consolidated by [../backlog-fr-inventory.md](../backlog-fr-inventory.md). Source quotes preserve backlog wording; classifications and applicability are review assessments, not product commitments.

## CAND-001

#### Source

> - [x] use `.env.example` instead of `.env.template`
>   - We use `.env.development.example` and `.env.production.example`

[Source lines 267-268](../../../../backlog/backlog.md#L267-L268)

#### Candidate behavior

No new requirement behavior asserted; this is repository naming.

#### Applicability

Actor: maintainer. Location: repository. Event kind: none. Interaction mode: configuration. Viewport: any. State: environment-template maintenance. Exclusions: runtime behavior.

#### Classification

implementation detail

#### Existing Requirements and Confidence

None. Confidence: inferred.

#### Disposition

Exclude from requirements.

#### Open Questions

None.

## CAND-002

#### Source

> - [x] multi-day <http://127.0.0.1:4173/e/5Ef6f>

[Source lines 269-269](../../../../backlog/backlog.md#L269-L269)

#### Candidate behavior

No new requirement behavior asserted; this is a test reference without an outcome.

#### Applicability

Actor: reviewer. Location: event page. Event kind: unconfirmed. Interaction mode: viewing. Viewport: unconfirmed. State: multi-day. Exclusions: all unspecified behavior.

#### Classification

bug or investigation

#### Existing Requirements and Confidence

Potentially FR-002 and FR-013. Confidence: needs product decision.

#### Disposition

Retain only as investigation provenance.

#### Open Questions

What behavior was being tested at this URL?

## CAND-003

#### Source

> - [x] Let's nest all toggles under Options dropdown and make it open by default
>   - Show best times is always visible, other can be opened via Options

[Source lines 270-271](../../../../backlog/backlog.md#L270-L271)

#### Candidate behavior

No new requirement behavior asserted; this is a completed UI arrangement with unclear enduring scope.

#### Applicability

Actor: event visitor. Location: event page. Event kind: unconfirmed. Interaction mode: viewing. Viewport: unconfirmed. State: responses present. Exclusions: no-response layout.

#### Classification

duplicate or refinement

#### Existing Requirements and Confidence

Potential overlap with FR-009 and FR-011. Confidence: inferred.

#### Disposition

Do not migrate unless product navigation policy is confirmed.

#### Open Questions

Which toggles and event states remain in scope?

## CAND-004

#### Source

> - [x] overlay availabilities - each slot has a solid frame <https://timeful.fun/e/c762cA>

[Source lines 272-272](../../../../backlog/backlog.md#L272-L272)

#### Candidate behavior

No new requirement behavior asserted; this is visual styling of an existing overlay.

#### Applicability

Actor: availability editor. Location: availability editor grid. Event kind: timed. Interaction mode: editing. Viewport: any. State: overlay visible. Exclusions: non-overlay cells.

#### Classification

duplicate or refinement

#### Existing Requirements and Confidence

FR-005 covers the overlay; FR-009 covers distinct grid states. Confidence: inferred.

#### Disposition

Fold into visual regression coverage, not a new requirement.

#### Open Questions

Is a solid frame a durable product constraint or implementation styling?

## CAND-005

#### Source

> - [x] there should be a space between grids for non-consecutive days

[Source lines 273-273](../../../../backlog/backlog.md#L273-L273)

#### Candidate behavior

No new requirement behavior asserted; this is grid layout styling.

#### Applicability

Actor: event visitor. Location: timed grid. Event kind: timed. Interaction mode: viewing. Viewport: any. State: non-consecutive displayed days. Exclusions: consecutive days.

#### Classification

needs product decision

#### Existing Requirements and Confidence

FR-002 defines columns, not spacing. Confidence: inferred.

#### Disposition

Exclude from requirements.

#### Open Questions

None.

## CAND-006

#### Source

> - [x] edit event button is missing

[Source lines 274-274](../../../../backlog/backlog.md#L274-L274)

#### Candidate behavior

No new requirement behavior asserted; a missing-control defect does not establish its permission or scope.

#### Applicability

Actor: unconfirmed. Location: event page. Event kind: unconfirmed. Interaction mode: viewing. Viewport: unconfirmed. State: unconfirmed. Exclusions: unconfirmed.

#### Classification

bug or investigation

#### Existing Requirements and Confidence

Potentially FR-027 for dates-only owners. Confidence: needs product decision.

#### Disposition

Retain as defect provenance.

#### Open Questions

Who should see the control, for which event kinds, and what may it edit?

## CAND-007

#### Source

> - [x] shown in should be the same size as the timezone and black

[Source lines 275-275](../../../../backlog/backlog.md#L275-L275)

#### Candidate behavior

No new requirement behavior asserted; this is control styling.

#### Applicability

Actor: event visitor. Location: event-page display controls. Event kind: timed. Interaction mode: viewing. Viewport: any. State: display timezone shown. Exclusions: editor controls.

#### Classification

needs product decision

#### Existing Requirements and Confidence

FR-013 establishes display-timezone behavior. Confidence: confirmed.

#### Disposition

Exclude from requirements.

#### Open Questions

None.

## CAND-008

#### Source

> - [x] the description text should be at the same vertical position when not editing and when editing

[Source lines 276-276](../../../../backlog/backlog.md#L276-L276)

#### Candidate behavior

No new requirement behavior asserted; this is layout consistency.

#### Applicability

Actor: event visitor. Location: event description. Event kind: any. Interaction mode: viewing and editing. Viewport: any. State: description present. Exclusions: text preservation.

#### Classification

needs product decision

#### Existing Requirements and Confidence

FR-019 covers multiline preservation, not alignment. Confidence: confirmed.

#### Disposition

Exclude from requirements.

#### Open Questions

None.

## CAND-009

#### Source

> - [x] replace the Create event button on the main page with the actual form
>   - No, there's additional useful info about the app on the main page

[Source lines 277-278](../../../../backlog/backlog.md#L277-L278)

#### Candidate behavior

No new requirement behavior asserted; the child records a rejected alternative.

#### Applicability

Actor: platform visitor. Location: main page. Event kind: none. Interaction mode: event creation entry. Viewport: any. State: landing page. Exclusions: event-editor validation.

#### Classification

ADR or decision

#### Existing Requirements and Confidence

FR-025 applies after the new-event form opens. Confidence: confirmed.

#### Disposition

Retain as a historical product-navigation decision.

#### Open Questions

Is this decision still current?

## CAND-010

#### Source

> - [x] on the event page, near "shown in", the underline colors for the timezone and time should be the same

[Source lines 279-279](../../../../backlog/backlog.md#L279-L279)

#### Candidate behavior

No new requirement behavior asserted; this is presentation styling.

#### Applicability

Actor: event visitor. Location: event-page display controls. Event kind: timed. Interaction mode: viewing. Viewport: any. State: controls visible. Exclusions: behavior of those controls.

#### Classification

needs product decision

#### Existing Requirements and Confidence

FR-013 and FR-024 cover control effects. Confidence: confirmed.

#### Disposition

Exclude from requirements.

#### Open Questions

None.

## CAND-011

#### Source

> - [x] move adr to the repo root
>   - No, keep adr for frontend inside frontend

[Source lines 280-281](../../../../backlog/backlog.md#L280-L281)

#### Candidate behavior

No new requirement behavior asserted; the child records a repository-documentation decision.

#### Applicability

Actor: maintainer. Location: repository documentation. Event kind: none. Interaction mode: architecture documentation. Viewport: any. State: ADR placement. Exclusions: product runtime.

#### Classification

ADR or decision

#### Existing Requirements and Confidence

None. Confidence: confirmed.

#### Disposition

Exclude from requirements.

#### Open Questions

None.

## CAND-012

#### Source

> - [x] use full "development" instead of "dev" and "production" instead of "prod"

[Source lines 282-282](../../../../backlog/backlog.md#L282-L282)

#### Candidate behavior

No new requirement behavior asserted; this is naming convention.

#### Applicability

Actor: maintainer. Location: repository. Event kind: none. Interaction mode: configuration. Viewport: any. State: environment naming. Exclusions: runtime behavior.

#### Classification

needs product decision

#### Existing Requirements and Confidence

None. Confidence: inferred.

#### Disposition

Exclude from requirements.

#### Open Questions

None.

## CAND-013

#### Source

> - [x] Can't save time for the first time

[Source lines 283-283](../../../../backlog/backlog.md#L283-L283)

#### Candidate behavior

No new requirement behavior asserted; the object and save path are ambiguous.

#### Applicability

Actor: unconfirmed. Location: unconfirmed. Event kind: unconfirmed. Interaction mode: saving. Viewport: unconfirmed. State: first save. Exclusions: unconfirmed.

#### Classification

needs product decision

#### Existing Requirements and Confidence

None. Confidence: needs product decision.

#### Disposition

Retain as defect provenance.

#### Open Questions

Does “time” mean scheduled event time, active-slot settings, or another field?

## CAND-014

#### Source

> - [x] In create event -> advanced options, "Timezone" should be black

[Source lines 284-284](../../../../backlog/backlog.md#L284-L284)

#### Candidate behavior

No new requirement behavior asserted; this is label styling.

#### Applicability

Actor: event creator. Location: new-event advanced options. Event kind: unconfirmed. Interaction mode: creation. Viewport: any. State: advanced options open. Exclusions: timezone semantics.

#### Classification

needs product decision

#### Existing Requirements and Confidence

FR-027 is limited to dates-only settings; no direct overlap. Confidence: inferred.

#### Disposition

Exclude from requirements.

#### Open Questions

None.

## CAND-015

#### Source

> - [x] Show full grid by default
>   - It's collapsed by default

[Source lines 285-286](../../../../backlog/backlog.md#L285-L286)

#### Candidate behavior

No new requirement behavior asserted; the child rejects the proposed default.

#### Applicability

Actor: event visitor. Location: timed grid. Event kind: timed. Interaction mode: viewing or scheduling. Viewport: any. State: initial grid state. Exclusions: availability editing and specific-times setting.

#### Classification

needs product decision

#### Existing Requirements and Confidence

FR-011 requires collapsed inactive runs when Show all hours is disabled. Confidence: confirmed.

#### Disposition

Map to FR-011; no migration.

#### Open Questions

None.

## CAND-016

#### Source

> - [x] add option to collapse unused hours (not hide)

[Source lines 287-287](../../../../backlog/backlog.md#L287-L287)

#### Candidate behavior

No new requirement behavior asserted; accepted behavior is already specified.

#### Applicability

Actor: event visitor. Location: timed grid. Event kind: timed. Interaction mode: viewing or scheduling. Viewport: any. State: inactive hour runs. Exclusions: availability editing and specific-times setting.

#### Classification

needs product decision

#### Existing Requirements and Confidence

FR-011 specifies collapsing rather than hiding inactive runs. Confidence: confirmed.

#### Disposition

Map to FR-011; no migration.

#### Open Questions

None.

## CAND-017

#### Source

> - [x] make each collapsed hours uncollapsible

[Source lines 288-288](../../../../backlog/backlog.md#L288-L288)

#### Candidate behavior

No new requirement behavior asserted; this is an interaction refinement of collapse behavior.

#### Applicability

Actor: event visitor. Location: timed grid. Event kind: timed. Interaction mode: viewing. Viewport: any. State: collapsed run. Exclusions: Show all hours control.

#### Classification

needs product decision

#### Existing Requirements and Confidence

FR-011 establishes collapse but not per-run expansion. Confidence: confirmed.

#### Disposition

Do not migrate pending product confirmation.

#### Open Questions

Should individual collapsed runs be expandable?

## CAND-018

#### Source

> - [x] mobile version - switching between 3 days and 7 days doesn't work
>   - Works now

[Source lines 289-290](../../../../backlog/backlog.md#L289-L290)

#### Candidate behavior

No new requirement behavior asserted; the child closes a mobile defect without defining expected behavior.

#### Applicability

Actor: mobile event visitor. Location: event page. Event kind: unconfirmed. Interaction mode: view-range switching. Viewport: mobile. State: 3-day or 7-day selection. Exclusions: desktop.

#### Classification

needs product decision

#### Existing Requirements and Confidence

None. Confidence: needs product decision.

#### Disposition

Retain as defect provenance.

#### Open Questions

What control, event kinds, and persistence behavior define the range switch?

## CAND-019

#### Source

> - [x] Add availability only over the responses section so that it can scroll
>   - Moved to the header

[Source lines 291-292](../../../../backlog/backlog.md#L291-L292)

#### Candidate behavior

No new requirement behavior asserted; the child records the implemented placement.

#### Applicability

Actor: event visitor. Location: event page. Event kind: any. Interaction mode: availability creation. Viewport: unconfirmed. State: responses section. Exclusions: availability authorization.

#### Classification

needs product decision

#### Existing Requirements and Confidence

FR-003 concerns deletion only. Confidence: inferred.

#### Disposition

Exclude from requirements.

#### Open Questions

None.

## CAND-020

#### Source

> - [x] When viewing an event, when clicking, dragging the box pointer in the red area, then unclicking, the box disappears

[Source lines 293-293](../../../../backlog/backlog.md#L293-L293)

#### Candidate behavior

No new requirement behavior asserted; this is a rendering defect report.

#### Applicability

Actor: event visitor. Location: event grid. Event kind: timed. Interaction mode: pointer drag. Viewport: unconfirmed. State: red area. Exclusions: outcome unspecified.

#### Classification

needs product decision

#### Existing Requirements and Confidence

Potentially FR-009. Confidence: inferred.

#### Disposition

Retain as defect provenance.

#### Open Questions

What selection or schedule state does the box represent?

## CAND-021

#### Source

> - [x] in a collapsed section, upper line overlaps the original line but the lower line doesn't

[Source lines 294-294](../../../../backlog/backlog.md#L294-L294)

#### Candidate behavior

No new requirement behavior asserted; this is a collapsed-grid visual defect.

#### Applicability

Actor: event visitor. Location: timed grid. Event kind: timed. Interaction mode: viewing. Viewport: any. State: collapsed section. Exclusions: non-collapsed grid.

#### Classification

needs product decision

#### Existing Requirements and Confidence

FR-011 specifies collapse bands and boundaries. Confidence: confirmed.

#### Disposition

Treat as FR-011 implementation regression.

#### Open Questions

None.

## CAND-022

#### Source

> - [x] use the same colors for all grid lines

[Source lines 295-295](../../../../backlog/backlog.md#L295-L295)

#### Candidate behavior

No new requirement behavior asserted; this is visual styling.

#### Applicability

Actor: event visitor. Location: grid lines. Event kind: timed. Interaction mode: viewing. Viewport: any. State: any grid. Exclusions: cell-state colors.

#### Classification

needs product decision

#### Existing Requirements and Confidence

FR-009 requires distinct cell states, not line colors. Confidence: confirmed.

#### Disposition

Exclude from requirements.

#### Open Questions

None.

## CAND-023

#### Source

> - [x] the selection box is almost invisible
>   - Now it's hatched

[Source lines 296-297](../../../../backlog/backlog.md#L296-L297)

#### Candidate behavior

No new requirement behavior asserted; the child names a visual implementation response.

#### Applicability

Actor: event visitor. Location: event grid. Event kind: timed. Interaction mode: selection. Viewport: any. State: selection visible. Exclusions: unavailable states.

#### Classification

needs product decision

#### Existing Requirements and Confidence

FR-009 requires visually distinct grid states. Confidence: inferred.

#### Disposition

Treat as visual regression coverage under FR-009.

#### Open Questions

Is selection visibility independently required beyond state distinction?

## CAND-024

#### Source

> - [x] use the same frequency of dashes for the selection box and the grid separator at half an hour
>   - we use solid selection box

[Source lines 298-299](../../../../backlog/backlog.md#L298-L299)

#### Candidate behavior

No new requirement behavior asserted; the child records a styling decision.

#### Applicability

Actor: event visitor. Location: timed grid. Event kind: timed. Interaction mode: selection. Viewport: any. State: selected box. Exclusions: grid semantics.

#### Classification

needs product decision

#### Existing Requirements and Confidence

FR-009 is limited to state distinction. Confidence: confirmed.

#### Disposition

Exclude from requirements.

#### Open Questions

None.

## CAND-025

#### Source

> - [x] when I click somewhere, the drop-down list in edit availability doesn't disappear

[Source lines 300-300](../../../../backlog/backlog.md#L300-L300)

#### Candidate behavior

No new requirement behavior asserted; this is a control-dismissal defect.

#### Applicability

Actor: availability editor. Location: edit availability. Event kind: any. Interaction mode: pointer click. Viewport: unconfirmed. State: dropdown open. Exclusions: keyboard dismissal.

#### Classification

needs product decision

#### Existing Requirements and Confidence

None. Confidence: inferred.

#### Disposition

Retain as defect provenance.

#### Open Questions

Which dropdown and outside-click boundary apply?

## CAND-026

#### Source

> - [x] let's collapse hours when they're at the start or at the end too. These hours are useless anyway

[Source lines 301-301](../../../../backlog/backlog.md#L301-L301)

#### Candidate behavior

No new requirement behavior asserted; accepted collapse behavior already includes leading and trailing runs.

#### Applicability

Actor: event visitor. Location: timed grid. Event kind: timed. Interaction mode: viewing or scheduling. Viewport: any. State: leading or trailing inactive hours. Exclusions: availability editing and specific-times setting.

#### Classification

needs product decision

#### Existing Requirements and Confidence

FR-011 explicitly covers leading and trailing inactive runs. Confidence: confirmed.

#### Disposition

Map to FR-011; no migration.

#### Open Questions

None.

## CAND-027

#### Source

> - [x] Check +3:30 and +5:45

[Source lines 302-302](../../../../backlog/backlog.md#L302-L302)

#### Candidate behavior

No new requirement behavior asserted; this is a timezone test prompt.

#### Applicability

Actor: reviewer. Location: timed-event rendering. Event kind: timed. Interaction mode: testing. Viewport: any. State: fractional-offset timezone. Exclusions: specified expected outcome.

#### Classification

needs product decision

#### Existing Requirements and Confidence

FR-013 concerns display-timezone projection. Confidence: inferred.

#### Disposition

Retain as test provenance.

#### Open Questions

What outcome and which timezone role were to be checked?

## CAND-028

#### Source

> - [x] when editing event, the week day every letter looks the same as the day of month

[Source lines 303-303](../../../../backlog/backlog.md#L303-L303)

#### Candidate behavior

No new requirement behavior asserted; this is date-picker styling.

#### Applicability

Actor: event editor. Location: event date picker. Event kind: any. Interaction mode: editing. Viewport: any. State: weekday and date labels. Exclusions: date selection behavior.

#### Classification

needs product decision

#### Existing Requirements and Confidence

FR-025 requires a selected date before creation. Confidence: confirmed.

#### Disposition

Exclude from requirements.

#### Open Questions

None.

## CAND-029

#### Source

> - [x] create event with specific availability in +2, 0-4 (day 1), 0-4 (day 2). When opened in 0:00, should see the previous date

[Source lines 304-304](../../../../backlog/backlog.md#L304-L304)

#### Candidate behavior

The event page should show the prior projected date when slots project there in the selected display timezone.

#### Applicability

Actor: event visitor. Location: event-page timed grid. Event kind: timed. Interaction mode: viewing. Viewport: any. State: timezone projection crosses midnight. Exclusions: event-timezone mutation.

#### Classification

needs product decision

#### Existing Requirements and Confidence

FR-002 requires projected date columns; FR-013 preserves instants across display-timezone changes. Confidence: confirmed.

#### Disposition

Map to FR-002 and FR-013; no migration.

#### Open Questions

None.

## CAND-030

#### Source

> - [x] box cursor doesn't follow the mouse in the specific times grid

[Source lines 305-305](../../../../backlog/backlog.md#L305-L305)

#### Candidate behavior

No new requirement behavior asserted; this is an interaction defect without a defined cursor contract.

#### Applicability

Actor: event editor. Location: specific-times grid. Event kind: timed. Interaction mode: pointer movement. Viewport: unconfirmed. State: setting specific times. Exclusions: availability editing.

#### Classification

needs product decision

#### Existing Requirements and Confidence

FR-011 excludes specific-times setting. Confidence: confirmed.

#### Disposition

Retain as defect provenance.

#### Open Questions

What selection outcome must follow the pointer?

## CAND-031

#### Source

> - [x] no Options when there are no responses
>   - now there's an option to show all hours

[Source lines 306-307](../../../../backlog/backlog.md#L306-L307)

#### Candidate behavior

No new requirement behavior asserted; the child states an implemented no-response control.

#### Applicability

Actor: event visitor. Location: event page. Event kind: timed. Interaction mode: viewing. Viewport: unconfirmed. State: no responses. Exclusions: responses present.

#### Classification

needs product decision

#### Existing Requirements and Confidence

FR-011 defines Show all hours; FR-014 defines its full-axis effect. Confidence: confirmed.

#### Disposition

Map Show all hours behavior to FR-011 and FR-014; do not migrate layout.

#### Open Questions

Should other Options controls be absent with no responses?

## CAND-032

#### Source

> - [x] disallow editing past dates
>   - Not disallowing because it's a feature

[Source lines 308-309](../../../../backlog/backlog.md#L308-L309)

#### Candidate behavior

No new requirement behavior asserted; the child rejects a restriction.

#### Applicability

Actor: event editor. Location: event date editing. Event kind: unconfirmed. Interaction mode: editing. Viewport: any. State: past date. Exclusions: future dates.

#### Classification

needs product decision

#### Existing Requirements and Confidence

FR-025 requires at least one picked date but has no past-date limit. Confidence: confirmed.

#### Disposition

Retain as a historical scope decision.

#### Open Questions

Is editing past dates intended for all event kinds and roles?

## CAND-033

#### Source

> - [x] after adding availability, the selected segments are dark green

[Source lines 310-310](../../../../backlog/backlog.md#L310-L310)

#### Candidate behavior

No new requirement behavior asserted; this is a visual-state observation.

#### Applicability

Actor: availability editor. Location: availability grid. Event kind: timed. Interaction mode: adding availability. Viewport: any. State: selected segments. Exclusions: schedule selection.

#### Classification

needs product decision

#### Existing Requirements and Confidence

FR-006 specifies state exclusivity; FR-009 specifies visible state distinction. Confidence: confirmed.

#### Disposition

Exclude from requirements.

#### Open Questions

None.

## CAND-034

#### Source

> - [x] Move "Copy link" closer to the event title

[Source lines 311-311](../../../../backlog/backlog.md#L311-L311)

#### Candidate behavior

No new requirement behavior asserted; this is layout placement.

#### Applicability

Actor: event visitor. Location: event header. Event kind: any. Interaction mode: viewing. Viewport: any. State: event link available. Exclusions: link-copy behavior.

#### Classification

needs product decision

#### Existing Requirements and Confidence

None. Confidence: inferred.

#### Disposition

Exclude from requirements.

#### Open Questions

None.

## CAND-035

#### Source

> - [x] too drastic width change at page width breakpoints (1/3 of the screen)

[Source lines 312-312](../../../../backlog/backlog.md#L312-L312)

#### Candidate behavior

No new requirement behavior asserted; this is responsive-layout tuning.

#### Applicability

Actor: platform visitor. Location: page layout. Event kind: any. Interaction mode: viewing. Viewport: breakpoint transition. State: responsive resize. Exclusions: functional controls.

#### Classification

needs product decision

#### Existing Requirements and Confidence

QR-008 may cover accessible coordination flows, but no explicit width criterion. Confidence: inferred.

#### Disposition

Retain as visual regression provenance.

#### Open Questions

What viewport widths and acceptable layout behavior apply?

## CAND-036

#### Source

> - [x] Make edit event a button with a pencil icon

[Source lines 313-313](../../../../backlog/backlog.md#L313-L313)

#### Candidate behavior

No new requirement behavior asserted; this is control presentation.

#### Applicability

Actor: event editor. Location: event page. Event kind: unconfirmed. Interaction mode: editing entry. Viewport: any. State: edit permitted. Exclusions: permission rules.

#### Classification

needs product decision

#### Existing Requirements and Confidence

FR-027 restricts dates-only timezone settings to the owner. Confidence: inferred.

#### Disposition

Exclude from requirements.

#### Open Questions

None.

## CAND-037

#### Source

> - [x] When no responses, show only add availability and show all hours

[Source lines 314-314](../../../../backlog/backlog.md#L314-L314)

#### Candidate behavior

No new requirement behavior asserted; this is no-response page composition.

#### Applicability

Actor: event visitor. Location: event page. Event kind: timed. Interaction mode: viewing. Viewport: unconfirmed. State: no responses. Exclusions: responses present.

#### Classification

needs product decision

#### Existing Requirements and Confidence

FR-011 and FR-014 define Show all hours behavior, not the complete control set. Confidence: confirmed.

#### Disposition

Do not migrate pending control-set confirmation.

#### Open Questions

Must all other controls be hidden for every no-response event state?

## CAND-038

#### Source

> - [x] show pencils

[Source lines 315-315](../../../../backlog/backlog.md#L315-L315)

#### Candidate behavior

No new requirement behavior asserted; this is ambiguous iconography.

#### Applicability

Actor: unconfirmed. Location: unconfirmed. Event kind: unconfirmed. Interaction mode: editing entry. Viewport: unconfirmed. State: unconfirmed. Exclusions: unconfirmed.

#### Classification

needs product decision

#### Existing Requirements and Confidence

None. Confidence: needs product decision.

#### Disposition

Retain as ambiguous source material.

#### Open Questions

Which controls require pencils and for whom?

## CAND-039

#### Source

> - [x] When one response, doesn't show best times

[Source lines 316-316](../../../../backlog/backlog.md#L316-L316)

#### Candidate behavior

No new requirement behavior asserted; this is a feature visibility defect with undefined product rule.

#### Applicability

Actor: event visitor. Location: event page. Event kind: unconfirmed. Interaction mode: viewing. Viewport: any. State: exactly one response. Exclusions: zero or multiple responses.

#### Classification

needs product decision

#### Existing Requirements and Confidence

No accepted requirement defines best-times visibility. Confidence: needs product decision.

#### Disposition

Retain as investigation provenance.

#### Open Questions

Should best times appear for one response, and what does it calculate?

## CAND-040

#### Source

> - [x] There's no source of truth for dates for specific times and event
>   - Added ADR-012

[Source lines 317-318](../../../../backlog/backlog.md#L317-L318)

#### Candidate behavior

No new requirement behavior asserted; the child directs the concern to an ADR.

#### Applicability

Actor: maintainer. Location: timed-event model. Event kind: timed. Interaction mode: specific-times configuration. Viewport: any. State: date source of truth. Exclusions: user-facing outcome unspecified.

#### Classification

ADR or decision

#### Existing Requirements and Confidence

FR-002, FR-010, and FR-014 use date/domain concepts. Confidence: inferred.

#### Disposition

Retain as ADR provenance; consult ADR-012 during consolidation.

#### Open Questions

What user-visible behavior, if any, does ADR-012 require?

## CAND-041

#### Source

> - [x] do we have recurring events?
>       Yes — it uses a custom TimedRecurrence model with two kinds: specific_dates (explicit date list) and weekly (day-of-week pattern). No iCalendar RRULE support.

[Source lines 319-320](../../../../backlog/backlog.md#L319-L320)

#### Candidate behavior

No new requirement behavior asserted; the child reports an implementation model and exclusion.

#### Applicability

Actor: maintainer. Location: recurrence model. Event kind: timed. Interaction mode: configuration. Viewport: any. State: recurrence. Exclusions: iCalendar RRULE support.

#### Classification

needs product decision

#### Existing Requirements and Confidence

No accepted requirement establishes recurrence. Confidence: inferred.

#### Disposition

Exclude from requirements pending product scope.

#### Open Questions

Is recurrence a supported product behavior and what are its observable rules?

## CAND-042

#### Source

> - [x] In the new event form, in Advanced options, The Timezone text isn't aligned horizontally with "Time increment"
>       The styling differs too

[Source lines 321-322](../../../../backlog/backlog.md#L321-L322)

#### Candidate behavior

No new requirement behavior asserted; this is form styling.

#### Applicability

Actor: event creator. Location: new-event advanced options. Event kind: timed. Interaction mode: creation. Viewport: any. State: controls visible. Exclusions: control semantics.

#### Classification

needs product decision

#### Existing Requirements and Confidence

None. Confidence: inferred.

#### Disposition

Exclude from requirements.

#### Open Questions

None.

## CAND-043

#### Source

> - [x] On the event page without responses, there should be only Add availability and Show all hours, not very wide Add availability and More options
>
>       "Show all hours" should be under "Add availability", as before.
>       The toggle and "Show all hours" text should be centered vertically within their box

[Source lines 323-326](../../../../backlog/backlog.md#L323-L326)

#### Candidate behavior

No new requirement behavior asserted; this combines a pending control-set choice and styling instructions.

#### Applicability

Actor: event visitor. Location: event page. Event kind: timed. Interaction mode: viewing. Viewport: unconfirmed. State: no responses. Exclusions: responses present.

#### Classification

needs product decision

#### Existing Requirements and Confidence

FR-011 and FR-014 cover Show all hours behavior only. Confidence: confirmed.

#### Disposition

Consolidate with CAND-037 and CAND-080 if product policy is confirmed.

#### Open Questions

Must this exact no-response layout apply on mobile and desktop?

## CAND-044

#### Source

> - [x] Create event with 9 - 17, timezone +9 (<http://127.0.0.1:4173/e/ee4Cb>), june 11 and 12.
>       Expected: june 10 and 11 are shown on the event page
>       Actual: june 11, 12, 13 are shown there

[Source lines 327-329](../../../../backlog/backlog.md#L327-L329)

#### Candidate behavior

The event page should render the dates to which timed slots project in the selected display timezone.

#### Applicability

Actor: event visitor. Location: event-page timed grid. Event kind: timed. Interaction mode: viewing. Viewport: any. State: timezone projection crosses dates. Exclusions: changing event timezone.

#### Classification

needs product decision

#### Existing Requirements and Confidence

FR-002 and FR-013 require projected-date columns and preserved instants. Confidence: confirmed.

#### Disposition

Treat as a regression scenario for FR-002 and FR-013.

#### Open Questions

None.

## CAND-045

#### Source

> - [x] Spacing between lines is different when editing and viewing event description

[Source lines 330-330](../../../../backlog/backlog.md#L330-L330)

#### Candidate behavior

No new requirement behavior asserted; this is layout consistency.

#### Applicability

Actor: event visitor. Location: event description. Event kind: any. Interaction mode: viewing and editing. Viewport: any. State: multiline description. Exclusions: text preservation.

#### Classification

needs product decision

#### Existing Requirements and Confidence

FR-019 preserves newline characters, not line spacing. Confidence: confirmed.

#### Disposition

Exclude from requirements.

#### Open Questions

None.

## CAND-046

#### Source

> - [x] Shown in shouldn't affect the time zone

[Source lines 331-331](../../../../backlog/backlog.md#L331-L331)

#### Candidate behavior

Changing the display timezone should not change the event timezone.

#### Applicability

Actor: event visitor. Location: event-page display controls. Event kind: timed. Interaction mode: display-timezone change. Viewport: any. State: timezone selected. Exclusions: explicit event-settings timezone edits.

#### Classification

needs product decision

#### Existing Requirements and Confidence

FR-013 explicitly preserves the event timezone. Confidence: confirmed.

#### Disposition

Map to FR-013; no migration.

#### Open Questions

None.

## CAND-047

#### Source

> - [x] Availability not rendered at <http://127.0.0.1:4173/e/Eb67A>

[Source lines 332-332](../../../../backlog/backlog.md#L332-L332)

#### Candidate behavior

No new requirement behavior asserted; this is a URL-specific rendering defect.

#### Applicability

Actor: event visitor. Location: event page. Event kind: unconfirmed. Interaction mode: viewing. Viewport: unconfirmed. State: referenced event. Exclusions: unspecified availability state.

#### Classification

needs product decision

#### Existing Requirements and Confidence

Potentially FR-005, FR-009, or FR-014. Confidence: needs product decision.

#### Disposition

Retain as defect provenance.

#### Open Questions

Which availability data and rendering mode failed?

## CAND-048

#### Source

> - [x] <http://127.0.0.1:4173/e/Eb67A> shown in GMT -7 shows blank grey columns

[Source lines 333-333](../../../../backlog/backlog.md#L333-L333)

#### Candidate behavior

No new requirement behavior asserted; this is a timezone-rendering defect report.

#### Applicability

Actor: event visitor. Location: event-page timed grid. Event kind: timed. Interaction mode: display-timezone change. Viewport: any. State: GMT -7 projection. Exclusions: event-timezone mutation.

#### Classification

needs product decision

#### Existing Requirements and Confidence

FR-002, FR-009, FR-013, and FR-014 are relevant. Confidence: inferred.

#### Disposition

Treat as regression provenance for those requirements.

#### Open Questions

What cell-state relationship is expected for blank projected columns?

## CAND-049

#### Source

> - [x] <http://127.0.0.1:4173/e/Eb67A> doesn't show the bottom separator

[Source lines 334-334](../../../../backlog/backlog.md#L334-L334)

#### Candidate behavior

No new requirement behavior asserted; this is a grid-border visual defect.

#### Applicability

Actor: event visitor. Location: event-page grid. Event kind: timed. Interaction mode: viewing. Viewport: any. State: referenced event. Exclusions: grid semantics.

#### Classification

needs product decision

#### Existing Requirements and Confidence

FR-009 requires state presentation, not separators. Confidence: inferred.

#### Disposition

Retain as visual regression provenance.

#### Open Questions

None.

## CAND-050

#### Source

> - [x] When there are no responses, Add availability and show all hours are too wide

[Source lines 335-335](../../../../backlog/backlog.md#L335-L335)

#### Candidate behavior

No new requirement behavior asserted; this is no-response layout tuning.

#### Applicability

Actor: event visitor. Location: event page controls. Event kind: timed. Interaction mode: viewing. Viewport: unconfirmed. State: no responses. Exclusions: responses present.

#### Classification

needs product decision

#### Existing Requirements and Confidence

FR-011 and FR-014 cover Show all hours function, not width. Confidence: confirmed.

#### Disposition

Exclude from requirements.

#### Open Questions

None.

## CAND-051

#### Source

> - [x] in What days might work, when sun and mon are selected, when enabling start on monday, both mon and sun must be selected

[Source lines 336-336](../../../../backlog/backlog.md#L336-L336)

#### Candidate behavior

No new requirement behavior asserted; this is an ambiguous date-selection rule.

#### Applicability

Actor: event editor. Location: What days might work. Event kind: unconfirmed. Interaction mode: week-start change. Viewport: any. State: Sunday and Monday selected. Exclusions: other selected dates.

#### Classification

needs product decision

#### Existing Requirements and Confidence

FR-025 requires at least one picked date; it does not define week-start effects. Confidence: confirmed.

#### Disposition

Do not migrate pending domain-rule confirmation.

#### Open Questions

Why should changing week start select dates, and does this apply to all date-picker modes?

## CAND-052

#### Source

> - [x] remove formerly known as schej and flag

[Source lines 337-337](../../../../backlog/backlog.md#L337-L337)

#### Candidate behavior

No new requirement behavior asserted; this is branding copy cleanup.

#### Applicability

Actor: platform visitor. Location: unconfirmed. Event kind: none. Interaction mode: viewing. Viewport: any. State: branding text present. Exclusions: product behavior.

#### Classification

needs product decision

#### Existing Requirements and Confidence

None. Confidence: inferred.

#### Disposition

Exclude from requirements.

#### Open Questions

None.

## CAND-053

#### Source

> - [x] on the event page, responses should be aligned top with the grid top

[Source lines 338-338](../../../../backlog/backlog.md#L338-L338)

#### Candidate behavior

No new requirement behavior asserted; this is page alignment styling.

#### Applicability

Actor: event visitor. Location: event page. Event kind: any. Interaction mode: viewing. Viewport: any. State: responses and grid shown. Exclusions: content behavior.

#### Classification

needs product decision

#### Existing Requirements and Confidence

None. Confidence: inferred.

#### Disposition

Exclude from requirements.

#### Open Questions

None.

## CAND-054

#### Source

> - [x] add flag to enable privacy policy

[Source lines 339-339](../../../../backlog/backlog.md#L339-L339)

#### Candidate behavior

No new requirement behavior asserted; this names a configuration mechanism without an observable policy outcome.

#### Applicability

Actor: maintainer. Location: deployment configuration. Event kind: none. Interaction mode: configuration. Viewport: any. State: privacy policy feature flag. Exclusions: policy content.

#### Classification

needs product decision

#### Existing Requirements and Confidence

None. Confidence: needs product decision.

#### Disposition

Do not migrate pending product and compliance scope.

#### Open Questions

When enabled, where must the privacy policy appear and who controls the flag?

## CAND-055

#### Source

<!-- prettier-ignore -->
> - [x] use the same palette for event page
>   (enabled - light-grey, active -red; disabled - dark-grey)

[Source lines 340-341](../../../../backlog/backlog.md#L340-L341)

#### Candidate behavior

No new requirement behavior asserted; the child gives styling tokens for existing state distinctions.

#### Applicability

Actor: event visitor. Location: event-page timed grid. Event kind: timed. Interaction mode: viewing. Viewport: any. State: enabled, active, or disabled cells. Exclusions: custom-domain editing colors where separately specified.

#### Classification

needs product decision

#### Existing Requirements and Confidence

FR-009 and FR-014 specify distinct state presentation, including light-grey and dark-grey treatments. Confidence: confirmed.

#### Disposition

Map behavioral state meaning to FR-009 and FR-014; retain exact palette as styling.

#### Open Questions

Does “active -red” mean unavailable active slots only?

## CAND-056

#### Source

> - [x] legend should be visible even if no responses. Show only enabled/active

[Source lines 342-342](../../../../backlog/backlog.md#L342-L342)

#### Candidate behavior

No new requirement behavior asserted; this refines legend visibility beyond accepted state rules.

#### Applicability

Actor: event visitor. Location: event-page legend. Event kind: timed. Interaction mode: viewing. Viewport: any. State: no responses. Exclusions: response-derived legend states.

#### Classification

needs product decision

#### Existing Requirements and Confidence

FR-009 requires a mode-appropriate legend but does not require it when no responses exist. Confidence: confirmed.

#### Disposition

Do not migrate pending no-response legend policy.

#### Open Questions

Must the legend be visible in every no-response viewport and grid mode?

## CAND-057

#### Source

> - [x] show all hours should work on the event page

[Source lines 343-343](../../../../backlog/backlog.md#L343-L343)

#### Candidate behavior

No new requirement behavior asserted; accepted requirements already define the control effect.

#### Applicability

Actor: event visitor. Location: event page. Event kind: timed. Interaction mode: viewing. Viewport: any. State: Show all hours enabled. Exclusions: availability editing.

#### Classification

needs product decision

#### Existing Requirements and Confidence

FR-011 and FR-014 require the full civil-day axis under Show all hours. Confidence: confirmed.

#### Disposition

Map to FR-011 and FR-014; no migration.

#### Open Questions

None.

## CAND-058

#### Source

> - [x] hide dates where there's nothing to pick?
>   - no, keep all dates

[Source lines 344-345](../../../../backlog/backlog.md#L344-L345)

#### Candidate behavior

No new requirement behavior asserted; the child rejects date hiding.

#### Applicability

Actor: event editor. Location: timed grid or date selection. Event kind: timed. Interaction mode: viewing or editing. Viewport: any. State: date without selectable slots. Exclusions: unknown grid mode.

#### Classification

needs product decision

#### Existing Requirements and Confidence

FR-010 requires every picked date to be shown in custom domain editing. Confidence: inferred.

#### Disposition

Map the established custom-domain scope to FR-010; retain broader scope as unresolved.

#### Open Questions

Does “all dates” apply outside custom domain editing?

## CAND-059

#### Source

<!-- prettier-ignore -->
> - [x] the day ends at 24
>    No, at `00:00` for consistency with 12AM in 12-hour format

[Source lines 346-347](../../../../backlog/backlog.md#L346-L347)

#### Candidate behavior

No new requirement behavior asserted; the child records time-label formatting choice.

#### Applicability

Actor: event visitor. Location: timed-grid axis. Event kind: timed. Interaction mode: viewing. Viewport: any. State: day boundary. Exclusions: time-format persistence.

#### Classification

needs product decision

#### Existing Requirements and Confidence

FR-024 separates event and display time formats. Confidence: inferred.

#### Disposition

Retain as presentation decision.

#### Open Questions

Does the `00:00` decision apply in both display formats and editor forms?

## CAND-060

#### Source

> - [x] there should be a label for each grid row that marks the start of an hour, even for collapsed

[Source lines 348-348](../../../../backlog/backlog.md#L348-L348)

#### Candidate behavior

No new requirement behavior asserted; accepted requirements already specify collapsed-band boundaries, not every row label.

#### Applicability

Actor: event visitor. Location: timed-grid time axis. Event kind: timed. Interaction mode: viewing or scheduling. Viewport: any. State: normal or collapsed rows. Exclusions: availability editing.

#### Classification

needs product decision

#### Existing Requirements and Confidence

FR-011 requires each collapsed band's start boundary. Confidence: confirmed.

#### Disposition

Do not migrate the every-row extension pending confirmation.

#### Open Questions

Must every visible hour row be labeled, or only collapsed boundaries?

## CAND-061

#### Source

<!-- prettier-ignore -->
> - [x] get rid of 12-hour format?
>    No

[Source lines 349-350](../../../../backlog/backlog.md#L349-L350)

#### Candidate behavior

No new requirement behavior asserted; the child retains 12-hour format as an option.

#### Applicability

Actor: event visitor or editor. Location: time-format controls. Event kind: timed. Interaction mode: format selection. Viewport: any. State: 12-hour format. Exclusions: default format unspecified.

#### Classification

needs product decision

#### Existing Requirements and Confidence

FR-024 defines independent event and display time formats. Confidence: confirmed.

#### Disposition

Map to FR-024; no migration.

#### Open Questions

None.

## CAND-062

#### Source

<!-- prettier-ignore -->
> - [x] remove Add availability, leave just edit availability
>    No. In this case, we won't be able to add availability for someone

[Source lines 351-352](../../../../backlog/backlog.md#L351-L352)

#### Candidate behavior

No new requirement behavior asserted; the child rejects removal based on a use case without defining authorization.

#### Applicability

Actor: unconfirmed. Location: event page. Event kind: any. Interaction mode: availability creation. Viewport: any. State: adding for another person. Exclusions: editing existing response.

#### Classification

needs product decision

#### Existing Requirements and Confidence

FR-003 covers deletion; no accepted requirement defines this entry-point policy. Confidence: needs product decision.

#### Disposition

Do not migrate pending actor and access-policy confirmation.

#### Open Questions

Who may add availability for someone else, and how is that response owned?

## CAND-063

#### Source

> - [x] don't uncollapse rows when scheduling

[Source lines 353-353](../../../../backlog/backlog.md#L353-L353)

#### Candidate behavior

No new requirement behavior asserted; accepted behavior already preserves collapsed bands during scheduling.

#### Applicability

Actor: event visitor. Location: timed grid. Event kind: timed. Interaction mode: scheduling. Viewport: any. State: collapsed inactive runs. Exclusions: availability editing and specific-times setting.

#### Classification

needs product decision

#### Existing Requirements and Confidence

FR-011 explicitly says scheduling does not expand collapsed bands. Confidence: confirmed.

#### Disposition

Map to FR-011; no migration.

#### Open Questions

None.

## CAND-064

#### Source

> - [x] show event when rescheduling

[Source lines 354-354](../../../../backlog/backlog.md#L354-L354)

#### Candidate behavior

No new requirement behavior asserted; “event” and rescheduling context are ambiguous.

#### Applicability

Actor: event visitor. Location: unconfirmed. Event kind: unconfirmed. Interaction mode: rescheduling. Viewport: any. State: scheduled event exists. Exclusions: initial scheduling.

#### Classification

needs product decision

#### Existing Requirements and Confidence

FR-012 allows saving, replacing, and clearing scheduled event time. Confidence: inferred.

#### Disposition

Do not migrate pending observable-outcome definition.

#### Open Questions

What must be shown when replacing a scheduled event time?

## CAND-065

#### Source

> - [x] Schedule button on mobile should be filled blueish when clicked

[Source lines 355-355](../../../../backlog/backlog.md#L355-L355)

#### Candidate behavior

No new requirement behavior asserted; this is mobile button styling.

#### Applicability

Actor: mobile event visitor. Location: schedule button. Event kind: timed. Interaction mode: pointer or touch click. Viewport: mobile. State: clicked. Exclusions: schedule action outcome.

#### Classification

needs product decision

#### Existing Requirements and Confidence

FR-012 covers scheduling persistence, not button color. Confidence: inferred.

#### Disposition

Exclude from requirements.

#### Open Questions

None.

## CAND-066

#### Source

> - [x] When there's a scheduled event and I edit availability, I shold draw only with green cells, not blue cells (scheduled)

[Source lines 356-356](../../../../backlog/backlog.md#L356-L356)

#### Candidate behavior

No new requirement behavior asserted; this is a state-presentation refinement with colors as implementation language.

#### Applicability

Actor: availability editor. Location: availability editor grid. Event kind: timed. Interaction mode: editing availability. Viewport: any. State: scheduled event exists. Exclusions: viewing or scheduling mode.

#### Classification

needs product decision

#### Existing Requirements and Confidence

FR-009 requires context-specific visual states; FR-012 defines scheduled event time. Confidence: inferred.

#### Disposition

Do not migrate pending editing-state semantics.

#### Open Questions

Should scheduled time be hidden, visually distinct, or non-interactive during availability editing?

## CAND-067

#### Source

> - [x] when editing availability, overlay availabilities should preserve the marked time slots and not shift

[Source lines 357-357](../../../../backlog/backlog.md#L357-L357)

#### Candidate behavior

The edited availability overlay should remain aligned with its marked timed slots during availability editing.

#### Applicability

Actor: availability editor. Location: availability editor grid. Event kind: timed. Interaction mode: editing availability. Viewport: any. State: overlay enabled. Exclusions: non-overlay rendering.

#### Classification

needs product decision

#### Existing Requirements and Confidence

FR-005 requires the edited response overlay; FR-013 preserves timed-slot instants under display-timezone changes. Confidence: inferred.

#### Disposition

Treat as a regression scenario for FR-005 and FR-013.

#### Open Questions

Does “not shift” specifically concern timezone changes, scrolling, or grid rerenders?

## CAND-068

#### Source

> - [x] set specific times - just select, don't draw scheduled event

[Source lines 358-358](../../../../backlog/backlog.md#L358-L358)

#### Candidate behavior

No new requirement behavior asserted; this separates two ambiguous modes without defining their outcomes.

#### Applicability

Actor: event editor. Location: specific-times grid. Event kind: timed. Interaction mode: setting specific times. Viewport: any. State: selection. Exclusions: scheduling mode.

#### Classification

needs product decision

#### Existing Requirements and Confidence

FR-011 excludes specific-times setting; FR-012 defines scheduled event time. Confidence: inferred.

#### Disposition

Do not migrate pending mode definitions.

#### Open Questions

What selection is saved by specific-times mode, and what visual feedback is required?

## CAND-069

#### Source

> - [x] on mobile, edit availability is permanently greenish although should be like that only after a click
>   - It's ok - it's disabled

[Source lines 359-360](../../../../backlog/backlog.md#L359-L360)

#### Candidate behavior

No new requirement behavior asserted; the child attributes the visual state to disabled status.

#### Applicability

Actor: mobile event visitor. Location: edit-availability control. Event kind: unconfirmed. Interaction mode: viewing. Viewport: mobile. State: disabled. Exclusions: enabled control feedback.

#### Classification

needs product decision

#### Existing Requirements and Confidence

None. Confidence: needs product decision.

#### Disposition

Retain as control-state provenance.

#### Open Questions

When and why is edit availability disabled on mobile?

## CAND-070

#### Source

> - [x] on mobile and desktop, there should be no delay between selecting the time slot and seeing the tooltip with time and date

[Source lines 361-361](../../../../backlog/backlog.md#L361-L361)

#### Candidate behavior

No new requirement behavior asserted; this requests a latency target without a measurable threshold.

#### Applicability

Actor: event visitor. Location: timed grid tooltip. Event kind: timed. Interaction mode: slot selection. Viewport: mobile and desktop. State: slot selected. Exclusions: no selected slot.

#### Classification

needs product decision

#### Existing Requirements and Confidence

FR-020 requires the mobile tooltip when the selected slot is visible; QR-009 is the relevant performance family. Confidence: needs product decision.

#### Disposition

Hold for a measurable responsiveness scenario.

#### Open Questions

What maximum selection-to-tooltip latency is acceptable, and which devices are in scope?

## CAND-071

#### Source

> - [x] on mobile, I shouldn't be able to click the grid through that Responses offcanvas

[Source lines 362-362](../../../../backlog/backlog.md#L362-L362)

#### Candidate behavior

On mobile, an open Responses offcanvas should prevent interactions with the grid behind it.

#### Applicability

Actor: mobile event visitor. Location: event page. Event kind: timed. Interaction mode: touch or click. Viewport: mobile. State: Responses offcanvas open. Exclusions: offcanvas closed.

#### Classification

needs product decision

#### Existing Requirements and Confidence

FR-028 concerns the selected-slot panel while editing but does not define interaction blocking. Confidence: inferred.

#### Disposition

Candidate for consolidation if the offcanvas scope is confirmed.

#### Open Questions

Does blocking apply to all background controls, keyboard actions, and scroll?

## CAND-072

#### Source

> - [x] on mobile, the tooltip should stay near selected timeslot and shouldn't move around
>   - now we store the info about the selected timeslot and render based on its position

[Source lines 363-364](../../../../backlog/backlog.md#L363-L364)

#### Candidate behavior

No new requirement behavior asserted; the child supplies an implementation approach for an already accepted adjacent-tooltip outcome.

#### Applicability

Actor: mobile event visitor. Location: timed grid tooltip. Event kind: timed. Interaction mode: slot selection. Viewport: mobile. State: slot selected and visible. Exclusions: no selected slot.

#### Classification

needs product decision

#### Existing Requirements and Confidence

FR-020 requires the visible selected-slot tooltip adjacent to that slot. Confidence: confirmed.

#### Disposition

Map to FR-020; exclude stored-position mechanism.

#### Open Questions

None.

## CAND-073

#### Source

> - [x] in commit skill, make the script for session identifier and time a single script

[Source lines 365-365](../../../../backlog/backlog.md#L365-L365)

#### Candidate behavior

No new requirement behavior asserted; this is tooling maintenance.

#### Applicability

Actor: maintainer. Location: commit skill. Event kind: none. Interaction mode: committing. Viewport: any. State: session metadata generation. Exclusions: product runtime.

#### Classification

needs product decision

#### Existing Requirements and Confidence

None. Confidence: confirmed.

#### Disposition

Exclude from requirements.

#### Open Questions

None.

## CAND-074

#### Source

> - [x] on mobile, when scrolling, the tooltip with time should scroll together with selected cell, not stay frozen at the same height of the screen while the screen scrolls

[Source lines 366-366](../../../../backlog/backlog.md#L366-L366)

#### Candidate behavior

No new requirement behavior asserted; accepted behavior already specifies tooltip availability adjacent to a visible selected slot after scrolling.

#### Applicability

Actor: mobile event visitor. Location: timed grid tooltip. Event kind: timed. Interaction mode: scrolling. Viewport: mobile. State: selected slot. Exclusions: no selected slot or selected slot not visible.

#### Classification

needs product decision

#### Existing Requirements and Confidence

FR-020 requires adjacency when the selected slot is visible. Confidence: confirmed.

#### Disposition

Map to FR-020; no migration.

#### Open Questions

None.

## CAND-075

#### Source

> - [x] on mobile, when no slot is selected, the tooltip shouldn't be visible

[Source lines 367-367](../../../../backlog/backlog.md#L367-L367)

#### Candidate behavior

No new requirement behavior asserted; accepted behavior already prohibits the tooltip without a selection.

#### Applicability

Actor: mobile event visitor. Location: timed grid tooltip. Event kind: timed. Interaction mode: viewing. Viewport: mobile. State: no slot selected. Exclusions: selected slot visible.

#### Classification

needs product decision

#### Existing Requirements and Confidence

FR-020 explicitly prohibits a tooltip when no timed slot is selected. Confidence: confirmed.

#### Disposition

Map to FR-020; no migration.

#### Open Questions

None.

## CAND-076

#### Source

> - [x] When a timeslot is selected, it's position is saved.
>       When the page is reloaded, the selection is rendered at the timeslot.
>       However, there's no tooltip near the selection.

[Source lines 368-370](../../../../backlog/backlog.md#L368-L370)

#### Candidate behavior

No new requirement behavior asserted; this is a persisted-selection tooltip regression, and persistence itself is not established.

#### Applicability

Actor: mobile event visitor. Location: timed grid tooltip. Event kind: timed. Interaction mode: reload. Viewport: mobile. State: previously selected slot restored and visible. Exclusions: selection persistence policy.

#### Classification

needs product decision

#### Existing Requirements and Confidence

FR-020 requires an adjacent tooltip for a visible selected slot, but not selection persistence. Confidence: confirmed.

#### Disposition

Treat missing tooltip as FR-020 regression; retain persistence scope for decision.

#### Open Questions

Should selected slots persist across reloads?

## CAND-077

#### Source

> - [x] on mobile, when long press inside the grid changes the selected timeslot, the tooltip must also appear near the selected timeslot
>   - currently, it stays near the previously selected timeslot

[Source lines 371-372](../../../../backlog/backlog.md#L371-L372)

#### Candidate behavior

No new requirement behavior asserted; accepted behavior already requires adjacency to the selected visible slot.

#### Applicability

Actor: mobile event visitor. Location: timed grid tooltip. Event kind: timed. Interaction mode: long press. Viewport: mobile. State: selected slot changes and is visible. Exclusions: no selected slot.

#### Classification

needs product decision

#### Existing Requirements and Confidence

FR-020 requires the tooltip adjacent to the visible selected slot. Confidence: confirmed.

#### Disposition

Treat as FR-020 regression.

#### Open Questions

None.

## CAND-078

#### Source

> - [x] on mobile, when long press inside the grid makes the timeslot selected, the tooltip must appear near it.
>   - currently, the tooltip doesn't appear

[Source lines 373-374](../../../../backlog/backlog.md#L373-L374)

#### Candidate behavior

No new requirement behavior asserted; accepted behavior already requires the tooltip for a visible selected slot.

#### Applicability

Actor: mobile event visitor. Location: timed grid tooltip. Event kind: timed. Interaction mode: long press. Viewport: mobile. State: slot newly selected and visible. Exclusions: no selected slot.

#### Classification

needs product decision

#### Existing Requirements and Confidence

FR-020 requires the adjacent tooltip. Confidence: confirmed.

#### Disposition

Treat as FR-020 regression.

#### Open Questions

None.

## CAND-079

#### Source

> - [x] on mobile, the tooltip should be under top navbar when the grid gets scrolled up

[Source lines 375-375](../../../../backlog/backlog.md#L375-L375)

#### Candidate behavior

No new requirement behavior asserted; this is a placement refinement not established by the accepted adjacency rule.

#### Applicability

Actor: mobile event visitor. Location: timed grid tooltip. Event kind: timed. Interaction mode: scrolling. Viewport: mobile. State: grid scrolled upward. Exclusions: unscrolled grid and no selected slot.

#### Classification

needs product decision

#### Existing Requirements and Confidence

FR-020 requires adjacency when the selected slot is visible, not a navbar relationship. Confidence: confirmed.

#### Disposition

Do not migrate pending placement policy.

#### Open Questions

If the selected slot is under the navbar, should the tooltip be hidden, clipped, or repositioned?

## CAND-080

#### Source

> - [x] On the event page without responses, there should be only Add availability and Show all hours, not very wide Add availability and More options
>
>       "Show all hours" should be under "Add availability", as before.
>       The toggle and "Show all hours" text should be centered vertically within their box

[Source lines 376-379](../../../../backlog/backlog.md#L376-L379)

#### Candidate behavior

No new requirement behavior asserted; this duplicates CAND-043 and combines pending policy with layout styling.

#### Applicability

Actor: event visitor. Location: event page. Event kind: timed. Interaction mode: viewing. Viewport: unconfirmed. State: no responses. Exclusions: responses present.

#### Classification

duplicate or refinement

#### Existing Requirements and Confidence

Same overlap as CAND-043: FR-011 and FR-014 cover Show all hours function only. Confidence: confirmed.

#### Disposition

Deduplicate into CAND-043 during consolidation.

#### Open Questions

See CAND-043.
