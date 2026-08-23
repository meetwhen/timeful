---
id: TASK-0050
title: Refresh Graphify semantic index
status: Done
assignee:
  - OpenCode
created_date: '2026-08-22 21:15'
updated_date: '2026-08-22 21:24'
labels: []
dependencies: []
references:
  - graphify-out/graph.json
documentation:
  - .agents/skills/graphify/references/update.md
modified_files:
  - graphify-out/
priority: medium
type: chore
ordinal: 43000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Refresh the repository Graphify index incrementally so its semantic graph reflects added, changed, and deleted corpus files.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [x] #1 The incremental scan identifies changed and deleted corpus files
- [x] #2 Changed semantic content is re-extracted and merged into the graph
- [x] #3 Graphify reports and visualization reflect the merged graph
- [x] #4 The graph health check completes and any integrity warnings are recorded
<!-- AC:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [x] #1 All acceptance criteria are satisfied
- [x] #2 All required unit tests pass. Documentation-only changes are exempt unless the user requests unit tests
- [x] #3 All required e2e tests pass. Documentation-only changes are exempt unless the user requests e2e tests
<!-- DOD:END -->

## Implementation Plan

<!-- SECTION:PLAN:BEGIN -->
1. Run Graphify incremental detection against the stored scan root. 2. If corpus files changed, extract changed code and semantic content, merge with the existing graph, and rebuild reports and HTML. 3. Run Graphify diagnostics and record graph-diff and health evidence. 4. If no changes are detected, report the index is already current without rewriting outputs.
<!-- SECTION:PLAN:END -->

## Implementation Notes

<!-- SECTION:NOTES:BEGIN -->
The stored Graphify interpreter inherited incompatible Nix NumPy packages. Updated `graphify-out/.graphify_python` to the isolated UV interpreter and invoked Graphify with `PYTHONPATH`/`PYTHONHOME` cleared. Initial merge re-extracted 80 documents, replaced 60 source nodes, pruned 7 nodes for `frontend/glossary.md`, and produced 4,272 nodes / 7,019 edges. A follow-up scan captured two concurrently changed Listmonk Go files; final graph is 4,281 nodes / 7,031 edges. `docs/requirements/constraints.md` was re-queued but semantically empty, so its cache stamp was intentionally cleared for a future retry. Diagnostics reported zero missing or dangling endpoints, self-loops, and collapsed edges.
<!-- SECTION:NOTES:END -->

## Final Summary

<!-- SECTION:FINAL_SUMMARY:BEGIN -->
Refreshed the Graphify semantic index incrementally. Re-extracted changed project and requirements documentation, merged changed Listmonk Go source, and pruned the deleted frontend glossary source. Regenerated `graphify-out/graph.json`, `GRAPH_REPORT.md`, and HTML visualization. Verified graph integrity: 4,281 nodes and 7,031 edges with no dangling or missing endpoints, self-loops, or collapsed edges. Frontend unit and E2E checks were not run because this task only updates generated Graphify metadata; the project Definition of Done exempts documentation-only work.
<!-- SECTION:FINAL_SUMMARY:END -->
