package templates

import "github.com/a-h/templ"

type contactPage struct {
}

func NewContactPage() Templifier {
	return &contactPage{}
}

func (p *contactPage) Templify() templ.Component {
	return ContactComponent()
}

func (p *contactPage) Title() string {
	return "Contact"
}

func (p *contactPage) Directory() string {
	return rootDir
}
