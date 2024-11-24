package templates

import (
	"fmt"
	"gh_static_portfolio/cmd/domain"
	"log"
	"os"
	"strings"

	"github.com/a-h/templ"
)

const (
	rootDir      = "./docs/"
	aboutDir     = "./docs/about/"
	educationDir = "./docs/about/education/"
	blogDir      = "./docs/blog/"
	coursesDir   = "./docs/courses/"
)

func courseRoutePath(course domain.Course) string {
	return fmt.Sprintf("%s%s", coursesDir, DirName(course))
}

func unitRoutePath(unit domain.Unit, course domain.Course) string {
	return fmt.Sprintf("%s%s%s", coursesDir, DirName(course), DirName(unit))
}

func lessonRoutePath(lesson domain.Lesson, unit domain.Unit, course domain.Course) string {
	return fmt.Sprintf("%s%s%s%s", coursesDir, DirName(course), DirName(unit), DirName(lesson))
}

func filesRoutePath(lesson domain.Lesson, unit domain.Unit, course domain.Course) string {
	return fmt.Sprintf("https://github.com/whlapinel/whlapinel.github.io/tree/main/docs/courses/%s%s%sfiles", DirName(course), DirName(unit), DirName(lesson))
}

func hasImage(path string) bool {
	files, err := os.ReadDir(path)
	if err != nil {
		log.Fatalf("error reading directory:%s", err)
	}
	for _, file := range files {
		if file.Name() == "image.png" {
			return true
		}
	}
	return false

}

func DirectoriesClearList() []string {
	return []string{
		rootDir,
		aboutDir,
		educationDir,
		blogDir,
		coursesDir,
	}
}

type Templifier interface {
	Templify() templ.Component
	Title() string
	Directory() string
}

type Titler interface {
	Title() string
}

type page struct {
	title     string
	directory string
	component templ.Component
}

func (p *page) Templify() templ.Component {
	return p.component
}

func (p *page) Title() string {
	return p.title
}

func (p *page) Directory() string {
	return p.directory
}

func NewHomePage() Templifier {
	return &page{
		title:     "Home",
		directory: rootDir,
		component: HomeComponent(),
	}
}

func NewContactPage() Templifier {
	return &page{
		title:     "Contact",
		directory: rootDir,
		component: ContactComponent(),
	}
}

func NewAboutPage() Templifier {
	return &page{
		title:     "About",
		directory: rootDir,
		component: AboutComponent(),
	}
}

func NewEducationPage(items []*domain.EducationItem) Templifier {
	return &page{
		title:     "Education",
		directory: aboutDir,
		component: EducationListComponent(items),
	}
}

func NewClassesPage(item *domain.EducationItem) Templifier {
	return &page{
		title:     item.School,
		directory: educationDir,
		component: ClassListComponent(item),
	}
}

func NewWorkHistoryPage(items []*domain.WorkHistoryItem) Templifier {
	return &page{
		title:     "Work History",
		directory: aboutDir,
		component: WorkHistoryComponent(items),
	}
}

func NewProjectsPage(items []*domain.ProjectItem) Templifier {
	return &page{
		title:     "Projects",
		directory: aboutDir,
		component: ProjectListComponent(items),
	}
}

func NewSkillsPage(items []*domain.SkillItem) Templifier {
	return &page{
		title:     "Skills",
		directory: aboutDir,
		component: SkillsListComponent(items),
	}
}

func NewBlogsListPage(items []*domain.Blog) Templifier {
	return &page{
		title:     "Blog",
		directory: rootDir,
		component: BlogsListComponent(items),
	}
}

func NewBlogPage(item *domain.Blog) Templifier {
	if item == nil {
		return &page{
			directory: blogDir,
			component: BlogComponent(nil),
		}
	}
	return &page{
		title:     item.Title(),
		directory: blogDir,
		component: BlogComponent(item),
	}
}

func NewPersonalPage() Templifier {
	return &page{
		title:     "Personal",
		directory: aboutDir,
		component: PersonalComponent(),
	}
}

func NewCoursesListPage(courses []domain.Course) Templifier {
	return &page{
		title:     "Courses",
		directory: rootDir,
		component: CoursesListComponent(courses),
	}

}

func NewCoursePage(course domain.Course) Templifier {
	return &page{
		title:     course.Title(),
		directory: courseRoutePath(course),
		component: CourseComponent(course),
	}
}

func NewUnitPage(unit domain.Unit, course domain.Course) Templifier {
	return &page{
		title:     unit.Title(),
		directory: unitRoutePath(unit, course),
		component: UnitComponent(unit, course),
	}
}

func NewLessonPage(lesson domain.Lesson, unit domain.Unit, course domain.Course) Templifier {
	return &page{
		title:     lesson.Title(),
		directory: lessonRoutePath(lesson, unit, course),
		component: LessonComponent(lesson, unit, course),
	}
}

// same as FileName() but with "/" at the end instead of the .html extension
func DirName(t Titler) string {
	return strings.ReplaceAll(strings.ToLower(t.Title()), " ", "-") + "/"
}

// replaces spaces with dashes, makes lowercase, and adds .html file extension
func FileName(t Titler) string {
	if t.Title() == "Home" {
		return "index.html"
	}
	return strings.ReplaceAll(strings.ToLower(t.Title()), " ", "-") + ".html"
}

func RemoveDocsFromPath(directory string) string {
	return strings.ReplaceAll(directory, "./docs", "")
}

func rootPages() []Templifier {
	return []Templifier{
		NewHomePage(),
		NewAboutPage(),
		NewBlogsListPage(nil),
		NewContactPage(),
		NewCoursesListPage(nil),
	}
}
