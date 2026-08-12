---
name: session-handoff
description: Session handoff, handoff notes, or handoff file: use when the user explicitly asks to capture durable context from the current session for later work.
---

# Session Handoff

Create a handoff only when the user explicitly requests one. Do not create one automatically at task completion.

<!--
Content guidance adapted from Matt Pocock's handoff skill:
https://github.com/mattpocock/skills/blob/main/skills/productivity/handoff/SKILL.md
Licensed under MIT.
-->

## Create The File

1. Immediately before writing the handoff, capture the most recently updated OpenCode session and the UTC creation timestamp. Use:

   ```sh
   python3 <<'PY'
   import json
   import subprocess
   import sys
   from datetime import UTC, datetime

   result = subprocess.run(
       ["opencode", "session", "list", "--format", "json", "--max-count", "1"],
       check=True,
       capture_output=True,
       text=True,
   )
   sessions = json.loads(result.stdout)
   session_id = sessions[0].get("id") if len(sessions) == 1 else None
   if not isinstance(session_id, str) or not session_id.startswith("ses_"):
       sys.exit("Could not determine the most recently updated OpenCode session.")

   print(f"session_id: {session_id}")
   print(f"created_at: {datetime.now(UTC):%Y-%m-%dT%H-%M-%SZ}")
   PY
   ```

   The output labels provide the `session_id` and `created_at` values to use below. The OpenCode CLI has no explicit current-session command or environment variable. This records the most recently updated session for the project, which is the best supported lookup. If another session in the same directory is updated concurrently, stop and ask the user to confirm the session ID rather than recording an uncertain ID.
2. Create `handoff/` if it does not exist.
3. Write the handoff to `handoff/handoff-<created_at>.md`, replacing `<created_at>` with the script output.
4. Report the resulting path.

## Content

Capture only facts established during the current session. The handoff must be useful without the chat transcript and must not invent completion, verification, risks, or decisions.

Do not duplicate material already captured in a specification, plan, ADR, issue, commit, diff, or test. Reference the artifact by repository path or URL and record only the current status or implication. Redact secrets and personal data; never include credentials, tokens, or other sensitive values.

Use this structure, omitting sections with no content:

```markdown
# Session Handoff

- Created session: `<OpenCode session ID>`

## Session Objective

## Motivation And Impact

## Current State

## Completed Work

## Important Decisions

## Decision Rationale

## Constraints And Assumptions

## Findings And Risks

## Dependencies And Prerequisites

## Validation

## Remaining Work

## Next Recommended Action

## Handoff Relationships

## Suggested Skills

## Relevant Files
```

## Section Guidance

- **Session Objective**: State the original problem, intended outcome, and its current state: completed, in progress, blocked, or not started. Describe the user's goal rather than only the implementation task.
- **Motivation And Impact**: Explain why the work is being done and the practical consequence of success or inaction. Include affected users, workflows, reliability, performance, or maintenance concerns only when established.
- **Current State**: Describe the relevant state at handoff time, especially for partial work: implemented but unvalidated, failing test behavior, active worktree changes, deployed state, or known runtime behavior. Do not duplicate **Completed Work**.
- **Completed Work**: List concrete work completed during the session. Reference the implementation, test, specification, issue, commit, or other durable artifact rather than duplicating its detail.
- **Important Decisions**: Record decisions that govern future work, including the selected approach and its consequence. Include only decisions made during the session or confirmed by the user.
- **Decision Rationale**: Explain why an important decision was made. Include rejected alternatives only when that context prevents a likely reversal or repeated investigation. Reference an ADR, issue, or plan instead of restating it.
- **Constraints And Assumptions**: Record scope boundaries, compatibility requirements, operational limitations, and assumptions that still need validation. Clearly distinguish confirmed constraints from assumptions.
- **Findings And Risks**: Capture discoveries that may affect correctness, delivery, or future investigation, plus concrete remaining risks. Avoid speculative risks without an actionable basis.
- **Dependencies And Prerequisites**: List external services, configuration, migrations, approvals, data state, related branches or issues, and ordering requirements needed to proceed. Redact sensitive values.
- **Validation**: Include commands actually run, their outcomes, and meaningful checks not run. Include exact commands when they materially establish confidence in the work.
- **Remaining Work**: Describe uncompleted work as concrete, actionable tasks. Include blockers and prerequisites. Tailor it to the user's stated next-session focus when one is provided, and do not repeat optional work already covered by a tracked issue without linking it.
- **Next Recommended Action**: State the single most useful first action for the next session, particularly when sequencing matters. Omit when **Remaining Work** is already unambiguous.
- **Handoff Relationships**: Describe how this handoff relates to earlier handoffs or durable planning artifacts. Use `Continues` for still-active prior work and `Supersedes` only when prior conclusions or next steps are no longer current. Omit when there is no meaningful relationship.
- **Suggested Skills**: Name only repository skills that are known to exist and directly useful for the remaining work.
- **Relevant Files**: Include changed or inspected paths that a future session needs to understand or continue the work. Add a brief reason when the path alone is not self-explanatory.

Keep entries concise, factual, and actionable.
