---
id: TASK-0056
title: Add SQL parsing to Graphify
status: To Do
assignee: []
created_date: '2026-08-24 15:08'
labels: []
dependencies: []
references:
  - flake.nix
  - backlog/tasks/task-0055 - Launch-Graphify-MCP-without-the-devshell.md
priority: low
type: enhancement
ordinal: 58000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Enable Graphify's optional SQL parser in the repository-managed environment so SQL files contribute database structure and relationship nodes to the local knowledge graph.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 The repository-managed Graphify environment includes the SQL parsing dependency
- [ ] #2 graphify update . parses the repository SQL files without the missing tree_sitter_sql warning
- [ ] #3 The refreshed local graph contains nodes or relationships extracted from supported SQL files
- [ ] #4 Nix flake validation passes
<!-- AC:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [ ] #1 All acceptance criteria are satisfied
- [ ] #2 All required unit tests pass. Documentation-only changes are exempt unless the user requests unit tests
- [ ] #3 All required e2e tests pass. Documentation-only changes are exempt unless the user requests e2e tests
<!-- DOD:END -->
