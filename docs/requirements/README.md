# Requirements

This directory is the canonical record of Timeful product requirements.
Requirements state durable, verifiable behavior or quality constraints.
They
are not implementation tasks, bug reports, design investigations, or decision
records.

<h2>Table of contents</h2>

- [Layout](#layout)
- [Requirement Format](#requirement-format)
- [Related Artifacts](#related-artifacts)
- [Terminology](#terminology)
- [Functional Requirements](#functional-requirements)
- [Quality Requirements](#quality-requirements)

## Layout

- `functional/` contains functional requirements.
- `quality/README.md` contains quality-requirement authoring guidance.
- `quality/qr/` contains quality requirements.
- `quality-requirements.md` is a temporary migration source.
  Do not add new
  requirements to it.

Each requirement has its own file.
Use an ID-only filename so links remain
stable when its title changes:

```text
functional/fr/FR-001.md
quality/qr/QR-001.md
```

The requirement ID and filename are permanent.
Titles may change.

## Requirement Format

Every requirement file starts with machine-readable YAML front matter followed
by a human-readable title and requirement statement:

```md
---
id: FR-001
title: Preserve guest availability records
type: functional
components:
  - frontend
  - backend
status: accepted
---

# FR-001: Preserve Guest Availability Records

A guest can add availability multiple times.

The system shall retain every response record in that guest's browser.
```

Use one of these `components` values:

- `frontend` for browser application behavior.
- `backend` for API and server behavior.
- `infrastructure` for deployment and runtime environment behavior.

List the components responsible for enforcing the requirement.
A requirement
that spans components remains one file and lists each applicable component;
do not duplicate it.

Quality requirements also declare exactly one ISO/IEC 25010 quality
characteristic and subcharacteristic:

```yaml
characteristic: security
subcharacteristic: confidentiality
```

See [`quality/README.md`](quality/README.md) for the allowed values and
quality-scenario authoring rules.

Every requirement has a concise title and is atomic: it is independently
understandable, decidable, and verifiable.
A requirement can contain several
acceptance criteria when they jointly specify one behavior, but unrelated
obligations belong in separate requirement files.

Describe observable outcomes rather than implementation mechanisms or
incidental UI structure.

The `status` can be one of these:

- proposed
- rejected
- accepted
- deprecated
- superseded by FR-005

## Related Artifacts

- Functional requirements (`FR-*`) state required system behavior.
- Quality requirements (`QR-*`) state required quality attributes or runtime
  constraints.
- Specifications (`SPEC-*`) describe a feature or domain in broader context
  and may link to multiple requirements.
- Architecture decision records (ADRs) record architectural decisions.
- FRs and QRs must remain self-contained and must not cite ADRs as normative
  dependencies.
  ADRs may link to the FRs and QRs they enable, constrain, or
  satisfy.
  Requirement provenance is recorded separately from the requirement
  statement.
- Backlog tasks track work to investigate, implement, or verify requirements.

## Terminology

Use the controlled terms and linking conventions in the
[terminology guide](../terminology/README.md).
When a requirement uses a
controlled term, link its first occurrence in each paragraph, list item, and
table cell to the matching [glossary](../terminology/glossary.md) entry.

## Functional Requirements

For guidance on authoring functional requirements, read
[`functional/README.md`](functional/README.md).

| ID                                | Title                                                                                    | Components        |
| --------------------------------- | ---------------------------------------------------------------------------------------- | ----------------- |
| [FR-001](functional/fr/FR-001.md) | Preserve event guest availability responses                                              | frontend          |
| [FR-002](functional/fr/FR-002.md) | Project timed slots into their projected date columns                                    | frontend          |
| [FR-003](functional/fr/FR-003.md) | Delete availability responses                                                            | frontend, backend |
| [FR-004](functional/fr/FR-004.md) | Require a non-empty availability response                                                | frontend          |
| [FR-005](functional/fr/FR-005.md) | Render the edited availability response as an overlay above other availability responses | frontend          |
| [FR-006](functional/fr/FR-006.md) | Keep availability states mutually exclusive                                              | frontend, backend |
| [FR-007](functional/fr/FR-007.md) | Hide sign-in entry points when sign-in is disabled                                       | frontend          |
| [FR-008](functional/fr/FR-008.md) | Initialize the ranged domain mode active slot range                                      | frontend          |
| [FR-009](functional/fr/FR-009.md) | Present timed grid states consistently                                                   | frontend          |
| [FR-010](functional/fr/FR-010.md) | Support custom domain editing                                                            | frontend          |
| [FR-011](functional/fr/FR-011.md) | Collapse inactive timed grid runs                                                        | frontend          |
| [FR-012](functional/fr/FR-012.md) | Maintain an optional scheduled event time                                                | frontend, backend |
| [FR-013](functional/fr/FR-013.md) | Preserve timed slot instants across display-timezone changes                             | frontend          |
| [FR-014](functional/fr/FR-014.md) | Derive timed grid cell states from event domains                                         | frontend          |
| [FR-015](functional/fr/FR-015.md) | Save active slots only within the enabled domain                                         | frontend, backend |
| [FR-016](functional/fr/FR-016.md) | Limit scheduling to active slots                                                         | frontend          |
| [FR-017](functional/fr/FR-017.md) | Preserve timed slots across daylight saving transitions                                  | frontend, backend |
| [FR-018](functional/fr/FR-018.md) | Restrict event settings editing to the event owner                                       | frontend, backend |
| [FR-019](functional/fr/FR-019.md) | Preserve multiline event descriptions                                                    | frontend, backend |
| [FR-020](functional/fr/FR-020.md) | Keep the mobile selected Timed Slot tooltip available after scrolling                    | frontend          |
| [FR-021](functional/fr/FR-021.md) | Initialize a new event timezone from the browser                                         | frontend          |
| [FR-022](functional/fr/FR-022.md) | Display adjacent-month days in the new-event date picker                                 | frontend          |
| [FR-023](functional/fr/FR-023.md) | Display email registration status during sign-in                                         | frontend, backend |
| [FR-024](functional/fr/FR-024.md) | Keep event and display time formats independent                                          | frontend          |
| [FR-025](functional/fr/FR-025.md) | Require event date selection before creating an event                                    | frontend          |
| [FR-026](functional/fr/FR-026.md) | Define event kind values                                                                 | frontend          |
| [FR-027](functional/fr/FR-027.md) | Restrict dates-only event timezone settings to the event owner                           | frontend          |
| [FR-028](functional/fr/FR-028.md) | Hide mobile response details while editing availability                                  | frontend          |
| [FR-029](functional/fr/FR-029.md) | Initialize the display timezone from the browser                                         | frontend          |
| [FR-030](functional/fr/FR-030.md) | Persist only active timed slots in browser storage                                       | frontend          |
| [FR-031](functional/fr/FR-031.md) | Send registration magic links by email                                                   | frontend, backend |
| [FR-032](functional/fr/FR-032.md) | Complete registration after a Magic Link sign-in                                         | frontend, backend |
| [FR-033](functional/fr/FR-033.md) | Configure the feedback link                                                              | frontend          |
| [FR-034](functional/fr/FR-034.md) | Gate the Discord banner with an environment flag                                         | frontend          |
| [FR-035](functional/fr/FR-035.md) | Allow the event owner to enter an event settings description during creation             | frontend          |
| [FR-036](functional/fr/FR-036.md) | Reset the Display Timezone to the Event Timezone                                         | frontend          |
| [FR-037](functional/fr/FR-037.md) | Do not reset the Event Timezone in the event form                                        | frontend          |
| [FR-038](functional/fr/FR-038.md) | Keep the mobile availability-name input above the keyboard                               | frontend          |
| [FR-039](functional/fr/FR-039.md) | Show event settings to non-owners                                                        | frontend          |
| [FR-040](functional/fr/FR-040.md) | Center the not-found page                                                                | frontend          |
| [FR-041](functional/fr/FR-041.md) | Cue availability editing from Unavailable timed grid areas                               | frontend          |
| [FR-042](functional/fr/FR-042.md) | Keep timed grid highlights within their cells                                            | frontend          |
| [FR-043](functional/fr/FR-043.md) | Place Timed Slot tooltips at the selected slot                                           | frontend          |
| [FR-044](functional/fr/FR-044.md) | Explain Protected Responses                                                              | frontend          |
| [FR-045](functional/fr/FR-045.md) | Show Event Guests as Unavailable for Unavailable timed grid states                       | frontend          |
| [FR-046](functional/fr/FR-046.md) | Label the Display Time Format control                                                    | frontend          |
| [FR-047](functional/fr/FR-047.md) | Label the Display Timezone control                                                       | frontend          |
| [FR-048](functional/fr/FR-048.md) | Center the no-response mobile Show all hours control                                     | frontend          |
| [FR-049](functional/fr/FR-049.md) | Arrange mobile event-page controls in two rows                                           | frontend          |
| [FR-050](functional/fr/FR-050.md) | Generate active slots for newly picked dates                                             | frontend, backend |
| [FR-051](functional/fr/FR-051.md) | Keep newly picked Custom Domain Mode dates inactive                                      | frontend, backend |
| [FR-052](functional/fr/FR-052.md) | Remove active slots for removed picked dates                                             | frontend, backend |
| [FR-053](functional/fr/FR-053.md) | Preserve active slots when converting to Custom Domain Mode                              | frontend, backend |
| [FR-054](functional/fr/FR-054.md) | Regenerate active slots when converting to Ranged Domain Mode                            | frontend, backend |
| [FR-055](functional/fr/FR-055.md) | Persist changed timed-event timezones                                                    | frontend, backend |
| [FR-056](functional/fr/FR-056.md) | Preserve Active Slot Settings on no-op saves                                             | frontend, backend |
| [FR-057](functional/fr/FR-057.md) | Apply timed-event timezone transition rules                                              | frontend, backend |
| [FR-058](functional/fr/FR-058.md) | Reject availability outside the event domain                                             | backend           |
| [FR-059](functional/fr/FR-059.md) | Validate and normalize Availability Response names                                       | frontend, backend |
| [FR-060](functional/fr/FR-060.md) | Create Protected Responses by default                                                    | frontend, backend |
| [FR-061](functional/fr/FR-061.md) | Allow Event Guests to open their Availability Responses                                  | frontend, backend |
| [FR-062](functional/fr/FR-062.md) | Restore Event Guest access after event sign-in                                           | frontend, backend |
| [FR-063](functional/fr/FR-063.md) | Associate the Event Owner after event sign-in                                            | frontend, backend |
| [FR-064](functional/fr/FR-064.md) | Keep protected availability responses in schedule overlap                                | frontend          |
| [FR-065](functional/fr/FR-065.md) | Show availability responses for availability editing                                     | frontend          |
| [FR-066](functional/fr/FR-066.md) | Filter the Availability Editor grid by availability response selection                   | frontend          |
| [FR-067](functional/fr/FR-067.md) | Create custom-domain timed events                                                        | frontend, backend |
| [FR-068](functional/fr/FR-068.md) | Regenerate active slots when the active slot range changes                               | frontend, backend |
| [FR-069](functional/fr/FR-069.md) | Calculate schedule overlap from availability states                                      | frontend          |
| [FR-070](functional/fr/FR-070.md) | Omit advertising when freemium is disabled                                               | frontend          |
| [FR-071](functional/fr/FR-071.md) | Remove advertising layout reservations when freemium is disabled                         | frontend          |
| [FR-072](functional/fr/FR-072.md) | Bypass freemium restrictions when freemium is disabled                                   | frontend          |
| [FR-073](functional/fr/FR-073.md) | Retain browser-local Event Visitor Identities                                            | frontend          |
| [FR-074](functional/fr/FR-074.md) | Derive the timed-event enabled domain                                                    | frontend, backend |
| [FR-075](functional/fr/FR-075.md) | Generate initial ranged timed-event active slots                                         | frontend, backend |
| [FR-076](functional/fr/FR-076.md) | Separate the mobile event action bar from the Timed Grid                                 | frontend          |
| [FR-077](functional/fr/FR-077.md) | Display optional event descriptions as read-only content                                 | frontend          |
| [FR-078](functional/fr/FR-078.md) | Provide mobile header navigation actions                                                 | frontend          |
| [FR-079](functional/fr/FR-079.md) | Associate Event Visitor Identities with Platform Identities                              | frontend, backend |
| [FR-080](functional/fr/FR-080.md) | Default signed-in response names from the account profile                                | frontend, backend |
| [FR-081](functional/fr/FR-081.md) | Transfer anonymous PostgreSQL guest authority to another browser                         | frontend, backend |
| [FR-082](functional/fr/FR-082.md) | Transfer platform sign-in from a PostgreSQL event to another browser                     | frontend, backend |
| [FR-083](functional/fr/FR-083.md) | Transfer anonymous PostgreSQL event-owner authority to another browser                   | frontend, backend |
| [FR-084](functional/fr/FR-084.md) | Preserve blind availability privacy                                                      | frontend, backend |

## Quality Requirements

Read [`quality/README.md`](quality/README.md) before creating or changing a
quality requirement.

| ID                             | Title                                          | Components              |
| ------------------------------ | ---------------------------------------------- | ----------------------- |
| [QR-004](quality/qr/QR-004.md) | Restrict shared-event read access              | frontend, backend       |
| [QR-005](quality/qr/QR-005.md) | Preserve shared-event modification integrity   | frontend, backend       |
| [QR-006](quality/qr/QR-006.md) | Authenticate anonymous edit credentials        | frontend, backend       |
| [QR-007](quality/qr/QR-007.md) | Exclude secrets from diagnostics               | backend, infrastructure |
| [QR-008](quality/qr/QR-008.md) | Support accessible coordination flows          | frontend                |
| [QR-009](quality/qr/QR-009.md) | Respond promptly for timed events              | frontend, backend       |
| [QR-010](quality/qr/QR-010.md) | Respond promptly for days-only events          | frontend, backend       |
| [QR-011](quality/qr/QR-011.md) | Support large coordination workloads           | frontend, backend       |
| [QR-012](quality/qr/QR-012.md) | Reject unsafe deployment configuration         | infrastructure          |
| [QR-013](quality/qr/QR-013.md) | Diagnose failed requests without exposing data | backend, infrastructure |
| [QR-014](quality/qr/QR-014.md) | Authenticate cross-device access transfers     | frontend, backend       |

When a requirement is migrated, add its row to the matching table with a
stable relative link, for example `[FR-001](functional/fr/FR-001.md)`.
