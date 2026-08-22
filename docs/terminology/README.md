# Terminology

This directory is the canonical home for Timeful's controlled terminology.
`glossary.md` records the canonical spelling, capitalization, and concise
definition of each controlled term.

## Authority

The glossary makes a term understandable at the point of use. Each entry must
link to its authoritative context: the ADR, requirement, or specification that
defines the term's complete behavior, constraints, and exceptions. The
authoritative context wins if it conflicts with the glossary; correct the
glossary rather than creating competing definitions.

Do not add a controlled term until it has a stable definition and an
authoritative context. A glossary entry may name approved aliases or rejected
variants when they are needed to keep documentation consistent.

## Linking Controlled Terms

In normative documentation, link the first occurrence of each controlled term
in every paragraph, list item, and table cell to its `glossary.md` anchor. Link
the term again when it first appears in a new paragraph, list item, or table
cell.

Do not add the link in headings, code blocks, link labels, or text that is
already linked. An author may omit a repeated link when it would make a dense
terminology table harder to read.

The rule makes the definition available where a reader needs it without
requiring a link on every repeated use. Documentation checks may verify that
the first applicable occurrence links to the matching glossary anchor.
