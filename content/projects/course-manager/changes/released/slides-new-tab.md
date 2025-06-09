
# Proposed Change: Add new button in web application lesson details page to view slides in new tab

## Status

Complete!

## Summary

Presenter view is blocked, presumably due to slides being inside an iframe. This means speaker notes cannot be viewed.

Viewing slides in a separate tab will allow presenter view to be accessed.

## Justification

Speaker notes are an essential component of instruction for many teachers, even though I don't often use them I do intend to in the future.

## Negative Impact

A small amount of time

## Required actions

- Add button, route handler should just send as a static file. File should be guaranteed already rendered, but I need to make sure about that.
- Maybe use existing getSlides route handler with query param like ?new-tab=true or something like that?
