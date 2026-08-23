---
id: TASK-0048
title: Align functional requirements with controlled terminology
status: Done
assignee:
  - OpenCode
created_date: '2026-08-22 20:27'
updated_date: '2026-08-22 20:46'
labels:
  - requirements
  - terminology
dependencies: []
references:
  - docs/requirements/README.md
  - docs/terminology/glossary.md
priority: high
type: docs
ordinal: 41000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Align the canonical functional requirements with the agreed event, availability, ownership, scheduling, and timed-domain terminology so the documentation is consistent and verifiable.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [x] #1 The glossary defines the agreed controlled terms and retired terminology is removed or redirected
- [x] #2 Functional requirements use and link controlled terminology at each required first occurrence
- [x] #3 FR-059 permits an Event Guest to own multiple Availability Responses
- [x] #4 FR-012 defines optional Scheduled Event Time behavior for timed and dates-only events
- [x] #5 FR-067 through FR-069 define approved custom-domain creation range updates and overlap-state behavior
- [x] #6 Requirement titles front matter and index rows remain synchronized
- [x] #7 Glossary anchors and requirement links validate successfully
<!-- AC:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [x] #1 All acceptance criteria are satisfied
- [ ] #2 All unit tests pass
- [ ] #3 All e2e tests pass
<!-- DOD:END -->

## Implementation Plan

<!-- SECTION:PLAN:BEGIN -->
1. Replace and add glossary entries for the approved event, ownership, availability, scheduling, and timed-domain vocabulary while preserving authoritative context links.
2. Update affected FR statements, acceptance criteria, titles, and README index entries; correct FR-059 response ownership and extend FR-012.
3. Add FR-067 through FR-069 for custom-domain creation, range-update regeneration, and overlap-state treatment.
4. Validate glossary anchors, relative links, front matter/title/index consistency, and required first-occurrence terminology links.
<!-- SECTION:PLAN:END -->

## Implementation Notes

<!-- SECTION:NOTES:BEGIN -->
Reconciled terminology definitions and requirement wording with the approved Event Kind, Timed Domain Mode, ownership, response, scheduling, and overlap models. Validation scripts confirmed glossary anchors, local documentation targets, and front matter/H1/index title consistency; a final terminology audit found no remaining legacy terms. Frontend unit and E2E suites were not run because this task changes documentation only.
<!-- SECTION:NOTES:END -->

## Final Summary

<!-- SECTION:FINAL_SUMMARY:BEGIN -->
Updated the glossary and 66 existing functional requirements to use the controlled event, timed-domain, ownership, availability, and scheduling terminology. Added FR-067 through FR-069 for custom-domain creation, active-slot regeneration after range changes, and availability-state overlap calculation. Corrected FR-059 so an Event Guest may own multiple Availability Responses while retaining the duplicate-display-name restriction. Verified glossary anchors, local requirement/glossary links, title synchronization, and clean documentation diffs. Frontend unit and E2E suites were not run because no application code changed.
<!-- SECTION:FINAL_SUMMARY:END -->
