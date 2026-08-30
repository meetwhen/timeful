import { vi, type MockInstance } from "vitest"
import { Temporal } from "temporal-polyfill"

type FakeNowValue = string | Temporal.Instant

const clockMethods = [
  "instant",
  "plainDateISO",
  "plainDateTimeISO",
  "plainTimeISO",
  "zonedDateTimeISO",
] as const

type ClockMethod = (typeof clockMethods)[number]

type NowClockFn = (timeZone?: string) => unknown

let activeSpies: MockInstance[] = []

export const restoreFakeTemporalNow = () => {
  for (const spy of activeSpies) {
    spy.mockRestore()
  }
  activeSpies = []
}

export const setFakeTemporalNow = (value: FakeNowValue) => {
  restoreFakeTemporalNow()
  const instant = typeof value === "string" ? Temporal.Instant.from(value) : value
  const timeZoneId = Temporal.Now.timeZoneId()
  const now = Temporal.Now as unknown as Record<ClockMethod, NowClockFn>

  for (const method of clockMethods) {
    const spy = vi.spyOn(now, method).mockImplementation((timeZone?: string) => {
      if (method === "instant") {
        return instant
      }
      const zoned = instant.toZonedDateTimeISO(timeZone ?? timeZoneId)
      if (method === "plainDateISO") {
        return zoned.toPlainDate()
      }
      if (method === "plainDateTimeISO") {
        return zoned.toPlainDateTime()
      }
      if (method === "plainTimeISO") {
        return zoned.toPlainTime()
      }
      return zoned
    })
    activeSpies.push(spy)
  }
}
