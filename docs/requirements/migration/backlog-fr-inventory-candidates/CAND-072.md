# CAND-072

#### Source

> - [x] on mobile, the tooltip should stay near selected timeslot and shouldn't move around
>   - now we store the info about the selected timeslot and render based on its position

[Source lines 363-364](../../../../backlog/backlog.md#L363-L364)

#### Candidate behavior

No new requirement behavior asserted; the child supplies an implementation approach for an already accepted adjacent-tooltip outcome.

#### Applicability

Actor: mobile event visitor. Location: timed grid tooltip. Event kind: timed. Interaction mode: slot selection. Viewport: mobile. State: slot selected and visible. Exclusions: no selected slot.

#### Classification

needs product decision

#### Existing Requirements and Confidence

FR-020 requires the visible selected-slot tooltip adjacent to that slot. Confidence: confirmed.

#### Disposition

Map to FR-020; exclude stored-position mechanism.

#### Open Questions

None.
