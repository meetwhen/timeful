---
id: TASK-0123
title: Extend neutral outline token to remaining app-wide neutral outlines
status: To Do
assignee: []
created_date: '2026-08-31 13:38'
labels:
  - frontend
  - styling
dependencies: []
priority: medium
type: enhancement
ordinal: 129300
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Follow-up to TASK-0122 (Done): extend the shared neutral outline token `--timeful-outline-neutral` (#bdbdbd, Tailwind alias `outline-neutral` → `tw-border-outline-neutral`) to the remaining neutral outlines app-wide, so all non-colored hairlines resolve to #bdbdbd instead of a mix of #dfdfdf, #f2f2f2, #4f4f4f1f, and palette gray.

Known sites (verify before changing; current values are durable facts):
- Views: Settings.vue (5x tw-border-light-gray-stroke, lines ~79-104), Landing.vue (~75), FAQ.vue (~6 toggled card border).
- Home: Dashboard.vue (~37 and ~155).
- Schedule overlap leftovers: RespondentsList.vue panel (~6) and legend/status swatch (~175, tw-border-gray; RespondentsList.test.ts asserts tw-border-gray on the status square), ColorLegend.vue swatches (5x tw-border-gray), ScheduleOverlapTimeGrid.vue pager buttons (~15, ~338), ScheduleOverlapDaysOnlyGrid.vue pager buttons (~6, ~17) — same pager family already converted in ScheduleOverlapSidebar.
- Sign-up: SignUpBlock.vue (~5) and SignUpCalendarBlock.vue (~4) neutral conditional arm only (colored arm is tw-border-light-green and must stay).
- CookieConsent.vue (~54, tw-border-gray button).
- settings/CalendarAccounts.vue (~20, tw-border-b tw-border-off-white divider).
- index.css: .timeful-elevated-button borders #dfdfdf (~59) and #f2f2f2 (~66); .timeful-solo-field border #4f4f4f1f (~97); .timeful-switch .v-switch__track raw #bdbdbd (~110, switch to var(--timeful-compact-switch-track-border)).

Visually prominent conversions to confirm with the user before applying: the #4f4f4f1f solo-field hairline (much more visible at #bdbdbd, affects all form fields), the #f2f2f2 white elevated-button border, and the off-white divider. Colored borders (green/yellow/blue/red, CalendarEventBlock white-on-blue) stay unchanged. The Tailwind gray palette itself is unchanged (still used for text/backgrounds).
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 Every remaining neutral border class/hex listed in the description is replaced by tw-border-outline-neutral or var(--timeful-outline-neutral), and no site listed in the description still carries tw-border-light-gray-stroke, tw-border-gray, tw-border-off-white, or a raw neutral border hex afterwards
- [ ] #2 The colored conditional arms (SignUpBlock and SignUpCalendarBlock tw-border-light-green, CalendarEventBlock tw-border-white on blue blocks) remain unchanged
- [ ] #3 The .timeful-switch raw track border in index.css uses var(--timeful-compact-switch-track-border) with no visual change
- [ ] #4 Visually prominent conversions (timeful-solo-field #4f4f4f1f hairline, timeful-elevated-button #dfdfdf/#f2f2f2 borders, settings CalendarAccounts off-white divider) are either confirmed with the user before conversion or explicitly kept, and the decision is recorded in task notes
- [ ] #5 Unit tests asserting the old classes (RespondentsList.test.ts tw-border-gray status-square assertion, any other affected suites) are updated and pass
- [ ] #6 npm run lint, typecheck, build, and test:unit pass; compiled CSS contains the token and generated utility
<!-- AC:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [ ] #1 All acceptance criteria are satisfied
- [ ] #2 All required unit tests pass. Documentation-only changes are exempt unless the user requests unit tests
- [ ] #3 All required e2e tests pass. Documentation-only changes are exempt unless the user requests e2e tests
- [ ] #4 Changed Markdown files are formatted with npm run format:markdown
<!-- DOD:END -->
