package domain

type CourseSOARepository interface {
	GetAll() ([]*CourseSOA, error)
	Save(*CourseSOA) error
}

type CourseSOA struct {
	ID           int
	Name         string
	Course       Course
	Day          []int // day of instruction number e.g. 1, 2, ...87
	UnitNum      []int // unit number e.g. 1, 2, 3
	UnitSequence []int
	UnitDescr    []string
	LessonNum    []int // lesson number e.g. 1, 2, 3
	LessonDescr  []string
	StandardNum  []int
	StdDescr     []string
}

func (c CourseSOA) GetTitle() string {
	return c.Name
}

func (curric CourseSOA) Truncate(end int) CourseSOA {
	curric.Day = curric.Day[:end]
	curric.UnitNum = curric.UnitNum[:end]
	curric.UnitDescr = curric.UnitDescr[:end]
	curric.LessonNum = curric.LessonNum[:end]
	curric.LessonDescr = curric.LessonDescr[:end]
	curric.StandardNum = curric.StandardNum[:end]
	return curric
}
