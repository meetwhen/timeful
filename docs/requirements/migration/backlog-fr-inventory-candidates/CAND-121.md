---
id: CAND-121
verdict: proposed-requirement
requirement_type: FR
related_requirements: [FR-049]
confidence: inferred
---

# CAND-121

## Source

> - [x] On desktop, on a dates only event page, in the second row, on the right:
>   - When there are no responses, show "Start on Monday" option in the second row
>   - When there are responses, show "Show best days" option and "More options" (which includes "Start on Monday" and "Hide if needed")

[Source lines 455-457](../../../../backlog/backlog.md#L455-L457)

## Candidate behavior

On desktop dates-only pages, the second row shows Start on Monday with no responses, or Show best days and More options with responses.

## Applicability

Actor: event visitor; Location: desktop dates-only event page second row; Event kind: dates-only; Interaction mode: viewing; Viewport: desktop; State: no responses or one or more responses; Exclusions: timed events and mobile.

## Classification

candidate FR

## Existing Requirements and Confidence

Overlap: proposed FR-049 organizes mobile controls, not this desktop dates-only conditional display.
Confidence: inferred.

## Disposition

Review as conditional controls behavior.

## Open Questions

Should More options always contain both listed options, and what does Hide if needed affect?
