# CAND-199: Mobile Availability Panel Details

### Source

> - [x] Bug:
>   - on mobile
>   - I add and save availability
>   - I click an active timeslot
>   - Responses panel with my response slides from the bottom
>   - I click Edit availability
>   - I click a grid timeslot
>   - Expected: on the panel that slides from the bottom, I see only Available/if needed
>   - Actual: on the panel, I see both Available/If needed and Responses

[Source lines 577-585](../../../../backlog/backlog.md#L577-L585)

### Candidate behavior

No new requirement behavior asserted; the source is a concrete defect report already covered by the mobile editing rule.

### Applicability

- Actor: event visitor
- Location: mobile event page selected-slot panel
- Event kind: timed
- Interaction mode: availability editing
- Viewport: mobile
- State: saved availability, selected active slot, then Edit availability
- Exclusions: read-mode response details

### Classification

existing requirement

### Existing Requirements and Confidence

Existing requirements: [FR-028](../../functional/fr/FR-028.md) requires availability controls and excludes response details from the mobile selected-slot panel while editing availability.

Confidence: confirmed

### Disposition

Map to FR-028 as a regression example; no new FR.
