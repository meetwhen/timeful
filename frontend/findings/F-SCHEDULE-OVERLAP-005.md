# F-SCHEDULE-OVERLAP-005

- Status: `fixed`
- Priority: `P1`
- Component: `src/composables/schedule_overlap/useAvailabilityData.ts`
- Problem: Timed event heatmap aggregation rebuilt response buckets from `day.dateObject.add(time.hoursOffset)` instead of the rendered slot projection. For timezone-shifted timed grids that use `absoluteMinutes` or other display-domain normalization, the fetched response slots were matched against the wrong instants, so the event page could show zero responded slots even when the response edit flow showed saved availability.
- Why it matters: Responded availability disappeared from the main event page for shifted timed events, which broke best-times and heatmap trust for owners and guests.
- Acceptance criteria: When a timed grid is rendered through a shifted local display domain, `responsesFormatted` should bucket fetched responses against the same slot instants produced by `getDateFromRowCol(...)`, so event-page heatmap cells and best-times counts reflect the saved responses.
- Root cause: `getResponsesFormatted()` reconstructed timed slots from raw `hoursOffset` values. That bypassed `useCalendarGrid`’s rendered-slot normalization, including `absoluteMinutes` handling for timezone-shifted views, so the response buckets landed on midnight-based instants instead of the visible 10:30-style slots.
- Fix: Build timed response buckets by iterating visible row and column indexes and calling `getDateFromRowCol(row, col)` for each rendered slot instead of recomputing instants from `hoursOffset`.
- Verification evidence: Added regression `"maps fetched timed responses onto rendered slots that use absolute local minutes"` in `src/utils/scheduleOverlap.regressions.test.ts`. Verified live against `http://127.0.0.1:4173/e/5B5af`: the page now shows four responded cells and no longer renders the “There’s no time when all 1 respondents are available.” note.
