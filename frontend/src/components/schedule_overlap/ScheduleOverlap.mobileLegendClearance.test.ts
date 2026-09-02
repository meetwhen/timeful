// @vitest-environment happy-dom

import { beforeEach, describe, expect, it } from "vitest"
import { defineComponent, nextTick } from "vue"
import {
  resetScheduleOverlapMocks,
  viewportWidth,
} from "./scheduleOverlapTestMocks"
import {
  installScheduleOverlapTestGlobals,
  mountScheduleOverlap,
} from "./scheduleOverlapTestUtils"

const MobileOverlayHeightStub = defineComponent({
  name: "ScheduleOverlapMobileOverlay",
  emits: ["overlayHeightChange"],
  mounted() {
    this.$emit("overlayHeightChange", 82)
  },
  template: "<div />",
})

const clearanceSpacer = (wrapper: ReturnType<typeof mountScheduleOverlap>) =>
  wrapper.find(".schedule-overlap__mobile-editing-clearance")

describe("ScheduleOverlap mobile legend clearance", () => {
  beforeEach(() => {
    resetScheduleOverlapMocks()
    installScheduleOverlapTestGlobals()
  })

  it("reserves space for the fixed editing panel stack while editing on phone", async () => {
    viewportWidth.value = 375
    const wrapper = mountScheduleOverlap({
      global: {
        stubs: {
          ScheduleOverlapMobileOverlay: MobileOverlayHeightStub,
        },
      },
    })

    expect(clearanceSpacer(wrapper).exists()).toBe(false)

    const vm = wrapper.vm as unknown as { startEditing: () => void }
    vm.startEditing()
    await nextTick()

    const spacer = clearanceSpacer(wrapper)
    expect(spacer.exists()).toBe(true)
    expect(spacer.attributes("style")).toContain("calc(82px + 4rem - 64px)")

    wrapper.unmount()
  })

  it("does not reserve editing panel space on desktop", async () => {
    viewportWidth.value = 1024
    const wrapper = mountScheduleOverlap({
      global: {
        stubs: {
          ScheduleOverlapMobileOverlay: MobileOverlayHeightStub,
        },
      },
    })

    const vm = wrapper.vm as unknown as { startEditing: () => void }
    vm.startEditing()
    await nextTick()

    expect(clearanceSpacer(wrapper).exists()).toBe(false)

    wrapper.unmount()
  })
})
