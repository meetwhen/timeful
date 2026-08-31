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
- [ ] add flag to disable sign-in
- [ ] document flags that disable features
- [ ] make it possible to disable sign-in on backend
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
- [ ] improve readme
  - [ ] update the site link
  - [ ] add warning about unstability and possible loss of information and under construction
  - [ ] update technologies
- [ ] in FAQ, align text and +
- [ ] in FAQ, don't mention calendars when sign in is disabled
- [ ] How it works section still exists?
- [ ] should be able to edit specific times again
- [ ] Possibly lazily load timed slots when viewing to avoid browser-storage bloat?
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
- [ ] Introduce UDR - universal decision records along with ADRs?
- [ ] Don't mention a particular ADR in AGENTS.md
- [ ] Check that docs don't contain stale references or too specific references to other files. Such references can misguide agents
- [ ] Track line and branch coverage
- [ ] Using Postgres needs an ADR?
- [ ] Update backlog instructions to not run e2e and unit tests for docs-only changes
- [ ] Research how to update graphify semantic index when using OpenCode
- [ ] Review terms used in FRs and link where the link is missing
- [ ] ADRs should reference QRs, FRs may reference ADRs
- [ ] move adr to architectural-decision-records/adr? and add README.md near adr?
- [ ] Cap the number of guests to a smaller number than one tested via the QR, e.g. 200 if tested 500
- [ ] Log rotation
- [ ] Flags for Sign up with Google, Outlook.
  - [ ] Disabled by default (not allowed in Russia)
  - [ ] Affects the sign in flow
- [ ] Support adding calendars by link
  - Can't share a link to a calendar in the Google Calendar Android app
  - Service status dashboard - like <https://status.openai.com/>
- [ ] Send alerts to the platform hoster via email
- [ ] Define "instance operator" - someone hosting Timeful
- [ ] Write a system-manager config
  - [ ] fail2ban
  - [ ] necessary deps like docker and its dns
- [ ] A guest shall see which responses are editable by them and which are protected
- [ ] ADR candidate - Use postgres instead of mongo for better maintainability and speed - clear schemas.
- [ ] Use SQL/PGQ from Postgres 19
- [ ] On mobile, on event page, when there are no responses, center Show all hours within its column
- [ ] On mobile, when user added availability and need to write the guest name, the input form shall be above the keyboard, near the top of the page
- [ ] ADR - Use stalwart instead of listmonk
- [ ] When there's no response from the guest, "Add availability" shall blink.
      Clicking a timeslot isn't necessary to trigger blinking
- [ ] On mobile, make the lower panel with buttons visually separate from the grid (e.g. elevated and white). A green panel fuses with the green grid background
- [ ] Change the legend label for red: "Everyone is unavailable at this timeslot; Add/Edit availability"
- [ ] The description shall be inside "Edit event"
- [ ] Show Help button somewhere at the header. Provide there instructions depending on the role
  - For the event owner: Copy link and send it to others
  - For a guest without a response: Click Add availability below to mark in the grid when you're available
  - For a guest with a response: Add availability or Edit availability
- [ ] Rename variables in the code and modules to match the glossary
- [ ] Make a data-driven decision?
  - Load-test the go version
  - If too slow, rewrite in Rust
- [ ] Bug: on mobile, when I add availability and select a timeslot and drag pointer down and my finger is on the collapsed strip then while I hold the finger, the pointer stays at the lowest timeslot
- [ ] Add glossary term - pointer (box that highlights a cell)
- [ ] Document the traceability structure
- [ ] Automate generating tables in
  - docs/requirements/README.md
  - docs/design/README.md
- [ ] Mark primary FRs as `derived_from: is_primary`, not just an empty array which might mean that the sources are just not determined
- [ ] On mobile, when the reset button is clicked, show a little tooltip (or a banner) that time zone was reset to the event time zone
- [ ] When dragging to schedule an event, the tooltip shall show the range of the event, not of a timeslot
- [ ] On mobile, in a timed event, when scheduling an event and dragging, the timeslot pointer shall be visible
- [ ] Make the handoff skill wording prompt to create a handoff more certainly
- [ ] FR-005 - what are "overlays"?
- [ ] Make create-handoff.sh an app in flake.nix
- [ ] Define slot cursor
- [ ] Only the event owner shall see edit event, guests will see "event info"
- [ ] Sign out button
- [ ] Candidate ADR - the platform visitor identity shall not be compromised when the event visitor identity is compromised
- [ ] Enabled scenario: use event visitor id to revoke access for that visitor to responses in that event
- [ ] Blind event (only the event owner sees others' responses)
- [ ] Decompose root .gitignore into per-directory gitignores (graphify, infra, etc.).
      Instruct agents to use per-directory gitignores
- [ ] bug: in Timed Domain Mode section, there's more than one sentence on a line. Need to fix formatter to detect such problems
- [x] review glossary:
  - [x] Availability Editing
  - [x] rename: Scheduled event time -> Event Scheduled Time
  - [x] Slot Increment - not an interval between timed slots
  - [x] rename: Timed Grid -> Timed Event Grid
- [ ] Review FR wording
  - [ ] Depending on the context, use timed grid or full "timed event page grid"
  - [ ] Take note of consequtive tems `] [`
- [ ] bug: FR-012 - only owner can edit Scheduled Span, not everyone
- [x] Need to review requirements from the pov of the ownership model
- [ ] In new event form and using range, Validate the range
- [ ] list only allowed aliases in glossary, remove rejected
- [x] Remove time format from event creation/editing form
  - Decided to keep it for convenience
- [ ] Mention DDD - at least for naming, not for design
- [ ] Give names to parts of the Desktop UI
  - [ ] Left part of the header (with event info and buttons)
  - [ ] Right part of the header (where buttons are grouped)
  - [ ] Sidebar (on the right of the grid)
  - [ ] Indicator - thing selecting an option inside a toggle
- [ ] On the Event Response Editing Page, move Calendar options outside of its section:
  - [ ] on desktop - to under time format and time zone row
  - [ ] on mobile - to under More options row
- [ ] On the Event Response Editing Page:
  - [ ] replace More options with the single option inside it (Show all hours)
- [ ] On the Event Response Editing Page:
  - [ ] Show input form for editing the respondent name instead of Editing availability as
- [ ] What connects StageName, StageResult, PropertyGroupName to the rest of the system? (1,745 weakly-connected nodes — possible doc gaps)
- [ ] graphify can't process sql files - tree_sitter_sql — pip install 'graphifyy[sql]'
- [ ] What is the exact relationship between Sign-In Link and Platform Sign-In? (AMBIGUOUS edge)
- [x] On Mobile, on the Event Response Editing Page, make the bottom panel elevated
- [ ] On mobile, On the Event Page, don't show the switch (3d 7d if there are less than 7 days)
- [x] On mobile, on the Event Response Editing Page, in the activity bar, the Save button shall be solid green and without shadow, should be just like Save on desktop
- [x] On mobile, the Legend shall be fully visible and not hidden under the Available/If needed panel
- [x] The descriptions shall be editable in the event form
- [ ] What did we decide about the calendar - should days before today be selectible?
- [ ] In the new event form, move set specific times per day to the start of the section What times might work, right under the title

## MUST

- [ ] event in +3, edit specific times in +9, some time slots are lost
- [ ] event in +3,
- [ ] specific times uses the default 9-18 after unselecting 9-18
- [ ] when autofill is disabled, only manually should be enabled
- [ ] RIIR
  - [ ] use dbfirst
  - [ ] don't fix sign in functionality in the Go version
  - [ ] keep the original code in comments for line-by-line rewriting
- [ ] Remove split-gap
- [ ] there's no time when all 8 respondents are available should be over responses
- [ ] Specify API response normalization and handling of legacy respondent-name rows
- [ ] On timed event page, Create an event and Give feedback should be bold like on dates-only page
- [ ] In glossary, define "guest", "anon"
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
  - Empty weekly events lack a durable anchor week.
  - numResponses is a drift-prone cached value.
  - Anonymous-event adoption after OAuth currently updates Mongo only.
  - Dates-only events use civil dates, not instants
- [ ] what is when2meetHref?
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
- [ ] landing on mobile - decide which buttons should go into the hamburger menu
- [ ] Given I edit availability, I should see Editing availability as - add input field to write the name above Available
- [ ] Remove "Shown in"
- [ ] What is "compact" for?
- [ ] Revive the inspect scripts?
  - Restructure the inspect scenario to not require signed-in session
  - "The inspect scenario's prepare needs a signed-in session to open the "New event" dialog."
- [ ] Set up graphify
  - Preferably add a successfully buildable package to flake.nix' devshell
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
- [ ] On the event page, Edit event and Copy link buttons shall look the same on mobile and desktop
- [ ] Specify the context for each FR (where is it applicable)

## SHOULD

- [ ] The grid lines should be black, not grey
- [ ] User settings for the time format
- [ ] Cookie consent overlay
- [ ] Add concurrency control
  - Rule: When one user edits an event, others can only view
  - Scenario: one user edits the event, another one edits the availability
  - Rule: When one user has saved event changes (info, availability), others receive them immediately
- [ ] When `typescript-eslint` supports the installed TypeScript native bridge version in its peer dependency range, verify `npm ci --dry-run` succeeds without overrides and remove `--legacy-peer-deps` from `frontend/Dockerfile`.
- [ ] In Create event form, I should be able to write the event description:
  - potential problem: the description will be formatted differently than on the event page
  - another one: might distract the event creator
  - maybe just let them know that only they can edit the description?

## COULD

- [ ] anon identity to save preferences, maybe sign in by password
- [ ] use @dicebear/identicon for avatars, not generic head?
      The downside is that after updating the name, the avatar will change too.

## MUST - Done

## SHOULD - Done

## COULD - Done
