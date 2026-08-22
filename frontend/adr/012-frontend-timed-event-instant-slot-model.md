# ADR-012: Frontend Timed Event Picked-Date Slot Model

Date: 2026-05-31

Status:

- Accepted

## Context

Timed events currently use two competing frontend models:

- ordinary timed events use selected dates plus one start or end window per day
- specific-times events use explicit slot instants

That split creates inconsistent behavior across creation, edit hydration, grid rendering, date summaries, and timezone changes. It also forces the frontend to switch between a continuous-window model and a slot-based model for the same class of timed scheduling behavior.

The product intent is to use one canonical slot model for all non-`daysOnly` timed events. Under that model:

- timed event availability is represented as slots
- timed specific-date membership is represented as picked dates
- range events and specific-times events are mutually exclusive timed modes
- the mode determines how newly picked dates initialize active slots

Within the frontend stack migration described in ADR-008, this is a scheduling-domain modeling problem rather than a local rendering bug. The frontend needs one explicit canonical model for all timed event semantics.

## Terminology

- `picked dates`: the canonical picked-date membership for timed specific-date events
- `enabled slots`: the full slot domain for the picked dates; always the full civil day (`00:00` through the next `00:00` exclusive, in the event timezone) of each picked date (specific-dates) or recurrence anchor-week instance day (weekly); a derived value, recomputed on the fly from picked dates and the event timezone — independent of the slot-generation window and never persisted
- `active slots`: the only persisted slot set; the canonical subset of enabled slots that respondents can actually answer on
- `event timezone`: the persisted timezone used to interpret picked dates and generate enabled slots
- `display timezone`: the timezone used to render timed slots in the UI; this is controlled by the `Shown in` selector and does not change canonical event state
- `slot-generation settings`: the persisted timed-event settings whose start/end window generates the initial active range on creation; the stored window also serves as active-range metadata. The enabled domain is always the full civil day and is independent of this window.
- `range event`: a timed event whose active slots are generated from its within-date start/end range for every picked date
- `specific-times event`: a timed event whose active slots are selected individually

This ADR standardizes on `active slots` and does not use `respondable slots` as a primary term.
The canonical invariant is `active slots ⊆ enabled slots`.

## Decision

For all non-`daysOnly` timed events, the frontend treats timed-event state as a picked-date slot model.

- enabled slots are a derived domain: the full civil day of each picked date (specific-dates) or anchor-week instance day (weekly) in the event timezone, independent of the slot-generation window
- active slots are the only persisted slot set
- the event timezone is persisted explicitly and used for projection and batch-edit operations
- slot-generation settings are persisted explicitly; the window generates the initial active range on creation and is preserved as active-range metadata. It does not bound the enabled domain.
- timed specific-date picked dates are persisted explicitly and used for date-picker selections and regeneration of the enabled slot domain
- display timezone only affects projection/rendering and does not change picked dates

For timed specific-date events, the date picker is a direct view of picked dates. Picking a date adds that date to membership and regenerates the full civil-day enabled slot domain for that date.

Range events and specific-times events are mutually exclusive modes over the same picked-date and enabled-slot model. A range event persists the slots generated from its within-date range; a specific-times event persists the slots selected individually. Any active instant outside the enabled full-day domain is dropped on save (the wipe rule); the sole exception is the timezone-change re-anchor rule below.

`daysOnly` events remain outside this model and continue using date-only semantics.

## Rules

### Canonical timed-event state

- Do not treat `event.dates`, `timeSeed`, or `duration` as the canonical source of truth for timed-event rendering or editing.
- Do not let compatibility dates, seed datetimes, duration windows, and slot instants compete as parallel authorities for the same timed-event UI.
- Treat `timedRecurrence.selectedDays` as the canonical picked-date membership for timed specific-date events.
- Keep active slots as the only persisted slot set for canonical timed-event state.
- Derive the enabled slot domain on the fly: a pure function of picked dates (specific-dates) or the recurrence anchor week (weekly), the event timezone, and the slot-generation settings. Never persist or transport it.
- For weekly events the anchor week is recovered from the earliest active slot instant; an event with no active instants falls back to `timeSeed` (frontend) or `time.Now()` in the event timezone (server), which may not match what the organizer last saw.
- Keep canonical timed-event state as picked dates plus instant-based active slots, not mixed civil-date-plus-window reconstructions.
- Treat range and specific-times as mutually exclusive modes. Keep enough persisted state to restore the selected mode and its editing behavior.

### Timezone semantics

- Persist an explicit event timezone for timed events as the authoring and projection timezone.
- Use IANA timezone identity as the long-term persisted timezone model for timed events.
- Treat fixed-offset-only timezone data as compatibility input when needed, not as the preferred long-term event-timezone model when a richer timezone identity is available.
- Keep picked dates stable as civil dates when the event timezone changes, except when a change would drop cross-midnight active instants: then picked dates re-anchor to the actives' local dates in the new event timezone so the instants survive (the re-anchor rule).
- Rebuild enabled slots from picked dates in the new event timezone (always the full civil day; the slot-generation window only supplies the increment and stored metadata).
- Filter active slots to the rebuilt enabled domain; an instant outside the rebuilt domain after a plain timezone change is dropped by the wipe rule.
- A display-timezone change only changes projection in the UI. It does not change picked dates, enabled slots, or active slots.

### Slot-generation semantics

- Timed events must keep explicit slot-generation settings sufficient to regenerate the initial active range for picked dates or timed recurrence-owned instances without inferring from viewer-local rendering.
- Those settings must be interpreted in the event timezone.
- Batch-added enabled slots for a picked date are the full civil day of that date in the event timezone; the slot-generation window is not a bound on the enabled domain.
- A default timed event created from selected local days plus start/end time must generate the full slot set for the configured window and persist that generated set as the active slots, keeping the window as the stored `slotGeneration`. The enabled domain for those dates is the full civil day and derives automatically.

### Timed-event mode semantics

- Range events and specific-times events are mutually exclusive.
- A range event generates active slots for each picked date at every configured increment whose start is at or after the range start and before the range end, interpreted in the event timezone. The range must not cross midnight in the event timezone.
- A specific-times event exposes slot-level edits over the canonical enabled-slot and active-slot state. Its enabled domain is the full civil day of its picked dates.
- Converting a range event to a specific-times event preserves the range-generated active slots as the initial manually editable selection.
- Converting a specific-times event to a range event replaces its manually selected active slots with the slots generated by the creator-selected within-date range for every picked date, interpreted in the event timezone.
- On save, active instants outside the enabled full-day domain are dropped (the wipe rule); the re-anchor rule is the only exception.

### Picked-date semantics

- For timed specific-date events, the date picker displays picked dates directly.
- Picking a date in a range event adds its full civil-day enabled domain and generates that date's range-defined active slots.
- Picking a date in a specific-times event adds its full civil-day enabled domain only; it does not initialize active slots. Added dates are enabled-only; existing active subsets are preserved unchanged.
- Unpicking a date removes that picked date and also removes any enabled or active slots on that date so the canonical subset invariant remains valid.

### DOW and group timed-event semantics

- Weekly and group timed events also use the timed-slot model.
- Only their active slots are persisted as instants; their enabled domain is derived from the anchor week (earliest active instant) in the event timezone just like specific-date timed events.
- Any recurrence-style creation or edit flow must use slot-generation settings to generate the relevant timed instances instead of relying on a separate canonical seed-plus-duration runtime model.

### Rendering and summary semantics

- Timed grids must derive rendered date columns from enabled slots projected into the display timezone. A projected slot belongs to its display-local calendar-date column; an adjacent column must be created when that date is otherwise absent.
- Specific-times grids additionally retain picked-date columns that have no projected enabled slots, so their full editable domain remains visible.
- Timed grids must derive rendered slot existence from enabled slots after projection into the display timezone; they must not render synthetic out-of-domain slots.
- Timed grids must derive participant-selectable cells from active slots.
- The read-only event-page grid collapses the axis to the saved active-range band: the band derives from active slots (with the enabled domain as fallback when there are no actives), so ranges without actives are not displayed. The full civil-day axis is shown only in the specific-times editor or when `Show all hours` is enabled.
- Specific-times editing may render a full-height rectangular grid around projected slots. Cells without a mapped enabled slot are disabled padding cells: they are non-editable and use a visually distinct unavailable treatment. Enabled inactive cells remain editable and use a separate treatment.
- When the grid already shows picked dates, duplicate date-summary UI should be hidden.
- Projection from instants to rendered local slots must stay centralized at scheduling boundaries instead of being rebuilt ad hoc in views or route-local helpers.

### Compatibility semantics

- Legacy timed payloads that only carry active timed instants decode with `enabled slots = active slots` until richer contract data (recurrence/timezone/slot-generation settings) is available.
- Legacy `event.dates`, `timeSeed`, and duration-style timed data are compatibility inputs only and must not override canonical picked dates, active slots, or the derived enabled domain once present.

## Consequences

- Timed events require explicit persisted timezone semantics instead of relying on timezone identity flattened into stored instants.
- Frontend transport and internal models persist only active slots; the enabled domain is derived by a centralized helper and, server-side, regenerated and validated against the persisted contract.
- Existing ordinary timed-event `dates + duration` behavior is retired as canonical state.
- Existing `dates` and `timeSeed` behavior for timed events should be treated as compatibility-only during migration, not as the long-term canonical model.
- Timed-event mode is a persistence and editing boundary. Conversion preserves or replaces active slots according to the selected direction.
- Edit hydration, date summaries, and grid rendering for all timed event types will need to stop reading competing date authorities for the same timed event.
- Timed events will keep explicit slot-generation settings: the window generates the initial active range and serves as stored active-range metadata; the enabled domain for membership days is always the full civil day.
- Regression tests should lock instant-preserving timezone changes (including the re-anchor rule), the full-day enabled domain, the wipe rule, and enabled-slot versus active-slot behavior before broader refactors proceed.

## Required Acceptance Scenarios

- Creating a normal timed event with selected days and a start/end window persists the full-window slot set as active slots and the window as `slotGeneration`; the enabled domain derives to the full civil day of the picked dates.
- Creating a specific-times event keeps a full-day enabled domain for every picked date and initializes an empty active subset, then enables slot-level selection.
- Adding a date to a range event generates active slots at each configured increment from the inclusive start through the exclusive end; adding a date to a specific-times event adds no active slots.
- Converting a range event to specific-times preserves range-generated active slots as the initial selection; converting specific-times to a range replaces manually selected active slots with the chosen within-date range for each picked date.
- A no-op specific-times save preserves the stored `slotGeneration` window, picked dates, timezone, and timed recurrence.
- An active instant outside the full-day enabled domain is dropped on save; e.g. the next-day `00:00`/`00:30` UTC instants of a cross-midnight window on a picked `2026-01-05` UTC event are wiped.
- Changing only the display timezone can shift active cells into different projected date columns without changing canonical picked-date membership.
- Adding a date in a range event generates its range-defined active slots; adding a date in a specific-times event adds only its full civil-day enabled slot domain.
- Removing a date removes the derived enabled slots and any active slots on that picked date.
- Changing the event timezone re-anchors picked dates to the actives' local dates in the new event timezone when the change would otherwise drop cross-midnight instants; plain timezone changes keep picked dates stable and filter the persisted active slots to the rebuilt domain.
- Cross-midnight enabled domains render slots in their display-local date columns after display-timezone changes, creating adjacent columns as needed.
- The read-only event-page grid collapses to the saved active-range band (derived from active slots, falling back to the enabled domain); the full civil-day axis appears only in the specific-times editor or with `Show all hours`.
- Specific-times editing distinguishes selected active slots, editable enabled inactive slots, and non-editable disabled padding cells in its legend and grid treatment.
- DST gap and overlap days derive the correct full civil-day enabled slots in the event timezone.
  - With the full civil-day domain the window-start-in-fold divergence is retired: derivation starts and ends at local midnight, which never falls inside an ambiguity fold, so the frontend Temporal and the Go server derive identical domains (locked by the shared `TestDeriveEnabledSlotsLordHoweThirtyMinuteDtstMatchesFrontendFixture` fixtures).
- DOW and group timed events derive displayed days and timed grids from the anchor-week enabled domain the same way as timed specific-date events.
- Legacy timed payloads without the full contract decode with `enabled slots = active slots` until migrated data is available.
- The main event header, timed grids, and edit dialog all derive from the same canonical timed-event state.
- `daysOnly` events remain outside this model and continue using date-only semantics.
