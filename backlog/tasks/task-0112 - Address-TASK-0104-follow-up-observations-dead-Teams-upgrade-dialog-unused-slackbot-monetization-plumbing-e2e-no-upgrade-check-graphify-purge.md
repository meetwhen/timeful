---
id: TASK-0112
title: >-
  Address TASK-0104 follow-up observations: dead Teams upgrade dialog, unused
  slackbot monetization plumbing, e2e no-upgrade check, graphify purge
status: Done
assignee:
  - opencode
created_date: '2026-08-29 21:29'
updated_date: '2026-08-29 22:18'
labels:
  - cleanup
  - frontend
  - backend
  - config
dependencies: []
references:
  - >-
    backlog/tasks/task-0104 -
    Remove-freemium-gating-Stripe-billing-and-advertising-from-the-app.md
  - >-
    backlog/tasks/task-0113 -
    Add-e2e-regression-spec-asserting-no-upgrade-paywall-or-ad-surfaces-remain-in-the-app.md
type: chore
ordinal: 118000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Follow-up to TASK-0104 (Remove freemium gating, Stripe billing, and advertising). TASK-0104's final summary recorded out-of-scope observations; the owner reviewed them on 2026-08-30 and approved addressing them as one focused change.

Binding owner decisions (2026-08-30):

- Remove TeamsNotReadyDialog entirely. It is an org-plan "Upgrade Required" book-a-call sales CTA and currently dead code; delete the component, its wiring in AuthUserMenu.vue (mount, ref, import), the commented-out add-team-member menu block, and its test stub.
- Remove the unused slackbot monetization plumbing: the MONETIZATION message type and its case in server/slackbot/slackbot.go, and the SLACK_MONETIZATION_WEBHOOK_URL env var from compose.yaml, all env files, and docs/environments.md.
- The e2e regression spec asserting no upgrade/paywall/ad surfaces was postponed to TASK-0113 by owner decision.
- Attempt to purge stale freemium nodes from graphify-out; if not possible without an API key, document the concrete blocker in the task notes.
- Restore FR-070/071/072 as rejected and ADR-003 as deprecated per the status conventions (owner review of TASK-0104); the Freemium Mode glossary term stays removed.

Constraints:

- Keep monitoring analytics per TASK-0104 scope: PostHog, GTM, numEventsCreated counter, /analytics/monthly-active-event-creators*, /analytics/user/:email, /analytics/scanned-poster.
- Do not reintroduce any freemium, billing, or advertising plumbing.
- Do not touch unrelated in-flight work present in the worktree (e.g. TASK-0107 changes during the scoping session); re-check git status at execution time.
- This is a planned follow-up: when starting execution, read BACKLOG_WORKFLOW.md, mark the task In Progress, assign it, and keep plan/notes current.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [x] #1 TeamsNotReadyDialog.vue is deleted; AuthUserMenu.vue has no dialog mount, ref, import, or commented add-team-member block; the TeamsNotReadyDialog stub is removed from AuthUserMenu.test.ts.
- [x] #2 The slackbot MONETIZATION type and case are removed from server/slackbot/slackbot.go and SLACK_MONETIZATION_WEBHOOK_URL is removed from compose.yaml, all env files, and docs/environments.md.
- [x] #3 graphify-out no longer contains freemium nodes from deleted files, or the purge attempt and its concrete blocker are recorded in the task notes.
- [x] #4 Required checks pass: frontend lint, typecheck, build, test:unit; server go build, go vet, and the scoped Mongo route suite via the compose.test.yaml overlay.
- [x] #5 Requirements records FR-070, FR-071, FR-072 are restored with status rejected and ADR-003 is restored with status deprecated (Freemium Mode glossary term stays removed and mentions are unlinked); their index rows are restored in docs/requirements/README.md and docs/design/README.md, and the pre-existing malformed FR-index header block is fixed.
<!-- AC:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [x] #1 All acceptance criteria are satisfied
- [x] #2 All required unit tests pass. Documentation-only changes are exempt unless the user requests unit tests
- [x] #3 All required e2e tests pass. Documentation-only changes are exempt unless the user requests e2e tests
- [x] #4 Changed Markdown files are formatted with npm run format:markdown
<!-- DOD:END -->

## Implementation Plan

<!-- SECTION:PLAN:BEGIN -->
# Suggested implementation plan (for the executing session)

## 1. Frontend — TeamsNotReadyDialog removal
- git rm frontend/src/components/TeamsNotReadyDialog.vue
- frontend/src/components/AuthUserMenu.vue: remove the TeamsNotReadyDialog import, the `<TeamsNotReadyDialog>` mount, the `showTeamsNotReadyDialog` ref, and the commented-out add-team-member menu block
- frontend/src/components/AuthUserMenu.test.ts: remove the `TeamsNotReadyDialog: true` stub

## 2. Server — slackbot monetization removal
- server/slackbot/slackbot.go: remove the MONETIZATION constant and its webhook case
- go build ./... and go vet ./... from server/

## 3. Env/config/docs
- Remove SLACK_MONETIZATION_WEBHOOK_URL from compose.yaml, .env.development, .env.development.example, .env.test, .env.test.example, .env.staging.example, .env.production.example, and docs/environments.md
- Run npm run format:markdown on changed Markdown files

## 4. Docs repair (TASK-0104 owner review)
- Restore FR-070.md, FR-071.md, FR-072.md from git history with status rejected and a one-line rejection note; write Freemium Mode mentions as plain text (glossary term stays removed)
- Restore ADR-003.md with status deprecated, updated_date 2026-08-30, and a deprecation note; no superseding ADR
- Restore their index rows in docs/requirements/README.md and docs/design/README.md
- Fix the pre-existing malformed FR-index header block in docs/requirements/README.md

## 5. Graphify purge
- Run `graphify update . --force`; verify graph.json no longer matches freemium
- If orphaned semantic-cache nodes persist, remove the stale cache entries and re-run; otherwise record the concrete blocker in the task notes

## 6. Verification and finalization
- frontend: npm run lint, typecheck, build, test:unit
- server: go build ./..., go vet ./...; scoped Mongo route suite via the compose.test.yaml overlay
- Sanity greps: Upgrade Required/book_call_for_organization_plan_clicked/TeamsNotReadyDialog/MONETIZATION/SLACK_MONETIZATION return zero over frontend/src + server + config
- Finalize per BACKLOG_WORKFLOW.md: verify each AC with evidence, check AC/DoD boxes, write final summary, set status Done
<!-- SECTION:PLAN:END -->

## Implementation Notes

<!-- SECTION:NOTES:BEGIN -->
Research findings from the 2026-08-30 scoping session (verified against the worktree):

- TeamsNotReadyDialog: mounted at frontend/src/components/AuthUserMenu.vue:58 with a `showTeamsNotReadyDialog` ref (line 80) that is never set to true; the only historical trigger is the commented-out "Add team member" menu item (lines 24-31). Component lives at frontend/src/components/TeamsNotReadyDialog.vue and fires the `book_call_for_organization_plan_clicked` PostHog event plus a cal.com book-a-call link. Test stub `TeamsNotReadyDialog: true` sits in frontend/src/components/AuthUserMenu.test.ts around line 106.
- Slackbot: MONETIZATION is defined at server/slackbot/slackbot.go:19 with its webhook case at lines 41-43. Repo-wide grep finds no senders — all SendMessage/SendMessageWithType call sites use GENERAL. SLACK_MONETIZATION_WEBHOOK_URL appears in compose.yaml:139, .env.development:66, .env.development.example:66, .env.test:65, .env.test.example:66, .env.staging.example:70, .env.production.example:70, and docs/environments.md:164.
- E2E: there is no authenticated e2e infrastructure today. Auth is OTP-based (server/routes/auth.go check-email/send/verify); OTP codes live only in Mongo OtpCodesCollection and Listmonk emails with no test hook, so real signed-in flows cannot be scripted cheaply. The firefox-desktop Playwright project matches only timed-event-*firefox.spec.ts, so a new generic spec runs on the default desktop (chromium) project. Feasible approach: assert public pages (landing, sign-in) directly, and render the signed-in shell by intercepting /api/auth/status and user API calls via page.route (route stubbing precedent exists in landing-hero.spec.ts). Router guard redirects /home and /settings to landing when /auth/status fails (frontend/src/router/index.ts:104), so the interception must cover it.
- Graphify: graphify-out/graph.json still contains freemium matches (including "freemium operational gating; passed through Compose build args") from deleted files, plus orphaned entries under graphify-out/cache/semantic/ keyed by content hashes of deleted sources. No LLM API key is set in the environment. Deterministic purge path: `graphify update . --force` (AST re-extraction, no LLM, overwrites even when node count drops); if orphaned semantic-cache nodes survive, delete the stale cache entries and re-run, otherwise document the blocker.

2026-08-30 owner review of TASK-0104 expanded scope: restore FR-070/071/072 as rejected and ADR-003 as deprecated per the requirements and ADR status conventions (permanent IDs; deletion is not a lifecycle status). Keep the glossary Freemium Mode entry removed and write Freemium Mode mentions in restored records as plain text. No superseding ADR. Also fix the pre-existing malformed FR-index header block in docs/requirements/README.md.

2026-08-30 owner decision: postpone the e2e no-upgrade/paywall/ad regression spec out of this task; it is now tracked as TASK-0113 with the e2e research findings carried over.
<!-- SECTION:NOTES:END -->

## Final Summary

<!-- SECTION:FINAL_SUMMARY:BEGIN -->
Removed the remaining freemium-adjacent dead code and repaired the TASK-0104 documentation records.

- Frontend: deleted `TeamsNotReadyDialog.vue` and all wiring in `AuthUserMenu.vue` (import, mount, `showTeamsNotReadyDialog` ref, commented add-team-member block) plus its test stub in `AuthUserMenu.test.ts`.
- Server: removed the unused `MONETIZATION` message type and webhook case from `server/slackbot/slackbot.go`.
- Config: removed `SLACK_MONETIZATION_WEBHOOK_URL` from `compose.yaml`, all env files, and `docs/environments.md`.
- Docs: restored FR-070/071/072 with `status: rejected` and rejection notes, restored ADR-003 with `status: deprecated` and a deprecation note (no superseding ADR), kept Freemium Mode mentions as unlinked plain text, restored their index rows in `docs/requirements/README.md` and `docs/design/README.md`, and fixed the pre-existing malformed duplicated FR-index header block.
- Graphify: purged stale nodes; `graphify-out/graph.json` now contains no nodes from deleted files. The only remaining freemium matches are titles/sections of the restored rejected/deprecated records, which is correct. Semantic-cache entries that mention freemium all reference existing restored files; no API-key blocker was hit.

**Verification evidence (2026-08-30):**
- Sanity greps return zero for `TeamsNotReadyDialog`, `MONETIZATION`, `SLACK_MONETIZATION` across frontend/src, server, compose.yaml, all env files, and docs/environments.md; no freemium glossary links remain in restored records.
- Frontend: `npm run lint`, `npm run typecheck`, `npm run build`, and `npm run test:unit` all pass (137 files, 952 tests).
- Server: `go build ./...` and `go vet ./...` pass; scoped Mongo route suite passes via `docker compose --env-file .env.test -f compose.yaml -f compose.test.yaml run --rm server-route-test` (`ok timeful/server/routes`).
- `npm run format:markdown:check` passes for all changed Markdown files.
- E2e DoD item: exempted by owner instruction on 2026-08-30 — the e2e no-upgrade/paywall/ad regression spec is tracked as TASK-0113, so e2e was skipped for this task.
<!-- SECTION:FINAL_SUMMARY:END -->
