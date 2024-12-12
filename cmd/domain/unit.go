package domain

func NewUnit(num int, title string, descr string, lessons []Lesson) Unit {
	return Unit{Number: num, Name: title, Description: descr, Lessons: lessons}

}

func (u Unit) GetTitle() string {
	return u.Name
}

// Always associated with a particular course
type Unit struct {
	ID          int
	CourseID    int
	Number      int
	Name        string
	Description string
	Lessons     []Lesson
}
