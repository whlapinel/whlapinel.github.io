package domain

import (
	"fmt"
	"time"
)

type Repository[T any] interface {
	Save(*T) error
	All() ([]*T, error)
	WriteToCSV(*T) error
	ReadFromCSV() ([]*T, error)
}

type BlogRepository interface {
	GetAll() ([]*Blog, error)
}

type Blog struct {
	title   string
	content string
}

func NewBlog(title, content string) *Blog {
	return &Blog{
		title:   title,
		content: content,
	}
}

func (b *Blog) GetTitle() string {
	return b.title
}

func (b *Blog) Content() string {
	return b.content
}

type EducationRepository interface {
	GetAll() ([]*EducationItem, error)
}

type EducationItem struct {
	Image   string
	Year    string
	School  string
	Degree  string
	Minor   string
	Classes []*ClassItem
}

func (e *EducationItem) GetTitle() string {
	return e.School
}

type ProjectRepository interface {
	GetAll() ([]*ProjectItem, error)
}

type ProjectItem struct {
	Image       string
	Title       string
	Subtitle    string
	Description string
	Link        string
}

type ClassRepository interface {
	GetAll() ([]*ClassItem, error)
}

type ClassItem struct {
	Image       string
	Number      string
	Title       string
	Subtitle    string
	Description string
	Link        string
}

type WorkHistoryRepository interface {
	GetAll() ([]*WorkHistoryItem, error)
}

type WorkHistoryItem struct {
	Image       string
	Year        string
	Company     string
	Position    string
	Description string
}

type SkillRepository interface {
	GetAll() ([]*SkillItem, error)
}

type SkillItem struct {
	Image       string
	Title       string
	Subtitle    string
	Description string
}

func NewCourse(title string, descr string, units []Unit, termName string) Course {
	return Course{Title: title, Description: descr, Units: units, TermName: termName}

}

type CourseRepository interface {
	GetAll() ([]Course, error)
}

// Courses I teach. this is the OOP version of CourseInstance. Bad wording I know.
type Course struct {
	ID          int
	Title       string
	Description string
	Units       []Unit
	TermName    string
}

func (c Course) GetTitle() string {
	return c.Title
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

type CourseSOARepository interface {
	GetAll() ([]*CourseSOA, error)
	Save(*CourseSOA) error
}

type CourseSOA struct {
	ID          int
	Name        string
	Course      Course
	Day         []int // day of instruction number e.g. 1, 2, ...87
	UnitNum     []int // unit number e.g. 1, 2, 3
	UnitName    []string
	LessonNum   []int // lesson number e.g. 1, 2, 3
	LessonName  []string
	StandardNum []int
	StdDescr    []string
}

func (c CourseSOA) GetTitle() string {
	return c.Name
}

func (curric CourseSOA) Truncate(end int) CourseSOA {
	curric.Day = curric.Day[:end]
	curric.UnitNum = curric.UnitNum[:end]
	curric.UnitName = curric.UnitName[:end]
	curric.LessonNum = curric.LessonNum[:end]
	curric.LessonName = curric.LessonName[:end]
	curric.StandardNum = curric.StandardNum[:end]
	return curric
}

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

func NewLesson(number int, title string, descr string, date time.Time) Lesson {
	return Lesson{Number: number, Title: title, Description: descr, Date: date}
}

type Lesson struct {
	Number      int
	Title       string
	Description string
	Date        time.Time
}

func (l Lesson) GetTitle() string {
	return l.Title
}
func NewTerm(start, end time.Time, nonInstructionalDays []time.Time, termType TermType, termID int, name string) (*Term, error) {
	if start.After(end) {
		return nil, fmt.Errorf("start must come before end: %s is after %s", start.String(), end.String())
	}
	term := Term{Start: start, End: end, NonInstructionalDays: nonInstructionalDays, TermType: termType, ID: termID, Name: name}
	return &term, nil
}

type TermType string

type Term struct {
	Start                time.Time
	End                  time.Time
	NonInstructionalDays []time.Time
	TermType             TermType
	ID                   int
	Name                 string
}

const (
	Semester = "semester"
	YearLong = "year_long"
)

type NonInstructionalDays struct {
	TermID []int
	Dates  []time.Time
}

// Create schedule from term and curriculum. If term is shorter than curriculum days, curriculum will be truncated
func NewCourseSchedule(term Term, curric CourseSOA) (*CourseInstance, error) {
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

func NonInstructionDayRange(start time.Time, end time.Time) []time.Time {
	dateRange := []time.Time{}
	currDate := start
	for !currDate.After(end) {
		dateRange = append(dateRange, currDate)
		currDate = currDate.Add(24 * time.Hour)
	}
	return dateRange
}

func InstructionDays(term Term) []time.Time {
	currDate := term.Start
	instructionDays := []time.Time{}
	for !currDate.After(term.End) {
		isInstructionDay := true
		if currDate.Weekday() == 0 || currDate.Weekday() == 6 {
			isInstructionDay = false
		}
		for _, day := range term.NonInstructionalDays {
			if IsSameDate(currDate, day) {
				isInstructionDay = false
			}
		}
		if isInstructionDay {
			instructionDays = append(instructionDays, currDate)
		}
		currDate = currDate.Add(24 * time.Hour)
	}
	return instructionDays
}

func IsSameDate(t1 time.Time, t2 time.Time) bool {
	y1, m1, d1 := t1.Date()
	y2, m2, d2 := t2.Date()
	if y1 == y2 && m1 == m2 && d1 == d2 {
		return true
	}
	return false
}
