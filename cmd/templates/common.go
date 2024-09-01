package templates

import (
	"context"
	"log"
	"os"
	"strings"

	"github.com/a-h/templ"
)

type PageData interface {
}

type Page struct {
	Title     string
	TemplFunc func(PageData) templ.Component
	PageData  interface{}
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
			PageData:  EducationItems,
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
	if page.PageData != nil {
		err = page.TemplFunc(page.PageData).Render(context.Background(), f)
	} else {
		err = page.TemplFunc(nil).Render(context.Background(), f)
	}
	if err != nil {
		log.Fatalf("failed to write output file: %v", err)
	}
	return nil
}

type EducationItem struct {
	Image   string
	Year    string
	School  string
	Degree  string
	Minor   string
	Classes []*ClassItem
}

type ProjectItem struct {
	Image       string
	Title       string
	Subtitle    string
	Description string
	Link        string
}

type ClassItem struct {
	Image       string
	Number      string
	Title       string
	Subtitle    string
	Description string
	Link        string
}

var EducationItems = []*EducationItem{
	{
		Image:  "https://via.placeholder.com/150",
		Year:   "2023 - 2024",
		School: "University of North Carolina at Charlotte",
		Degree: "Master of Science in Information Technology",
		Minor:  "Web Development",
		Classes: []*ClassItem{
			{
				Image:       "https://via.placeholder.com/150",
				Number:      "ITIS 5154",
				Title:       "Applied Machine Learning",
				Subtitle:    "Graduate",
				Description: "This course covers the theory and practice of machine learning, including supervised and unsupervised learning, neural networks, and deep learning.",
				Link:        "https://catalog.uncc.edu/",
			},
			{
				Image:       "https://via.placeholder.com/150",
				Number:      "ITIS 5153",
				Title:       "Applied Artificial Intelligence",
				Subtitle:    "Graduate",
				Description: "This course covers the theory and practice of artificial intelligence, including expert systems, natural language processing, and computer vision.",
				Link:        "https://catalog.uncc.edu/",
			},
			{
				Image:       "https://via.placeholder.com/150",
				Number:      "ITIS 6200",
				Title:       "Principles of Information Security and Privacy",
				Subtitle:    "Graduate",
				Description: "This course covers the theory and practice of information security, including cryptography, network security, and risk management.",
			},
		},
	},
	{
		Image:  "https://via.placeholder.com/150",
		Year:   "2005 - 2009",
		School: "United States Naval Academy",
		Degree: "Bachelor of Science",
		Minor:  "History",
	},
	{
		Image:  "https://via.placeholder.com/150",
		Year:   "2001",
		School: "Northwest School of the Arts",
		Degree: "High School Diploma",
	},
}
