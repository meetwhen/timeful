---
id: CAND-215
title: More-Options Toggle Aggregation
verdict: proposed-requirement
requirement_type: FR
related_requirements:
  - FR-107
confidence: confirmed
---

# CAND-215: More-Options Toggle Aggregation

## Source

> The More options button is used only when more than one hideable toggle exists; otherwise those toggles render directly instead of the button.

Recorded from the session triage decision round of 2026-08-26; this record has no retained backlog source, so the quoted ruling is its provenance.

## Candidate behavior

The event page shows the `More options` button only when more than one hideable toggle exists there; with a single hideable toggle, that toggle renders directly instead of the button.

## Applicability

Actor: event visitor.
Location: event page options area.
Event kind: timed and dates-only.
Interaction mode: viewing.
Viewport: any.
State: hideable toggles present near `More options`.
Exclusions: contexts with no hideable toggles.

## Classification

candidate FR

## Existing Requirements and Confidence

No accepted or proposed FR specifies the aggregation rule; the sibling orders FR-105 and FR-106 sit under it.
Confidence: confirmed.

## Disposition

Promoted to [FR-107](../../functional/fr/FR-107.md); the sibling orders FR-105 and FR-106 sit under it.

## Open Questions

None.
