// @vitest-environment happy-dom

import { mount } from "@vue/test-utils"
import { describe, expect, it } from "vitest"
import EditingAvailabilityAs from "./EditingAvailabilityAs.vue"
import {
  buildEditingAvailabilityAsViewModel,
  scheduleOverlapGlobalStubs,
} from "./scheduleOverlapTestUtils"

describe("EditingAvailabilityAs", () => {
  const dialogContentStubs = {
    "v-dialog": { template: "<div><slot /></div>" },
    "v-card": { template: "<div><slot /></div>" },
    "v-card-text": { template: "<div><slot /></div>" },
    "v-card-actions": { template: "<div><slot /></div>" },
    "v-btn": { template: "<button><slot /></button>" },
  }

  const mountIndicator = (
    editingAsOverrides: Partial<
      ReturnType<typeof buildEditingAvailabilityAsViewModel>
    > = {},
    propsOverride: Record<string, unknown> = {},
  ) =>
    mount(EditingAvailabilityAs, {
      props: {
        editingAs: {
          ...buildEditingAvailabilityAsViewModel(),
          ...editingAsOverrides,
        },
        editGuestNameDialog: false,
        newGuestName: "",
        ...propsOverride,
      },
      global: {
        stubs: {
          ...scheduleOverlapGlobalStubs,
          ...dialogContentStubs,
        },
      },
    })

  it("renders the plain actor fallback when no editable guest is targeted", () => {
    const wrapper = mountIndicator({
      actionText: "Adding",
      actorName: "a guest",
    })

    expect(wrapper.text()).toContain("Adding availability as a guest")
    expect(wrapper.find(".editing-availability-as__guest").exists()).toBe(false)
  })

  it("renders the editable guest name with a pencil affordance that opens the dialog", async () => {
    const wrapper = mountIndicator({
      actionText: "Editing",
      editableGuestName: "Dana",
    })

    expect(wrapper.text()).toContain("Editing availability as")
    expect(wrapper.text()).toContain("Dana")

    await wrapper.get(".editing-availability-as__guest").trigger("click")

    expect(wrapper.emitted("openEditGuestNameDialog")).toHaveLength(1)
  })

  it("relays the guest name dialog cancel action", async () => {
    const wrapper = mountIndicator(
      { editableGuestName: "Dana" },
      { editGuestNameDialog: true, newGuestName: "Dee" },
    )

    expect(wrapper.find("v-text-field-stub").exists()).toBe(true)

    const cancelButton = wrapper
      .findAll("button")
      .find((node) => node.text() === "Cancel")
    if (!cancelButton) {
      throw new Error("Expected dialog Cancel button to be rendered")
    }
    await cancelButton.trigger("click")

    expect(wrapper.emitted("update:editGuestNameDialog")).toEqual([[false]])
  })

  it("renders the sentence variant by default", () => {
    const wrapper = mountIndicator({
      actionText: "Editing",
      editableGuestName: "Dana",
    })

    expect(wrapper.find(".editing-availability-as--chip").exists()).toBe(false)
    expect(wrapper.find(".editing-availability-as__guest").exists()).toBe(true)
  })

  it("renders the chip variant as a non-italic block with a rectangular chip button", () => {
    const wrapper = mountIndicator(
      { actionText: "Editing", editableGuestName: "John Doe" },
      { variant: "chip" },
    )

    const indicator = wrapper.get(".editing-availability-as--chip")
    expect(indicator.classes()).not.toContain("tw-justify-end")
    expect(indicator.classes()).toContain("tw-not-italic")
    expect(indicator.classes()).toContain("tw-flex-wrap")

    const row = wrapper.get(".editing-availability-as__chip-row")
    expect(row.classes()).not.toContain("tw-justify-end")
    expect(row.classes()).toContain("tw-flex-wrap")
    expect(row.text()).toContain("Editing availability as")

    const chip = wrapper.get(".editing-availability-as__guest-chip")
    expect(chip.element.tagName).toBe("BUTTON")
    expect(chip.classes()).toContain("tw-grow")
    expect(chip.text()).toContain("John Doe")
    expect(chip.find("v-icon-stub").exists()).toBe(true)
    expect(chip.classes()).toContain("tw-rounded")
    expect(chip.classes()).not.toContain("tw-rounded-full")
    expect(chip.classes()).toContain("tw-text-left")
    const name = wrapper.get(".editing-availability-as__guest-name")
    expect(name.classes()).not.toContain("tw-underline")
    expect(name.classes()).toContain("tw-grow")
  })

  it("emits the open-dialog event when the name chip is clicked", async () => {
    const wrapper = mountIndicator(
      { actionText: "Editing", editableGuestName: "Dana" },
      { variant: "chip" },
    )

    await wrapper.get(".editing-availability-as__guest-chip").trigger("click")

    expect(wrapper.emitted("openEditGuestNameDialog")).toHaveLength(1)
  })

  it("relays the guest name dialog actions unchanged in the chip variant", async () => {
    const wrapper = mountIndicator(
      { actionText: "Editing", editableGuestName: "Dana" },
      { variant: "chip", editGuestNameDialog: true, newGuestName: "Dee" },
    )

    const saveButton = wrapper
      .findAll("button")
      .find((node) => node.text() === "Save")
    if (!saveButton) {
      throw new Error("Expected dialog Save button to be rendered")
    }
    await saveButton.trigger("click")

    expect(wrapper.emitted("saveGuestName")).toHaveLength(1)

    const cancelButton = wrapper
      .findAll("button")
      .find((node) => node.text() === "Cancel")
    if (!cancelButton) {
      throw new Error("Expected dialog Cancel button to be rendered")
    }
    await cancelButton.trigger("click")

    expect(wrapper.emitted("update:editGuestNameDialog")).toEqual([[false]])
  })

  it("lets the chip drop below the label and the name break within the chip", () => {
    const wrapper = mountIndicator(
      {
        actionText: "Editing",
        editableGuestName: "Bartholomew Montgomery Fitzwilliam",
      },
      { variant: "chip" },
    )

    expect(
      wrapper.get(".editing-availability-as__chip-row").classes(),
    ).toContain("tw-flex-wrap")
    expect(
      wrapper.get(".editing-availability-as__guest-chip").classes(),
    ).toContain("tw-grow")
    expect(
      wrapper.get(".editing-availability-as__guest-chip").classes(),
    ).toContain("tw-min-w-0")
    const name = wrapper.get(".editing-availability-as__guest-name")
    expect(name.classes()).toContain("tw-break-words")
    expect(name.classes()).toContain("tw-grow")
    expect(name.text()).toContain("Bartholomew Montgomery Fitzwilliam")
  })

  it("shows the respondent-name placeholder in the chip when the editable name is empty", () => {
    const wrapper = mountIndicator(
      { actionText: "Editing", editableGuestName: "" },
      { variant: "chip" },
    )

    expect(
      wrapper.get(".editing-availability-as__guest-chip").text(),
    ).toContain("Respondent name")
  })

  it("renders non-editable actors as plain text in the chip variant", () => {
    const wrapper = mountIndicator(
      { actionText: "Adding", actorName: "a guest", editableGuestName: null },
      { variant: "chip" },
    )

    const indicator = wrapper.get(".editing-availability-as--chip")
    expect(indicator.text()).toContain("Adding availability as a guest")
    expect(wrapper.find(".editing-availability-as__guest-chip").exists()).toBe(
      false,
    )
    expect(wrapper.find(".editing-availability-as__guest").exists()).toBe(false)
    expect(indicator.findAll("v-icon-stub")).toHaveLength(0)
  })
})
