package main

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
)

func main() {
	err := WriteScheduleToCSV()
	if err != nil {
		log.Fatalf("error writing schedule to csv, %v", err)
	}

}

func WriteScheduleToCSV() error {
	py1curric, err := ImportCurriculumFromCSV("Python I Programming Honors")
	if err != nil {
		return err
	}
	py2curric, err := ImportCurriculumFromCSV("Python II Programming Honors")
	if err != nil {
		return err
	}
	terms, err := TermsLoader()
	if err != nil {
		return err
	}

	var currentTerm domain.Term
	for _, term := range terms {
		if time.Now().After(term.Start) && time.Now().Before(term.End) {
			currentTerm = term
		}
	}
	var schedules []domain.CourseSchedule
	py1currentSchedule, err := domain.NewCourseSchedule(currentTerm, *py1curric)
	if err != nil {
		return err
	}
	py2currentSchedule, err := domain.NewCourseSchedule(currentTerm, *py2curric)
	if err != nil {
		return err
	}
	schedules = append(schedules, *py1currentSchedule, *py2currentSchedule)
	err = os.Remove("./cmd/data/csv_files/schedules.csv")
	if err != nil {
		return err
	}
	file, err := os.Create("./cmd/data/csv_files/schedules.csv")
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
	for _, schedule := range schedules {
		for i := range schedule.Day {
			courseName := schedule.Course.Title()
			day := strconv.Itoa(schedule.Day[i])
			unitNum := strconv.Itoa(schedule.UnitNum[i])
			unitDescr := schedule.UnitDescr[i]
			lessonNum := strconv.Itoa(schedule.LessonNum[i])
			lessonDescr := schedule.LessonDescr[i]
			std := strconv.Itoa(schedule.StandardNum[i])
			stdDescr := schedule.StdDescr[i]
			date := schedule.Date[i].String()[:10]
			termID := strconv.Itoa(schedule.TermID)
			termName := schedule.TermName
			rows = append(rows, []string{
				courseName,
				day,
				unitNum,
				unitDescr,
				lessonNum,
				lessonDescr,
				std,
				stdDescr,
				date,
				termID,
				termName,
			})

		}
	}
	writer.WriteAll(rows)
	return nil

}

const curricCsvDir = "/home/whlapinel/personal_projects/github_portfolio_site/whlapinel.github.io/cmd/data/csv_files/curricula.csv"

// 11/24: copied and pasted this from teacher app, I wanted to do something more proper but couldn't immediately think of how
func ImportCurriculumFromCSV(courseName string) (*domain.Curriculum, error) {
	curric := domain.Curriculum{
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
			curric.UnitDescr = append(curric.UnitDescr, record[unitDescrCol])
			lessonNum := 0
			if record[lessonNumCol] != "" {
				lessonNum, err = strconv.Atoi(record[lessonNumCol])
				if err != nil {
					return nil, fmt.Errorf("error reading lesson number from csv")
				}
			}
			curric.LessonNum = append(curric.LessonNum, lessonNum)
			curric.LessonDescr = append(curric.LessonDescr, record[lessonDescrCol])
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

const termsPath = "/home/whlapinel/personal_projects/github_portfolio_site/whlapinel.github.io/cmd/data/csv_files/terms.csv"
const nonIDaysPath = "/home/whlapinel/personal_projects/github_portfolio_site/whlapinel.github.io/cmd/data/csv_files/non_instruct_days.csv"

const (
	termIDCol = iota
	termStartCol
	termEndCol
	termTypeCol
	termNameCol
)

const (
	nonInstructionalDateCol = 1
)

func nonInstructionalDaysLoader() (*domain.NonInstructionalDays, error) {
	dates := domain.NonInstructionalDays{}
	file, err := os.Open(nonIDaysPath)
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
		if i == 0 {
			continue
		}
		termID, err := strconv.Atoi(record[termIDCol])
		if err != nil {
			return nil, err
		}
		dates.TermID = append(dates.TermID, termID)
		date, err := time.Parse(time.DateOnly, record[nonInstructionalDateCol])
		if err != nil {
			return nil, err
		}
		dates.Dates = append(dates.Dates, date)
	}
	return &dates, nil
}

func filterNonInstructionalDates(termID int, dates *domain.NonInstructionalDays) []time.Time {
	filtered := []time.Time{}
	for i, date := range dates.Dates {
		if dates.TermID[i] == termID {
			filtered = append(filtered, date)
		}
	}
	return filtered
}

func TermsLoader() ([]domain.Term, error) {
	file, err := os.Open(termsPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	dates, err := nonInstructionalDaysLoader()
	if err != nil {
		return nil, err
	}
	terms := []domain.Term{}
	for i, record := range records {
		if i == 0 {
			continue
		}
		startDate, err := time.Parse(time.DateOnly, record[termStartCol])
		if err != nil {
			return nil, err
		}
		endDate, err := time.Parse(time.DateOnly, record[termEndCol])
		if err != nil {
			return nil, err
		}
		termID, err := strconv.Atoi(record[termIDCol])
		if err != nil {
			return nil, err
		}
		termDates := filterNonInstructionalDates(termID, dates)
		termType := record[termTypeCol]

		termName := record[termNameCol]
		term, err := domain.NewTerm(startDate, endDate, termDates, domain.TermType(termType), termID, termName)
		if err != nil {
			return nil, err
		}
		terms = append(terms, *term)

	}
	return terms, nil

}
