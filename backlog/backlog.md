# Backlog

## Inbox

Semi-structured TODO list

- [ ] Don't mention "legacy" in the code
- [ ] Add adr about using vue 3 and vuetify
- [ ] Speed up eslint and reduce memory consumption from 11GB
- [ ] remove `as unknown as`
- [ ] add eslint rule for `as unknown`
- [ ] is `$el` idiomatic modern syntax?
- [ ] Check against composition API
- [ ] Use the right palette consistently for dropdowns, selects, buttons, switches
- [ ] use clean layout-based fixes
- [ ] avoid !important
- [ ] make a design system
- [ ] refactoring - get rid of duplication
- [ ] add more instrumentation?
- [ ] don't modify vuetify internals (deep)
- [ ] load all routes lazily
- [ ] make more functions for business logic pure
- [ ] Get rid of eslint-disable-*
  - [ ] `eslint-disable vue/one-component-per-file`
- [ ] optional password for restoring. edit own responses and open for editing, can click the lock button to enter password and edit others' responses
- [ ] add instructions for the agent to write scripts for the browser and edit it instead of inline scripts?
- [ ] Resolve `schedule-overlap-mobile-scroll:8` under chromium-mobile: read-only grid collapses to a single `data-row="0"` 15px row after the create-dialog flow; unit contract is `ScheduleOverlap.collapsedHours.test.ts` ("collapses the read-only specific-times band to the saved active subset") — decide spec fix vs product fix
- [ ] Stabilize `ensureSpecificTimesEditorMode` / `isSpecificTimesEnabled` E2E toggling in `firefox-touch` and `chromium-mobile`; remove force-click fallback only after reliable native interaction coverage
- [ ] Investigate intermittent `firefox-touch` mobile tooltip/navbar layering failure. Verify Responses tap isolation, outside dismissal, and visibility/`mouseleave` behavior on a real Firefox mobile device; Playwright touch emulation is insufficient
- [ ] In `frontend/e2e/sign-up-form-event.spec.ts`, capture a non-OK event POST's status and response body when it recurs, then compare its payload with `server/routes/events.go` timed-payload normalization
- [ ] Before production deployment, run and record the preflight status of `server/scripts/20260724_canonical_timed_events`; deploy the server before a frontend release that requires derived enabled-slot validation
- [ ] If unit tests become a CI bottleneck, profile with `npm run test:unit:profile`; retain default workers and `slowTestThreshold: 100` unless target-CI benchmarks justify a change. Convert suites away from `happy-dom` only when verified Node-safe
- [ ] specify in the docs how the color is calculated at:
  - overlapping slots
  - best times
- [ ] Create an event for may 28 with availability from 0 to 4 and timezone +02:00. If you open the date picker in +0:00, should you see two days marked in the date picker?
- [ ] When adding availability, cancel and save should be aligned to the right
- [ ] Who are the group respondents?
- [ ] Make available/If needed on mobile higher to cover the "Adding availability" text
- [ ] Remove magic constants in CSS
- [ ] what is group (NewGroup.vue)?
- [ ] Handle narrow mobile screens (e.g. iPhone 17)
- [ ] "Editing availability as" shouldn't be in italic
- [ ] Collapse hours when editing response availability
- [ ] Mismatch between event page and specific times <http://127.0.0.1:4173/e/Eb67A>
- [ ] When creating event with specific times and setting timezone, on the specific times page, the timezone should be set to the one specified in the event (during creation). For a user, shown in should be the same on the event page and specific times page
- [ ] Lost slots on the specific times page when selected all slots at +7, then switched to +6
- [ ] <http://127.0.0.1:4173/e/EE2Fc> When editing event specific times, when "shown in" is +7, all slots for jun 15 are selected. When it's set to +6, jun 14 is duplicated and jun 15 is missing and there's a gap between jun 14 and jun 14.
- [ ] When I set specific times  0-4 on jun 14 and jun 16 when shown in is +5, then switch to +3,
      jun 15 doesn't appear although it should take some timeslots from jun 16
- [ ] what is the source of truth for enabled slots?
  - date picker may show dates that aren't shown when setting specific times.
    specific times editor should show these dates even if there are no enabled slots at these dates.
- [ ] Make a demo screenshot of <http://127.0.0.1:4173/e/6df78>
- [ ] At <http://127.0.0.1:4173/e/6df78>, when I hover over Maya Patel, I see if needed (yellow) for jun 20, 13:45-14:30 but it's available (green) when I edit her availability.
- [ ] make the feedback link configurable via .env and point to issues in my repo
- [ ] add flag to disable sign-in
- [ ] document flags that disable features
- [ ] make it possible to disable sign-in on backend
- [ ] add flag for enabling discord banner at the top (disabled by default)
- [ ] demo
  - first convert png to webp <https://picflow.com/convert/png-to-webp>
  - remember scale and screen size and make full page
  - then <https://ezgif.com/webp-maker>
    - 600 ms delay
    - 1200 ms for the last one
    - quality 90
  - <https://ezgif.com/webp-maker/ezgif-62a1b9a1b6704abf-split.html>
- [ ] rich landing enabled flag - enable more than just the title and demo
- [ ] font size of hours and days of week too large on mobile
- [ ] tooltip with time should appear where I hover cursor or release it
- [ ] on click on lock icon, show a tooltip with explanation why not editable
- [ ] improve readme
  - [ ] update the site link
  - [ ] add warning about unstability and possible loss of information and under construction
  - [ ] update technologies
- [ ] in FAQ, align text and +
- [ ] in FAQ, don't mention calendars when sign in is disabled
- [ ] How it works section still exists?
- [ ] should be able to edit specific times again
- [ ] For recurring events and normal events, only active slots are stored in the localstorage.
  - [ ] Possibly lazily loaded when viewing to not bloat localstorage?
- [ ] Everyone should be unavailable in responses when hover over red, light-grey, or dark-grey
- [ ] switch to when2meet in the repo
- [ ] make the app name configurable and when2meet by default
- [ ] why grey without grid in <http://127.0.0.1:4173/e/Eb67A> at gmt+9?
- [ ] reduce the number of columns on mobile so that the event can fit into that
- [ ] Summarize feedback on <https://www.reddit.com/r/schej/>
- [ ] Document the architecture and integration with external systems,
    e.g. Google Cloud project
- [ ] enable tests with chromium in addition to Firefox
- [ ] Refactor .env files
  - Group values
  - Shift the most important higher
- [ ] Schedule:
  - Phase 1 - everyone
  - Phase 2 - only the event owner, be it a registered or an anon user
- [ ] Move Show all hours to over Overlay availabilities on desktop and mobile
- [ ] Given an event was scheduled and time zone switched so that event slots shifted to another date, mark them blue in relevant dates
- [ ] Introduce a log of non-architectural decisions with SPEC-NNN identifiers
- [ ] Introduce an index that tracks the status of SPECs
- [ ] Set up CI/CD (maybe CD on releases only)
- [ ] On desktop, on the event page, when I scroll the grid and the top border of the grid isn't visible, then the space above sidebar shall be collapsed because it's not needed to separate Schedule event from the time format switch
- [ ] For <http://127.0.0.1:4173/e/C9ZC3WZS>, Edit availability is disabled. However, I can edit responses by clicking the pencil icon in responses
- [ ] During the sign-in, use green color for links
- [ ] `.env.staging.example` and `.env.production.example` shouldn't have `VITE_PREVIEW_*`

## MUST

- [ ] Event can be edited only by the creator. The description too
- [ ] event in +3, edit specific times in +9, some time slots are lost
- [ ] event in +3,
- [ ] specific times uses the default 9-18 after unselecting 9-18
- [ ] when autofill is disabled, only manually should be enabled
- [ ] RIIR
  - [ ] use dbfirst
  - [ ] don't fix sign in functionality in the Go version
  - [ ] keep the original code in comments for line-by-line rewriting
- [ ] Remove split-gap
- [ ] given new days are added in edit event with specific times, when the specific times page is opened, then the new days should be available for editing
- [ ] there's no time when all 8 respondents are available should be over responses
- [ ] On a non-existing page, 404 must be centered vertically and horizontally
- [ ] Given I'm on the event page in read mode, when I click somewhere in the red zone, then the Add availability and/or Edit availability must blink
- [ ] when user clicks edit availability, they should see a drop-down list of all respondents whose availability they can change. own availabilities first, open for editing next, people with password last
- [ ] Enforce that the user name is always non-empty

      The core of this is a functional requirement:

      - which guest names are accepted or rejected
      - how respondent names are normalized
      - what data the API returns
      - how duplicates and legacy rows are handled

      It has some non-functional aspects:

      - data quality and consistency
      - robustness against malformed input
      - maintainability of one canonical contract
- [ ] The event that spans two dates in the display time zone must appear on both dates
- [ ] On timed event page, Create an event and Give feedback should be bold like on dates-only page
- [ ] In glossary, define "guest", "anon"
- [ ] Password protection in responses for anon guests
- [ ] For date-specific events, make the dates not disappear
  - For <http://127.0.0.1:4173/e/JTGTEFXY>, the dates disappeared several times,
    maybe because the agent run a container with another db volume
- [ ] Install buildx on the VM.
      Logs when deploying:

      ```text
      time="2026-08-13T23:29:13+03:00" level=warning msg="Docker Compose is configured to build using Bake, but buildx isn't installed"
      ```
- [ ] Product questions
  - Anonymous event metadata is publicly editable.
  - Selected schedules are publicly replaceable and clearable.
  - Blind-response viewing accepts guest ID/name without an edit token.
  - Guest edit tokens are stored in plaintext.
  - Dates-only values are stored as instants rather than explicit civil dates.
  - Responses may contain slots outside the event's active slots or dates.
  - Timed enabled slots span the full civil day, not the configured local-time window.
  - Empty weekly events lack a durable anchor week.
  - numResponses is a drift-prone cached value.
  - Anonymous-event adoption after OAuth currently updates Mongo only.
  - Dates-only events use civil dates, not instants
- [ ] what is when2meetHref?
- [ ] Allow duplicate response names?
- [ ] For dates-only events, when I edit event and enable Overlay availability, then I shall see all responses.
  - Respondent's response is overlaid
  - Others' responses are shown
  
  Problem:
  - Currently, I see none of the responses
- [ ] On mobile, for dates-only events, show the buttons Overlay availability and Start on Monday
- [ ] Failing e2e test "sign-up blocks are visible on the event page" — unrelated pre-existing: the spec seeds a sign-up event with legacy fields (duration, dates, timeIncrement, startOnMonday), which POST /api/events has rejected with 400 legacy-timed-event-field:* since commit f7817601 (2026-07-30). Reproduced via direct curl against the test API.
- [ ] add i18n (Russian, German)
- [ ] Update demo
  - Hover over participants, then grid, then select best times, then create event on timeful
- [ ] Only response creator can edit these responses if they're not publicly editable - protected with a password or editable only on the creator's device
- [ ] landing on mobile - decide which buttons should go into the hamburger menu
- [ ] Given I edit availability, I should see Editing availability as - add input field to write the name above Available
- [ ] On the event page, the timeslot highlighting box borders shall not overlap with the timeslot borders
- [ ] In the sidebar, show:
  - "Display time format" above the time format switch
  - "Display time zone" abovt the time zone switch
- [ ] On mobile, on the event page when there are no responses, the Show all hours toggle shall be centered
- [ ] Display time zone resets to the event time zone
- [ ] In the event form, there should be no reset button for the event time zone
- [ ] What happens to the availabilities when the time zone changes?
  - Re-anchoring?
  - Or, forced removal of availabilities outside of the active time slots?
  - Or, keep availabilities and re-anchor only active time slots?
- [ ] Add View event form for non-editors of an event to see the event properties.
- [ ] Remove "Shown in"
- [ ] What is "compact" for?
- [ ] Revive the inspect scripts?
  - Restructure the inspect scenario to not require signed-in session
  - "The inspect scenario's prepare needs a signed-in session to open the "New event" dialog."
- [ ] Set up graphify
  - Preferably add a successfully buildable package to flake.nix' devshell
- [ ] On mobile, on event page, when there are no responses, Show all hours shall be centered inside its column
- [ ] On mobile, when I have added availability and clicked Save and need to input the name, the input field shall be fully visible above the keyboard so that I see what I type
- [ ] Refactor ADRs
  - [ ] Move ADRs from `frontend/adr` to `adr` (repo root)
  - [ ] Add README that explains the ADR format
  - [ ] Specify the scope inside the ADR - frontend/backend
  - [ ] in the README, generate tables with ADRs (ID, title) for frontend and backend
    - Maybe use mdsh
  - [ ] Document important quality attributes - performance, maintainability, reliability, security, usability
  - [ ] Select true ADRs that affect important quality attributes
- [ ] Introduce specs
  - [ ] Identifiers start with `SPEC-`
  - [ ] Current not "true" ADRs can be the first specs
  - [ ] SPECs are affected by ADRs
- [ ] ADR candidate:
  - (Scope: frontend, backend): backend handles only particular paths for initial HTML with essential metadata
  - Using Crockford base32 encoding (8 characters) with repeated probings for event identifiers (less collisions)
- [ ] Use the following sign-up flow:
  - User enters email
  - Timeful sends a magic link to the email
  - In email, user clicks the link
  - User is redirected to the Timeful site
  - User enters a password, first name, last name
  - Timeful saves email, password, first name, last name
  - Timeful signs in the user
- [ ] Use the following sign-in flow:
  - User enters email and password
  - If user forgot a password:
    - User clicks Forgot password?
    - Timeful sends a magic link to the email
    - Timeful suggests to use it to restore the password
    - User opens the email and clicks the link
    - User is redirected to the Timeful site
    - User enters a new password
  - If email is not recognized, Timeful suggests to sign up
  - If email is recognized, sign in succeeds
- [ ] Document features of a good email sending service
- [ ] Document DNS records
- [ ] On mobile, to mirror the Desktop version:
  - Show best times/Show best days and More options shall be in the first row
  - Time format, time zone, and the number of days shall be in the second row
- [ ] On the event page, Edit event and Copy link buttons shall look the same on mobile and desktop

## SHOULD

- [ ] The grid lines should be black, not grey
- [ ] User settings for the time format
- [ ] Cookie consent overlay
- [ ] Add concurrency control
  - Rule: When one user edits an event, others can only view
  - Scenario: one user edits the event, another one edits the availability
  - Rule: When one user has saved event changes (info, availability), others receive them immediately
- [ ] When `typescript-eslint` supports the installed TypeScript native bridge version in its peer dependency range, verify `npm ci --dry-run` succeeds without overrides and remove `--legacy-peer-deps` from `frontend/Dockerfile`.
- [ ] In Create event form, I should be able to write the event description
  - potential problem: the description will be formatted differently than on the event page
  - another one: might distract the event creator
  - maybe just let them know that only they can edit the description?

## COULD

- [ ] anon identity to save preferences, maybe sign in by password
- [ ] use @dicebear/identicon for avatars, not generic head?
      The downside is that after updating the name, the avatar will change too.

## MUST - Done

- [x] use `.env.example` instead of `.env.template`
  - We use `.env.development.example` and `.env.production.example`
- [x] multi-day <http://127.0.0.1:4173/e/5Ef6f>
- [x] Let's nest all toggles under Options dropdown and make it open by default
  - Show best times is always visible, other can be opened via Options
- [x] overlay availabilities - each slot has a solid frame <https://timeful.fun/e/c762cA>
- [x] there should be a space between grids for non-consecutive days
- [x] edit event button is missing
- [x] shown in should be the same size as the timezone and black
- [x] the description text should be at the same vertical position when not editing and when editing
- [x] replace the Create event button on the main page with the actual form
  - No, there's additional useful info about the app on the main page
- [x] on the event page, near "shown in", the underline colors for the timezone and time should be the same
- [x] move adr to the repo root
  - No, keep adr for frontend inside frontend
- [x] use full "development" instead of "dev" and "production" instead of "prod"
- [x] Can't save time for the first time
- [x] In create event -> advanced options, "Timezone" should be black
- [x] Show full grid by default
  - It's collapsed by default
- [x] add option to collapse unused hours (not hide)
- [x] make each collapsed hours uncollapsible
- [x] mobile version - switching between 3 days and 7 days doesn't work
  - Works now
- [x] Add availability only over the responses section so that it can scroll
  - Moved to the header
- [x] When viewing an event, when clicking, dragging the box pointer in the red area, then unclicking, the box disappears
- [x] in a collapsed section, upper line overlaps the original line but the lower line doesn't
- [x] use the same colors for all grid lines
- [x] the selection box is almost invisible
  - Now it's hatched
- [x] use the same frequency of dashes for the selection box and the grid separator at half an hour
  - we use solid selection box
- [x] when I click somewhere, the drop-down list in edit availability doesn't disappear
- [x] let's collapse hours when they're at the start or at the end too. These hours are useless anyway
- [x] Check +3:30 and +5:45
- [x] when editing event, the week day every letter looks the same as the day of month
- [x] create event with specific availability in +2, 0-4 (day 1), 0-4 (day 2). When opened in 0:00, should see the previous date
- [x] box cursor doesn't follow the mouse in the specific times grid
- [x] no Options when there are no responses
  - now there's an option to show all hours
- [x] disallow editing past dates
  - Not disallowing because it's a feature
- [x] after adding availability, the selected segments are dark green
- [x] Move "Copy link" closer to the event title
- [x] too drastic width change at page width breakpoints (1/3 of the screen)
- [x] Make edit event a button with a pencil icon
- [x] When no responses, show only add availability and show all hours
- [x] show pencils
- [x] When one response, doesn't show best times
- [x] There's no source of truth for dates for specific times and event
  - Added ADR-012
- [x] do we have recurring events?
      Yes — it uses a custom TimedRecurrence model with two kinds: specific_dates (explicit date list) and weekly (day-of-week pattern). No iCalendar RRULE support.
- [x] In the new event form, in Advanced options, The Timezone text isn't aligned horizontally with "Time increment"
      The styling differs too
- [x] On the event page without responses, there should be only Add availability and Show all hours, not very wide Add availability and More options

      "Show all hours" should be under "Add availability", as before.
      The toggle and "Show all hours" text should be centered vertically within their box
- [x] Create event with 9 - 17, timezone +9 (<http://127.0.0.1:4173/e/ee4Cb>), june 11 and 12.
      Expected: june 10 and 11 are shown on the event page
      Actual: june 11, 12, 13 are shown there
- [x] Spacing between lines is different when editing and viewing event description
- [x] Shown in shouldn't affect the time zone
- [x] Availability not rendered at <http://127.0.0.1:4173/e/Eb67A>
- [x] <http://127.0.0.1:4173/e/Eb67A> shown in GMT -7 shows blank grey columns
- [x] <http://127.0.0.1:4173/e/Eb67A> doesn't show the bottom separator
- [x] When there are no responses, Add availability and show all hours are too wide
- [x] in What days might work, when sun and mon are selected, when enabling start on monday, both mon and sun must be selected
- [x] remove formerly known as schej and flag
- [x] on the event page, responses should be aligned top with the grid top
- [x] add flag to enable privacy policy
- [x] use the same palette for event page
  (enabled - light-grey, active -red; disabled - dark-grey)
- [x] legend should be visible even if no responses. Show only enabled/active
- [x] show all hours should work on the event page
- [x] hide dates where there's nothing to pick?
  - no, keep all dates
- [x] the day ends at 24
   No, at `00:00` for consistency with 12AM in 12-hour format
- [x] there should be a label for each grid row that marks the start of an hour, even for collapsed
- [x] get rid of 12-hour format?
   No
- [x] remove Add availability, leave just edit availability
   No. In this case, we won't be able to add availability for someone
- [x] don't uncollapse rows when scheduling
- [x] show event when rescheduling
- [x] Schedule button on mobile should be filled blueish when clicked
- [x] When there's a scheduled event and I edit availability, I shold draw only with green cells, not blue cells (scheduled)
- [x] when editing availability, overlay availabilities should preserve the marked time slots and not shift
- [x] set specific times - just select, don't draw scheduled event
- [x] on mobile, edit availability is permanently greenish although should be like that only after a click
  - It's ok - it's disabled
- [x] on mobile and desktop, there should be no delay between selecting the time slot and seeing the tooltip with time and date
- [x] on mobile, I shouldn't be able to click the grid through that Responses offcanvas
- [x] on mobile, the tooltip should stay near selected timeslot and shouldn't move around
  - now we store the info about the selected timeslot and render based on its position
- [x] in commit skill, make the script for session identifier and time a single script
- [x] on mobile, when scrolling, the tooltip with time should scroll together with selected cell, not stay frozen at the same height of the screen while the screen scrolls
- [x] on mobile, when no slot is selected, the tooltip shouldn't be visible
- [x] When a timeslot is selected, it's position is saved.
      When the page is reloaded, the selection is rendered at the timeslot.
      However, there's no tooltip near the selection.
- [x] on mobile, when long press inside the grid changes the selected timeslot, the tooltip must also appear near the selected timeslot
  - currently, it stays near the previously selected timeslot
- [x] on mobile, when long press inside the grid makes the timeslot selected, the tooltip must appear near it.
  - currently, the tooltip doesn't appear
- [x] on mobile, the tooltip should be under top navbar when the grid gets scrolled up
- [x] On the event page without responses, there should be only Add availability and Show all hours, not very wide Add availability and More options

      "Show all hours" should be under "Add availability", as before.
      The toggle and "Show all hours" text should be centered vertically within their box

- [x] "Best times" toggle should appear when there is at least one response, not more than one
- [x] When Best times and More options are both visible, they should be side by side
- [x] On mobile, on the event page, when swiping inside the grid, the grid should scroll just like when swiping outside
- [x] "select in Add/Edit availability" -> "change in ..."
- [x] On mobile, on the event page, when not interacting with the grid, make it scrollable with finger on the grid
- [x] Disabled padding cells - "Unavailable, outside the event dates in the event timezone"
- [x] In desktop app, in firefox, the selected timeslot must follow the mouse when it moves inside the grid, no matter clicks.
- [x] doesn't show time on hover after clicking and moving the cursor - <http://127.0.0.1:4173/e/dEeaF>
  - Can't reproduce
- [x] on mobile, tooltip should be below the Responses offcanvas panel
- [x] get rid of the comparator, leave just inspect and update docs
- [x] Given I'm on the event page, when I move the mouse cursor out of the grid, the Responses must show just the number of responses and not show who's available
- [x] In the desktop version, the alignment of rows in add/edit availability should be similar to the event page.

      - event title - Cancel, Save
      - Edit event - Show all hours
      - Add description - shouldn't be visible, actually
- [x] make scheduled event color #76AFF2 so that it's visible on the green background
- [x] Given on the event page, when scrolled down, then clicked Edit event, then the navbar (timeful, create an event, etc.) must not move higher
- [x] On the event page, align:
      - event title with Add availability/Edit availability
      - "Edit event" and Copy link with Show best times and More options
      - The event description with Schedule/Reschedule event
- [x] When changing availability, there should be these buttons:
  
  - Cancel | Save
  - Overlay availability | More options (Show all hours)
  - Delete
- [x] show all hours should show all hours, not trimmed. Currently, it trims wrong
  - can't reproduce
- [x] Given on the edit availability page, when no timeslot is marked as available/if needed, then the Save button should be disabled
- [x] Allow multiple scheduled events.

    Remove clear in Reschedule event because one can just wipe it by dragging the event cursor.

    Should behave just like the availability cursor.

    Given I clicked Reschedule event,
    when I click a slot that belongs to the existing event, the slot gets cleared
    and when I click a free slot, it becomes an event slot
    and adjacent slots merge into a single event

  - Won't implement because it makes the interface much more complicated.
    If there's 0 slots selected, how to save?
    If there's several events selected, need to warn that exactly one event is required to schedule on Google Calendar
- [x] On desktop, the width of buttons when adding availability should be the same as on the event page
- [x] On desktop, when rescheduling event, add availability, edit availability, edit event should be disabled
- [x] On desktop, Shown in and timezone should be above Responses and scroll with it so that when all timeslots are shown
     and one selects a timeslot and see the tooltip with the time, they can see which timezone that time belongs to
- [x] Use monospace font for times
  - [x] in the grid
  - [x] in the tooltip
- [x] Don't show the cursor and tooltip in enabled, inactive or disabled cells
- [x] On mobile, given I'm on the event page and a timeslot is selected and the tooltip is visible, when I click the responses offcanvas, the selected timeslot and its tooltip must not disappear
- [x] On mobile, make the grid labels fit on the screen
- [x] Given I'm on desktop and on the event page, when I hover over collapsed hours,
  - the responses should show 0/N (behave similarly to disabled cells)
  - and mark everyone unavailable
  - and the selection at the highlight at the last hovered timeslot must be cleared
- [x] Given the selected dates are non-consecutive and make the grid split into sub-grids (e.g. Aug 6 and 9), when I click (desktop, mobile) or hover (desktop) in the space between sub-grids, the highlight and tooltip at the last selected timeslot must be cleared like when clicking or hovering inactive timeslots
- [x] On desktop, given I'm on the event page, when I hover an active timeslot or hover outside the grid and then hover over an inactive slot, then the Responses sidebar should show (0/N) and all responses crossed-out and the status square should have the corresponding color (light-grey for enabled, inactive etc.)
- [x] On mobile, buttons with arrows for switching pages when there are several days, buttons should be the same size
- [x] On mobile, when I click a disabled timeslot, the selection should disappear
- [x] On desktop, given I'm on the event page, when I hover or click inside the grid outside active cells, the highlight at the last timeslot and any tooltip should not be visible
- [x] Create a glossary and oblige agents to read it in AGENTS.md
- [x] In responses,
  - when I hover or click in a grid, show the square for the status of a person at the corresponding timeslot (available, if needed, unavailable, etc.) instead of a profile image;
  - when I hover or click outside of the grid, replace the status square with a profile icon like it's now
- [x] Add "Disabled, collapsed" legend item
- [x] dashed outline in the legend bullet for Disabled, collapsed
- [x] On desktop, in the new event form and on the new event page, when I scroll the time zone menu, the width is always 520 and doesn't change based on the content length.
- [x] Schedule event on Google Calendar should happen in the display timezone
- [x] Given I'm on the event page and there's a scheduled event, when I click Reschedule event, the Schedule button is active
- [x] On the Event not found page, when using tab-navigation, the rectangle should coincide with the Back to home button
- [x] On desktop, on a dates only event page, in the second row, on the right:
  - When there are no responses, show "Start on Monday" option in the second row
  - When there are responses, show "Show best days" option and "More options" (which includes "Start on Monday" and "Hide if needed")
- [x] For dates-only events, the combination of days of week and dates in the calendar must match the reality
- [x] For dates-only events, the color of disabled dates must be dark-grey like for disabled padding cells in timed event
- [x] On the Event not found page, the "Back to home" button must have a black shadow (like on the home page), not greenish glow
- [x] "If needed" - show in the legend to explain the status color in responses
- [x] On dates-only event page, the Responses top of the text must horizontally coincide with the top edge of the grid
- [x] Keep yellow status for "if needed" responses but don't:
  - highlight such responses text
  - show asterisk ("\*") near such responses
  - show "* if needed" under responses
- [x] On the read-only timed event page, when I switch the timezone, the grid shouldn't collapse
- [x] In timed range events, the Legend shall show the light-gray "Disabled, change in Edit event"
- [x] Address linter warnings
- [x] In "New event" form, when "Dates only" is selected, in "Advanced options", there should be no "Time increment"
- [x] On the dates-only event page, when I click the "Edit event" button, then the edit form opens
- [x] On dates-only event page, there must be space between the grid and Responses sidebar
- [x] On the dates-only event page, show the Schedule button even if there are no responses, like on the timed event page.
- [x] The width of the description field must be the same as on the timed event page
- [x] On the dates-only event page, "Start on Monday" must be centered vertically with "Add availability"
- [x] When VITE_ENABLE_SIGN_IN is not false, on the landing page, there should be the Sign in button.
      When I click that button, "Sign in" page opens with options like Google Calendar / Email.
- [x] On each page, Sign in must be the left-most button
- [x] For date-specific events, when I edit the timezone in the form, it shall be persisted.
  - Currently, it's always set to GMT+0:00
- [x] Given I sign in, when I enter an unregistered email and click Continue with email:
  - The input field is highlighted red
  - The error appears like on accounts.google.com: (red alert icon) "Couldn’t find this account. Create account"
- [x] Consistently use "Sign in" and "Sign up" in the sign in flow.
- [x] Don't provide support text in the Sign up form because such text is redundant.
- [x] Use TypeScript 7 to speed up compilation
- [x] Use oxlint to speed up linting
- [x] Update references from deemp/timeful to whensync/timeful
  - Reserved whensync just in case I'd like to rename the project in future
- [x] Fix event description card
  - Multi-line event descriptions save with newline characters preserved.
  - Saved multi-line descriptions render across multiple lines.
  - Read-only description pencil is in the upper-right corner.
  - Edit mode preserves the description’s text width and wrapping.
  - Edit mode cancel and save buttons are in the upper-right corner.
  - Description action buttons are reduced to 32px and visually centered for single-line descriptions.
- [x] Hovering a disabled date on a dates-only event now gives respondents a disabled status, not active/unavailable.
- [x] The dates-only legend now shows Unavailable, change in Add/Edit availability, matching timed events.
- [x] Fix highlighting in dates-only events
  - Enabled dates-only grid cells receive a black inset square frame on hover.
  - The frame stays inside the cell and does not overlap neighboring cells or the date number.
  - The date text does not change from border-box layout shifts.
  - Hovering a disabled cell removes the prior enabled-cell frame.
  - Hovering a disabled cell shows disabled status in Responses.
  - Leaving the grid removes the current frame.
  - On mobile, tapping a disabled cell removes the current frame while showing disabled Responses status.
  - On mobile, tapping outside the grid removes the current frame and clears its state.
- [x] Fix gap on the right for dates-only events.
  - Dates-only grid fits within narrow viewports from 320px through 639px.
  - Grid has a 16px left and right gutter on phone layouts.
  - No horizontal document overflow at 320, 390, 410, 480, 639, and 640px.
  - Below 640px, the sidebar remains stacked beneath the calendar and fits the viewport.
  - At 640px, grid and sidebar remain side-by-side with their intended 16–20px gap.
  - The full-width dates-only grid no longer adds external margin beyond its pane.
- [x] On the dates-only event page, the date cells must be rectangular 1:2 (height:width) to better fit the screen
- [x] Improve separation between environments (development, test, staging, production)
- [x] Don't run vite in a container to keep things simple and fast during development
- [x] <http://127.0.0.1:4173/e/aB3BE>
  - ok in gmt+3
  - padding gets added in +3:30 although not needed
  - grid split in +12 although shouldn't
  
  Status: can't reproduce
- [x] Add `flake.nix`
- [x] rename PR (<https://github.com/schej-it/timeful.app/pull/250>) to `Modernize the app`
- [x] Use Node 26.5.0 everywhere
- [x] On the specific times page <http://127.0.0.1:4173/e/Eb67A>, when I switch timezone from +5 to +6, the left-most upper-most enabled slot should become disabled
- [x] On the specific times page <http://127.0.0.1:4173/e/Eb67A>, when I switch timezone from +5 to +4 and on june 14, 0-4 are selected, jun 13 should appear
- [x] In the new timed event form, Advanced options must show the Time increment.
- [x] On specific times edit page, show the time format and timezone menu between the instructions (Click and drag ...) and the Legend
- [x] On specific times edit page, align top edge of the date (month + day) with top edge of the instruction "Click and drag ..."
- [x] "Legend" font size must be the same as "Responses" font size
- [x] Use "Legend" instead of "Legend:"
- [x] Align upper edge of the sidebar with the upper edge of the grid
- [x] The right edge of the navigation and header (buttons in the upper-right corner) and the sidebar must coincide
- [x] Reduce the gap between the grid and Responses
- [x] Make compose files fully configurable via required values from corresponding .env files
- [x] Switch to Postgres for new anon events
- [x] Make 8-character Crockford base32 ids point to events in Postgres
- [x] Migrate to postgres and report new links:
  - <http://127.0.0.1:4173/e/BFfB4> -> <http://127.0.0.1:4173/e/M7FVZFYP>
  - <http://127.0.0.1:4173/e/6df78> -> <http://127.0.0.1:4173/e/C9ZC3WZS>
- [x] On the timed event page, in More options, "Show all hours" must be above "Hide if needed times"
- [x] On the dates-only event page, in More options, "Start on Monday" must be above "Hide if needed days"
- [x] Spaces between elements in the sidebar must be the same (Timezone-Responses-Legend)
- [x] Set up caching to speed up E2E tests
- [x] Collapsed hours should work in range events too, not only specific-times events
- [x] On dates-only event page, when I click the "Edit event" button, a form for editing the event opens
- [x] Store only active timeslots and calculate all other types (enabled inactive and disabled) on the fly
- [x] add Playwright e2e tests from the comparator to the main repository
- [x] commit comparator code to the main repository until it works as expected
  - not included because we don't use a comparator anymore
- [x] Use the same Node 26.5.0 for frontend in dockerfile and in dev
- [x] Check whether NODE_ENV and GIN_MODE are in the example .env files
  - NODE_ENV isn't there because it's not used
  - GIN_MODE is there
- [x] introduce staging environment
- [x] On the event page:
  - when there are no responses, the Edit availability shall be not visible.
  - When there are responses, shall be visible
- [x] On the event page, Edit event button shall be under the event title
- [x] On the event page, Github repo button shall be in the navbar, right-most button
- [x] In the README, the Timeful icon shall be readable both in night and light modes
- [x] On the event page, when I hover over collapsed hours strip, neither the strip nor any cell is highlighted like active cells are highlighted on hover.
- [x] Bug: when scrolling the gri, the No responses yet changes the position
  - Can't reproduce
- [x] On desktop, on the timed event page, time format and timezone shall be above Responses in the sidebar
- [x] Create an event with specific times for dates Aug 30, 31, mark hours 0-4 for both dates, edit event, set dates for 28, 29, click next. See May 28, 30, 31 in specific times page, and May 30, 31 on the event page.
  - Can't reproduce. I see Aug 28, Aug 29 on the event page
- [x] Create an event with two dates, mark timeslots for only one day in specific times, save, edit again and see only one day on the event page
  - Can't reproduce. I see both days.
- [x] On mobile, on the timed event page, the timezone control keeps a fixed width (112px) so the reset button fits inside without resizing the control, with the time format and days-per-page buttons on either side
- [x] On the Create your account form, when I click Continue and the OTP can't be sent, a report about that is visible
- [x] The icon to reset the timezone shall have a counter-clockwise arrow
- [x] On the timed event page, timezone button and time format button shall together span the full row in the sidebar, the time format aligned right.
- [x] On the timed event page, the time format shall be on the left from the time zone so that the timezone button can get a reset button or stay the same freely
- [x] On mobile, on the event page, tooltip shall be fully on-screen
- [x] On mobile, when adding availability, Show all hours should be an option at its normal place - under other options between the event description and the grid
- [x] Bug:
  - on mobile
  - I add and save availability
  - I click an active timeslot
  - Responses panel with my response slides from the bottom
  - I click Edit availability
  - I click a grid timeslot
  - Expected: on the panel that slides from the bottom, I see only Available/if needed
  - Actual: on the panel, I see both Available/If needed and Responses

## SHOULD - Done

## COULD - Done
