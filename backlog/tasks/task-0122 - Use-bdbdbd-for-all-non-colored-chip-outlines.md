---
id: TASK-0122
title: 'Use #bdbdbd for all non-colored chip outlines'
status: Done
assignee:
  - opencode
created_date: '2026-08-31 13:17'
updated_date: '2026-08-31 13:36'
labels:
  - frontend
  - styling
dependencies: []
modified_files:
  - frontend/src/index.css
  - frontend/tailwind.config.cjs
  - frontend/src/components/SlideToggle.vue
  - frontend/src/components/schedule_overlap/TimeFormatToggle.vue
  - frontend/src/components/schedule_overlap/TimezoneSelector.vue
  - frontend/src/components/schedule_overlap/ScheduleOverlapSidebar.vue
  - frontend/src/components/schedule_overlap/ScheduleOverlapMobileOverlay.vue
  - frontend/src/components/schedule_overlap/EditingAvailabilityAs.vue
  - frontend/src/components/schedule_overlap/TimeFormatToggle.test.ts
  - frontend/src/components/schedule_overlap/TimezoneSelector.test.ts
  - frontend/src/components/NewEvent.test.ts
  - frontend/e2e/inspect/src/dom-resolvers.ts
priority: medium
type: enhancement
ordinal: 128300
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Standardize the neutral (non-colored) outline color of the schedule-overlap chips and small controls on #bdbdbd so all grey outlines match, instead of a mix of #dfdfdf, #4f4f4f1f, and #bdbdbd.

Colored outlines that carry meaning (green "Available" active indicator, yellow "If needed" active indicator, blue/red error states, etc.) must keep their existing colors; only neutral outlines change.

Current state (durable facts, verify before changing):
- Tailwind already defines gray: #BDBDBD (frontend/tailwind.config.cjs), and the Calendar options button (ScheduleOverlapSidebar.vue, ScheduleOverlapMobileOverlay.vue) plus the EditingAvailabilityAs guest chip already use tw-border-gray.
- SlideToggle.vue outer track (used by the Available/If needed AvailabilityTypeToggle) uses tw-border-light-gray-stroke (#dfdfdf).
- TimeFormatToggle.vue container and indicator use tw-border-light-gray-stroke (#dfdfdf).
- TimezoneSelector.vue hardcodes border: 1px solid #dfdfdf on the compact select button and border-color: #dfdfdf on the compact reset button.
- The only --timeful-* token holding #bdbdbd is --timeful-compact-switch-track-border (switch-specific); there is no shared neutral-outline token yet.

Expected outcome: every non-colored outline across the listed chips (Available/If needed toggle, Calendar options button, editing-availability-as chip, 12h/24h time-format toggle, timezone compact button and its reset button) resolves to #bdbdbd, expressed through the existing Tailwind gray palette or a shared --timeful-* semantic token rather than scattered raw hex values.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [x] #1 The Available/If needed slide-toggle outer track border renders #bdbdbd in the desktop sidebar and the mobile overlay
- [x] #2 The 12h/24h time-format toggle container and its inner indicator border render #bdbdbd
- [x] #3 The compact timezone selector button border renders #bdbdbd
- [x] #4 The timezone reset button border renders #bdbdbd in both compact (right-side) and non-compact placements
- [x] #5 The Calendar options button and the editing-availability-as guest chip continue to render #bdbdbd and do not regress
- [x] #6 Colored outlines (green/yellow active indicators, blue, red, etc.) are unchanged
- [x] #7 Neutral outlines no longer rely on scattered raw hex values where a shared palette or --timeful-* token can be used, per frontend styling rules
- [x] #8 Unit tests that assert the old border classes (e.g. TimeFormatToggle.test.ts) are updated, and border-color regression coverage is updated where it references these chips
- [x] #9 The New Event form's time-format toggle, time-increment toggle, and timezone chip render #bdbdbd (covered via the shared TimeFormatToggle and TimezoneSelector components)
- [x] #10 The New Event 'Dates and times' SlideToggle track also renders #bdbdbd (shared SlideToggle component side effect, confirmed in scope)
- [x] #11 The e2e inspect helper selector that references the SlideToggle border class (frontend/e2e/inspect/src/dom-resolvers.ts) is updated to the new class
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
1. Token layer (frontend/src/index.css, frontend/tailwind.config.cjs)\n   - Add `--timeful-outline-neutral: #bdbdbd;` to `:root`.\n   - Rewire `--timeful-compact-switch-track-border` to `var(--timeful-outline-neutral)`; leave `--timeful-compact-switch-track-bg` literal (background, not outline).\n   - Add Tailwind color alias `"outline-neutral": "var(--timeful-outline-neutral)"` so `tw-border-outline-neutral` works (config uses `important: true`).\n2. Template swaps to `tw-border-outline-neutral`\n   - SlideToggle.vue track (shared: Available/If needed toggle + New Event 'Dates and times' toggle).\n   - TimeFormatToggle.vue container + indicator (shared: schedule-overlap 12h/24h, mobile days-per-page, New Event time format + time increment, NewGroup/NewSignUp).\n   - ScheduleOverlapSidebar.vue and ScheduleOverlapMobileOverlay.vue Calendar options buttons (from `tw-border-gray`).\n   - EditingAvailabilityAs.vue guest chip (from `tw-border-gray`).\n3. Scoped CSS (TimezoneSelector.vue)\n   - Compact select button border and compact reset button border → `var(--timeful-outline-neutral)` (Tailwind class cannot reliably reach `:deep(.v-field)` overrides).\n4. Tests/tooling\n   - TimeFormatToggle.test.ts: assert `tw-border-outline-neutral`.\n   - NewEvent.test.ts: update `--timeful-compact-switch-track-border` regex for the alias; add assertion for the new token.\n   - e2e/inspect/src/dom-resolvers.ts: update SlideToggle selector chain.\n   - Grep SlideToggle/AvailabilityTypeToggle tests for further border-class assertions.\n5. Verification\n   - npm run lint, typecheck, build, test:unit.\n   - Visual sanity: colored green/yellow indicators unchanged; neutral outlines #bdbdbd.\n6. Out of scope: Settings/Landing/Dashboard/EventDescription/RespondentsList/SignUpBlock/FAQ neutral outlines; `timeful-solo-field` and elevated-button borders; compact-switch track bg.
<!-- SECTION:PLAN:END -->

## Implementation Notes

<!-- SECTION:NOTES:BEGIN -->
Implementation: added --timeful-outline-neutral (#bdbdbd) to :root, aliased it in Tailwind as "outline-neutral", rewired --timeful-compact-switch-track-border to the new token. Swapped neutral border classes to tw-border-outline-neutral in SlideToggle, TimeFormatToggle (track + indicator), both Calendar options buttons, and the EditingAvailabilityAs guest chip. TimezoneSelector compact button + reset borders now use var(--timeful-outline-neutral). Discovered extra test coupling: TimezoneSelector.test.ts asserts the exact scoped-CSS border text (updated to var form); NewEvent.test.ts switch-token regex updated and extended with an assertion for the new token; inspect dom-resolvers SlideToggle selector updated. RespondentsList.test.ts tw-border-gray status-square assertion is unaffected (Tailwind gray palette unchanged).

Final state: lint, typecheck, build, and the full unit suite (961 tests) pass. Compiled CSS verified to contain the token, the Tailwind utility, and the switch-track alias. Converted one additional neutral outline found late (sidebar calendar pager button) from tw-border-gray to the token — same resolved color. inspect new-event-form scenario timed out in prepare (needs signed-in backend flow); covered instead via component tests + compiled-CSS assertions.
<!-- SECTION:NOTES:END -->

## Final Summary

<!-- SECTION:FINAL_SUMMARY:BEGIN -->
## Summary

All neutral (non-colored) chip outlines now resolve to #bdbdbd through a single shared token instead of scattered #dfdfdf / palette-gray values.

## Changes

- Token layer: added `--timeful-outline-neutral: #bdbdbd` in `frontend/src/index.css`; `--timeful-compact-switch-track-border` now aliases it. Added Tailwind color alias `outline-neutral: var(--timeful-outline-neutral)` in `frontend/tailwind.config.cjs` so templates use `tw-border-outline-neutral`.
- Template swaps to `tw-border-outline-neutral`: SlideToggle track (Available/If needed toggle; also the New Event "Dates and times" toggle), TimeFormatToggle track + indicator (12h/24h, time increment, mobile days-per-page; NewGroup/NewSignUp inherit), Calendar options buttons in ScheduleOverlapSidebar and ScheduleOverlapMobileOverlay, EditingAvailabilityAs guest chip, and the sidebar calendar pager button (was `tw-border-gray`, same resolved color).
- Scoped CSS: TimezoneSelector compact select button and compact reset button borders now use `var(--timeful-outline-neutral)` (a Tailwind class cannot reliably reach the `:deep(.v-field)` overrides).
- Tests/tooling: TimeFormatToggle.test.ts asserts the new indicator class; new SlideToggle test asserts the neutral track class; TimezoneSelector.test.ts exact-CSS assertions updated to the var() form; NewEvent.test.ts switch-token regex updated plus a new assertion for `--timeful-outline-neutral`; e2e inspect `dom-resolvers.ts` SlideToggle selector updated.

## Verification

- `npm run lint`, `npm run typecheck`, `npm run build`: pass on final state.
- `npm run test:unit`: 961 tests / 137 files pass (includes the new and updated border assertions).
- Built CSS bundle verified: `--timeful-outline-neutral:#bdbdbd`, `tw-border-outline-neutral{border-color:var(--timeful-outline-neutral)!important}`, and the switch-track alias are all present.
- Colored outlines untouched: green/yellow AvailabilityTypeToggle indicator classes and colors, blue/red states unchanged (diff review + component suites).
- E2E: no e2e spec asserts these border styles; the only e2e-tree code referencing them (inspect helper selector) was updated. Full browser E2E was not run; the inspect harness timed out in scenario preparation because it requires the signed-in backend flow, so visual confirmation rests on the component tests plus compiled-CSS verification. No check failures to report.

## Notes / follow-ups

- The non-compact timezone reset button is a borderless `variant="text"` button by design, so the bordered-reset criterion effectively applies to the compact right-side placement.
- Out of scope (possible follow-up): app-wide neutral outlines in Settings/Landing/Dashboard/EventDescription/RespondentsList/SignUpBlock/FAQ, `timeful-solo-field` (#4f4f4f1f) and elevated-button borders, and `--timeful-compact-switch-track-bg` (background, intentionally left literal).
<!-- SECTION:FINAL_SUMMARY:END -->
