---
id: CAND-202
title: Preserve Event Timezone During Display-Timezone Changes
verdict: covered
related_requirements:
  - FR-013
confidence: confirmed
---

# CAND-202: Preserve Event Timezone During Display-Timezone Changes

## Source

> Setting "Shown in" shouldn't affect the initial event time zone.
>
> display time zone shan't affect the event time zone

## Candidate behavior

Changing the [Display Timezone](../../../terminology/glossary.md#display-timezone) shall not change the [Event Timezone](../../../terminology/glossary.md#event-timezone).

## Applicability

Actor: event visitor.
Location: timed event-page display controls.
Event kind: timed.
Interaction mode: display-timezone change.
Viewport: any.
State: a display timezone is selected.
Exclusions: explicit event-settings timezone edits.

## Classification

existing requirement

## Existing Requirements and Confidence

Accepted [FR-013](../../functional/fr/FR-013.md) directly specifies this behavior.
Confidence: confirmed.

## Disposition

Map to FR-013; retain as temporary-source provenance.
