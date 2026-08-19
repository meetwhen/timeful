---
id: TASK-0036
title: 'F-USER-ITEM-001: move showEventNames persistence behind explicit boundary'
status: Done
assignee: []
created_date: '2026-08-19 13:40'
updated_date: '2026-08-19 14:01'
labels:
  - findings
dependencies: []
references:
  - frontend/findings/F-USER-ITEM-001.md
priority: high
type: task
ordinal: 36000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Component: src/components/UserItem.vue

Problem: showEventNames is initialized from localStorage, then a watcher persists and emits from the same reactive path.

Why it matters: Browser storage currently owns user-facing state that should have a clearer boundary, which makes behavior harder to test and reason about.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 Move storage reads and writes behind an explicit boundary or composable
- [ ] #2 Keep the component focused on rendering and user interaction
- [ ] #3 Preserve current behavior
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
- Added `src/composables/useShowEventNamesPreference.ts` so `UserItem.vue` no longer reads or writes `localStorage` directly.
- Added `src/components/UserItem.test.ts` covering stored initialization, persisted updates, and emitted change behavior.
- Ran `npm run lint`, `npm run typecheck`, `npm run build`, and `npm run test:unit`.

**Implementation notes:**
- Confirm current persistence behavior first and add targeted regression coverage around initialization and updates if practical.
---
<!-- COMMENTS:END -->
