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
	CrudderHandler
}
type courseInstanceHandler struct {
	service services.CourseInstanceService
	e       *echo.Echo
}

func NewCourseInstanceHandler(service services.CourseInstanceService, e *echo.Echo) CourseInstanceHandler {
	return courseInstanceHandler{service: service, e: e}
}

const (
	CourseInstanceHandlerCreate      RouteName = "POST: /course-instances"
	CourseInstanceHandlerList        RouteName = "GET: /course-instances"
	CourseInstanceHandlerReadFromCSV RouteName = "GET: /course-instances/csv"
	CourseInstanceHandlerUpdate      RouteName = "PUT: /course-instances/:id"
	CourseInstanceHandlerDelete      RouteName = "DELETE: /course-instances/:id"
)

const (
	CourseID ParamName = "course-id"
	TermID   ParamName = "term-id"
)

func (h courseInstanceHandler) Mount() {
	nameRoute(h.e.POST("/course-instances", h.Create), CourseInstanceHandlerCreate)
	nameRoute(h.e.GET("/course-instances", h.List), CourseInstanceHandlerList)
	nameRoute(h.e.GET("/course-instances/csv", h.ReadFromCSV), CourseInstanceHandlerReadFromCSV)
	nameRoute(h.e.PUT("/course-instances/:id", h.Update), CourseInstanceHandlerUpdate)
	nameRoute(h.e.DELETE("/course-instances/:id", h.Delete), CourseInstanceHandlerDelete)
}

func (h courseInstanceHandler) Create(c echo.Context) error {
	c.String(http.StatusOK, "not implemented")
	return nil
}

func (h courseInstanceHandler) List(c echo.Context) error {
	courseID, err := getIntegerParam(c, CourseID)
	if err != nil {
		return nil
	}
	termID, err := getIntegerParam(c, TermID)
	if err != nil {
		return err
	}
	courses, err := h.service.FetchOne(courseID, termID)
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
