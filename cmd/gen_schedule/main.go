package main

import (
	"encoding/csv"
	"gh_static_portfolio/cmd/data"
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
	py1curric, err := data.ImportCurriculumFromCSV("Python I Programming Honors")
	if err != nil {
		return err
	}
	py2curric, err := data.ImportCurriculumFromCSV("Python II Programming Honors")
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
	var schedules []domain.CourseInstance
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
			unitDescr := schedule.UnitName[i]
			lessonNum := strconv.Itoa(schedule.LessonNum[i])
			lessonDescr := schedule.LessonName[i]
			std := strconv.Itoa(schedule.StandardNum[i])
			stdDescr := schedule.StdDescr[i]
			date := schedule.Date[i].String()[:10]
			termID := strconv.Itoa(schedule.Term.ID)
			termName := schedule.Term.Name
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
