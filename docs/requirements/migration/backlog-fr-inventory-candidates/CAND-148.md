---
id: CAND-148
verdict: proposed-requirement
requirement_type: FR
related_requirements: [FR-042, FR-045]
confidence: inferred
---

# CAND-148

### Source

> - [x] Fix highlighting in dates-only events
>   - Enabled dates-only grid cells receive a black inset square frame on hover.
>   - The frame stays inside the cell and does not overlap neighboring cells or the date number.
>   - The date text does not change from border-box layout shifts.
>   - Hovering a disabled cell removes the prior enabled-cell frame.
>   - Hovering a disabled cell shows disabled status in Responses.
>   - Leaving the grid removes the current frame.
>   - On mobile, tapping a disabled cell removes the current frame while showing disabled Responses status.
>   - On mobile, tapping outside the grid removes the current frame and clears its state.

[Source lines 497-505](../../../../backlog/backlog.md#L497-L505)

### Candidate behavior

Hovering an enabled dates-only grid cell displays an inset frame that remains inside that cell without shifting its date text.

### Applicability

Actor: respondent. Location: dates-only grid and Responses. Event kind: dates-only. Interaction mode: hover or mobile tap. Viewport: desktop and mobile. State: enabled or disabled cell, including leaving the grid. Exclusions: timed-grid highlights.

### Classification

candidate FR

### Existing Requirements and Confidence

FR-042 proposes analogous timed-grid containment and FR-045 proposes analogous Responses treatment, but neither confirms dates-only behavior. Confidence: inferred.

### Disposition

Hold as a compound interaction candidate; split visual containment from response-state behavior during consolidation.

### Open Questions

Are the black color and square frame durable, and what does “clears its state” include?
