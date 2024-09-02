package templates

import (
	"context"
	"gh_static_portfolio/cmd/data"
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
			TemplFunc: BlogsListComponent,
		},
	}
	AboutSections = []*AboutSection{
		{
			Title:     "Skills",
			TemplFunc: SkillsListComponent,
		},
		{
			Title:     "Work History",
			TemplFunc: WorkHistoryComponent,
		},
		{
			Title:     "Education",
			TemplFunc: EducationListComponent,
		},
		{
			Title:     "Projects",
			TemplFunc: ProjectListComponent,
		},
		{
			Title:     "Personal",
			TemplFunc: PersonalComponent,
		},
	}
	ClearHTMLFiles()
	for _, page := range Pages {
		RenderPage(page)
	}
	for _, section := range AboutSections {
		RenderAboutSection(section)
	}
	err := RenderBlogs()
	if err != nil {
		log.Fatalf("failed to render blogs: %v", err)
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

func RenderAboutSection(s *AboutSection) error {

	f, err := os.Create("./docs/about/" + s.FileName())
	if err != nil {
		log.Fatalf("failed to create output file: %v", err)
	}
	err = s.TemplFunc().Render(context.Background(), f)
	if err != nil {
		log.Fatalf("failed to write output file: %v", err)
	}
	return nil
}

func RenderPage(page Page) error {
	f, err := os.Create("./docs/" + page.FileName())
	if err != nil {
		log.Fatalf("failed to create output file: %v", err)
	}
	if page.Title == "About" {
		log.Println("About page")
		log.Println("len(AboutSections): ", len(AboutSections))
	}
	err = page.TemplFunc().Render(context.Background(), f)
	if err != nil {
		log.Fatalf("failed to write output file: %v", err)
	}
	return nil
}

func RenderBlogs() error {
	for _, blog := range data.Blogs {
		f, err := os.Create("./docs/blogs/" + strings.ReplaceAll(strings.ToLower(blog.Title), " ", "-") + ".html")
		if err != nil {
			log.Fatalf("failed to create output file: %v", err)
		}
		err = BlogComponent(&blog).Render(context.Background(), f)
		if err != nil {
			log.Fatalf("failed to write output file: %v", err)
		}
	}
	return nil
}
