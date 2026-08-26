import { describe, expect, it } from 'vitest'
import { Linter } from 'eslint'
import markdown from '@eslint/markdown'
import noSplitSentence from './no-split-sentence.js'

const linter = new Linter({ configType: 'flat' })

const config = [
  {
    files: ['**/*.md'],
    plugins: {
      markdown,
      local: {
        rules: {
          'no-split-sentence': noSplitSentence,
        },
      },
    },
    language: 'markdown/gfm',
    languageOptions: {
      frontmatter: 'yaml',
    },
    rules: {
      'local/no-split-sentence': 'error',
    },
  },
]

function verify(code) {
  return linter.verify(code, config, 'fixture.md')
}

function fix(code) {
  return linter.verifyAndFix(code, config, 'fixture.md')
}

describe('no-split-sentence', () => {
  it('reports a prose sentence that continues across a physical line and fixes it', () => {
    const code =
      'The deployment finished early\nand the queue drained cleanly.\n'

    const messages = verify(code)
    expect(messages).toHaveLength(1)
    expect(messages[0].ruleId).toBe('local/no-split-sentence')
    expect(messages[0].messageId).toBe('sentenceContinues')
    expect(messages[0].fix).toBeDefined()

    const result = fix(code)
    expect(result.fixed).toBe(true)
    expect(result.messages).toHaveLength(0)
    expect(result.output).toBe(
      'The deployment finished early and the queue drained cleanly.\n',
    )
  })

  it('does not report lines that end with sentence-terminal punctuation', () => {
    const code =
      'First line ends here.\nSecond ends with energy!\nThird asks nicely?\nFourth quotes "the term."\n'

    expect(verify(code)).toHaveLength(0)
  })

  it('reports abbreviation endings as continuations and fixes them', () => {
    const code = 'The memo cites Dr.\nWu and the appendix.\n'

    const messages = verify(code)
    expect(messages).toHaveLength(1)
    expect(messages[0].fix).toBeDefined()

    const result = fix(code)
    expect(result.output).toBe('The memo cites Dr. Wu and the appendix.\n')
  })

  it('reports initialism endings as continuations and fixes them', () => {
    const code = 'The spec references U.S.\nfederal guidance today.\n'

    const messages = verify(code)
    expect(messages).toHaveLength(1)

    const result = fix(code)
    expect(result.output).toBe('The spec references U.S. federal guidance today.\n')
  })

  it('treats digit periods as sentence ends without reporting', () => {
    const code = 'Version 2.\nShips today.\n'

    expect(verify(code)).toHaveLength(0)
  })

  it('reports one diagnostic per soft-wrap boundary and joins all of them', () => {
    const code = 'Alpha began here\nbeta middle chunk\ngamma wrapped up.\n'

    const messages = verify(code)
    expect(messages).toHaveLength(2)

    const result = fix(code)
    expect(result.output).toBe('Alpha began here beta middle chunk gamma wrapped up.\n')
  })

  it('reports but does not autofix boundaries with unsafe edge characters', () => {
    const code = 'A well-\nknown fact appears.\n'

    const messages = verify(code)
    expect(messages).toHaveLength(1)
    expect(messages[0].fix).toBeUndefined()

    const result = fix(code)
    expect(result.fixed).toBe(false)
    expect(result.output).toBe(code)
  })

  it('excludes structural Markdown content from reporting', () => {
    const structural = [
      '---',
      'title: Wrapped',
      'prose sample',
      '---',
      '',
      '# Heading text that',
      'continues below',
      '',
      'Fenced code sample:',
      '',
      '```js',
      'const first = 1',
      'const second = 2',
      '```',
      '',
      '    Indented code line one',
      '    indented code line two continues',
      '',
      '| Column | Value |',
      '| ------ | ----- |',
      '| alpha wrapped | one |',
      '',
      '- Wrapped bullet item',
      '  continuing on the next line',
      '- Parent item',
      '  - Nested wrapped item',
      '    continuing deeper',
      '',
      '1. Ordered wrapped item',
      '   continues here',
      '',
      '- [ ] Task item wrapped',
      '  continues too',
      'Lazy unindented continuation',
      'of the preceding list item stays quiet',
      '',
      '> Quoted prose wrapped',
      '> across physical lines',
      '',
      'A definition style term',
      '---',
      '',
      'Double space hard break next  ',
      'still excluded here',
      '',
      'Backslash hard break next\\',
      'also excluded entirely',
      '',
      'Inline code span `wrapped',
      'across lines` stays quiet',
      '',
      'A [link label wrapped',
      'across lines](https://example.com) stays quiet',
      '',
      '*Emphasis span wrapped',
      'across lines* stays quiet',
      '',
    ].join('\n')

    expect(verify(structural)).toHaveLength(0)
  })
})
