# Completed Backlog Requirement Candidate Inventory

This consolidated inventory is non-normative. Accepted `FR-*` and `QR-*`
records are canonical requirements. `CAND-*` identifiers are temporary review
and traceability identifiers and never permanent requirement IDs. The four
durable review artifacts are [Batch 01](backlog-fr-inventory-batches/01-layout-and-timed-grid.md),
[Batch 02](backlog-fr-inventory-batches/02-event-page-interaction.md),
[Batch 03](backlog-fr-inventory-batches/03-dates-only-and-platform.md), and
[Batch 04](backlog-fr-inventory-batches/04-infrastructure-and-final-ui.md).
See the [migration inventory guide](README.md) for authority, schema, and
review rules.

## CAND-001

#### Source

> - [x] use `.env.example` instead of `.env.template`
>   - We use `.env.development.example` and `.env.production.example`

[Source lines 267-268](../../../backlog/backlog.md#L267-L268)

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

[Source lines 269-269](../../../backlog/backlog.md#L269-L269)

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

[Source lines 270-271](../../../backlog/backlog.md#L270-L271)

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

[Source lines 272-272](../../../backlog/backlog.md#L272-L272)

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

[Source lines 273-273](../../../backlog/backlog.md#L273-L273)

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

[Source lines 274-274](../../../backlog/backlog.md#L274-L274)

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

[Source lines 275-275](../../../backlog/backlog.md#L275-L275)

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

[Source lines 276-276](../../../backlog/backlog.md#L276-L276)

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

[Source lines 277-278](../../../backlog/backlog.md#L277-L278)

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

[Source lines 279-279](../../../backlog/backlog.md#L279-L279)

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

[Source lines 280-281](../../../backlog/backlog.md#L280-L281)

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

[Source lines 282-282](../../../backlog/backlog.md#L282-L282)

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

[Source lines 283-283](../../../backlog/backlog.md#L283-L283)

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

[Source lines 284-284](../../../backlog/backlog.md#L284-L284)

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

[Source lines 285-286](../../../backlog/backlog.md#L285-L286)

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

[Source lines 287-287](../../../backlog/backlog.md#L287-L287)

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

[Source lines 288-288](../../../backlog/backlog.md#L288-L288)

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

[Source lines 289-290](../../../backlog/backlog.md#L289-L290)

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

[Source lines 291-292](../../../backlog/backlog.md#L291-L292)

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

[Source lines 293-293](../../../backlog/backlog.md#L293-L293)

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

[Source lines 294-294](../../../backlog/backlog.md#L294-L294)

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

[Source lines 295-295](../../../backlog/backlog.md#L295-L295)

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

[Source lines 296-297](../../../backlog/backlog.md#L296-L297)

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

[Source lines 298-299](../../../backlog/backlog.md#L298-L299)

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

[Source lines 300-300](../../../backlog/backlog.md#L300-L300)

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

[Source lines 301-301](../../../backlog/backlog.md#L301-L301)

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

[Source lines 302-302](../../../backlog/backlog.md#L302-L302)

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

[Source lines 303-303](../../../backlog/backlog.md#L303-L303)

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

[Source lines 304-304](../../../backlog/backlog.md#L304-L304)

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

[Source lines 305-305](../../../backlog/backlog.md#L305-L305)

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

[Source lines 306-307](../../../backlog/backlog.md#L306-L307)

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

[Source lines 308-309](../../../backlog/backlog.md#L308-L309)

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

[Source lines 310-310](../../../backlog/backlog.md#L310-L310)

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

[Source lines 311-311](../../../backlog/backlog.md#L311-L311)

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

[Source lines 312-312](../../../backlog/backlog.md#L312-L312)

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

[Source lines 313-313](../../../backlog/backlog.md#L313-L313)

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

[Source lines 314-314](../../../backlog/backlog.md#L314-L314)

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

[Source lines 315-315](../../../backlog/backlog.md#L315-L315)

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

[Source lines 316-316](../../../backlog/backlog.md#L316-L316)

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

[Source lines 317-318](../../../backlog/backlog.md#L317-L318)

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

[Source lines 319-320](../../../backlog/backlog.md#L319-L320)

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

[Source lines 321-322](../../../backlog/backlog.md#L321-L322)

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

[Source lines 323-326](../../../backlog/backlog.md#L323-L326)

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

[Source lines 327-329](../../../backlog/backlog.md#L327-L329)

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

[Source lines 330-330](../../../backlog/backlog.md#L330-L330)

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

[Source lines 331-331](../../../backlog/backlog.md#L331-L331)

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

[Source lines 332-332](../../../backlog/backlog.md#L332-L332)

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

[Source lines 333-333](../../../backlog/backlog.md#L333-L333)

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

[Source lines 334-334](../../../backlog/backlog.md#L334-L334)

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

[Source lines 335-335](../../../backlog/backlog.md#L335-L335)

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

[Source lines 336-336](../../../backlog/backlog.md#L336-L336)

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

[Source lines 337-337](../../../backlog/backlog.md#L337-L337)

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

[Source lines 338-338](../../../backlog/backlog.md#L338-L338)

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

[Source lines 339-339](../../../backlog/backlog.md#L339-L339)

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

[Source lines 340-341](../../../backlog/backlog.md#L340-L341)

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

[Source lines 342-342](../../../backlog/backlog.md#L342-L342)

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

[Source lines 343-343](../../../backlog/backlog.md#L343-L343)

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

[Source lines 344-345](../../../backlog/backlog.md#L344-L345)

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

[Source lines 346-347](../../../backlog/backlog.md#L346-L347)

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

[Source lines 348-348](../../../backlog/backlog.md#L348-L348)

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

[Source lines 349-350](../../../backlog/backlog.md#L349-L350)

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

[Source lines 351-352](../../../backlog/backlog.md#L351-L352)

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

[Source lines 353-353](../../../backlog/backlog.md#L353-L353)

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

[Source lines 354-354](../../../backlog/backlog.md#L354-L354)

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

[Source lines 355-355](../../../backlog/backlog.md#L355-L355)

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

[Source lines 356-356](../../../backlog/backlog.md#L356-L356)

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

[Source lines 357-357](../../../backlog/backlog.md#L357-L357)

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

[Source lines 358-358](../../../backlog/backlog.md#L358-L358)

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

[Source lines 359-360](../../../backlog/backlog.md#L359-L360)

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

[Source lines 361-361](../../../backlog/backlog.md#L361-L361)

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

[Source lines 362-362](../../../backlog/backlog.md#L362-L362)

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

[Source lines 363-364](../../../backlog/backlog.md#L363-L364)

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

[Source lines 365-365](../../../backlog/backlog.md#L365-L365)

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

[Source lines 366-366](../../../backlog/backlog.md#L366-L366)

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

[Source lines 367-367](../../../backlog/backlog.md#L367-L367)

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

[Source lines 368-370](../../../backlog/backlog.md#L368-L370)

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

[Source lines 371-372](../../../backlog/backlog.md#L371-L372)

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

[Source lines 373-374](../../../backlog/backlog.md#L373-L374)

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

[Source lines 375-375](../../../backlog/backlog.md#L375-L375)

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

[Source lines 376-379](../../../backlog/backlog.md#L376-L379)

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

## CAND-081

#### Source

> - [x] "Best times" toggle should appear when there is at least one response, not more than one

[Source lines 381-381](../../../backlog/backlog.md#L381-L381)

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

[Source lines 382-382](../../../backlog/backlog.md#L382-L382)

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

[Source lines 383-383](../../../backlog/backlog.md#L383-L383)

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

[Source lines 384-384](../../../backlog/backlog.md#L384-L384)

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

[Source lines 385-385](../../../backlog/backlog.md#L385-L385)

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

[Source lines 386-386](../../../backlog/backlog.md#L386-L386)

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

[Source lines 387-387](../../../backlog/backlog.md#L387-L387)

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

[Source lines 388-389](../../../backlog/backlog.md#L388-L389)

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

[Source lines 390-390](../../../backlog/backlog.md#L390-L390)

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

[Source lines 391-391](../../../backlog/backlog.md#L391-L391)

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

[Source lines 392-392](../../../backlog/backlog.md#L392-L392)

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

[Source lines 393-397](../../../backlog/backlog.md#L393-L397)

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

[Source lines 398-398](../../../backlog/backlog.md#L398-L398)

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

[Source lines 399-399](../../../backlog/backlog.md#L399-L399)

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

[Source lines 400-403](../../../backlog/backlog.md#L400-L403)

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

[Source lines 404-408](../../../backlog/backlog.md#L404-L408)

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

[Source lines 409-410](../../../backlog/backlog.md#L409-L410)

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

[Source lines 411-411](../../../backlog/backlog.md#L411-L411)

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

[Source lines 412-425](../../../backlog/backlog.md#L412-L425)

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

[Source lines 426-426](../../../backlog/backlog.md#L426-L426)

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

[Source lines 427-427](../../../backlog/backlog.md#L427-L427)

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

[Source lines 428-429](../../../backlog/backlog.md#L428-L429)

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

[Source lines 430-432](../../../backlog/backlog.md#L430-L432)

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

[Source lines 433-433](../../../backlog/backlog.md#L433-L433)

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

[Source lines 434-434](../../../backlog/backlog.md#L434-L434)

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

[Source lines 435-435](../../../backlog/backlog.md#L435-L435)

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

[Source lines 436-439](../../../backlog/backlog.md#L436-L439)

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

[Source lines 440-440](../../../backlog/backlog.md#L440-L440)

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

[Source lines 441-441](../../../backlog/backlog.md#L441-L441)

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

[Source lines 442-442](../../../backlog/backlog.md#L442-L442)

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

[Source lines 443-443](../../../backlog/backlog.md#L443-L443)

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

[Source lines 444-444](../../../backlog/backlog.md#L444-L444)

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

[Source lines 445-445](../../../backlog/backlog.md#L445-L445)

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

[Source lines 446-448](../../../backlog/backlog.md#L446-L448)

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

[Source lines 449-449](../../../backlog/backlog.md#L449-L449)

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

[Source lines 450-450](../../../backlog/backlog.md#L450-L450)

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

[Source lines 451-451](../../../backlog/backlog.md#L451-L451)

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

[Source lines 452-452](../../../backlog/backlog.md#L452-L452)

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

[Source lines 453-453](../../../backlog/backlog.md#L453-L453)

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

[Source lines 454-454](../../../backlog/backlog.md#L454-L454)

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

[Source lines 455-457](../../../backlog/backlog.md#L455-L457)

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

[Source lines 458-458](../../../backlog/backlog.md#L458-L458)

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

[Source lines 459-459](../../../backlog/backlog.md#L459-L459)

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

[Source lines 460-460](../../../backlog/backlog.md#L460-L460)

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

[Source lines 461-461](../../../backlog/backlog.md#L461-L461)

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

[Source lines 462-462](../../../backlog/backlog.md#L462-L462)

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

[Source lines 463-466](../../../backlog/backlog.md#L463-L466)

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

[Source lines 467-467](../../../backlog/backlog.md#L467-L467)

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

[Source lines 468-468](../../../backlog/backlog.md#L468-L468)

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

[Source lines 469-469](../../../backlog/backlog.md#L469-L469)

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

[Source lines 470-470](../../../backlog/backlog.md#L470-L470)

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

[Source lines 471-471](../../../backlog/backlog.md#L471-L471)

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

[Source lines 472-472](../../../backlog/backlog.md#L472-L472)

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

[Source lines 473-473](../../../backlog/backlog.md#L473-L473)

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

[Source lines 474-474](../../../backlog/backlog.md#L474-L474)

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

[Source lines 475-475](../../../backlog/backlog.md#L475-L475)

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

[Source lines 476-477](../../../backlog/backlog.md#L476-L477)

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

[Source lines 478-478](../../../backlog/backlog.md#L478-L478)

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

[Source lines 479-481](../../../backlog/backlog.md#L479-L481)

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

## CAND-140

### Source

> - [x] Consistently use "Sign in" and "Sign up" in the sign in flow.

[Source lines 482-482](../../../backlog/backlog.md#L482-L482)

### Candidate behavior

The sign-in flow consistently displays `Sign in` and `Sign up`.

### Applicability

Actor: user. Location: sign-in flow. Event kind: any. Interaction mode: sign-in or sign-up. Viewport: any. State: sign-in enabled. Exclusions: disabled sign-in.

### Classification

candidate FR

### Existing Requirements and Confidence

No accepted FR or QR governs enabled sign-in-flow terminology. Confidence: inferred.

### Disposition

Hold for consolidation as a terminology refinement.

### Open Questions

Does “sign in flow” include routes, headings, buttons, and messages?

## CAND-141

### Source

> - [x] Don't provide support text in the Sign up form because such text is redundant.

[Source lines 483-483](../../../backlog/backlog.md#L483-L483)

### Candidate behavior

The sign-up form omits redundant support text.

### Applicability

Actor: user. Location: sign-up form. Event kind: any. Interaction mode: viewing the form. Viewport: any. State: sign-in enabled. Exclusions: unspecified support text that is not redundant.

### Classification

needs product decision

### Existing Requirements and Confidence

No accepted FR or QR identifies the support text. Confidence: needs product decision.

### Disposition

Do not consolidate until the text and redundancy criterion are defined.

### Open Questions

Which support text is in scope, and who determines redundancy?

## CAND-142

### Source

> - [x] Use TypeScript 7 to speed up compilation

[Source lines 484-484](../../../backlog/backlog.md#L484-L484)

### Candidate behavior

No new requirement behavior asserted; this names a compiler version as a means to improve compilation speed.

### Applicability

Actor: developer. Location: frontend tooling. Event kind: any. Interaction mode: compilation. Viewport: not applicable. State: development. Exclusions: runtime behavior.

### Classification

implementation detail

### Existing Requirements and Confidence

No accepted FR or QR requires a TypeScript version. Confidence: confirmed.

### Disposition

Exclude from requirements consolidation.

## CAND-143

### Source

> - [x] Use oxlint to speed up linting

[Source lines 485-485](../../../backlog/backlog.md#L485-L485)

### Candidate behavior

No new requirement behavior asserted; this selects a linting tool to improve developer workflow speed.

### Applicability

Actor: developer. Location: frontend tooling. Event kind: any. Interaction mode: linting. Viewport: not applicable. State: development. Exclusions: product behavior.

### Classification

implementation detail

### Existing Requirements and Confidence

No accepted FR or QR requires a linting tool. Confidence: confirmed.

### Disposition

Exclude from requirements consolidation.

## CAND-144

### Source

> - [x] Update references from deemp/timeful to whensync/timeful
>   - Reserved whensync just in case I'd like to rename the project in future

[Source lines 486-487](../../../backlog/backlog.md#L486-L487)

### Candidate behavior

No new requirement behavior asserted; this updates repository references and records a possible future naming decision.

### Applicability

Actor: maintainer. Location: repository references. Event kind: any. Interaction mode: maintenance. Viewport: not applicable. State: current project name. Exclusions: a future product rename.

### Classification

ADR or decision

### Existing Requirements and Confidence

No accepted FR or QR governs repository naming. Confidence: inferred.

### Disposition

Retain only as project-history context; do not consolidate into a requirement.

### Open Questions

Was a product rename decided, or was only the repository namespace changed?

## CAND-145

### Source

> - [x] Fix event description card
>   - Multi-line event descriptions save with newline characters preserved.
>   - Saved multi-line descriptions render across multiple lines.
>   - Read-only description pencil is in the upper-right corner.
>   - Edit mode preserves the description’s text width and wrapping.
>   - Edit mode cancel and save buttons are in the upper-right corner.
>   - Description action buttons are reduced to 32px and visually centered for single-line descriptions.

[Source lines 488-494](../../../backlog/backlog.md#L488-L494)

### Candidate behavior

No new requirement behavior asserted for newline preservation because FR-019 already requires it; the remaining card layout details are an unconfirmed refinement.

### Applicability

Actor: event owner or reader. Location: event description card. Event kind: any. Interaction mode: read-only and edit. Viewport: any. State: multiline or single-line description. Exclusions: description persistence and rendering already covered by FR-019.

### Classification

duplicate or refinement

### Existing Requirements and Confidence

FR-019 accepts preservation and multiline rendering of event descriptions; it does not specify controls, dimensions, or placement. Confidence: confirmed.

### Disposition

Retain the layout details as a possible refinement separate from FR-019.

### Open Questions

Which actors can edit, and are the 32px and placement details durable product intent?

## CAND-146

### Source

> - [x] Hovering a disabled date on a dates-only event now gives respondents a disabled status, not active/unavailable.

[Source lines 495-495](../../../backlog/backlog.md#L495-L495)

### Candidate behavior

Hovering a disabled date on a dates-only event shows respondents a disabled status rather than active or unavailable.

### Applicability

Actor: respondent. Location: dates-only event page. Event kind: dates-only. Interaction mode: pointer hover. Viewport: pointer-capable. State: disabled date. Exclusions: enabled dates and touch-only interaction.

### Classification

candidate FR

### Existing Requirements and Confidence

FR-009 defines timed-grid state presentation and does not confirm dates-only hover status. Confidence: inferred.

### Disposition

Hold as a dates-only state-presentation candidate.

### Open Questions

What exact status text and response-view treatment are required?

## CAND-147

### Source

> - [x] The dates-only legend now shows Unavailable, change in Add/Edit availability, matching timed events.

[Source lines 496-496](../../../backlog/backlog.md#L496-L496)

### Candidate behavior

The dates-only legend shows `Unavailable, change in Add/Edit availability`.

### Applicability

Actor: respondent. Location: dates-only event legend. Event kind: dates-only. Interaction mode: viewing availability. Viewport: any. State: unavailable date. Exclusions: timed-event legend behavior already addressed by FR-009.

### Classification

candidate FR

### Existing Requirements and Confidence

FR-009 accepts the corresponding timed-grid label but does not state dates-only coverage. Confidence: inferred.

### Disposition

Hold as a dates-only extension of FR-009.

### Open Questions

Does the label apply in every dates-only availability mode?

## CAND-148

### Source

> - [x] Fix highlighting in dates-only events
>   - Enabled dates-only grid cells receive a black inset square frame on hover.
>   - The frame stays inside the cell and does not overlap neighboring cells or the date number.
>   - The date text does not change from border-box layout shifts.
>   - Hovering a disabled cell removes the prior enabled-cell frame.
>   - Hovering a disabled cell shows disabled status in Responses.
>   - Leaving the grid removes the current frame.
>   - On mobile, tapping a disabled cell removes the current frame while showing disabled Responses status.
>   - On mobile, tapping outside the grid removes the current frame and clears its state.

[Source lines 497-505](../../../backlog/backlog.md#L497-L505)

### Candidate behavior

Hovering an enabled dates-only grid cell displays an inset frame that remains inside that cell without shifting its date text.

### Applicability

Actor: respondent. Location: dates-only grid and Responses. Event kind: dates-only. Interaction mode: hover or mobile tap. Viewport: desktop and mobile. State: enabled or disabled cell, including leaving the grid. Exclusions: timed-grid highlights.

### Classification

candidate FR

### Existing Requirements and Confidence

FR-042 proposes analogous timed-grid containment and is not accepted; no accepted FR covers dates-only highlights. Confidence: inferred.

### Disposition

Hold as a compound interaction candidate; split visual containment from response-state behavior during consolidation.

### Open Questions

Are the black color and square frame durable, and what does “clears its state” include?

## CAND-149

### Source

> - [x] Fix gap on the right for dates-only events.
>   - Dates-only grid fits within narrow viewports from 320px through 639px.
>   - Grid has a 16px left and right gutter on phone layouts.
>   - No horizontal document overflow at 320, 390, 410, 480, 639, and 640px.
>   - Below 640px, the sidebar remains stacked beneath the calendar and fits the viewport.
>   - At 640px, grid and sidebar remain side-by-side with their intended 16–20px gap.
>   - The full-width dates-only grid no longer adds external margin beyond its pane.

[Source lines 506-512](../../../backlog/backlog.md#L506-L512)

### Candidate behavior

On 320px through 639px viewports, a dates-only grid and its stacked sidebar fit without horizontal document overflow.

### Applicability

Actor: user. Location: dates-only event page. Event kind: dates-only. Interaction mode: viewing. Viewport: 320px through 640px. State: responsive layout. Exclusions: wider layouts except the stated 640px boundary.

### Classification

candidate FR

### Existing Requirements and Confidence

No accepted FR or QR specifies dates-only responsive layout. Confidence: inferred.

### Disposition

Hold for consolidation; separate fit behavior from pixel-gutter implementation details.

### Open Questions

Is the 640px side-by-side boundary and 16–20px gap a durable product decision?

## CAND-150

### Source

> - [x] On the dates-only event page, the date cells must be rectangular 1:2 (height:width) to better fit the screen

[Source lines 513-513](../../../backlog/backlog.md#L513-L513)

### Candidate behavior

Dates-only event-page date cells have a 1:2 height-to-width ratio.

### Applicability

Actor: user. Location: dates-only event page. Event kind: dates-only. Interaction mode: viewing. Viewport: unspecified. State: calendar visible. Exclusions: timed grids.

### Classification

candidate FR

### Existing Requirements and Confidence

No accepted FR or QR specifies dates-only cell geometry. Confidence: inferred.

### Disposition

Hold as a visual-layout candidate.

### Open Questions

Does the ratio apply at every viewport and to all calendar states?

## CAND-151

### Source

> - [x] Improve separation between environments (development, test, staging, production)

[Source lines 514-514](../../../backlog/backlog.md#L514-L514)

### Candidate behavior

No new requirement behavior asserted; “improve separation” does not specify an observable environment boundary.

### Applicability

Actor: operator or developer. Location: deployment environments. Event kind: any. Interaction mode: development and deployment. Viewport: not applicable. State: development, test, staging, or production. Exclusions: unspecified.

### Classification

needs product decision

### Existing Requirements and Confidence

No accepted FR or QR defines environment separation. Confidence: needs product decision.

### Disposition

Do not consolidate without defined isolation outcomes.

### Open Questions

Which resources must be separated: data, credentials, hosts, services, or configuration?

## CAND-152

### Source

> - [x] Don't run vite in a container to keep things simple and fast during development

[Source lines 515-515](../../../backlog/backlog.md#L515-L515)

### Candidate behavior

No new requirement behavior asserted; this is a development-environment implementation choice.

### Applicability

Actor: developer. Location: local frontend development. Event kind: any. Interaction mode: running Vite. Viewport: not applicable. State: development. Exclusions: production deployment.

### Classification

ADR or decision

### Existing Requirements and Confidence

No accepted FR or QR mandates the Vite execution environment. Confidence: confirmed.

### Disposition

Exclude from requirements consolidation.

## CAND-153

### Source

> - [x] <http://127.0.0.1:4173/e/aB3BE>
>   - ok in gmt+3
>   - padding gets added in +3:30 although not needed
>   - grid split in +12 although shouldn't
>
>   Status: can't reproduce

[Source lines 516-521](../../../backlog/backlog.md#L516-L521)

### Candidate behavior

No new requirement behavior asserted; this is an unreproduced timezone-specific investigation report.

### Applicability

Actor: user. Location: a specific event page. Event kind: unspecified. Interaction mode: viewing after timezone selection. Viewport: unspecified. State: GMT+3, +3:30, or +12. Exclusions: reproducible behavior not established.

### Classification

bug or investigation

### Existing Requirements and Confidence

FR-013 accepts preservation of timed-slot instants across display-timezone changes, but does not confirm the reported padding or split. Confidence: inferred.

### Disposition

Exclude from requirements unless reproduced and recast as a durable outcome.

### Open Questions

What event configuration, display timezone, and viewport reproduce the report?

## CAND-154

### Source

> - [x] Add `flake.nix`

[Source lines 522-522](../../../backlog/backlog.md#L522-L522)

### Candidate behavior

No new requirement behavior asserted; this identifies a development-environment file.

### Applicability

Actor: developer. Location: repository root. Event kind: any. Interaction mode: environment setup. Viewport: not applicable. State: development. Exclusions: product runtime behavior.

### Classification

implementation detail

### Existing Requirements and Confidence

No accepted FR or QR requires Nix tooling. Confidence: confirmed.

### Disposition

Exclude from requirements consolidation.

## CAND-155

### Source

> - [x] rename PR (<https://github.com/schej-it/timeful.app/pull/250>) to `Modernize the app`

[Source lines 523-523](../../../backlog/backlog.md#L523-L523)

### Candidate behavior

No new requirement behavior asserted; this changes review metadata.

### Applicability

Actor: maintainer. Location: pull request metadata. Event kind: any. Interaction mode: code review. Viewport: not applicable. State: pull request. Exclusions: product behavior.

### Classification

implementation detail

### Existing Requirements and Confidence

No accepted FR or QR governs pull request titles. Confidence: confirmed.

### Disposition

Exclude from requirements consolidation.

## CAND-156

### Source

> - [x] Use Node 26.5.0 everywhere

[Source lines 524-524](../../../backlog/backlog.md#L524-L524)

### Candidate behavior

No new requirement behavior asserted; this pins a runtime version for project tooling.

### Applicability

Actor: developer or operator. Location: project runtime tooling. Event kind: any. Interaction mode: build, test, or deployment. Viewport: not applicable. State: all environments. Exclusions: user-visible runtime behavior.

### Classification

implementation detail

### Existing Requirements and Confidence

No accepted FR or QR requires a Node version. Confidence: confirmed.

### Disposition

Exclude from requirements consolidation.

## CAND-157

### Source

> - [x] On the specific times page <http://127.0.0.1:4173/e/Eb67A>, when I switch timezone from +5 to +6, the left-most upper-most enabled slot should become disabled

[Source lines 525-525](../../../backlog/backlog.md#L525-L525)

### Candidate behavior

After changing the timezone from +5 to +6 on the stated specific-times page, the left-most upper-most enabled slot becomes disabled.

### Applicability

Actor: user. Location: specific-times page. Event kind: timed. Interaction mode: changing timezone. Viewport: unspecified. State: transition from +5 to +6. Exclusions: other transitions and event configurations.

### Classification

duplicate or refinement

### Existing Requirements and Confidence

FR-013 accepts preservation of timed-slot instants across display-timezone changes; this is a specific expected projection. Confidence: inferred.

### Disposition

Treat as a regression example for FR-013 rather than a new requirement.

### Open Questions

Is the changed timezone the display timezone, and what initial slot domain produces this result?

## CAND-158

### Source

> - [x] On the specific times page <http://127.0.0.1:4173/e/Eb67A>, when I switch timezone from +5 to +4 and on june 14, 0-4 are selected, jun 13 should appear

[Source lines 526-526](../../../backlog/backlog.md#L526-L526)

### Candidate behavior

After changing the timezone from +5 to +4 with June 14 0-4 selected, June 13 appears on the stated specific-times page.

### Applicability

Actor: user. Location: specific-times page. Event kind: timed. Interaction mode: changing timezone. Viewport: unspecified. State: +5 to +4 with the stated selection. Exclusions: other transitions and selections.

### Classification

duplicate or refinement

### Existing Requirements and Confidence

FR-013 accepts preservation of timed-slot instants across display-timezone changes; this is a specific expected projection. Confidence: inferred.

### Disposition

Treat as a regression example for FR-013 rather than a new requirement.

### Open Questions

Is “June 13 should appear” a date column, a grid row, or an enabled slot?

## CAND-159

### Source

> - [x] In the new timed event form, Advanced options must show the Time increment.

[Source lines 527-527](../../../backlog/backlog.md#L527-L527)

### Candidate behavior

Advanced options in the new timed-event form show the time increment.

### Applicability

Actor: event owner. Location: new timed-event form. Event kind: timed. Interaction mode: creating an event. Viewport: any. State: advanced options visible. Exclusions: editing existing events.

### Classification

candidate FR

### Existing Requirements and Confidence

FR-024 governs time-format controls, not time-increment visibility. Confidence: inferred.

### Disposition

Hold as a form-content candidate.

### Open Questions

Is the control editable there, and is it visible only in a particular domain mode?

## CAND-160

### Source

> - [x] On specific times edit page, show the time format and timezone menu between the instructions (Click and drag ...) and the Legend

[Source lines 528-528](../../../backlog/backlog.md#L528-L528)

### Candidate behavior

The specific-times edit page shows the time-format and timezone menu between its instructions and Legend.

### Applicability

Actor: event owner. Location: specific-times edit page. Event kind: timed. Interaction mode: editing. Viewport: unspecified. State: instructions and Legend visible. Exclusions: non-editing pages.

### Classification

candidate FR

### Existing Requirements and Confidence

FR-024 governs independent event and display time formats; it does not require this menu placement. Confidence: inferred.

### Disposition

Hold as a layout refinement candidate.

### Open Questions

Which timezone control is intended, and does placement vary by viewport?

## CAND-161

### Source

> - [x] On specific times edit page, align top edge of the date (month + day) with top edge of the instruction "Click and drag ..."

[Source lines 529-529](../../../backlog/backlog.md#L529-L529)

### Candidate behavior

On the specific-times edit page, the date top edge aligns with the `Click and drag ...` instruction top edge.

### Applicability

Actor: event owner. Location: specific-times edit page. Event kind: timed. Interaction mode: editing. Viewport: unspecified. State: page header visible. Exclusions: other pages.

### Classification

candidate FR

### Existing Requirements and Confidence

No accepted FR or QR specifies this alignment. Confidence: inferred.

### Disposition

Hold as a visual-layout candidate.

### Open Questions

Does the alignment apply across responsive breakpoints?

## CAND-162

### Source

> - [x] "Legend" font size must be the same as "Responses" font size

[Source lines 530-530](../../../backlog/backlog.md#L530-L530)

### Candidate behavior

The `Legend` and `Responses` labels have the same font size.

### Applicability

Actor: user. Location: event-page sidebar. Event kind: unspecified. Interaction mode: viewing. Viewport: unspecified. State: both labels visible. Exclusions: pages without both labels.

### Classification

candidate FR

### Existing Requirements and Confidence

No accepted FR or QR specifies relative sidebar typography. Confidence: inferred.

### Disposition

Hold as a visual-consistency candidate.

### Open Questions

Does this apply to dates-only and timed pages at all breakpoints?

## CAND-163

### Source

> - [x] Use "Legend" instead of "Legend:"

[Source lines 531-531](../../../backlog/backlog.md#L531-L531)

### Candidate behavior

The label displays `Legend` rather than `Legend:`.

### Applicability

Actor: user. Location: event-page sidebar. Event kind: unspecified. Interaction mode: viewing. Viewport: any. State: legend visible. Exclusions: other labels.

### Classification

candidate FR

### Existing Requirements and Confidence

No accepted FR or QR specifies the Legend label. Confidence: inferred.

### Disposition

Hold as a terminology refinement candidate.

### Open Questions

Does this change apply to every legend, including non-event pages?

## CAND-164

### Source

> - [x] Align upper edge of the sidebar with the upper edge of the grid

[Source lines 532-532](../../../backlog/backlog.md#L532-L532)

### Candidate behavior

The sidebar upper edge aligns with the grid upper edge.

### Applicability

Actor: user. Location: event page. Event kind: unspecified. Interaction mode: viewing. Viewport: side-by-side layout. State: sidebar and grid visible. Exclusions: stacked layouts.

### Classification

candidate FR

### Existing Requirements and Confidence

No accepted FR or QR specifies sidebar-to-grid alignment. Confidence: inferred.

### Disposition

Hold as a visual-layout candidate.

### Open Questions

Which event pages and breakpoint define the side-by-side layout?

## CAND-165

### Source

> - [x] The right edge of the navigation and header (buttons in the upper-right corner) and the sidebar must coincide

[Source lines 533-533](../../../backlog/backlog.md#L533-L533)

### Candidate behavior

The navigation/header right edge and sidebar right edge coincide.

### Applicability

Actor: user. Location: event page navigation, header, and sidebar. Event kind: unspecified. Interaction mode: viewing. Viewport: unspecified. State: sidebar visible. Exclusions: layouts without a sidebar.

### Classification

candidate FR

### Existing Requirements and Confidence

No accepted FR or QR specifies these shared right edges. Confidence: inferred.

### Disposition

Hold as a visual-layout candidate.

### Open Questions

Does this apply at mobile breakpoints where the sidebar is stacked?

## CAND-166

### Source

> - [x] Reduce the gap between the grid and Responses

[Source lines 534-534](../../../backlog/backlog.md#L534-L534)

### Candidate behavior

No new requirement behavior asserted; the requested reduction gives no measurable target or baseline.

### Applicability

Actor: user. Location: grid and Responses area. Event kind: unspecified. Interaction mode: viewing. Viewport: unspecified. State: both areas visible. Exclusions: unspecified.

### Classification

needs product decision

### Existing Requirements and Confidence

No accepted FR or QR specifies this gap. Confidence: needs product decision.

### Disposition

Do not consolidate until the intended spacing or observable layout outcome is defined.

### Open Questions

What target gap, viewport range, and event pages are intended?

## CAND-167

### Source

> - [x] Make compose files fully configurable via required values from corresponding .env files

[Source lines 535-535](../../../backlog/backlog.md#L535-L535)

### Candidate behavior

Compose configuration obtains required values from its corresponding `.env` files.

### Applicability

Actor: operator or developer. Location: Compose configuration. Event kind: any. Interaction mode: starting services. Viewport: not applicable. State: configured environment. Exclusions: unspecified optional values.

### Classification

candidate QR

### Existing Requirements and Confidence

QR-012 requires staging and production to reject missing or unsafe required configuration, but does not require Compose or `.env` files. Confidence: inferred.

### Disposition

Hold as an installability/configuration candidate; consolidate only if the required environments and failure behavior are defined.

### Open Questions

Which Compose files and environments are in scope, and must missing values prevent startup?

## CAND-168: Postgres Anonymous-Event Storage

### Source

> - [x] Switch to Postgres for new anon events

[Source lines 536-536](../../../backlog/backlog.md#L536)

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

[Source lines 537-537](../../../backlog/backlog.md#L537)

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

[Source lines 538-540](../../../backlog/backlog.md#L538-L540)

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

[Source lines 541-541](../../../backlog/backlog.md#L541)

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

Existing requirements: [FR-011](../functional/fr/FR-011.md) establishes collapsed timed-grid behavior but not this control order.

Confidence: inferred

### Disposition

Consolidate only if control order is a durable product behavior rather than incidental layout.

### Open Questions

- Is the required order intended for every viewport and control state?

## CAND-172: Dates-Only More-Options Order

### Source

> - [x] On the dates-only event page, in More options, "Start on Monday" must be above "Hide if needed days"

[Source lines 542-542](../../../backlog/backlog.md#L542)

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

[Source lines 543-543](../../../backlog/backlog.md#L543)

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

[Source lines 544-544](../../../backlog/backlog.md#L544)

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

Existing requirements: [QR-009](../quality/qr/QR-009.md) concerns timed-event runtime response time, not test execution.

Confidence: confirmed

### Disposition

Retain as implementation provenance; no QR is asserted.

## CAND-175: Range-Event Collapsed Hours

### Source

> - [x] Collapsed hours should work in range events too, not only specific-times events

[Source lines 545-545](../../../backlog/backlog.md#L545)

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

Existing requirements: [FR-011](../functional/fr/FR-011.md) already requires collapsed inactive runs for timed-event viewing and scheduling.

Confidence: confirmed

### Disposition

Map to FR-011; no new FR.

## CAND-176: Dates-Only Edit-Event Entry

### Source

> - [x] On dates-only event page, when I click the "Edit event" button, a form for editing the event opens

[Source lines 546-546](../../../backlog/backlog.md#L546)

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

Existing requirements: [FR-018](../functional/fr/FR-018.md) covers authorization to edit event settings, not this entry point.

Confidence: inferred

### Disposition

Consolidate with event-settings navigation if that navigation is durable.

### Open Questions

- Is this entry point owner-only, and what should non-owners see?

## CAND-177: Active-Slot Persistence Model

### Source

> - [x] Store only active timeslots and calculate all other types (enabled inactive and disabled) on the fly

[Source lines 547-547](../../../backlog/backlog.md#L547)

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

Existing requirements: [FR-009](../functional/fr/FR-009.md) specifies visible grid states; [FR-074](../functional/fr/FR-074.md) is proposed and specifies enabled-domain derivation.

Confidence: confirmed

### Disposition

Retain as design provenance; do not create an FR or QR from storage wording.

## CAND-178: Comparator E2E Test Import

### Source

> - [x] add Playwright e2e tests from the comparator to the main repository

[Source lines 548-548](../../../backlog/backlog.md#L548)

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

[Source lines 549-550](../../../backlog/backlog.md#L549-L550)

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

[Source lines 551-551](../../../backlog/backlog.md#L551)

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

[Source lines 552-554](../../../backlog/backlog.md#L552-L554)

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

[Source lines 555-555](../../../backlog/backlog.md#L555)

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

[Source lines 556-558](../../../backlog/backlog.md#L556-L558)

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

Existing requirements: [FR-065](../functional/fr/FR-065.md) is proposed and concerns availability-editing response lists, not action visibility.

Confidence: inferred

### Disposition

Consolidate only after confirming whether “responses” means all event responses or editable responses.

### Open Questions

- Which response count controls the action for a visitor without edit authorization?

## CAND-184: Edit-Event Button Placement

### Source

> - [x] On the event page, Edit event button shall be under the event title

[Source lines 559-559](../../../backlog/backlog.md#L559)

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

Existing requirements: [FR-018](../functional/fr/FR-018.md) covers edit authorization, not button placement.

Confidence: inferred

### Disposition

Consolidate only if placement is durable across responsive layouts.

### Open Questions

- Does “under” apply to both desktop and mobile headers?

## CAND-185: GitHub Repository Navbar Placement

### Source

> - [x] On the event page, Github repo button shall be in the navbar, right-most button

[Source lines 560-560](../../../backlog/backlog.md#L560)

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

[Source lines 561-561](../../../backlog/backlog.md#L561)

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

Existing requirements: [QR-008](../quality/qr/QR-008.md) is proposed and broadly concerns accessible coordination flows; no accepted QR covers README icon readability.

Confidence: inferred

### Disposition

Consolidate as documentation accessibility guidance only if a contrast criterion and rendering targets are confirmed.

### Open Questions

- Which README renderers and contrast threshold define “readable”?

## CAND-187: Collapsed-Strip Hover Treatment

### Source

> - [x] On the event page, when I hover over collapsed hours strip, neither the strip nor any cell is highlighted like active cells are highlighted on hover.

[Source lines 562-562](../../../backlog/backlog.md#L562)

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

Existing requirements: [FR-011](../functional/fr/FR-011.md) establishes collapsed timed-grid runs but does not specify hover treatment.

Confidence: inferred

### Disposition

Consolidate as a grid interaction refinement if the treatment is intentionally observable.

### Open Questions

- Should touch interaction have an equivalent non-highlight rule?

## CAND-188: Scrolling No-Responses Position

### Source

> - [x] Bug: when scrolling the gri, the No responses yet changes the position
>   - Can't reproduce

[Source lines 563-564](../../../backlog/backlog.md#L563-L564)

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

[Source lines 565-565](../../../backlog/backlog.md#L565)

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

Existing requirements: None identified; [FR-046](../functional/fr/FR-046.md) and [FR-047](../functional/fr/FR-047.md) are proposed labels, not ordering.

Confidence: inferred

### Disposition

Consolidate with timed-page sidebar layout if the desktop breakpoint is confirmed.

### Open Questions

- What viewport threshold defines desktop?

## CAND-190: Specific-Times Date Regression Report

### Source

> - [x] Create an event with specific times for dates Aug 30, 31, mark hours 0-4 for both dates, edit event, set dates for 28, 29, click next. See May 28, 30, 31 in specific times page, and May 30, 31 on the event page.
>   - Can't reproduce. I see Aug 28, Aug 29 on the event page

[Source lines 566-567](../../../backlog/backlog.md#L566-L567)

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

Existing requirements: [FR-050](../functional/fr/FR-050.md), [FR-051](../functional/fr/FR-051.md), and [FR-052](../functional/fr/FR-052.md) are proposed date-domain rules; none confirms this report.

Confidence: confirmed

### Disposition

Retain as investigation history; no FR or QR.

### Open Questions

- Which timezone, date picker state, and saved data are required to reproduce the reported May dates?

## CAND-191: Specific-Times Missing-Date Regression Report

### Source

> - [x] Create an event with two dates, mark timeslots for only one day in specific times, save, edit again and see only one day on the event page
>   - Can't reproduce. I see both days.

[Source lines 568-569](../../../backlog/backlog.md#L568-L569)

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

Existing requirements: [FR-074](../functional/fr/FR-074.md) is proposed and derives the enabled domain for each picked date; it does not confirm the report.

Confidence: confirmed

### Disposition

Retain as investigation history; no FR or QR.

### Open Questions

- What persisted data or transition would cause one selected date to disappear?

## CAND-192: Mobile Timezone-Control Width

### Source

> - [x] On mobile, on the timed event page, the timezone control keeps a fixed width (112px) so the reset button fits inside without resizing the control, with the time format and days-per-page buttons on either side

[Source lines 570-570](../../../backlog/backlog.md#L570)

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

Existing requirements: [FR-049](../functional/fr/FR-049.md) is proposed and places the mobile controls in a row but does not specify width or order.

Confidence: inferred

### Disposition

Consolidate only if the 112px measurement is a durable compatibility constraint.

### Open Questions

- What mobile width range and localization widths must the fixed control support?

## CAND-193: OTP Send-Failure Report

### Source

> - [x] On the Create your account form, when I click Continue and the OTP can't be sent, a report about that is visible

[Source lines 571-571](../../../backlog/backlog.md#L571)

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

Existing requirements: [FR-023](../functional/fr/FR-023.md) is accepted and covers registration status during sign-in, but not an OTP send failure.

Confidence: inferred

### Disposition

Consolidate with registration feedback behavior if OTP is confirmed as the current registration mechanism.

### Open Questions

- What failures receive a report, and what information may it disclose?

## CAND-194: Timezone-Reset Icon Direction

### Source

> - [x] The icon to reset the timezone shall have a counter-clockwise arrow

[Source lines 572-572](../../../backlog/backlog.md#L572)

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

[Source lines 573-573](../../../backlog/backlog.md#L573)

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

Existing requirements: [FR-046](../functional/fr/FR-046.md) and [FR-047](../functional/fr/FR-047.md) are proposed label requirements, not row layout. Overlap: CAND-196 specifies the separate time-control order; no accepted FR/QR overlap.

Confidence: inferred

### Disposition

Consolidate only after resolving the source's potentially ambiguous alignment wording.

### Open Questions

- Does “time format aligned right” mean its text, button, or allocated area?

## CAND-196: Timed Sidebar Time-Control Order

### Source

> - [x] On the timed event page, the time format shall be on the left from the time zone so that the timezone button can get a reset button or stay the same freely

[Source lines 574-574](../../../backlog/backlog.md#L574)

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

Existing requirements: [FR-046](../functional/fr/FR-046.md) and [FR-047](../functional/fr/FR-047.md) are proposed label requirements, not control order. Overlap: CAND-195 specifies the separate shared-row and alignment behavior; no accepted FR/QR overlap.

Confidence: inferred

### Disposition

Consolidate with timed-page sidebar layout if the order is durable across responsive layouts.

### Open Questions

- Does this order apply when the timezone reset control is absent?

## CAND-197: Mobile Tooltip Screen Containment

### Source

> - [x] On mobile, on the event page, tooltip shall be fully on-screen

[Source lines 575-575](../../../backlog/backlog.md#L575)

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

Existing requirements: [FR-020](../functional/fr/FR-020.md) already requires a visible mobile selected-slot tooltip adjacent to a visible selected slot; it does not expressly require screen containment.

Confidence: inferred

### Disposition

Refine FR-020 if containment is confirmed; do not create a separate requirement.

### Open Questions

- Does this apply to every tooltip or only selected-slot tooltips?

## CAND-198: Mobile Availability Show-All-Hours Placement

### Source

> - [x] On mobile, when adding availability, Show all hours should be an option at its normal place - under other options between the event description and the grid

[Source lines 576-576](../../../backlog/backlog.md#L576)

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

Existing requirements: [FR-011](../functional/fr/FR-011.md) covers the expansion behavior; [FR-048](../functional/fr/FR-048.md) is proposed and covers centering in a no-response state, not this placement.

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

[Source lines 577-585](../../../backlog/backlog.md#L577-L585)

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

Existing requirements: [FR-028](../functional/fr/FR-028.md) requires availability controls and excludes response details from the mobile selected-slot panel while editing availability.

Confidence: confirmed

### Disposition

Map to FR-028 as a regression example; no new FR.
