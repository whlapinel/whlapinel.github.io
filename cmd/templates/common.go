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
	log.Println("len(AboutSections): ", len(AboutSections))
	for _, page := range Pages {
		RenderPage(page)
	}
	for _, section := range AboutSections {
		RenderAboutSection(section)
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

type WorkHistoryItem struct {
	Image       string
	Year        string
	Company     string
	Position    string
	Description string
}

type SkillItem struct {
	Image       string
	Title       string
	Subtitle    string
	Description string
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

var WorkHistoryItems = []*WorkHistoryItem{
	{
		Image:       "https://via.placeholder.com/150",
		Year:        "2021 - Present",
		Company:     "Company 1",
		Position:    "Position 1",
		Description: "This is a description of my work at Company 1.",
	},
	{
		Image:       "https://via.placeholder.com/150",
		Year:        "2019 - 2021",
		Company:     "Company 2",
		Position:    "Position 2",
		Description: "This is a description of my work at Company 2.",
	},
}

var ProjectItems = []*ProjectItem{
	{
		Image:       "https://via.placeholder.com/150",
		Title:       "Project 1",
		Subtitle:    "Subtitle 1",
		Description: "This is a description of project 1.",
		Link:        "",
	},
}

var SkillItems = []*SkillItem{
	{
		Image:       "https://via.placeholder.com/150",
		Title:       "Skill 1",
		Subtitle:    "Subtitle 1",
		Description: "This is a description of skill 1.",
	},
}
