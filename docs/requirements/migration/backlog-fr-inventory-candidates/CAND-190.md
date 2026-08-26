---
id: CAND-190
verdict: excluded
related_requirements:
  - FR-050
  - FR-051
  - FR-052
confidence: confirmed
---

# CAND-190: Specific-Times Date Regression Report

### Source

> - [x] Create an event with specific times for dates Aug 30, 31, mark hours 0-4 for both dates, edit event, set dates for 28, 29, click next.
>       See May 28, 30, 31 in specific times page, and May 30, 31 on the event page.
>   - Can't reproduce.
>     I see Aug 28, Aug 29 on the event page

[Source lines 566-567](../../../../backlog/backlog.md#L566-L567)

### Candidate behavior

No new requirement behavior asserted; this is a non-reproduced date-editing regression report.

### Applicability

- Actor: event owner
- Location: specific-times editing and event page
- Event kind: timed
- Interaction mode: event settings editing
- Viewport: unspecified
- State: changed date selection
- Exclusions: confirmed date-domain behavior

### Classification

bug or investigation

### Existing Requirements and Confidence

Existing requirements: [FR-050](../../functional/fr/FR-050.md), [FR-051](../../functional/fr/FR-051.md), and [FR-052](../../functional/fr/FR-052.md) are proposed date-domain rules; none confirms this report.

Confidence: confirmed

### Disposition

Retain as investigation history; no FR or QR.

### Open Questions

- Which timezone, date picker state, and saved data are required to reproduce the reported May dates?
