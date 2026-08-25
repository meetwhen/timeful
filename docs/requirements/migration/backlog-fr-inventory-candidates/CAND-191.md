---
id: CAND-191
verdict: excluded
related_requirements:
  - FR-074
confidence: confirmed
---

# CAND-191: Specific-Times Missing-Date Regression Report

### Source

> - [x] Create an event with two dates, mark timeslots for only one day in specific times, save, edit again and see only one day on the event page
>   - Can't reproduce. I see both days.

[Source lines 568-569](../../../../backlog/backlog.md#L568-L569)

### Candidate behavior

No new requirement behavior asserted; this is a non-reproduced saved-date regression report.

### Applicability

- Actor: event owner
- Location: specific-times editing and event page
- Event kind: timed
- Interaction mode: event settings editing
- Viewport: unspecified
- State: two selected dates with slots marked on one date
- Exclusions: confirmed date persistence behavior

### Classification

bug or investigation

### Existing Requirements and Confidence

Existing requirements: [FR-074](../../functional/fr/FR-074.md) is proposed and derives the enabled domain for each picked date; it does not confirm the report.

Confidence: confirmed

### Disposition

Retain as investigation history; no FR or QR.

### Open Questions

- What persisted data or transition would cause one selected date to disappear?
