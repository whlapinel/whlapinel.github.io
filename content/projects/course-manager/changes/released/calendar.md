+++
title="Improve Course Calendar"
+++
## Release comments

Still needs some cosmetic changes but implemented and deployed. Now calendar is monthly with next/previous month buttons instead of loading entire semester at once. Big win! One important result of this is there isn't as much layout shift when using shift lesson or add occasion functionality.

## Proposed Change

Make the app course and term calendars show monthly view instead of entire semester with scrolling.

## Summary

The static site calendar is much better than the app side in terms of appearance and ease of use. Need to bring the app side up to speed by using Tailwind calendar template like I did on the static side.

## Details

## Justification

## Negative Impact

## Required actions

## Possible actions

## Actions taken so far

- Create new route path with month and year params
- Added to showCourseCalendar route handler to parse these params and send the calendar
- Currently only displays the same information as static side, no interactivity yet possible
- Need to add all functionality currently available in app side calendar
- Need to implement on term calendar as well
- Try to eliminate complexity with interfaces in current setup
-
