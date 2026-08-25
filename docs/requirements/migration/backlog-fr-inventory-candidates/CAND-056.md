---
id: CAND-056
verdict: proposed-requirement
requirement_type: FR
related_requirements:
  - FR-009
confidence: confirmed
---

# CAND-056

## Source

> - [x] legend should be visible even if no responses. Show only enabled/active

## Candidate behavior

When a timed event has no Availability Responses, its event-page legend remains visible and shows only enabled or active states.

## Applicability

Actor: event visitor. Location: event-page legend. Event kind: timed. Interaction mode: viewing. Viewport: any. State: no responses. Exclusions: response-derived legend states.

## Classification

candidate FR

## Existing Requirements and Confidence

FR-009 is related but does not require a legend when no responses exist. Confidence: confirmed.

## Disposition

Retain as a candidate FR for no-response legend visibility and state scope.

## Open Questions

Must the legend be visible in every no-response viewport and grid mode?
