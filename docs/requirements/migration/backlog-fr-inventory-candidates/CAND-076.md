# CAND-076

#### Source

> - [x] When a timeslot is selected, it's position is saved.
>       When the page is reloaded, the selection is rendered at the timeslot.
>       However, there's no tooltip near the selection.

[Source lines 368-370](../../../../backlog/backlog.md#L368-L370)

#### Candidate behavior

No new requirement behavior asserted; this is a persisted-selection tooltip regression, and persistence itself is not established.

#### Applicability

Actor: mobile event visitor. Location: timed grid tooltip. Event kind: timed. Interaction mode: reload. Viewport: mobile. State: previously selected slot restored and visible. Exclusions: selection persistence policy.

#### Classification

needs product decision

#### Existing Requirements and Confidence

FR-020 requires an adjacent tooltip for a visible selected slot, but not selection persistence. Confidence: confirmed.

#### Disposition

Treat missing tooltip as FR-020 regression; retain persistence scope for decision.

#### Open Questions

Should selected slots persist across reloads?
