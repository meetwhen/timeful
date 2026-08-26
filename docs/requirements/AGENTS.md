# Requirements Authoring

`README.md` is the canonical guide for this directory.
Read it and relevant existing requirements before creating, changing, or migrating a requirement.
For an `FR-*` record, also read `functional/README.md`.
For a `QR-*` record, also read `quality/README.md`.

## Migrating From Backlog

- Treat backlog tasks as source material, not as requirement records.
  Extract
  the durable expected behavior and discard implementation plans, investigation
  notes, bug history, and completion details.
- Recover applicability context from the task, linked artifacts, and current
  behavior.
  Do not generalize a requirement beyond the page, [Event Kind](../terminology/glossary.md#event-kind), mode,
  role, or state the source establishes.
- Preserve the requirement's permanent ID and ID-only filename.
  Use the YAML
  front matter defined in `README.md`, including every component responsible
  for enforcing the behavior.
- Keep cross-component behavior in one requirement with multiple `components`;
  do not duplicate the requirement per component.
- Add or update the corresponding row in the README index with a stable
  relative link.

## Review Checklist

- Does the record conform to the README's format, status, component, and index
  conventions?
