---
id: CAND-137
verdict: proposed-requirement
requirement_type: FR
related_requirements: [FR-007, FR-031]
confidence: inferred
---

# CAND-137

## Source

> - [x] When VITE_ENABLE_SIGN_IN is not false, on the landing page, there should be the Sign in button.
>       When I click that button, "Sign in" page opens with options like Google Calendar / Email.

[Source lines 476-477](../../../../backlog/backlog.md#L476-L477)

## Candidate behavior

When sign-in is enabled, the landing page shows Sign in, which opens a page with Google Calendar and Email options.

## Applicability

Actor: landing-page visitor; Location: landing page and Sign in page; Event kind: not applicable; Interaction mode: select Sign in; Viewport: unspecified; State: `VITE_ENABLE_SIGN_IN` is not false; Exclusions: sign-in disabled.

## Classification

candidate FR

## Existing Requirements and Confidence

Overlap: proposed FR-007 gates sign-in entry points when disabled; proposed FR-031 sends email magic links, but neither establishes this enabled-state entry point. Confidence: inferred.

## Disposition

Review as sign-in entry and provider-choice behavior.

## Open Questions

Are Google Calendar and Email the complete option set, and is this condition valid for every deployment?
