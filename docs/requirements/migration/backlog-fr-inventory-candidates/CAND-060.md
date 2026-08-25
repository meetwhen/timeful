---
id: CAND-060
verdict: needs-decision
related_requirements:
  - FR-011
confidence: needs-product-decision
---

# CAND-060

## Source

> - [x] there should be a label for each grid row that marks the start of an hour, even for collapsed

## Candidate behavior

No new requirement behavior asserted; accepted requirements already specify collapsed-band boundaries, not every row label.

## Applicability

Actor: event visitor. Location: timed-grid time axis. Event kind: timed. Interaction mode: viewing or scheduling. Viewport: any. State: normal or collapsed rows. Exclusions: availability editing.

## Classification

needs product decision

## Existing Requirements and Confidence

FR-011 requires each collapsed band's start boundary. Confidence: needs product decision.

## Disposition

Do not migrate the every-row extension pending confirmation.

## Open Questions

Must every visible hour row be labeled, or only collapsed boundaries?
