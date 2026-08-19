---
id: TASK-0032
title: 'F-TIMEZONE-SELECTOR-001: move timezone persistence behind owned boundary'
status: Done
assignee: []
created_date: '2026-08-19 13:40'
updated_date: '2026-08-19 14:00'
labels:
  - findings
dependencies: []
references:
  - frontend/findings/F-TIMEZONE-SELECTOR-001.md
priority: high
type: task
ordinal: 32000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Component: src/components/schedule_overlap/TimezoneSelector.vue

Problem: The control mirrors props into local state and persists selection through localStorage from inside the selector.

Why it matters: The selector currently owns both input state and persistence side effects, which blurs the boundary between view logic and application state.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 Keep one clear owner for timezone state
- [ ] #2 Move persistence behind an explicit boundary
- [ ] #3 Preserve existing user-visible behavior
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
- `npm run lint`, `npm run typecheck`, `npm run build`, and `npm run test:unit` passed in `frontend/`.
- Added `src/composables/timezone/useOwnedTimezone.test.ts` and updated `src/components/schedule_overlap/TimezoneSelector.test.ts`.

**Implementation notes:**
- `TimezoneSelector` is now a controlled view-only component that emits `update:modelValue` and `reset`, while owner components handle persistence and reset semantics through `useOwnedTimezone`.
---
<!-- COMMENTS:END -->
