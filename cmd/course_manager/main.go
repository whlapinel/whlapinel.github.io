package main

import (
	"context"
	"database/sql"
	"embed"
	"gh_static_portfolio/cmd/course_manager/database"
	"gh_static_portfolio/cmd/course_manager/templates"
	"gh_static_portfolio/cmd/data"
	"gh_static_portfolio/cmd/domain"
	"log"

	"github.com/a-h/templ"

	_ "embed"

	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
)

//go:embed schema.sql
var ddl string

//go:embed assets/**
var embeddedFiles embed.FS

func main() {
	var queries *database.Queries
	ctx := context.Background()
	db, err := sql.Open("sqlite3", "course_manager.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	if _, err := db.ExecContext(ctx, ddl); err != nil {
		log.Fatal(err)
	}
	// Enable foreign keys
	_, err = db.ExecContext(ctx, "PRAGMA foreign_keys = ON;")
	if err != nil {
		log.Fatal("Failed to enable foreign keys:", err)
	}

	// Check if foreign keys are enabled
	var foreignKeysEnabled int
	err = db.QueryRow("PRAGMA foreign_keys;").Scan(&foreignKeysEnabled)
	if err != nil {
		log.Fatal("Failed to check foreign_keys status:", err)
	}

	log.Println("Foreign keys enabled:", foreignKeysEnabled)
	queries = database.New(db)
	curricula, err := loadCurriculaFromCSV()
	if err != nil {
		log.Fatal(err)
	}
	for _, curriculum := range curricula {
		_, err := saveCurriculum(curriculum, ctx, queries)
		if err != nil {
			log.Fatal(err)
		}

	}

	e := echo.New()
	fs := echo.MustSubFS(embeddedFiles, "assets")
	e.StaticFS("/static", fs)
	e.GET("/", func(c echo.Context) error {
		component := templates.ManagerMainMenu()
		return RenderTempl(component, c, 200)
	})
	e.GET("/courses", func(c echo.Context) error {
		courseRows, err := queries.GetCourses(ctx)
		if err != nil {
			log.Fatal(err)
		}
		courseSOA, err := CourseRowsToCourseSOA(courseRows)
		if err != nil {
			log.Fatal(err)
		}
		component := templates.ManageCourseComponent(courseSOA)
		return RenderTempl(component, c, 200)
	})
	e.Logger.Fatal(e.Start(":1323"))
}

func loadCurriculaFromCSV() ([]*domain.CourseSOA, error) {
	curriculum1, err := data.ImportCurriculumFromCSV("Python I Programming Honors")
	if err != nil {
		return nil, err

	}
	curriculum2, err := data.ImportCurriculumFromCSV("Python II Programming Honors")
	if err != nil {
		return nil, err

	}
	return []*domain.CourseSOA{
		curriculum1,
		curriculum2,
	}, nil

}

func saveCurriculum(curric *domain.CourseSOA, ctx context.Context, queries *database.Queries) (int64, error) {
	course, err := queries.SaveCourse(ctx, database.SaveCourseParams{Name: curric.Course.Title()})
	if err != nil {
		return 0, err
	}
	var currUnit *database.Unit
	var currLesson *database.Lesson
	var currDayNumber *database.DayNumber
	for i, currDay := range curric.Day {
		if currUnit == nil || int64(curric.UnitNum[i]) != currUnit.Number {
			log.Println("creating and saving unit")
			currUnit = &database.Unit{
				Number:   int64(curric.UnitNum[i]),
				Name:     curric.UnitName[i],
				CourseID: course.ID,
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
		if currLesson == nil || int64(curric.LessonNum[i]) != currLesson.Number {
			log.Println("Lesson Name: ", curric.LessonName[i])
			currLesson = &database.Lesson{
				Number: int64(curric.LessonNum[i]),
				Name: sql.NullString{
					String: curric.LessonName[i],
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
		if currDayNumber == nil || currDay != int(currDayNumber.Day) {
			currDayNumber = &database.DayNumber{
				Day:      int64(currDay),
				LessonID: currLesson.ID,
			}
			*currDayNumber, err = queries.SaveDayNumber(ctx, database.SaveDayNumberParams{
				LessonID: currLesson.ID,
				Day:      currDayNumber.Day,
			})
			if err != nil {
				log.Fatal(err)
			}

		}

	}
	return course.ID, nil

}

func RenderTempl(component templ.Component, c echo.Context, statusCode int) error {
	c.Response().WriteHeader(statusCode)
	return component.Render(c.Request().Context(), c.Response().Writer)
}

func CourseRowsToCourseSOA(rows []database.GetCoursesRow) ([]*domain.CourseSOA, error) {
	var curricula = []*domain.CourseSOA{}
	curricMap := make(map[int]*domain.CourseSOA)
	var currID int64 = 0
	for _, row := range rows {
		currID = row.CourseID
		curriculum := curricMap[int(currID)]
		if curriculum == nil {
			curriculum = &domain.CourseSOA{}
			curriculum.Course = domain.NewCourse(row.CourseName, "", []domain.Unit{}, "")
			curriculum.Course.ID = int(currID)

		}
		curriculum.Day = append(curriculum.Day, int(row.DayNumber))
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
