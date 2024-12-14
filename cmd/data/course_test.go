package data

import (
	"log"
	"testing"
	"time"
)

func TestImportCoursesFromCSV(t *testing.T) {
	courses, err := importCoursesFromCSV()
	if err != nil {
		t.Errorf("error importing courses: %s", err)
	}
	for _, course := range courses {
		log.Println(course.Name)
		log.Println("num units: ", len(course.Units))
		for _, unit := range course.Units {
			log.Println(unit.Name)
			log.Println("num lessons: ", len(unit.Lessons))
			for _, lesson := range unit.Lessons {
				log.Println(lesson.Name, lesson.Description, lesson.Date.Format(time.DateOnly))
			}
		}
	}

}
