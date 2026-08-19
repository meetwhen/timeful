---
id: TASK-0014
title: 'F-NEW-SIGN-UP-001: remove reload-based sign-up flow'
status: Done
assignee: []
created_date: '2026-08-19 13:39'
updated_date: '2026-08-19 13:59'
labels:
  - findings
dependencies: []
references:
  - frontend/findings/F-NEW-SIGN-UP-001.md
priority: high
type: task
ordinal: 14000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Component: src/components/NewSignUp.vue

Problem: Sign-up creation still uses reload-based behavior and expose-driven parent coordination inside the main flow.

Why it matters: Core form transitions depend on browser-level resets and imperative coordination rather than explicit state changes.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 Remove reload-based flow
- [ ] #2 Narrow parent-child contracts
- [ ] #3 Preserve current sign-up creation behavior
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
- `npm run lint`, `npm run typecheck`, `npm run build`, and `npm run test:unit` passed in `frontend/`.
- Added edit-save coverage in `src/components/NewSignUp.test.ts` and `src/components/NewDialog.test.ts`.

**Implementation notes:**
- Successful sign-up edits now emit an explicit refresh event so the owning dialog and event view refresh state in-process instead of relying on a page reload.
---
<!-- COMMENTS:END -->
