# CAND-211: Clarify End-Of-Day Time-Range Labels

#### Source

> Time-range picker labels shall represent the end-of-day boundary clearly:
>
> - In 24-hour mode, labels shall be zero-padded from `00:00` through `23:00`, followed by `24:00`.
> - In 12-hour mode, the end-of-day option shall be `12 AM`.
> - Selecting the end-of-day boundary shall render as `00:00` in the next date column of the grid.

[Source FR-022](../../functional-requirements.md#fr-022)

#### Candidate behavior

The time-range picker represents the end-of-day boundary as `24:00` in 24-hour mode and `12 AM` in 12-hour mode; selecting it renders as `00:00` in the next date column of the timed grid.

#### Applicability

Actor: event owner. Location: timed event form time-range picker and timed grid. Event kind: timed. Interaction mode: editing a time range. Viewport: unspecified. State: selecting the end-of-day boundary. Exclusions: non-boundary time labels and dates-only events.

#### Classification

candidate FR

#### Existing Requirements and Confidence

Accepted FR-024 keeps event and display time formats independent but does not specify end-of-day picker labels or grid rendering. Confidence: inferred.

#### Disposition

Hold as a time-range boundary behavior candidate.
