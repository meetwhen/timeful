---
id: CAND-177
verdict: excluded
related_requirements: [FR-009, FR-074]
confidence: confirmed
---

# CAND-177: Active-Slot Persistence Model

### Source

> - [x] Store only active timeslots and calculate all other types (enabled inactive and disabled) on the fly

[Source lines 547-547](../../../../backlog/backlog.md#L547)

### Candidate behavior

No new requirement behavior asserted; the source prescribes a persistence and derivation model rather than an observable outcome.

### Applicability

- Actor: system implementation
- Location: timed-event storage and grid derivation
- Event kind: timed
- Interaction mode: not applicable
- Viewport: not applicable
- State: persistence and rendering
- Exclusions: user-visible state semantics already specified elsewhere

### Classification

implementation detail

### Existing Requirements and Confidence

Existing requirements: [FR-009](../../functional/fr/FR-009.md) specifies visible grid states; [FR-074](../../functional/fr/FR-074.md) is proposed and specifies enabled-domain derivation.

Confidence: confirmed

### Disposition

Retain as design provenance; do not create an FR or QR from storage wording.
