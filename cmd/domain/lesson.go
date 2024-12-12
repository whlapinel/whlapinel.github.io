package domain

import "time"

func NewLesson(number int, title string, descr string, date time.Time) Lesson {
	return Lesson{Number: number, Name: title, Description: descr, Date: date}
}

type Lesson struct {
	ID          int
	Number      int
	Name        string
	Description string
	Date        time.Time
}

func (l Lesson) GetTitle() string {
	return l.Name
}
