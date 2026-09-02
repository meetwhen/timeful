---
id: TASK-0086
title: 'Commit skill: guide model identifier resolution with helper script'
status: Done
assignee:
  - OpenCode
created_date: '2026-08-27 11:20'
updated_date: '2026-08-27 11:22'
labels: []
dependencies: []
priority: medium
type: task
ordinal: 94000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Goal: make the `Model:` metadata line of commit messages reliable instead of guessed.

Context from completed research (durable facts):
- opencode injects an identity statement into each session system prompt ("You are powered by the model named ... The exact model ID is <provider>/<model>"). This is the primary source.
- Config-based sources are unreliable: resolved config shows no `model` here because models are selected per-session in the TUI; no env var reaches agent tool environments.
- Reliable fallback: get the current session id via `opencode session list --format json --max-count 1`, then query `~/.local/share/opencode/opencode.db` (SQLite, table `message`, columns `session_id` and JSON `data`) read-only for the newest assistant message's `providerID`/`modelID`.
- Verified live in session ses_fbd15dddaffe07uLgK63JW3srq: output was `openrouter/z-ai/glm-5.3-flash`.

Scope constraints:
- Helper script goes to `scripts/opencode/current-model.sh` (repo-level utility, consistent with `scripts/handoff/create-handoff.sh`), NOT inside `.agents/skills/commit/`.
- Follow existing bash conventions (`#!/usr/bin/env bash`, `set -euo pipefail`); use SQLite in read-only URI mode.
- Commit skill stays minimal: identity line first, script as cross-check, ask-the-user guard for other harnesses.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [x] #1 The commit skill states that the Model trailer must be the exact <provider>/<model> identifier, not a bare model name
- [x] #2 The commit skill names the harness-injected identity line as the primary source for the model identifier
- [x] #3 The commit skill references scripts/opencode/current-model.sh as the verification fallback and never suggests deriving the value from opencode.json, opencode debug config, or environment variables
- [x] #4 scripts/opencode/current-model.sh resolves the most recently updated opencode session, queries its latest assistant message from the local session database read-only, and prints providerID/modelID
- [x] #5 The script exits non-zero with a clear stderr message when the session cannot be resolved, the database is missing, or no assistant message exists
- [x] #6 Running the script inside this live opencode session prints openrouter/z-ai/glm-5.3-flash, recorded as verification evidence
- [x] #7 Changed Markdown files pass npm run format:markdown
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
1. Create `scripts/opencode/current-model.sh` following `scripts/handoff/create-handoff.sh` conventions: `#!/usr/bin/env bash`, `set -euo pipefail`.
2. Script flow:
   - Resolve session id from `opencode session list --format json --max-count 1`, validate `ses_` prefix.
   - Open `${OPENCODE_DATA_DIR:-$HOME/.local/share/opencode}/opencode.db` with SQLite read-only URI mode (`file:...?mode=ro`).
   - Query latest assistant message for that session: `SELECT data FROM message WHERE session_id = ? AND json_extract(data,'$.role')='assistant' ORDER BY rowid DESC LIMIT 1`.
   - Print `<providerID>/<modelID>`; exit non-zero with stderr messages for each failure mode.
3. Edit `.agents/skills/commit/SKILL.md`: replace the vague "verify model name current" bullet with a Model-resolution block (identity line primary; script fallback; anti-patterns config/env; ask-user guard for other harnesses); keep one-sentence-per-line Markdown style.
4. Verify: run script live (expect `openrouter/z-ai/glm-5.3-flash`), `bash -n`, shellcheck if present.
5. Run `npm run format:markdown` (frontend workspace) for changed Markdown files.
<!-- SECTION:PLAN:END -->

## Final Summary

<!-- SECTION:FINAL_SUMMARY:BEGIN -->
Added `scripts/opencode/current-model.sh` (executable, bash -n clean), which resolves the most recently updated opencode session via `opencode session list --format json --max-count 1`, then reads the newest assistant message for that session from `$HOME/.local/share/opencode/opencode.db` (override with OPENCODE_DATA_DIR) in SQLite read-only mode and prints `<providerID>/<modelID>`. Failure modes exit non-zero with clear stderr messages; missing-db and no-assistant-message paths were exercised directly (exit 1 each). Rewrote `.agents/skills/commit/SKILL.md` model guidance: Model trailer must be the exact provider/model ID; primary source is the harness-injected identity line ("You are powered by ... The exact model ID is ..."); cross-check fallback is the new repo-root script; explicitly rules out opencode.json / opencode debug config / env vars because model selection is per-session and config-resolved model is often unset; other harnesses must ask the user rather than guess. Verification: running the script inside the live session printed openrouter/z-ai/glm-5.3-flash, matching the identity line. Changed Markdown files pass `npm run format:markdown`. Docs-and-tooling-only change; unit/e2e tests exempt per DoD policy.
<!-- SECTION:FINAL_SUMMARY:END -->
