---
id: TASK-0051
title: Require Backlog policy review for non-trivial work
status: To Do
assignee: []
created_date: '2026-08-23 11:26'
labels: []
dependencies: []
references:
  - AGENTS.md
  - BACKLOG_WORKFLOW.md
  - TASK-0049
modified_files:
  - AGENTS.md
priority: medium
type: docs
ordinal: 45000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Ensure root agent instructions require agents to consult the repository-owned Backlog policy before taking non-trivial implementation action, so work requiring investigation, design, sequencing, or implementation decisions is evaluated for Backlog tracking without imposing unnecessary overhead on questions or mechanical edits.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 Root agent instructions require review of the repository-owned Backlog policy before non-trivial implementation action
- [ ] #2 The instruction directs agents to use the policy when deciding whether work requires a Backlog task
- [ ] #3 Questions, exploration, and obvious mechanical changes are explicitly excluded from the review requirement
- [ ] #4 The instruction preserves Backlog MCP as the required transport for managed records
<!-- AC:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [ ] #1 All acceptance criteria are satisfied
- [ ] #2 All required unit tests pass. Documentation-only changes are exempt unless the user requests unit tests
- [ ] #3 All required e2e tests pass. Documentation-only changes are exempt unless the user requests e2e tests
<!-- DOD:END -->
