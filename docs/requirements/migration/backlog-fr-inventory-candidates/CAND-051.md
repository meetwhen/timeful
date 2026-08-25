# CAND-051

## Source

> - [x] in What days might work, when sun and mon are selected, when enabling start on monday, both mon and sun must be selected

## Candidate behavior

No new requirement behavior asserted; this is an ambiguous date-selection rule.

## Applicability

Actor: event editor. Location: What days might work. Event kind: unconfirmed. Interaction mode: week-start change. Viewport: any. State: Sunday and Monday selected. Exclusions: other selected dates.

## Classification

needs product decision

## Existing Requirements and Confidence

FR-025 requires at least one picked date; it does not define week-start effects. Confidence: confirmed.

## Disposition

Do not migrate pending domain-rule confirmation.

## Open Questions

Why should changing week start select dates, and does this apply to all date-picker modes?
