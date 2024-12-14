package domain

func NewCourse(title string, descr string, units []Unit, termName string) Course {
	return Course{Name: title, Description: descr, Units: units, TermName: termName}

}

type CourseRepo interface {
	Repository[Course]
	GetTemplates() ([]*Course, error)
	GetInstances() ([]*Course, error)
}

// Courses I teach. this is the OOP version of CourseInstance. Bad wording I know.
type Course struct {
	ID          int
	TemplateID  int
	Name        string
	Description string
	Units       []Unit
	TermID      int
	TermName    string
}

type CourseType int

type CourseInstanceOOP struct {
}

func (c Course) GetTitle() string {
	return c.Name
}

func (c *Course) AddUnit(unit Unit) {
	c.Units = append(c.Units, unit)
}
