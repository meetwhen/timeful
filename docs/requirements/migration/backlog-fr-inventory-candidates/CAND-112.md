# CAND-112

#### Source

> - [x] On desktop, given I'm on the event page, when I hover or click inside the grid outside active cells, the highlight at the last timeslot and any tooltip should not be visible

[Source lines 444-444](../../../../backlog/backlog.md#L444-L444)

#### Candidate behavior

On desktop, hovering or clicking outside active grid cells clears the prior highlight and hides its tooltip.

#### Applicability

Actor: event visitor; Location: desktop timed grid; Event kind: timed; Interaction mode: hover or click; Viewport: desktop; State: outside active cells; Exclusions: active cells.

#### Classification

duplicate or refinement

#### Existing Requirements and Confidence

Overlap: CAND-104 and CAND-111 specify the same visibility and clearing rule by cell state; no accepted FR/QR overlap. Confidence: inferred.

#### Disposition

Consolidate into one cross-viewport inactive-cell requirement.

#### Open Questions

Does `outside active cells` include inter-grid space and collapsed hours?
