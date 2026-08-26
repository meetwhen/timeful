---
id: CAND-192
verdict: proposed-requirement
requirement_type: FR
related_requirements:
  - FR-049
  - FR-108
confidence: inferred
---

# CAND-192: Mobile Timezone-Control Width

### Source

> - [x] On mobile, on the timed event page, the timezone control keeps a fixed width (112px) so the reset button fits inside without resizing the control, with the time format and days-per-page buttons on either side

[Source lines 570-570](../../../../backlog/backlog.md#L570)

### Candidate behavior

On mobile timed event pages, the timezone control remains 112px wide so its reset button fits without resizing the control, between the time-format and days-per-page buttons.

### Applicability

- Actor: event visitor
- Location: timed event page controls
- Event kind: timed
- Interaction mode: viewing
- Viewport: mobile
- State: timezone reset button available
- Exclusions: desktop layout and dates-only pages

### Classification

candidate FR

### Existing Requirements and Confidence

Existing requirements: [FR-049](../../functional/fr/FR-049.md) is proposed and places the mobile controls in a row but does not specify width or order.

Confidence: inferred

### Disposition

Promoted to [FR-108](../../functional/fr/FR-108.md), phrased as a reserved-width outcome so the implementation can keep the reservation configurable.

### Open Questions

- What mobile width range and localization widths must the fixed control support?
