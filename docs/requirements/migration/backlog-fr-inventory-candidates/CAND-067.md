# CAND-067

#### Source

> - [x] when editing availability, overlay availabilities should preserve the marked time slots and not shift

[Source lines 357-357](../../../../backlog/backlog.md#L357-L357)

#### Candidate behavior

The edited availability overlay should remain aligned with its marked timed slots during availability editing.

#### Applicability

Actor: availability editor. Location: availability editor grid. Event kind: timed. Interaction mode: editing availability. Viewport: any. State: overlay enabled. Exclusions: non-overlay rendering.

#### Classification

needs product decision

#### Existing Requirements and Confidence

FR-005 requires the edited response overlay; FR-013 preserves timed-slot instants under display-timezone changes. Confidence: inferred.

#### Disposition

Treat as a regression scenario for FR-005 and FR-013.

#### Open Questions

Does “not shift” specifically concern timezone changes, scrolling, or grid rerenders?
