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

candidate FR

### Existing Requirements and Confidence

Existing requirements: [FR-023](../../functional/fr/FR-023.md) is accepted and covers registration status during sign-in, but not an OTP send failure.

Confidence: inferred

### Disposition

Consolidate with registration feedback behavior if OTP is confirmed as the current registration mechanism.

### Open Questions

- What failures receive a report, and what information may it disclose?
