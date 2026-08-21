# Functional Requirement Authoring

Read `../README.md` for the shared requirement format, metadata, component, and
index conventions before creating or changing a functional requirement.

## Independently Verifiable Behavior

A functional requirement states one durable system behavior that a reader can
understand and verify without consulting the originating backlog task,
implementation, or bug report.

## Applicability And Boundaries

- Name the actor when behavior differs by viewer, respondent, creator, editor,
  or another role.
- Name the location when behavior applies only to a page, view, control, or
  frontend entry point.
- Name the event kind, viewport, interaction mode, permission state, response
  state, or other condition when it limits applicability.
- State exclusions when a similar mode or context could otherwise appear to be
  covered.

For example, `The system shall show a tooltip for locked responses` omits who
encounters it and where. `In the event-page Responses view, selecting the
locked icon for a response the current viewer cannot edit shall display a
tooltip that explains why the response is not editable` identifies the actor,
location, trigger, and outcome.

## Review

- Can a reader identify who experiences the behavior, where it applies, and
  under which relevant conditions without opening the source task?
- Are similar but excluded contexts explicitly ruled out when ambiguity is
  likely?
