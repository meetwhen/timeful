# CAND-048

#### Source

> - [x] <http://127.0.0.1:4173/e/Eb67A> shown in GMT -7 shows blank grey columns

[Source lines 333-333](../../../../backlog/backlog.md#L333-L333)

#### Candidate behavior

No new requirement behavior asserted; this is a timezone-rendering defect report.

#### Applicability

Actor: event visitor. Location: event-page timed grid. Event kind: timed. Interaction mode: display-timezone change. Viewport: any. State: GMT -7 projection. Exclusions: event-timezone mutation.

#### Classification

needs product decision

#### Existing Requirements and Confidence

FR-002, FR-009, FR-013, and FR-014 are relevant. Confidence: inferred.

#### Disposition

Treat as regression provenance for those requirements.

#### Open Questions

What cell-state relationship is expected for blank projected columns?
