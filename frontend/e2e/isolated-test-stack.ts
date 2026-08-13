import { execFile } from "node:child_process"
import path from "node:path"
import { fileURLToPath } from "node:url"
import { promisify } from "node:util"
import { Temporal } from "temporal-polyfill"
import { loadEnv } from "vite"

const execFileAsync = promisify(execFile)
const frontendRoot = path.dirname(path.dirname(fileURLToPath(import.meta.url)))
const repositoryRoot = path.dirname(frontendRoot)
const testEnv = loadEnv("test", repositoryRoot, "")

function composeArguments(...args: string[]): string[] {
  return [
    "compose",
    "--env-file",
    path.join(repositoryRoot, ".env.test"),
    "-f",
    path.join(repositoryRoot, "compose.yaml"),
    "-f",
    path.join(repositoryRoot, "compose.test.yaml"),
    ...args,
  ]
}

async function runCompose(...args: string[]): Promise<void> {
  await execFileAsync("docker", composeArguments(...args), {
    cwd: repositoryRoot,
  })
}

async function waitForHealthcheck(): Promise<void> {
  const url = "http://127.0.0.1:3005/api/health"
  const deadline = Temporal.Now.instant().epochMilliseconds + 120_000

  while (Temporal.Now.instant().epochMilliseconds < deadline) {
    try {
      const response = await fetch(url)
      if (response.ok) {
        return
      }
    } catch {
      // The server can still be compiling or waiting for MongoDB.
    }
    await new Promise((resolve) => setTimeout(resolve, 500))
  }

  throw new Error(`Timed out waiting for isolated E2E server at ${url}`)
}

async function start(): Promise<void> {
  await runCompose("up", "-d", "--build", "mongo-test", "server-test")
  await waitForHealthcheck()
}

async function stop(): Promise<void> {
  if (testEnv.TEST_MONGO_PERSIST.toLowerCase() === "false") {
    await runCompose("down", "-v")
    return
  }

  await runCompose("stop", "server-test")
}

export default async function isolatedTestStack(): Promise<
  () => Promise<void>
> {
  await start()
  return stop
}
