import { markdown, sentencesPerLine } from './frontend/eslint/markdown'

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
