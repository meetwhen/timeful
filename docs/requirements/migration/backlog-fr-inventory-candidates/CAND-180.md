---
id: CAND-180
verdict: excluded
related_requirements: []
confidence: confirmed
---

# CAND-180: Development Node Version Consistency

### Source

> - [x] Use the same Node 26.5.0 for frontend in dockerfile and in dev

[Source lines 551-551](../../../../backlog/backlog.md#L551)

### Candidate behavior

No new requirement behavior asserted; the source selects a development-environment Node version.

### Applicability

- Actor: development team
- Location: frontend Dockerfile and local development
- Event kind: not applicable
- Interaction mode: not applicable
- Viewport: not applicable
- State: development configuration
- Exclusions: example environment configuration and deployed product behavior

### Classification

ADR or decision

### Existing Requirements and Confidence

Existing requirements: None identified.
Overlap: CAND-181 records the separate example-environment configuration check; no requirement overlap identified.

Confidence: confirmed

### Disposition

Retain as configuration provenance; no FR or QR.
