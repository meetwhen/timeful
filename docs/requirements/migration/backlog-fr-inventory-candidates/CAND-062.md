---
id: CAND-062
verdict: covered
requirement_type: null
related_requirements:
  - FR-001
confidence: confirmed
---

# CAND-062

## Source

<!-- prettier-ignore -->
> - [x] remove Add availability, leave just edit availability
>

No.
In this case, we won't be able to add availability for someone

## Candidate behavior

An [Event Guest](../../../terminology/glossary.md#event-guest) may create an [Availability Response](../../../terminology/glossary.md#availability-response) labeled for another person; the creating [Event Visitor Identity](../../../terminology/glossary.md#event-visitor-identity) owns the response independently of its display name.

## Applicability

Actor: event guest.
Location: event page.
Event kind: any.
Interaction mode: availability creation.
Viewport: any.
State: adding for another person.
Exclusions: editing an existing response.

## Classification

existing requirement

## Existing Requirements and Confidence

Proposed FR-001 directly covers creating and owning multiple availability responses independently of their display names.
Confidence: confirmed.

## Disposition

Covered by proposed FR-001; retain as provenance for the response display-name boundary.

## Open Questions

None.
