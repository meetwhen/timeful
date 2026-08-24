# Backlog Review Inventory: Dates-Only and Platform

This durable batch artifact is non-normative. `CAND-*` IDs are temporary and
will be consolidated by [../backlog-fr-inventory.md](../backlog-fr-inventory.md).

## CAND-140

### Source

> - [x] Consistently use "Sign in" and "Sign up" in the sign in flow.

[Source lines 482-482](../../../../backlog/backlog.md#L482-L482)

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

[Source lines 483-483](../../../../backlog/backlog.md#L483-L483)

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

[Source lines 484-484](../../../../backlog/backlog.md#L484-L484)

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

[Source lines 485-485](../../../../backlog/backlog.md#L485-L485)

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

[Source lines 486-487](../../../../backlog/backlog.md#L486-L487)

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

[Source lines 488-494](../../../../backlog/backlog.md#L488-L494)

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

[Source lines 495-495](../../../../backlog/backlog.md#L495-L495)

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

[Source lines 496-496](../../../../backlog/backlog.md#L496-L496)

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

[Source lines 497-505](../../../../backlog/backlog.md#L497-L505)

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

[Source lines 506-512](../../../../backlog/backlog.md#L506-L512)

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

[Source lines 513-513](../../../../backlog/backlog.md#L513-L513)

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

[Source lines 514-514](../../../../backlog/backlog.md#L514-L514)

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

[Source lines 515-515](../../../../backlog/backlog.md#L515-L515)

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

[Source lines 516-521](../../../../backlog/backlog.md#L516-L521)

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

[Source lines 522-522](../../../../backlog/backlog.md#L522-L522)

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

[Source lines 523-523](../../../../backlog/backlog.md#L523-L523)

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

[Source lines 524-524](../../../../backlog/backlog.md#L524-L524)

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

[Source lines 525-525](../../../../backlog/backlog.md#L525-L525)

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

[Source lines 526-526](../../../../backlog/backlog.md#L526-L526)

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

[Source lines 527-527](../../../../backlog/backlog.md#L527-L527)

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

[Source lines 528-528](../../../../backlog/backlog.md#L528-L528)

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

[Source lines 529-529](../../../../backlog/backlog.md#L529-L529)

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

[Source lines 530-530](../../../../backlog/backlog.md#L530-L530)

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

[Source lines 531-531](../../../../backlog/backlog.md#L531-L531)

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

[Source lines 532-532](../../../../backlog/backlog.md#L532-L532)

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

[Source lines 533-533](../../../../backlog/backlog.md#L533-L533)

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

[Source lines 534-534](../../../../backlog/backlog.md#L534-L534)

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

[Source lines 535-535](../../../../backlog/backlog.md#L535-L535)

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
