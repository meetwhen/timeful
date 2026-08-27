# Timeful Glossary

This glossary records controlled terms used in Timeful documentation.
Its definitions are concise references; the linked authoritative context defines the complete behavior and wins if the two conflict.

## Text Terms

### Unicode Normalization Form C (NFC)

A Unicode normalization form that applies canonical composition, producing a consistent composed code-point sequence for canonically equivalent text.

Authoritative context: [Unicode Standard Annex #15](https://www.unicode.org/reports/tr15/).

## Timed-slot Terms

### Event Kind

The event's top-level scheduling model.
A [Timed Event](#timed-event) collects availability for time slots; a [Dates-Only Event](#dates-only-event) collects availability for calendar dates.
Their UI labels are `Dates and times` and `Dates only`, respectively.

Authoritative context: [FR-026](../requirements/functional/fr/FR-026.md).

### Timed Event

An [Event Kind](#event-kind) whose [Enabled Domain](#enabled-domain) consists of time slots.
A **Timed Event** has a [Timed Domain Mode](#timed-domain-mode) and may have a [Timed Event Scheduled Span](#timed-event-scheduled-span).

Authoritative context: [FR-026](../requirements/functional/fr/FR-026.md) and [FR-074](../requirements/functional/fr/FR-074.md).

### Timed Slot

A discrete availability unit in a [Timed Event](#timed-event).
A **Timed Slot** has an [Instant](#instant) and is generated at the event's [Slot Increment](#slot-increment) within its [Enabled Domain](#enabled-domain).

Authoritative context: [FR-074](../requirements/functional/fr/FR-074.md) and [FR-017](../requirements/functional/fr/FR-017.md).

### Instant

An absolute point in time that identifies a [Timed Slot](#timed-slot) independently of its rendered timezone, [Civil Date](#civil-date), or clock label.

Authoritative context: [FR-013](../requirements/functional/fr/FR-013.md) and [FR-017](../requirements/functional/fr/FR-017.md).

### Civil Date

A calendar date interpreted in a specified timezone.
A **Civil Date** has no time or timezone meaning on its own.

Authoritative context: [FR-002](../requirements/functional/fr/FR-002.md) and [FR-074](../requirements/functional/fr/FR-074.md).

### Slot Increment

The configured interval between [Timed Slots](#timed-slot) generated for a [Timed Event](#timed-event).

Authoritative context: [FR-074](../requirements/functional/fr/FR-074.md) and [FR-075](../requirements/functional/fr/FR-075.md).

### Dates-Only Event

An [Event Kind](#event-kind) whose [Enabled Domain](#enabled-domain) consists of calendar dates.
A **Dates-Only Event** may have a [Dates-Only Event Scheduled Span](#dates-only-event-scheduled-span).

Authoritative context: [FR-026](../requirements/functional/fr/FR-026.md).

### Timed Domain Mode

The persisted configuration that determines how a [Timed Event's](#timed-event) set of [Active Slots](#active-slots) is maintained.
[Ranged Domain Mode](#ranged-domain-mode) configures it from an [Active Slot Range](#active-slot-range); [Custom Domain Mode](#custom-domain-mode) configures it through [Custom Domain Editing](#custom-domain-editing).

Authoritative context: [FR-067](../requirements/functional/fr/FR-067.md), [FR-053](../requirements/functional/fr/FR-053.md), and [FR-054](../requirements/functional/fr/FR-054.md).

### Ranged Domain Mode

A [Timed Domain Mode](#timed-domain-mode) that creates or restores the set of [Active Slots](#active-slots) from the [Active Slot Range](#active-slot-range) for the event's current [Picked Dates](#picked-dates).

Authoritative context: [FR-050](../requirements/functional/fr/FR-050.md) and [FR-075](../requirements/functional/fr/FR-075.md).

### Custom Domain Mode

A [Timed Domain Mode](#timed-domain-mode) that preserves a custom subset of [Enabled Slots](#enabled-slots) and exposes [Custom Domain Editing](#custom-domain-editing).
Switching to [Ranged Domain Mode](#ranged-domain-mode) restores the set of [Active Slots](#active-slots) from the event's [Active Slot Range](#active-slot-range).

Authoritative context: [FR-010](../requirements/functional/fr/FR-010.md), [FR-053](../requirements/functional/fr/FR-053.md), and [FR-054](../requirements/functional/fr/FR-054.md).

### Custom Domain Editing

The slot-level editing UI available only in [Custom Domain Mode](#custom-domain-mode).
It adds or removes [Active Slots](#active-slots) without changing [Picked Dates](#picked-dates) or the canonical persistence model.

Authoritative context: [FR-010](../requirements/functional/fr/FR-010.md).

### Picked Dates

The canonical selected-date membership for an event.
For [Timed Events](#timed-event), the date picker is a direct view of **Picked Dates** and picking a date regenerates the [Enabled Slots](#enabled-slots) for that date.

Authoritative context: [FR-050](../requirements/functional/fr/FR-050.md), [FR-051](../requirements/functional/fr/FR-051.md), and [FR-052](../requirements/functional/fr/FR-052.md).

### Enabled Slots

For a [Timed Event](#timed-event), the full slot domain for each [Picked Date](#picked-dates): the full civil day (`00:00` through the next `00:00` exclusive, in the [Event Timezone](#event-timezone)).
**Enabled Slots** are derived from **Picked Dates** and **Event Timezone**.

Authoritative context: [FR-074](../requirements/functional/fr/FR-074.md).

### Active Slots

The canonical subset of [Enabled Slots](#enabled-slots) that respondents can answer on.
The canonical invariant is `active slots subset of enabled slots`.

Authoritative context: [FR-016](../requirements/functional/fr/FR-016.md) and [FR-074](../requirements/functional/fr/FR-074.md).

### Event Timezone

The persisted timezone used to interpret [Picked Dates](#picked-dates) and generate [Enabled Slots](#enabled-slots).
Changing it preserves **Picked Dates**, rebuilds the [Enabled Domain](#enabled-domain), and filters [Active Slots](#active-slots) to the rebuilt domain.

Authoritative context: [FR-055](../requirements/functional/fr/FR-055.md) and [FR-057](../requirements/functional/fr/FR-057.md).

### Display Timezone

The timezone used to render [Timed Slots](#timed-slot) in the UI, controlled by the `Shown in` selector.
It affects projection and rendering only; it does not change [Picked Dates](#picked-dates), [Enabled Slots](#enabled-slots), or [Active Slots](#active-slots).

Authoritative context: [FR-013](../requirements/functional/fr/FR-013.md).

### Active Slot Settings

The persisted [Timed Event](#timed-event) configuration that determines [Active Slots](#active-slots).
It contains the [Timed Domain Mode](#timed-domain-mode) and, depending on that mode, either an [Active Slot Range](#active-slot-range) or the selected **Active Slots**.

Authoritative context: [FR-056](../requirements/functional/fr/FR-056.md) and [FR-067](../requirements/functional/fr/FR-067.md).

### Active Slot Range

The start/end window in [Active Slot Settings](#active-slot-settings).
In [Ranged Domain Mode](#ranged-domain-mode) it defines the [Active Slots](#active-slots) generated for the event's current [Picked Dates](#picked-dates).

Authoritative context: [FR-050](../requirements/functional/fr/FR-050.md) and [FR-075](../requirements/functional/fr/FR-075.md).

### Wipe Rule

On save, an active [Instant](#instant) outside the [Enabled Domain](#enabled-domain) is dropped.
For example, this includes next-day `00:00` and `00:30` instants from a cross-midnight window on a picked UTC date.
Changing the [Event Timezone](#event-timezone) does not change [Picked Dates'](#picked-dates) membership to preserve an otherwise out-of-domain slot.

Authoritative context: [FR-015](../requirements/functional/fr/FR-015.md) and [FR-057](../requirements/functional/fr/FR-057.md).

### End-of-Day Boundary

The picker option closing a [Civil Date](#civil-date), labeled `24:00` in 24-hour mode and `12 AM` in 12-hour mode.
Selecting it renders as `00:00` in the next [Projected Date Column](#projected-date-column).
The [Wipe Rule](#wipe-rule) drops next-day instants beyond the enabled day this boundary closes.

Authoritative context: [FR-091](../requirements/functional/fr/FR-091.md) and [FR-015](../requirements/functional/fr/FR-015.md).

## Domain Terms

### Active Domain

The answerable subset of the [Enabled Domain](#enabled-domain): what respondents may record availability on.
For a [Timed Event](#timed-event), it coincides with the set of [Active Slots](#active-slots) and obeys the domain-mode machinery.

Authoritative context: [FR-016](../requirements/functional/fr/FR-016.md) and [FR-053](../requirements/functional/fr/FR-053.md).

### Disabled Domain

Within the rendered grid axis, everything outside the [Enabled Domain](#enabled-domain).
Its visible representatives are [Disabled Padding Cells](#disabled-padding-cell) on [Timed Grids](#timed-grid) and grey non-picked dates carrying the `Disabled` status.

Authoritative context: [FR-014](../requirements/functional/fr/FR-014.md) and [FR-099](../requirements/functional/fr/FR-099.md).

### Enabled Domain

Everything generated from [Picked Dates](#picked-dates).
For a [Timed Event](#timed-event), these are full civil-day slots at the [Slot Increment](#slot-increment) through the next `00:00` exclusive in the [Event Timezone](#event-timezone).
For a [Dates-Only Event](#dates-only-event), these are the **Picked Dates** themselves.
Together the [Active Domain](#active-domain) and the [Inactive Domain](#inactive-domain) constitute the **Enabled Domain**; the [Disabled Domain](#disabled-domain) lies entirely outside it.
It generalizes the **Timed Event** definition recorded in [Enabled Slots](#enabled-slots).

Authoritative context: [FR-074](../requirements/functional/fr/FR-074.md) and [FR-075](../requirements/functional/fr/FR-075.md).

### Inactive Domain

The remainder of the [Enabled Domain](#enabled-domain) minus the [Active Domain](#active-domain).
It may be empty for any [Event Kind](#event-kind).
Its units coincide with [Enabled Inactive Slots](#enabled-inactive-slot).

Authoritative context: [FR-014](../requirements/functional/fr/FR-014.md).

## Time Format Terms

### Event Time Format

The 12/24-hour format used to render times inside event-editor forms, including the `What times might work?` and `Time range` dropdowns and start/end time fields.
It persists independently of the [Display Time Format](#display-time-format) under `eventTimeType`, defaults to 24-hour, and changing it during event creation saves the new preference.

Authoritative context: [browser date preferences](../../frontend/src/utils/browserDatePreferences.ts) and [event editor state](../../frontend/src/composables/event/useEventEditorState.ts).

### Display Time Format

The format used to render times in the event-page schedule grid and tooltips, controlled by the event-page `12h/24h` switch and persisted under `timeType`.
It defaults to 24-hour and affects event-page rendering only, not event-editor forms.

Authoritative context: [browser date preferences](../../frontend/src/utils/browserDatePreferences.ts) and [calendar grid](../../frontend/src/composables/schedule_overlap/useCalendarGrid.ts).

## Event Page Terms

### Dates-Only Event Owner Page

One physical route renders every event-page surface in this section.
This page shows a [Dates-Only Event](#dates-only-event) as seen by its [Event Owner](#event-owner): `Edit Event` is available and [Event Settings](#event-settings) are editable rather than presented read-only to others.
Timezone settings are owner-restricted here.

Authoritative context: [FR-018](../requirements/functional/fr/FR-018.md) and [FR-027](../requirements/functional/fr/FR-027.md).

### Dates-Only Event Page

One physical route renders every event-page surface in this section.
This page is the main page of a [Dates-Only Event](#dates-only-event), where every [Availability Response](#availability-response) of the event is visible.
It is the umbrella for the dates-only owner, scheduling, response editing, and response creation variants below.

Authoritative context: [FR-092](../requirements/functional/fr/FR-092.md) and [FR-101](../requirements/functional/fr/FR-101.md).

### Dates-Only Event Response Creation Page

One physical route renders every event-page surface in this section.
This is the page where one adds a new [Availability Response](#availability-response) to a [Dates-Only Event](#dates-only-event).
Saving requires at least one [Available](#available) or [If needed](#if-needed) date.

Authoritative context: [FR-004](../requirements/functional/fr/FR-004.md) and [FR-001](../requirements/functional/fr/FR-001.md).

### Dates-Only Event Response Editing Page

One physical route renders every event-page surface in this section.
This is the page where one can potentially edit an [Availability Response](#availability-response) of a [Dates-Only Event](#dates-only-event); authority varies with access mode ([Protected Response](#protected-response) or [Open Response](#open-response)) and [Platform Identity](#platform-identity) recovery.
It hosts the [Availability Editor](#availability-editor).

Authoritative context: [FR-060](../requirements/functional/fr/FR-060.md) and [FR-061](../requirements/functional/fr/FR-061.md).

### Dates-Only Event Scheduling Page

One physical route renders every event-page surface in this section.
This owner-only page schedules or reschedules a [Dates-Only Event](#dates-only-event), and its drawn or selected input is a single scheduled date that sets or clears the [Dates-Only Event Scheduled Span](#dates-only-event-scheduled-span) directly.
Only the [Event Owner](#event-owner) may schedule or reschedule the event here.
The phrases "specific-times page", "specific-times mode", and "specific-times editing" are rejected aliases for this page.

Authoritative context: [FR-012](../requirements/functional/fr/FR-012.md) and [FR-027](../requirements/functional/fr/FR-027.md).

### Timed Event Owner Page

One physical route renders every event-page surface in this section.
This page shows a [Timed Event](#timed-event) as seen by its [Event Owner](#event-owner): `Edit Event` is available and [Event Settings](#event-settings) are editable, in contrast with their read-only presentation to others.
Visibility differences are material under [Blind Availability](#blind-availability).

Authoritative context: [FR-018](../requirements/functional/fr/FR-018.md), [FR-039](../requirements/functional/fr/FR-039.md), and [FR-084](../requirements/functional/fr/FR-084.md).

### Timed Event Page

One physical route renders every event-page surface in this section.
This page is the main page of a [Timed Event](#timed-event), where every [Availability Response](#availability-response) of the event is visible.
It is the umbrella for the owner, scheduling, response editing, and response creation variants below.

Authoritative context: [FR-002](../requirements/functional/fr/FR-002.md) and [FR-064](../requirements/functional/fr/FR-064.md).

### Timed Event Response Creation Page

One physical route renders every event-page surface in this section.
This is the page where one adds a new [Availability Response](#availability-response) to a [Timed Event](#timed-event).
Saving requires at least one [Available](#available) or [If needed](#if-needed) slot.

Authoritative context: [FR-004](../requirements/functional/fr/FR-004.md).

### Timed Event Response Editing Page

One physical route renders every event-page surface in this section.
This is the page where one can potentially edit an [Availability Response](#availability-response) of a [Timed Event](#timed-event); authority varies with access mode ([Protected Response](#protected-response) or [Open Response](#open-response)) and [Platform Identity](#platform-identity) recovery.
It hosts the [Availability Editor](#availability-editor).

Authoritative context: [FR-060](../requirements/functional/fr/FR-060.md), [FR-061](../requirements/functional/fr/FR-061.md), and [FR-066](../requirements/functional/fr/FR-066.md).

### Timed Event Scheduling Page

One physical route renders every event-page surface in this section.
This owner-only page schedules or reschedules a [Timed Event](#timed-event): unlike the owner's default reading presentation, it accepts drawn input, and dragging sets or clears the [Timed Event Scheduled Span](#timed-event-scheduled-span) directly without painting availability.
Only the [Event Owner](#event-owner) may schedule or reschedule the event here.
The phrases "specific-times page", "specific-times mode", and "specific-times editing" are rejected aliases for this page.

Authoritative context: [FR-016](../requirements/functional/fr/FR-016.md), [FR-085](../requirements/functional/fr/FR-085.md), and [FR-086](../requirements/functional/fr/FR-086.md).

## Response Access Terms

### Event Owner

An [Event Guest](#event-guest) with additional authority to manage an event's [Event Settings](#event-settings).
An **Event Owner** may associate that ownership with a [Platform Identity](#platform-identity) through [Event Sign-In](#event-sign-in).

Authoritative context: [FR-018](../requirements/functional/fr/FR-018.md) and [FR-063](../requirements/functional/fr/FR-063.md).

### Platform Identity

The durable authenticated account identity for a person.
It can be associated with [Event Visitor Identities](#event-visitor-identity) for cross-browser recovery, but its raw value is not an event or [Availability Response](#availability-response) identifier.

Authoritative context: [FR-079](../requirements/functional/fr/FR-079.md).

### Platform Visitor

A person viewing any Timeful page.

Authoritative context: [FR-072](../requirements/functional/fr/FR-072.md).

### Event Visitor

A [Platform Visitor](#platform-visitor) viewing an event page.
Every **Event Visitor** has an [Event Visitor Identity](#event-visitor-identity) for that event.

Authoritative context: [FR-061](../requirements/functional/fr/FR-061.md) and [FR-073](../requirements/functional/fr/FR-073.md).

### Authenticated Event Visitor

An [Event Visitor](#event-visitor) signed in to a [Platform Identity](#platform-identity).

Authoritative context: [FR-062](../requirements/functional/fr/FR-062.md), [FR-063](../requirements/functional/fr/FR-063.md), and [FR-079](../requirements/functional/fr/FR-079.md).

### Anonymous Event Visitor

An [Event Visitor](#event-visitor) who is not signed in to a Timeful account.

Authoritative context: [FR-062](../requirements/functional/fr/FR-062.md) and [FR-063](../requirements/functional/fr/FR-063.md).

### Event Visitor Identity

The opaque, browser-local, event-scoped, non-authorizing identity established for every [Event Visitor](#event-visitor).
It persists across browser sessions and sign-out unless browser-local data is cleared.
A [Platform Identity](#platform-identity) may be associated with more than one **Event Visitor Identity** for an event, but is distinct from it.
The browser establishes an [Event Owner's](#event-owner) identity when event creation begins.

Authoritative context: [FR-073](../requirements/functional/fr/FR-073.md) and [FR-079](../requirements/functional/fr/FR-079.md).

### Event Visitor Control Credential (EVCC)

A private browser-held credential that authorizes a PostgreSQL [Event Visitor](#event-visitor) to manage every [Availability Response](#availability-response) owned by the visitor's [Event Visitor Identity](#event-visitor-identity) for that event.
It is distinct from the public, non-authorizing `eventVisitorId`, a [Platform Identity](#platform-identity), and legacy MongoDB response credentials.

Authoritative context: [FR-081](../requirements/functional/fr/FR-081.md), [QR-006](../requirements/quality/qr/QR-006.md), and [ADR-010](../design/architecture/adr/ADR-010.md).

### Granted Event Visitor Control Credential (Granted EVCC)

A private, browser-local, event-scoped credential issued to a target browser after a source-approved [Cross-Device Access Transfer](#cross-device-access-transfer).
It is distinct from the source EVCC, preserves the source's delegated role without transferring ownership, and remains usable until the target browser's data is cleared or the source revokes it.

Authoritative context: [FR-081](../requirements/functional/fr/FR-081.md), [FR-083](../requirements/functional/fr/FR-083.md), and [ADR-010](../design/architecture/adr/ADR-010.md).

### Event Guest

An [Event Visitor](#event-visitor) who can create and own multiple [Availability Responses](#availability-response) through an [Event Visitor Identity](#event-visitor-identity).
An [Event Owner](#event-owner) is an **Event Guest** with the additional authority to edit [Event Settings](#event-settings).

Authoritative context: [FR-001](../requirements/functional/fr/FR-001.md), [FR-018](../requirements/functional/fr/FR-018.md), and [FR-073](../requirements/functional/fr/FR-073.md).

### Event Owner Edit Token

An opaque, event-scoped credential that authorizes an [Event Owner's](#event-owner) [Event Settings](#event-settings) edits.
It is distinct from guest-response credentials and does not authorize guest-response edits.

Authoritative context: [FR-018](../requirements/functional/fr/FR-018.md) and [FR-063](../requirements/functional/fr/FR-063.md).

### Availability Response Edit Credential

The applicable opaque MongoDB credential that proves authority to edit a [Protected Response](#protected-response) before its owner is recoverable through an associated [Platform Identity](#platform-identity).
It is distinct from an [Event Owner Edit Token](#event-owner-edit-token) and the PostgreSQL EVCC model.

Authoritative context: [FR-062](../requirements/functional/fr/FR-062.md).

### Cross-Device Access Transfer

A PostgreSQL-only, source-confirmed browser process for granting another browser either a [Platform Identity](#platform-identity) session or delegated event authority.
The target displays a matching code that the source approves; the pending transfer is single-use and expires five minutes after creation.

Authoritative context: [FR-081](../requirements/functional/fr/FR-081.md), [FR-082](../requirements/functional/fr/FR-082.md), and [ADR-010](../design/architecture/adr/ADR-010.md).

### Blind Availability

An event privacy mode in which a non-owner may view and manage only the [Availability Responses](#availability-response) the non-owner is authorized to manage, without learning about other responses or their count.
The [Event Owner](#event-owner) may view every **Availability Response**.

Authoritative context: [FR-084](../requirements/functional/fr/FR-084.md).

### Availability Response

An [Event Guest's](#event-guest) recorded availability for an event, including its display name, access mode, and [Availability States](#availability-state).
Its ownership belongs to the creating [Event Visitor Identity](#event-visitor-identity), not its display name.

Authoritative context: [FR-001](../requirements/functional/fr/FR-001.md), [FR-060](../requirements/functional/fr/FR-060.md), and [FR-061](../requirements/functional/fr/FR-061.md).

### Availability State

The state an [Availability Response](#availability-response) assigns to a slot or date.
[Available](#available) and [If needed](#if-needed) count equally in overlap calculations.
[Unavailable](#unavailable) does not.

Authoritative context: [FR-006](../requirements/functional/fr/FR-006.md) and [FR-069](../requirements/functional/fr/FR-069.md).

#### Available

The [Availability State](#availability-state) value recording that a slot or date works for the respondent.
It contributes to overlap calculations exactly like [If needed](#if-needed) does.
Its canonical value spelling is the single capitalized word shown in this heading.

Authoritative context: [FR-006](../requirements/functional/fr/FR-006.md) and [FR-069](../requirements/functional/fr/FR-069.md).

#### If needed

The [Availability State](#availability-state) value recording that a slot or date works for the respondent when needed.
It contributes to overlap calculations exactly like [Available](#available) does.
The state label is always the two-word spelling shown in this heading; the hyphenated form appears only as an adjectival inflection, never as the state label itself.

Authoritative context: [FR-006](../requirements/functional/fr/FR-006.md) and [FR-069](../requirements/functional/fr/FR-069.md).

#### Unavailable

The [Availability State](#availability-state) value recording that a slot or date does not work for the respondent.
It does not contribute to overlap calculations.
Its canonical value spelling is the single capitalized word shown in this heading.

Authoritative context: [FR-006](../requirements/functional/fr/FR-006.md) and [FR-069](../requirements/functional/fr/FR-069.md).

### Availability Response Overlay

The elevated visual representation of the [Availability Response](#availability-response) being edited, rendered above other **Availability Responses** so its [Availability States](#availability-state) remain visible and editable.

Authoritative context: [FR-005](../requirements/functional/fr/FR-005.md).

### Protected Response

The default [Availability Response](#availability-response) access mode.
Only the [Event Guest](#event-guest) that owns the response may edit it through the applicable PostgreSQL EVCC authority, MongoDB [Availability Response Edit Credential](#availability-response-edit-credential), or associated [Platform Identity](#platform-identity).

Authoritative context: [FR-060](../requirements/functional/fr/FR-060.md), [FR-062](../requirements/functional/fr/FR-062.md), and [FR-073](../requirements/functional/fr/FR-073.md).

### Open Response

An [Availability Response](#availability-response) whose owning [Event Guest](#event-guest) has explicitly allowed any [Event Visitor](#event-visitor) to edit it.

Authoritative context: [FR-061](../requirements/functional/fr/FR-061.md).

### Shared Link

An event-scoped link whose validity gates disclosure of event metadata, respondent names, and availability data.
Without a valid **Shared Link**, the system shall not disclose that information.

Authoritative context: [QR-001](../requirements/quality/qr/QR-001.md) and [QR-002](../requirements/quality/qr/QR-002.md).

## Grid Rendering Terms

### Timed Grid

The event-page grid that renders [Timed Slots](#timed-slot) and their availability or scheduling states.

Authoritative context: [FR-002](../requirements/functional/fr/FR-002.md) and [FR-014](../requirements/functional/fr/FR-014.md).

### Enabled Inactive Slot

An [Enabled Slot](#enabled-slots) that is not active.
It remains editable in [Custom Domain Editing](#custom-domain-editing) but is not respondent-selectable; it has a treatment separate from both [Active Slots](#active-slots) and [Disabled Padding Cells](#disabled-padding-cell).

Authoritative context: [FR-014](../requirements/functional/fr/FR-014.md).

### Disabled Padding Cell

A grid cell without a mapped [Enabled Slot](#enabled-slots).
It is non-editable and uses a visually distinct unavailable treatment; it is not a slot.

Authoritative context: [FR-014](../requirements/functional/fr/FR-014.md).

### Grid Pointer

The interactive highlight marking the grid cell currently under the pointer during [Availability Editing](#availability-editing) or scheduling.
It is never rendered on collapsed-hours strips.

Authoritative context: [FR-095](../requirements/functional/fr/FR-095.md).

### Projected Date Column

A grid column derived from [Enabled Slots](#enabled-slots) projected into the [Display Timezone](#display-timezone).
A projected slot belongs to its display-local calendar-date column; an adjacent column is created when that date is otherwise absent, and a slot crossing midnight is not duplicated.

Authoritative context: [FR-002](../requirements/functional/fr/FR-002.md).

### Saved Active-Range Band

The read-only event-page grid's collapsed axis, derived from [Active Slots](#active-slots) and falling back to the [Enabled Domain](#enabled-domain) when there are no **Active Slots**.
The full civil-day axis appears only in [Custom Domain Editing](#custom-domain-editing) or with `Show all hours`.

Authoritative context: [FR-014](../requirements/functional/fr/FR-014.md).

### Event Scheduled Span

The optional scheduled occurrence of an event.
It can be saved, replaced, or cleared, and its concrete shape depends on the event's kind.
A [Timed Event](#timed-event) schedules a [Timed Event Scheduled Span](#timed-event-scheduled-span).
A [Dates-Only Event](#dates-only-event) schedules a [Dates-Only Event Scheduled Span](#dates-only-event-scheduled-span).
The phrase "Scheduled Event Time" is a rejected alias for this concept.

Authoritative context: [FR-012](../requirements/functional/fr/FR-012.md).

### Timed Event Scheduled Span

The [Event Scheduled Span](#event-scheduled-span) of a [Timed Event](#timed-event).
It is a time range selected as the event's scheduled occurrence, and saving requires a non-empty range.

Authoritative context: [FR-012](../requirements/functional/fr/FR-012.md), [FR-086](../requirements/functional/fr/FR-086.md), and [FR-113](../requirements/functional/fr/FR-113.md).

### Dates-Only Event Scheduled Span

The [Event Scheduled Span](#event-scheduled-span) of a [Dates-Only Event](#dates-only-event).
It is a single date recorded as the event's scheduled occurrence.

Authoritative context: [FR-012](../requirements/functional/fr/FR-012.md).

### Schedule Overlap

The availability calculated across the [Availability Responses](#availability-response) included in a **Schedule Overlap** view.

Authoritative context: [FR-064](../requirements/functional/fr/FR-064.md) and [FR-069](../requirements/functional/fr/FR-069.md).

### Event Settings

The settings that configure an event itself, rather than an individual [Availability Response](#availability-response).
Their scope includes the description, [Event Kind](#event-kind), [Picked Dates](#picked-dates), the [Event Timezone](#event-timezone) and [Display Timezone](#display-timezone), the [Event Time Format](#event-time-format) and [Display Time Format](#display-time-format), [Active Slot Settings](#active-slot-settings), and the [Active Slot Range](#active-slot-range).

Authoritative context: [FR-018](../requirements/functional/fr/FR-018.md).

### Dates-Only Grid

The grid that renders a [Dates-Only Event's](#dates-only-event) dates and their availability or scheduling states.
It is the counterpart of the [Timed Grid](#timed-grid).

Authoritative context: [FR-099](../requirements/functional/fr/FR-099.md) and [FR-100](../requirements/functional/fr/FR-100.md).

### Disabled Status

The `Disabled` label and treatment shown when interacting with non-answerable cells.
In a [Dates-Only Grid](#dates-only-grid) it appears when hovering a disabled date, mirroring the [Timed Grid's](#timed-grid) disabled wording and treatment.

Authoritative context: [FR-099](../requirements/functional/fr/FR-099.md).

## Availability Editing Terms

### Availability Editing

The activity of viewing, selecting, and editing [Availability Responses](#availability-response).
It takes place on the [Timed Event Response Creation Page](#timed-event-response-creation-page), [Dates-Only Event Response Creation Page](#dates-only-event-response-creation-page), [Timed Event Response Editing Page](#timed-event-response-editing-page), or [Dates-Only Event Response Editing Page](#dates-only-event-response-editing-page).

Authoritative context: [FR-041](../requirements/functional/fr/FR-041.md), [FR-065](../requirements/functional/fr/FR-065.md), and [FR-066](../requirements/functional/fr/FR-066.md).

### Availability Editor

The editing UI shown while an [Availability Response](#availability-response) is being added or edited.
It is present on the [Timed Event Response Creation Page](#timed-event-response-creation-page), [Dates-Only Event Response Creation Page](#dates-only-event-response-creation-page), [Timed Event Response Editing Page](#timed-event-response-editing-page), and [Dates-Only Event Response Editing Page](#dates-only-event-response-editing-page).

Authoritative context: [FR-005](../requirements/functional/fr/FR-005.md) and [FR-066](../requirements/functional/fr/FR-066.md).

## Authentication Terms

### Event Sign-In

Sign-in performed within an event-page context, associating the [Event Visitor Identity](#event-visitor-identity) and restoring guest access, or associating event ownership with a [Platform Identity](#platform-identity) when a valid token is presented.
It is distinct from platform-wide sign-in, whose availability is gated application-wide.

Authoritative context: [FR-062](../requirements/functional/fr/FR-062.md), [FR-063](../requirements/functional/fr/FR-063.md), and [FR-007](../requirements/functional/fr/FR-007.md).

### Magic Link

A registration or sign-in link sent by email that authenticates its recipient for the linked flow.

Authoritative context: [FR-031](../requirements/functional/fr/FR-031.md) and [FR-032](../requirements/functional/fr/FR-032.md).

## Product-Mode Terms

### Freemium

The product mode that enables advertising, access restrictions, and upgrade prompts.
When **Freemium** is disabled, those behaviors are omitted or bypassed.

Authoritative context: [FR-070](../requirements/functional/fr/FR-070.md), [FR-071](../requirements/functional/fr/FR-071.md), and [FR-072](../requirements/functional/fr/FR-072.md).

## UI Elements

### Legend

The event-page block labeled `Legend` explaining grid and response state colors.
It remains visible with zero responses and shows only states possible in the current mode.
It includes the [If needed](#if-needed) item for both [Event Kinds](#event-kind).

Authoritative context: [FR-009](../requirements/functional/fr/FR-009.md), [FR-092](../requirements/functional/fr/FR-092.md), and [FR-104](../requirements/functional/fr/FR-104.md).

### Show all hours

The event-page control that expands a [Timed Grid](#timed-grid) from its [Saved Active-Range Band](#saved-active-range-band) to its full civil-day axis.

Authoritative context: [FR-011](../requirements/functional/fr/FR-011.md), [FR-014](../requirements/functional/fr/FR-014.md), and [FR-048](../requirements/functional/fr/FR-048.md).
