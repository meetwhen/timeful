---
id: CAND-157
verdict: needs-decision
related_requirements: [FR-013, FR-057]
confidence: needs-product-decision
---

# CAND-157

### Source

> - [x] On the specific times page <http://127.0.0.1:4173/e/Eb67A>, when I switch timezone from +5 to +6, the left-most upper-most enabled slot should become disabled

[Source lines 525-525](../../../../backlog/backlog.md#L525-L525)

### Candidate behavior

After changing the timezone from +5 to +6 on the stated specific-times page, the left-most upper-most enabled slot becomes disabled.

### Applicability

Actor: user.
Location: specific-times page.
Event kind: timed.
Interaction mode: changing timezone.
Viewport: unspecified.
State: transition from +5 to +6.
Exclusions: other transitions and event configurations.

### Classification

needs product decision

### Existing Requirements and Confidence

FR-013 applies to display-timezone changes and FR-057 applies to event-timezone changes, but the source does not identify which timezone changes.
Confidence: needs product decision.

### Disposition

Do not consolidate or map this behavior until the changed timezone and initial event configuration are identified.

### Open Questions

Is the changed timezone the display timezone, and what initial slot domain produces this result?
