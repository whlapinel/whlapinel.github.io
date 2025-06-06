+++
date = '2025-06-06T00:03:40-04:00'
draft = false
title = 'Change Log'
+++

<!-- Created on 4/6/25 -->

[Pending Development](#pending-development)

[Pending Release](#pending-release)

[Released](#released)

## Pending Decision
<!-- this means the decision has not been made whether to move into development phase -->

## Pending Development

- In web app lesson details interface, add button to open slides in new tab so that presenter view can be accessed. When viewing slides in an iframe the presenter view is blocked due to javascript being contained in iframe. [details](./changes/slides-new-tab.md)
- Change assessments so that instructions are contained in database existing column `assessments.instructions` rather than in file referenced in `assessments.file` [details](./changes/assessment_files.md)

## Pending Release

- Added 4/6/25 DB migration: add course_id and unit_id to assessments, remove NOT NULL constraint from lesson_id [details](./changes/assignments.md)

## Released
