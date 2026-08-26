---
id: CAND-213
title: Permit Development MongoDB Access Without Authentication
verdict: needs-decision
related_requirements: []
confidence: needs-product-decision
---

# CAND-213: Permit Development MongoDB Access Without Authentication

## Source

> Dev MongoDB shall allow unauthenticated access.

## Candidate behavior

The source asserts that a development MongoDB deployment permits unauthenticated access, but does not establish the environment boundary or applicable safeguards.

## Applicability

Actor: developer.
Location: development MongoDB deployment.
Event kind: not applicable.
Interaction mode: local development configuration.
Viewport: not applicable.
State: connecting to MongoDB during development.
Exclusions: staging, production, and test environments.

## Classification

needs product decision

## Existing Requirements and Confidence

None identified.
Confidence: needs product decision.

## Disposition

Retain for later quality-requirement triage.

## Open Questions

- Which environments qualify as development, and is unauthenticated access required or merely permitted?
- Which network, data, and credential safeguards apply to a development deployment?
