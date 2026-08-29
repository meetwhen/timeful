// @vitest-environment happy-dom

import { beforeEach, describe, expect, it, vi } from "vitest"
import { createLocalStorageMock } from "@/test/localStorage"
import {
  cookieConsentVersion,
  getCookieConsent,
  getCookieConsentPreferences,
  setCookieConsent,
} from "./cookie_utils"

describe("cookie_utils", () => {
  beforeEach(() => {
    vi.stubGlobal("localStorage", createLocalStorageMock())
    vi.spyOn(console, "error").mockImplementation(() => undefined)
  })

  it("returns default preferences when consent is missing or corrupt", () => {
    expect(getCookieConsent()).toBeNull()
    expect(getCookieConsentPreferences()).toEqual({
      necessary: true,
      analytics: true,
    })

    localStorage.setItem("cookieConsent", "{not-json")

    expect(getCookieConsent()).toBeNull()
    expect(getCookieConsentPreferences()).toEqual({
      necessary: true,
      analytics: true,
    })
  })

  it("normalizes persisted consent data and increments the shared version", () => {
    const versionBeforeSave = cookieConsentVersion.value

    const consent = setCookieConsent({
      necessary: false,
      analytics: 0,
    })

    expect(consent.preferences).toEqual({
      necessary: true,
      analytics: false,
    })
    expect(getCookieConsentPreferences()).toEqual({
      necessary: true,
      analytics: false,
    })
    expect(cookieConsentVersion.value).toBe(versionBeforeSave + 1)
  })
})
