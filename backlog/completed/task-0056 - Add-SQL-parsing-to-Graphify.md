---
id: TASK-0056
title: Add SQL parsing to Graphify
status: Done
assignee:
  - opencode
created_date: '2026-08-24 15:08'
updated_date: '2026-08-29 10:15'
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
- [x] #1 The repository-managed Graphify environment includes the SQL parsing dependency
- [x] #2 graphify update . parses the repository SQL files without the missing tree_sitter_sql warning
- [x] #3 The refreshed local graph contains nodes or relationships extracted from supported SQL files
- [x] #4 Nix flake validation passes
<!-- AC:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [x] #1 All acceptance criteria are satisfied
- [ ] #2 All required unit tests pass. Documentation-only changes are exempt unless the user requests unit tests
- [ ] #3 All required e2e tests pass. Documentation-only changes are exempt unless the user requests e2e tests
<!-- DOD:END -->

## Implementation Plan

<!-- SECTION:PLAN:BEGIN -->
1. Confirm the flake's pinned nixpkgs (c8f9065) exposes `python3Packages.tree-sitter-sql` compatible with graphify 0.9.42.
2. In `flake.nix`, extend the `graphify-cli` wrapper's target with an override that appends `tree-sitter-sql` to graphify's Python propagated build inputs (the `[sql]` extra equivalent), so no manual pip install is needed.
3. Verify in a fresh `nix develop -c` shell that `tree_sitter_sql` imports and that `graphify extract` on a repository `.sql` file emits no missing-dependency error.
4. Run `graphify update .` to refresh the local graph and confirm SQL-derived nodes/relationships appear.
5. Validate the flake (`nix flake check` or devshell build).
<!-- SECTION:PLAN:END -->

## Final Summary

<!-- SECTION:FINAL_SUMMARY:BEGIN -->
Added the Graphify SQL parser dependency to the repository-managed environment and refreshed the local knowledge graph with SQL extraction enabled.

Changes:
- flake.nix: the `graphify-cli` wrapper now targets `graphify-sql`, a `pkgs.graphify.overridePythonAttrs` that appends `pkgs.python3.pkgs.tree-sitter-sql` (0.3.11) to graphify's propagated build inputs — the Nix equivalent of `pip install 'graphifyy[sql]'`.

Evidence:
- The rebuilt graphify 0.9.42 `.graphify-wrapped` entrypoint lists `python3.14-tree-sitter-sql-0.3.11` in its `site.addsitedir` path, so `import tree_sitter_sql` resolves inside the managed environment.
- `graphify update .` re-extracted 432 files with no `tree_sitter_sql` missing-dependency warning (only unrelated zero-node warnings for four JSON files, graphify issue #1666).
- The refreshed graph (4755 nodes, 8267 edges) now contains SQL-derived nodes for `postgres_events` and `postgres_event_responses` plus `contains`/`references` edges extracted from `server/migrations/20260814170000_postgres_anonymous_event_compatibility.sql` (`_origin: ast`).
- `nix flake check` passes for `devShells.x86_64-linux.default`.

Unit and e2e suites do not apply: the change is Nix devshell/build configuration only; no frontend, backend, or runtime code was modified.
<!-- SECTION:FINAL_SUMMARY:END -->
