import * as base from 'prettier-plugin-sentences-per-line'

const tableNodeTypes = new Set(['table', 'tableRow', 'tableCell'])

function restoreTableWhitespace(node, insideTable = false) {
  const nestedInTable = insideTable || tableNodeTypes.has(node.type)
  if (!Array.isArray(node.children)) return

  for (let index = 0; index < node.children.length; index += 1) {
    const child = node.children[index]
    if (nestedInTable && child.type === 'sentenceBreak') {
      // The upstream plugin removed this whitespace when it inserted the break.
      node.children[index] = { type: 'whitespace', value: ' ' }
      continue
    }
    restoreTableWhitespace(child, nestedInTable)
  }
}

export const options = base.options
export const parsers = base.parsers
export const printers = {
  mdast: {
    ...base.printers.mdast,
    async preprocess(ast, options) {
      const processed = await base.printers.mdast.preprocess(ast, options)
      restoreTableWhitespace(processed)
      return processed
    },
  },
}
