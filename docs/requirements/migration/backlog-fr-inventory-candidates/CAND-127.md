# CAND-127

## Source

<!-- prettier-ignore -->
> - [x] Keep yellow status for "if needed" responses but don't:
>   - highlight such responses text
>   - show asterisk ("\*") near such responses
>   - show "* if needed" under responses

[Source lines 463-466](../../../../backlog/backlog.md#L463-L466)

## Candidate behavior

If needed responses retain yellow status but omit text highlight, an asterisk, and the `* if needed` note.

## Applicability

Actor: event visitor; Location: Responses; Event kind: unspecified; Interaction mode: viewing; Viewport: unspecified; State: If needed response; Exclusions: other response statuses.

## Classification

candidate FR

## Existing Requirements and Confidence

Overlap: accepted FR-006 defines availability-state exclusivity, not its presentation. Confidence: inferred.

## Disposition

Review as one response-status presentation requirement.

## Open Questions

Is yellow defined by a shared status-color token and is it accessible in every response view?
