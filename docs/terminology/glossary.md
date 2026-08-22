# Timeful Glossary

This glossary records controlled terms used in Timeful documentation. Its
definitions are concise references; the linked authoritative context defines
the complete behavior and wins if the two conflict.

## Timed-slot Terms

### Picked Dates

The canonical picked-date membership for timed specific-date events. The date
picker is a direct view of picked dates; picking a date regenerates the enabled
slot domain for that date.

Authoritative context: [ADR-012 terminology](../../frontend/adr/012-frontend-timed-event-instant-slot-model.md#terminology) and [picked-date semantics](../../frontend/adr/012-frontend-timed-event-instant-slot-model.md#picked-date-semantics).

### Enabled Slots

The full slot domain for the picked dates: the full civil day (`00:00` through
the next `00:00` exclusive, in the event timezone) of each picked date for
specific-date events or recurrence anchor-week instance day for weekly events.
Enabled slots are derived from picked dates and the event timezone, independent
of the slot-generation window, and are never persisted or transported.

Authoritative context: [ADR-012 terminology](../../frontend/adr/012-frontend-timed-event-instant-slot-model.md#terminology) and [canonical timed-event state](../../frontend/adr/012-frontend-timed-event-instant-slot-model.md#canonical-timed-event-state).

### Active Slots

The canonical subset of enabled slots that respondents can answer on. The
canonical invariant is `active slots subset of enabled slots`.

Authoritative context: [ADR-012 terminology](../../frontend/adr/012-frontend-timed-event-instant-slot-model.md#terminology).

### Event Timezone

The persisted timezone used to interpret picked dates and generate enabled
slots. Changing it preserves picked dates, rebuilds the enabled domain, and
filters active slots to the rebuilt domain.

Authoritative context: [ADR-012 terminology](../../frontend/adr/012-frontend-timed-event-instant-slot-model.md#terminology) and [timezone semantics](../../frontend/adr/012-frontend-timed-event-instant-slot-model.md#timezone-semantics).

### Display Timezone

The timezone used to render timed slots in the UI, controlled by the `Shown in`
selector. It affects projection and rendering only; it does not change picked
dates, enabled slots, or active slots.

Authoritative context: [ADR-012 terminology](../../frontend/adr/012-frontend-timed-event-instant-slot-model.md#terminology).

### Slot-Generation Settings

The persisted timed-event settings whose start/end window generates the initial
active range on creation; the stored window also serves as active-range
metadata. They are interpreted in the event timezone. The enabled domain is
always the full civil day and is independent of this window.

Authoritative context: [ADR-012 terminology](../../frontend/adr/012-frontend-timed-event-instant-slot-model.md#terminology) and [slot-generation semantics](../../frontend/adr/012-frontend-timed-event-instant-slot-model.md#slot-generation-semantics).

### Wipe Rule

On save, an active instant outside the enabled full-day domain is dropped. For
example, this includes next-day `00:00` and `00:30` instants from a
cross-midnight window on a picked UTC date. The timezone-change re-anchor rule
is the sole exception.

Authoritative context: [ADR-012 advanced slot editing semantics](../../frontend/adr/012-frontend-timed-event-instant-slot-model.md#advanced-slot-editing-semantics).

### Re-Anchor Rule

When an event-timezone change would drop cross-midnight active instants, picked
dates move to the actives' local dates in the new event timezone so those
instants survive. A plain timezone change keeps picked dates stable.

Authoritative context: [ADR-012 timezone semantics](../../frontend/adr/012-frontend-timed-event-instant-slot-model.md#timezone-semantics).

### Advanced Slot Editing

The UI mode exposed by the specific-times toggle. It allows slot-level edits to
active slots without changing the canonical persistence model. Disabling it
restores `active slots = enabled slots`: the full civil day of the current
picked-date domain.

Authoritative context: [ADR-012 terminology](../../frontend/adr/012-frontend-timed-event-instant-slot-model.md#terminology) and [advanced slot editing semantics](../../frontend/adr/012-frontend-timed-event-instant-slot-model.md#advanced-slot-editing-semantics).

### Days-Only Events

Events that remain outside the timed-slot model and use date-only semantics.

Authoritative context: [ADR-012 decision](../../frontend/adr/012-frontend-timed-event-instant-slot-model.md#decision).

## Time Format Terms

### Event Time Format

The 12/24-hour format used to render times inside event-editor forms, including
the `What times might work?` and `Time range` dropdowns and start/end time
fields. It persists independently of the display time format under
`eventTimeType`, defaults to 24-hour, and changing it during event creation
saves the new preference.

Authoritative context: [browser date preferences](../../frontend/src/utils/browserDatePreferences.ts) and [event editor state](../../frontend/src/composables/event/useEventEditorState.ts).

### Display Time Format

The format used to render times in the event-page schedule grid and tooltips,
controlled by the event-page `12h/24h` switch and persisted under `timeType`.
It defaults to 24-hour and affects event-page rendering only, not event-editor
forms.

Authoritative context: [browser date preferences](../../frontend/src/utils/browserDatePreferences.ts) and [calendar grid](../../frontend/src/composables/schedule_overlap/useCalendarGrid.ts).

## Timed-Grid Rendering Terms

### Enabled Inactive Slot

An enabled slot that is not active. It remains editable in specific-times
editing but is not respondent-selectable; it has a treatment separate from
both active slots and disabled padding cells.

Authoritative context: [ADR-012 rendering rules](../../frontend/adr/012-frontend-timed-event-instant-slot-model.md#rendering-and-summary-semantics) and [FR-026](../requirements/functional/FR-026.md).

### Disabled Padding Cell

A grid cell without a mapped enabled slot. It is non-editable and uses a
visually distinct unavailable treatment; it is not a slot.

Authoritative context: [ADR-012 rendering rules](../../frontend/adr/012-frontend-timed-event-instant-slot-model.md#rendering-and-summary-semantics) and [FR-026](../requirements/functional/FR-026.md).

### Projected Date Column

A grid column derived from enabled slots projected into the display timezone. A
projected slot belongs to its display-local calendar-date column; an adjacent
column is created when that date is otherwise absent, and a slot crossing
midnight is not duplicated.

Authoritative context: [ADR-012 rendering rules](../../frontend/adr/012-frontend-timed-event-instant-slot-model.md#rendering-and-summary-semantics), [FR-025](../requirements/functional/FR-025.md), and [FR-026](../requirements/functional/FR-026.md).

### Saved Active-Range Band

The read-only event-page grid's collapsed axis, derived from active slots and
falling back to the enabled domain when there are no active slots. The full
civil-day axis appears only in the specific-times editor or with `Show all
hours`.

Authoritative context: [ADR-012 rendering rules](../../frontend/adr/012-frontend-timed-event-instant-slot-model.md#rendering-and-summary-semantics) and [FR-026](../requirements/functional/FR-026.md).

### Range-Created Event

A timed event created from selected days and a start/end window. The generated
range is persisted as active slots and the window is stored as `slotGeneration`.
The enabled domain is always the full civil day of the picked dates, so there is
no separate range-generated versus full-day enabled-domain distinction.

Authoritative context: [ADR-012 decision](../../frontend/adr/012-frontend-timed-event-instant-slot-model.md#decision) and [FR-026](../requirements/functional/FR-026.md).
