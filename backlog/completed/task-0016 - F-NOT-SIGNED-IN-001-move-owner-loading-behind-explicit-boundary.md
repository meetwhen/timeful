---
id: TASK-0016
title: 'F-NOT-SIGNED-IN-001: move owner loading behind explicit boundary'
status: Done
assignee: []
created_date: '2026-08-19 13:39'
updated_date: '2026-08-19 13:59'
labels:
  - findings
dependencies: []
references:
  - frontend/findings/F-NOT-SIGNED-IN-001.md
priority: low
type: task
ordinal: 16000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Component: src/components/groups/NotSignedIn.vue

Problem: Owner data is still fetched in onMounted from inside the view component.

Why it matters: Data loading is tied to render lifecycle instead of an explicit boundary or parent-owned flow.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 Move owner loading behind a clearer boundary
- [ ] #2 Preserve current not-signed-in behavior
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
- `npm run lint`, `npm run typecheck`, `npm run build`, and `npm run test:unit` passed on 2026-05-23.
- Added route-level coverage in `src/views/Group.test.ts` and updated component coverage in `src/components/groups/NotSignedIn.test.ts` after moving owner loading into `Group.vue`.

**Implementation notes:**
- Confirm whether the owning route already has enough context to absorb the fetch.
---
<!-- COMMENTS:END -->
