---
id: CAND-149
verdict: covered
related_requirements: [QR-005]
confidence: inferred
---

# CAND-149

### Source

> - [x] Fix gap on the right for dates-only events.
>   - Dates-only grid fits within narrow viewports from 320px through 639px.
>   - Grid has a 16px left and right gutter on phone layouts.
>   - No horizontal document overflow at 320, 390, 410, 480, 639, and 640px.
>   - Below 640px, the sidebar remains stacked beneath the calendar and fits the viewport.
>   - At 640px, grid and sidebar remain side-by-side with their intended 16–20px gap.
>   - The full-width dates-only grid no longer adds external margin beyond its pane.

[Source lines 506-512](../../../../backlog/backlog.md#L506-L512)

### Candidate behavior

On 320px through 639px viewports, a dates-only grid and its stacked sidebar fit without horizontal document overflow.

### Applicability

Actor: user.
Location: dates-only event page.
Event kind: dates-only.
Interaction mode: viewing.
Viewport: 320px through 640px.
State: responsive layout.
Exclusions: wider layouts except the stated 640px boundary.

### Classification

existing requirement

### Existing Requirements and Confidence

QR-005 requires WCAG 2.2 AA conformance for supported mobile event-view flows, which covers responsive reflow of the dates-only page.
Confidence: inferred.

### Disposition

Map to QR-005; retain the pixel gutters and 640px layout boundary only if separately needed.

### Open Questions

Is the 640px side-by-side boundary and 16–20px gap a durable product decision?
