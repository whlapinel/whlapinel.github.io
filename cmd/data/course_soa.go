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
)

type courseSOARepo struct {
	queries *database.Queries
}

func (c *courseSOARepo) All() ([]*domain.CourseSOA, error) {
	ctx := context.Background()
	rows, err := c.queries.GetCourses(ctx)
	if err != nil {
		return nil, err
	}
	return CourseRowsToCourseSOA(rows)
}

func (c *courseSOARepo) Save(course *domain.CourseSOA) error {
	ctx := context.Background()
	dbCourse, err := c.queries.SaveCourse(ctx, database.SaveCourseParams{Name: course.Name})
	if err != nil {
		return err
	}
	var currUnit *database.Unit
	var currLesson *database.Lesson
	for i := range course.Day {
		if currUnit == nil || int64(course.UnitNum[i]) != currUnit.Number {
			log.Println("creating and saving unit")
			currUnit = &database.Unit{
				Number:   int64(course.UnitNum[i]),
				Name:     course.UnitName[i],
				CourseID: dbCourse.ID,
			}
			*currUnit, err = c.queries.SaveUnit(ctx, database.SaveUnitParams{
				Number:   currUnit.Number,
				Name:     currUnit.Name,
				CourseID: currUnit.CourseID,
			})
			if err != nil {
				log.Fatal(err)
			}
		}
		if currLesson == nil || int64(course.LessonNum[i]) != currLesson.Number {
			log.Println("Lesson Name: ", course.LessonName[i])
			currLesson = &database.Lesson{
				Number: int64(course.LessonNum[i]),
				Name: sql.NullString{
					String: course.LessonName[i],
					Valid:  true,
				},
			}
			*currLesson, err = c.queries.SaveLesson(ctx, database.SaveLessonParams{
				Number: currLesson.Number,
				Name:   currLesson.Name,
				UnitID: currUnit.ID,
			})
			if err != nil {
				log.Fatal(err)
			}
		}

	}
	return nil
}

func (c *courseSOARepo) ReadFromCSV() ([]*domain.CourseSOA, error) {
	courses := []*domain.CourseSOA{}
	file, err := os.Open(curricCsvDir)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	courseMap := make(map[string]*domain.CourseSOA)
	for i, record := range records[1:] {
		courseName := record[courseNameCol]

		course, ok := courseMap[courseName]
		if !ok {
			course = &domain.CourseSOA{Name: courseName}
		}
		dayNum := i + 1
		if record[dayNumCol] != "" {
			dayNum, err = strconv.Atoi(record[dayNumCol])
			if err != nil {
				return nil, fmt.Errorf("error reading day number from csv: %v", err)
			}
		}
		course.Day = append(course.Day, dayNum)
		unitNum := 0
		if record[unitNumCol] != "" {
			unitNum, err = strconv.Atoi(record[unitNumCol])
			if err != nil {
				return nil, fmt.Errorf("error reading unit number from csv")
			}
		}
		course.UnitNum = append(course.UnitNum, unitNum)
		course.UnitName = append(course.UnitName, record[unitDescrCol])
		lessonNum := 0
		if record[lessonNumCol] != "" {
			lessonNum, err = strconv.Atoi(record[lessonNumCol])
			if err != nil {
				return nil, fmt.Errorf("error reading lesson number from csv")
			}
		}
		course.LessonNum = append(course.LessonNum, lessonNum)
		course.LessonName = append(course.LessonName, record[lessonDescrCol])
		stdNum := 0
		if record[stdNumCol] != "" {
			stdNum, err = strconv.Atoi(record[stdNumCol])
			if err != nil {
				return nil, fmt.Errorf("error reading std number from csv: %s", err)
			}
		}
		course.StandardNum = append(course.StandardNum, stdNum)
		course.StdDescr = append(course.StdDescr, record[stdDescrCol])
		courseMap[courseName] = course

	}
	for _, course := range courseMap {
		log.Println("importCoursesFromCSV(): looping through courseMap: Name:", course.Name)
		courses = append(courses, course)
	}
	return courses, nil
}

func (c *courseSOARepo) WriteToCSV(course *domain.CourseSOA) error {
	return fmt.Errorf("not implemented")
}

func NewCourseSOARepo(db *database.Queries) domain.Repository[domain.CourseSOA] {
	return &courseSOARepo{db}
}

func SaveCourse(course *domain.CourseSOA, ctx context.Context, queries *database.Queries) (int64, error) {
	dbCourse, err := queries.SaveCourse(ctx, database.SaveCourseParams{Name: course.Name})
	if err != nil {
		return 0, err
	}
	var currUnit *database.Unit
	var currLesson *database.Lesson
	for i, _ := range course.Day {
		if currUnit == nil || int64(course.UnitNum[i]) != currUnit.Number {
			log.Println("creating and saving unit")
			currUnit = &database.Unit{
				Number:   int64(course.UnitNum[i]),
				Name:     course.UnitName[i],
				CourseID: dbCourse.ID,
			}
			*currUnit, err = queries.SaveUnit(ctx, database.SaveUnitParams{
				Number:   currUnit.Number,
				Name:     currUnit.Name,
				CourseID: currUnit.CourseID,
			})
			if err != nil {
				log.Fatal(err)
			}
		}
		if currLesson == nil || int64(course.LessonNum[i]) != currLesson.Number {
			log.Println("Lesson Name: ", course.LessonName[i])
			currLesson = &database.Lesson{
				Number: int64(course.LessonNum[i]),
				Name: sql.NullString{
					String: course.LessonName[i],
					Valid:  true,
				},
			}
			*currLesson, err = queries.SaveLesson(ctx, database.SaveLessonParams{
				Number: currLesson.Number,
				Name:   currLesson.Name,
				UnitID: currUnit.ID,
			})
			if err != nil {
				log.Fatal(err)
			}
		}

	}
	return dbCourse.ID, nil

}

func CourseRowsToCourseSOA(rows []database.GetCoursesRow) ([]*domain.CourseSOA, error) {
	var curricula = []*domain.CourseSOA{}
	curricMap := make(map[int]*domain.CourseSOA)
	var currID int64 = 0
	var dayCount = 0
	for _, row := range rows {
		dayCount++
		currID = row.CourseID
		curriculum := curricMap[int(currID)]
		curriculum.Day = append(curriculum.Day, dayCount)
		curriculum.UnitNum = append(curriculum.UnitNum, int(row.UnitNumber))
		curriculum.UnitName = append(curriculum.UnitName, row.UnitName)
		curriculum.LessonNum = append(curriculum.LessonNum, int(row.LessonNumber))
		if row.LessonName.Valid {
			curriculum.LessonName = append(curriculum.LessonName, row.LessonName.String)
		}
		curricMap[int(currID)] = curriculum

	}
	for _, val := range curricMap {
		curricula = append(curricula, val)
	}
	return curricula, nil

}

const curricCsvDir = "/home/whlapinel/personal_projects/github_portfolio_site/whlapinel.github.io/cmd/data/csv_files/curricula.csv"

func loadCoursesFromCSV() ([]*domain.CourseSOA, error) {
	courses := []*domain.CourseSOA{}
	file, err := os.Open(curricCsvDir)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	courseMap := make(map[string]*domain.CourseSOA)
	for i, record := range records[1:] {
		courseName := record[courseNameCol]

		course, ok := courseMap[courseName]
		if !ok {
			course = &domain.CourseSOA{Name: courseName}
		}
		dayNum := i + 1
		if record[dayNumCol] != "" {
			dayNum, err = strconv.Atoi(record[dayNumCol])
			if err != nil {
				return nil, fmt.Errorf("error reading day number from csv: %v", err)
			}
		}
		course.Day = append(course.Day, dayNum)
		unitNum := 0
		if record[unitNumCol] != "" {
			unitNum, err = strconv.Atoi(record[unitNumCol])
			if err != nil {
				return nil, fmt.Errorf("error reading unit number from csv")
			}
		}
		course.UnitNum = append(course.UnitNum, unitNum)
		course.UnitName = append(course.UnitName, record[unitDescrCol])
		lessonNum := 0
		if record[lessonNumCol] != "" {
			lessonNum, err = strconv.Atoi(record[lessonNumCol])
			if err != nil {
				return nil, fmt.Errorf("error reading lesson number from csv")
			}
		}
		course.LessonNum = append(course.LessonNum, lessonNum)
		course.LessonName = append(course.LessonName, record[lessonDescrCol])
		stdNum := 0
		if record[stdNumCol] != "" {
			stdNum, err = strconv.Atoi(record[stdNumCol])
			if err != nil {
				return nil, fmt.Errorf("error reading std number from csv: %s", err)
			}
		}
		course.StandardNum = append(course.StandardNum, stdNum)
		course.StdDescr = append(course.StdDescr, record[stdDescrCol])
		courseMap[courseName] = course

	}
	for _, course := range courseMap {
		log.Println("importCoursesFromCSV(): looping through courseMap: Name:", course.Name)
		courses = append(courses, course)
	}
	return courses, nil
}
