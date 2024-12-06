package domain

import "time"

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
