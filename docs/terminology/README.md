# Terminology

This directory is the canonical home for Timeful's controlled terminology.
`glossary.md` records the canonical spelling, capitalization, and concise definition of each controlled term.

## Authority

The glossary makes a term understandable at the point of use.
Each entry must link to its authoritative context: the ADR, requirement, or specification that defines the term's complete behavior, constraints, and exceptions.
The authoritative context wins if it conflicts with the glossary; correct the glossary rather than creating competing definitions.

Do not add a controlled term until it has a stable definition and an authoritative context.
A glossary entry may name approved aliases or rejected variants when they are needed to keep documentation consistent.

## Canonical Form in Prose

When prose refers to a defined concept, write its controlled term with the spelling and capitalization of its `glossary.md` heading.
Do not infer a general capitalization style: most entries use title case, but `Show all hours` does not.
Natural singular, plural, and possessive inflections are allowed when they preserve the capitalization of the corresponding glossary words, such as `Picked Date` from `Picked Dates` and `Event Timezone's` from `Event Timezone`.
This rule does not apply to generic lowercase prose, code blocks or spans, backticked UI labels, or the glossary's own definitional prose.

## Linking Controlled Terms

In normative documentation, link the first occurrence of each controlled term in every paragraph, list item, and table cell to its `glossary.md` anchor.
Link the term again when it first appears in a new paragraph, list item, or table cell.

Do not add the link in headings, code blocks, link labels, or text that is already linked.
An author may omit a repeated link when it would make a dense terminology table harder to read.

The rule makes the definition available where a reader needs it without requiring a link on every repeated use.
Documentation checks may verify that the first applicable occurrence links to the matching glossary anchor.
