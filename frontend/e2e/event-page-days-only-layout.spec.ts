import { expect, test } from "@playwright/test"
import {
  openEventPage,
  seedCanonicalTimedEvent,
} from "./helpers/timed-event-helpers"
import { Temporal } from "temporal-polyfill"

test.describe.configure({ mode: "serial" })

test("days-only event page without responses shows an inline Start on Monday switch aligned with Add availability", async ({
  page,
  request,
}, testInfo) => {
  test.skip(
    testInfo.project.name === "chromium-mobile",
    "Desktop-only header layout",
  )

  const now = Temporal.Now.instant()
  const today = now.toZonedDateTimeISO("UTC").toPlainDate().toString()
  const tomorrow = now
    .toZonedDateTimeISO("UTC")
    .toPlainDate()
    .add({ days: 1 })
    .toString()

  const seed = await seedCanonicalTimedEvent(request, {
    name: `Days-only layout test ${String(now.epochMilliseconds)}`,
    type: "specific_dates",
    daysOnly: true,
    dates: [`${today}T00:00:00.000Z`, `${tomorrow}T00:00:00.000Z`],
    enabledSlots: [],
    eventTimezone: "UTC",
    slotGeneration: {
      startTimeLocal: "09:00",
      endTimeLocal: "17:00",
      timeIncrementMinutes: 60,
    },
    timedRecurrence: {
      kind: "specific_dates",
      selectedDays: [today, tomorrow],
      selectedDaysOfWeek: [],
      startOnMonday: false,
    },
  })

  await openEventPage(page, seed.shortId)

  const addAvailabilityBtn = page.locator("#desktop-primary-availability-btn")
  await expect(addAvailabilityBtn).toBeVisible()
  await expect(addAvailabilityBtn).toHaveText(/Add availability/i)

  const startOnMondayToggle = page.locator("#start-calendar-on-monday-toggle")
  await expect(startOnMondayToggle).toBeVisible()

  const moreOptions = page.locator("#desktop-header-more-options")
  await expect(moreOptions).not.toBeVisible()

  if (testInfo.project.name === "chromium-desktop") {
    const startOnMondaySwitch = page.locator(
      "#desktop-header-start-calendar-on-monday .v-input",
    )
    const [addAvailabilityBox, startOnMondayBox] = await Promise.all([
      addAvailabilityBtn.boundingBox(),
      startOnMondaySwitch.boundingBox(),
    ])

    if (addAvailabilityBox === null || startOnMondayBox === null) {
      throw new Error(
        "Expected Add availability and Start on Monday to have boxes",
      )
    }

    expect(
      Math.abs(
        addAvailabilityBox.x +
          addAvailabilityBox.width -
          (startOnMondayBox.x + startOnMondayBox.width),
      ),
    ).toBeLessThanOrEqual(2)
  }
})