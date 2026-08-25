---
id: CAND-208
title: Follow Scheduling Drag With Slot Tooltip
verdict: covered
related_requirements:
  - FR-043
confidence: confirmed
---

# CAND-208: Follow Scheduling Drag With Slot Tooltip

## Source

> When scheduling an event, the tooltip with the info about the time slot should follow the mouse cursor and not be above the slot where scheduling the event started

## Candidate behavior

While dragging to schedule a [Timed Event](../../../terminology/glossary.md#timed-event), the [Timed Slot](../../../terminology/glossary.md#timed-slot) tooltip follows the cursor rather than remaining at the drag-start slot.

## Applicability

Actor: event visitor. Location: timed event-page grid. Event kind: timed. Interaction mode: scheduling drag. Viewport: pointer-capable. State: scheduling is active and the pointer is over a timed slot. Exclusions: availability editing, touch interactions, and a stationary pointer.

## Classification

existing requirement

## Existing Requirements and Confidence

Proposed FR-043 directly requires a Timed Slot tooltip at the location where the user hovers or releases a grid selection. Confidence: confirmed.

## Disposition

Map to FR-043; retain as temporary-source provenance.
