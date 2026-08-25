# CAND-195: Timed Sidebar Time-Control Row

### Source

> - [x] On the timed event page, timezone button and time format button shall together span the full row in the sidebar, the time format aligned right.

[Source lines 573-573](../../../../backlog/backlog.md#L573)

### Candidate behavior

On the timed event page, the timezone and time-format buttons together span the full sidebar row, with the time format aligned right.

### Applicability

- Actor: event visitor
- Location: timed event page sidebar
- Event kind: timed
- Interaction mode: viewing
- Viewport: unspecified
- State: time controls visible
- Exclusions: time-control order, dates-only pages, and reset-control behavior

### Classification

candidate FR

### Existing Requirements and Confidence

Existing requirements: [FR-046](../../functional/fr/FR-046.md) and [FR-047](../../functional/fr/FR-047.md) are proposed label requirements, not row layout. Overlap: CAND-196 specifies the separate time-control order; no accepted FR/QR overlap.

Confidence: inferred

### Disposition

Consolidate only after resolving the source's potentially ambiguous alignment wording.

### Open Questions

- Does “time format aligned right” mean its text, button, or allocated area?
