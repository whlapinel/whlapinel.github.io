+++
date = '2025-09-17T00:03:40-04:00'
draft = false
title = 'Private Slides'
+++

I would like to have the option to make slides private and there's currently no way to create slides that won't be published to the static site. For example if I want to administer an assessment via slides, I have to use Google Slides or some other service because all slides are published (and there is only one slides per lesson). I need to make sure I understand how slides are created and rendered on both the app and static side first.

The process for slides is as follows:

When slides are created via the app, they are written to .slides.md. The `.` in the file name is to prevent Hugo from including as content and rendering via goldmark. On the Hugo side, the slides are fetched as a resource and added to the page. (see _content.gotmpl)

1. infrastructure/marpclient 
    - this simply takes a url and sends a request via network call to the marp server running in a container, and returns the content as a slice of bytes, stripping some of the js that causes issues.
1. slides service
    - `func (s *SlidesService) GetSlides(nodes ...ports.Node) ([]byte, error)`
    - `func (svc *SlidesService) SlidesContent(nodes ports.Nodes) ([]byte, error)`
    - `func (svc *SlidesService) UpdateSlides(nodes ports.Nodes, content []byte) error`
    - `func (s *SlidesService) dataPathToMarpURL(path string) (string, error)`
    - this writes the markdown sent in on post request to .slides.md
    - this takes the bytes from the client and writes to a file called .slides.html
1. ui


If I make a new "private slides" field and write this to .private_slides.md, I could reproduce the whole system on the app side and not worry about it getting published on the static site. Perhaps adding a "private" boolean to the service methods would be all I need to modify the logic accordingly.

On the client side, I would just need to duplicate the UI and somehow flag it as private so requests to the server are flagged similarly.

The handler would need to check for the private flag as well.

