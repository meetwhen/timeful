---
id: TASK-0108
title: Silence Node experimental localStorage warning in frontend unit tests
status: Done
assignee:
  - opencode
created_date: '2026-08-29 14:26'
updated_date: '2026-08-29 14:40'
labels:
  - frontend
  - testing
dependencies: []
priority: low
type: chore
ordinal: 114000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Unit test output for the frontend shows the Node warning "ExperimentalWarning: localStorage is not available because --localstorage-file was not provided" once per run. Node >= 22.4 (repo runs v26.x) exposes an experimental built-in localStorage on globalThis in the default node Vitest environment; node-environment test files (no happy-dom pragma) and the source modules they import touch globalThis.localStorage — including bare typeof guards (e.g. src/utils/timezone_utils.ts, src/utils/browserDatePreferences.ts) — which triggers the warning because no --localstorage-file is provided. This is expected, intentional behavior (the code treats missing storage as a valid state), so the goal is only to silence this specific noise without changing test semantics. Known constraints: do NOT pass --localstorage-file (that would make typeof localStorage return "object" and silently stop covering the storage-unavailable fallback paths); the warning goes to raw stderr via Node's default warning handler, so Vitest onConsoleLog does not see it — suppress at process level, matching only this exact warning message and forwarding everything else. The fix should be a small setup file wired via Vitest test.setupFiles, carrying a removal note (delete when the minimum supported Node no longer emits the warning; verify with node -e 'typeof globalThis.localStorage' showing no stderr output).
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [x] #1 Running npm run test:unit in frontend produces no 'ExperimentalWarning: localStorage is not available because --localstorage-file was not provided' output
- [x] #2 Only that exact warning is suppressed; any other Node warning (including other ExperimentalWarnings) still reaches the output
- [x] #3 Test semantics are unchanged: node-environment tests still observe 'typeof localStorage === "undefined"' and the storage-unavailable fallback paths in source modules
- [x] #4 The suppression carries a durable removal note stating when it can be deleted (minimum supported Node no longer emits the warning) and how to verify
- [x] #5 npm run lint, npm run typecheck, npm run build, and npm run test:unit all pass in frontend
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
1. Add `frontend/src/test/silenceNodeLocalStorageWarning.ts`: capture `process.listeners(\"warning\")` (exactly one Node default handler at worker bootstrap), `process.removeAllListeners(\"warning\")`, then install a filter listener that swallows only the exact message \"localStorage is not available because --localstorage-file was not provided.\" (name ExperimentalWarning) and re-dispatches every other warning to the captured default listeners. Guard with `typeof process !== \"undefined\"`. Header comment documents removal trigger and rejected alternatives.\n2. Wire it in `frontend/vitest.config.mjs` via `test.setupFiles: [\"./src/test/silenceNodeLocalStorageWarning.ts\"]`.\n3. Revised constraint from execution research: patching `process.emitWarning` does NOT suppress the warning (the internal webstorage emit triggers the default stderr print independently of the public wrapper; verified with node fork probes). Listener replacement is the working mechanism. Do not pass `--localstorage-file` (would flip `typeof localStorage` to \"object\" and silently lose coverage of storage-unavailable fallback paths); do not use `--disable-warning=ExperimentalWarning` (hides all experimental warnings).\n4. Verify: full `npm run test:unit` shows zero warning matches with all tests passing; in-context probe that a non-target warning still prints; lint, typecheck, build pass.
<!-- SECTION:PLAN:END -->

## Implementation Notes

<!-- SECTION:NOTES:BEGIN -->
Plan adjustment (verified empirically, within scope): patching process.emitWarning does NOT suppress the warning - re-examined the first probe and ran a fork probe; the internal webstorage emit triggers the default stderr print independently of the public wrapper. process.on('warning') listeners receive the warning but the default handler is an independent listener that always runs. Working mechanism: capture process.listeners('warning') (exactly 1 default handler at bootstrap), removeAllListeners('warning'), install a filter listener that swallows only the exact message 'localStorage is not available because --localstorage-file was not provided.' and re-dispatches every other warning to the captured default listeners. Node probes (/tmp/opencode/lsprobe) confirmed: target suppressed, other ExperimentalWarning and DeprecationWarning still print with original formatting, typeof globalThis.localStorage stays undefined.

Also rejected --disable-warning=ExperimentalWarning: it would hide every experimental warning, violating AC #2; noted in the setup file header so nobody 'simplifies' into it later.

Verification evidence: full npm run test:unit = 138 files / 958 tests passed with zero warning matches in output (same 958 count as the pre-change run, supporting AC #3 semantics unchanged). In-context AC #2 probe: temporary test emitted a DeprecationWarning while the filter was active and it printed to output; probe file removed after use.

Debug instrumentation (writeSync marker) added during diagnosis was removed from the final setup file.
<!-- SECTION:NOTES:END -->

## Final Summary

<!-- SECTION:FINAL_SUMMARY:BEGIN -->
Silenced the Node `ExperimentalWarning: localStorage is not available because --localstorage-file was not provided` noise in frontend unit tests without changing any test semantics.\n\n## What changed\n- Added `frontend/src/test/silenceNodeLocalStorageWarning.ts`, wired via `test.setupFiles` in `frontend/vitest.config.mjs`. It replaces the process `warning` listener set with a filter that drops only the exact warning message `\"localStorage is not available because --localstorage-file was not provided.\"` and re-dispatches every other warning to the captured Node default handlers. A header comment documents the removal trigger (delete once the minimum supported Node no longer emits the warning; verify with `node -e 'typeof globalThis.localStorage'` printing nothing to stderr) and the rejected alternatives (`--localstorage-file` would flip `typeof localStorage` to `\"object\"` and silently drop storage-unavailable fallback coverage; `--disable-warning=ExperimentalWarning` would hide all experimental warnings).\n\n## Why this mechanism\nResearch during execution showed the warning bypasses the public `process.emitWarning` (patching it does not suppress the default stderr print) and that adding a `process.on(\"warning\")` listener cannot suppress Node's independent default handler. Listener-set replacement was verified with standalone Node fork probes before being applied.\n\n## Verification\n- `npm run test:unit`: 138 files / 958 tests passed, zero warning matches in output (same test count as before the change).\n- AC #2 probe: a temporary test emitted a `DeprecationWarning` while the filter was active and it printed to output as before; probe removed afterward.\n- `npm run lint`, `npm run typecheck`, `npm run build`: all pass.\n\n## Not run\n- Browser e2e was not run: the change only touches the unit-test setup path (`test.setupFiles`), which Playwright E2E and the production build never load. No runtime, build, or deployment configuration is affected.
<!-- SECTION:FINAL_SUMMARY:END -->
