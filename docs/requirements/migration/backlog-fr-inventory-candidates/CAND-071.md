# CAND-071

#### Source

> - [x] on mobile, I shouldn't be able to click the grid through that Responses offcanvas

[Source lines 362-362](../../../../backlog/backlog.md#L362-L362)

#### Candidate behavior

On mobile, an open Responses offcanvas should prevent interactions with the grid behind it.

#### Applicability

Actor: mobile event visitor. Location: event page. Event kind: timed. Interaction mode: touch or click. Viewport: mobile. State: Responses offcanvas open. Exclusions: offcanvas closed.

#### Classification

needs product decision

#### Existing Requirements and Confidence

FR-028 concerns the selected-slot panel while editing but does not define interaction blocking. Confidence: inferred.

#### Disposition

Candidate for consolidation if the offcanvas scope is confirmed.

#### Open Questions

Does blocking apply to all background controls, keyboard actions, and scroll?
