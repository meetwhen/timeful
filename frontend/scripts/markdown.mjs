import { execFileSync, spawnSync } from 'node:child_process'
import { resolve } from 'node:path'
import { fileURLToPath } from 'node:url'

const frontend = fileURLToPath(new URL('..', import.meta.url))
const root = fileURLToPath(new URL('../..', import.meta.url))
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

const commands = {
  lint: {
    executable: new URL('../node_modules/.bin/eslint', import.meta.url).pathname,
    args: ['--config', `${root}eslint.config.ts`, '--cache', '--concurrency', '1'],
    cwd: root,
  },
  format: {
    executable: new URL('../node_modules/.bin/prettier', import.meta.url).pathname,
    args: ['--config', `${frontend}.prettierrc`, '--write'],
    cwd: frontend,
  },
  'format:check': {
    executable: new URL('../node_modules/.bin/prettier', import.meta.url).pathname,
    args: ['--config', `${frontend}.prettierrc`, '--check'],
    cwd: frontend,
  },
}

const selected = commands[command]

if (!selected) {
  throw new Error(`Unknown Markdown command: ${command}`)
}

const result = spawnSync(selected.executable, [...selected.args, ...additionalArgs, ...files], {
  cwd: selected.cwd,
  stdio: 'inherit',
})

process.exit(result.status ?? 1)
