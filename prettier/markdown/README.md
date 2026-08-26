# Local Markdown sentences-per-line formatter

This wrapper preserves the upstream `prettier-plugin-sentences-per-line` behavior for prose and leaves Markdown table rows intact.
The upstream plugin inserts sentence breaks inside table cells, which invalidates rows because Markdown tables require each row to occupy one physical source line.
The wrapper restores a single whitespace node for every upstream sentence break inside a table before printing.
`scripts/markdown.mjs` invokes Prettier through its JavaScript API because Prettier's CLI loads its built-in Markdown printer after configured plugins and ignores sentence-per-line printer overrides.
