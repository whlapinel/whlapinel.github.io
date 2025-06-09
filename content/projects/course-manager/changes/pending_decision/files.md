# Proposed Change: Mount files to content folder and use hugo to generate pages rather than doing so manually in _content.gotmpl

## Status

Pending decision

## Summary

Currently the content adapter adds pages manually. It works fine, but the problem is there is no way of setting and getting front matter fields such as title, etc.

## Details

Since the file itself doesn't store metadata the only place I could put that would be in the database. I really don't want to deal with having file information in the database because that seems like a giant can of worms. Instead, I could add form fields such as title to the file creation/upload interface for markdown files, and then write that information to the file. I would also need to hide the front matter when rendering the markdown as well as allowing editing of the markdown

## Justification

Having a way to include file title and other metadata would be a nice thing to have, and it would allow me to take more advantage of Hugo features in the future.

## Negative Impact

Substantial amount of work.

## Required actions

- Mount user data directory to content directory and set
  - includeFiles='**/files/*.md'
  - excludeFiles='**/slides.md'
- Remove page creation (.AddPage) portion of _content.gotmpl as pages will now be created by Hugo automatically from content directory
- Modify view markdown handler in app so that TOML front matter portion is not shown to user
- Modify edit markdown handler in app so that TOML front matter portion is not shown to user
- Form for editing file metadata should be included when editing markdown file (just start with title)
- Update, save file handlers should be modified so that metadata captured from the form is written to file as TOML
