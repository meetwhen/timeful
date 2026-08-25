# CAND-104

## Source

> - [x] Don't show the cursor and tooltip in enabled, inactive or disabled cells

## Candidate behavior

The grid shows neither cursor nor tooltip for enabled-inactive or disabled cells.

## Applicability

Actor: event visitor; Location: timed grid; Event kind: timed; Interaction mode: hover or selection; Viewport: unspecified; State: enabled-inactive or disabled cell; Exclusions: active cells.

## Classification

candidate FR

## Existing Requirements and Confidence

Overlap: accepted FR-014 derives grid cell states from event domains but does not prescribe cursor or tooltip visibility. Confidence: inferred.

## Disposition

Review with CAND-112 as inactive-cell interaction behavior.

## Open Questions

Does `cursor` mean a grid selection indicator, pointer cursor, or scheduled-event cursor?
