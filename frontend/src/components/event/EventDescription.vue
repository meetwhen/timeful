<template>
  <div class="tw-mt-1 tw-max-w-full sm:tw-mt-2">
    <div
      v-if="showDescription"
      class="event-description-shell tw-relative tw-w-full tw-cursor-pointer tw-rounded-md tw-border tw-border-light-gray-stroke tw-bg-light-gray tw-p-2 tw-font-normal tw-text-very-dark-gray hover:tw-bg-[#eeeeee]"
    >
      <div
        class="event-description-copy event-description-text tw-whitespace-pre-wrap tw-break-words"
        :class="canEdit ? 'tw-pr-[72px]' : ''"
      >
        {{ event.description }}
      </div>
      <div
        v-if="canEdit"
        class="event-description-actions tw-absolute tw-right-2 tw-top-1"
      >
        <v-btn
          key="edit-description-btn"
          class="event-description-action-button event-description-edit-button tw-h-8 tw-w-8"
          icon
          variant="text"
          size="small"
          @click="beginEditing"
        >
          <v-icon size="24">mdi-pencil</v-icon>
        </v-btn>
      </div>
    </div>

    <button
      v-else-if="canEdit && !isEditing"
      type="button"
      class="event-description-add-trigger event-description-text tw-mt-0 tw-px-2 tw-py-2 tw-text-dark-gray"
      @click="beginEditing"
    >
      + Add description
    </button>
    <div
      v-else-if="canEdit && isEditing"
      class="event-description-edit-shell tw-relative tw-w-full tw-px-2 tw-py-2 tw-font-normal tw-text-very-dark-gray"
    >
      <div class="event-description-editor tw-pr-[72px]">
        <div
          ref="descriptionEditor"
          class="event-description-copy event-description-text event-description-editor-field tw-border-0 tw-border-b tw-border-solid tw-bg-transparent tw-outline-none"
          :contenteditable="isEditing ? 'true' : 'false'"
          :data-placeholder="descriptionPlaceholder"
          role="textbox"
          aria-multiline="true"
          @input="syncDraftDescription"
        ></div>
      </div>
      <div
        class="event-description-actions tw-absolute tw-right-2 tw-top-1 tw-flex tw-gap-2"
      >
        <v-btn
          class="event-description-action-button event-description-cancel-button tw-h-8 tw-w-8"
          icon
          variant="text"
          size="small"
          @click="cancelEditing"
        >
          <v-icon size="24">mdi-close</v-icon>
        </v-btn>
        <v-btn
          class="event-description-action-button event-description-save-button tw-h-8 tw-w-8"
          icon
          variant="text"
          size="small"
          color="primary"
          @click="saveDescription"
        >
          <v-icon size="24">mdi-check</v-icon>
        </v-btn>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, nextTick, ref } from "vue"
import { put } from "@/utils"
import { useMainStore } from "@/stores/main"
import type { Event } from "@/types"
import { toEventPatchPayload } from "@/composables/event/eventMutationBoundary"

const props = defineProps<{
  event: Event
  canEdit: boolean
}>()

const emit = defineEmits<{
  "update:event": [event: Event]
}>()

const mainStore = useMainStore()

const isEditing = ref(false)
const draftDescription = ref("")
const descriptionEditor = ref<HTMLDivElement | null>(null)
const descriptionPlaceholder = "Enter a description..."

const showDescription = computed(
  () => props.event.description && !isEditing.value
)

const syncEditorContent = (value: string) => {
  const editor = descriptionEditor.value

  if (!editor) {
    return
  }

  if (editor.textContent !== value) {
    editor.textContent = value
  }
}

const contenteditableText = (element: HTMLElement): string => {
  const isBlock = (node: Node) =>
    node.nodeType === Node.ELEMENT_NODE &&
    ["DIV", "LI", "P"].includes((node as HTMLElement).tagName)

  const textFromChildren = (parent: Node) => {
    let text = ""
    let previousNode: Node | undefined

    for (const node of parent.childNodes) {
      if (
        previousNode &&
        (isBlock(previousNode) || isBlock(node)) &&
        !text.endsWith("\n")
      ) {
        text += "\n"
      }

      if (node.nodeType === Node.TEXT_NODE) {
        text += node.textContent ?? ""
      } else if (
        node.nodeType === Node.ELEMENT_NODE &&
        (node as HTMLElement).tagName === "BR"
      ) {
        text += "\n"
      } else {
        text += textFromChildren(node)
      }

      previousNode = node
    }

    return text
  }

  return textFromChildren(element).replace(/\r\n?/g, "\n")
}

const syncDraftDescription = () => {
  draftDescription.value = descriptionEditor.value
    ? contenteditableText(descriptionEditor.value)
    : ""
}

const beginEditing = async () => {
  draftDescription.value = props.event.description ?? ""
  isEditing.value = true

  await nextTick()
  syncEditorContent(draftDescription.value)
  descriptionEditor.value?.focus()
}

const saveDescription = () => {
  const oldEvent = { ...props.event }
  const newEvent = { ...props.event, description: draftDescription.value }

  const eventPayload = toEventPatchPayload({
    name: props.event.name,
    duration: props.event.duration,
    type: props.event.type,
    description: draftDescription.value,
    dates: props.event.dates,
    timeSeed: props.event.timeSeed,
    activeSlots: props.event.activeSlots,
    eventTimezone: props.event.eventTimezone,
    slotGeneration: props.event.slotGeneration,
    timedRecurrence: props.event.timedRecurrence,
  })

  emit("update:event", newEvent)
  isEditing.value = false
  put(`/events/${props.event._id ?? ""}`, eventPayload).catch((err: unknown) => {
    console.error(err)
    mainStore.showError(
      "Failed to save description! Please try again later."
    )
    emit("update:event", { ...oldEvent })
  })
}

const cancelEditing = () => {
  isEditing.value = false
}
</script>

<style scoped>
.event-description-text {
  font-size: 0.75rem;
  line-height: 1.5rem;
}

@media (min-width: 640px) {
  .event-description-text {
    font-size: 0.875rem;
    line-height: 1.5rem;
  }
}

.event-description-copy {
  min-height: 1.5rem;
}

.event-description-add-trigger {
  background: none;
  border: 0;
  color: inherit;
  cursor: pointer;
  display: block;
  font-family: inherit;
  font-weight: inherit;
  margin: 0;
  text-align: left;
}

.event-description-add-trigger:focus-visible {
  outline: 2px solid rgb(0 153 76 / 1);
  outline-offset: 2px;
}

.event-description-editor-field {
  border-bottom-color: var(--timeful-grid-hour-separator);
  color: inherit;
  font-family: inherit;
  font-weight: inherit;
  letter-spacing: inherit;
  white-space: pre-wrap;
  word-break: break-word;
}

.event-description-editor-field:empty::before {
  content: attr(data-placeholder);
  color: var(--timeful-disabled-foreground);
}
</style>
