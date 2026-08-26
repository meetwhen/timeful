# Timeful Project Context

## Purpose

Timeful helps groups select a meeting time with minimal communication.
A creator shares an event, participants mark their availability in a shared grid, and the group uses the resulting view to choose a time that works.

## Vision

Make planning a meeting fast and low effort for groups of any size.
Timeful makes availability visible and comparable while leaving the meeting-time decision to the people involved.

## Primary Users

- Friends coordinating personal meetings.
- Teams, including student teams and teams within companies.
- Event creators, who configure an event and select a meeting time.
- Participants, who provide their availability.

## Stakeholders

- Event creators and participants, who need an understandable, dependable
  coordination flow.
- Instance operators, who deploy and administer self-hosted Timeful instances.
- Timeful maintainers, who develop the project and operate the supported hosted
  offering.
- Core and third-party plugin developers, who may build future browser
  integrations.

## Product Principles

- Availability should be quick to express in a shared grid.
- Timeful should reduce coordination effort without autonomously choosing a
  meeting time.
- The product supports both a hosted service and self-hosted instances, giving
  users deployment and data-control choices.
- Time-zone-aware scheduling should work for distributed groups.

## Current Scope

- Responsive web experiences for desktop and mobile browsers.
- Creating and coordinating events in Timeful.
- Collecting and comparing participant availability.
- Scheduling a Timeful event in a calendar.
- Operating Timeful through the supported hosted site or a self-hosted instance.

## Out of Scope Today

- Automatically recommending or selecting a meeting time on behalf of a group.
- Importing availability or event information from external calendars.
- Browser-plugin integrations that load availability or event information.

Calendar and browser-plugin import are future product directions, not current product capabilities.

## Success Measures

Timeful succeeds when groups:

- Plan meetings faster.
- Spend less effort coordinating a meeting.

Supporting indicators include fewer coordination messages and reliable completion of availability polls by creators and participants.

## Related Documentation

- [Product requirements](requirements/README.md)
- [Environment and deployment configuration](environments.md)
- [Browser plugin API](../PLUGIN_API_README.md)
- [Deployment guide](../DEPLOYMENT.md)
