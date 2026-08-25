---
id: CAND-210
title: Label Whole-Hour Timed-Grid Lines
verdict: proposed-requirement
requirement_type: FR
related_requirements:
  - FR-011
confidence: inferred
---

# CAND-210: Label Whole-Hour Timed-Grid Lines

## Source

> Each full-hour line in the grid should have a label on the left, including top of the collapsed hours rectangle.
>
> each full-hour horizontal line in the grid shall be labelled with the hour on the left

## Candidate behavior

Each whole-hour horizontal line in a timed grid shall display its hour label on the left, including a line at the top of a collapsed-hours strip.

## Applicability

Actor: event visitor. Location: timed grid. Event kind: timed. Interaction mode: viewing or scheduling. Viewport: unspecified. State: whole-hour lines are visible, including collapsed hours. Exclusions: half-hour lines and dates-only event grids.

## Classification

candidate FR

## Existing Requirements and Confidence

Accepted FR-011 requires each collapsed band's start boundary on the time axis but does not require labels on every whole-hour line. Confidence: inferred.

## Disposition

Proposed requirement related to FR-011; it adds labels for every whole-hour line.
