package templates

import "github.com/a-h/templ"

type homePage struct {
}

func NewHomePage() Templifier {
	return &homePage{}
}

func (p *homePage) Templify() templ.Component {
	return HomeComponent()
}

func (p *homePage) Title() string {
	return "Home"
}

func (p *homePage) Directory() string {
	return string(rootDir)
}
