---
id: CAND-188
verdict: excluded
related_requirements: []
confidence: confirmed
---

# CAND-188: Scrolling No-Responses Position

### Source

> - [x] Bug: when scrolling the gri, the No responses yet changes the position
>   - Can't reproduce

[Source lines 563-564](../../../../backlog/backlog.md#L563-L564)

### Candidate behavior

No new requirement behavior asserted; the reported defect could not be reproduced and gives no expected stable position.

### Applicability

- Actor: event visitor
- Location: grid and no-responses message
- Event kind: unspecified
- Interaction mode: scrolling
- Viewport: unspecified
- State: no responses
- Exclusions: reproduced layout defects

### Classification

bug or investigation

### Existing Requirements and Confidence

Existing requirements: None identified.

Confidence: confirmed

### Disposition

Retain as non-reproduced investigation history; no FR or QR.

### Open Questions

- What grid position and scroll sequence originally exposed the issue?
