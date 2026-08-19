---
id: TASK-0006
title: 'F-DASHBOARD-001: move folder-open persistence behind explicit boundary'
status: Done
assignee: []
created_date: '2026-08-19 13:39'
updated_date: '2026-08-19 13:58'
labels:
  - findings
dependencies: []
references:
  - frontend/findings/F-DASHBOARD-001.md
priority: low
type: task
ordinal: 6000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Component: src/components/home/Dashboard.vue

Problem: Folder open state is bootstrapped from and persisted back to localStorage through setup logic and watchers inside the component.

Why it matters: Ambient browser state currently owns a cross-session dashboard behavior that should have a clearer boundary.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 Move folder-state persistence behind an explicit boundary or composable
- [ ] #2 Preserve existing dashboard behavior
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
- Reconfirmed that `src/components/home/Dashboard.vue` directly read and wrote `localStorage` in setup through `folderOpenState` initialization plus deep watchers before this change.
- Moved folder-open persistence behind `src/components/home/useDashboardFolderOpenState.ts`.
- Added focused unit coverage in `src/components/home/useDashboardFolderOpenState.test.ts` for restoration, persistence, and default-open handling for newly seen folders.
- Ran `npm run test:unit -- src/components/home/Dashboard.test.ts src/components/home/useDashboardFolderOpenState.test.ts` in `frontend` on 2026-05-23; 4 tests passed.

**Implementation notes:**
- Kept dashboard toggle behavior and the default-open behavior for newly discovered folders while moving browser-storage ownership behind a dashboard-local composable boundary.
---
<!-- COMMENTS:END -->
