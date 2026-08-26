import markdown from '@eslint/markdown'
import sentencesPerLine from 'eslint-plugin-sentences-per-line'
import noSplitSentence from './eslint/markdown/no-split-sentence.js'

export default [
  {
    files: ['**/*.md'],
    ignores: ['backlog/**', '.opencode/**', '.agents/**'],
    plugins: {
      markdown,
      'sentences-per-line': sentencesPerLine,
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
      'sentences-per-line/one': 'error',
      'local/no-split-sentence': 'off',
    },
  },
]
