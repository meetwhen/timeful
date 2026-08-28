---
id: TASK-0058
title: Isolate Graphify Python dependencies from development shell
status: Done
assignee:
  - OpenCode
created_date: '2026-08-24 16:22'
updated_date: '2026-08-24 16:25'
labels: []
dependencies: []
references:
  - flake.nix
modified_files:
  - flake.nix
  - graphify-out/.graphify_labels.json
priority: medium
type: bug
ordinal: 60000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Ensure entering the Nix development shell does not export Graphify's Nix Python dependency closure globally, so the repository's UV-managed Graphify interpreter remains independent and can import its compatible NumPy build.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [x] #1 The default development shell does not expose Graphify's propagated Python dependencies through global PYTHONPATH
- [x] #2 The graphify command remains available and executes through its Nix-compatible runtime environment
- [x] #3 The saved Graphify interpreter can import NumPy successfully inside nix develop
- [x] #4 Graphify clustering completes successfully inside nix develop
<!-- AC:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [x] #1 All acceptance criteria are satisfied
- [x] #2 All required unit tests pass. Documentation-only changes are exempt unless the user requests unit tests
- [x] #3 All required e2e tests pass. Documentation-only changes are exempt unless the user requests e2e tests
<!-- DOD:END -->

## Implementation Plan

<!-- SECTION:PLAN:BEGIN -->
1. Add a non-propagating graphify launcher package that delegates to the Nix-wrapped pkgs.graphify executable by absolute path.
2. Use that launcher in the default development shell instead of pkgs.graphify, retaining the existing explicit Python path for graphify-mcp.
3. Verify PYTHONPATH is absent from the development shell, the saved UV interpreter imports NumPy, and graphify clustering runs.
<!-- SECTION:PLAN:END -->

## Implementation Notes

<!-- SECTION:NOTES:BEGIN -->
Replaced the propagated package in the default shell with a lightweight launcher that delegates to the Nix-wrapped Graphify executable. The MCP app retains its explicit private Python path.

Verification: `nix flake check --no-build` passed; `nix develop --command sh -c 'test -z "${PYTHONPATH:-}" && "$(cat graphify-out/.graphify_python)" -c "import numpy; print(numpy.__version__)"'` exited successfully; `nix develop --command sh -c 'graphify cluster-only . --no-viz'` completed with 391 communities. `graphify update .` refreshed the required graph output. `nix fmt` could not run because this flake defines no formatter.
<!-- SECTION:NOTES:END -->

## Final Summary

<!-- SECTION:FINAL_SUMMARY:BEGIN -->
Replaced the direct `pkgs.graphify` shell package with a non-propagating launcher that invokes Graphify's existing Nix-wrapped executable by absolute path. This keeps the development shell free of Graphify's global Python closure while preserving Graphify's compatible runtime. The Graphify MCP app continues to construct its own explicit Python path. Verified flake evaluation, the saved interpreter's NumPy import with no PYTHONPATH, and Graphify clustering inside `nix develop`; refreshed graph metadata with `graphify update .`.
<!-- SECTION:FINAL_SUMMARY:END -->
