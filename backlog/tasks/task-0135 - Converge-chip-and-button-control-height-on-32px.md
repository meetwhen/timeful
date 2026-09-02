---
id: TASK-0135
title: Converge chip and button control height on 32px
status: To Do
assignee: []
created_date: '2026-09-02 10:27'
labels:
  - frontend
  - styling
  - design-tokens
dependencies: []
references:
  - frontend/src/App.vue
  - frontend/src/index.css
  - frontend/src/components/SignInGoogleBtn.vue
  - frontend/src/components/ExpandableSection.vue
  - frontend/src/components/TimeRangePicker.vue
  - frontend/src/components/general/UserChip.vue
  - frontend/src/components/EventItem.vue
  - frontend/src/components/home/Dashboard.vue
  - frontend/src/views/Event.vue
  - frontend/src/components/schedule_overlap/TimezoneSelector.vue
  - frontend/src/components/schedule_overlap/ScheduleOverlapCompactSwitch.css
  - frontend/src/components/TimeRangePicker.test.ts
  - frontend/src/components/schedule_overlap/TimezoneSelector.test.ts
documentation:
  - frontend/AGENTS.md
  - docs/design/architecture/adr/ADR-001.md
priority: medium
type: enhancement
ordinal: 148300
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
## Outcome

Converge the standard control height for chips and buttons on **32px** so chips, buttons, and the existing compact fields form one size scale. This decision was made in a sizing-inventory session (Sept 2026): 32px is Vuetify 3's native chip height and the app is chip-heavy, so buttons move to chips instead of overriding chips away from the framework.

## Verified current state (frontend/src)

Vuetify 3 defaults: v-btn 36px, v-chip 32px (chip "small" would be 26px, x-small 20px). Real heights in the app today: 20 / 26 / 28 / 32 / 36 / 38 / 40 / 58px, with two silent no-ops and one dead rule.

- `App.vue:361-369` — `.v-btn:not(.v-btn--round, .v-btn-toggle > .v-btn).v-size--default { height: 38px !important; border-radius: 0.375rem !important }` is dead CSS: `.v-size--default` is a Vuetify 2 class that does not exist in Vuetify 3, so nothing matches. Buttons actually render at the Vuetify default 36px.
- `components/SignInGoogleBtn.vue:63` — custom 40px button (the only 40px control).
- `components/ExpandableSection.vue:79` — 38px min-height compact accordion header.
- `components/schedule_overlap/TimezoneSelector.vue` — compact fields/buttons already 32px (its test asserts these values); inline selects are 26px. No change needed; they converge automatically.
- `components/schedule_overlap/ScheduleOverlapCompactSwitch.css` — 28px switch track; not in scope.
- Legacy no-op chip props (Vuetify 2 booleans, inert attributes in Vuetify 3, so these chips render default 32px): `components/EventItem.vue:48` (`small`), `components/home/Dashboard.vue:36` (`small`), `views/Event.vue:168` (`:small="isPhone"`, whose original intent was a smaller chip on phone).
- `components/general/UserChip.vue:4` — `size="x-small"` (20px) is the intended dense step for user chips; keep.
- `components/TimeRangePicker.vue:112` — `--time-range-chip-height: 58px` is an intentional two-line field, already pinned by `TimeRangePicker.test.ts`; keep as the documented exception.

## Constraints

- Follow `frontend/AGENTS.md` styling rules: use existing `--timeful-*` semantic tokens from `frontend/src/index.css` rather than component-local raw values; prefer layout-based fixes and plain selectors in non-scoped style blocks; use `:deep(...)` only in scoped styles; verify rendered selectors in the browser.
- Introduce one shared token (for example `--timeful-control-height: 32px`) in `index.css :root` and wire it into the button/chip sizing rules, so future changes are one-line. Vuetify exposes `--v-btn-height` / `--v-chip-height` custom properties that sizing rules can pin.
- Note: migrating the no-op `small` props to working `size="small"` would change rendered size (32 -> 26px). Preserve current rendered size unless the original phone-smaller intent is explicitly wanted; when in doubt keep 32px and flag in the task notes.
- Mobile touch targets (44px) are explicitly out of scope; a phone-density bump for primary actions is a possible follow-up task.
- After changing route/API annotations this would need swag regen, but this task is CSS-only; no swag or server changes expected.

## Suggested starting points

Read `frontend/AGENTS.md` first, then `frontend/src/index.css` (token home), `frontend/src/App.vue` (dead rule), and the component files listed above. The CSS-assertion test patterns in `frontend/src/components/TimeRangePicker.test.ts` and `frontend/src/components/schedule_overlap/TimezoneSelector.test.ts` show how sizing rules are pinned in this repo.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 All standard v-btn and v-chip controls render at a shared 32px height; no standard control remains at 36px, 38px, or 40px.
- [ ] #2 The dead .v-size--default rule in App.vue is removed or replaced with a Vuetify 3-valid rule driven by the shared token; button border-radius matches the app's existing 0.375rem convention (0.375rem is used in TimezoneSelector and .timeful-solo-field) with the rendered result verified in the browser.
- [ ] #3 The shared height token is defined in frontend/src/index.css :root as a --timeful-* semantic token and consumed by the sizing rules; no new scattered height literals and no new !important overrides beyond what existing overrides already use.
- [ ] #4 All legacy Vuetify 2 boolean chip size props (small / :small="isPhone") are migrated to explicit Vuetify 3 size props; no inert size attributes remain in the codebase.
- [ ] #5 The 58px TimeRangePicker chip and the x-small (20px) dense chips/icon buttons keep their current rendering; TimeRangePicker.test.ts and TimezoneSelector.test.ts stay green.
- [ ] #6 Sizing rules are covered by unit-test assertions in the style of the existing CSS-assertion specs (TimeRangePicker.test.ts), including a check that pins the shared token value.
- [ ] #7 npm run lint, npm run typecheck, npm run build, and npm run test:unit all pass in frontend/.
<!-- AC:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [ ] #1 All acceptance criteria are satisfied
- [ ] #2 All required unit tests pass. Documentation-only changes are exempt unless the user requests unit tests
- [ ] #3 All required e2e tests pass. Documentation-only changes are exempt unless the user requests e2e tests
- [ ] #4 Changed Markdown files are formatted with npm run format:markdown
<!-- DOD:END -->
