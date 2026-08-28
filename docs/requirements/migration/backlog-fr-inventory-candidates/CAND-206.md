---
id: CAND-206
title: Edit-Availability Action for Non-Editable Responses
verdict: needs-decision
related_requirements: []
confidence: needs-product-decision
---

# CAND-206: Edit-Availability Action for Non-Editable Responses

## Source

> When there are responses but no responses to edit, the user should see disabled Edit availability button.
>
> when a user doesn't have responses to edit, the edit availability button shouldn't be visible (or should be disabled - is it better UX?)

## Candidate behavior

No new requirement behavior asserted; the source identifies an unresolved choice between hiding and disabling `Edit availability` when the current visitor has no editable [Event Response](../../../terminology/glossary.md#event-response).

## Applicability

Actor: event visitor.
Location: event page.
Event kind: unspecified.
Interaction mode: viewing.
Viewport: unspecified.
State: one or more responses exist but none is editable by the current visitor.
Exclusions: zero-response behavior and add-availability visibility.

## Classification

needs product decision

## Existing Requirements and Confidence

No accepted or proposed FR or QR resolves this editable-response state.
Confidence: needs product decision.

## Disposition

Needs product decision; do not consolidate until the hidden-versus-disabled outcome is chosen.

## Open Questions

Should `Edit availability` be hidden or disabled, and if disabled, what explanation is shown?
