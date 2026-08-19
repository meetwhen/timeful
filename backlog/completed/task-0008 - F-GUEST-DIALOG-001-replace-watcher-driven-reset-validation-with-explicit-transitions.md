---
id: TASK-0008
title: >-
  F-GUEST-DIALOG-001: replace watcher-driven reset/validation with explicit
  transitions
status: Done
assignee: []
created_date: '2026-08-19 13:39'
updated_date: '2026-08-19 13:58'
labels:
  - findings
dependencies: []
references:
  - frontend/findings/F-GUEST-DIALOG-001.md
priority: medium
type: task
ordinal: 8000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Component: src/components/GuestDialog.vue

Problem: Open-state watchers still reset form state and rebuild validation rules, with nextTick coordination keeping UI in sync.

Why it matters: Dialog lifecycle and form rules are more implicit than they should be, which makes small changes risky.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 Replace watcher-driven reset and validation orchestration with clearer state transitions
- [ ] #2 Preserve guest-flow behavior
<!-- AC:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [ ] #1 All acceptance criteria are satisfied
- [ ] #2 All unit tests pass
- [ ] #3 All e2e tests pass
<!-- DOD:END -->

## Comments

<!-- COMMENTS:BEGIN -->
created: 2026-08-19 13:58
---
**Verification evidence:**
- `npm run lint`, `npm run typecheck`, `npm run build`, and `npm run test:unit` passed on 2026-05-23.
- Existing and updated unit coverage in `src/components/GuestDialog.test.ts` now exercises the dialog with computed rule state instead of watcher-driven reset and `nextTick` validation orchestration.

**Implementation notes:**
- Keep this aligned with the similar sign-up-for-slot dialog pattern to avoid two divergent fixes.
---
<!-- COMMENTS:END -->
