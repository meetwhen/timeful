---
id: CAND-111
verdict: proposed-requirement
requirement_type: FR
related_requirements: [FR-020]
confidence: inferred
---

# CAND-111

## Source

> - [x] On mobile, when I click a disabled timeslot, the selection should disappear

[Source lines 443-443](../../../../backlog/backlog.md#L443-L443)

## Candidate behavior

On mobile, selecting a disabled timeslot clears the existing selection.

## Applicability

Actor: event visitor; Location: mobile timed grid; Event kind: timed; Interaction mode: tap; Viewport: mobile; State: disabled timeslot; Exclusions: enabled-inactive cells not explicit.

## Classification

candidate FR

## Existing Requirements and Confidence

Overlap: proposed FR-020 requires no tooltip when no Timed Slot is selected but does not require a disabled-timeslot tap to clear selection.
Confidence: inferred.

## Disposition

Review with CAND-104 and CAND-112.

## Open Questions

Should tap show no replacement selection immediately, and does this include padding cells?
