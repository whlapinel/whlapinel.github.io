package data

import (
	"encoding/csv"
	"fmt"
	"gh_static_portfolio/cmd/domain"
	"log"
	"os"
	"strconv"
	"time"
)

const (
	courseNameCol = iota
	dayNumCol
	unitNumCol
	unitDescrCol
	lessonNumCol
	lessonDescrCol
	stdNumCol
	stdDescrCol
	scheduleDateCol
	termIdCol
	termNameCol
)

const scheduleCsvDir = "./cmd/data/csv_files/schedules.csv"

// 11/24: copied and pasted this from teacher app, I wanted to do something more proper but couldn't immediately think of how
func ImportScheduleFromCSV(courseName string) (*domain.CourseInstance, error) {
	instance := domain.CourseInstance{
		CourseSOA: &domain.CourseSOA{Course: domain.NewCourse(courseName, "", []domain.Unit{}, "")},
	}
	file, err := os.Open(scheduleCsvDir)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	if len(records) == 0 {
		log.Fatalf("file is empty")
	}
	for i, record := range records {
		if instance.Course.Title() == record[courseNameCol] {
			dayNum := i + 1
			if record[dayNumCol] != "" {
				dayNum, err = strconv.Atoi(record[dayNumCol])
				if err != nil {
					return nil, fmt.Errorf("error reading day number from csv")
				}
			}
			instance.Day = append(instance.Day, dayNum)
			unitNum := 0
			if record[unitNumCol] != "" {
				unitNum, err = strconv.Atoi(record[unitNumCol])
				if err != nil {
					return nil, fmt.Errorf("error reading unit number from csv")
				}
			}
			instance.UnitNum = append(instance.UnitNum, unitNum)
			instance.UnitName = append(instance.UnitName, record[unitDescrCol])
			lessonNum := 0
			if record[lessonNumCol] != "" {
				lessonNum, err = strconv.Atoi(record[lessonNumCol])
				if err != nil {
					return nil, fmt.Errorf("error reading lesson number from csv")
				}
			}
			instance.LessonNum = append(instance.LessonNum, lessonNum)
			instance.LessonName = append(instance.LessonName, record[lessonDescrCol])
			stdNum := 0
			if record[stdNumCol] != "" {
				stdNum, err = strconv.Atoi(record[stdNumCol])
				if err != nil {
					return nil, fmt.Errorf("error reading std number from csv: %s", err)
				}
			}
			instance.StandardNum = append(instance.StandardNum, stdNum)
			instance.StdDescr = append(instance.StdDescr, record[stdDescrCol])
			date, err := time.Parse(time.DateOnly, record[scheduleDateCol])
			if err != nil {
				return nil, err
			}
			instance.Date = append(instance.Date, date)
			termID, err := strconv.Atoi(record[termIdCol])
			if err != nil {
				return nil, err
			}
			instance.Term = &domain.Term{}
			instance.Term.ID = termID
			instance.Term.Name = record[termNameCol]

		}
	}
	return &instance, nil
}

// convert schedule SOA to nested objects since that's how I wrote it originally
func courseScheduleSoaToOop(schedule domain.CourseInstance) domain.Course {
	unitLessonCounter := make(map[int]int) // unit number and lesson count
	var currUnit domain.Unit

	for i, unitNum := range schedule.UnitNum {
		log.Println("curricSoaToOop: i=", i, "unitNum=", unitNum, "currUnit=", currUnit)
		if _, exists := unitLessonCounter[unitNum]; !exists {
			if currUnit.Title() != "" {
				schedule.Course.Units = append(schedule.Course.Units, currUnit) // if we've hit a new unit we should append the previous unit since it's complete (crucial assumption is that everything is in order)
			}
			unitLessonCounter[unitNum] = 1
			lessonTitle := fmt.Sprintf("Lesson %d.%d", unitNum, schedule.LessonNum[i])
			if schedule.UnitNum[i] < 0 {
				lessonTitle = fmt.Sprintf("%s Day %d", schedule.UnitName[i], schedule.LessonNum[i])
			}
			lessons := []domain.Lesson{
				domain.NewLesson(schedule.LessonNum[i], lessonTitle, schedule.LessonName[i], schedule.Date[i]),
			}
			unitTitle := fmt.Sprintf("Unit %d", unitNum)
			if unitNum < 0 {
				unitTitle = schedule.UnitName[i]
			}
			currUnit = domain.NewUnit(unitNum, unitTitle, schedule.UnitName[i], lessons)
		} else {
			// if it's the same as current unit, create a new lesson and increment the count
			unitLessonCounter[unitNum]++
			lessonTitle := fmt.Sprintf("Lesson %d.%d", unitNum, schedule.LessonNum[i])
			if schedule.UnitNum[i] < 0 {
				lessonTitle = fmt.Sprintf("%s Day %d", schedule.UnitName[i], schedule.LessonNum[i])
			}
			lesson := domain.NewLesson(schedule.LessonNum[i], lessonTitle, schedule.LessonName[i], schedule.Date[i])
			currUnit.Lessons = append(currUnit.Lessons, lesson)
		}
	}
	schedule.Course.Units = append(schedule.Course.Units, currUnit) // append the final unit
	schedule.Course.TermName = schedule.Term.Name
	return schedule.Course

}
