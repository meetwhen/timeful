---
id: CAND-026
verdict: covered
related_requirements: [FR-011]
confidence: confirmed
---

# CAND-026

## Source

> - [x] let's collapse hours when they're at the start or at the end too. These hours are useless anyway

## Candidate behavior

No new requirement behavior asserted; accepted collapse behavior already includes leading and trailing runs.

## Applicability

Actor: event visitor. Location: timed grid. Event kind: timed. Interaction mode: viewing or scheduling. Viewport: any. State: leading or trailing inactive hours. Exclusions: availability editing and specific-times setting.

## Classification

existing requirement

## Existing Requirements and Confidence

FR-011 explicitly covers leading and trailing inactive runs. Confidence: confirmed.

## Disposition

Map to FR-011; no migration.

## Open Questions

None.
