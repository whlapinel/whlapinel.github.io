package data

import (
	"context"
	"database/sql"
	"encoding/csv"
	"fmt"
	"gh_static_portfolio/cmd/data/database"
	"gh_static_portfolio/cmd/domain"
	"log"
	"os"
	"strconv"
	"time"
)

type courseRepo struct {
	queries *database.Queries
}

// GetInstances implements domain.CourseRepository.
func (c *courseRepo) GetInstances() ([]*domain.Course, error) {
	panic("unimplemented")
}

// GetTemplates implements domain.CourseRepository.
func (c *courseRepo) GetTemplates() ([]*domain.Course, error) {
	dbCourses, err := c.queries.GetTemplates(context.Background())
	if err != nil {
		return nil, err
	}
	var courses []*domain.Course
	for _, dbCourse := range dbCourses {
		course := &domain.Course{
			ID:          int(dbCourse.CourseID),
			Name:        dbCourse.CourseName,
			Description: dbCourse.CourseDescr.String,
		}
		dbUnits, err := c.queries.GetUnits(context.Background(), dbCourse.CourseID)
		if err != nil {
			return nil, err
		}
		var units []domain.Unit
		for _, dbUnit := range dbUnits {
			unit := domain.Unit{
				ID:          int(dbUnit.ID),
				CourseID:    int(dbCourse.CourseID),
				Number:      int(dbUnit.Number),
				Name:        dbUnit.Name,
				Description: dbUnit.Description.String,
			}
			dbLessons, err := c.queries.GetLessons(context.Background(), dbUnit.ID)
			if err != nil {
				return nil, err
			}
			var lessons []domain.Lesson
			for _, dbLesson := range dbLessons {
				lesson := domain.Lesson{
					ID:          int(dbLesson.ID),
					Number:      int(dbLesson.Number),
					Name:        dbLesson.Name.String,
					Description: dbLesson.Description.String,
				}
				lessons = append(lessons, lesson)
			}
			unit.Lessons = lessons
			units = append(units, unit)
		}
		course.Units = units
		courses = append(courses, course)
	}
	return courses, nil
}

func NewCourseRepo(db *database.Queries) domain.CourseRepo {
	return &courseRepo{queries: db}
}

func (c *courseRepo) All() ([]*domain.Course, error) {
	panic("not implemented")
}

func (c *courseRepo) WriteToCSV(course *domain.Course) error {
	return fmt.Errorf("not implemented")
}

// TODO: Convert this to import directly from CSV rather than using instance import
func (c *courseRepo) ReadFromCSV() ([]*domain.Course, error) {

	// instances, err := importInstancesFromCSV()
	// if err != nil {
	// 	return nil, err
	// }
	// var courses []*domain.Course
	// for _, instance := range instances {
	// 	courses = append(courses, courseInstanceSoaToOopV2(*instance))
	// }
	// return courses, nil
	courses, err := importCoursesFromCSV()
	if err != nil {
		return nil, err
	}
	return courses, nil

}

// Save updates the ID for Course, Units, and Lessons
func (c *courseRepo) Save(course *domain.Course) error {
	ctx := context.Background()
	dbCourse, err := c.queries.SaveCourse(ctx, database.SaveCourseParams{Name: course.Name})
	if err != nil {
		return err
	}
	course.ID = int(dbCourse.ID)
	for i, unit := range course.Units {
		var hasDescr = unit.Description == ""
		currUnit := database.Unit{
			CourseID: int64(course.ID),
			Number:   int64(unit.Number),
			Name:     unit.Name,
			Description: sql.NullString{
				String: unit.Description,
				Valid:  hasDescr,
			},
		}
		currUnit, err = c.queries.SaveUnit(ctx, database.SaveUnitParams{
			Number:      currUnit.Number,
			Name:        currUnit.Name,
			Description: currUnit.Description,
			CourseID:    currUnit.CourseID,
		})
		if err != nil {
			return fmt.Errorf("courseRepo.Save(): %s", err)
		}
		unit.ID = int(currUnit.ID)
		for j, lesson := range unit.Lessons {
			log.Println("lesson name:", lesson.Name, "lesson number: ", lesson.Number)
			currLesson := database.Lesson{
				UnitID: int64(unit.ID),
				Number: int64(lesson.Number),
				Name: sql.NullString{
					String: lesson.Name,
					Valid:  lesson.Name != "",
				},
				Description: sql.NullString{
					String: lesson.Description,
					Valid:  lesson.Description != "",
				},
			}
			currLesson, err = c.queries.SaveLesson(ctx, database.SaveLessonParams{
				Number:      currLesson.Number,
				Name:        currLesson.Name,
				Description: currLesson.Description,
				UnitID:      currLesson.UnitID,
			})
			if err != nil {
				return fmt.Errorf("courseRepo.Save(), SaveLesson(): %v, unit id: %s, lesson#: %s", err, strconv.Itoa(int(currUnit.ID)), strconv.Itoa(int(currLesson.Number)))
			}
			course.Units[i].Lessons[j].ID = int(currLesson.ID)
		}

	}
	return nil

}

func importCoursesFromCSV() ([]*domain.Course, error) {

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
	type UnitMap map[int]domain.Unit
	type CourseHolder struct {
		Course domain.Course
		Units  UnitMap
	}
	courseMap := make(map[string]CourseHolder)
	for _, record := range records[1:] {
		courseName := record[courseNameCol]
		courseHolder, exists := courseMap[courseName]
		termName := record[scheduleTermNameCol]
		if !exists {
			course := domain.Course{
				Name:     courseName,
				TermName: termName,
			}
			holder := CourseHolder{
				Course: course,
				Units:  UnitMap{},
			}
			courseMap[courseName] = holder
			courseHolder = holder
		}
		unitNum := 0
		if record[unitNumCol] != "" {
			unitNum, err = strconv.Atoi(record[unitNumCol])
			if err != nil {
				return nil, fmt.Errorf("error reading unit number from csv")
			}
		} else {
			return nil, fmt.Errorf("unit number field is blank")
		}
		unitSequence, err := strconv.Atoi(record[unitSequenceCol])
		if err != nil {
			return nil, fmt.Errorf("error reading unit sequence number from csv")
		}
		unitName := fmt.Sprintf("Unit %d", unitNum)
		if unitNum < 0 {
			unitName = record[unitDescrCol]
		}
		unitDescr := record[unitDescrCol]
		unit, exists := courseHolder.Units[unitSequence]
		if !exists {
			courseHolder.Units[unitSequence] = domain.Unit{
				Number:      unitNum,
				SequenceNum: unitSequence,
				Name:        unitName,
				Description: unitDescr,
			}
			unit = courseHolder.Units[unitSequence]
		}
		lessonNum := 0
		if record[lessonNumCol] != "" {
			lessonNum, err = strconv.Atoi(record[lessonNumCol])
			if err != nil {
				return nil, fmt.Errorf("error reading lesson number from csv")
			}
		}
		lessonName := fmt.Sprintf("Lesson %d.%d", unitNum, lessonNum)
		if unitNum < 0 {
			lessonName = fmt.Sprintf("%s Day %d", unitName, lessonNum)

		}
		lessonDescr := record[lessonDescrCol]
		lessonDate, err := time.Parse(time.DateOnly, record[scheduleDateCol])
		if err != nil {
			return nil, fmt.Errorf("error parsing date: %s", err)
		}
		lesson := domain.NewLesson(lessonNum, lessonName, lessonDescr, lessonDate)
		unit.Lessons = append(unit.Lessons, lesson)
		courseHolder.Units[unitSequence] = unit
		courseMap[courseName] = courseHolder
	}
	var courses []*domain.Course
	for _, holder := range courseMap {
		course := holder.Course
		nums := sortedBySequence(holder.Units)
		for _, num := range nums {
			course.Units = append(course.Units, holder.Units[num])
		}
		courses = append(courses, &course)

	}
	return courses, nil
}

// This imports a course template and a term and generates a course instance
func GenerateCourseInstancesFromCSV2() ([]*domain.Course, error) {
	courses, err := importCoursesFromCSV()
	if err != nil {
		return nil, err
	}
	terms, err := TermsLoader()
	if err != nil {
		return nil, err
	}
	var currentTerm *domain.Term
	for _, term := range terms {
		if time.Now().After(term.Start) && time.Now().Before(term.End) {
			currentTerm = term
		}
	}
	currDate := currentTerm.InstructionalDays[0]
	dateNum := 0
courseLoop:
	for i, course := range courses {
		for j, unit := range course.Units {
			for k, lesson := range unit.Lessons {
				log.Printf("Assigning %v to lesson %v", currDate, lesson.Name) // Log assignment
				courses[i].Units[j].Lessons[k].Date = currDate
				if dateNum != len(currentTerm.InstructionalDays)-1 {
					dateNum++
					currDate = currentTerm.InstructionalDays[dateNum]
				} else {
					break courseLoop
				}
			}
		}
	}
	return courses, nil
}

func WriteCourseInstancesToCSV(instances []*domain.Course) error {
	file, err := os.Create(newScheduleDir)
	if err != nil {
		return err
	}
	writer := csv.NewWriter(file)
	var rows [][]string
	rows = append(rows, []string{
		"course_name",
		"day_num",
		"unit_num",
		"unit_descr",
		"lesson_num",
		"lesson_descr",
		"std_num",
		"std_descr",
		"date",
		"term_id",
		"term_name",
	})
	dayNum := 1
	var currDate time.Time
	for _, instance := range instances {
		for _, unit := range instance.Units {
			for _, lesson := range unit.Lessons {
				emptyTime := time.Time{}
				if currDate == emptyTime {
					currDate = lesson.Date
				}
				if lesson.Date != currDate {
					dayNum++
				}
				courseName := instance.Name
				unitNum := unit.Number
				unitDescr := unit.Description
				lessonNum := lesson.Number
				stdNum := ""
				stdDescr := ""
				date := currDate
				termID := ""
				termName := instance.TermName
				row := []string{
					courseName,
					strconv.Itoa(unitNum),
					unitDescr,
					strconv.Itoa(lessonNum),
					stdNum,
					stdDescr,
					date.Format(time.DateOnly),
					termID,
					termName,
				}
				rows = append(rows, row)

			}

		}
	}
	writer.WriteAll(rows)
	return nil

}
