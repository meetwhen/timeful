# CAND-196: Timed Sidebar Time-Control Order

### Source

> - [x] On the timed event page, the time format shall be on the left from the time zone so that the timezone button can get a reset button or stay the same freely

[Source lines 574-574](../../../../backlog/backlog.md#L574)

### Candidate behavior

On the timed event page, the time-format button is left of the timezone button so the timezone button can accommodate a reset button or remain unchanged.

### Applicability

- Actor: event visitor
- Location: timed event page sidebar
- Event kind: timed
- Interaction mode: viewing
- Viewport: unspecified
- State: time controls visible
- Exclusions: row spanning and alignment, dates-only pages, and reset-control behavior

### Classification

candidate FR

### Existing Requirements and Confidence

Existing requirements: [FR-046](../../functional/fr/FR-046.md) and [FR-047](../../functional/fr/FR-047.md) are proposed label requirements, not control order. Overlap: CAND-195 specifies the separate shared-row and alignment behavior; no accepted FR/QR overlap.

Confidence: inferred

### Disposition

Consolidate with timed-page sidebar layout if the order is durable across responsive layouts.

### Open Questions

- Does this order apply when the timezone reset control is absent?
