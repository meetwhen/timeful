---
id: TASK-0055
title: Launch Graphify MCP without the devshell
status: Done
assignee:
  - OpenCode
created_date: '2026-08-24 14:57'
updated_date: '2026-08-24 15:02'
labels: []
dependencies: []
references:
  - flake.nix
  - opencode.json
  - backlog/handoffs/handoff-2026-08-24T14-49-57Z.md
modified_files:
  - flake.nix
  - opencode.json
  - graphify-out/.graphify_labels.json
priority: medium
type: chore
ordinal: 57000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Provide a lean, reproducible repository-local Nix entry point for the Graphify MCP server so OpenCode does not start the full development shell for each MCP process.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [x] #1 OpenCode launches the local Graphify MCP server through a dedicated Nix run target
- [x] #2 The MCP target supplies the pinned Graphify and MCP Python dependencies without entering the full development shell
- [x] #3 The Graphify MCP server starts successfully against graphify-out/graph.json
- [x] #4 Nix flake validation passes
<!-- AC:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [x] #1 All acceptance criteria are satisfied
- [x] #2 All required unit tests pass. Documentation-only changes are exempt unless the user requests unit tests
- [x] #3 All required e2e tests pass. Documentation-only changes are exempt unless the user requests e2e tests
<!-- DOD:END -->

## Implementation Plan

<!-- SECTION:PLAN:BEGIN -->
1. Define one `apps.graphify-mcp` flake output that runs a wrapper with only the pinned Python interpreter, Graphify, and MCP Python paths.
2. Change the OpenCode local MCP command from `nix develop . --command ...` to `nix run .#graphify-mcp`.
3. Verify the wrapper imports Graphify and MCP, starts the stdio server against the repository graph, and passes `nix flake check`.

Research: Graphify is a top-level nixpkgs application rather than `python3Packages.graphify`; MCP is `python3Packages.mcp`. The wrapper will use `pkgs.makePythonPath` to make both packages importable by the pinned Python interpreter without constructing the devshell.

Implementation detail: used `pkgs.writeShellScriptBin` rather than `writeShellApplication` so the MCP app does not add ShellCheck to its runtime/build closure. The wrapper uses `pkgs.python3Packages.makePythonPath` and the pinned `pkgs.python3` executable.
<!-- SECTION:PLAN:END -->

## Implementation Notes

<!-- SECTION:NOTES:BEGIN -->
Implemented `apps.graphify-mcp` and changed the OpenCode MCP command to `nix run .#graphify-mcp`. `opencode mcp list` reports Graphify connected. `timeout 5s nix run .#graphify-mcp` remained running with no startup error, demonstrating the stdio server started. `nix flake check` passed. Ran `graphify update .` as required by repository guidance; it rebuilt the ignored local graph and updated the tracked label cache, reporting existing source/SQL extraction warnings unrelated to this configuration change.
<!-- SECTION:NOTES:END -->

## Final Summary

<!-- SECTION:FINAL_SUMMARY:BEGIN -->
Added a dedicated `apps.graphify-mcp` Nix app that starts Graphify's stdio MCP server with only the pinned Python, Graphify, and MCP module paths. OpenCode now runs this app with `nix run .#graphify-mcp` instead of constructing the full development shell. Verified the app remains running as the MCP server, `opencode mcp list` connects successfully, and `nix flake check` passes. No frontend unit or e2e suite applies to this root Nix/OpenCode configuration change.
<!-- SECTION:FINAL_SUMMARY:END -->
