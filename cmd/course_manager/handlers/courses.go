package handlers

import (
	"gh_static_portfolio/cmd/course_manager/services"
	"gh_static_portfolio/cmd/course_manager/templates"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CourseHandler interface {
	Handler
	ReadFromCSV(c echo.Context) error
}
type courseHandler struct {
	service services.CourseService
	e       *echo.Echo
}

func NewCourseHandler(service services.CourseService, e *echo.Echo) CourseHandler {
	return courseHandler{service: service, e: e}
}

const (
	CourseHandlerCreate      = "POST: /courses"
	CourseHandlerList        = "GET: /courses"
	CourseHandlerReadFromCSV = "GET: /courses/csv"
	CourseHandlerUpdate      = "PUT: /courses/:id"
	CourseHandlerDelete      = "DELETE: /courses/:id"
)

func (h courseHandler) Mount() {
	h.e.POST("/courses", h.Create).Name = CourseHandlerCreate
	h.e.GET("/courses", h.List).Name = CourseHandlerList
	h.e.GET("/courses/csv", h.ReadFromCSV).Name = CourseHandlerReadFromCSV
	h.e.PUT("/courses/:id", h.Update).Name = CourseHandlerUpdate
	h.e.DELETE("/courses/:id", h.Delete).Name = CourseHandlerDelete
}

func (h courseHandler) Create(c echo.Context) error {
	c.String(http.StatusOK, "not implemented")
	return nil
}
func (h courseHandler) List(c echo.Context) error {
	courses, err := h.service.All()
	if err != nil {
		return err
	}
	component := templates.ManageCourseComponent(courses)
	return RenderTempl(component, c, 200)
}

func (h courseHandler) ReadFromCSV(c echo.Context) error {
	courses, err := h.service.ReadFromCSV()
	if err != nil {
		return err
	}
	for _, course := range courses {
		err := h.service.Create(course)
		if err != nil {
			return err
		}
	}
	component := templates.ManageCourseComponent(courses)
	return RenderTempl(component, c, 200)

}

func (h courseHandler) Delete(c echo.Context) error {
	c.String(http.StatusOK, "not implemented")
	return nil
}

func (h courseHandler) Update(c echo.Context) error {
	c.String(http.StatusOK, "not implemented")
	return nil
}
