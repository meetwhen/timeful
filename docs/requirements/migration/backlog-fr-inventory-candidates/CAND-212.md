---
id: CAND-212
title: Secure Production And Staging MongoDB Access
verdict: needs-decision
related_requirements: []
confidence: needs-product-decision
---

# CAND-212: Secure Production And Staging MongoDB Access

## Source

> Production and staging MongoDB shall disallow unauthenticated access.
>
> - Root user + limited app user
> - Authentication required
> - Credentials stored in corresponding .env files

## Candidate behavior

The source asserts that production and staging MongoDB deployments require authenticated access, but does not establish independently testable configuration or credential-handling boundaries.

## Applicability

Actor: operator.
Location: production and staging MongoDB deployments.
Event kind: not applicable.
Interaction mode: deployment configuration.
Viewport: not applicable.
State: configuring database access.
Exclusions: development and test environments.

## Classification

needs product decision

## Existing Requirements and Confidence

None identified.
Confidence: needs product decision.

## Disposition

Retain for later quality-requirement triage.

## Open Questions

- Which deployments, authentication mechanisms, application roles, and credential-storage boundaries are required?
- Does the existing deployment-configuration QR cover any part of this outcome?
