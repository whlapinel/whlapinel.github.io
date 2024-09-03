package main

import (
	"context"
	"gh_static_portfolio/cmd/data"
	"gh_static_portfolio/cmd/templates"
	"log"
	"os"
	"strings"
)

func main() {
	directories := templates.DirectoryList()
	for _, directory := range directories {
		ClearHTMLFiles(directory)
	}
	educationRepo := data.NewEducationRepo()
	educationItems, err := educationRepo.GetAll()
	if err != nil {
		log.Fatalf("failed to get education items: %v", err)
	}
	educationPage := templates.NewEducationPage(educationItems)
	err = RenderPage(educationPage)
	if err != nil {
		log.Fatalf("failed to render pages: %v", err)
	}
	skillsRepo := data.NewSkillRepo()
	skillItems, err := skillsRepo.GetAll()
	if err != nil {
		log.Fatalf("failed to get skill items: %v", err)
	}
	skillsPage := templates.NewSkillsPage(skillItems)
	err = RenderPage(skillsPage)
	if err != nil {
		log.Fatalf("failed to render pages: %v", err)
	}
	workHistoryRepo := data.NewWorkHistoryRepo()
	workHistoryItems, err := workHistoryRepo.GetAll()
	if err != nil {
		log.Fatalf("failed to get work history items: %v", err)
	}
	workHistoryPage := templates.NewWorkHistoryPage(workHistoryItems)
	err = RenderPage(workHistoryPage)
	if err != nil {
		log.Fatalf("failed to render pages: %v", err)
	}
	projectsRepo := data.NewProjectRepo()
	projectItems, err := projectsRepo.GetAll()
	if err != nil {
		log.Fatalf("failed to get project items: %v", err)
	}
	projectsPage := templates.NewProjectsPage(projectItems)
	err = RenderPage(projectsPage)
	if err != nil {
		log.Fatalf("failed to render pages: %v", err)
	}
	blogRepo := data.NewBlogRepo()
	blogItems, err := blogRepo.GetAll()
	if err != nil {
		log.Fatalf("failed to get blog items: %v", err)
	}
	blogsPage := templates.NewBlogsPage(blogItems)
	err = RenderPage(blogsPage)
	if err != nil {
		log.Fatalf("failed to render pages: %v", err)
	}
	homePage := templates.NewHomePage()
	err = RenderPage(homePage)
	if err != nil {
		log.Fatalf("failed to render pages: %v", err)
	}
	aboutPage := templates.NewAboutPage()
	err = RenderPage(aboutPage)
	if err != nil {
		log.Fatalf("failed to render pages: %v", err)
	}
	contactPage := templates.NewContactPage()
	err = RenderPage(contactPage)
	if err != nil {
		log.Fatalf("failed to render pages: %v", err)
	}
	personalPage := templates.NewPersonalPage()
	err = RenderPage(personalPage)
	if err != nil {
		log.Fatalf("failed to render pages: %v", err)
	}
}

func ClearHTMLFiles(directory string) {
	files, err := os.ReadDir(directory)
	if err != nil {
		log.Fatalf("failed to read directory: %v", err)
	}
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".html") {
			err := os.Remove(directory + file.Name())
			if err != nil {
				log.Fatalf("failed to delete file: %v", err)
			}
		}
	}
}

func RenderPage(t templates.Templifier) error {
	err := os.MkdirAll(t.Directory(), os.ModePerm)
	if err != nil {
		log.Fatalf("failed to create directory: %v", err)
	}
	f, err := os.Create(t.Directory() + templates.FileName(t))
	if err != nil {
		log.Fatalf("failed to create output file: %v", err)
	}
	err = t.Templify().Render(context.Background(), f)
	if err != nil {
		log.Fatalf("failed to write output file: %v", err)
	}
	return nil
}
