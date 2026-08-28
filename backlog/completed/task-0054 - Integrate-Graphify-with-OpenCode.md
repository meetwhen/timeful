---
id: TASK-0054
title: Integrate Graphify with OpenCode
status: Done
assignee:
  - OpenCode
created_date: '2026-08-24 14:27'
updated_date: '2026-08-24 14:47'
labels: []
dependencies: []
references:
  - 'https://github.com/Graphify-Labs/graphify#install'
  - 'https://opencode.ai/docs/mcp-servers/'
documentation:
  - backlog/handoffs/handoff-2026-08-24T14-22-40Z.md
modified_files:
  - AGENTS.md
  - .gitattributes
  - .opencode/opencode.json
  - .opencode/plugins/graphify.js
  - .opencode/skills/graphify/.graphify_version
  - .opencode/skills/graphify/SKILL.md
  - flake.nix
  - opencode.json
priority: medium
type: enhancement
ordinal: 56000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Provide this repository with Graphify's supported project-scoped OpenCode guidance, local graph MCP access, and code-change freshness automation so OpenCode sessions can query the semantic graph reliably.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [x] #1 OpenCode has Graphify query-first guidance in the repository
- [x] #2 OpenCode can connect to the local Graphify MCP server for graph queries
- [x] #3 Graphify code-change refresh automation is installed and reports healthy status
- [x] #4 The configured integration is verified against the installed Graphify and OpenCode CLIs
<!-- AC:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [x] #1 All acceptance criteria are satisfied
- [x] #2 All required unit tests pass. Documentation-only changes are exempt unless the user requests unit tests
- [x] #3 All required e2e tests pass. Documentation-only changes are exempt unless the user requests e2e tests
<!-- DOD:END -->

## Implementation Plan

<!-- SECTION:PLAN:BEGIN -->
1. Retain Graphify's supported project-scoped OpenCode guidance and code-change hooks already installed.
2. Make the Nix devshell the reproducible Graphify MCP runtime by providing Graphify's Python MCP dependency from the pinned nixpkgs input.
3. Replace the machine-local uv interpreter in opencode.json with a nix develop launcher that starts graphify.serve against graphify-out/graph.json.
4. Verify imports in the devshell, OpenCode config and MCP connectivity, a graph query, hook status, and nix flake check. Do not install Graphify or MCP globally with uv.
<!-- SECTION:PLAN:END -->

## Implementation Notes

<!-- SECTION:NOTES:BEGIN -->
2026-08-24: User approved expanding TASK-0054 to cover the Nix devshell reproducibility needed for its local MCP acceptance criterion. The pinned devshell's Graphify 0.9.42 installation cannot import the optional mcp module; the current opencode.json command points to a machine-local uv archive interpreter. The planned correction is a Nix-provided mcp dependency and portable nix develop launcher.

2026-08-24: Added pkgs.python3Packages.mcp to the pinned Nix devshell and replaced the machine-local uv archive interpreter in opencode.json with `nix develop . --command python3 -m graphify.serve graphify-out/graph.json` (60-second MCP startup timeout). Devshell verification imports Graphify 0.9.42 and MCP 1.29.0. `opencode mcp list` reports both Graphify and Backlog connected. `graphify hook status` reports post-commit and post-checkout installed plus merge driver registered. `nix flake check` passes for the x86_64-linux devshell.
<!-- SECTION:NOTES:END -->

## Final Summary

<!-- SECTION:FINAL_SUMMARY:BEGIN -->
Integrated Graphify with OpenCode using Graphify's project-scoped guidance and hooks, then made the local Graphify MCP runtime reproducible through the pinned Nix devshell. The devshell now provides `pkgs.python3Packages.mcp`; `opencode.json` launches `graphify.serve` through `nix develop .` rather than a machine-local uv archive interpreter.

Verification: `nix develop . --command python3` imported Graphify 0.9.42 and MCP 1.29.0; `opencode debug config` resolved the local MCP declaration; `opencode mcp list` reported Graphify connected; a direct MCP client enumerated Graphify tools and successfully called `query_graph`; `graphify hook status` reported both hooks installed and the merge driver registered; `nix flake check` passed for x86_64-linux. No frontend or runtime source changed, so frontend unit and e2e suites were not applicable.
<!-- SECTION:FINAL_SUMMARY:END -->
