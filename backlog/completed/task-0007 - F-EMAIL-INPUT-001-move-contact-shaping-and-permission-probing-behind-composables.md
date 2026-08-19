---
id: TASK-0007
title: >-
  F-EMAIL-INPUT-001: move contact shaping and permission probing behind
  composables
status: Done
assignee: []
created_date: '2026-08-19 13:39'
updated_date: '2026-08-19 13:58'
labels:
  - findings
dependencies: []
references:
  - frontend/findings/F-EMAIL-INPUT-001.md
priority: high
type: task
ordinal: 7000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Component: src/components/event/EmailInput.vue

Problem: The component mixes Contact | string local state, mutates fetched contact objects, probes permissions on mount, and exposes an imperative reset() API.

Why it matters: Transport shaping and async side effects currently live inside the view, which conflicts with the frontend boundary-model rules.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 Move contact shaping and permission probing behind explicit helpers or composables
- [ ] #2 Keep the component contract typed
- [ ] #3 Remove the imperative reset surface if possible
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
- Added unit coverage in `src/components/event/EmailInput.test.ts` for parent-driven resync while the component now uses typed internal entries, extracted contact helpers/composables, and no exposed `reset()` API.

**Implementation notes:**
- Read ADR 001 before implementation and add unit coverage around any extracted contact-shaping logic.
---
<!-- COMMENTS:END -->
