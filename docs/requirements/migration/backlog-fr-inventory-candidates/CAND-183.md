---
id: CAND-183
verdict: needs-decision
requirement_type: null
related_requirements:
  - FR-065
confidence: needs-product-decision
---

# CAND-183: Edit-Availability Visibility

### Source

> - [x] On the event page:
>   - when there are no responses, the Edit availability shall be not visible.
>   - When there are responses, shall be visible

[Source lines 556-558](../../../../backlog/backlog.md#L556-L558)

### Candidate behavior

No new requirement behavior asserted; the source does not resolve whether
availability-action visibility depends on all event responses or responses the
current visitor may edit as [Availability Responses](../../../terminology/glossary.md#availability-response).

### Applicability

- Actor: event visitor
- Location: event page
- Event kind: unspecified
- Interaction mode: viewing
- Viewport: unspecified
- State: zero responses or one or more responses
- Exclusions: add-availability visibility and authorization

### Classification

needs product decision

### Existing Requirements and Confidence

Existing requirements: [FR-065](../../functional/fr/FR-065.md) is proposed and concerns availability-editing response lists, not action visibility.

Confidence: needs product decision

### Disposition

Needs product decision; do not consolidate until the controlling response count and no-editable-response behavior are chosen.

### Open Questions

- Does the action depend on all event responses or only responses editable by the current visitor?
- When no response is editable, should the action be hidden or disabled?
