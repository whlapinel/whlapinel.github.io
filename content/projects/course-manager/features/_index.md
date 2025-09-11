+++
date = '2025-06-06T00:03:40-04:00'
draft = false
title = 'Features'
+++

## Pending (keep in priority order)

- Add "copy link" to file viewer so you can get static site links from the app side
- change title page for files especially assessment files if possible
- favicons for static site and server
- Add hx-indicator to everything so user gets confirmation that request is pending
- CourseNode files capabilities:
  - 1/22/25 User can delete node files
  - 1/22/25 User can rename node files
  - 1/22/25 User can move node files
  - 1/22/25 User can duplicate node files
- Automatic site regeneration for every change (??)
- Publish/unpublish for all nodes and assessments
- ~~Password protection for assessment files~~ (postponed indefinitely)
  - Add table for assessment files
    - columns: title, relative path (relative to node files), password
- Add password to assessment file
- Allow user to modify unit sequence
- Calendar should jump to current month
- From lesson details page: next lesson or previous lesson (nice to have)
- Static and manager calendars should show assessments
- Copy course must alter or at least zero out the assessment dates (assigned and due)
- Share course with user (read-only)
- Share course with user (with edit privileges)
- Copy course from user
- ~~Add view standards tab to lesson details~~ (postponed indefinitely)

## Complete

- Add course occasions so that occasions can be specific to a course instead of applying to an entire term
- change title page of static site
- From lesson details: back to calendar
- Markdown file editor (non slides)
- 1/23/25 In calendar, user can add and remove lessons from a date
  - Add: new screen to select lesson from course
  - Remove: remove date from lesson dates
- From lesson details page, user can edit assessment details
- Static site should show assessments for lesson along with assessment details
- 1/22/25 User can view lesson files
- 1/22/25 User can upload lesson files
- 12/30/24 Go command to generate site should include other build commands e.g. tailwind, templ, etc (all the things that the task command does). Should probably also pass in the data instead of Generator using its own connection and fetching data?
- 12/30/24 Create `New Lesson` UI
- 12/30/24 Create `New Unit` UI
- 12/30/24 Create `New Course` interface
- 12/30/24 Create Copy Course to Term interface (this will use the Course.FitToTerm method)
