package main

import (
	"context"
	"database/sql"
	"embed"
	"gh_static_portfolio/cmd/course_manager/handlers"
	"gh_static_portfolio/cmd/course_manager/services"
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
	e := echo.New()
	queries = database.New(db)
	courseRepo := data.NewCourseSOARepo(queries)
	courseService := services.NewCourseService(courseRepo)
	courseHandler := handlers.NewCourseHandler(courseService, e)
	termRepo := data.NewTermRepo(queries)
	instanceRepo := data.NewInstanceRepo(queries)
	instanceService := services.NewcourseInstanceService(instanceRepo, courseRepo, termRepo)
	instanceHandler := handlers.NewCourseInstanceHandler(instanceService, e)
	mainMenuHandler := handlers.NewMainMenuHandler(e)
	handlers := []handlers.Handler{
		courseHandler, instanceHandler, mainMenuHandler,
	}
	MountHandlers(handlers)
	fs := echo.MustSubFS(embeddedFiles, "assets")
	e.StaticFS("/static", fs)
	e.Logger.Fatal(e.Start(":1323"))
}

func RenderTempl(component templ.Component, c echo.Context, statusCode int) error {
	c.Response().WriteHeader(statusCode)
	return component.Render(c.Request().Context(), c.Response().Writer)
}

func MountHandlers(handlers []handlers.Handler) {
	for _, handler := range handlers {
		handler.Mount()
	}

}
