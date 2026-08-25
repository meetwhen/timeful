# CAND-128

#### Source

> - [x] On the read-only timed event page, when I switch the timezone, the grid shouldn't collapse

[Source lines 467-467](../../../../backlog/backlog.md#L467-L467)

#### Candidate behavior

Changing timezone on a read-only timed-event page does not collapse the grid.

#### Applicability

Actor: read-only event visitor; Location: timed-event page; Event kind: timed; Interaction mode: change timezone; Viewport: unspecified; State: read-only; Exclusions: editable pages.

#### Classification

candidate FR

#### Existing Requirements and Confidence

Overlap: accepted FR-013 preserves timed-slot instants across display-timezone changes; it does not specify grid collapse. Confidence: inferred.

#### Disposition

Review as display-timezone interaction behavior.

#### Open Questions

Which collapse state must be preserved, and does this apply after every timezone change?
