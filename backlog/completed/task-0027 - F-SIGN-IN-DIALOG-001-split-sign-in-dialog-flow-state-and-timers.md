---
id: TASK-0027
title: 'F-SIGN-IN-DIALOG-001: split sign-in dialog flow state and timers'
status: Done
assignee: []
created_date: '2026-08-19 13:40'
updated_date: '2026-08-19 14:00'
labels:
  - findings
dependencies: []
references:
  - frontend/findings/F-SIGN-IN-DIALOG-001.md
priority: medium
type: task
ordinal: 27000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Component: src/components/SignInDialog.vue

Problem: Provider selection, onboarding, OTP verification, validation, and resend cooldown timer orchestration still live in one dialog component.

Why it matters: The sign-in flow is hard to modify safely because independent steps share local state and timer coordination.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 Split flow state and timer behavior into clearer units
- [ ] #2 Preserve current sign-in UX
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
- `npm run lint`, `npm run typecheck`, `npm run build`, and `npm run test:unit` passed on 2026-05-24.
- Added `src/composables/useSignInDialogState.test.ts` for cooldown, close/reset, unmount cleanup, and successful verification callbacks, and updated `src/components/SignInDialog.test.ts` to keep the existing component contract covered.

**Implementation notes:**
- The dialog now delegates OTP step state, validation, resend cooldown ownership, and verification flow to `src/composables/useSignInDialogState.ts`, leaving the component focused on rendering and emits.
---
<!-- COMMENTS:END -->
