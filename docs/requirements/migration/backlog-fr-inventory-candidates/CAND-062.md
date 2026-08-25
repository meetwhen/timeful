---
id: CAND-062
verdict: needs-decision
requirement_type: null
related_requirements:
  - FR-001
confidence: needs-product-decision
---

# CAND-062

## Source

<!-- prettier-ignore -->
> - [x] remove Add availability, leave just edit availability
>    No. In this case, we won't be able to add availability for someone

## Candidate behavior

No new requirement behavior asserted; the child rejects removal based on a use case without defining authorization.

## Applicability

Actor: unconfirmed. Location: event page. Event kind: any. Interaction mode: availability creation. Viewport: any. State: adding for another person. Exclusions: editing existing response.

## Classification

needs product decision

## Existing Requirements and Confidence

FR-001 permits an Event Guest to own multiple Availability Responses; no canonical requirement defines who may create a response for someone else. Confidence: needs product decision.

## Disposition

Do not migrate pending actor and access-policy confirmation.

## Open Questions

Who may add availability for someone else, and how is that response owned?
