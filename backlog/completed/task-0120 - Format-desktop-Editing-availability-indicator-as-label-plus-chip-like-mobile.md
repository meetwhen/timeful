---
id: TASK-0120
title: Format desktop Editing availability indicator as label plus chip like mobile
status: Done
assignee: []
created_date: '2026-08-31 12:24'
updated_date: '2026-08-31 12:30'
labels: []
dependencies: []
priority: medium
type: enhancement
ordinal: 127300
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
The mobile editing panel renders the Editing availability indicator as a non-italic label plus a rectangular editable-name chip (EditingAvailabilityAs variant="chip", TASK-0118). The desktop sidebar still renders the italic sentence variant.

Restyle the desktop sidebar indicator to match the mobile formatting:
- label (non-italic) + chip
- the indicator spans the full sidebar width
- the label and chip align to their left edge (not the mobile right alignment)

Keep the mobile overlay right-aligned and keep the phone sidebar sentence rendering for group events unchanged.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [x] #1 The desktop sidebar Editing availability indicator renders the chip variant: a non-italic label with the editable guest name in a rectangular chip button
- [x] #2 The desktop chip indicator spans the full sidebar width with the label and chip aligned to the left edge instead of right-aligned
- [x] #3 The mobile editing panel indicator keeps its current right-aligned chip presentation
- [x] #4 The phone sidebar indicator for group events keeps the sentence rendering
- [x] #5 npm run lint, npm run typecheck, npm run build, and npm run test:unit pass in frontend/
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
1. EditingAvailabilityAs.vue: drop `tw-justify-end` from the chip variant root so chip alignment is owned by the usage site.
2. ScheduleOverlapMobileOverlay.vue: pass `class="tw-justify-end"` at the usage site to keep mobile right-aligned.
3. ScheduleOverlapSidebar.vue: pass `variant="chip"` only when `!sidebar.isPhone` so desktop uses chip and phone group sidebar keeps sentence.
4. Tests: update EditingAvailabilityAs chip tests (chip no longer inherently right-aligned), update sidebar desktop test to assert chip variant and left alignment, assert phone group sidebar stays sentence, keep mobile overlay assertions (justify-end now via fallthrough).
5. Run lint, typecheck, build, test:unit.
<!-- SECTION:PLAN:END -->

## Final Summary

<!-- SECTION:FINAL_SUMMARY:BEGIN -->
Desktop now renders the Editing availability indicator like mobile: a non-italic "…availability as" label with the editable guest name in a rectangular chip button, spanning the full sidebar width with the label and chip left-aligned.

- EditingAvailabilityAs.vue: the chip variant no longer hard-codes right alignment (`tw-justify-end` removed from the root class binding), so alignment is owned by the usage site.
- ScheduleOverlapMobileOverlay.vue: passes `class="tw-justify-end"` at the usage site, keeping the mobile panel's right-aligned chip presentation unchanged.
- ScheduleOverlapSidebar.vue: passes `variant="chip"` when `!sidebar.isPhone`; the phone sidebar keeps the sentence rendering for group events.

Tests:
- EditingAvailabilityAs.test.ts: chip-variant test now asserts the block is non-italic and not right-aligned; default-sentence test renamed to match the invariant.
- ScheduleOverlapSidebar.test.ts: new desktop test asserts the chip variant, non-italic left-aligned block, and the name chip content; the phone-group test asserts the sentence variant with no chip block.
- ScheduleOverlapMobileOverlay.test.ts: existing right-alignment assertions still pass via the usage-site fallthrough class.

Checks: `npm run lint`, `npm run typecheck`, `npm run build`, and `npm run test:unit` (137 files, 960 tests) all pass in `frontend/`; `npm run format:markdown` run with no changes. No e2e specs cover this indicator and none were required; both usage sites are covered by unit tests.
<!-- SECTION:FINAL_SUMMARY:END -->
