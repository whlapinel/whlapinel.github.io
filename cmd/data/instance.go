package data

import (
	"encoding/csv"
	"fmt"
	"gh_static_portfolio/cmd/data/database"
	"gh_static_portfolio/cmd/domain"
	"log"
	"os"
	"slices"
	"strconv"
	"time"
)

const (
	courseNameCol = iota
	dayNumCol
	unitNumCol
	unitSequenceCol
	unitDescrCol
	lessonNumCol
	lessonDescrCol
	stdNumCol
	stdDescrCol
	scheduleDateCol
	scheduleTermIdCol
	scheduleTermNameCol
)

type courseInstanceRepo struct {
	queries *database.Queries
}

// FetchOne implements domain.CourseInstanceRepository.
func (r courseInstanceRepo) FetchOne(courseID int, termID int) ([]*domain.CourseInstance, error) {
	// ctx := context.Background()
	// params := database.GetCourseInstancesParams{
	// 	ID:     int64(courseID),
	// 	TermID: int64(termID),
	// }
	// rows, err := r.queries.GetCourseInstances(ctx, params)
	// if err != nil {
	// 	return nil, err
	// }
	// instances := []*domain.CourseInstance{}
	// for _, row := range
	panic("unimplemented")
}

func NewInstanceRepo(queries *database.Queries) domain.CourseInstanceRepository {
	return courseInstanceRepo{queries}
}

func (r courseInstanceRepo) ReadFromCSV() ([]*domain.CourseInstance, error) {
	return importInstancesFromCSV()
}

func (r courseInstanceRepo) WriteToCSV(instance *domain.CourseInstance) error {
	instances := []*domain.CourseInstance{instance}
	return WriteInstancesToCSV(instances)

}

// Save to Database
func (r courseInstanceRepo) Save(instance *domain.CourseInstance) error {
	return fmt.Errorf("not implemented")
}

const scheduleCsvDir = "/home/whlapinel/personal_projects/github_portfolio_site/whlapinel.github.io/cmd/data/csv_files/schedules.csv"

func GenerateInstancesFromCSV() ([]*domain.CourseInstance, error) {
	courses, err := loadCoursesFromCSV()
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
	var schedules []*domain.CourseInstance
	for _, course := range courses {
		schedule, err := domain.NewCourseInstance(*currentTerm, *course)
		if err != nil {
			return nil, err
		}
		schedules = append(schedules, schedule)
	}
	return schedules, nil

}

func WriteInstancesToCSV(instances []*domain.CourseInstance) error {
	err := os.Remove("./cmd/data/csv_files/schedules.csv")
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
	for _, instance := range instances {
		for i := range instance.Day {
			courseName := instance.Course.GetTitle()
			day := strconv.Itoa(instance.Day[i])
			unitNum := strconv.Itoa(instance.UnitNum[i])
			unitDescr := instance.UnitDescr[i]
			lessonNum := strconv.Itoa(instance.LessonNum[i])
			lessonDescr := instance.LessonDescr[i]
			std := strconv.Itoa(instance.StandardNum[i])
			stdDescr := instance.StdDescr[i]
			date := instance.Date[i].String()[:10]
			termID := strconv.Itoa(instance.Term.ID)
			termName := instance.Term.Name
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

func importInstancesFromCSV() ([]*domain.CourseInstance, error) {
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
	instanceMap := make(map[string]*domain.CourseInstance)
	for i, record := range records[1:] {
		courseName := record[courseNameCol]
		instance, exists := instanceMap[courseName]
		if !exists {
			instance = &domain.CourseInstance{
				CourseSOA: &domain.CourseSOA{
					Name: courseName}}
		}
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
		unitSequence, err := strconv.Atoi(record[unitSequenceCol])
		if err != nil {
			return nil, fmt.Errorf("error reading unit sequence number from csv")
		}
		instance.UnitSequence = append(instance.UnitSequence, unitSequence)
		instance.UnitDescr = append(instance.UnitDescr, record[unitDescrCol])
		lessonNum := 0
		if record[lessonNumCol] != "" {
			lessonNum, err = strconv.Atoi(record[lessonNumCol])
			if err != nil {
				return nil, fmt.Errorf("error reading lesson number from csv")
			}
		}
		instance.LessonNum = append(instance.LessonNum, lessonNum)
		instance.LessonDescr = append(instance.LessonDescr, record[lessonDescrCol])
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
		termID, err := strconv.Atoi(record[scheduleTermIdCol])
		if err != nil {
			return nil, err
		}
		instance.Term = &domain.Term{}
		instance.Term.ID = termID
		instance.Term.Name = record[scheduleTermNameCol]
		instanceMap[courseName] = instance

	}
	var instances []*domain.CourseInstance
	for _, instance := range instanceMap {
		instances = append(instances, instance)

	}
	return instances, nil
}

// // convert schedule SOA to nested objects since that's how I wrote it originally
// func courseInstanceSoaToOop(instance domain.CourseInstance) *domain.Course {
// 	unitLessonCounter := make(map[int]int) // unit number and lesson count
// 	var currUnit domain.Unit

// 	for i, unitNum := range instance.UnitNum {
// 		if _, exists := unitLessonCounter[unitNum]; !exists {
// 			if currUnit.GetTitle() != "" {
// 				instance.Course.Units = append(instance.Course.Units, currUnit) // if we've hit a new unit we should append the previous unit since it's complete (crucial assumption is that everything is in order)
// 			}
// 			unitLessonCounter[unitNum] = 1
// 			lessonName := fmt.Sprintf("Lesson %d.%d", unitNum, instance.LessonNum[i])
// 			if instance.UnitNum[i] < 0 {
// 				lessonName = fmt.Sprintf("%s Day %d", instance.UnitDescr[i], instance.LessonNum[i])
// 			}
// 			lessons := []domain.Lesson{
// 				domain.NewLesson(instance.LessonNum[i], lessonName, instance.LessonDescr[i], instance.Date[i]),
// 			}
// 			unitName := fmt.Sprintf("Unit %d", unitNum)
// 			if unitNum < 0 {
// 				unitName = instance.UnitDescr[i]
// 			}
// 			currUnit = domain.NewUnit(unitNum, unitName, instance.UnitDescr[i], lessons)
// 		} else {
// 			// if it's the same as current unit, create a new lesson and increment the count
// 			unitLessonCounter[unitNum]++
// 			lessonTitle := fmt.Sprintf("Lesson %d.%d", unitNum, instance.LessonNum[i])
// 			if instance.UnitNum[i] < 0 {
// 				lessonTitle = fmt.Sprintf("%s Day %d", instance.UnitDescr[i], instance.LessonNum[i])
// 			}
// 			lesson := domain.NewLesson(instance.LessonNum[i], lessonTitle, instance.LessonDescr[i], instance.Date[i])
// 			currUnit.Lessons = append(currUnit.Lessons, lesson)
// 		}
// 	}
// 	instance.Course.Units = append(instance.Course.Units, currUnit) // append the final unit
// 	instance.Course.TermName = instance.Term.Name
// 	instance.Course.Name = instance.CourseSOA.Name
// 	return &instance.Course

// }

func courseInstanceSoaToOopV2(instance domain.CourseInstance) *domain.Course {
	unitMap := make(map[int]domain.Unit)
	course := &domain.Course{
		Name:     instance.CourseSOA.Name,
		TermName: instance.Term.Name,
	}

	for i, sequence := range instance.UnitSequence {
		var unit domain.Unit
		if _, exists := unitMap[sequence]; exists {
			unit = unitMap[sequence]
		} else {
			unitName := fmt.Sprintf("Unit %d", instance.UnitNum[i])
			if instance.UnitNum[i] < 0 {
				unitName = instance.UnitDescr[i]
			}
			unit = domain.NewUnit(instance.UnitNum[i], sequence, unitName, instance.UnitDescr[i], []domain.Lesson{})
		}
		lessonNum := instance.LessonNum[i]
		lessonName := fmt.Sprintf("Lesson %d.%d", instance.UnitNum[i], instance.LessonNum[i])
		if instance.UnitNum[i] < 0 {
			lessonName = fmt.Sprintf("%s Day %d", instance.UnitDescr[i], instance.LessonNum[i])
		}
		lessonDescr := instance.LessonDescr[i]
		lessonDate := instance.Date[i]
		lesson := domain.NewLesson(lessonNum, lessonName, lessonDescr, lessonDate)
		unit.Lessons = append(unit.Lessons, lesson)
		unitMap[sequence] = unit
	}
	units := sortedBySequence(unitMap)
	for _, unitNum := range units {
		unit := unitMap[unitNum]
		course.Units = append(course.Units, unit)
	}
	return course

}

func sortedBySequence(unitMap map[int]domain.Unit) []int {
	keys := make([]int, 0, len(unitMap))
	for sequence := range unitMap {
		keys = append(keys, sequence)
	}

	slices.Sort(keys)
	return keys

}

