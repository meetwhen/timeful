---
id: CAND-151
verdict: needs-decision
related_requirements: [QR-012]
confidence: needs-product-decision
---

# CAND-151

### Source

> - [x] Improve separation between environments (development, test, staging, production)

[Source lines 514-514](../../../../backlog/backlog.md#L514-L514)

### Candidate behavior

No new requirement behavior asserted; “improve separation” does not specify an observable environment boundary.

### Applicability

Actor: operator or developer.
Location: deployment environments.
Event kind: any.
Interaction mode: development and deployment.
Viewport: not applicable.
State: development, test, staging, or production.
Exclusions: unspecified.

### Classification

needs product decision

### Existing Requirements and Confidence

QR-012 distinguishes staging and production from intentionally isolated development and test configuration, but does not define environment separation outcomes.
Confidence: needs product decision.

### Disposition

Do not consolidate without defined isolation outcomes.

### Open Questions

Which resources must be separated: data, credentials, hosts, services, or configuration?
