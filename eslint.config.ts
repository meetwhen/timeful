import markdown from '@eslint/markdown'
import sentencesPerLine from 'eslint-plugin-sentences-per-line'

export default [
  {
    files: ['**/*.md'],
    ignores: ['backlog/**', '.opencode/**', '.agents/**'],
    plugins: {
      markdown,
      'sentences-per-line': sentencesPerLine,
    },
    language: 'markdown/gfm',
    rules: {
      'sentences-per-line/one': 'error',
    },
  },
]
