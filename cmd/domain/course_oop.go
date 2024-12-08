package domain

func NewCourse(title string, descr string, units []Unit, termName string) CourseOOP {
	return CourseOOP{Title: title, Description: descr, Units: units, TermName: termName}

}

type CourseRepository interface {
	GetAll() ([]CourseOOP, error)
}

// Courses I teach. this is the OOP version of CourseInstance. Bad wording I know.
type CourseOOP struct {
	ID          int
	Instance    bool
	Title       string
	Description string
	Units       []Unit
	TermName    string
}

type CourseType int

type CourseInstanceOOP struct {
}

func (c CourseOOP) GetTitle() string {
	return c.Title
}
