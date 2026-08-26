---
id: CAND-147
verdict: proposed-requirement
requirement_type: FR
related_requirements: [FR-009, FR-101, FR-102]
confidence: inferred
---

# CAND-147

### Source

> - [x] The dates-only legend now shows Unavailable, change in Add/Edit availability, matching timed events.

[Source lines 496-496](../../../../backlog/backlog.md#L496-L496)

### Candidate behavior

The dates-only legend shows `Unavailable, change in Add/Edit availability`.

### Applicability

Actor: respondent.
Location: dates-only event legend.
Event kind: dates-only.
Interaction mode: viewing availability.
Viewport: any.
State: unavailable date.
Exclusions: timed-event legend behavior already addressed by FR-009.

### Classification

candidate FR

### Existing Requirements and Confidence

FR-009 accepts the corresponding timed-grid label but does not state dates-only coverage.
Confidence: inferred.

### Disposition

Superseded by [FR-101](../../functional/fr/FR-101.md) and [FR-102](../../functional/fr/FR-102.md): the split summary wording replaces this record's legacy legend label approach.
Retained as provenance.

### Open Questions

Does the label apply in every dates-only availability mode?
