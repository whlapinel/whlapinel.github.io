+++
title="Leverage Hugo features more for markdown files"
+++

## Proposed Change

- Mount markdown files to content folder and use hugo to generate pages rather than doing so manually in _content.gotmpl
- Use front matter in markdown files either programatically and/or entered by user via editor

## Summary

Currently the content adapter adds pages manually. It works fine, but the problem is there is no way of setting and getting front matter fields such as title, etc.

## Details

Since the file itself doesn't store metadata the only place I could put that would be in the database. I really don't want to deal with having file information in the database because that seems like a giant can of worms.

- Instead, I could either add form fields such as title to the file creation/upload interface for markdown files, and then write that information to the file.
- There are some fields that I want to set programmatically such as URL, while other fields that I could allow the user to set e.g. title.
- I would also need to hide the front matter when rendering the markdown as well as allowing editing of the markdown (done, by adding extension to goldmark)

## Justification

Having a way to include file title and other metadata would be a nice thing to have, and it would allow me to take more advantage of Hugo features in the future. Would be a nice way for user to customize their files.

- Just realized, now the user can control publishing of documents using front matter (!!) 😲😲😲 which means I don't have to write a custom way of handling that.
  - Still need to handle this issue for non-markdown files however.

## Negative Impact

Substantial amount of work.

## Required actions

- [X] Mount user data directory to content directory and set
  - includeFiles='**/files/*.md'
  - excludeFiles='**/slides.md' -- front matter here should be handled separately if at all. Currently it is yaml entered manually.
- [X] Remove page creation (.AddPage) portion of _content.gotmpl as pages will now be created by Hugo automatically from content directory
- [X] Modify goldmark renderer so that front matter portion is not shown to user
- ~~[] Modify edit markdown handler in app so that TOML front matter portion is not shown to user~~
- ~~Form for editing file metadata should be included when editing markdown file (just start with title)~~
- ~~Update, save file handlers should be modified so that metadata captured from the form is written to file as TOML~~
- [X] Update and create markdown file handlers should be modified so that front matter is validated and set for certain fields e.g. URL, page type ("standalone")
- ~~[] Deferrable to future: add .meta for each non-markdown file to allow control of publishing (needs modification of infrastructure/hugo code as well as _content.gotmpl)~~
  - For now we can just auto-publish all files and instead of listing in files page, only serve via links in markdown files. So they'll be published
  - [X] As a solution, I am thinking of the following, but this should be a separate change log entry:
    - for each standalone page, in front matter there could be a Params.files, a list of files that should be linked in the page
    - in the standalone page layout, range over .Params.files and for each file call ~~.GetResource or resources.GetMatch~~ or whatever the right method is (see below)
    - Ok this is implemented now. Standalone pages can link to files using a shortcode like this: [Example]{{< /*parentresource "my_file.txt"*/>}} which gets the resource from parent page. I should probably auto-insert an html comment with an example of this for all markdown files
    - A question remaining is that I am not sure if resources are added automatically or only if they are linked in a page. A little experimentation will answer this quickly. If they are only added when linked then all my problems are solved here, no need for creating metadata for each file or anything like that.

## Possible actions

- [X] in go code and/or in _content.gotmpl, change files page (this is a list of both markdown files (rendered as pages) as well as resources) from kind page to kind section. This allows me to use range .RegularPages perhaps?

## Actions taken so far

- Removed the filePage creation portion of _content.gotmpl
- Mounted UserDataDir i.e. internal/data/users/<user-id>/terms... to the content directory
  - This results in the pages being built automatically by Hugo according to front matter rather than according to _content.gotmpl
- In post file handler (currently sent only by file upload there) or in the service used by that handler, I've added some code to write front matter.
  - Title can be entered by user
  - URL should be handled programmatically in file service
