package data

import (
	"encoding/csv"
	"fmt"
	"gh_static_portfolio/cmd/domain"
	"log"
	"os"
	"strconv"
)

type Curriculum struct {
	Course      domain.Course
	Day         []int // day of instruction number e.g. 1, 2, ...87
	UnitNum     []int // unit number e.g. 1, 2, 3
	UnitDescr   []string
	LessonNum   []int // lesson number e.g. 1, 2, 3
	LessonDescr []string
	Standard    []int
}

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

const (
	filesRootDir = "https://github.com/whlapinel"
)

var (
	coursePath = map[string]string{
		"Python II Programming Honors": "python-2-course",
		"Python I Programming Honors":  "python-1-course",
	}
)

const curricCsvDir = "/home/whlapinel/pob/teacher_app_v3/data/curricula.csv"

func lessonUrl(courseName string, unit, lesson int) string {
	return fmt.Sprintf("%s/%s/blob/main/unit%d/lesson%d_%d", filesRootDir, coursePath[courseName], unit, unit, lesson)
}

// 11/24: copied and pasted this from teacher app, I wanted to do something more proper but couldn't immediately think of how
func ImportCurriculumFromCSV(courseName string) (*Curriculum, error) {
	curric := Curriculum{
		Course: domain.NewCourse(courseName, "", []domain.Unit{}),
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
			curric.Standard = append(curric.Standard, stdNum)

		}
	}
	return &curric, nil
}

// convert curriculum SOA to nested objects since that's how I wrote it originally
func curricSoaToOop(curric Curriculum) domain.Course {
	unitLessonCounter := make(map[int]int) // unit number and lesson count
	var currUnit domain.Unit

	for i, unitNum := range curric.UnitNum {
		log.Println("curricSoaToOop: i=", i, "unitNum=", unitNum, "currUnit=", currUnit)
		if _, exists := unitLessonCounter[unitNum]; !exists {
			if currUnit.Title() != "" {
				curric.Course.Units = append(curric.Course.Units, currUnit) // if we've hit a new unit we should append the previous unit since it's complete (crucial assumption is that everything is in order)
			}
			unitLessonCounter[unitNum] = 1
			lessons := []domain.Lesson{
				domain.NewLesson(curric.LessonNum[i], fmt.Sprintf("Lesson %d.%d", unitNum, curric.LessonNum[i]), curric.LessonDescr[i]),
			}
			unitTitle := fmt.Sprintf("Unit %d", unitNum)
			if unitNum < 0 {
				unitTitle = curric.UnitDescr[i]
			}
			currUnit = domain.NewUnit(unitNum, unitTitle, curric.UnitDescr[i], lessons)
		} else {
			// if it's the same as current unit, create a new lesson and increment the count
			unitLessonCounter[unitNum]++
			lesson := domain.NewLesson(curric.LessonNum[i], fmt.Sprintf("Lesson %d.%d", unitNum, curric.LessonNum[i]), curric.LessonDescr[i])
			currUnit.Lessons = append(currUnit.Lessons, lesson)
		}
	}
	curric.Course.Units = append(curric.Course.Units, currUnit) // append the final unit
	return curric.Course

}
