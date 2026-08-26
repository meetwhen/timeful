---
id: CAND-139
verdict: covered
related_requirements: [FR-023]
confidence: confirmed
---

# CAND-139

## Source

> - [x] Given I sign in, when I enter an unregistered email and click Continue with email:
>   - The input field is highlighted red
>   - The error appears like on accounts.google.com: (red alert icon) "Couldn’t find this account.
>     Create account"

[Source lines 479-481](../../../../backlog/backlog.md#L479-L481)

## Candidate behavior

An unregistered email submission highlights the input red and shows `Couldn’t find this account. Create account` with a red alert icon.

## Applicability

Actor: sign-in visitor; Location: Sign in page email flow; Event kind: not applicable; Interaction mode: submit unregistered email; Viewport: unspecified; State: email has no account; Exclusions: registered email and other providers.

## Classification

existing requirement

## Existing Requirements and Confidence

Overlap: accepted FR-023 directly covers the unregistered-email result and exact message.
Confidence: confirmed.

## Disposition

Map to FR-023; retain the red input and alert-icon presentation as provenance only.

## Open Questions

Is the exact accounts.google.com-like visual treatment a durable requirement, and does Create account initiate registration?
