package domain

import (
	"time"
)

type CourseInstanceRepository interface {
	FetchOne(courseID int, termID int) ([]*CourseInstance, error)
	Saver[CourseInstance]
	CSVer[CourseInstance]
}

// Create schedule from term and curriculum. If term is shorter than curriculum days, curriculum will be truncated
func NewCourseInstance(term Term, curric CourseSOA) (*CourseInstance, error) {
	dates := InstructionDays(term)
	var cs CourseInstance
	if len(curric.Day) > len(dates) {
		curric = curric.Truncate(len(dates))
	}
	cs = CourseInstance{
		CourseSOA: &curric,
		Date:      dates,
		Term:      &Term{ID: term.ID, Name: term.Name},
	}
	return &cs, nil
}

// combine a term and a curriculum and you get a schedule
type CourseInstance struct {
	*CourseSOA
	*Term
	Date []time.Time
}

func (c CourseInstance) GetTitle() string {
	return c.CourseSOA.Name
}
