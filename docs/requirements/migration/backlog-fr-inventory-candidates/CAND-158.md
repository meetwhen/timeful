# CAND-158

### Source

> - [x] On the specific times page <http://127.0.0.1:4173/e/Eb67A>, when I switch timezone from +5 to +4 and on june 14, 0-4 are selected, jun 13 should appear

[Source lines 526-526](../../../../backlog/backlog.md#L526-L526)

### Candidate behavior

After changing the timezone from +5 to +4 with June 14 0-4 selected, June 13 appears on the stated specific-times page.

### Applicability

Actor: user. Location: specific-times page. Event kind: timed. Interaction mode: changing timezone. Viewport: unspecified. State: +5 to +4 with the stated selection. Exclusions: other transitions and selections.

### Classification

duplicate or refinement

### Existing Requirements and Confidence

FR-013 accepts preservation of timed-slot instants across display-timezone changes; this is a specific expected projection. Confidence: inferred.

### Disposition

Treat as a regression example for FR-013 rather than a new requirement.

### Open Questions

Is “June 13 should appear” a date column, a grid row, or an enabled slot?
