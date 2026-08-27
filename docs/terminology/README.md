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
This rule does not apply to generic lowercase prose, code blocks or spans, or backticked UI labels.

## Linking and Styling Terms

In normative documentation, style a controlled term according to its position.

Link the first occurrence of each controlled term in every paragraph, list item, and table cell to its `glossary.md` anchor.
Use the inflection that fits the surrounding sentence as the link label, such as `[Availability Responses](glossary.md#availability-response)`; the anchor always resolves to the canonical heading even when the label differs in number or adds a possessive.
Style further occurrences of the same term within the same paragraph, list item, or table cell in bold, covering the entire inflected form inside the markers, such as `**Event Timezone's**`; never split the markup around an inflection, such as `**Event Timezone**'s`.

Do not link or style terms in headings, code blocks or spans, backticked UI labels, existing link labels, or text that is already linked.
A glossary entry may name its own term in bold without linking to itself.
An author may omit a repeated link when it would make a dense terminology table harder to read.

The first link makes the definition available where a reader needs it, and bold keeps repeats recognizable without requiring a link everywhere.
The glossary applies this same rule to cross-references in its definitions.
