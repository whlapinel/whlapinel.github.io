# Proposed Change: Full re-design of code structure

## Status

Completed!

## Summary

Rewrite project to conform with feature-based + clean architecture.

## Justification
Things are starting to feel unweildy and I would like to conform with modern architectural practices to the extent feasible. I would like to maximize the maintainability of my codebase because I have come to think that perhaps it is much much harder to add features or fix bugs than it probably should be. I am really starting to dislike the hacky way that I wrote things in the beginning and want to make it all nice and neat.

## Negative Impact
This will be a lot of work.

## Required actions

- Decide on design.

## Additional info:

 This is just a sketch of the new layout:

```text
/cmd/
  server/
    main.go                  ← calls app.New() and starts the server

/internal/
  /app/                      ← 📌 This is your **application layer**
    container.go             ← builds and wires services, repos, handlers
    router.go                ← mounts handler routes (handlers come from features)
    dto/                     ← composed types (e.g., TermWithCourses)
      term_with_courses.go
    orchestration/           ← application-specific workflows across domains
      schedule_term.go       ← e.g., create term + calendar + courses
      clone_course.go        ← e.g., duplicate course structure across terms

  /presentation/             ← shared templates and rendering helpers
    layout.templ
    base_components.templ
    shared_helpers.go

  /shared/                   ← domain-agnostic utilities
    node.go
    image.go
    middleware.go
    auth.go

  /infrastructure/           ← adapters and system integrations
    db.go
    postgres/
      course_repo.go
      term_repo.go

  /features/
    /course/
      model.go
      handler.go
      service.go
      repository.go
      repo_postgres.go

    /term/
      model.go
      handler.go
      service.go
      repository.go
      repo_postgres.go
```
