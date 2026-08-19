---
id: TASK-0015
title: 'F-NEW-SIGN-UP-002: split sign-up creator into view/state boundaries'
status: Done
assignee: []
created_date: '2026-08-19 13:39'
updated_date: '2026-08-19 13:59'
labels:
  - findings
dependencies: []
references:
  - frontend/findings/F-NEW-SIGN-UP-002.md
priority: medium
type: task
ordinal: 15000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Component: src/components/NewSignUp.vue

Problem: The sign-up creator still combines validation, reset logic, timers, and parent coordination in one watcher-heavy component.

Why it matters: Refactoring or extending sign-up behavior remains risky because the file owns too many concerns.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 Decompose the flow into clearer view and state boundaries
- [ ] #2 Preserve sign-up behavior
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
- Moved shared sign-up editor ownership into `src/composables/event/useEventEditorState.ts` and reused the shared schedule builder in `src/composables/event/eventEditorSchedule.ts`.
- Preserved `NewSignUp.vue` public props, emits, and exposed methods while removing duplicate watcher-heavy reset and hydration logic from the component.
- Added focused regression coverage in `src/composables/event/useEventEditorState.test.ts` and `src/composables/event/eventEditorSchedule.test.ts`.
- Passed `npm run lint`, `npm run typecheck`, `npm run build`, and `npm run test:unit`.

**Implementation notes:**
- Handle reload-based flow first under `F-NEW-SIGN-UP-001` so the later split is simpler.
---
<!-- COMMENTS:END -->
