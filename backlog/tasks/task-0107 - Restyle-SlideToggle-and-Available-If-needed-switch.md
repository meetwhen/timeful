---
id: TASK-0107
title: Restyle SlideToggle and Available/If needed switch
status: Done
assignee: []
created_date: '2026-08-29 13:02'
updated_date: '2026-08-29 14:41'
labels:
  - frontend
  - styling
dependencies: []
type: enhancement
ordinal: 113000
---

## Description

<!-- SECTION:DESCRIPTION:BEGIN -->
Restyle the shared SlideToggle segmented switch used by AvailabilityTypeToggle (mobile elevated panel next to the Calendar options button; desktop sidebar above it), NewEvent daysOnly, and UpgradeDialog billing.

Requirements:
- Fixed height 36px like the Calendar options v-btn.
- Moving indicator inset 3px inside the box with 2px corner radius so indicator arcs are concentric quarter-circles with the 6px outer box radius (6 - 1 border - 3 gap); no corners visible.
- Indicator styled with tinted background plus colored border from the selected option; remove the glow (boxShadow) entirely.
- Keep the Available (green) / If needed (yellow/orange) color distinction: colored text on the active cell, tinted indicator background, colored indicator border.
- Apply the new default look to all SlideToggle usages; do not change TimeFormatToggle (12h/24h).
- No layout changes: mobile toggle stays right of the Calendar options button on the elevated panel; desktop stays full-width above it.
<!-- SECTION:DESCRIPTION:END -->

## Acceptance Criteria
<!-- AC:BEGIN -->
- [x] #1 SlideToggle renders at a fixed 36px height with a white box background and transparent option cells so no square corners are visible
- [x] #2 The moving indicator is inset 3px inside the toggle box with a 5px corner radius and no corner artifacts
- [x] #3 The indicator uses a solid fill with a colored border and no box shadow is rendered anywhere in the toggle; default fill is tw-bg-green/10
- [x] #4 Active option text is visible above the solid indicator: option cells paint above it and the indicator is pointer-events-none
- [x] #5 Active option text color matches its border color per option (Available green, If needed dark-yellow); inactive text is dark-gray and brightens to black on hover
- [x] #6 AvailabilityTypeToggle uses tw-bg-green/10 and tw-bg-yellow/10 indicator fills with green and dark-yellow borders; the soft-green and pale-yellow tailwind tokens are removed
- [x] #7 NewEvent daysOnly toggle and UpgradeDialog billing toggle render correctly with the new shared default look
- [x] #8 SlideToggle unit tests cover the indicator geometry, colors, no-shadow behavior, and text-above-indicator paint order; AvailabilityTypeToggle has unit coverage for option colors
- [x] #9 npm run lint, typecheck, build, and test:unit pass in frontend
<!-- AC:END -->

## Definition of Done
<!-- DOD:BEGIN -->
- [ ] #1 All acceptance criteria are satisfied
- [ ] #2 All required unit tests pass. Documentation-only changes are exempt unless the user requests unit tests
- [ ] #3 All required e2e tests pass. Documentation-only changes are exempt unless the user requests e2e tests
- [ ] #4 Changed Markdown files are formatted with npm run format:markdown
<!-- DOD:END -->

## Implementation Notes

<!-- SECTION:NOTES:BEGIN -->
Naming: use "indicator" (not "nut") for the moving element.

Geometry model (SlideToggle):
- Box: rounded-md (6px outer), 1px border, fixed h-9 (36px).
- Indicator: inset top/left/bottom 3px, rounded-[2px] = 6 - 1 - 3 so arcs are concentric.
- Indicator width calc(100%/n - 6px) with transform translateX(calc(i*100% + i*6px)) keeps 3px edge gaps and 6px gap between indicator positions; translateX % stays relative to the indicator's own width.
- Option cells drop py-2.5 and stretch to the fixed box height; inactive bg-off-white still fills the box.
- indicatorBgClass option field carries the tint; activeClass becomes text-color-only; borderClass/borderColor color the indicator border; boxShadow/borderStyle removed everywhere.

Plan:
1. SlideToggle.vue markup + script + scoped style changes.
2. AvailabilityTypeToggle.vue option definitions.
3. Update SlideToggle.test.ts geometry assertions; add AvailabilityTypeToggle.test.ts.
4. Required frontend checks; visual sanity on NewEvent/UpgradeDialog defaults.

Design revision after first implementation pass (user feedback on rendered screenshots): container background is now white (like TimeFormatToggle) with transparent cells, which removes the visible square corners of the inactive cell over the rounded box; indicator radius raised from concentric 2px to 5px (same curvature as the box inner corner) per user preference; indicator fill is now solid legend colors (#88cfab green, #FFE8B8 yellow via tw-bg-yellow token) with green/orange borders kept; active option text is black for both options (supersedes earlier idea of matching the If needed text to its outline color); inactive text is dark-gray with hover to black like the 12h/24h toggle.

Round 2-3 design refinements (user-confirmed): container gets white background with transparent cells (removes visible square corners on the inactive option); indicator radius increased to 5px (same curvature as box inner corner, no longer strictly concentric); indicator fills are solid legend-family colors via new tailwind tokens soft-green #88CFAB and pale-yellow #FFF6E0; active text color matches border per option (Available green #00994C, If needed dark-yellow #997700 replacing orange); inactive text dark-gray with hover to black like TimeFormatToggle; TimeFormatToggle itself unchanged.

Revision round 3 (user feedback on rendered toggle): (1) Active option text was invisible - root cause is paint order, not color: the indicator is absolutely positioned with a now-solid fill and paints over the non-positioned option cell text; fix by making option cells tw-relative (they are later in DOM so they paint above the indicator) and adding tw-pointer-events-none to the indicator. (2) Replace solid soft-green/pale-yellow indicator fills with opacity tints of the existing legend colors: Available = tw-bg-green/10, If needed = tw-bg-yellow/10; SlideToggle default fill becomes tw-bg-green/10. (3) Remove the now-unused tailwind tokens soft-green and pale-yellow. Colored active text (green/dark-yellow) is kept and becomes visible again with the paint-order fix.
<!-- SECTION:NOTES:END -->

## Final Summary

<!-- SECTION:FINAL_SUMMARY:BEGIN -->
Restyled the shared SlideToggle segmented switch and all of its usages (AvailabilityTypeToggle on the mobile elevated panel and desktop sidebar, NewEvent daysOnly toggle, UpgradeDialog billing toggle). TimeFormatToggle (12h/24h) was intentionally left untouched.

Changes:
- frontend/src/components/SlideToggle.vue: fixed 36px (h-9) box with white background and transparent, full-height option cells; glow (boxShadow) and the borderStyle option field removed entirely; moving indicator inset 3px with 5px radius, solid fill via a new indicatorBgClass option field, colored border from borderClass/borderColor, sliding via width calc(100%/n - 6px) and translateX(calc(i*100% + i*6px)); default accent is green text, tw-bg-green/10 fill, green border; inactive text dark-gray with hover to black. Revision round 3: option cells are now tw-relative (slide-toggle__option) so active option text paints above the solid indicator instead of being hidden underneath it, and the indicator is tw-pointer-events-none so clicks always reach the cells.
- frontend/src/components/schedule_overlap/AvailabilityTypeToggle.vue: Available = green text + green border (#00994C) + tw-bg-green/10 fill; If needed = dark-yellow text + dark-yellow border (#997700) + tw-bg-yellow/10 fill (opacity tints of the existing legend colors, per user preference over the earlier solid soft-green/pale-yellow fills).
- frontend/tailwind.config.cjs: removed the now-unused soft-green (#88CFAB) and pale-yellow (#FFF6E0) tokens introduced earlier in this task.
- frontend/src/components/SlideToggle.test.ts: covers the 36px white box geometry, rounded-[5px] inset indicator, no shadow, default green/10 accent, inactive hover, and a new text-above-indicator paint-order test (cells relative, indicator pointer-events-none).
- frontend/src/components/schedule_overlap/AvailabilityTypeToggle.test.ts: covers option texts, green/dark-yellow text-border pairing, green/10 and yellow/10 fills, pointer-events-none indicator, no-shadow, and emitted values.

Verification: npm run lint, typecheck, build, and test:unit all pass (138 files / 959 tests). Visual confirmation of the toggle on mobile and desktop was done interactively by the user during design iterations; no e2e changes were required.
<!-- SECTION:FINAL_SUMMARY:END -->
