---
id: TASK-0070
title: Document proposed platform and event identity model
status: Done
assignee:
  - OpenCode
created_date: "2026-08-25 14:55"
updated_date: "2026-08-25 15:09"
labels:
  - requirements
  - terminology
  - identity
dependencies: []
references:
  - docs/requirements/functional/fr/FR-001.md
  - docs/requirements/functional/fr/FR-018.md
  - docs/requirements/functional/fr/FR-060.md
  - docs/requirements/functional/fr/FR-061.md
  - docs/requirements/functional/fr/FR-062.md
  - docs/requirements/functional/fr/FR-063.md
  - docs/requirements/functional/fr/FR-073.md
  - docs/design/architecture/adr/ADR-008.md
  - docs/design/architecture/adr/ADR-009.md
  - docs/terminology/glossary.md
  - docs/requirements/migration/backlog-fr-inventory-candidates/CAND-062.md
  - docs/requirements/migration/backlog-fr-inventory-candidates/CAND-132.md
  - docs/requirements/migration/backlog-fr-inventory-candidates/CAND-176.md
  - docs/requirements/migration/backlog-fr-inventory-candidates/CAND-183.md
  - docs/requirements/migration/backlog-fr-inventory-candidates/CAND-184.md
  - docs/requirements/migration/backlog-fr-inventory-candidates/CAND-200.md
  - docs/requirements/migration/backlog-fr-inventory-candidates/CAND-206.md
documentation:
  - docs/requirements/README.md
  - docs/requirements/functional/README.md
  - docs/requirements/migration/README.md
  - docs/terminology/README.md
modified_files:
  - docs/requirements/README.md
  - docs/requirements/functional/fr/FR-001.md
  - docs/requirements/functional/fr/FR-018.md
  - docs/requirements/functional/fr/FR-060.md
  - docs/requirements/functional/fr/FR-061.md
  - docs/requirements/functional/fr/FR-062.md
  - docs/requirements/functional/fr/FR-063.md
  - docs/requirements/functional/fr/FR-073.md
  - docs/requirements/functional/fr/FR-079.md
  - docs/requirements/functional/fr/FR-080.md
  - docs/terminology/glossary.md
  - docs/design/architecture/adr/ADR-008.md
  - docs/design/architecture/adr/ADR-009.md
  - docs/requirements/migration/backlog-fr-inventory-candidates/CAND-062.md
  - docs/requirements/migration/backlog-fr-inventory-candidates/CAND-132.md
  - docs/requirements/migration/backlog-fr-inventory-candidates/CAND-176.md
  - docs/requirements/migration/backlog-fr-inventory-candidates/CAND-183.md
  - docs/requirements/migration/backlog-fr-inventory-candidates/CAND-184.md
  - docs/requirements/migration/backlog-fr-inventory-candidates/CAND-200.md
  - backlog/handoffs/handoff-2026-08-25T15-08-00Z.md
priority: medium
type: docs
ordinal: 72000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->

Define the approved Platform Identity, Event Visitor Identity, event-owner, event-guest, response-ownership, and signed-in response-name semantics in canonical requirements, terminology, architecture decisions, and the migration candidate inventory. Keep all new and revised functional requirements in proposed status; this task documents product behavior and does not implement runtime identity changes.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria

<!-- AC:BEGIN -->

- [x] #1 Canonical requirements define Platform Identity and Event Visitor Identity without using raw Platform Identity values as event or response identifiers
- [x] #2 All new and revised functional requirements remain proposed
- [x] #3 Requirements define that event owners are event guests with additional event-settings authority
- [x] #4 Requirements define multiple response ownership by an Event Visitor Identity independently of response display names
- [x] #5 Requirements define signed-in response-name defaults and preserve existing response names after account-profile name changes
- [x] #6 Requirements define cross-device owner and response recovery with credential proof for anonymous identity association
- [x] #7 The glossary and relevant ADRs distinguish Platform Identity Event Visitor Identity event ownership and response credentials
- [x] #8 CAND-062 CAND-132 CAND-176 CAND-184 CAND-200 CAND-183 and CAND-206 reflect the resolved or remaining triage outcomes

<!-- AC:END -->

## Definition of Done

<!-- DOD:BEGIN -->

- [x] #1 All acceptance criteria are satisfied
- [x] #2 All required unit tests pass. Documentation-only changes are exempt unless the user requests unit tests
- [x] #3 All required e2e tests pass. Documentation-only changes are exempt unless the user requests e2e tests

<!-- DOD:END -->

## Implementation Plan

<!-- SECTION:PLAN:BEGIN -->

1. Complete FR-080 so it explicitly preserves existing response display names when the signed-in account profile changes.
2. Restore CAND-132 and CAND-176 as proposed requirements because their dates-only Edit event trigger is not covered by FR-018 or FR-039.
3. Correct TASK-0070 metadata and final summary to distinguish CAND-206's retained disposition from edited candidates.
4. Create an append-only successor handoff recording that the prior handoff predates documentation completion, then run Markdown formatting and staged-diff checks.

<!-- SECTION:PLAN:END -->

## Implementation Notes

<!-- SECTION:NOTES:BEGIN -->

Confirmed the task has no dependencies and remains documentation-only. Current records show the owner/guest glossary conflict, token-only owner authority wording, and legacy `guestId` wording that must not be presented as the future Event Visitor Identity model.

Validated the edited Markdown with Prettier and `git diff --check`. No unit or e2e tests were run because this task changes documentation only; the task's documented exemption applies.

Review corrections: completed FR-080's profile-name-change preservation statement. Restored CAND-132 and CAND-176 as proposed requirements because FR-018 and FR-039 do not define the dates-only Edit event entry point. CAND-206 was retained, not modified. Created a successor handoff to clarify the archived pre-completion handoff.
<!-- SECTION:NOTES:END -->

## Final Summary

<!-- SECTION:FINAL_SUMMARY:BEGIN -->

Documented the proposed Platform Identity and Event Visitor Identity model without runtime changes. Added proposed FR-079 for platform/event-visitor association and FR-080 for signed-in response-name defaults, including preservation of existing response display names after Platform Identity profile-name changes. Revised the proposed ownership, protection, and recovery requirements, and corrected glossary and ADR distinctions between the proposed model and legacy `guestId` handling. Candidate dispositions: CAND-062 is covered; CAND-132 and CAND-176 remain proposed requirements because their dates-only Edit event entry behavior is not covered by FR-018 or FR-039; CAND-184 and CAND-200 are excluded presentation details; CAND-183 and the retained CAND-206 remain needs-decision. Created a successor handoff that supersedes the earlier handoff's pre-completion current-state guidance. Validation: `npx prettier --check` for corrected Markdown and `git diff HEAD --check`. Unit and e2e tests were exempt for documentation-only work.
<!-- SECTION:FINAL_SUMMARY:END -->
