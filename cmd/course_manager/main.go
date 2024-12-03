package main

import (
	"context"
	"database/sql"
	"embed"
	"gh_static_portfolio/cmd/course_manager/templates"
	"gh_static_portfolio/cmd/data"
	"gh_static_portfolio/cmd/data/database"
	"log"

	"github.com/a-h/templ"

	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
)

//go:embed assets/**
var embeddedFiles embed.FS

func main() {
	var queries *database.Queries
	ctx := context.Background()
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	if _, err := db.ExecContext(ctx, database.DDL); err != nil {
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
	coursesRepo := data.NewCourseSOARepo(queries)
	courses, err := coursesRepo.ReadFromCSV()
	if err != nil {
		log.Fatal(err)
	}
	for _, course := range courses {
		log.Println("CourseManager main(): saving course: ", course.Name)
		_, err := data.SaveCourse(course, ctx, queries)
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
		courses, err := coursesRepo.All()
		if err != nil {
			return err
		}
		component := templates.ManageCourseComponent(courses)
		return RenderTempl(component, c, 200)
	})
	e.Logger.Fatal(e.Start(":1323"))
}

func RenderTempl(component templ.Component, c echo.Context, statusCode int) error {
	c.Response().WriteHeader(statusCode)
	return component.Render(c.Request().Context(), c.Response().Writer)
}
