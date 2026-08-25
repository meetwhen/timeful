---
id: CAND-124
verdict: excluded
requirement_type: null
related_requirements: [FR-040]
confidence: confirmed
---

# CAND-124

## Source

> - [x] On the Event not found page, the "Back to home" button must have a black shadow (like on the home page), not greenish glow

[Source lines 460-460](../../../../backlog/backlog.md#L460-L460)

## Candidate behavior

The Event not found Back to home button uses the home page's black shadow rather than a greenish glow.

## Applicability

Actor: event visitor; Location: Event not found page; Event kind: not applicable; Interaction mode: viewing; Viewport: unspecified; State: button visible; Exclusions: other buttons.

## Classification

implementation detail

## Existing Requirements and Confidence

Overlap: proposed FR-040 centers the not-found page, not button styling. Confidence: confirmed.

## Disposition

Exclude from requirement migration; the page-specific shadow is incidental presentation.

## Open Questions

Is matching the home-page token sufficient, or is black shadow an exact visual constraint?
