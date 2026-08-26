const knownAbbreviations = new Set([
  'mr',
  'mrs',
  'ms',
  'dr',
  'prof',
  'sr',
  'jr',
  'st',
  'mt',
  'vs',
  'etc',
  'cf',
  'al',
  'approx',
  'fig',
  'inc',
  'ltd',
])

const closingPairsPattern = /[)"'”’\]]+$/
const terminalPunctuationPattern = /[.!?]$/
const digitPeriodPattern = /\d\.$/
const initialismPattern = /(?:^|[\s("'])(?:[A-Za-z]\.)+$/
const finalWordPeriodPattern = /([A-Za-z]+)\.$/
const trailingPaddingPattern = /[ \t]+$/
const leadingPaddingPattern = /^[ \t]+/
const fixableTailEndPattern = /[A-Za-z0-9,;:%)"'”’\].!?]$/
const fixableHeadStartPattern = /^[A-Za-z0-9("'“‘[\]]/

function analyzeTail(rawTail) {
  const tail = rawTail.replace(trailingPaddingPattern, '')
  if (!tail) {
    return null
  }
  const stripped = tail.replace(closingPairsPattern, '')
  if (!terminalPunctuationPattern.test(stripped)) {
    return { tailLength: tail.length }
  }
  if (digitPeriodPattern.test(stripped)) {
    return null
  }
  if (initialismPattern.test(stripped)) {
    return { tailLength: tail.length }
  }
  const wordMatch = stripped.match(finalWordPeriodPattern)
  if (wordMatch && knownAbbreviations.has(wordMatch[1].toLowerCase())) {
    return { tailLength: tail.length }
  }
  return null
}

export default {
  meta: {
    type: 'layout',
    docs: {
      description:
        'Disallow Markdown prose sentences that continue across physical source lines',
      recommended: false,
    },
    fixable: 'whitespace',
    messages: {
      sentenceContinues:
        'Expected the sentence to end on this line, but it continues on the next physical line.',
    },
    schema: [],
  },
  create(context) {
    const { sourceCode } = context
    return {
      paragraph(node, parent) {
        if (!parent || parent.type !== 'root') {
          return
        }
        for (const child of node.children) {
          if (child.type !== 'text' || !child.value.includes('\n')) {
            continue
          }
          const value = child.value
          const base = child.position.start.offset
          let lineStart = 0
          for (;;) {
            const newline = value.indexOf('\n', lineStart)
            if (newline === -1) {
              break
            }
            const nextNewline = value.indexOf('\n', newline + 1)
            const rawTail = value.slice(lineStart, newline)
            const head =
              nextNewline === -1
                ? value.slice(newline + 1)
                : value.slice(newline + 1, nextNewline)
            const analysis = analyzeTail(rawTail)
            if (analysis) {
              const trimmedTail = rawTail.replace(trailingPaddingPattern, '')
              const tailEnd = base + lineStart + trimmedTail.length
              const headLead = head.match(leadingPaddingPattern)?.[0].length ?? 0
              const headStart = base + newline + 1 + headLead
              const fixable =
                fixableTailEndPattern.test(trimmedTail) &&
                fixableHeadStartPattern.test(head.trimStart())
              context.report({
                loc: {
                  start: sourceCode.getLocFromIndex(tailEnd),
                  end: sourceCode.getLocFromIndex(Math.min(headStart, tailEnd + 1)),
                },
                messageId: 'sentenceContinues',
                ...(fixable
                  ? {
                      fix(fixer) {
                        return fixer.replaceTextRange([tailEnd, headStart], ' ')
                      },
                    }
                  : {}),
              })
            }
            lineStart = newline + 1
          }
        }
      },
    }
  },
}
