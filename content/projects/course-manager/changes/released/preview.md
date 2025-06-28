+++
title="Replace App-Side Markdown Rendering With Static Site Preview"
+++

## Proposed Change

- Replace app-side goldmark rendering with simpler static site preview window

## Summary

- Currently the user can view markdown files rendered to HTML with goldmark by the application itself. But there are some problems with this. Namely, I have to change rendering in 2 different places or it might not appear the same. Second and more importantly, if the markdown has shortcodes in it, they won't work and the code will be shown as code.

## Details

- Remove the goldmark renderer
- Add a site preview link

## Justification

## Negative Impact

- Small amount of work

## Required actions

Decide how to ensure build shown to user is not stale
- by triggering build on preview request?

## Possible actions

## Actions taken so far

- Commented out the app-side rendering UI (everything is still there on the backend however)
- Added a site preview interface in the markdown service (app layer service). This simply provids the static site url for a given file.
- Cleaned up the code in this service so that the front matter url setter uses the same code as the preview URL provider

## Final notes

- Initially I was using an iframe and anchor element which allowed a nice preview of the entire site but in order to specifically show the document I changed it to an htmx button and a div. For now I'm just selecting the page out of the response and putting that in the preview div. It works. I just need to make sure I keep markdown styles in sync between the static site and the app.
- I had some issues on the client-side with CORS after changing the preview link to an HTMX button since HTMX sends Ajax requests. Had to modify the caddyfile for both dev and prod to respond to OPTIONS method requests with allowed origins and headers.
- Added generate site button to node details page. Component simply checks if URL field is not empty before rendering button. Handler provides url for course, unit and lesson details.
- One issue that needs to be resolved somehow: links within a document won't work since they're not displayed in an iframe
  - I either need to go back to simply showing the site in an iframe, or somehow rendering the content with htmx in the iframe, or just show a warning that links won't work (not ideal)
