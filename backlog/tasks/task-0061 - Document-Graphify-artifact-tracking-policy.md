---
id: TASK-0061
title: Document Graphify artifact tracking policy
status: Done
assignee:
  - OpenCode
created_date: '2026-08-24 20:02'
updated_date: '2026-08-24 20:03'
labels: []
dependencies: []
references:
  - backlog/handoffs/handoff-2026-08-24T19-58-59Z.md
  - .gitignore
modified_files:
  - docs/graphify.md
priority: low
type: docs
ordinal: 63000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Record the repository policy for Graphify-derived outputs so collaborators understand which artifacts are shared immediately and which are regenerated locally. Community labels and membership signatures must remain untracked to avoid post-commit worktree churn; the semantic index is tracked for immediate collaborator availability; syntactic output remains untracked because it rebuilds quickly.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [x] #1 `docs/graphify.md` documents the policy for Graphify labels and membership signatures
- [x] #2 The document states that the semantic index is tracked for immediate collaborator use
- [x] #3 The document states that syntactic output is untracked because it can be indexed quickly
- [x] #4 The documented policy is consistent with `.gitignore`
<!-- AC:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [x] #1 All acceptance criteria are satisfied
- [x] #2 All required unit tests pass. Documentation-only changes are exempt unless the user requests unit tests
- [x] #3 All required e2e tests pass. Documentation-only changes are exempt unless the user requests e2e tests
<!-- DOD:END -->

## Implementation Plan

<!-- SECTION:PLAN:BEGIN -->
1. Add `docs/graphify.md` describing which Graphify artifacts are tracked and why.
2. State that regenerated labels and membership signatures remain untracked because their post-commit churn is not useful source history.
3. State that the semantic cache is tracked for immediate collaborator access, while syntactic output remains untracked because it is quick to rebuild.
4. Verify the document matches the `graphify-out/*` exceptions in `.gitignore`.
<!-- SECTION:PLAN:END -->

## Implementation Notes

<!-- SECTION:NOTES:BEGIN -->
Added `docs/graphify.md` and verified the ignore policy with `git check-ignore -v`. Labels, signatures, and syntactic cache output are ignored; the semantic cache remains unignored. `git diff --check` completed without output.
<!-- SECTION:NOTES:END -->

## Final Summary

<!-- SECTION:FINAL_SUMMARY:BEGIN -->
Added `docs/graphify.md` to document the Graphify artifact policy. The document explains why regenerated community labels and membership signatures remain untracked, why the semantic index is shared for immediate collaborator use, and why syntactic output is rebuilt locally. Verified the policy against `.gitignore` with `git check-ignore -v`; documentation-only unit and e2e tests are exempt.
<!-- SECTION:FINAL_SUMMARY:END -->
