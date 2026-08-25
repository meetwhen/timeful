<template>
  <div v-if="event.description?.trim()" class="tw-mt-1 tw-max-w-full sm:tw-mt-2">
    <div class="event-description-shell tw-w-full tw-rounded-md tw-border tw-border-light-gray-stroke tw-bg-light-gray tw-p-2 tw-font-normal tw-text-very-dark-gray">
      <div class="event-description-copy event-description-text tw-whitespace-pre-wrap tw-break-words">
        {{ event.description }}
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import type { Event } from "@/types"

defineProps<{
  event: Event
}>()
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

</style>
