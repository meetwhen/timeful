---
id: TASK-0052
title: Migrate session handoffs into the Backlog archive
status: Done
assignee:
  - '@OpenCode'
created_date: '2026-08-23 11:46'
updated_date: '2026-08-23 11:53'
labels:
  - handoff
  - workflow
dependencies: []
references:
  - scripts/handoff/create-handoff.sh
  - .agents/skills/handoff/SKILL.md
  - AGENTS.md
  - backlog/handoffs/
modified_files:
  - scripts/handoff/create-handoff.sh
  - .agents/skills/handoff/SKILL.md
  - AGENTS.md
  - backlog/handoffs/
priority: medium
type: docs
ordinal: 46000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Move the repository's session-handoff archive from handoff/ to the custom append-only backlog/handoffs/ directory. Add repository-owned generation tooling and expose the workflow through the /handoff skill without treating handoffs as Backlog-managed documents or task records.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [x] #1 The repository script creates a timestamped session-handoff template in backlog/handoffs
- [x] #2 The archive creation script refuses to overwrite an existing handoff file
- [x] #3 The public handoff skill is named handoff and uses the repository script
- [x] #4 All existing handoff markdown files are migrated to backlog/handoffs with navigational references updated
- [x] #5 The archive is documented as append-only and not Backlog-managed
- [x] #6 The prior handoff directory contains no migrated handoff markdown files
<!-- AC:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [x] #1 All acceptance criteria are satisfied
- [ ] #2 All required unit tests pass. Documentation-only changes are exempt unless the user requests unit tests
- [ ] #3 All required e2e tests pass. Documentation-only changes are exempt unless the user requests e2e tests
<!-- DOD:END -->

## Implementation Plan

<!-- SECTION:PLAN:BEGIN -->
1. Add an executable repository script that determines the latest OpenCode session and creates a unique, non-overwritable Markdown skeleton under backlog/handoffs.
2. Rename the session-handoff skill to handoff and make it invoke the script before completing the generated template with factual session context.
3. Move every handoff/handoff-*.md record into backlog/handoffs, changing only archive-navigation paths and the renamed skill references; leave historical statements and handoff/answers untouched.
4. Document that backlog/handoffs is a custom append-only archive, not a Backlog-managed document or task location.
5. Validate script output and overwrite protection, migrated archive count and references, shell syntax, Markdown formatting, and diff whitespace.
<!-- SECTION:PLAN:END -->

## Implementation Notes

<!-- SECTION:NOTES:BEGIN -->
Added an executable append-only creator that writes the standard handoff template, atomically reserves its timestamped filename with a hard link, and refuses collisions. The normal destination is backlog/handoffs; HANDOFF_ARCHIVE_DIR supports isolated validation without creating a production handoff. Renamed the public skill to handoff and removed its duplicate template. Migrated all 15 handoff Markdown files, updated archive links and skill references, and retained handoff/answers as untouched user content. Validation: concurrent isolated creator invocations produced exactly one template and one collision failure; bash -n, archive count/reference checks, Prettier checks for changed guidance, and git diff --check passed. The full historical archive was not reformatted to preserve its existing prose; five legacy files fail a whole-archive Prettier check.
<!-- SECTION:NOTES:END -->

## Final Summary

<!-- SECTION:FINAL_SUMMARY:BEGIN -->
Moved the session-handoff workflow into the repository-owned, append-only backlog/handoffs archive. Added scripts/handoff/create-handoff.sh to generate the timestamped session template with atomic no-overwrite behavior, renamed the skill to /handoff, and documented the custom archive boundary in AGENTS.md. Migrated all 15 existing handoffs and updated archive navigation and skill references; handoff/answers remains untouched. Verified concurrent isolated creation yields one archive entry and one collision failure, plus shell syntax, executable bit, archive integrity, Prettier for changed guidance, and git diff --check. Unit and e2e tests are exempt because this is documentation and shell tooling only; legacy archive formatting was intentionally preserved.
<!-- SECTION:FINAL_SUMMARY:END -->
