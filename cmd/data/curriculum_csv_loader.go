package data

import (
	"encoding/csv"
	"fmt"
	"gh_static_portfolio/cmd/domain"
	"os"
	"strconv"
)

const curricCsvDir = "/home/whlapinel/personal_projects/github_portfolio_site/whlapinel.github.io/cmd/data/csv_files/curricula.csv"

// 11/24: copied and pasted this from teacher app, I wanted to do something more proper but couldn't immediately think of how
func ImportCurriculumFromCSV(courseName string) (*domain.CourseSOA, error) {
	curric := domain.CourseSOA{
		Course: domain.NewCourse(courseName, "", []domain.Unit{}, ""),
	}
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
	for i, record := range records {
		if curric.Course.Title() == record[courseNameCol] {
			dayNum := i + 1
			if record[dayNumCol] != "" {
				dayNum, err = strconv.Atoi(record[dayNumCol])
				if err != nil {
					return nil, fmt.Errorf("error reading day number from csv")
				}
			}
			curric.Day = append(curric.Day, dayNum)
			unitNum := 0
			if record[unitNumCol] != "" {
				unitNum, err = strconv.Atoi(record[unitNumCol])
				if err != nil {
					return nil, fmt.Errorf("error reading unit number from csv")
				}
			}
			curric.UnitNum = append(curric.UnitNum, unitNum)
			curric.UnitName = append(curric.UnitName, record[unitDescrCol])
			lessonNum := 0
			if record[lessonNumCol] != "" {
				lessonNum, err = strconv.Atoi(record[lessonNumCol])
				if err != nil {
					return nil, fmt.Errorf("error reading lesson number from csv")
				}
			}
			curric.LessonNum = append(curric.LessonNum, lessonNum)
			curric.LessonName = append(curric.LessonName, record[lessonDescrCol])
			stdNum := 0
			if record[stdNumCol] != "" {
				stdNum, err = strconv.Atoi(record[stdNumCol])
				if err != nil {
					return nil, fmt.Errorf("error reading std number from csv: %s", err)
				}
			}
			curric.StandardNum = append(curric.StandardNum, stdNum)
			curric.StdDescr = append(curric.StdDescr, record[stdDescrCol])

		}
	}
	return &curric, nil
}
