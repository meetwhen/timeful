import { execFileSync, spawnSync } from 'node:child_process'
import { existsSync, readFileSync, writeFileSync } from 'node:fs'
import { relative, resolve } from 'node:path'
import { fileURLToPath } from 'node:url'
import * as prettier from 'prettier'
import * as sentencesPerLine from '../prettier/markdown/sentences-per-line.js'

const root = fileURLToPath(new URL('..', import.meta.url))
const [command, ...additionalArgs] = process.argv.slice(2)
const files = execFileSync(
  'git',
  [
    'ls-files',
    '-z',
    '--',
    '*.md',
    ':(exclude)backlog/**',
    ':(exclude).opencode/**',
    ':(exclude).agents/**',
  ],
  { cwd: root },
)
  .toString()
  .split('\0')
  .filter(Boolean)
  .map((file) => resolve(root, file))
  .filter(existsSync)

if (!['lint', 'format', 'format:check'].includes(command)) {
  throw new Error(`Unknown Markdown command: ${command}`)
}

if (command === 'lint') {
  const result = spawnSync(
    new URL('../node_modules/.bin/eslint', import.meta.url).pathname,
    [
      '--config',
      `${root}eslint.config.ts`,
      '--cache',
      '--concurrency',
      '1',
      ...additionalArgs,
      ...files,
    ],
    { cwd: root, stdio: 'inherit' },
  )

  process.exit(result.status ?? 1)
}

if (additionalArgs.length > 0) {
  throw new Error(`${command} does not accept additional arguments`)
}

const unformattedFiles = []

for (const file of files) {
  const source = readFileSync(file, 'utf8')
  const config = await prettier.resolveConfig(file)
  const formatted = await prettier.format(source, {
    ...config,
    filepath: file,
    plugins: [sentencesPerLine],
  })

  if (formatted === source) continue

  const displayPath = relative(root, file)
  if (command === 'format') {
    writeFileSync(file, formatted)
    console.log(displayPath)
  } else {
    unformattedFiles.push(displayPath)
  }
}

if (command === 'format:check' && unformattedFiles.length > 0) {
  for (const file of unformattedFiles) console.error(file)
  console.error(
    `Code style issues found in ${unformattedFiles.length} file(s). Run npm run format:markdown.`,
  )
  process.exitCode = 1
}
