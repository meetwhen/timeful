---
id: TASK-0012
title: 'F-NEW-EVENT-002: split event editor along shared composables'
status: Done
assignee: []
created_date: '2026-08-19 13:39'
updated_date: '2026-08-19 13:59'
labels:
  - findings
dependencies: []
references:
  - frontend/findings/F-NEW-EVENT-002.md
priority: medium
type: task
ordinal: 12000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Component: src/components/NewEvent.vue

Problem: The event editor still mixes many watchers, refs, computed values, and child coordination responsibilities in one component.

Why it matters: The file is too broad to maintain confidently, especially while migration parity work is still active.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 Split the editor along stable subflows or composables
- [ ] #2 Preserve behavior
- [ ] #3 Keep extracted logic covered with targeted tests
<!-- AC:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [ ] #1 All acceptance criteria are satisfied
- [ ] #2 All unit tests pass
- [ ] #3 All e2e tests pass
<!-- DOD:END -->

## Comments

<!-- COMMENTS:BEGIN -->
created: 2026-08-19 13:59
---
**Verification evidence:**
- Extracted shared editor state ownership into `src/composables/event/useEventEditorState.ts`.
- Extracted shared date/time schedule shaping into `src/composables/event/eventEditorSchedule.ts`.
- Preserved `NewEvent.vue` public props, emits, and exposed methods while moving reset, hydration, and edit-tracking logic behind the shared composable.
- Added focused regression coverage in `src/composables/event/useEventEditorState.test.ts` and `src/composables/event/eventEditorSchedule.test.ts`.
- Passed `npm run lint`, `npm run typecheck`, `npm run build`, and `npm run test:unit`.

**Implementation notes:**
- Resolve reload and storage ownership issues from `F-NEW-EVENT-001` before optional structural cleanup.
---
<!-- COMMENTS:END -->
