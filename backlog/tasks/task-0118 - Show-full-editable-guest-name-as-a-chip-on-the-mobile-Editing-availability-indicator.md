---
id: TASK-0118
title: >-
  Show full editable guest name as a chip on the mobile Editing availability
  indicator
status: Done
assignee:
  - opencode
created_date: '2026-08-30 20:37'
updated_date: '2026-08-31 11:38'
labels:
  - mobile
  - ui
  - schedule-overlap
dependencies: []
modified_files:
  - frontend/src/components/schedule_overlap/EditingAvailabilityAs.vue
  - frontend/src/components/schedule_overlap/EditingAvailabilityAs.test.ts
  - frontend/src/components/schedule_overlap/ScheduleOverlapMobileOverlay.vue
  - >-
    frontend/src/components/schedule_overlap/ScheduleOverlapMobileOverlay.test.ts
priority: medium
type: enhancement
ordinal: 125300
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
## Problem

On phone-width screens the editing panel shows "Editing availability as d..." with the guest name truncated. The editable name affordance (pencil icon) is easy to miss, and the long sentence wastes the row.

## Shipped design (Aug 30, 2026)

The mobile indicator is a right-aligned row: label "Editing availability as" + inline chip button (underlined italic name, pencil icon) that opens the existing edit-guest-name dialog. Wrap behavior drops the chip below the label; long names break inside the chip. Placeholder "Respondent name" for empty names; non-editable states stay plain text.

## Revision (Aug 31, 2026) — styling adjustments requested by the owner

The first chip styling needs correction; behavior (dialog, emits, placeholder, non-editable text, wrap) stays as shipped:

- The chip is rectangular with rounded corners like other buttons (Vuetify `v-btn` default 4px radius → `tw-rounded`), not pill-shaped; remove `tw-rounded-full`.
- No underline on the chip name; remove `tw-underline` from the chip name span.
- Alignment: the label and the chip form one block that is always aligned to the right edge of the row; within that block, the label and the chip are aligned by their left edges (outer indicator row right-aligned via `tw-justify-end`, inner label+chip row left-aligned with flex-wrap so the chip wraps under the label sharing its left edge).
- Chip text aligns left, not centered (`tw-text-left` on the chip button).
- The chip-variant indicator is not italicized (`tw-not-italic` on the chip root); the default sentence variant keeps its current italic rendering.

## Revision 2 (Aug 31, 2026) — pencil pinned to the chip's right edge

Within the chip button, the pencil icon must always sit at the chip's right edge (inside the right padding), not immediately after the last character of the name. The name span grows to fill the space between the name text and the icon so the icon is pinned right for both short and long names. All prior behavior (dialog, emits, placeholder, non-editable text, wrap, right-aligned block) stays unchanged.

### Design goals behind the chip

- Use most of the available width: the chip grows to fill the row so the long guest name has room and the tap target is large.
- Keep the pencil clickable with the right hand: pinning the pencil at the chip's right edge puts it at a natural right-thumb reach on phone-width screens.
- Make the chip look like a button: the bordered, white-background, rounded style (with hover state) signals that the name must be clicked to edit it.

## Out of scope

- Desktop sidebar presentation, group-event presentation, view-model contract changes, and the Delete/Cancel/Save mobile action bar.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [x] #1 On the phone-width editing panel the label "Editing availability as" (or "Adding availability as") and the chip form one block aligned to the right edge of the row, and within that block the label and the chip are aligned by their left edges (outer row right-aligned, inner label+chip row left-aligned with flex-wrap)
- [x] #2 The chip is a rectangle with rounded corners like other buttons (tw-rounded, not tw-rounded-full) and keeps the pencil icon
- [x] #3 The chip name has no underline
- [x] #4 Chip text aligns left, not centered
- [x] #5 The chip-variant indicator is not italicized; the default sentence variant rendering is unchanged (still italic)
- [x] #6 Behavior from the shipped design is unchanged: clicking the chip opens the edit-guest-name dialog, "Respondent name" placeholder when the editable name is empty, non-editable states render plain text, and long names wrap/break without truncating
- [x] #7 Unit tests updated: EditingAvailabilityAs.test.ts and ScheduleOverlapMobileOverlay.test.ts assert the revised styling (right-aligned block, left-aligned label+chip row, rectangular chip, no underline, left-aligned text, not italic)
- [x] #8 npm run lint, npm run typecheck, npm run build, and npm run test:unit all pass in frontend/
- [x] #9 The pencil icon is always at the right edge of the chip (inside the right padding, not adjacent to the name text), for both short and long names
- [x] #10 Tests assert the chip name span grows so the pencil is pinned to the chip's right edge, and all required frontend checks pass
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
## Revision plan (Aug 31, 2026, updated after owner clarification)

Research confirmed:
- Vuetify `v-btn` default border-radius is 4px, so `tw-rounded` (4px) matches the "rounded corners like other buttons" requirement; the current pill uses `tw-rounded-full`.
- Final alignment (clarified by owner): the label + chip form one block always aligned to the right edge of the row; inside the block, label and chip share their left edge. This needs two levels: outer indicator row right-aligned (`tw-justify-end`), inner label+chip row left-aligned with `tw-flex-wrap` so the wrapped chip lands under the label sharing its left edge.
- Chip button has no explicit text alignment, so button text defaults to centered; needs `tw-text-left`.
- Label italic comes from the root `tw-italic` shared by both variants; chip variant needs `tw-not-italic` on the root.

Steps:
1. EditingAvailabilityAs.vue: chip variant root keeps `tw-justify-end` (right-aligned block) and adds `tw-not-italic`; chip content moves into an inner `editing-availability-as__chip-row` (`tw-flex tw-flex-wrap tw-items-baseline tw-gap-1`, left-aligned) holding the label text, the chip button (`tw-rounded`, `tw-text-left`, no underline on the name span), or the plain actor span. Sentence variant markup stays DOM-identical (wrapped in `<template v-else>`, no extra DOM); dialog and emits shared outside both branches.
2. EditingAvailabilityAs.test.ts: update chip-variant assertions — indicator has tw-justify-end + tw-not-italic; `__chip-row` exists without tw-justify-end and with tw-flex-wrap; chip has tw-rounded + tw-text-left and not tw-rounded-full; name has no tw-underline; keep emit, placeholder, and non-editable tests.
3. ScheduleOverlapMobileOverlay.test.ts: indicator asserts tw-justify-end + tw-not-italic on the row, left-aligned `__chip-row`, variant prop, text, and ordering above the toggle row.
4. Run required checks from frontend/: npm run lint, npm run typecheck, npm run build, npm run test:unit.
5. Finalize: verify each AC, record final summary, mark Done.

Risks/checkpoints:
- The sentence branch keeps identical rendered output (the `<template v-else>` wrapper produces no DOM); sidebar guard tests must pass unchanged.
- Wrap behavior classes (tw-flex-wrap on the chip row, tw-min-w-0, tw-break-words, tw-max-w-full) must remain so non-truncation stays satisfied.
<!-- SECTION:PLAN:END -->

## Implementation Notes

<!-- SECTION:NOTES:BEGIN -->
Implementation notes:
- EditingAvailabilityAs.vue: optional `variant?: "sentence" | "chip"` prop (default "sentence"). Chip mode adds `editing-availability-as--chip tw-justify-end` to the root (which already has tw-flex-wrap); the editable name renders as `<button type="button" class="editing-availability-as__guest-chip">` with the calendar-options outlined treatment (tw-rounded-full tw-border tw-border-solid tw-border-gray tw-bg-white, hover:tw-bg-light-gray, tw-appearance-none/tw-shadow-none resets since Tailwind preflight is off), pencil v-icon, and a `tw-underline` name span (`editing-availability-as__guest-name`). Chip + name span carry tw-min-w-0 (chip also tw-max-w-full) and the name tw-break-words so long names break within the chip and the chip drops below the right-aligned label. Chip text falls back to "Respondent name" when editableGuestName is "". Non-editable states keep the unchanged plain actor span. Dialog and all four emits are shared by both variants, outside the variant branches.
- The label was first wrapped in a span for a test hook; this broke `text()` spacing ("asa guest") because Vue condense removes whitespace-only nodes between elements — reverted to the original bare text node (identical default rendering; sidebar guard tests unchanged and passing). Flex tw-gap-1 provides visual spacing in both variants.
- ScheduleOverlapMobileOverlay.vue: passes `variant="chip"` to the EditingAvailabilityAs usage; the Calendar-options + AvailabilityTypeToggle row is untouched.
- Tests: 7 new EditingAvailabilityAs cases (default-sentence guard, right-aligned chip row with button/pencil, click emit, chip dialog Save/Cancel relay, wrap/break classes, "Respondent name" placeholder, non-editable plain text with no chip/icon). ScheduleOverlapMobileOverlay tests: indicator test now asserts variant="chip" prop, tw-justify-end row, and ordering above the toggle row; re-emit test uses the chip selector.
- One lint iteration: DOMWrapper.get() has no exists() — used find().exists().

Revision (Aug 31, 2026): EditingAvailabilityAs.vue chip variant restructured into two levels — root keeps tw-justify-end (the label+chip block sits at the right edge of the row) and gains tw-not-italic; chip content moved into an inner .editing-availability-as__chip-row (tw-flex tw-flex-wrap tw-items-baseline tw-gap-1, left-aligned) so the label and chip share their left edge and the wrapped chip lands under the label. Chip button: tw-rounded-full → tw-rounded (matches Vuetify v-btn default 4px radius) and added tw-text-left; name span dropped tw-underline. Sentence variant markup is DOM-identical (nodes wrapped in <template v-else>, which produces no DOM); dialog and emits stay shared outside both branches.

Owner clarified alignment mid-review: not fully left-aligned — the label+chip construct forms a right-aligned block whose internal left edges align; AC #1 and the plan were updated accordingly.

Tests: EditingAvailabilityAs.test.ts chip-row test asserts indicator has tw-justify-end + tw-not-italic, __chip-row exists without tw-justify-end and with tw-flex-wrap, chip has tw-rounded + tw-text-left (not tw-rounded-full), name has no tw-underline; wrap test now targets __chip-row. ScheduleOverlapMobileOverlay.test.ts asserts right-aligned block + left-aligned chip row + tw-not-italic. Sidebar guard tests pass unchanged (44 tests across the three touched spec files).

Revision 2 (Aug 31, 2026): owner asked for the pencil to align to the right edge of the chip instead of following the text end. Root cause: the chip button is a flex row (name span, then v-icon) and the name span had no flex-grow, so leftover space stayed between the icon and the chip's right edge. Fix: added tw-grow to the name span (kept tw-min-w-0/tw-break-words) so the span fills the available width and pins the icon right; text stays left-aligned via the button's tw-text-left. Tests: tw-grow asserted on the name span in the chip-styling test (short name) and the wrap test (long name). lint, typecheck, build, and test:unit (959 tests) all pass; e2e not re-run (no selector/behavior changes; same feature passed firefox-desktop e2e earlier today).
<!-- SECTION:NOTES:END -->

## Final Summary

<!-- SECTION:FINAL_SUMMARY:BEGIN -->
## Summary

Applied the owner's styling revision to the mobile editing-availability chip (TASK-0118 revision, Aug 31, 2026).

### Changes

- `frontend/src/components/schedule_overlap/EditingAvailabilityAs.vue` (chip variant only): the label + chip now form one block that is always aligned to the right edge of the row, with the label and chip sharing their left edge inside it (root keeps `tw-justify-end`, content moved into a left-aligned `editing-availability-as__chip-row` with `tw-flex-wrap`); chip button is rectangular with rounded corners like other buttons (`tw-rounded`, was `tw-rounded-full`) with left-aligned text (`tw-text-left`); the chip name has no underline; the chip-variant indicator is not italic (`tw-not-italic`). The sentence variant's rendered output is unchanged (nodes wrapped in `<template v-else>`, which produces no DOM); dialog and emits are shared and unchanged.
- Tests: `EditingAvailabilityAs.test.ts` chip assertions updated (right-aligned block + left-aligned `__chip-row`, `tw-rounded` not `tw-rounded-full`, `tw-text-left`, no `tw-underline`, `tw-not-italic`; wrap test targets `__chip-row`); `ScheduleOverlapMobileOverlay.test.ts` asserts the right-aligned block, left-aligned chip row, and `tw-not-italic`. Sidebar sentence-rendering guard tests pass unchanged.

### Check results (all from frontend/)

- `npm run lint` PASS
- `npm run typecheck` PASS
- `npm run build` PASS
- `npm run test:unit` PASS: 137 files, 959 tests
- `npm run test:e2e -- --project=firefox-desktop` PASS: 22 passed, 1 skipped (by design); Playwright owned the isolated test stack (mongo-test, postgres-test, server-test on 3003, Vite on 4174)

No Markdown files changed, so the format:markdown DoD item is vacuous.

## Revision 2 (Aug 31, 2026) — pencil pinned to the chip's right edge

### Changes

- `frontend/src/components/schedule_overlap/EditingAvailabilityAs.vue`: the chip name span (`editing-availability-as__guest-name`) gained `tw-grow`, so it fills the space between the name text and the pencil icon inside the flex chip button. The pencil is therefore always pinned at the chip's right edge (inside the right padding) instead of sitting immediately after the last character, for both short and long names. `tw-min-w-0`/`tw-break-words` remain, so long names still wrap within the chip; no other classes, behavior, or the sentence variant changed.
- Tests: `EditingAvailabilityAs.test.ts` asserts `tw-grow` on the name span in both the chip-styling test (short name) and the wrap test (long name), covering AC #9 for both cases.

### Check results (all from frontend/)

- `npm run lint` PASS
- `npm run typecheck` PASS
- `npm run build` PASS
- `npm run test:unit` PASS: 137 files, 959 tests

E2E was not re-run for this one-class layout change: it adds no selector or behavior changes, the new layout is asserted at the unit level, and the firefox-desktop e2e suite passed on the prior revision of this same feature earlier today.
<!-- SECTION:FINAL_SUMMARY:END -->
