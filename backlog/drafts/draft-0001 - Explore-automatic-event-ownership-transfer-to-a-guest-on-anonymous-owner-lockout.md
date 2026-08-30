---
id: DRAFT-0001
title: >-
  Explore automatic event-ownership transfer to a guest on anonymous-owner
  lockout
status: Draft
assignee: []
created_date: '2026-08-30 20:06'
labels: []
dependencies:
  - TASK-0117
type: spike
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
The event ownership model has an accepted lockout: if an anonymous Event Owner's browser-local data is cleared, the Event Owner Edit Token is lost and the event becomes permanently unmanageable. TASK-0117 documents this as accepted current behavior. The product owner wants a design exploration of a recovery option: automatically transferring event ownership to one of the event's guests when the owner is locked out.

Open questions the exploration must address:
- Guest-selection rule: which guest becomes the owner (first respondent, most active, oldest response, or none)? Can the rule be deterministic and abuse-resistant, given guests are effectively anonymous Event Visitor Identities?
- Security: promotion gives the chosen guest full Event Settings control, including visibility changes that matter under Blind Availability Mode (FR-084). Should promotion require explicit consent from the promoted guest, or happen without their knowledge?
- Interaction with FR-063: ownership binds to exactly one Platform Visitor Identity and moves when the Event Owner Edit Token is proven by another PVI. Define how automatic promotion coexists with a later token proof.
- Applicability: PostgreSQL events only, or also legacy MongoDB events?

This is a spike: research and a recommendation only, no code or requirements changes. Record findings in the task's implementation notes or a linked document, and end with a concrete recommendation.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [ ] #1 A written exploration document proposes whether and how anonymous Event Owner authority could transfer automatically to a guest when the owner's browser-local data is cleared, or recommends against it
- [ ] #2 The exploration evaluates guest-selection rules (for example, first respondent, most active guest, or no selection) and states the trade-offs of each option
- [ ] #3 The exploration evaluates the security consequence that a promoted guest gains Event Settings control, including whether explicit guest consent is required before promotion
- [ ] #4 The exploration states how any proposal interacts with single-Platform Visitor Identity ownership binding and owner-association rules in FR-063
- [ ] #5 The exploration ends with a concrete recommendation and, if positive, the outline of a future requirements task; no implementation changes are made in this task
<!-- AC:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [ ] #1 All acceptance criteria are satisfied
- [ ] #2 All required unit tests pass. Documentation-only changes are exempt unless the user requests unit tests
- [ ] #3 All required e2e tests pass. Documentation-only changes are exempt unless the user requests e2e tests
- [ ] #4 Changed Markdown files are formatted with npm run format:markdown
<!-- DOD:END -->
