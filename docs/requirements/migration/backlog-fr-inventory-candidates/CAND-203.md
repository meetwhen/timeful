---
id: CAND-203
title: Require Scheduled Timed Event Slots
verdict: proposed-requirement
requirement_type: FR
related_requirements:
  - FR-012
confidence: inferred
---

# CAND-203: Require Scheduled Timed Event Slots

## Source

> When scheduling an event, it can't be empty.
>
> a scheduled event time shall have at least one slot

## Candidate behavior

A [Scheduled Event Time](../../../terminology/glossary.md#scheduled-event-time) for a [Timed Event](../../../terminology/glossary.md#timed-event) shall contain at least one [Timed Slot](../../../terminology/glossary.md#timed-slot).

## Applicability

Actor: event visitor.
Location: timed event-page scheduling controls.
Event kind: timed.
Interaction mode: scheduling or rescheduling.
Viewport: unspecified.
State: saving a scheduled event time.
Exclusions: clearing a scheduled event time and dates-only event scheduling.

## Classification

candidate FR

## Existing Requirements and Confidence

Accepted [FR-012](../../functional/fr/FR-012.md) permits saving a scheduled event time but does not require a non-empty timed range.
Confidence: inferred.

## Disposition

Proposed requirement related to FR-012; it adds a non-empty timed-range validation boundary.

## Open Questions

Does this prohibit saving any zero-duration timed range, including one produced by a direct API request?
