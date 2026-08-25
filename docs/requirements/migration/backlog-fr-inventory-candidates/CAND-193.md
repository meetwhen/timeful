---
id: CAND-193
verdict: needs-decision
related_requirements:
  - FR-031
confidence: needs-product-decision
---

# CAND-193: OTP Send-Failure Report

### Source

> - [x] On the Create your account form, when I click Continue and the OTP can't be sent, a report about that is visible

[Source lines 571-571](../../../../backlog/backlog.md#L571)

### Candidate behavior

When sending an OTP fails after `Continue` on the Create your account form, a failure report is visible.

### Applicability

- Actor: account registrant
- Location: Create your account form
- Event kind: not applicable
- Interaction mode: registration
- Viewport: unspecified
- State: OTP send failure
- Exclusions: successful sends and failure-report wording

### Classification

needs product decision

### Existing Requirements and Confidence

Existing requirements: [FR-031](../../functional/fr/FR-031.md) specifies registration magic links, not OTP; it does not specify send-failure feedback.

Confidence: needs product decision

### Disposition

Do not migrate pending a decision on the registration mechanism and send-failure feedback scope.

### Open Questions

- What failures receive a report, and what information may it disclose?
