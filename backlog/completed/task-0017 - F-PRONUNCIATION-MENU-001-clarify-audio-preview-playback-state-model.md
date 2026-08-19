---
id: TASK-0017
title: 'F-PRONUNCIATION-MENU-001: clarify audio preview playback state model'
status: Done
assignee: []
created_date: '2026-08-19 13:39'
updated_date: '2026-08-19 13:59'
labels:
  - findings
dependencies: []
references:
  - frontend/findings/F-PRONUNCIATION-MENU-001.md
priority: low
type: task
ordinal: 17000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Component: src/components/PronunciationMenu.vue

Problem: Audio preview still uses timers and imperative media-element control through a template ref.

Why it matters: Media behavior is more DOM-driven than necessary, which makes the component awkward to reason about.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 Clarify the playback state model
- [ ] #2 Reduce direct element control where practical
- [ ] #3 Preserve current preview behavior
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
- `npm run lint`, `npm run typecheck`, `npm run build`, and `npm run test:unit` passed on 2026-05-24.
- Added `src/composables/usePronunciationMenuPlayback.test.ts` to cover animation progression, close/reset behavior, unmount cleanup, and blocked-audio playback handling.

**Implementation notes:**
- Playback state, timer cleanup, and audio reset behavior now live in `src/composables/usePronunciationMenuPlayback.ts`, while the component keeps the existing menu template and assets.
---
<!-- COMMENTS:END -->
