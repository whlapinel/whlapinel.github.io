+++
date = '2025-06-06T00:03:40-04:00'
draft = false
title = 'Course Manager'
+++

Course Manager is a web application currently in testing mode that allows teachers to develop curricular materials in a systematic and coherent fashion.

The problem this is meant to address is one that I personally experienced as a teacher. I found the current processes of generating and delivering materials to students on a daily basis to be unnecessarily cumbersome and frustrating. Indeed, this began as a personal tool. At first, it was a command-line-interface static site generator. I would manually push the site to Github pages. Eventually I made the interface graphical using Fyne, and then in December 2024 I switched to creating a web application.

Now the web service hosts users' static sites directly which negates the step of publishing to a separate platform.

Features:

- Quickly create slides from markdown using marp
- Quickly create documents from markdown rendered to html using goldmark
- Site generation for each course using Hugo

[https://course-manager.app](https://course-manager.app)
