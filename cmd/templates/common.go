package templates

import (
	"context"
	"log"
	"os"
	"strings"

	"github.com/a-h/templ"
)

type Page struct {
	Title     string
	TemplFunc func() templ.Component
}

func (p *Page) FileName() string {
	if p.Title != "Home" {
		return strings.ToLower(p.Title) + ".html"
	}
	return "index.html"
}

var Pages []Page

func RenderPages() error {
	Pages = []Page{
		{
			Title:     "Home",
			TemplFunc: HomeComponent,
		},
		{
			Title:     "About",
			TemplFunc: AboutComponent,
		},
		{
			Title:     "Contact",
			TemplFunc: ContactComponent,
		},
		{
			Title:     "Blog",
			TemplFunc: BlogComponent,
		},
	}
	ClearHTMLFiles()
	for _, page := range Pages {
		RenderPage(page)
	}
	return nil
}

func ClearHTMLFiles() {
	files, err := os.ReadDir("./docs")
	if err != nil {
		log.Fatalf("failed to read directory: %v", err)
	}
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".html") {
			err := os.Remove("./docs/" + file.Name())
			if err != nil {
				log.Fatalf("failed to delete file: %v", err)
			}
		}
	}
}

func RenderPage(page Page) error {
	f, err := os.Create("./docs/" + page.FileName())
	if err != nil {
		log.Fatalf("failed to create output file: %v", err)
	}
	err = page.TemplFunc().Render(context.Background(), f)
	if err != nil {
		log.Fatalf("failed to write output file: %v", err)
	}
	return nil
}
