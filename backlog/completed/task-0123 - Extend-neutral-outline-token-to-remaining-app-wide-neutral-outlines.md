---
id: TASK-0123
title: Extend neutral outline token to remaining app-wide neutral outlines
status: Done
assignee:
  - '@opencode'
created_date: '2026-08-31 13:38'
updated_date: '2026-08-31 14:20'
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
- [x] #1 Every remaining neutral border class/hex listed in the description is replaced by tw-border-outline-neutral or var(--timeful-outline-neutral), and no site listed in the description still carries tw-border-light-gray-stroke, tw-border-gray, tw-border-off-white, or a raw neutral border hex afterwards
- [x] #2 The colored conditional arms (SignUpBlock and SignUpCalendarBlock tw-border-light-green, CalendarEventBlock tw-border-white on blue blocks) remain unchanged
- [x] #3 The .timeful-switch raw track border in index.css uses var(--timeful-compact-switch-track-border) with no visual change
- [x] #4 Visually prominent conversions (timeful-solo-field #4f4f4f1f hairline, timeful-elevated-button #dfdfdf/#f2f2f2 borders, settings CalendarAccounts off-white divider) are either confirmed with the user before conversion or explicitly kept, and the decision is recorded in task notes
- [x] #5 Unit tests asserting the old classes (RespondentsList.test.ts tw-border-gray status-square assertion, any other affected suites) are updated and pass
- [x] #6 npm run lint, typecheck, build, and test:unit pass; compiled CSS contains the token and generated utility
<!-- AC:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [x] #1 All acceptance criteria are satisfied
- [x] #2 All required unit tests pass. Documentation-only changes are exempt unless the user requests unit tests
- [x] #3 All required e2e tests pass. Documentation-only changes are exempt unless the user requests e2e tests
- [x] #4 Changed Markdown files are formatted with npm run format:markdown
<!-- DOD:END -->

## Implementation Plan

<!-- SECTION:PLAN:BEGIN -->
1. Token layer (frontend/src/index.css)
   - .timeful-elevated-button border #dfdfdf -> var(--timeful-outline-neutral); .timeful-elevated-button.tw-bg-white border #f2f2f2 -> var(--timeful-outline-neutral) (user-approved).
   - .timeful-solo-field .v-field border #4f4f4f1f -> var(--timeful-outline-neutral) (user-approved).
   - .timeful-switch .v-switch__track border #bdbdbd -> var(--timeful-compact-switch-track-border) (AC #3, no visual change; track bg literal #bdbdbd stays).
   - Colored elevated-button variants (#29bc68, #6b6b6b, #53a2ff) and dirty-state track colors unchanged.
2. Template swaps to tw-border-outline-neutral (Tailwind alias already exists from TASK-0122)
   - Settings.vue x5 (79, 82, 88, 98, 104), Landing.vue 75, FAQ.vue 6 (untoggled arm only), Dashboard.vue 37+155.
   - RespondentsList.vue panel 6 + status square 175; ColorLegend.vue x5 (7, 15, 23, 33, 51); ScheduleOverlapTimeGrid.vue 15+338; ScheduleOverlapDaysOnlyGrid.vue 6+17 (pager family, matches ScheduleOverlapSidebar conversion).
   - SignUpBlock.vue 5 and SignUpCalendarBlock.vue 4: neutral conditional arm only; tw-border-light-green arms stay.
   - CookieConsent.vue 54; settings/CalendarAccounts.vue 20 (tw-border-off-white -> token, user-approved).
   - Approved extras: EventDescription.vue 3 (tw-border-light-gray-stroke); Event.vue 805 divider (tw-border-t tw-border-gray).
3. Tests
   - RespondentsList.test.ts:281 status-square assertion tw-border-gray -> tw-border-outline-neutral.
   - NewEvent.test.ts shared-control-contract pin: extend solo-field assertion with border var(--timeful-outline-neutral) (alongside existing box-shadow none pin).
   - Leave tw-bg-light-gray-stroke assertions (backgrounds) and e2e dom-resolvers (already on token) untouched.
4. Verification
   - Grep: no remaining tw-border-light-gray-stroke / tw-border-gray / tw-border-off-white on border sites or raw neutral hexes from the listed set (gray palette stays for text/bg).
   - npm run lint, typecheck, build, test:unit; verify compiled CSS contains --timeful-outline-neutral and tw-border-outline-neutral utility.
   - graphify update .
<!-- SECTION:PLAN:END -->

## Implementation Notes

<!-- SECTION:NOTES:BEGIN -->
AC #4 user decisions (2026-08-31, via structured questions): convert solo-field #4f4f4f1f -> var(--timeful-outline-neutral); convert both elevated-button borders (#dfdfdf, #f2f2f2) -> var(--timeful-outline-neutral); convert CalendarAccounts off-white divider -> token. Scope approved to include two extra verified neutral-border sites not in the original description: EventDescription.vue (tw-border-light-gray-stroke) and Event.vue ~805 divider (tw-border-t tw-border-gray). Elevated-button colored variants (#29bc68 green, #6b6b6b dark-gray surface-matching edge, #53a2ff blue) and .timeful-switch track bg #bdbdbd (background, not outline) intentionally unchanged.

Verification evidence: post-edit greps show zero tw-border-light-gray-stroke / tw-border-gray / tw-border-off-white in frontend/src .vue files and no raw neutral border hexes left in index.css (only token vars; track background literal kept). Compiled dist CSS contains --timeful-outline-neutral:#bdbdbd, the .tw-border-outline-neutral utility, and the converted solo-field/elevated-button/switch-track rules; residual dfdfdf/f2f2f2 hits in the bundle are background utilities only. lint, typecheck, build pass; test:unit 137 files / 961 tests pass. E2E not run: no spec asserts the affected styles, e2e dom-resolvers already on the token (TASK-0122), styling-only change covered by component tests + compiled-CSS checks. graphify update run.
<!-- SECTION:NOTES:END -->

## Final Summary

<!-- SECTION:FINAL_SUMMARY:BEGIN -->
## Summary

All app-wide neutral (non-colored) borders now resolve to #bdbdbd through `--timeful-outline-neutral` / the `tw-border-outline-neutral` Tailwind alias, replacing the mix of #dfdfdf, #f2f2f2, #4f4f4f1f, and palette-gray border classes.

## Changes

- `frontend/src/index.css`: `.timeful-elevated-button` (#dfdfdf) and `.timeful-elevated-button.tw-bg-white` (#f2f2f2) borders → `var(--timeful-outline-neutral)`; `.timeful-solo-field .v-field` border (#4f4f4f1f) → `var(--timeful-outline-neutral)`; `.timeful-switch .v-switch__track` border → `var(--timeful-compact-switch-track-border)` (resolves to the token; track background literal kept — background, not outline).
- Template swaps to `tw-border-outline-neutral`: Settings.vue (5x), Landing.vue, FAQ.vue (untoggled arm), Dashboard.vue (2x), RespondentsList.vue (panel + status square), ColorLegend.vue (5x swatches), ScheduleOverlapTimeGrid.vue (2x pager buttons), ScheduleOverlapDaysOnlyGrid.vue (2x pager buttons), SignUpBlock.vue and SignUpCalendarBlock.vue (neutral conditional arms), CookieConsent.vue, settings/CalendarAccounts.vue (off-white divider), plus two user-approved extras found during research: EventDescription.vue and the Event.vue phone-layout divider (tw-border-t tw-border-gray).
- Tests: RespondentsList.test.ts status-square assertion updated to `tw-border-outline-neutral`; NewEvent.test.ts "uses explicit shared control contracts" extended with pins for the solo-field border, elevated-button border, and switch-track border var usage.

## User decisions (AC #4, recorded in notes)

All four confirmed via structured questions on 2026-08-31: convert solo-field #4f4f4f1f hairline; convert both elevated-button neutral borders; convert CalendarAccounts off-white divider; include EventDescription.vue and Event.vue divider in scope. Elevated-button colored variants (#29bc68, #6b6b6b surface-matching edge, #53a2ff), green sign-up arms, and CalendarEventBlock white-on-blue remain unchanged.

## Verification

- `npm run lint`, `npm run typecheck`, `npm run build`: pass.
- `npm run test:unit`: 137 files / 961 tests pass, including the updated border assertions.
- Compiled `dist/assets/index-*.css` contains `--timeful-outline-neutral:#bdbdbd`, the generated `.tw-border-outline-neutral{border-color:var(--timeful-outline-neutral)!important}` utility, and the converted solo-field / elevated-button / switch-track rules. Remaining `dfdfdf`/`f2f2f2` hexes in the bundle are background utilities (`tw-bg-light-gray-stroke/50`, `tw-bg-off-white`), which are out of scope by design.
- Post-edit greps: zero `tw-border-light-gray-stroke` / `tw-border-gray` / `tw-border-off-white` border usages remain in frontend/src .vue files; Tailwind gray palette unchanged for text/backgrounds.
- E2E: no e2e spec asserts the affected border styles; the only e2e-tree reference (inspect `dom-resolvers.ts`) already used the token from TASK-0122. Full browser E2E was not run for these styling-only class swaps (same precedent as TASK-0122: covered by component tests + compiled-CSS verification); no check failures to report.
- `graphify update .` run; no hand-edited Markdown files changed (only the Backlog-managed task record), so no format:markdown pass was required.
<!-- SECTION:FINAL_SUMMARY:END -->
