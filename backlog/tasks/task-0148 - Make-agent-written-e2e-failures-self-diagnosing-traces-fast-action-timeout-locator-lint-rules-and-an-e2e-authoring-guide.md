---
id: TASK-0148
title: >-
  Make agent-written e2e failures self-diagnosing: traces, fast action timeout,
  locator lint rules, and an e2e authoring guide
status: Done
assignee:
  - opencode
created_date: '2026-09-02 22:58'
updated_date: '2026-09-02 23:08'
labels:
  - e2e
  - tooling
  - agent-workflow
dependencies: []
priority: medium
type: chore
ordinal: 161300
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Agents writing browser e2e tests in frontend/e2e lose significant time diagnosing click timeouts and wrong-element matches. Root causes in the current setup: (1) playwright.config.ts uses trace "on-first-retry" while local runs have retries: 0, so traces are never captured where agents actually work; (2) no action timeout, so broken selectors fail at the 30s test timeout far from the cause; (3) nothing forbids brittle selector APIs (page.$, page.waitForSelector, raw waitForTimeout), so generated tests copy fragile patterns from helpers; (4) no written guidance tells an agent how to diagnose a failing spec (isolate, read the action call log, open the trace).

Outcome: agent-written e2e failures surface in seconds with a trace and a precise action log, and new specs follow locator hygiene rules enforced mechanically by lint instead of prose alone.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [x] #1 Local e2e runs capture a trace for every failing test and view it with playwright show-trace, instead of traces only being produced on CI retries.
- [x] #2 playwright.config.ts sets an action timeout shorter than the 30s test timeout so broken selectors fail fast at the exact step.
- [x] #3 ESLint error-level rules under frontend/e2e forbid page.$, page.$$, page.pause, page.waitForSelector, and raw page.waitForTimeout, with one documented settle helper as the only fixed-delay escape hatch.
- [x] #4 Existing e2e specs and helpers pass the new rules through behavior-preserving fixes only (no rewritten waits, changed timing, or changed assertions).
- [x] #5 frontend/e2e/AGENTS.md documents the failure-diagnosis loop (isolate failing test with --grep and --project, read the action call log, open the trace) and authoring conventions (web-first expect assertions, user-facing locators, toHaveCount for ambiguity, test.step, API seeding).
- [x] #6 frontend/AGENTS.md Browser Verification section links to frontend/e2e/AGENTS.md.
- [x] #7 npm run lint, fmt:check, typecheck, build, and test:unit pass.
- [x] #8 At least one representative e2e spec (for example landing-hero) passes end-to-end with the updated config to validate trace and timeout changes.
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
## Implementation Plan

### 1. playwright.config.ts (AC #1, #2)
- `use.trace`: "on-first-retry" → "retain-on-failure" (local runs have `retries: 0`, so on-first-retry never captures anything locally; retain-on-failure also still covers CI retries).
- Add `use.actionTimeout: 15_000` — shorter than the 30s default test timeout so a broken selector fails at the exact action. 15s (not 5–10s) leaves margin for cold first navigations against the on-demand Vite server that e2e owns; explicit per-call timeouts in existing specs override it.

### 2. ESLint locator-hygiene rules (AC #3, #4)
Add a flat-config block in `frontend/eslint.config.ts` scoped to `e2e/**/*.ts` with error-level `no-restricted-syntax` selectors:
- `CallExpression[callee.property.name='waitForSelector']` → use `locator.waitFor()` instead
- `CallExpression[callee.property.name='waitForTimeout']` → use the shared settle helper
- `CallExpression[callee.property.name='$']` and `[callee.property.name='$$']` → use locators
- `CallExpression[callee.property.name='pause']` → use `--debug`/`--ui` or traces
Then a follow-up block exempting only the new `e2e/helpers/settle.ts` from the waitForTimeout selector (config-level exemption, no inline disables). Oxlint already lints e2e separately; these selectors are eslint-only. Note: eslint lint script already ignores `e2e/inspect/**`; repro/ scripts are diagnostic entrypoints — extend the lint ignore or exempt block only if they trip rules (they do use waitForSelector/waitForTimeout; they run via `npm run inspect`-style tooling, not `test:e2e`, so exempting them from these rules is consistent with their diagnostic purpose).

### 3. Behavior-preserving fixes to existing code (AC #4)
All mechanical, no timing changes:
- `firefox-timed-event-harness.ts` waitForSelector ×4 → `page.locator(sel).waitFor()` (same default timeout semantics)
- `timed-event-helpers.ts` waitForTimeout ×2, `firefox-timed-event-harness.ts` ×1, `timed-event-timerange-width-firefox.spec.ts` ×2, `timed-event-timezone-menu-firefox.spec.ts` ×1 → new `settlePage(page, ms)` helper in `e2e/helpers/settle.ts` (wraps waitForTimeout; single documented escape hatch)

### 4. e2e authoring guide (AC #5, #6)
- New `frontend/e2e/AGENTS.md`: failure-diagnosis loop (run failing test in isolation with `--grep` + single `--project`, read the action call log in the error output, open trace with `npx playwright show-trace <trace.zip under frontend/tmp/playwright>`), authoring conventions (web-first `expect(locator).toBeVisible()/toHaveCount()`, user-facing locators getByRole/getByLabel/getByTestId, add data-testid in app code when no accessible role exists, `test.step()` for journeys, API seeding via request fixture, one behavior per test), timeout policy (actionTimeout 15s default, pass explicit timeout only with a reason), and escape hatches (settlePage, DEBUG=pw:api).
- Link it from `frontend/AGENTS.md` Browser Verification section.

### 5. Verification (AC #7, #8)
- `npm run lint`, `fmt:check`, `typecheck`, `build`, `test:unit` in frontend.
- Root `npm run format:markdown` for changed Markdown (DoD).
- Run one representative spec through `npm run test:e2e -- --project=chromium-desktop landing-hero` to validate config changes end-to-end.

### Risks
- actionTimeout 15s vs cold Vite transforms on first test of a run → mitigated by 15s choice and explicit timeouts in existing specs; validated by AC #8 run.
- repro/inspect exemption decision → keep rules scoped to what `npm run lint` eslint actually covers (`e2e` minus `e2e/inspect`); if repro trips rules, exempt it in config with a comment-free named block.
<!-- SECTION:PLAN:END -->

## Implementation Notes

<!-- SECTION:NOTES:BEGIN -->
actionTimeout set to 15s rather than the 5-10s discussed in planning: e2e owns a Vite server that transforms modules on demand, so the first navigation of a run can be slow; the validated landing-hero run was a cold start and passed comfortably.

Existing waitForTimeout call sites are deliberate settle waits (fallback-open pacing, consent-click settle, resize settle); they were routed through a single settlePage helper with a config-level lint exemption instead of being rewritten into state waits, to keep the change strictly behavior-preserving.

e2e/repro/** is exempt from the locator-hygiene selectors because those scripts are diagnostic entrypoints run outside test:e2e; e2e/inspect/** was already eslint-ignored via the lint script.

AC #1 and #8 verified with a live run: intentional-failure probe produced trace.zip + error-context.md and the error output printed the show-trace usage; landing-hero passed on both chromium-desktop and chromium-mobile projects.

Pre-existing lint warnings (vue/one-component-per-file in NewEvent.test.ts) and lightningcss :deep build warnings are unrelated to this task.
<!-- SECTION:NOTES:END -->

## Final Summary

<!-- SECTION:FINAL_SUMMARY:BEGIN -->
## What changed

Made agent-written e2e failures self-diagnosing: failures now surface in seconds at the exact failing step with a full trace, and locator hygiene is enforced mechanically instead of by prose.

### frontend/playwright.config.ts
- `use.trace`: "on-first-retry" → "retain-on-failure" — local runs have `retries: 0`, so traces were never captured locally; now every failed test produces a trace, and CI retries are still covered.
- Added `use.actionTimeout: 15_000` — broken selectors now fail at the action with its call log instead of the 30s test timeout; 15s leaves margin for cold Vite on-demand transforms.

### frontend/eslint.config.ts
- Extracted the Temporal `no-restricted-syntax` selectors into a shared const (no rule changes) and added an error-level e2e locator-hygiene block for `e2e/**/*.ts` forbidding `page.waitForSelector`, raw `page.waitForTimeout`, `page.$`, `page.$$`, and `page.pause`.
- Config-level exemptions (no inline disables): `e2e/helpers/settle.ts` (the single fixed-delay escape hatch) and `e2e/repro/**` (diagnostic entrypoints outside test:e2e, consistent with `e2e/inspect/**` already being eslint-ignored).

### Behavior-preserving e2e fixes (no timing, wait, or assertion changes)
- All `page.waitForSelector(sel)` → `page.locator(sel).waitFor()` (identical semantics): 4 sites in `firefox-timed-event-harness.ts`.
- All raw `page.waitForTimeout` → new `settlePage(page, ms)` helper in `e2e/helpers/settle.ts`: 6 sites across `timed-event-helpers.ts`, `firefox-timed-event-harness.ts`, `timed-event-timerange-width-firefox.spec.ts`, `timed-event-timezone-menu-firefox.spec.ts`, `timed-event-days-only-timezone-firefox.spec.ts`.

### Documentation
- New `frontend/e2e/AGENTS.md`: failure-diagnosis loop (isolate with `--grep` + `--project`, read the action call log, open the trace with show-trace, never fix timeouts with sleeps), authoring rules (user-facing locators, web-first expectations, `toHaveCount` for ambiguity, explicit timeouts only with a reason, `test.step()`, API seeding, `settlePage` as the only fixed-delay escape), and environment ownership.
- `frontend/AGENTS.md` Browser Verification links to it.

## Verification
- `npm run lint` (0 errors; 2 pre-existing warnings in NewEvent.test.ts), `fmt:check`, `typecheck`, `build`, `test:unit` (138 files, 990 tests) all pass.
- `npm run test:e2e -- --project=chromium-desktop landing-hero.spec.ts`: desktop contract test passed (13.4s) against the fresh isolated stack; `--project=chromium-mobile` mobile test passed (5.6s).
- Failure-trace evidence: a temporary intentionally-failing spec run locally (retries 0) produced `trace.zip` and `error-context.md` under `tmp/playwright/` with the exact `npx playwright show-trace <path>` usage printed in the error output; probe spec and artifacts were removed afterward.
- Root `npm run format:markdown` run; `format:check` clean.

## Risks / follow-ups
- actionTimeout 15s could flake on unusually slow cold starts; the validated run was a cold start and passed. Raise per-call with a documented reason if a legitimately slow action appears.
- A failure-context fixture (auto-attaching browser console/page errors on failure) was considered and intentionally deferred; traces already capture console and network.
<!-- SECTION:FINAL_SUMMARY:END -->
