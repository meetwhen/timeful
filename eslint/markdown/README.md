# Local Markdown split-sentence prototype

This folder contains the experimental `no-split-sentence` ESLint rule for Markdown prose and its fixture suite.

## Purpose

The rule detects prose sentences that continue across physical Markdown source lines without terminal punctuation on the earlier line.
It complements `sentences-per-line/one`, which already prevents multiple sentences from sharing one line.
TASK-0076 found that external sentence linters are unsafe for this repository, so detection uses the existing `@eslint/markdown` AST integration instead of line heuristics.

## Detection boundary

Only paragraphs whose direct AST parent is the document root are eligible.
Headings, front matter, fenced code, indented code, tables, blockquotes, HTML blocks, and list items never produce reports because they are different node types or live inside other containers.
Lazy unindented continuations belong to their list item in the AST, which closes the gap that made `textlint-rule-idiomatic-lines` unsafe in TASK-0076.01.
Inside eligible paragraphs the rule inspects soft wraps only, because hard breaks made by trailing backslashes or two spaces become dedicated break nodes.
Inline structures such as code spans, links, and emphasis that wrap across lines hold their own text nodes, so those boundaries are skipped conservatively.

A boundary is reported when the earlier line does not end with confident sentence-terminal punctuation.
Digit periods such as version numbers count as sentence ends and are never reported.
Whitelisted abbreviations such as `Dr.` and initialisms such as `U.S.` still report as continuations.

## Safe-fix boundary

The autofix joins an offending boundary by replacing the newline and surrounding padding with a single space, which preserves rendered Markdown meaning.
A boundary is fixed only when both edges are plain characters, so tails ending in hyphens and heads starting with Markdown syntax stay unfixed while still reporting.

## Running the fixtures

Run `npm run test:markdown-rules` at the repository root to execute the fixture suite.

## Future enablement

TASK-0075 owns normalizing the existing corpus before any enforcement begins.
After that work completes, the enforcement task can flip `local/no-split-sentence` to `error` in the root `eslint.config.ts`, where the rule is registered but explicitly off today.
Enablement should pair the rule with the `frontmatter` language option used by the fixture configuration so front matter stays excluded deterministically.
