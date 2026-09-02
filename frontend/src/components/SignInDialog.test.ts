// @vitest-environment happy-dom

import { flushPromises, mount } from "@vue/test-utils"
import { afterEach, beforeEach, describe, expect, it, vi } from "vitest"
import type * as UtilsModule from "@/utils"
import {
  buttonStubWithDisabled,
  mergeComponentStubs,
  nullStub,
  passThroughStub,
  vTextFieldStub,
} from "@/test/componentStubs"
import SignInDialog from "./SignInDialog.vue"

const { postMock, verifyOtpSignInMock } = vi.hoisted(() => ({
  postMock: vi.fn(),
  verifyOtpSignInMock: vi.fn(),
}))

vi.mock("@/utils", async () => {
  const actual = await vi.importActual<typeof UtilsModule>("@/utils")

  return {
    ...actual,
    post: postMock,
  }
})

vi.mock("@/utils/services/UserService", () => ({
  verifyOtpSignIn: verifyOtpSignInMock,
}))

const dialogStubs = mergeComponentStubs({
  "router-link": passThroughStub,
  "v-btn": buttonStubWithDisabled,
  "v-card": passThroughStub,
  "v-card-text": passThroughStub,
  "v-card-title": passThroughStub,
  "v-dialog": passThroughStub,
  "v-divider": nullStub,
  "v-icon": nullStub,
  "v-img": nullStub,
  "v-spacer": nullStub,
  "v-text-field": vTextFieldStub,
})

const mountDialog = () =>
  mount(SignInDialog, {
    props: {
      modelValue: true,
    },
    global: {
      stubs: dialogStubs,
    },
  })

const findTextFieldByPlaceholder = (
  wrapper: ReturnType<typeof mountDialog>,
  placeholder: string,
) => {
  const field = wrapper
    .findAllComponents(vTextFieldStub)
    .find((component) => component.props("placeholder") === placeholder)

  if (field == null) {
    throw new Error(`Expected text field with placeholder "${placeholder}"`)
  }

  return field
}

const findButtonByText = (
  wrapper: ReturnType<typeof mountDialog>,
  text: string,
) => {
  const button = wrapper
    .findAll("button")
    .find((candidate) => candidate.text().includes(text))

  if (button == null) {
    throw new Error(`Expected button containing "${text}"`)
  }

  return button
}

describe("SignInDialog", () => {
  afterEach(() => {
    vi.useRealTimers()
  })

  beforeEach(() => {
    postMock.mockReset()
    verifyOtpSignInMock.mockReset()
  })

  it("uses variant solo for all credential fields", async () => {
    const wrapper = mountDialog()

    expect(
      findTextFieldByPlaceholder(wrapper, "Enter your email...").props(
        "variant",
      ),
    ).toBe("solo")

    postMock.mockResolvedValueOnce({ isNewUser: false })
    postMock.mockResolvedValueOnce(undefined)

    await wrapper
      .get('input[placeholder="Enter your email..."]')
      .setValue("existing@example.com")
    await findButtonByText(wrapper, "Continue with Email").trigger("click")
    await flushPromises()

    expect(
      findTextFieldByPlaceholder(wrapper, "Enter 6-digit code...").props(
        "variant",
      ),
    ).toBe("solo")
  })

  it("offers account creation inline when sign-in cannot find an email", async () => {
    postMock.mockResolvedValueOnce({ isNewUser: true })

    const wrapper = mountDialog()

    await wrapper
      .get('input[placeholder="Enter your email..."]')
      .setValue("new@example.com")
    await findButtonByText(wrapper, "Continue with Email").trigger("click")
    await flushPromises()

    expect(wrapper.text()).toContain("Couldn’t find this account.")
    expect(wrapper.text()).toContain("Sign up")
    expect(wrapper.find('input[placeholder="First name"]').exists()).toBe(false)
    expect(postMock).toHaveBeenCalledTimes(1)
  })

  it("keeps email validation errors gating OTP send", async () => {
    const wrapper = mountDialog()

    await wrapper
      .get('input[placeholder="Enter your email..."]')
      .setValue("invalid-email")
    await findButtonByText(wrapper, "Continue with Email").trigger("click")

    expect(postMock).not.toHaveBeenCalled()
    expect(
      findTextFieldByPlaceholder(wrapper, "Enter your email...").props(
        "errorMessages",
      ),
    ).toBe("Please enter a valid email address.")
  })

  it("keeps plus-alias email validation gating OTP send", async () => {
    const wrapper = mountDialog()

    await wrapper
      .get('input[placeholder="Enter your email..."]')
      .setValue("name+alias@example.com")
    await findButtonByText(wrapper, "Continue with Email").trigger("click")

    expect(postMock).not.toHaveBeenCalled()
    expect(
      findTextFieldByPlaceholder(wrapper, "Enter your email...").props(
        "errorMessages",
      ),
    ).toBe("Email aliases with '+' are not allowed.")
  })

  it("does not send an OTP for an unknown email", async () => {
    postMock.mockResolvedValueOnce({ isNewUser: true })

    const wrapper = mountDialog()

    await wrapper
      .get('input[placeholder="Enter your email..."]')
      .setValue("new@example.com")
    await findButtonByText(wrapper, "Continue with Email").trigger("click")
    await flushPromises()

    expect(wrapper.text()).toContain("Couldn’t find this account.")
    expect(postMock).toHaveBeenCalledTimes(1)
  })

  it("keeps OTP verify gated until six digits are present", async () => {
    postMock.mockResolvedValueOnce({ isNewUser: false })
    postMock.mockResolvedValueOnce(undefined)
    verifyOtpSignInMock.mockResolvedValue({
      _id: "user-1",
      email: "existing@example.com",
    })

    const wrapper = mountDialog()

    await wrapper
      .get('input[placeholder="Enter your email..."]')
      .setValue("existing@example.com")
    await findButtonByText(wrapper, "Continue with Email").trigger("click")
    await flushPromises()

    const verifyButton = findButtonByText(wrapper, "Verify")
    expect(verifyButton.attributes("disabled")).toBeDefined()

    await verifyButton.trigger("click")
    expect(verifyOtpSignInMock).not.toHaveBeenCalled()

    await wrapper
      .get('input[placeholder="Enter 6-digit code..."]')
      .setValue("123456")

    expect(verifyButton.attributes("disabled")).toBeUndefined()

    await verifyButton.trigger("click")
    await flushPromises()

    expect(verifyOtpSignInMock).toHaveBeenCalledTimes(1)
  })

  it("emits the existing success contract after OTP verification", async () => {
    postMock.mockResolvedValueOnce({ isNewUser: false })
    postMock.mockResolvedValueOnce(undefined)
    verifyOtpSignInMock.mockResolvedValue({
      _id: "user-1",
      email: "existing@example.com",
    })

    const wrapper = mountDialog()

    await wrapper
      .get('input[placeholder="Enter your email..."]')
      .setValue("existing@example.com")
    await findButtonByText(wrapper, "Continue with Email").trigger("click")
    await flushPromises()
    await wrapper
      .get('input[placeholder="Enter 6-digit code..."]')
      .setValue("123456")
    await findButtonByText(wrapper, "Verify").trigger("click")
    await flushPromises()

    expect(wrapper.emitted("emailSignIn")).toEqual([
      [
        {
          _id: "user-1",
          email: "existing@example.com",
        },
      ],
    ])
    expect(wrapper.emitted("update:modelValue")).toContainEqual([false])
  })

  it("shows an error and report link when resending the code fails", async () => {
    vi.useFakeTimers()

    postMock.mockResolvedValueOnce({ isNewUser: false })
    postMock.mockResolvedValueOnce(undefined)
    postMock.mockRejectedValueOnce(
      new Error("OTP email service is not configured"),
    )

    const wrapper = mountDialog()

    await wrapper
      .get('input[placeholder="Enter your email..."]')
      .setValue("existing@example.com")
    await findButtonByText(wrapper, "Continue with Email").trigger("click")
    await flushPromises()

    const resendButton = findButtonByText(wrapper, "Resend code")
    expect(resendButton.attributes("disabled")).toBeDefined()

    await vi.advanceTimersByTimeAsync(30_000)
    await flushPromises()

    await resendButton.trigger("click")
    await flushPromises()

    expect(wrapper.text()).toContain(
      "We couldn’t resend the verification code.",
    )

    const reportLink = wrapper.get("a")
    expect(reportLink.text()).toContain("Report this problem")
    expect(reportLink.attributes("href")).toBeTruthy()
  })
})
