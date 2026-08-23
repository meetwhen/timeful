# Timeful Glossary

This glossary records controlled terms used in Timeful documentation. Its
definitions are concise references; the linked authoritative context defines
the complete behavior and wins if the two conflict.

## Text Terms

### Unicode Normalization Form C (NFC)

A Unicode normalization form that applies canonical composition, producing a
consistent composed code-point sequence for canonically equivalent text.

Authoritative context: [Unicode Standard Annex #15](https://www.unicode.org/reports/tr15/).

## Timed-slot Terms

### Event Kind

The event's top-level scheduling model. A **Timed Event** collects availability
for time slots; a **Dates-Only Event** collects availability for calendar dates.

Authoritative context: [ADR-012 decision](../../frontend/adr/012-frontend-timed-event-instant-slot-model.md#decision).

### Timed Event

An event kind whose availability domain consists of time slots. A timed event
has a timed domain mode and may have a scheduled event time range.

Authoritative context: [ADR-012 decision](../../frontend/adr/012-frontend-timed-event-instant-slot-model.md#decision).

### Dates-Only Event

An event kind whose availability domain consists of calendar dates. A
dates-only event may have a scheduled event date.

Authoritative context: [ADR-012 decision](../../frontend/adr/012-frontend-timed-event-instant-slot-model.md#decision).

### Timed Domain Mode

The persisted configuration that determines how a timed event's active-slot
domain is maintained. **Ranged Domain Mode** configures it from an active-slot
range; **Custom Domain Mode** configures it through custom domain editing.

Authoritative context: [ADR-012 terminology](../../frontend/adr/012-frontend-timed-event-instant-slot-model.md#terminology) and [advanced slot editing semantics](../../frontend/adr/012-frontend-timed-event-instant-slot-model.md#advanced-slot-editing-semantics).

### Ranged Domain Mode

A timed domain mode that creates or restores the active-slot domain from the
active-slot range for the current picked-date domain.

Authoritative context: [ADR-012 terminology](../../frontend/adr/012-frontend-timed-event-instant-slot-model.md#terminology) and [slot-generation semantics](../../frontend/adr/012-frontend-timed-event-instant-slot-model.md#slot-generation-semantics).

### Custom Domain Mode

A timed domain mode that preserves a custom subset of enabled slots and exposes
custom domain editing. Switching to ranged domain mode restores the active-slot
domain from its active-slot range.

Authoritative context: [ADR-012 terminology](../../frontend/adr/012-frontend-timed-event-instant-slot-model.md#terminology) and [advanced slot editing semantics](../../frontend/adr/012-frontend-timed-event-instant-slot-model.md#advanced-slot-editing-semantics).

### Custom Domain Editing

The slot-level editing UI available only in custom domain mode. It adds or
removes active slots without changing picked dates or the canonical persistence
model.

Authoritative context: [ADR-012 terminology](../../frontend/adr/012-frontend-timed-event-instant-slot-model.md#terminology) and [advanced slot editing semantics](../../frontend/adr/012-frontend-timed-event-instant-slot-model.md#advanced-slot-editing-semantics).

### Picked Dates

The canonical selected-date membership for an event. For timed events, the date
picker is a direct view of picked dates and picking a date regenerates the
enabled-slot domain for that date.

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

### Active Slot Settings

The persisted timed-event configuration that determines active slots. It
contains the timed domain mode and, depending on that mode, either an active
slot range or the selected active slots.

Authoritative context: [ADR-012 terminology](../../frontend/adr/012-frontend-timed-event-instant-slot-model.md#terminology) and [slot-generation semantics](../../frontend/adr/012-frontend-timed-event-instant-slot-model.md#slot-generation-semantics).

### Active Slot Range

The start/end window in active slot settings. In ranged domain mode it defines
the active slots generated for the picked-date domain.

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

## Response Access Terms

### Event Owner

The person authorized to manage an event's event settings. An anonymous event
owner proves browser-local ownership and may associate that ownership with an
authenticated account through event sign-in.

Authoritative context: [FR-018](../requirements/functional/fr/FR-018.md) and [FR-063](../requirements/functional/fr/FR-063.md).

### Event Guest

The browser-local credential representing a person responding to an event. An
event guest can own multiple availability responses and can be associated with
an authenticated account through event sign-in.

Authoritative context: [ADR-011](../../frontend/adr/011-frontend-guest-response-ownership-semantics.md) and [FR-062](../requirements/functional/fr/FR-062.md).

### Availability Response

An event guest's recorded availability for an event, including its access mode
and availability states.

Authoritative context: [ADR-011](../../frontend/adr/011-frontend-guest-response-ownership-semantics.md), [FR-060](../requirements/functional/fr/FR-060.md), and [FR-061](../requirements/functional/fr/FR-061.md).

### Availability State

The state an availability response assigns to a slot or date. **Available** and
**If needed** count equally in overlap calculations; **Unavailable** does not.

Authoritative context: [FR-006](../requirements/functional/fr/FR-006.md) and [FR-069](../requirements/functional/fr/FR-069.md).

### Availability Overlay

The availability response being edited, rendered above other responses so its
availability states remain the editable layer.

Authoritative context: [FR-005](../requirements/functional/fr/FR-005.md).

### Protected Response

The default availability-response access mode. Only the event guest that owns
the response may edit it through that guest's browser-local credential or its
associated authenticated account.

Authoritative context: [ADR-011](../../frontend/adr/011-frontend-guest-response-ownership-semantics.md) and [FR-060](../requirements/functional/fr/FR-060.md).

### Open Response

An availability response whose owning event guest has explicitly allowed any
event visitor to edit it.

Authoritative context: [ADR-011](../../frontend/adr/011-frontend-guest-response-ownership-semantics.md) and [FR-061](../requirements/functional/fr/FR-061.md).

## Timed-Grid Rendering Terms

### Enabled Inactive Slot

An enabled slot that is not active. It remains editable in custom domain
editing but is not respondent-selectable; it has a treatment separate from
both active slots and disabled padding cells.

Authoritative context: [ADR-012 rendering rules](../../frontend/adr/012-frontend-timed-event-instant-slot-model.md#rendering-and-summary-semantics) and [FR-026](../requirements/functional/fr/FR-026.md).

### Disabled Padding Cell

A grid cell without a mapped enabled slot. It is non-editable and uses a
visually distinct unavailable treatment; it is not a slot.

Authoritative context: [ADR-012 rendering rules](../../frontend/adr/012-frontend-timed-event-instant-slot-model.md#rendering-and-summary-semantics) and [FR-026](../requirements/functional/fr/FR-026.md).

### Projected Date Column

A grid column derived from enabled slots projected into the display timezone. A
projected slot belongs to its display-local calendar-date column; an adjacent
column is created when that date is otherwise absent, and a slot crossing
midnight is not duplicated.

Authoritative context: [ADR-012 rendering rules](../../frontend/adr/012-frontend-timed-event-instant-slot-model.md#rendering-and-summary-semantics), [FR-025](../requirements/functional/fr/FR-025.md), and [FR-026](../requirements/functional/fr/FR-026.md).

### Saved Active-Range Band

The read-only event-page grid's collapsed axis, derived from active slots and
falling back to the enabled domain when there are no active slots. The full
civil-day axis appears only in custom domain editing or with `Show all
hours`.

Authoritative context: [ADR-012 rendering rules](../../frontend/adr/012-frontend-timed-event-instant-slot-model.md#rendering-and-summary-semantics) and [FR-026](../requirements/functional/fr/FR-026.md).

### Scheduled Event Time

The optional scheduled occurrence for an event. It can be replaced or cleared;
it is a time range for a timed event and one date for a dates-only event.

Authoritative context: [FR-012](../requirements/functional/fr/FR-012.md).

### Event Settings

The settings that configure an event itself, rather than an individual
availability response. Their scope includes the description, event kind, date
selection, event and display timezones, event and display time formats, active
slot settings, and active slot range.

Authoritative context: [FR-018](../requirements/functional/fr/FR-018.md).
