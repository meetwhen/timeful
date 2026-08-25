# CAND-167

### Source

> - [x] Make compose files fully configurable via required values from corresponding .env files

[Source lines 535-535](../../../../backlog/backlog.md#L535-L535)

### Candidate behavior

Compose configuration obtains required values from its corresponding `.env` files.

### Applicability

Actor: operator or developer. Location: Compose configuration. Event kind: any. Interaction mode: starting services. Viewport: not applicable. State: configured environment. Exclusions: unspecified optional values.

### Classification

candidate QR

### Existing Requirements and Confidence

QR-012 requires staging and production to reject missing or unsafe required configuration, but does not require Compose or `.env` files. Confidence: inferred.

### Disposition

Hold as an installability/configuration candidate; consolidate only if the required environments and failure behavior are defined.

### Open Questions

Which Compose files and environments are in scope, and must missing values prevent startup?
