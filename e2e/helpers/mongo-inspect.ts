import { execFileSync } from "node:child_process"
import { fileURLToPath } from "node:url"

const repositoryRoot = fileURLToPath(new URL("../../", import.meta.url))
const mongodbDatabase = "timeful-test"

const composeArguments = [
  "compose",
  "--env-file",
  ".env.test",
  "-f",
  "compose.yaml",
  "-f",
  "compose.test.yaml",
]

// Evaluates a mongosh expression against the retained-data MongoDB in the
// isolated stack and returns its printed output. It never reaches a
// development database because the mongo-test service is Playwright-owned.
function mongoEval(expression: string): string {
  return execFileSync(
    "docker",
    [
      ...composeArguments,
      "exec",
      "-T",
      "mongo-test",
      "mongosh",
      "--quiet",
      `mongodb://localhost:27017/${mongodbDatabase}`,
      "--eval",
      `print(${expression})`,
    ],
    { cwd: repositoryRoot, encoding: "utf8" },
  ).trim()
}

// Returns the retained MongoDB calendar account map keys for an account email.
// Calendar connections are PostgreSQL-authoritative, so a non-empty result from
// this helper means the retained document was used as a second authority.
export function retainedCalendarAccountKeys(email: string): string[] {
  const output = mongoEval(
    `JSON.stringify(Object.keys(db.users.findOne({email:${JSON.stringify(email)}})?.calendarAccounts ?? {}))`,
  )
  return JSON.parse(output) as string[]
}
