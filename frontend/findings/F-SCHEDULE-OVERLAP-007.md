# F-SCHEDULE-OVERLAP-007

- Status: `fixed`
- Priority: `P1`
- Component: `src/composables/schedule_overlap/useDragPaint.ts`, `src/components/schedule_overlap/ScheduleOverlap.vue`
- Problem: Schedule-event drags could end in enabled but inactive specific-time cells, and timed blocks retained their full height when their range crossed a collapsed-hours row.
- Root cause: Drag movement and pointer-up did not validate endpoints against active slots. Timed blocks only translated their top position to visible rows, leaving their base-grid duration height unchanged.
- Resolution: Schedule drags retain their last active endpoint. Timed calendar and scheduled-event blocks now render visible contiguous fragments, skipping collapsed rows.
- Regression coverage: `useDragPaint.test.ts` covers inactive endpoint rejection; `scheduleOverlapRendering.test.ts` covers block fragmentation around collapsed rows.
