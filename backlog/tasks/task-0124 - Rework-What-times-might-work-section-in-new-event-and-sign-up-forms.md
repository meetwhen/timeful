---
id: TASK-0124
title: Rework "What times might work?" section in new event and sign-up forms
status: In Progress
assignee: []
created_date: '2026-08-31 14:40'
updated_date: '2026-08-31 20:04'
labels:
  - frontend
  - ui
dependencies: []
priority: medium
type: enhancement
ordinal: 131300
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Restructure the "What times might work?" section in NewEvent.vue and NewSignUp.vue per user-approved decisions:

- Move the "Set specific times per day" switch to the top of the section (first row under the heading), keeping its helper message.
- Toggle off: 12h/24h format toggle and the start-to-end range selects share one row (format toggle left, range right). Toggle on: the whole row collapses (animated), message stays.
- Range chips match the 32px TimeFormatToggle height; dropdown items ~40px.
- Selects and menus shrink to fit one row inside the card (~384px inner width).
- Apply the same one-row layout and heights to NewSignUp.vue (no specific-times toggle there).

Decisions confirmed with user: toggle under heading; shrink selects; menu items ~40px; toggle left / range right; match both forms; keep helper message.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 In NewEvent.vue, the "Set specific times per day" switch is the first row under the "What times might work?" heading, keeping data-testid="specific-times-toggle" and the "Click the Next button below" helper message
- [ ] #2 In NewEvent.vue, when the switch is off, the 12h/24h toggle and the start-to-end selects render in one row (toggle left, range right); when on, the entire row is hidden with a collapse transition
- [ ] #3 Range select chips and the "to" separator are 32px tall matching TimeFormatToggle; time-range dropdown list items are ~40px min-height in both NewEvent.vue and NewSignUp.vue
- [ ] #4 The combined row fits inside the new event card without overflow or wrapping at default width, with selects and dropdown menus shrunk accordingly
- [ ] #5 NewSignUp.vue renders the time-format toggle and range in the same one-row layout with the same chip and menu-item heights
- [ ] #6 npm run lint, typecheck, build, and test:unit pass in frontend/
- [ ] #7 Targeted firefox e2e specs for timed-event create and specific-times flows pass
<!-- AC:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [ ] #1 All acceptance criteria are satisfied
- [ ] #2 All required unit tests pass. Documentation-only changes are exempt unless the user requests unit tests
- [ ] #3 All required e2e tests pass. Documentation-only changes are exempt unless the user requests e2e tests
- [ ] #4 Changed Markdown files are formatted with npm run format:markdown
<!-- DOD:END -->

## Implementation Notes

<!-- SECTION:NOTES:BEGIN -->
## Decisions confirmed with user (2026-08-31)
- Specific-times toggle goes directly under the "What times might work?" heading as the first row.
- Shrink the two range selects (and menus) to fit one row; menu items ~40px (chips 32px matching TimeFormatToggle).
- Combined row: format toggle left, "start to end" right (space-between).
- Apply same one-row layout + heights to NewSignUp.vue (no specific-times toggle there).
- When toggle enabled: whole format+range row collapses (v-expand-transition), "Click the Next button below" message stays.

## Implementation (done)
frontend/src/components/NewEvent.vue:
- Moved the specific-times switch block (keeps data-testid="specific-times-toggle", compact-switch-grid classes, helper message) to be the first row under the heading.
- Merged TimeFormatToggle + start/end selects into a single `.time-range-row` (tw-flex tw-items-center tw-justify-between) inside v-expand-transition with v-if="!specificTimesEnabled" (previously only the range row collapsed and the format toggle stayed visible).
- Range select menu-props width 176 -> 132.
- Styles: `.time-range-row` vars: --time-range-control-height 44->32px, --time-range-select-width 176->132px; added `.time-range-row .v-input.time-range-select { --v-input-control-height: 32px; width }` and `.time-range-row .v-input.time-range-select .v-field { --v-field-input-padding-top: 0px; --v-field-input-padding-bottom: 4px }`; `.time-range-select-item` min-height 48->40px.
- WHY this shape: Vuetify field min-height = max(--v-input-control-height, 1.5rem + input-padding-top + input-padding-bottom). Defaults are set directly on the elements: `.v-input--density-default { --v-input-control-height: 56px }` (0,2,0) and `.v-field { --v-field-input-padding-* }`. The wrapper-scoped selectors (0,3,0)/(0,4,0) beat them without relying on source-order tie-breaks. Result: max(32, 24+0+4)=32px chip.

frontend/src/components/NewSignUp.vue:
- Removed the separate TimeFormatToggle row; same one-row layout and classes as NewEvent.
- Replaced legacy `item-color="green"` on the two range selects with the same `#item` slot markup + `.time-range-select-item--active` semantic tokens as NewEvent (gives 40px items + active highlight). The date-option select still uses item-color="green".
- Same style updates as NewEvent (32px chip, 132px width, 40px items).

Known shared-class side effect: the date-option dropdown menu items in both files also use `.time-range-select-item`, so their items are now 40px too (deemed consistent).

Width math: card inner width 384px (max-w-28rem 448px - sm:px-8); toggle 80px + 2x132px selects + "to" separator ~16px + 2x8px gaps = 376px, fits with slack.

## Test changes (done)
- NewEvent.test.ts: menuProps assertions 176->132 (2 places); added 3 behavioral tests: toggle-first ordering (uses compareDocumentPosition), one-row layout when disabled, row hidden + message kept when enabled.
- NewSignUp.test.ts: itemColor/menuProps assertions updated; added one-row layout test.
- Unit-test gotchas discovered: `v-expand-transition` is a passThroughStub that wraps the row in a plain div, so element.previousElementSibling is null - use compareDocumentPosition for ordering. `[data-testid='time-format-toggle-stub']` matches BOTH the section toggle and the advanced-options increment toggle - scope finders to `.time-range-row`. shallowMount auto-stub tag name is not `timeformat-toggle-stub`; use findComponent({ name: "TimeFormatToggle" }). In NewEvent.test.ts keep `if (!heading) throw` (not toBeDefined) to satisfy TS18048.

## Verification status
- npm run lint: PASS
- npm run typecheck: PASS
- npm run build: PASS
- npm run test:unit: PASS (137 files / 965 tests)

## E2E (in progress - the only remaining work)
- Ran `npm run test:e2e -- --project=firefox-desktop timed-event-create-firefox.spec.ts` twice.
  - Run 1: failed at page.goto("/") 30s timeout (cold isolated stack).
  - Run 2: 1 passed; 1 failed at spec line 55 with the SAME page.goto timeout (http://127.0.0.1:4174/ domcontentloaded); 4 did not run.
- Both failures are goto timeouts waiting for the Vite server on 4174, not UI assertion failures - environment flakiness, likely Vite/test-stack instability mid-run.

### Next steps for the next session
1. Re-run `npm run test:e2e -- --project=firefox-desktop timed-event-create-firefox.spec.ts` until green or a real assertion failure appears; if goto timeouts persist, inspect the Vite process output/log in the Playwright webServer config.
2. Run `timed-event-specific-times-edit-firefox.spec.ts` and `timed-event-time-format-toggle-alignment-firefox.spec.ts` (the alignment spec only asserts toggle internals, DOM order change should be safe).
3. Browser-verify: one-row fit at 384px inner width without wrap; range chip height == 32px toggle height; menu items 40px; dropdown menu width 132px. Use `npm run inspect -- --target <scenario>` per e2e/inspect/AGENTS.md.
4. Check all acceptance criteria, mark them, add final summary, set task Done.
<!-- SECTION:NOTES:END -->
