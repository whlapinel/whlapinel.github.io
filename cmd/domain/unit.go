package domain

func NewUnit(num int, title string, descr string, lessons []Lesson) Unit {
	return Unit{num, title, descr, lessons}

}

func (u Unit) GetTitle() string {
	return u.Title
}

type Unit struct {
	Number      int
	Title       string
	Description string
	Lessons     []Lesson
}
