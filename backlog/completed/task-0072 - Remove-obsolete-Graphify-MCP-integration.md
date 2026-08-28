---
id: TASK-0072
title: Remove obsolete Graphify MCP integration
status: Done
assignee: []
created_date: '2026-08-25 21:54'
updated_date: '2026-08-25 21:55'
labels: []
dependencies: []
modified_files:
  - opencode.json
  - flake.nix
priority: low
type: chore
ordinal: 76000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Retire the nonfunctional Graphify MCP launcher from OpenCode and the Nix flake. Graphify remains available through its supported OpenCode plugin, which was verified in a new OpenCode session.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [x] #1 OpenCode no longer launches Graphify through the obsolete MCP configuration
- [x] #2 The Nix flake no longer exposes the nonfunctional graphify-mcp app
- [x] #3 Graphify plugin integration works in a newly started OpenCode session
<!-- AC:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [x] #1 All acceptance criteria are satisfied
- [x] #2 All required unit tests pass. Documentation-only changes are exempt unless the user requests unit tests
- [x] #3 All required e2e tests pass. Documentation-only changes are exempt unless the user requests e2e tests
<!-- DOD:END -->

## Implementation Plan

<!-- SECTION:PLAN:BEGIN -->
1. Reproduce the MCP startup failure. 2. Remove the obsolete OpenCode MCP configuration and Nix app. 3. Validate the flake and confirm the supported plugin integration.
<!-- SECTION:PLAN:END -->

## Implementation Notes

<!-- SECTION:NOTES:BEGIN -->
Reproduced `nix run .#graphify-mcp` failure: `python3 -m graphify.serve` exited with `ModuleNotFoundError` because the generated launcher omitted an importable Graphify module.

Removed the obsolete MCP configuration from `opencode.json` and the unused `graphify-mcp` flake app. The repository Graphify plugin remains unchanged.

Verified with `jq empty opencode.json`, `git diff --check`, and `nix flake show --no-write-lock-file`. The user verified Graphify operates as a plugin in a new OpenCode session.
<!-- SECTION:NOTES:END -->

## Final Summary

<!-- SECTION:FINAL_SUMMARY:BEGIN -->
Removed the obsolete Graphify MCP configuration from `opencode.json` and the nonfunctional `graphify-mcp` Nix app from `flake.nix`. Preserved the supported Graphify OpenCode plugin integration. Verification: `jq empty opencode.json`, `git diff --check`, and `nix flake show --no-write-lock-file` passed; the user confirmed the plugin works in a fresh OpenCode session.
<!-- SECTION:FINAL_SUMMARY:END -->
