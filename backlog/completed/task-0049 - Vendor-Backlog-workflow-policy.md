---
id: TASK-0049
title: Vendor Backlog workflow policy
status: Done
assignee:
  - OpenCode
created_date: '2026-08-22 21:05'
updated_date: '2026-08-22 21:06'
labels: []
dependencies: []
references:
  - AGENTS.md
  - backlog/backlog.md
priority: medium
type: docs
ordinal: 42000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Make the repository, rather than mutable MCP workflow resources, the authoritative source for Backlog working instructions. Keep Backlog MCP as the management transport and update the project Definition of Done so documentation-only changes do not require unit or e2e tests unless requested by the user.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [x] #1 A repository-owned Backlog workflow document defines the project policy for task management
- [x] #2 Root agent instructions direct agents to the repository-owned workflow document instead of requiring MCP workflow resources
- [x] #3 Project Definition of Done exempts documentation-only changes from unit tests unless the user requests them
- [x] #4 Project Definition of Done exempts documentation-only changes from e2e tests unless the user requests them
<!-- AC:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [x] #1 All acceptance criteria are satisfied
<!-- DOD:END -->

## Implementation Plan

<!-- SECTION:PLAN:BEGIN -->
1. Add BACKLOG_WORKFLOW.md at the repository root as the authoritative policy for Backlog task management. Capture task triage, search, creation, execution, finalization, scope changes, and the acceptance-criteria versus DoD distinction.
2. Replace the generated Backlog MCP guidance in AGENTS.md with a concise pointer to the vendored policy. State that MCP is the management transport and its workflow resources are not repository policy.
3. Replace project DoD defaults through Backlog MCP. Retain acceptance-criteria completion and add separate documentation-only exemptions for unit and e2e tests when the user has not requested them.
4. Review the two files and re-read DoD configuration. Because this is documentation and Backlog configuration only, do not run test suites.
<!-- SECTION:PLAN:END -->

## Implementation Notes

<!-- SECTION:NOTES:BEGIN -->
Added `BACKLOG_WORKFLOW.md` as the repository-owned policy and changed `AGENTS.md` to reference it. Verified the rendered local policy and root reference directly.

Re-read the project DoD defaults after updating them. The task was created before the new defaults, so its inherited unit and e2e requirements were removed as the documented docs-only exception. No test suite was run because this task changed only documentation and Backlog configuration.
<!-- SECTION:NOTES:END -->

## Final Summary

<!-- SECTION:FINAL_SUMMARY:BEGIN -->
Vendored the Backlog operating policy in `BACKLOG_WORKFLOW.md` and made it the root agent-instruction source of truth. `AGENTS.md` now directs agents to the local policy while retaining Backlog MCP as the management transport. Updated project DoD defaults so documentation-only changes are exempt from unit and e2e tests unless the user requests the respective check. Verification: reviewed the rendered policy and root reference, re-read configured DoD defaults, and ran `git diff --check`; no test suite was run because the change is documentation and Backlog configuration only.
<!-- SECTION:FINAL_SUMMARY:END -->
