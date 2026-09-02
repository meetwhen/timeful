/*
 * Node >= 22.4 exposes an experimental built-in localStorage on globalThis and
 * emits `ExperimentalWarning: localStorage is not available because
 * --localstorage-file was not provided` when it is touched in the default node
 * Vitest environment. The codebase intentionally treats missing storage as a
 * valid state, so this warning is test noise. Only that exact warning is
 * dropped here; every other warning is re-dispatched to the original default
 * handlers unchanged.
 *
 * Mechanism note: the internal webstorage warning does not respect a patched
 * `process.emitWarning` (the default stderr print happens independently of the
 * public wrapper), and it cannot be suppressed by merely adding a
 * `process.on("warning")` listener because Node's default handler is an
 * independent listener that always runs. The reliable suppression point is to
 * replace the process "warning" listener set with a filter that swallows only
 * this warning and forwards all others to the captured default listeners.
 *
 * Removal note: delete this file and its vitest setupFiles entry once the
 * minimum supported Node version no longer emits the warning. Verify with
 * `node -e 'typeof globalThis.localStorage'` printing nothing to stderr.
 * Do not provide --localstorage-file instead: that would make
 * `typeof localStorage` return "object" and silently stop covering the
 * storage-unavailable fallback paths in node-environment tests. Do not use
 * --disable-warning=ExperimentalWarning either: that would hide every other
 * experimental warning as well.
 */

const localStorageWarningMessage =
  "localStorage is not available because --localstorage-file was not provided."

if (typeof process !== "undefined") {
  const defaultWarningListeners = process.listeners("warning")
  process.removeAllListeners("warning")
  process.on("warning", (warning) => {
    if (
      warning.name === "ExperimentalWarning" &&
      warning.message === localStorageWarningMessage
    ) {
      return
    }
    for (const listener of defaultWarningListeners) {
      listener(warning)
    }
  })
}
