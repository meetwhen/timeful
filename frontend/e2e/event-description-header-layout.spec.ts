import { expect, test, type Page } from "@playwright/test"
import {
  buildSpecificDateSeed,
  openEventPage,
  seedCanonicalTimedEvent,
} from "./helpers/timed-event-helpers"
import { Temporal } from "temporal-polyfill"

test.describe.configure({ mode: "serial" })

async function saveDescription(page: Page, description: string) {
  await page.getByRole("button", { name: /^\+\s*add description$/i }).click()
  const editor = page.locator('[role="textbox"]')
  await editor.fill(description)
  await page.locator(".event-description-save-button").click()
  await expect(page.locator(".event-description-edit-button")).toBeVisible()
}

test("multiline description keeps its edit button in the upper-right corner", async ({
  page,
  request,
}) => {
  const now = Temporal.Now.instant()
  const today = now.toZonedDateTimeISO("UTC").toPlainDate().toString()
  const seed = await seedCanonicalTimedEvent(request, {
    ...buildSpecificDateSeed({
      name: `Multiline description ${String(now.epochMilliseconds)}`,
      selectedDays: [today],
      activeSlots: [`${today}T09:00:00.000Z`],
      eventTimezone: "UTC",
      startTimeLocal: "09:00",
      endTimeLocal: "17:00",
      timeIncrementMinutes: 60,
    }),
    description: "",
  })

  await openEventPage(page, seed.shortId)
  await saveDescription(page, "First line\nSecond line\nThird line")

  const positions = await page.evaluate(() => {
    const shell = document.querySelector<HTMLElement>(".event-description-shell")
    const button = document.querySelector<HTMLElement>(
      ".event-description-edit-button",
    )

    if (!shell || !button) {
      return null
    }

    const shellRect = shell.getBoundingClientRect()
    const buttonRect = button.getBoundingClientRect()
    return {
      shellTop: shellRect.top,
      shellRight: shellRect.right,
      buttonTop: buttonRect.top,
      buttonRight: buttonRect.right,
    }
  })

  expect(positions).not.toBeNull()
  if (!positions) {
    throw new Error("Expected multiline description action metrics")
  }

  expect(Math.abs(positions.buttonTop - positions.shellTop)).toBeLessThanOrEqual(12)
  expect(Math.abs(positions.buttonRight - positions.shellRight)).toBeLessThanOrEqual(16)
})

test("editing preserves the description text width and wrapping", async ({
  page,
  request,
}) => {
  const now = Temporal.Now.instant()
  const today = now.toZonedDateTimeISO("UTC").toPlainDate().toString()
  const seed = await seedCanonicalTimedEvent(request, {
    ...buildSpecificDateSeed({
      name: `Description wrapping ${String(now.epochMilliseconds)}`,
      selectedDays: [today],
      activeSlots: [`${today}T09:00:00.000Z`],
      eventTimezone: "UTC",
      startTimeLocal: "09:00",
      endTimeLocal: "17:00",
      timeIncrementMinutes: 60,
    }),
    description: "",
  })

  await openEventPage(page, seed.shortId)
  await saveDescription(
    page,
    "A long description that wraps onto multiple lines without changing when editing begins.",
  )

  const preview = page.locator(".event-description-copy")
  const previewWidth = await preview.evaluate((element) => {
    const style = window.getComputedStyle(element)
    return (
      element.getBoundingClientRect().width -
      Number.parseFloat(style.paddingLeft) -
      Number.parseFloat(style.paddingRight)
    )
  })

  await page.locator(".event-description-edit-button").click()
  const editor = page.locator('[role="textbox"]')
  await expect(editor).toBeVisible()

  await expect(editor).toHaveText(
    "A long description that wraps onto multiple lines without changing when editing begins.",
  )
  const editorWidth = await editor.evaluate((element) =>
    element.getBoundingClientRect().width,
  )

  expect(editorWidth).toBeCloseTo(previewWidth, 0)
})

test("multiline editor keeps its cancel and save buttons in the upper-right corner", async ({
  page,
  request,
}) => {
  const now = Temporal.Now.instant()
  const today = now.toZonedDateTimeISO("UTC").toPlainDate().toString()
  const seed = await seedCanonicalTimedEvent(request, {
    ...buildSpecificDateSeed({
      name: `Description edit actions ${String(now.epochMilliseconds)}`,
      selectedDays: [today],
      activeSlots: [`${today}T09:00:00.000Z`],
      eventTimezone: "UTC",
      startTimeLocal: "09:00",
      endTimeLocal: "17:00",
      timeIncrementMinutes: 60,
    }),
    description: "",
  })

  await openEventPage(page, seed.shortId)
  await saveDescription(page, "First line\nSecond line\nThird line")
  await page.locator(".event-description-edit-button").click()

  const positions = await page.evaluate(() => {
    const shell = document.querySelector<HTMLElement>(
      ".event-description-edit-shell",
    )
    const cancel = document.querySelector<HTMLElement>(
      ".event-description-cancel-button",
    )
    const save = document.querySelector<HTMLElement>(
      ".event-description-save-button",
    )

    if (!shell || !cancel || !save) {
      return null
    }

    const shellRect = shell.getBoundingClientRect()
    const cancelRect = cancel.getBoundingClientRect()
    const saveRect = save.getBoundingClientRect()
    return {
      shellTop: shellRect.top,
      shellRight: shellRect.right,
      cancelTop: cancelRect.top,
      saveTop: saveRect.top,
      saveRight: saveRect.right,
    }
  })

  expect(positions).not.toBeNull()
  if (!positions) {
    throw new Error("Expected multiline editor action metrics")
  }

  expect(Math.abs(positions.cancelTop - positions.shellTop)).toBeLessThanOrEqual(12)
  expect(Math.abs(positions.saveTop - positions.shellTop)).toBeLessThanOrEqual(12)
  expect(Math.abs(positions.saveRight - positions.shellRight)).toBeLessThanOrEqual(16)
})

test("event description stays aligned to the left header column on desktop", async (
  { page, request },
  testInfo,
) => {
  test.skip(
    testInfo.project.name.includes("mobile"),
    "Desktop header actions are not rendered on mobile.",
  )

  const now = Temporal.Now.instant()
  const today = now.toZonedDateTimeISO("UTC").toPlainDate().toString()

  const seed = await seedCanonicalTimedEvent(request, {
    ...buildSpecificDateSeed({
      name: `Description layout ${String(now.epochMilliseconds)}`,
      selectedDays: [today],
      activeSlots: [`${today}T09:00:00.000Z`, `${today}T10:00:00.000Z`],
      eventTimezone: "UTC",
      startTimeLocal: "09:00",
      endTimeLocal: "17:00",
      timeIncrementMinutes: 60,
    }),
    description: "",
  })

  await openEventPage(page, seed.shortId)

  const addDescriptionButton = page.getByRole("button", {
    name: /^\+\s*add description$/i,
  })
  await expect(addDescriptionButton).toBeVisible()
  await expect(page.locator("#event-header-actions")).toBeVisible()

  interface MetadataActionAlignmentMetrics {
    buttonRowCenter: number
    secondaryActionsCenter: number
  }

  const metadataActionAlignment =
    await page.evaluate<MetadataActionAlignmentMetrics | null>(() => {
      const buttonRow = document.querySelector<HTMLElement>(
        "#event-header-button-row",
      )
      const secondaryActions = document.querySelector<HTMLElement>(
        "#event-header-meta-row > .desktop-event-header-actions",
      )

      if (!buttonRow || !secondaryActions) {
        return null
      }

      const buttonRowRect = buttonRow.getBoundingClientRect()
      const secondaryActionsRect = secondaryActions.getBoundingClientRect()

      return {
        buttonRowCenter: buttonRowRect.top + buttonRowRect.height / 2,
        secondaryActionsCenter:
          secondaryActionsRect.top + secondaryActionsRect.height / 2,
      }
    })

  expect(metadataActionAlignment).not.toBeNull()
  if (!metadataActionAlignment) {
    throw new Error("Expected metadata action alignment metrics")
  }
  expect(
    Math.abs(
      metadataActionAlignment.buttonRowCenter -
        metadataActionAlignment.secondaryActionsCenter,
    ),
  ).toBeLessThanOrEqual(8)

  interface AddButtonMetrics {
    buttonRowLeft: number
    addButtonLeft: number
    addButtonTextTop: number
    addButtonFontSize: string
    addButtonLineHeight: string
  }

  interface DescriptionLayoutMetrics {
    titleToMetaGap: number
    metaToDescriptionGap: number
    descriptionRight: number
    editorRight: number
    editorTop: number
    editorFontSize: string
    editorLineHeight: string
    actionsLeft: number
  }

  const addButtonMetrics = await page.evaluate<AddButtonMetrics | null>(() => {
    const buttonRow = document.querySelector<HTMLElement>(
      "#event-header-button-row",
    )
    const addButton = Array.from(
      document.querySelectorAll<HTMLElement>("button"),
    ).find(
      (button) =>
        button.textContent.replace(/\s+/g, " ").trim() === "+ Add description",
    )

    if (!buttonRow || !addButton) {
      return null
    }

    const buttonRowRect = buttonRow.getBoundingClientRect()
    const addButtonRect = addButton.getBoundingClientRect()
    const addButtonStyle = window.getComputedStyle(addButton)

    return {
      buttonRowLeft: buttonRowRect.left,
      addButtonLeft: addButtonRect.left,
      addButtonTextTop:
        addButtonRect.top + Number.parseFloat(addButtonStyle.paddingTop || "0"),
      addButtonFontSize: addButtonStyle.fontSize,
      addButtonLineHeight: addButtonStyle.lineHeight,
    }
  })

  expect(addButtonMetrics).not.toBeNull()
  if (!addButtonMetrics) {
    throw new Error("Expected add description button metrics")
  }
  expect(
    Math.abs(addButtonMetrics.buttonRowLeft - addButtonMetrics.addButtonLeft),
  ).toBeLessThanOrEqual(16)

  await addDescriptionButton.click()
  await expect(page.locator('[role="textbox"]')).toBeVisible()

  const layoutMetrics = await page.evaluate<DescriptionLayoutMetrics | null>(
    () => {
      const titleBlock = document.querySelector<HTMLElement>(
        "#event-header > .event-header-row:first-child > .tw-min-w-0.tw-flex-1 > div:first-child",
      )
      const metaRow = document.querySelector<HTMLElement>(
        "#event-header-meta-row",
      )
      const descriptionShell = document.querySelector<HTMLElement>(
        ".event-description-edit-shell",
      )
      const descriptionEditor =
        document.querySelector<HTMLElement>('[role="textbox"]')
      const headerActions = document.querySelector<HTMLElement>(
        "#event-header-actions",
      )

      if (
        !titleBlock ||
        !metaRow ||
        !descriptionShell ||
        !descriptionEditor ||
        !headerActions
      ) {
        return null
      }

      const titleRect = titleBlock.getBoundingClientRect()
      const metaRect = metaRow.getBoundingClientRect()
      const descriptionRect = descriptionShell.getBoundingClientRect()
      const editorRect = descriptionEditor.getBoundingClientRect()
      const actionsRect = headerActions.getBoundingClientRect()
      const editorStyle = window.getComputedStyle(descriptionEditor)

      return {
        titleToMetaGap: metaRect.top - titleRect.bottom,
        metaToDescriptionGap: descriptionRect.top - metaRect.bottom,
        descriptionRight: descriptionRect.right,
        editorRight: editorRect.right,
        editorTop: editorRect.top,
        editorFontSize: editorStyle.fontSize,
        editorLineHeight: editorStyle.lineHeight,
        actionsLeft: actionsRect.left,
      }
    },
  )

  expect(layoutMetrics).not.toBeNull()
  if (!layoutMetrics) {
    throw new Error("Expected description layout metrics")
  }
  expect(layoutMetrics.titleToMetaGap).toBeGreaterThan(0)
  expect(layoutMetrics.metaToDescriptionGap).toBeGreaterThan(0)
  expect(layoutMetrics.titleToMetaGap).toBeLessThanOrEqual(64)
  expect(layoutMetrics.metaToDescriptionGap).toBeLessThanOrEqual(64)
  expect(addButtonMetrics.addButtonFontSize).toBe(layoutMetrics.editorFontSize)
  expect(addButtonMetrics.addButtonLineHeight).toBe(
    layoutMetrics.editorLineHeight,
  )
  expect(
    Math.abs(addButtonMetrics.addButtonTextTop - layoutMetrics.editorTop),
  ).toBeLessThanOrEqual(8)
  expect(layoutMetrics.descriptionRight).toBeLessThanOrEqual(
    layoutMetrics.actionsLeft + 16,
  )
  expect(layoutMetrics.editorRight).toBeLessThanOrEqual(
    layoutMetrics.actionsLeft + 16,
  )
})
