# ADR-009: Frontend Semantic Styling Tokens

Date: 2026-05-09

Status:

- Accepted

## Context

Shared visual states had drifted across component-local raw palette values and
framework-specific overrides.

## Decision

The frontend uses app-level semantic styling tokens as its shared visual-state
contract:

- Global CSS custom properties define shared semantic tokens with the
  `--timeful-*` prefix.
- Components consume semantic tokens for shared visual meanings instead of
  repeating raw state-palette literals.
- Token names describe intent, not a feature implementation.
- Raw palette values remain implementation details of the token-definition
  layer.
- Component-local variables may serve private concerns but do not replace
  shared semantic tokens.
- Framework overrides use semantic tokens when unavoidable and are exceptions,
  not the default styling mechanism.

## Consequences

Shared visual states can be reused across feature views and library-rendered
surfaces without making raw palette choices part of component contracts.
