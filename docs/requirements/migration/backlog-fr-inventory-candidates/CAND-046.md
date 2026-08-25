# CAND-046

#### Source

> - [x] Shown in shouldn't affect the time zone

[Source lines 331-331](../../../../backlog/backlog.md#L331-L331)

#### Candidate behavior

Changing the display timezone should not change the event timezone.

#### Applicability

Actor: event visitor. Location: event-page display controls. Event kind: timed. Interaction mode: display-timezone change. Viewport: any. State: timezone selected. Exclusions: explicit event-settings timezone edits.

#### Classification

needs product decision

#### Existing Requirements and Confidence

FR-013 explicitly preserves the event timezone. Confidence: confirmed.

#### Disposition

Map to FR-013; no migration.

#### Open Questions

None.
