---
id: TASK-0004
title: >-
  F-COOKIE-CONSENT-001: centralize consent persistence and remove reload-based
  flow
status: Done
assignee: []
created_date: '2026-08-19 13:39'
updated_date: '2026-08-19 13:58'
labels:
  - findings
dependencies: []
references:
  - frontend/findings/F-COOKIE-CONSENT-001.md
priority: high
type: task
ordinal: 4000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Component: src/components/CookieConsent.vue

Problem: Consent bootstrapping runs directly from browser storage, and accept-flow behavior still uses window.location.reload().

Why it matters: Consent state should cross an explicit boundary instead of being coordinated through ad hoc browser-side effects.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 Centralize consent persistence and state transitions behind a clear boundary
- [ ] #2 Remove reload-based control flow
- [ ] #3 Preserve current consent behavior
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
- Reconfirmed that `src/components/CookieConsent.vue` read `localStorage` directly during setup and called `window.location.reload()` after persisting consent before this change.
- Centralized consent reads, normalized preference defaults, and write notifications in `src/utils/cookie_utils.ts`, then updated `CookieConsent.vue` to hydrate and react through that shared boundary instead of direct browser-side effects.
- Added focused regression coverage in `src/components/CookieConsent.test.ts` and `src/utils/cookie_utils.test.ts` for corrupt or missing consent, shared consent updates, persisted preference normalization, and reload removal.
- Ran `npm run lint`, `npm run typecheck`, `npm run build`, and `npm run test:unit` in `frontend/` on 2026-05-24; all checks passed.

**Implementation notes:**
- Preserved the current default consent preference shape and kept live GTM/PostHog reconfiguration out of scope for this ownership fix.
---
<!-- COMMENTS:END -->
