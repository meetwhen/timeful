import { describe, expect, it } from "vitest"
import { buildAllAvailableNote } from "./useScheduleOverlapViewModels"

const baseInput = {
  editingAvailability: false,
  loadingResponses: false,
  fetchedResponsesCount: 2,
  max: 1,
  respondentsLength: 3,
  daysOnly: false,
  isGroup: false,
}

describe("buildAllAvailableNote", () => {
  it("builds the timed-event note when best availability is below the respondent count", () => {
    expect(buildAllAvailableNote(baseInput)).toBe(
      "Note: There's no time when all 3 respondents are available.",
    )
  })

  it("uses the day wording for days-only events", () => {
    expect(buildAllAvailableNote({ ...baseInput, daysOnly: true })).toBe(
      "Note: There's no day when all 3 respondents are available.",
    )
  })

  it("uses the members wording for group events", () => {
    expect(buildAllAvailableNote({ ...baseInput, isGroup: true })).toBe(
      "Note: There's no time when all 3 members are available.",
    )
    expect(
      buildAllAvailableNote({ ...baseInput, daysOnly: true, isGroup: true }),
    ).toBe("Note: There's no day when all 3 members are available.")
  })

  it("is hidden while editing availability", () => {
    expect(
      buildAllAvailableNote({ ...baseInput, editingAvailability: true }),
    ).toBeNull()
  })

  it("is hidden while responses are loading", () => {
    expect(
      buildAllAvailableNote({ ...baseInput, loadingResponses: true }),
    ).toBeNull()
  })

  it("is hidden when no responses have been fetched", () => {
    expect(
      buildAllAvailableNote({ ...baseInput, fetchedResponsesCount: 0 }),
    ).toBeNull()
  })

  it("is hidden when a slot exists where everyone is available", () => {
    expect(
      buildAllAvailableNote({ ...baseInput, max: 3, respondentsLength: 3 }),
    ).toBeNull()
    expect(
      buildAllAvailableNote({ ...baseInput, max: 4, respondentsLength: 3 }),
    ).toBeNull()
  })
})
