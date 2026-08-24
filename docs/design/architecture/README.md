# Architecture Decision Records

Read `../README.md` for the design-record catalog and the relationship between
ADRs, requirements, specifications, and Backlog tasks.

## ADR Format

Each ADR has a permanent, zero-padded identifier and an ID-only filename:

```text
adr/ADR-001.md
```

IDs are assigned sequentially across the repository, not per component. An ADR
declares its scope in the title or context and includes context, decision,
consequences, quality attributes, and related requirements.

An ADR may link to the functional requirements (FRs) and quality requirements
(QRs) that it enables, constrains, or satisfies. FRs and QRs must remain
self-contained specifications and must not cite ADRs as normative dependencies.
Requirement provenance is recorded separately from the requirement statement.
