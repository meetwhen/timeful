# CAND-183: Edit-Availability Visibility

### Source

> - [x] On the event page:
>   - when there are no responses, the Edit availability shall be not visible.
>   - When there are responses, shall be visible

[Source lines 556-558](../../../../backlog/backlog.md#L556-L558)

### Candidate behavior

The event page hides `Edit availability` when there are no responses and shows it when responses exist.

### Applicability

- Actor: event visitor
- Location: event page
- Event kind: unspecified
- Interaction mode: viewing
- Viewport: unspecified
- State: zero responses or one or more responses
- Exclusions: add-availability visibility and authorization

### Classification

candidate FR

### Existing Requirements and Confidence

Existing requirements: [FR-065](../../functional/fr/FR-065.md) is proposed and concerns availability-editing response lists, not action visibility.

Confidence: inferred

### Disposition

Consolidate only after confirming whether “responses” means all event responses or editable responses.

### Open Questions

- Which response count controls the action for a visitor without edit authorization?
