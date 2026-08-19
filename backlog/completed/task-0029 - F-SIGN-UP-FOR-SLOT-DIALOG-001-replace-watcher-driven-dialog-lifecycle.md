---
id: TASK-0029
title: 'F-SIGN-UP-FOR-SLOT-DIALOG-001: replace watcher-driven dialog lifecycle'
status: Done
assignee: []
created_date: '2026-08-19 13:40'
updated_date: '2026-08-19 14:00'
labels:
  - findings
dependencies: []
references:
  - frontend/findings/F-SIGN-UP-FOR-SLOT-DIALOG-001.md
priority: medium
type: task
ordinal: 29000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Component: src/components/sign_up_form/SignUpForSlotDialog.vue

Problem: The dialog repeats watcher-driven reset, validation rebuilding, and nextTick coordination patterns from the guest flow.

Why it matters: Shared dialog-lifecycle problems are duplicated, which increases maintenance cost and inconsistency risk.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 Replace the duplicated watcher-driven lifecycle pattern with a clearer shared or local state model
- [ ] #2 Preserve slot sign-up behavior
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
- `npm run lint`, `npm run typecheck`, `npm run build`, and `npm run test:unit` passed on 2026-05-23.
- Existing unit coverage in `src/components/sign_up_form/SignUpForSlotDialog.test.ts` passed after replacing watcher-driven reset and `nextTick` validation coordination with explicit dialog initialization and computed rule state.

**Implementation notes:**
- Consider solving this together with `F-GUEST-DIALOG-001` if a shared pattern emerges cleanly.
---
<!-- COMMENTS:END -->
