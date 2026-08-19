---
id: TASK-0028
title: 'F-SIGN-UP-BLOCKS-LIST-001: extract scroll/sizing into composable contracts'
status: Done
assignee: []
created_date: '2026-08-19 13:40'
updated_date: '2026-08-19 14:00'
labels:
  - findings
dependencies: []
references:
  - frontend/findings/F-SIGN-UP-BLOCKS-LIST-001.md
priority: high
type: task
ordinal: 28000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Component: src/components/sign_up_form/SignUpBlocksList.vue

Problem: The list owns query-selector-based scrolling, global resize handling, imperative max-height injection, and an exposed scroll API.

Why it matters: View behavior is tightly coupled to DOM orchestration, which makes slot-list behavior fragile across layout changes.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 Extract scroll and sizing behavior into clearer contracts or composables
- [ ] #2 Remove imperative style injection where practical
- [ ] #3 Preserve current list behavior
<!-- AC:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [ ] #1 All acceptance criteria are satisfied
- [ ] #2 All unit tests pass
- [ ] #3 All e2e tests pass
<!-- DOD:END -->

## Comments

<!-- COMMENTS:BEGIN -->
created: 2026-08-19 14:00
---
**Verification evidence:**
- Extracted the viewport-specific max-height, resize subscription, mount gating, and `scrollToSignUpBlock(id)` behavior into `src/components/sign_up_form/useSignUpBlocksListViewport.ts`, keeping `SignUpBlocksList.vue` focused on rendering and the existing exposed scroll contract.
- Replaced the list root's inline string interpolation with a typed computed style binding while preserving the same desktop max-height behavior and `OverflowGradient` wiring.
- Added `src/components/sign_up_form/SignUpBlocksList.test.ts` coverage for desktop max-height calculation, resize-driven updates, the exposed scroll bridge, missing-target no-op behavior, and desktop-only gradient rendering.
- Ran `npm run test:unit -- SignUpBlocksList.test.ts` in `frontend` on 2026-05-24; 5 tests passed.
- Ran `npm run lint`, `npm run typecheck`, `npm run build`, and `npm run test:unit` in `frontend` on 2026-05-24; all passed.

**Implementation notes:**
- Comparator evidence was not needed because the refactor preserved the existing layout and overflow contracts without changing visual styling.
---
<!-- COMMENTS:END -->
