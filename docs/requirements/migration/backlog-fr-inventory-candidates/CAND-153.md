# CAND-153

### Source

<!-- prettier-ignore -->
> - [x] <http://127.0.0.1:4173/e/aB3BE>
>   - ok in gmt+3
>   - padding gets added in +3:30 although not needed
>   - grid split in +12 although shouldn't
>
>   Status: can't reproduce

[Source lines 516-521](../../../../backlog/backlog.md#L516-L521)

### Candidate behavior

No new requirement behavior asserted; this is an unreproduced timezone-specific investigation report.

### Applicability

Actor: user. Location: a specific event page. Event kind: unspecified. Interaction mode: viewing after timezone selection. Viewport: unspecified. State: GMT+3, +3:30, or +12. Exclusions: reproducible behavior not established.

### Classification

bug or investigation

### Existing Requirements and Confidence

FR-013 accepts preservation of timed-slot instants across display-timezone changes, but does not confirm the reported padding or split. Confidence: inferred.

### Disposition

Exclude from requirements unless reproduced and recast as a durable outcome.

### Open Questions

What event configuration, display timezone, and viewport reproduce the report?
