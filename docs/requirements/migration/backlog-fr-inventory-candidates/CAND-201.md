---
id: CAND-201
title: Event-Page Action Visual Hierarchy
verdict: proposed-requirement
requirement_type: FR
related_requirements:
  - FR-112
confidence: inferred
---

# CAND-201: Event-Page Action Visual Hierarchy

## Source

> Button styles on the event page follow Material Design:
>
> - `Add availability`: filled primary, subtle shadow if needed for separation, no persistent glow
> - `Edit availability`: filled primary, calmer than add, no glow
> - `Edit event` and `Copy link`: outlined, no shadow

## Candidate behavior

The event page renders `Add availability` and `Edit availability` as filled primary actions, with `Add availability` visually more prominent, and renders `Edit event` and `Copy link` as outlined actions without shadows.

## Applicability

Actor: event visitor.
Location: event-page actions.
Event kind: unspecified.
Interaction mode: viewing.
Viewport: unspecified.
State: named actions visible.
Exclusions: actions not named by the source and transient interaction feedback.

## Classification

candidate FR

## Existing Requirements and Confidence

No accepted or proposed FR or QR specifies this action hierarchy.
Confidence: inferred.

## Disposition

Promoted to [FR-112](../../functional/fr/FR-112.md) with the owner's correction: only `Edit availability` is filled; `Add availability` is not filled.

## Open Questions

What observable criterion distinguishes `Add availability` as more prominent than `Edit availability`?
