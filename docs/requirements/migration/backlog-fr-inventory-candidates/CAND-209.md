---
id: CAND-209
title: Size Collapsed-Hours Strips
verdict: needs-decision
related_requirements:
  - FR-011
confidence: needs-product-decision
---

# CAND-209: Size Collapsed-Hours Strips

## Source

> Collapsed hours rectangle height should be the same as half-hour line.
>
> Collapsed hours strip height shall be 60% of a row between whole hours

## Candidate behavior

A collapsed-hours strip shall have a height equal to 60 percent of a timed-grid row between whole-hour lines.

## Applicability

Actor: event visitor.
Location: timed event-page grid.
Event kind: timed.
Interaction mode: viewing or scheduling.
Viewport: unspecified.
State: collapsed hours are visible.
Exclusions: expanded timed-grid rows and availability editing.

## Classification

needs product decision

## Existing Requirements and Confidence

Accepted FR-011 establishes collapsed inactive runs but does not specify strip height.
The source gives conflicting height rules.
Confidence: needs product decision.

## Disposition

Needs product decision; do not migrate until the collapsed-hours strip height is chosen.
