# CAND-145

### Source

> - [x] Fix event description card
>   - Multi-line event descriptions save with newline characters preserved.
>   - Saved multi-line descriptions render across multiple lines.
>   - Read-only description pencil is in the upper-right corner.
>   - Edit mode preserves the description’s text width and wrapping.
>   - Edit mode cancel and save buttons are in the upper-right corner.
>   - Description action buttons are reduced to 32px and visually centered for single-line descriptions.

[Source lines 488-494](../../../../backlog/backlog.md#L488-L494)

### Candidate behavior

No new requirement behavior asserted for newline preservation because FR-019 already requires it; the remaining card layout details are an unconfirmed refinement.

### Applicability

Actor: event owner or reader. Location: event description card. Event kind: any. Interaction mode: read-only and edit. Viewport: any. State: multiline or single-line description. Exclusions: description persistence and rendering already covered by FR-019.

### Classification

duplicate or refinement

### Existing Requirements and Confidence

FR-019 accepts preservation and multiline rendering of event descriptions; it does not specify controls, dimensions, or placement. Confidence: confirmed.

### Disposition

Retain the layout details as a possible refinement separate from FR-019.

### Open Questions

Which actors can edit, and are the 32px and placement details durable product intent?
