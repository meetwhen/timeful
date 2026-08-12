import { expect, test } from "@playwright/test"
import {
  openEditDialog,
  openEventPage,
  seedCanonicalTimedEvent,
} from "./helpers/timed-event-helpers"
import { Temporal } from "temporal-polyfill"

test.describe.configure({ mode: "serial" })

test("dates-only event Edit event opens the dates-only editor", async ({
  page,
  request,
}) => {
  const now = Temporal.Now.instant()
  const today = now.toZonedDateTimeISO("UTC").toPlainDate().toString()
  const tomorrow = now
    .toZonedDateTimeISO("UTC")
    .toPlainDate()
    .add({ days: 1 })
    .toString()

  const seed = await seedCanonicalTimedEvent(request, {
    name: `Dates-only edit dialog ${String(now.epochMilliseconds)}`,
    type: "specific_dates",
    daysOnly: true,
    dates: [`${today}T00:00:00.000Z`, `${tomorrow}T00:00:00.000Z`],
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
  await expect(page.locator("#edit-event-btn")).toBeVisible()

  const editorCard = await openEditDialog(page)
  await expect(
    page.getByRole("dialog").getByText("Edit event", { exact: true })
  ).toBeVisible()
  await expect(editorCard.getByText("What dates might work?")).toBeVisible()
  await expect(editorCard.getByText("Drag to select multiple dates")).toBeVisible()
  await expect(editorCard.getByText("What times might work?")).not.toBeVisible()
})

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
  await expect(
    page.getByRole("button", { name: /^Schedule event$/i }),
  ).toBeVisible()

  const startOnMondayToggle = page.locator("#start-calendar-on-monday-toggle")
  await expect(startOnMondayToggle).toBeVisible()

  const moreOptions = page.locator("#desktop-header-more-options")
  await expect(moreOptions).not.toBeVisible()

  if (testInfo.project.name === "chromium-desktop") {
    const startOnMondaySwitch = page.locator(
      "#desktop-header-start-calendar-on-monday .v-input",
    )
    await page.getByRole("button", { name: /^\+\s*add description$/i }).click()
    const descriptionEditor = page.locator('[role="textbox"]')
    await expect(descriptionEditor).toBeVisible()
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
        (addAvailabilityBox.x + addAvailabilityBox.width / 2) -
          (startOnMondayBox.x + startOnMondayBox.width / 2),
      ),
    ).toBeLessThanOrEqual(2)

    const [descriptionBox, scheduleBox] = await Promise.all([
      descriptionEditor.boundingBox(),
      page.getByRole("button", { name: /^Schedule event$/i }).boundingBox(),
    ])

    if (descriptionBox === null || scheduleBox === null) {
      throw new Error("Expected the description editor and Schedule event button")
    }

    expect(descriptionBox.x + descriptionBox.width).toBeLessThanOrEqual(
      scheduleBox.x + 16,
    )
  }
})

test("dates-only Responses heading top edge stays aligned with the grid top edge", async ({
  page,
  request,
}, testInfo) => {
  test.skip(
    testInfo.project.name === "chromium-mobile",
    "Desktop-only sidebar layout",
  )

  const now = Temporal.Now.instant()
  const today = now.toZonedDateTimeISO("UTC").toPlainDate().toString()
  const tomorrow = now
    .toZonedDateTimeISO("UTC")
    .toPlainDate()
    .add({ days: 1 })
    .toString()

  const seed = await seedCanonicalTimedEvent(request, {
    name: `Days-only responses alignment ${String(now.epochMilliseconds)}`,
    type: "specific_dates",
    daysOnly: true,
    dates: [`${today}T00:00:00.000Z`, `${tomorrow}T00:00:00.000Z`],
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

  const monthGrid = page.locator(".schedule-overlap-days-only-grid__month")
  await expect(monthGrid).toBeVisible()

  const responsesHeading = page.getByText("Responses", { exact: true })
  await expect(responsesHeading).toBeVisible()

  const [monthBox, headingBox] = await Promise.all([
    monthGrid.boundingBox(),
    responsesHeading.boundingBox(),
  ])

  if (monthBox === null || headingBox === null) {
    throw new Error(
      "Expected the days-only grid and Responses heading to have boxes",
    )
  }

  const gridTopMinusHeadingTop = monthBox.y - headingBox.y
  const gridRightToSidebarLeft = headingBox.x - (monthBox.x + monthBox.width)

  expect(gridTopMinusHeadingTop).toBeGreaterThanOrEqual(0)
  expect(gridTopMinusHeadingTop).toBeLessThanOrEqual(8)
  expect(gridRightToSidebarLeft).toBeGreaterThanOrEqual(16)
  expect(gridRightToSidebarLeft).toBeLessThanOrEqual(20)
})

test("dates-only calendar cells are twice as wide as they are tall", async ({
  page,
  request,
}) => {
  const now = Temporal.Now.instant()
  const today = now.toZonedDateTimeISO("UTC").toPlainDate().toString()
  const tomorrow = now
    .toZonedDateTimeISO("UTC")
    .toPlainDate()
    .add({ days: 1 })
    .toString()

  const seed = await seedCanonicalTimedEvent(request, {
    name: `Days-only rectangular cells ${String(now.epochMilliseconds)}`,
    type: "specific_dates",
    daysOnly: true,
    dates: [`${today}T00:00:00.000Z`, `${tomorrow}T00:00:00.000Z`],
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

  const cellBox = await page.locator(".schedule-overlap-days-only-grid .timeslot").first().boundingBox()
  if (cellBox === null) {
    throw new Error("Expected a dates-only calendar cell to have a box")
  }

  expect(cellBox.width / cellBox.height).toBeCloseTo(2, 1)
})

test("dates-only grid keeps gutters across narrow viewports", async ({
  page,
  request,
}, testInfo) => {
  test.skip(
    testInfo.project.name === "chromium-mobile",
    "Runs an explicit viewport matrix on desktop Chromium",
  )

  const now = Temporal.Now.instant()
  const today = now.toZonedDateTimeISO("UTC").toPlainDate().toString()
  const tomorrow = now
    .toZonedDateTimeISO("UTC")
    .toPlainDate()
    .add({ days: 1 })
    .toString()

  const seed = await seedCanonicalTimedEvent(request, {
    name: `Days-only narrow layout ${String(now.epochMilliseconds)}`,
    type: "specific_dates",
    daysOnly: true,
    dates: [`${today}T00:00:00.000Z`, `${tomorrow}T00:00:00.000Z`],
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

  for (const width of [320, 390, 410, 480, 639, 640]) {
    await page.setViewportSize({ width, height: 900 })
    await openEventPage(page, seed.shortId)

    const monthGrid = page.locator(".schedule-overlap-days-only-grid__month")
    await expect(monthGrid).toBeVisible()
    await expect
      .poll(() =>
        page.evaluate(() =>
          document.documentElement.scrollWidth <= window.innerWidth
        )
      )
      .toBe(true)

    const monthBox = await monthGrid.boundingBox()
    if (monthBox === null) {
      throw new Error("Expected the days-only grid to have a box")
    }

    expect(monthBox.x).toBeGreaterThanOrEqual(16)
    expect(monthBox.x + monthBox.width).toBeLessThanOrEqual(width - 16)

    if (width < 640) {
      const sidebar = page.locator(".schedule-overlap-sidebar")
      await expect(sidebar).toBeVisible()
      const sidebarBox = await sidebar.boundingBox()

      if (sidebarBox === null) {
        throw new Error("Expected the sidebar to have a box")
      }

      expect(sidebarBox.y).toBeGreaterThanOrEqual(monthBox.y + monthBox.height)
      expect(sidebarBox.x + sidebarBox.width).toBeLessThanOrEqual(width)
      continue
    }

    const responsesHeading = page.getByText("Responses", { exact: true })
    await expect(responsesHeading).toBeVisible()

    const headingBox = await responsesHeading.boundingBox()

    if (headingBox === null) {
      throw new Error("Expected the Responses heading to have a box")
    }

    const gridRightToSidebarLeft = headingBox.x - (monthBox.x + monthBox.width)
    expect(gridRightToSidebarLeft).toBeGreaterThanOrEqual(16)
    expect(gridRightToSidebarLeft).toBeLessThanOrEqual(20)
  }
})
