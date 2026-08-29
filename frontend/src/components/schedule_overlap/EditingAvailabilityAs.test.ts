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
})
