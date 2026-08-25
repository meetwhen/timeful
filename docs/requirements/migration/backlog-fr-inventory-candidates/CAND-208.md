# CAND-208: Follow Scheduling Drag With Slot Tooltip

#### Source

> When scheduling an event, the tooltip with the info about the time slot should follow the mouse cursor and not be above the slot where scheduling the event started

[Source FR-016](../../functional-requirements.md#fr-016)

#### Candidate behavior

While dragging to schedule a [Timed Event](../../../terminology/glossary.md#timed-event), the [Timed Slot](../../../terminology/glossary.md#timed-slot) tooltip follows the cursor rather than remaining at the drag-start slot.

#### Applicability

Actor: event visitor. Location: timed event-page grid. Event kind: timed. Interaction mode: scheduling drag. Viewport: pointer-capable. State: scheduling is active and the pointer is over a timed slot. Exclusions: availability editing, touch interactions, and a stationary pointer.

#### Classification

candidate FR

#### Existing Requirements and Confidence

Accepted FR-020 concerns mobile selected-slot tooltips, not desktop scheduling-drag tooltip tracking. Confidence: inferred.

#### Disposition

Hold as a timed scheduling interaction candidate.
