# CAND-055

#### Source

<!-- prettier-ignore -->
> - [x] use the same palette for event page
>   (enabled - light-grey, active -red; disabled - dark-grey)

[Source lines 340-341](../../../../backlog/backlog.md#L340-L341)

#### Candidate behavior

No new requirement behavior asserted; the child gives styling tokens for existing state distinctions.

#### Applicability

Actor: event visitor. Location: event-page timed grid. Event kind: timed. Interaction mode: viewing. Viewport: any. State: enabled, active, or disabled cells. Exclusions: custom-domain editing colors where separately specified.

#### Classification

needs product decision

#### Existing Requirements and Confidence

FR-009 and FR-014 specify distinct state presentation, including light-grey and dark-grey treatments. Confidence: confirmed.

#### Disposition

Map behavioral state meaning to FR-009 and FR-014; retain exact palette as styling.

#### Open Questions

Does “active -red” mean unavailable active slots only?
