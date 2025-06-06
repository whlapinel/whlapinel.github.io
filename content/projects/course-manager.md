+++
date = '2025-04-07T11:48:39-04:00'
draft = false
title = 'Course Manager'
+++

Course Manager is a web application currently in testing mode that allows teachers to develop curricular materials in a systematic and coherent fashion.

The problem this is meant to address is one that I personally experienced as a teacher. Indeed, this began as a personal tool. At first, it was a command-line-interface static site generator. I would manually push the site to Github pages. Eventually I made the interface graphical using Fyne, and then in December 2024 I switched to creating a web application.

Now the web service hosts users' static sites directly which negates the step of publishing to a separate platform.

Features:

- Quickly create slides from markdown using marp
- Quickly create documents from markdown rendered to html using goldmark
- Site generation for each course using Hugo

Technologies/Dependencies:

- Go
- Echo (web framework for Go)
- Templ
- Caddy (for static sites, as well as reverse proxy for web app)
- Tailwind
- Hugo
- Marp
- Goldmark
- Task (like Makefiles for Go, this is for managing recurring commands and scripts)

Deployment architecture:

Everything is currently running on a single Virtual Private Server running Ubuntu. Inside this VPS, I am using

- Docker compose with separate containers for
  - Caddy
  - Web app with binary (running web app server) and also Hugo for building static sites
  - Marp server (for fetching html rendering of markdown slides)

[https://course-manager.app](https://course-manager.app)
