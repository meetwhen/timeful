# CAND-066

#### Source

> - [x] When there's a scheduled event and I edit availability, I shold draw only with green cells, not blue cells (scheduled)

[Source lines 356-356](../../../../backlog/backlog.md#L356-L356)

#### Candidate behavior

No new requirement behavior asserted; this is a state-presentation refinement with colors as implementation language.

#### Applicability

Actor: availability editor. Location: availability editor grid. Event kind: timed. Interaction mode: editing availability. Viewport: any. State: scheduled event exists. Exclusions: viewing or scheduling mode.

#### Classification

needs product decision

#### Existing Requirements and Confidence

FR-009 requires context-specific visual states; FR-012 defines scheduled event time. Confidence: inferred.

#### Disposition

Do not migrate pending editing-state semantics.

#### Open Questions

Should scheduled time be hidden, visually distinct, or non-interactive during availability editing?
