package main

import (
	"context"
	"fmt"
	"gh_static_portfolio/cmd/data"
	"gh_static_portfolio/cmd/templates"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"
)

func main() {
	directories := templates.DirectoriesClearList()
	for _, directory := range directories {
		ClearHTMLFiles(directory)
	}
	// Generate education page using education data
	// educationRepo := data.NewEducationRepo()
	// educationItems, err := educationRepo.GetAll()
	// if err != nil {
	// 	log.Fatalf("failed to get education items: %v", err)
	// }
	// educationPage := templates.NewEducationPage(educationItems)
	// err = RenderPage(educationPage)
	// if err != nil {
	// 	log.Fatalf("failed to render pages: %v", err)
	// }
	// // Generate "classes attended" list for education
	// for _, item := range educationItems {
	// 	if len(item.Classes) > 0 {
	// 		educationItemPage := templates.NewClassesPage(item)
	// 		err = RenderPage(educationItemPage)
	// 		if err != nil {
	// 			log.Fatalf("failed to render pages: %v", err)
	// 		}
	// 	}
	// }
	// // Generate skills page using skills data
	// skillsRepo := data.NewSkillRepo()
	// skillItems, err := skillsRepo.GetAll()
	// if err != nil {
	// 	log.Fatalf("failed to get skill items: %v", err)
	// }
	// skillsPage := templates.NewSkillsPage(skillItems)
	// err = RenderPage(skillsPage)
	// if err != nil {
	// 	log.Fatalf("failed to render pages: %v", err)
	// }
	// // Generate work history page using work history data
	// workHistoryRepo := data.NewWorkHistoryRepo()
	// workHistoryItems, err := workHistoryRepo.GetAll()
	// if err != nil {
	// 	log.Fatalf("failed to get work history items: %v", err)
	// }
	// workHistoryPage := templates.NewWorkHistoryPage(workHistoryItems)
	// err = RenderPage(workHistoryPage)
	// if err != nil {
	// 	log.Fatalf("failed to render pages: %v", err)
	// }
	// // Generate projects page using projects data
	// projectsRepo := data.NewProjectRepo()
	// projectItems, err := projectsRepo.GetAll()
	// if err != nil {
	// 	log.Fatalf("failed to get project items: %v", err)
	// }
	// projectsPage := templates.NewProjectsPage(projectItems)
	// err = RenderPage(projectsPage)
	// if err != nil {
	// 	log.Fatalf("failed to render pages: %v", err)
	// }
	// // Generate blogs list using blogs data
	// blogRepo := data.NewBlogRepo()
	// blogItems, err := blogRepo.GetAll()
	// if err != nil {
	// 	log.Fatalf("failed to get blog items: %v", err)
	// }
	// blogsPage := templates.NewBlogsListPage(blogItems)
	// err = RenderPage(blogsPage)
	// if err != nil {
	// 	log.Fatalf("failed to render pages: %v", err)
	// }
	// err = RenderPage(blogsPage)
	// if err != nil {
	// 	log.Fatalf("failed to render pages: %v", err)
	// }
	// // Generate individual blogs pages
	// for _, blog := range blogItems {
	// 	blogPage := templates.NewBlogPage(blog)
	// 	err = RenderPage(blogPage)
	// 	if err != nil {
	// 		log.Fatalf("failed to render pages: %v", err)
	// 	}
	// }
	// Generate home page
	homePage := templates.NewHomePage()
	err := RenderPage(homePage)
	if err != nil {
		log.Fatalf("failed to render pages: %v", err)
	}
	// Generate about page
	aboutPage := templates.NewAboutPage()
	err = RenderPage(aboutPage)
	if err != nil {
		log.Fatalf("failed to render pages: %v", err)
	}
	// Generate contact page
	contactPage := templates.NewContactPage()
	err = RenderPage(contactPage)
	if err != nil {
		log.Fatalf("failed to render pages: %v", err)
	}
	// // Generate personal page
	// personalPage := templates.NewPersonalPage()
	// err = RenderPage(personalPage)
	// if err != nil {
	// 	log.Fatalf("failed to render pages: %v", err)
	// }

	// Generate "courses I teach" list page
	courseRepo := data.NewCoursesRepo()
	courses, err := courseRepo.ReadFromCSV()
	if err != nil {
		log.Fatalf("Error getting courses: %v", err)
	}
	coursesPage := templates.NewCoursesListPage(courses)
	err = RenderPage(coursesPage)
	if err != nil {
		log.Fatalf("failed to render pages: %v", err)
	}
	// Generate page for each course
	for _, course := range courses {
		log.Println("Site generator main() course: Name: ", course.Title)
		coursePage := templates.NewCoursePage(course)
		err = RenderPage(coursePage)
		if err != nil {
			log.Fatalf("failed to render pages: %v", err)
		}
		// Generate calendar page for each course
		calendarPage := templates.NewCourseCalendarPage(*course)
		err = RenderPage(calendarPage)
		if err != nil {
			log.Fatalf("failed to render pages: %v", err)
		}
		// Generate page for each unit
		for _, unit := range course.Units {
			log.Println("looping through units in main():", unit.Title)
			unitPage := templates.NewUnitPage(unit, *course)
			err = RenderPage(unitPage)
			if err != nil {
				log.Fatalf("failed to render pages: %v", err)
			}
			// Generate page for each lesson
			for _, lesson := range unit.Lessons {
				lessonPage := templates.NewLessonPage(lesson, unit, *course)
				GenerateSlides(lessonPage.Directory())
				err = RenderPage(lessonPage)
				if err != nil {
					log.Fatalf("failed to render pages: %v", err)
				}
			}
		}
	}
}

func ClearHTMLFiles(directory string) {
	files, err := os.ReadDir(directory)
	if err != nil {
		return
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
	log.Println("RenderPage: ", t.GetTitle())
	err := os.MkdirAll(t.Directory(), os.ModePerm)
	if err != nil {
		log.Fatalf("failed to create directory: %v", err)
	}
	// ClearHTMLFiles(t.Directory())
	f, err := os.Create(t.Directory() + templates.FileName(t))
	if err != nil {
		log.Fatalf("failed to create output file: %v", err)
	}
	// added 11/24: remove all html files before writing. Probably not necessary, so commenting out for now
	// files, err := os.ReadDir(t.Directory())
	// if err != nil {
	// 	return err
	// }
	// for _, file := range files {
	// 	if strings.HasSuffix(file.Name(), ".html") {
	// 		err := os.Remove(t.Directory() + file.Name())
	// 		if err != nil {
	// 			log.Fatalf("failed to delete file: %v", err)
	// 		}
	// 	}
	// }
	err = t.Templify().Render(context.Background(), f)
	if err != nil {
		log.Fatalf("failed to write output file: %v", err)
	}
	return nil
}

func GenerateSlides(dir string) {
	// File paths
	markdownFile := "slides.md"
	htmlFile := "slides.html"
	markdownPath := path.Join(dir, markdownFile)
	htmlPath := path.Join(dir, htmlFile)

	// Get file information
	mdInfo, err := os.Stat(markdownPath)
	if err != nil {
		log.Printf("Error: Could not access %s: %v\n", markdownPath, err)
		return
	}

	htmlInfo, err := os.Stat(htmlPath)
	if os.IsNotExist(err) {
		// If HTML file doesn't exist, regenerate it
		log.Println("HTML file does not exist. Generating...")
		regenerateHTML(markdownPath, htmlPath)
		return
	} else if err != nil {
		log.Printf("Error: Could not access %s: %v\n", htmlPath, err)
		return
	}

	// Compare modification times
	mdModTime := mdInfo.ModTime()
	htmlModTime := htmlInfo.ModTime()

	if mdModTime.After(htmlModTime) {
		log.Println("Markdown file is newer. Regenerating HTML...")
		regenerateHTML(markdownPath, htmlPath)
	} else {
		log.Println("No need to regenerate. HTML is up-to-date.")
	}
}

// Regenerate HTML from the Markdown file
func regenerateHTML(markdownFile, htmlFile string) {
	cmd := exec.Command("marp", markdownFile, "-o", htmlFile)
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error: Failed to regenerate HTML: %v\n", err)
		return
	}
	fmt.Printf("HTML file %s successfully regenerated from %s\n", htmlFile, markdownFile)
}
