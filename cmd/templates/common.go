package templates

import (
	"gh_static_portfolio/cmd/domain"
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

func DirectoryList() []string {
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
		directory: coursesDir,
		component: CourseComponent(course),
	}
}

func NewPersonalPage() Templifier {
	return &page{
		title:     "Personal",
		directory: aboutDir,
		component: PersonalComponent(),
	}
}

func FileName(t Titler) string {
	if t.Title() == "Home" {
		return "index.html"
	}
	return strings.ReplaceAll(strings.ToLower(t.Title()), " ", "-") + ".html"
}

func RoutePath(directory string) string {
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
