---
id: TASK-0145
title: >-
  Fix Backend CI: docker compose wait race fails on completed one-shot migrate
  container
status: Done
assignee: []
created_date: '2026-09-02 17:01'
updated_date: '2026-09-03 14:46'
labels:
  - ci
  - docker-compose
dependencies: []
references:
  - 'https://github.com/whensync/timeful/actions/runs/33316150651/job/99269864485'
modified_files:
  - .github/workflows/backend-ci.yml
  - docs/environments.md
  - AGENTS.md
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
- [x] #6 backend-ci.yml ensures the external timeful-test-go-build-cache volume exists before any compose step that mounts it (server-route-test), matching the E2E harness behavior
- [x] #7 docs/environments.md route-test snippet includes creating the external Go build cache volume so a fresh machine does not hit the missing-volume error
- [x] #8 Local verification covers the fresh-runner condition: with the external volume removed, the CI run command reproduces the external volume error, and after creating it the exact CI sequence (up, then run --rm server-route-test go test ./... -count=1) passes
- [x] #9 Backend CI green on the next run for the run-tests step
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
Implementation plan (round 2 - reopened for external volume failure):

1. `.github/workflows/backend-ci.yml`: add an "Ensure external Go build cache volume exists" step (`docker volume create timeful-test-go-build-cache`, idempotent) before "Start MongoDB and PostgreSQL test services", mirroring `ensureGoBuildCacheVolume` in `frontend/e2e/isolated-test-stack.ts`. A comment should note why: compose never creates external volumes and fresh runners have none.
2. `docs/environments.md` Route tests snippet: add the `docker volume create timeful-test-go-build-cache` line before the compose up line, and extend the build-cache paragraph to say the volume must be created before first use. Keep each sentence on one physical source line.
3. `AGENTS.md` Server Test Workflow: add one line noting the external build-cache volume must exist (create/reset commands) so agent-driven fresh setups succeed. (Pending user confirmation in summary; part of the same defect class.)
4. Verify the fresh-runner condition locally: `docker volume rm timeful-test-go-build-cache`, run the CI up step (no service mounts the volume, so it still succeeds), run the CI run step to reproduce `external volume "timeful-test-go-build-cache" not found`, then create the volume and re-run the exact CI sequence (`up -d --build ...`, `run --rm server-route-test go test ./... -count=1`) green, then `down -v` (retains the external volume).
5. Run `npm run format:markdown:check` (root) for the changed Markdown; verify ACs, record final summary, mark Done.

Round 4 plan - provision CI browser deps via Nix (AC #10):

1. flake.nix: add a frontend-e2e writeShellScriptBin whose closure contains only nodejs_26 + pkgs.playwright-driver.browsers (not the devshell: no go/python/backlog/graphify). Script: set -euo pipefail, export PLAYWRIGHT_BROWSERS_PATH to the nix browsers store path, resolve repo root via git rev-parse --show-toplevel, cd frontend, npm ci, exec npm run test:e2e -- "$@". Expose as packages.frontend-e2e plus an apps.frontend-e2e entry so `nix run .#frontend-e2e` works.

2. backend-ci.yml: replace the inline `E2E_...=true npm ci && E2E_...=true npm run test:e2e ...` step with `nix run .#frontend-e2e -- --project=firefox-desktop timed-event-postgres-plugin-firefox.spec.ts`, keeping E2E_POSTGRES_ANONYMOUS_EVENT_CREATION_ENABLED as step env; remove stray trailing-whitespace lines left by the cache step edit.

3. Pre-verified facts (2026-09-03): pinned nixpkgs playwright-driver is 1.61.1 with browsers revision firefox-1532, exactly matching frontend playwright-core 1.61.0 registry; nix firefox binary interpreter and libs all resolve inside the Nix store (self-contained, no apt browser deps needed on ubuntu-latest); browsers closure 2.15 GB fits the 5G gc-max-store-size cache budget.

4. Verify: actionlint on backend-ci.yml; run the exact CI command locally (`E2E_POSTGRES_ANONYMOUS_EVENT_CREATION_ENABLED=true nix run .#frontend-e2e -- --project=firefox-desktop timed-event-postgres-plugin-firefox.spec.ts`) against the isolated E2E stack (owns mongo-test/postgres-test/server-test on 3003 and Vite on 4174; creates the external Go build cache itself).

5. Leave AC #10 unchecked until the next CI run completes end to end; keep the task In Progress (same pattern as round 3).

Round 4 (2026-09-03): the Nix browser-provisioning plan (minimal frontend-e2e flake entry plus workflow nix run step) moved to TASK-0149; this task closes with rounds 1-2 (wait race fix, external Go build cache volume) and the green run-tests step.
<!-- SECTION:PLAN:END -->

## Implementation Notes

<!-- SECTION:NOTES:BEGIN -->
Execution notes:
- CI run 33316150651 job 99269864485: up step succeeded (postgres-test-migrate Started 14:12:11.6229), wait step failed 14:12:11.7659 with `no containers for project "timeful-test"`.
- compose Wait() source: pkg/compose/wait.go uses getContainers(..., oneOffInclude, false /* all */, services...) so exited containers are invisible; identical behavior in v2.29.7, v2.39.3, and main.
- compose run: pkg/compose/run.go calls waitDependencies (run.go:174), which honors depends_on conditions including service_completed_successfully.
- Migration set: server/migrations contains 2 SQL files (20260814170000, 20260815100000), so migrate completes in well under a second on CI.
- Every Backend CI run (2026-08-25, 26, 29, 30) failed at the wait step; local E2E bypasses wait entirely via isolated-test-stack.ts healthcheck polling.

Round 2 verification (2026-09-03): volume removed → CI up step succeeded → CI run command reproduced `external volume "timeful-test-go-build-cache" not found` exactly → `docker volume create` → exact CI sequence green (`go test ./... -count=1`, exit 0, cold cache) → `down -v` retained the external volume. `docker compose config --quiet` and `npm run format:markdown:check` pass.

Round 3 (2026-09-03): run-tests step is green after the volume fix; overall CI still fails because Playwright browsers are unavailable on the runner (not installed by the workflow). Planned resolution: provision browsers via Nix later. Task reopened (In Progress) until the workflow completes end to end.
<!-- SECTION:NOTES:END -->

## Comments

<!-- COMMENTS:BEGIN -->
author: agent
created: 2026-09-03 11:25
---
Reopened after Backend CI failed at **Run backend tests** with `external volume "timeful-test-go-build-cache" not found`.

Root cause: `compose.test.yaml` declares the Go build cache as an external volume (`go_build_cache`, name `timeful-test-go-build-cache`). Compose never creates external volumes, so they must pre-exist. The local E2E harness creates it (`ensureGoBuildCacheVolume`, `frontend/e2e/isolated-test-stack.ts`), so every local run had it, which is why the previous verification passed. Backend CI had never reached the `server-route-test` step before (the removed wait step always failed first), so the missing volume only surfaced now. The up step does not hit it because none of `mongo-test`/`postgres-test`/`postgres-test-bootstrap`/`postgres-test-migrate` mounts the cache volume.

Fix direction: mirror the E2E harness in CI — create the external volume (idempotently) before the compose up/run steps, and add the volume-create line to the documented route-test snippet.
---

author: agent
created: 2026-09-03 11:25
---
Verification gap from round 1: the local verification reused a machine where the volume already existed (created 2026-08-15 by E2E runs), so the fresh-runner condition was never exercised. Round 2 verification must first remove the volume, reproduce the error, then run the full CI sequence green.
---
<!-- COMMENTS:END -->

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

## Round 2 (reopened 2026-09-03): external Go build cache volume

## Problem

Backend CI failed at "Run backend tests" with `external volume "timeful-test-go-build-cache" not found` once the round-1 fix removed the wait step that had always failed first.

## Root cause

`compose.test.yaml` declares the Go build cache as an external volume (`go_build_cache`, name `timeful-test-go-build-cache`); compose never creates external volumes, and fresh CI runners have none. The up step succeeds because none of `mongo-test`/`postgres-test`/`postgres-test-bootstrap`/`postgres-test-migrate` mounts the volume; only `server-route-test` (and `server-test`) do. Round-1 local verification passed because the machine's volume already existed (created 2026-08-15 by E2E runs), so the fresh-runner condition was never exercised.

## Fix

- `.github/workflows/backend-ci.yml`: added an idempotent "Ensure external Go build cache volume exists" step (`docker volume create timeful-test-go-build-cache`) before the compose up/run steps, mirroring `ensureGoBuildCacheVolume` in `frontend/e2e/isolated-test-stack.ts`.
- `docs/environments.md`: added the volume-create line to the route-test snippet and documented why it is needed.
- `AGENTS.md`: Server Test Workflow now instructs creating the external volume first.

## Verification (fresh-runner condition, isolated test stack)

- Removed the external volume, then ran the CI up step: succeeded (no mounted volume), matching CI.
- CI run command without the volume: failed with the exact CI error `external volume "timeful-test-go-build-cache" not found`.
- After `docker volume create timeful-test-go-build-cache`, the exact CI sequence (`up -d --build ...`, then `run --rm server-route-test go test ./... -count=1`) passed green with a cold cache (exit 0).
- `down -v` removed the stack and retained the external volume; `docker compose config --quiet` and `npm run format:markdown:check` pass.

## Open

AC #9 (Backend CI green on the run-tests step) is confirmed by the next CI run after this commit; reopen the task if it fails. Browser E2E remains exempt: the change touches CI configuration and docs only, and the E2E harness already creates the volume itself.

## Round 3 status (2026-09-03): kept open

The round-2 volume fix holds: the run-tests step (go test) no longer fails on the external volume. CI now fails later in the workflow at the browser lifecycle step because Playwright browsers are not available on the runner; browsers will be brought via Nix in later work. Task stays open until CI completes end to end.

## Round 4 close (2026-09-03): run-tests goal achieved, Nix work split out

The run-tests step is green on CI after the round-2 external-volume fix (commit 75729b93), completing this task's goal for the backend test path; AC #9 is checked on that evidence.

The remaining open item (former AC #10, browsers for the browser lifecycle step) is a separate concern and moved to TASK-0149 - Run Backend CI browser E2E with Nix-provided dependencies, which carries the verified Nix facts (pinned playwright-driver 1.61.1 / firefox-1532 matches playwright-core 1.61.0; Nix browsers are self-contained patched builds needing no apt packages; browsers closure 2.15 GB) and the plan for a minimal `frontend-e2e` flake entry plus the cache-nix-action groundwork already in the workflow.

DoD #3 (e2e): the e2e-level verification for this CI-configuration change is the exact CI sequence run green against the isolated test stack (round 2); the browser lifecycle spec itself leaves this task's scope with the AC #10 split and is TASK-0149's acceptance criterion #5. Closing as Done.
<!-- SECTION:FINAL_SUMMARY:END -->
