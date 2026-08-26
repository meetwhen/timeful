---
id: CAND-158
verdict: covered
related_requirements: [FR-002, FR-013]
confidence: confirmed
---

# CAND-158

### Source

> - [x] On the specific times page <http://127.0.0.1:4173/e/Eb67A>, when I switch timezone from +5 to +4 and on june 14, 0-4 are selected, jun 13 should appear

[Source lines 526-526](../../../../backlog/backlog.md#L526-L526)

### Candidate behavior

After changing the timezone from +5 to +4 with June 14 0-4 selected, June 13 appears on the stated specific-times page.

### Applicability

Actor: user.
Location: specific-times page.
Event kind: timed.
Interaction mode: changing timezone.
Viewport: unspecified.
State: +5 to +4 with the stated selection.
Exclusions: other transitions and selections.

### Classification

existing requirement

### Existing Requirements and Confidence

FR-002 requires adjacent projected date columns when timed slots project across midnight, and FR-013 governs that projection after a display-timezone change.
Confidence: confirmed.

### Disposition

Map to FR-002 and FR-013 as a regression example.

### Open Questions

None.
