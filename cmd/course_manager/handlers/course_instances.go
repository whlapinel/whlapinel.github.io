package handlers

import (
	"gh_static_portfolio/cmd/course_manager/services"
	"gh_static_portfolio/cmd/course_manager/templates"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CourseInstanceHandler interface {
	Handler
	ReadFromCSV(c echo.Context) error
}
type courseInstanceHandler struct {
	service services.CourseInstanceService
	e       *echo.Echo
}

func NewCourseInstanceHandler(service services.CourseInstanceService, e *echo.Echo) CourseInstanceHandler {
	return courseInstanceHandler{service: service, e: e}
}

const (
	CourseInstanceHandlerCreate      = "POST: /course-instances"
	CourseInstanceHandlerList        = "GET: /course-instances"
	CourseInstanceHandlerReadFromCSV = "GET: /course-instances/csv"
	CourseInstanceHandlerUpdate      = "PUT: /course-instances/:id"
	CourseInstanceHandlerDelete      = "DELETE: /course-instances/:id"
)

func (h courseInstanceHandler) Mount() {
	h.e.POST("/course-instances", h.Create).Name = CourseInstanceHandlerCreate
	h.e.GET("/course-instances", h.List).Name = CourseInstanceHandlerList
	h.e.GET("/course-instances/csv", h.ReadFromCSV).Name = CourseInstanceHandlerReadFromCSV
	h.e.PUT("/course-instances/:id", h.Update).Name = CourseInstanceHandlerUpdate
	h.e.DELETE("/course-instances/:id", h.Delete).Name = CourseInstanceHandlerDelete
}

func (h courseInstanceHandler) Create(c echo.Context) error {
	c.String(http.StatusOK, "not implemented")
	return nil
}
func (h courseInstanceHandler) List(c echo.Context) error {
	courses, err := h.service.All()
	if err != nil {
		return err
	}
	component := templates.ManageInstanceComponent(courses)
	return RenderTempl(component, c, 200)
}

func (h courseInstanceHandler) ReadFromCSV(c echo.Context) error {
	instances, err := h.service.ReadFromCSV()
	if err != nil {
		return err
	}
	for _, instance := range instances {
		err := h.service.Create(instance)
		if err != nil {
			log.Println("error: ", err)
			return err
		}
	}
	component := templates.ManageInstanceComponent(instances)
	return RenderTempl(component, c, 200)

}

func (h courseInstanceHandler) Delete(c echo.Context) error {
	c.String(http.StatusOK, "not implemented")
	return nil
}

func (h courseInstanceHandler) Update(c echo.Context) error {
	c.String(http.StatusOK, "not implemented")
	return nil
}
