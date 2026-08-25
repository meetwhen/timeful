---
id: CAND-182
verdict: needs-decision
related_requirements:
  - QR-012
confidence: needs-product-decision
---

# CAND-182: Staging Environment

### Source

> - [x] introduce staging environment

[Source lines 555-555](../../../../backlog/backlog.md#L555)

### Candidate behavior

No new requirement behavior asserted; the source names an environment but does not specify its purpose, parity, access, or lifecycle.

### Applicability

- Actor: operations team
- Location: deployment environments
- Event kind: not applicable
- Interaction mode: not applicable
- Viewport: not applicable
- State: deployment and release validation
- Exclusions: unspecified production behavior

### Classification

needs product decision

### Existing Requirements and Confidence

Existing requirements: [QR-012](../../quality/qr/QR-012.md) applies safe configuration validation to staging deployments, but does not establish staging's purpose, parity, access, or lifecycle.

Confidence: needs product decision

### Disposition

Require an environment decision before considering a QR or ADR.

### Open Questions

- What release-validation purpose and production parity must staging provide?
- Who can deploy to and access staging?
