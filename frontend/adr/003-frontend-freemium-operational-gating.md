# ADR-003: Frontend Freemium Operational Gating

Date: 2026-05-05

Status:

- Accepted

## Context

Frontend freemium behavior spans advertising, access checks, and upgrade
surfaces. Scattered environment reads could make those paths interpret the
operational state differently or confuse it with persisted product premium
state.

## Decision

The frontend has one centralized freemium operational boundary:

- `VITE_ENABLE_FREEMIUM` controls frontend operational gating.
- Shared freemium helpers own environment parsing and viewer-access semantics.
- Persisted owner and user premium state remains distinct from operational
  viewer access.
- Components and stores use the shared helpers rather than directly reading
  freemium environment values.

## Product Behavior

The disabled-freemium outcomes are defined by [FR-070](../../docs/requirements/functional/fr/FR-070.md),
[FR-071](../../docs/requirements/functional/fr/FR-071.md), and [FR-072](../../docs/requirements/functional/fr/FR-072.md).

## Consequences

Freemium behavior has one operational interpretation while real product state
continues to describe persisted account and owner data.
