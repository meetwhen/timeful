# Graphify Artifact Policy

Graphify produces derived output under `graphify-out/`.
The repository tracks
only `graphify-out/cache/semantic/`.

## Tracked Semantic Index

Track the semantic index so collaborators can use a shared semantic view of the
repository immediately, without first rebuilding it locally.

## Untracked Derived Output

Do not track community labels (`.graphify_labels.json`) or their membership
signatures (`.graphify_labels.json.sig`).
Graphify regenerates these files after
commits, and versioning them creates distracting worktree churn without
preserving source material.

Do not track syntactic output.
It is quick to index locally, so committing it
does not provide enough collaboration value to justify the repository churn.
