---
id: TASK-0033
title: 'F-TOOLTIP-001: centralize tooltip placement and visibility logic'
status: Done
assignee: []
created_date: '2026-08-19 13:40'
updated_date: '2026-08-19 14:01'
labels:
  - findings
dependencies: []
references:
  - frontend/findings/F-TOOLTIP-001.md
priority: high
type: task
ordinal: 33000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Component: src/components/Tooltip.vue

Problem: Tooltip visibility and placement still depend on manual DOM listeners, inline style objects, timers, and mount-coupled setup.

Why it matters: Behavior is difficult to verify and easy to regress because state ownership is split across reactivity and direct DOM orchestration.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 Centralize placement and visibility logic behind a composable or helper
- [ ] #2 Remove manual listener wiring from the view where practical
- [ ] #3 Preserve current tooltip behavior
<!-- AC:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [ ] #1 All acceptance criteria are satisfied
- [ ] #2 All unit tests pass
- [ ] #3 All e2e tests pass
<!-- DOD:END -->

## Comments

<!-- COMMENTS:BEGIN -->
created: 2026-08-19 14:01
---
**Verification evidence:**
- Added `src/composables/useTooltipState.ts` to centralize delayed visibility, pointer placement, and tooltip style derivation.
- Updated `src/components/Tooltip.vue` to use declarative template mouse events instead of manual listener registration in lifecycle hooks.
- Added `src/components/Tooltip.test.ts` covering the listener boundary, delayed visibility, and pointer-driven placement output.
- Ran `npm run lint`, `npm run typecheck`, `npm run build`, and `npm run test:unit`.

**Implementation notes:**
- Favor targeted unit coverage for placement and visibility logic before touching layout glue.
---
<!-- COMMENTS:END -->
