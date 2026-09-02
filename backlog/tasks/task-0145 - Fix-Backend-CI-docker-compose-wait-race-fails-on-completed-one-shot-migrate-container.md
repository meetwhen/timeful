---
id: TASK-0145
title: >-
  Fix Backend CI: docker compose wait race fails on completed one-shot migrate
  container
status: Done
assignee: []
created_date: '2026-09-02 17:01'
updated_date: '2026-09-02 17:12'
labels:
  - ci
  - docker-compose
dependencies: []
references:
  - 'https://github.com/whensync/timeful/actions/runs/33316150651/job/99269864485'
modified_files:
  - .github/workflows/backend-ci.yml
  - docs/environments.md
priority: high
type: bug
ordinal: 158300
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
The Backend CI workflow fails on every run at the "Wait for PostgreSQL migrations" step with `no containers for project "timeful-test"`.

Root cause (verified against docker/compose source v2.29.7 through main):

1. `docker compose up -d --build ... postgres-test-migrate` starts the one-shot migrate container and returns immediately after "Started"; detached `up` only blocks on dependency conditions (bootstrap's `service_completed_successfully`), not on the requested service's own exit. CI log confirms `timeful-test-postgres-test-migrate-1 Started` at 14:12:11.622 and the up step ending at 14:12:11.631.
2. `docker compose wait postgres-test-migrate` errors with `no containers for project "timeful-test"` even though the container exists. Compose's `wait` implementation (`pkg/compose/wait.go`) lists containers with `ContainerList{All: false}`, i.e. running containers only. If the target already exited (exit 0 or 1), the list is empty and the command fails.
3. The migrate container only runs goose over two small SQL migration files plus a psql grant, so it finishes in roughly 10-150 ms on the CI runner, always before the wait step queries (the 2026-08-29 run had a 9 ms gap between "Started" and the wait step starting).
4. Local E2E never hits this because `frontend/e2e/isolated-test-stack.ts` never uses `wait`; it starts the stack and polls the API healthcheck, and `server-test` depends on `postgres-test-migrate: service_completed_successfully` inside Compose.

Additional defect: the same behavior masks real migration failures, because a failed goose run produces the identical "no containers" error instead of the true exit code.

Fix direction: remove the redundant wait step. `docker compose run --rm server-route-test` already honors `depends_on: postgres-test-migrate: service_completed_successfully` (compose `run` calls `waitDependencies`, `pkg/compose/run.go`), so migrations are enforced at the point of consumption and a failed migration surfaces as a dependency failure. Add a failure-conditioned log dump for the one-shot services and update the documented route-test snippet.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [x] #1 backend-ci.yml no longer runs docker compose wait postgres-test-migrate
- [x] #2 Migration completion is still enforced before route tests via server-route-test depends_on postgres-test-migrate service_completed_successfully
- [x] #3 A failure-conditioned step dumps postgres-test-migrate and postgres-test-bootstrap logs so goose failures surface
- [x] #4 docs/environments.md route-test snippet no longer documents the racy wait command
- [x] #5 Running the CI sequence locally (compose up mongo-test postgres-test postgres-test-bootstrap postgres-test-migrate, then compose run --rm server-route-test go test ./...) passes against the isolated test stack
<!-- AC:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [x] #1 All acceptance criteria are satisfied
- [x] #2 All required unit tests pass. Documentation-only changes are exempt unless the user requests unit tests
- [ ] #3 All required e2e tests pass. Documentation-only changes are exempt unless the user requests e2e tests
- [x] #4 Changed Markdown files are formatted with npm run format:markdown
<!-- DOD:END -->

## Implementation Plan

<!-- SECTION:PLAN:BEGIN -->
Implementation plan:
1. `.github/workflows/backend-ci.yml`: delete the "Wait for PostgreSQL migrations" step (docker compose ... wait postgres-test-migrate). Keep the up step unchanged (mongo-test, postgres-test, postgres-test-bootstrap, postgres-test-migrate).
2. `.github/workflows/backend-ci.yml`: add an `if: failure()` step after "Run backend tests" that dumps `docker compose ... logs postgres-test-migrate postgres-test-bootstrap` so goose/bootstrap failures surface with real output.
3. `docs/environments.md` Route tests snippet: remove the `docker compose ... wait postgres-test-migrate` line; keep the up and run lines. Keep each sentence on one physical source line.
4. Verify locally with the isolated test stack: `docker compose --env-file .env.test -f compose.yaml -f compose.test.yaml config --quiet`, then run the exact CI sequence (up -d --build mongo-test postgres-test postgres-test-bootstrap postgres-test-migrate; run --rm server-route-test go test ./... -count=1), then down -v.
5. Verify acceptance criteria, record final summary, mark task Done.
<!-- SECTION:PLAN:END -->

## Implementation Notes

<!-- SECTION:NOTES:BEGIN -->
Execution notes:
- CI run 33316150651 job 99269864485: up step succeeded (postgres-test-migrate Started 14:12:11.6229), wait step failed 14:12:11.7659 with `no containers for project "timeful-test"`.
- compose Wait() source: pkg/compose/wait.go uses getContainers(..., oneOffInclude, false /* all */, services...) so exited containers are invisible; identical behavior in v2.29.7, v2.39.3, and main.
- compose run: pkg/compose/run.go calls waitDependencies (run.go:174), which honors depends_on conditions including service_completed_successfully.
- Migration set: server/migrations contains 2 SQL files (20260814170000, 20260815100000), so migrate completes in well under a second on CI.
- Every Backend CI run (2026-08-25, 26, 29, 30) failed at the wait step; local E2E bypasses wait entirely via isolated-test-stack.ts healthcheck polling.
<!-- SECTION:NOTES:END -->

## Final Summary

<!-- SECTION:FINAL_SUMMARY:BEGIN -->
## Problem

Every Backend CI run failed at "Wait for PostgreSQL migrations" with `no containers for project "timeful-test"` (e.g. run 33316150651).

## Root cause

`docker compose up -d` returns immediately after starting the one-shot `postgres-test-migrate` container (it only blocks on dependency conditions, such as bootstrap's `service_completed_successfully`). Compose's `wait` implementation (`pkg/compose/wait.go`, identical in v2.29.7 through main) lists containers with `ContainerList{All: false}`, i.e. running containers only, and errors when the list is empty. The migrate container runs goose over two small migration files plus a psql grant and finishes in well under a second, so it had always already exited when the wait step queried it. The same behavior masks real migration failures: a failed goose run produces the identical "no containers" error instead of the true exit code. Local E2E never hit this because `frontend/e2e/isolated-test-stack.ts` never uses `wait`; it polls the API healthcheck and lets Compose enforce `depends_on` for `server-test`.

## Fix

- `.github/workflows/backend-ci.yml`: removed the racy "Wait for PostgreSQL migrations" step. Migration completion remains enforced by `server-route-test`'s `depends_on: postgres-test-migrate: service_completed_successfully`, which `docker compose run` honors via `waitDependencies` (`pkg/compose/run.go`). Added an `if: failure()` step that dumps `postgres-test-bootstrap`/`postgres-test-migrate` logs so goose failures surface with real output.
- `docs/environments.md`: removed the `wait` line from the route-test snippet and documented why the command must not be used as a gate.

## Verification (all against the isolated test stack, then `down -v`)

- Reproduced the bug locally: with migrate exited 0, `docker compose wait postgres-test-migrate` fails with the exact CI error `no containers for project "timeful-test"`.
- Happy path, exact CI sequence: `up -d --build mongo-test postgres-test postgres-test-bootstrap postgres-test-migrate` then `run --rm server-route-test go test ./... -count=1` — full suite passed.
- Failure path: with a broken migrator password, `compose run server-route-test` exits 1 with `service "postgres-test-migrate" didn't complete successfully: exit 1` (a real failure signal), and the new log-dump command surfaces goose's actual error (`FATAL: password authentication failed for user "timeful_migrator"`).
- `docker compose config --quiet` passes; `npm run format:markdown:check` passes.

## Not run

The scoped browser E2E spec (`timed-event-postgres-plugin-firefox.spec.ts`) was not run locally: the change does not touch the E2E startup path (`isolated-test-stack.ts` never used `wait`) and CI runs it on push.
<!-- SECTION:FINAL_SUMMARY:END -->
