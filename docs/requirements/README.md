# Requirements

This directory is the canonical record of Timeful product requirements.
Requirements state durable, verifiable behavior or quality constraints. They
are not implementation tasks, bug reports, design investigations, or decision
records.

<h2>Table of contents</h2>

- [Layout](#layout)
- [Requirement Format](#requirement-format)
- [Related Artifacts](#related-artifacts)
- [Functional Requirements](#functional-requirements)
- [Quality Requirements](#quality-requirements)

## Layout

- `functional/` contains functional requirements.
- `quality/` contains quality requirements.
- `functional-requirements.md` and `quality-requirements.md` are temporary
  migration sources. Do not add new requirements to them.

Each requirement has its own file. Use an ID-only filename so links remain
stable when its title changes:

```text
functional/FR-001.md
quality/QR-001.md
```

The requirement ID and filename are permanent. Titles may change.

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

List the components responsible for enforcing the requirement. A requirement
that spans components remains one file and lists each applicable component;
do not duplicate it.

Every requirement has a concise title and is atomic: it is independently
understandable, decidable, and verifiable. A requirement can contain several
acceptance criteria when they jointly specify one behavior, but unrelated
obligations belong in separate requirement files.

## Related Artifacts

- Functional requirements (`FR-*`) state required system behavior.
- Quality requirements (`QR-*`) state required quality attributes or runtime
  constraints.
- Specifications (`SPEC-*`) describe a feature or domain in broader context
  and may link to multiple requirements.
- Architecture decision records (ADRs) record architectural decisions.
- Backlog tasks track work to investigate, implement, or verify requirements.

## Functional Requirements

| ID | Title | Components |
| --- | --- | --- |
| _No canonical per-file requirements have been migrated yet._ | | |

## Quality Requirements

| ID | Title | Components |
| --- | --- | --- |
| _No canonical per-file requirements have been migrated yet._ | | |

When a requirement is migrated, add its row to the matching table with a
stable relative link, for example `[FR-001](functional/FR-001.md)`.
