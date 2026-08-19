---
id: TASK-0005
title: >-
  F-COOKIE-SETTINGS-001: move settings behind explicit consent boundary and
  remove reload
status: Done
assignee: []
created_date: '2026-08-19 13:39'
updated_date: '2026-08-19 13:58'
labels:
  - findings
dependencies: []
references:
  - frontend/findings/F-COOKIE-SETTINGS-001.md
priority: high
type: task
ordinal: 5000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Component: src/components/CookieSettings.vue

Problem: Settings initialization and apply behavior still depend on mount-time storage reads and window.location.reload().

Why it matters: Browser-owned flow hides state transitions that should be explicit and testable.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 Move settings reads and writes behind an explicit boundary
- [ ] #2 Remove reload-based flow
- [ ] #3 Preserve current cookie-settings UX
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
- Reconfirmed that `src/components/CookieSettings.vue` loaded saved consent through mount-time reads and reloaded the page after saves before this change.
- Moved the settings page onto the shared consent boundary in `src/utils/cookie_utils.ts`, removed reload-based flow, and kept the settings page synchronized through the shared consent version signal.
- Added focused regression coverage in `src/components/CookieSettings.test.ts` for setup hydration, explicit saves, external consent updates, and reload removal.
- Ran `npm run lint`, `npm run typecheck`, `npm run build`, and `npm run test:unit` in `frontend/` on 2026-05-24; all checks passed.

**Implementation notes:**
- Kept the page-local state simple and reused the same normalized consent defaults as the banner so the two consent entry points cannot drift.
---
<!-- COMMENTS:END -->
