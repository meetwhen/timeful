---
id: CAND-207
title: Derive Enabled Domain From Picked Dates
verdict: covered
related_requirements:
  - FR-074
confidence: confirmed
---

# CAND-207: Derive Enabled Domain From Picked Dates

## Source

> Dates picked in the date picker shall be the source of truth for enabled time slots.
>
> picked dates define the enabled domain

## Candidate behavior

For a [Timed Event](../../../terminology/glossary.md#timed-event), [Event Picked Dates](../../../terminology/glossary.md#event-picked-dates) define the [Enabled Domain](../../../terminology/glossary.md#enabled-domain).

## Applicability

Actor: event owner.
Location: timed event date picker and enabled-slot domain.
Event kind: timed.
Interaction mode: selecting dates or editing event settings.
Viewport: unspecified.
State: picked dates are present.
Exclusions: active-slot selection within the enabled domain and dates-only events.

## Classification

existing requirement

## Existing Requirements and Confidence

Proposed [FR-074](../../functional/fr/FR-074.md) directly derives enabled slots for each picked date.
Confidence: confirmed.

## Disposition

Map to FR-074; retain as temporary-source provenance.
