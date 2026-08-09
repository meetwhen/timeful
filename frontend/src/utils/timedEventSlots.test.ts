import { describe, expect, it } from "vitest"
import { Temporal } from "temporal-polyfill"
import { UTC } from "@/constants"
import type { Event } from "@/types"
import {
  buildTimedDateSeeds,
  generateTimedSlotsForDay,
  getEventEnabledSlots,
  getEventWindowRangeSlots,
  getTimedSlotForMembershipDay,
  getTimedSlotCoverage,
  getTimedSlotGeneration,
  getTimedWeekDays,
  mergeActiveSlotsByMembershipDay,
  normalizeActiveSlots,
  projectSlotsToMembershipDays,
  projectSlotsToLocalDays,
} from "./timedEventSlots"

// Cross-side derivation fixtures. Mirrored as identical expected instant
// lists in server/routes/event_timed_slots_test.go so Go and Temporal
// derivation are asserted against the same instants (Q3 parity suite).

const zdt = (value: string): Temporal.ZonedDateTime =>
  Temporal.ZonedDateTime.from(value)

const instant = (slot: Temporal.ZonedDateTime): string =>
  slot.toInstant().toString()

const contractEvent = (
  overrides: Partial<
    Pick<
      Event,
      | "daysOnly"
      | "timedRecurrence"
      | "eventTimezone"
      | "slotGeneration"
      | "activeSlots"
      | "timeIncrement"
    >
  >
): Event =>
  ({
    daysOnly: false,
    timedRecurrence: {
      kind: "specific_dates",
      selectedDays: [Temporal.PlainDate.from("2026-01-05")],
      selectedDaysOfWeek: [],
      startOnMonday: false,
    },
    eventTimezone: "America/New_York",
    slotGeneration: {
      startTimeLocal: Temporal.PlainTime.from("09:00"),
      endTimeLocal: Temporal.PlainTime.from("10:00"),
      timeIncrement: Temporal.Duration.from({ minutes: 15 }),
    },
    activeSlots: [],
    ...overrides,
  })

describe("timedEventSlots", () => {
  it("getTimedSlotGeneration returns correct values for midnight start (00:00-01:00)", () => {
    const result = getTimedSlotGeneration({
      slotGeneration: {
        startTimeLocal: Temporal.PlainTime.from("00:00"),
        endTimeLocal: Temporal.PlainTime.from("01:00"),
        timeIncrement: Temporal.Duration.from({ minutes: 15 }),
      },
      activeSlots: [],
    })

    expect(result.startTimeLocal.toString()).toBe("00:00:00")
    expect(result.endTimeLocal.toString()).toBe("01:00:00")
    expect(result.timeIncrement.total("minutes")).toBe(15)
  })

  it("getTimedSlotGeneration falls back to timeSeed and activeSlots when slotGeneration is missing", () => {
    const result = getTimedSlotGeneration({
      timeSeed: zdt("2026-01-05T10:00:00+00:00[UTC]"),
      activeSlots: [zdt("2026-01-05T10:00:00+00:00[UTC]")],
    })

    // Falls back to the timeSeed time + increment = 10:00 + 15min = 10:15
    expect(result.startTimeLocal.toString()).toBe("10:00:00")
    expect(result.endTimeLocal.toString()).toBe("10:15:00")
  })

  it("getTimedSlotGeneration falls back to defaults when slotGeneration, timeSeed, and activeSlots are missing", () => {
    const result = getTimedSlotGeneration({})

    expect(result.startTimeLocal.toString()).toBe("09:00:00")
    expect(result.endTimeLocal.toString()).toBe("17:00:00")
  })

  it("drops out-of-domain active slots during normalization", () => {
    const enabledSlots = [
      zdt("2026-01-05T09:00:00+00:00[UTC]"),
      zdt("2026-01-05T09:15:00+00:00[UTC]"),
    ]
    const activeSlots = [
      zdt("2026-01-05T09:15:00+00:00[UTC]"),
      zdt("2026-01-05T10:00:00+00:00[UTC]"),
    ]

    expect(
      normalizeActiveSlots({ enabledSlots, activeSlots }).activeSlots.map((slot) =>
        slot.toString()
      )
    ).toEqual(["2026-01-05T09:15:00+00:00[UTC]"])
  })

  it("projects slots to shifted local membership days", () => {
    const slots = [
      zdt("2026-01-05T00:30:00+00:00[UTC]"),
      zdt("2026-01-06T00:30:00+00:00[UTC]"),
    ]

    expect(
      projectSlotsToLocalDays(slots, "America/Los_Angeles").map((day) =>
        day.toString()
      )
    ).toEqual(["2026-01-04", "2026-01-05"])
  })

  it("projects slots to membership days with wrapped local windows", () => {
    const slots = [
      zdt("2026-01-05T23:00:00+00:00[UTC]"),
      zdt("2026-01-06T00:30:00+00:00[UTC]"),
    ]

    expect(
      projectSlotsToMembershipDays({
        slots,
        timeZone: UTC,
        slotGeneration: {
          startTimeLocal: Temporal.PlainTime.from("23:00:00"),
          endTimeLocal: Temporal.PlainTime.from("01:00:00"),
          timeIncrement: Temporal.Duration.from({ minutes: 30 }),
        },
      }).map((day) => day.toString())
    ).toEqual(["2026-01-05"])
  })

  it("preserves the prior subset and defaults newly added membership days to fully active", () => {
    const merged = mergeActiveSlotsByMembershipDay({
      priorEnabledSlots: [
        zdt("2026-01-05T09:00:00+00:00[UTC]"),
        zdt("2026-01-05T09:15:00+00:00[UTC]"),
      ],
      priorActiveSlots: [zdt("2026-01-05T09:15:00+00:00[UTC]")],
      nextEnabledSlots: [
        zdt("2026-01-05T09:00:00+00:00[UTC]"),
        zdt("2026-01-05T09:15:00+00:00[UTC]"),
        zdt("2026-01-06T09:00:00+00:00[UTC]"),
        zdt("2026-01-06T09:15:00+00:00[UTC]"),
      ],
      timeZone: UTC,
      slotGeneration: {
        startTimeLocal: Temporal.PlainTime.from("09:00:00"),
        endTimeLocal: Temporal.PlainTime.from("09:30:00"),
        timeIncrement: Temporal.Duration.from({ minutes: 15 }),
      },
      priorMembershipDays: [Temporal.PlainDate.from("2026-01-05")],
      nextMembershipDays: [
        Temporal.PlainDate.from("2026-01-05"),
        Temporal.PlainDate.from("2026-01-06"),
      ],
    })

    expect(merged.map((slot) => slot.toString())).toEqual([
      "2026-01-05T09:15:00+00:00[UTC]",
      "2026-01-06T09:00:00+00:00[UTC]",
      "2026-01-06T09:15:00+00:00[UTC]",
    ])
  })

  it("generates slots across DST spring-forward gaps", () => {
    const slots = generateTimedSlotsForDay({
      day: Temporal.PlainDate.from("2026-03-08"),
      timeZone: "America/Los_Angeles",
      slotGeneration: {
        startTimeLocal: Temporal.PlainTime.from("01:30:00"),
        endTimeLocal: Temporal.PlainTime.from("03:30:00"),
        timeIncrement: Temporal.Duration.from({ minutes: 15 }),
      },
    })

    expect(slots.map((slot) => slot.toInstant().toString())).toEqual([
      "2026-03-08T09:30:00Z",
      "2026-03-08T09:45:00Z",
      "2026-03-08T10:00:00Z",
      "2026-03-08T10:15:00Z",
    ])
  })

  it("generates slots across DST fall-back overlaps", () => {
    const slots = generateTimedSlotsForDay({
      day: Temporal.PlainDate.from("2026-11-01"),
      timeZone: "America/Los_Angeles",
      slotGeneration: {
        startTimeLocal: Temporal.PlainTime.from("01:00:00"),
        endTimeLocal: Temporal.PlainTime.from("02:00:00"),
        timeIncrement: Temporal.Duration.from({ minutes: 30 }),
      },
    })

    expect(slots.map((slot) => slot.toInstant().toString())).toEqual([
      "2026-11-01T08:00:00Z",
      "2026-11-01T08:30:00Z",
      "2026-11-01T09:00:00Z",
      "2026-11-01T09:30:00Z",
    ])
  })

  it("generates wrapped cross-midnight slots into the next local day", () => {
    const slots = generateTimedSlotsForDay({
      day: Temporal.PlainDate.from("2026-01-05"),
      timeZone: UTC,
      slotGeneration: {
        startTimeLocal: Temporal.PlainTime.from("23:00:00"),
        endTimeLocal: Temporal.PlainTime.from("01:00:00"),
        timeIncrement: Temporal.Duration.from({ minutes: 30 }),
      },
    })

    expect(slots.map((slot) => slot.toInstant().toString())).toEqual([
      "2026-01-05T23:00:00Z",
      "2026-01-05T23:30:00Z",
      "2026-01-06T00:00:00Z",
      "2026-01-06T00:30:00Z",
    ])
  })

  it("maps membership-day grid rows to canonical UTC instants in the event timezone", () => {
    expect(
      getTimedSlotForMembershipDay({
        day: Temporal.PlainDate.from("2026-05-28"),
        timeZone: UTC,
        absoluteMinutes: 9 * 60,
      }).toInstant().toString()
    ).toBe("2026-05-28T09:00:00Z")

    expect(
      getTimedSlotForMembershipDay({
        day: Temporal.PlainDate.from("2026-01-05"),
        timeZone: "America/Los_Angeles",
        absoluteMinutes: 23 * 60 + 30,
      }).toInstant().toString()
    ).toBe("2026-01-06T07:30:00Z")
  })

  it("builds canonical membership seeds for wrapped specific-date windows", () => {
    const seeds = buildTimedDateSeeds({
      daysOnly: false,
      activeSlots: [],
      eventTimezone: UTC,
      slotGeneration: {
        startTimeLocal: Temporal.PlainTime.from("23:00:00"),
        endTimeLocal: Temporal.PlainTime.from("01:00:00"),
        timeIncrement: Temporal.Duration.from({ minutes: 30 }),
      },
      timedRecurrence: {
        kind: "specific_dates",
        selectedDays: [Temporal.PlainDate.from("2026-01-05")],
        selectedDaysOfWeek: [],
        startOnMonday: false,
      },
      type: "specific_dates",
    })

    expect(seeds.map((slot) => slot.toString())).toEqual([
      "2026-01-05T23:00:00+00:00[UTC]",
    ])
  })

  it("computes monotonic slot coverage for sparse and wrapped legacy domains", () => {
    expect(
      getTimedSlotCoverage({
        times: [
          zdt("2026-01-05T09:30:00+00:00[UTC]"),
          zdt("2026-01-05T12:00:00+00:00[UTC]"),
        ],
        eventTimezone: UTC,
        slotGeneration: {
          startTimeLocal: Temporal.PlainTime.from("09:00:00"),
          endTimeLocal: Temporal.PlainTime.from("13:00:00"),
          timeIncrement: Temporal.Duration.from({ minutes: 30 }),
        },
      })
    ).toEqual({
      minTime: Temporal.PlainTime.from("09:30:00"),
      maxTime: Temporal.PlainTime.from("12:30:00"),
    })

    expect(
      getTimedSlotCoverage({
        times: [
          zdt("2026-01-05T23:30:00+00:00[UTC]"),
          zdt("2026-01-06T00:00:00+00:00[UTC]"),
        ],
        eventTimezone: UTC,
        slotGeneration: {
          startTimeLocal: Temporal.PlainTime.from("23:00:00"),
          endTimeLocal: Temporal.PlainTime.from("01:00:00"),
          timeIncrement: Temporal.Duration.from({ minutes: 30 }),
        },
      })
    ).toEqual({
      minTime: Temporal.PlainTime.from("23:30:00"),
      maxTime: Temporal.PlainTime.from("00:30:00"),
    })
  })
})

describe("getEventEnabledSlots", () => {
  it("returns an empty domain for days-only events", () => {
    expect(
      getEventEnabledSlots({ daysOnly: true, activeSlots: [] })
    ).toEqual([])
  })

  it("returns the legacy folded domain when there is no contract", () => {
    const times = [
      zdt("2026-01-05T09:00:00+00:00[UTC]"),
      zdt("2026-01-05T09:15:00+00:00[UTC]"),
    ]

    expect(getEventEnabledSlots({ times }).map(instant)).toEqual([
      "2026-01-05T09:00:00Z",
      "2026-01-05T09:15:00Z",
    ])
  })

  it("derives the full civil day per picked date, independent of the window (matches the Go fixture)", () => {
    const event = contractEvent({})

    const derived = getEventEnabledSlots(event)
    expect(derived).toHaveLength(96)
    expect(derived[0]?.toInstant().toString()).toBe("2026-01-05T05:00:00Z")
    expect(derived[48]?.toInstant().toString()).toBe("2026-01-05T17:00:00Z")
    expect(derived.at(-1)?.toInstant().toString()).toBe("2026-01-06T04:45:00Z")
    expect(
      new Set(
        derived.map((slot) =>
          slot.withTimeZone("America/New_York").toPlainDate().toString(),
        ),
      ),
    ).toEqual(new Set(["2026-01-05"]))

    // The window does not bound the enabled domain.
    const wrappedWindow = contractEvent({
      slotGeneration: {
        startTimeLocal: Temporal.PlainTime.from("22:30"),
        endTimeLocal: Temporal.PlainTime.from("01:30"),
        timeIncrement: Temporal.Duration.from({ minutes: 15 }),
      },
    })
    expect(getEventEnabledSlots(wrappedWindow)).toEqual(derived)
  })

  it("anchors the weekly domain on the earliest active instant (matches the Go fixture)", () => {
    const event = contractEvent({
      timedRecurrence: {
        kind: "weekly",
        selectedDays: [],
        selectedDaysOfWeek: [1, 3],
        startOnMonday: true,
      },
      eventTimezone: "America/Los_Angeles",
      slotGeneration: {
        startTimeLocal: Temporal.PlainTime.from("09:00"),
        endTimeLocal: Temporal.PlainTime.from("11:00"),
        timeIncrement: Temporal.Duration.from({ minutes: 30 }),
      },
      activeSlots: [
        zdt("2026-01-05T17:00:00+00:00[UTC]"),
        zdt("2026-01-05T17:30:00+00:00[UTC]"),
        zdt("2026-01-07T17:00:00+00:00[UTC]"),
      ],
    })

    const derived = getEventEnabledSlots(event)
    expect(derived).toHaveLength(96)
    expect(derived[0]?.toInstant().toString()).toBe("2026-01-05T08:00:00Z")
    expect(derived[18]?.toInstant().toString()).toBe("2026-01-05T17:00:00Z")
    expect(derived[48]?.toInstant().toString()).toBe("2026-01-07T08:00:00Z")
    expect(derived[66]?.toInstant().toString()).toBe("2026-01-07T17:00:00Z")
    expect(derived.at(-1)?.toInstant().toString()).toBe("2026-01-08T07:30:00Z")
  })

  it("anchors the weekly domain from a Saturday anchor on the following week (matches the Go fixture)", () => {
    const event = contractEvent({
      timedRecurrence: {
        kind: "weekly",
        selectedDays: [],
        selectedDaysOfWeek: [1, 3],
        startOnMonday: true,
      },
      eventTimezone: "America/Los_Angeles",
      slotGeneration: {
        startTimeLocal: Temporal.PlainTime.from("09:00"),
        endTimeLocal: Temporal.PlainTime.from("10:00"),
        timeIncrement: Temporal.Duration.from({ minutes: 15 }),
      },
      activeSlots: [zdt("2026-01-10T09:00:00+00:00[UTC]")],
    })

    const derived = getEventEnabledSlots(event)
    expect(derived).toHaveLength(192)
    expect(derived[0]?.toInstant().toString()).toBe("2026-01-12T08:00:00Z")
    expect(derived[96]?.toInstant().toString()).toBe("2026-01-14T08:00:00Z")
    expect(derived.at(-1)?.toInstant().toString()).toBe("2026-01-15T07:45:00Z")
  })

  it("filters Sunday indexes out of startOnMonday weekly domains (matches the Go fixture)", () => {
    const event = contractEvent({
      timedRecurrence: {
        kind: "weekly",
        selectedDays: [],
        selectedDaysOfWeek: [0, 1, 7],
        startOnMonday: true,
      },
      eventTimezone: "America/Los_Angeles",
      slotGeneration: {
        startTimeLocal: Temporal.PlainTime.from("09:00"),
        endTimeLocal: Temporal.PlainTime.from("10:00"),
        timeIncrement: Temporal.Duration.from({ minutes: 15 }),
      },
      activeSlots: [zdt("2026-01-11T09:00:00+00:00[UTC]")],
    })

    const derived = getEventEnabledSlots(event)
    expect(derived).toHaveLength(192)
    expect(derived[0]?.toInstant().toString()).toBe("2026-01-11T08:00:00Z")
    expect(derived[96]?.toInstant().toString()).toBe("2026-01-12T08:00:00Z")
    expect(derived.at(-1)?.toInstant().toString()).toBe("2026-01-13T07:45:00Z")
  })

  it("derives the NY spring-forward gap day as a full civil day (matches the Go fixture)", () => {
    const event = contractEvent({
      eventTimezone: "America/New_York",
      timedRecurrence: {
        kind: "specific_dates",
        selectedDays: [Temporal.PlainDate.from("2026-03-08")],
        selectedDaysOfWeek: [],
        startOnMonday: false,
      },
      slotGeneration: {
        startTimeLocal: Temporal.PlainTime.from("01:30"),
        endTimeLocal: Temporal.PlainTime.from("03:00"),
        timeIncrement: Temporal.Duration.from({ minutes: 15 }),
      },
    })

    const derived = getEventEnabledSlots(event)
    expect(derived).toHaveLength(92)
    expect(derived[0]?.toInstant().toString()).toBe("2026-03-08T05:00:00Z")
    expect(derived.at(-1)?.toInstant().toString()).toBe("2026-03-09T03:45:00Z")
  })

  it("derives the NY fall-back overlap day as a full civil day (matches the Go fixture)", () => {
    const event = contractEvent({
      eventTimezone: "America/New_York",
      timedRecurrence: {
        kind: "specific_dates",
        selectedDays: [Temporal.PlainDate.from("2026-11-01")],
        selectedDaysOfWeek: [],
        startOnMonday: false,
      },
      slotGeneration: {
        startTimeLocal: Temporal.PlainTime.from("01:00"),
        endTimeLocal: Temporal.PlainTime.from("02:30"),
        timeIncrement: Temporal.Duration.from({ minutes: 15 }),
      },
    })

    const derived = getEventEnabledSlots(event)
    expect(derived).toHaveLength(100)
    expect(derived[0]?.toInstant().toString()).toBe("2026-11-01T04:00:00Z")
    expect(derived.at(-1)?.toInstant().toString()).toBe("2026-11-02T04:45:00Z")
    // 01:30 EDT (05:30Z) and 01:30 EST (06:30Z) both occur on the overlap day.
    expect(derived).toContainEqual(zdt("2026-11-01T05:30:00+00:00[UTC]"))
    expect(derived).toContainEqual(zdt("2026-11-01T06:30:00+00:00[UTC]"))
  })

  it("derives the Lord Howe 30-minute DST days as full civil days (matches the Go fixtures)", () => {
    const spring = contractEvent({
      eventTimezone: "Australia/Lord_Howe",
      timedRecurrence: {
        kind: "specific_dates",
        selectedDays: [Temporal.PlainDate.from("2026-10-04")],
        selectedDaysOfWeek: [],
        startOnMonday: false,
      },
      slotGeneration: {
        startTimeLocal: Temporal.PlainTime.from("02:00"),
        endTimeLocal: Temporal.PlainTime.from("03:30"),
        timeIncrement: Temporal.Duration.from({ minutes: 15 }),
      },
    })

    const springDerived = getEventEnabledSlots(spring)
    expect(springDerived).toHaveLength(94)
    expect(springDerived[0]?.toInstant().toString()).toBe(
      "2026-10-03T13:30:00Z",
    )
    expect(springDerived.at(-1)?.toInstant().toString()).toBe(
      "2026-10-04T12:45:00Z",
    )

    // 00:00 through the next 00:00 exclusive derives identically on both
    // sides because the domain never starts inside the 02:00 fold.
    const fall = contractEvent({
      eventTimezone: "Australia/Lord_Howe",
      timedRecurrence: {
        kind: "specific_dates",
        selectedDays: [Temporal.PlainDate.from("2026-04-05")],
        selectedDaysOfWeek: [],
        startOnMonday: false,
      },
      slotGeneration: {
        startTimeLocal: Temporal.PlainTime.from("00:30"),
        endTimeLocal: Temporal.PlainTime.from("03:00"),
        timeIncrement: Temporal.Duration.from({ minutes: 15 }),
      },
    })

    const fallDerived = getEventEnabledSlots(fall)
    expect(fallDerived).toHaveLength(98)
    expect(fallDerived[0]?.toInstant().toString()).toBe(
      "2026-04-04T13:00:00Z",
    )
    expect(fallDerived.at(-1)?.toInstant().toString()).toBe(
      "2026-04-05T13:15:00Z",
    )
  })

  it("ignores wrapped slot-generation windows and derives the full civil day (matches the Go fixture)", () => {
    const event = contractEvent({
      eventTimezone: UTC,
      timedRecurrence: {
        kind: "specific_dates",
        selectedDays: [Temporal.PlainDate.from("2026-05-10")],
        selectedDaysOfWeek: [],
        startOnMonday: false,
      },
      slotGeneration: {
        startTimeLocal: Temporal.PlainTime.from("22:30"),
        endTimeLocal: Temporal.PlainTime.from("01:30"),
        timeIncrement: Temporal.Duration.from({ minutes: 30 }),
      },
    })

    const derived = getEventEnabledSlots(event)
    expect(derived).toHaveLength(48)
    expect(derived[0]?.toInstant().toString()).toBe("2026-05-10T00:00:00Z")
    expect(derived.at(-1)?.toInstant().toString()).toBe("2026-05-10T23:30:00Z")
  })

  it("deduplicates and sorts the derived domain (matches the Go fixture)", () => {
    const event = contractEvent({
      eventTimezone: UTC,
      timedRecurrence: {
        kind: "specific_dates",
        selectedDays: [
          Temporal.PlainDate.from("2026-01-05"),
          Temporal.PlainDate.from("2026-01-05"),
        ],
        selectedDaysOfWeek: [],
        startOnMonday: false,
      },
      slotGeneration: {
        startTimeLocal: Temporal.PlainTime.from("09:00"),
        endTimeLocal: Temporal.PlainTime.from("09:30"),
        timeIncrement: Temporal.Duration.from({ minutes: 15 }),
      },
    })

    const derived = getEventEnabledSlots(event)
    expect(derived).toHaveLength(96)
    expect(derived[0]?.toInstant().toString()).toBe("2026-01-05T00:00:00Z")
    expect(derived.at(-1)?.toInstant().toString()).toBe("2026-01-05T23:45:00Z")
  })

  it("anchors empty-active weekly events on the current week's selected days", () => {
    const event = contractEvent({
      timedRecurrence: {
        kind: "weekly",
        selectedDays: [],
        selectedDaysOfWeek: [1, 3],
        startOnMonday: true,
      },
      eventTimezone: "America/Los_Angeles",
      slotGeneration: {
        startTimeLocal: Temporal.PlainTime.from("09:00"),
        endTimeLocal: Temporal.PlainTime.from("10:00"),
        timeIncrement: Temporal.Duration.from({ minutes: 15 }),
      },
      activeSlots: [],
    })

    const weekDays = getTimedWeekDays({
      activeSlots: [],
      timeZone: "America/Los_Angeles",
      timedRecurrence: {
        kind: "weekly",
        selectedDays: [],
        selectedDaysOfWeek: [1, 3],
        startOnMonday: true,
      },
    })

    const derived = getEventEnabledSlots(event)
    const derivedDays = derived.map((slot) =>
      slot.withTimeZone("America/Los_Angeles").toPlainDate().toString()
    )

    expect(weekDays.map((day) => day.toString())).toEqual([
      ...new Set(derivedDays),
    ])
    expect(derived).toHaveLength(192)
  })
})

describe("getEventWindowRangeSlots", () => {
  it("derives the persisted slot-generation window range for range-event discrimination", () => {
    const event = contractEvent({
      eventTimezone: "America/New_York",
      slotGeneration: {
        startTimeLocal: Temporal.PlainTime.from("09:00"),
        endTimeLocal: Temporal.PlainTime.from("10:00"),
        timeIncrement: Temporal.Duration.from({ minutes: 15 }),
      },
    })

    expect(getEventWindowRangeSlots(event).map(instant)).toEqual([
      "2026-01-05T14:00:00Z",
      "2026-01-05T14:15:00Z",
      "2026-01-05T14:30:00Z",
      "2026-01-05T14:45:00Z",
    ])
  })

  it("returns an empty range when there is no canonical contract", () => {
    expect(getEventWindowRangeSlots({ daysOnly: true })).toEqual([])
    expect(getEventWindowRangeSlots({ activeSlots: [] })).toEqual([])
  })

  it("derives the wrapped window range across midnight per membership day", () => {
    const event = contractEvent({
      eventTimezone: UTC,
      slotGeneration: {
        startTimeLocal: Temporal.PlainTime.from("23:00"),
        endTimeLocal: Temporal.PlainTime.from("01:00"),
        timeIncrement: Temporal.Duration.from({ minutes: 30 }),
      },
    })

    expect(getEventWindowRangeSlots(event).map(instant)).toEqual([
      "2026-01-05T23:00:00Z",
      "2026-01-05T23:30:00Z",
      "2026-01-06T00:00:00Z",
      "2026-01-06T00:30:00Z",
    ])
  })
})