---
id: CAND-138
verdict: excluded
requirement_type: null
related_requirements: [FR-007]
confidence: confirmed
---

# CAND-138

## Source

> - [x] On each page, Sign in must be the left-most button

[Source lines 478-478](../../../../backlog/backlog.md#L478-L478)

## Candidate behavior

When present, Sign in is the left-most button on each page.

## Applicability

Actor: visitor; Location: every page header or action area; Event kind: not applicable; Interaction mode: viewing; Viewport: unspecified; State: Sign in available; Exclusions: pages without Sign in not established.

## Classification

implementation detail

## Existing Requirements and Confidence

Overlap: proposed FR-007 governs disabled-state visibility, not ordering.
Confidence: confirmed.

## Disposition

Exclude from requirement migration; global button ordering is incidental navigation layout.

## Open Questions

Does left-most mean visual order in RTL layouts and on mobile menus?
