---
id: CAND-214
title: Keep The Server As An Application Rather Than A Library
verdict: needs-decision
related_requirements: []
confidence: needs-product-decision
---

# CAND-214: Keep The Server As An Application Rather Than A Library

## Source

> The server shall not used as a libarary.
>
> Thus:
>
> - Can use an internal name prefixed with `timeful` instead of prefixed with `github.com/deemp/timeful`

## Candidate behavior

The source asserts that the server is not intended for use as a library, but does not establish a verifiable compatibility or module-boundary outcome.

## Applicability

Actor: server maintainer.
Location: server module and internal package boundaries.
Event kind: not applicable.
Interaction mode: importing or building server code.
Viewport: not applicable.
State: consuming server packages.
Exclusions: application runtime behavior.

## Classification

needs product decision

## Existing Requirements and Confidence

None identified.
Confidence: needs product decision.

## Disposition

Retain for later quality-requirement triage.

## Open Questions

- Is this a product quality outcome, a module-structure decision, or an implementation convention?
- Which external imports, if any, must the server support?
