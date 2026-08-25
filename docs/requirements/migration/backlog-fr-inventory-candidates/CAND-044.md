# CAND-044

#### Source

> - [x] Create event with 9 - 17, timezone +9 (<http://127.0.0.1:4173/e/ee4Cb>), june 11 and 12.
>       Expected: june 10 and 11 are shown on the event page
>       Actual: june 11, 12, 13 are shown there

[Source lines 327-329](../../../../backlog/backlog.md#L327-L329)

#### Candidate behavior

The event page should render the dates to which timed slots project in the selected display timezone.

#### Applicability

Actor: event visitor. Location: event-page timed grid. Event kind: timed. Interaction mode: viewing. Viewport: any. State: timezone projection crosses dates. Exclusions: changing event timezone.

#### Classification

needs product decision

#### Existing Requirements and Confidence

FR-002 and FR-013 require projected-date columns and preserved instants. Confidence: confirmed.

#### Disposition

Treat as a regression scenario for FR-002 and FR-013.

#### Open Questions

None.
