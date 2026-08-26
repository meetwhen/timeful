---
id: CAND-146
verdict: proposed-requirement
requirement_type: FR
related_requirements: [FR-009, FR-045, FR-099]
confidence: inferred
---

# CAND-146

### Source

> - [x] Hovering a disabled date on a dates-only event now gives respondents a disabled status, not active/unavailable.

[Source lines 495-495](../../../../backlog/backlog.md#L495-L495)

### Candidate behavior

Hovering a disabled date on a dates-only event shows respondents a disabled status rather than active or unavailable.

### Applicability

Actor: respondent.
Location: dates-only event page.
Event kind: dates-only.
Interaction mode: pointer hover.
Viewport: pointer-capable.
State: disabled date.
Exclusions: enabled dates and touch-only interaction.

### Classification

candidate FR

### Existing Requirements and Confidence

FR-009 and FR-045 define analogous timed-grid state presentation and Responses treatment, but do not confirm dates-only hover status.
Confidence: inferred.

### Disposition

Promoted to [FR-099](../../functional/fr/FR-099.md), mirroring the timed-grid Disabled wording and treatment.

### Open Questions

What exact status text and response-view treatment are required?
