package handlers

import (
	"gh_static_portfolio/cmd/course_manager/services"
	"gh_static_portfolio/cmd/course_manager/templates"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CourseOOPHandler interface {
	Handler
	CrudderHandler
}
type courseOOPHandler struct {
	service services.CourseInstanceService
	e       *echo.Echo
}

func NewCourseOOPHandler(service services.CourseInstanceService, e *echo.Echo) CourseOOPHandler {
	return courseOOPHandler{service: service, e: e}
}

const (
	CourseOOPHandlerCreate      RouteName = "POST: /course-instances"
	CourseOOPHandlerList        RouteName = "GET: /course-instances"
	CourseOOPHandlerReadFromCSV RouteName = "GET: /course-instances/csv"
	CourseOOPHandlerUpdate      RouteName = "PUT: /course-instances/:id"
	CourseOOPHandlerDelete      RouteName = "DELETE: /course-instances/:id"
)


func (h courseOOPHandler) Mount() {
	nameRoute(h.e.POST("/course-instances", h.Create), CourseOOPHandlerCreate)
	nameRoute(h.e.GET("/course-instances", h.List), CourseOOPHandlerList)
	nameRoute(h.e.GET("/course-instances/csv", h.ReadFromCSV), CourseOOPHandlerReadFromCSV)
	nameRoute(h.e.PUT("/course-instances/:id", h.Update), CourseOOPHandlerUpdate)
	nameRoute(h.e.DELETE("/course-instances/:id", h.Delete), CourseOOPHandlerDelete)
}

func (h courseOOPHandler) Create(c echo.Context) error {
	c.String(http.StatusOK, "not implemented")
	return nil
}

func (h courseOOPHandler) List(c echo.Context) error {
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

func (h courseOOPHandler) ReadFromCSV(c echo.Context) error {
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

func (h courseOOPHandler) Delete(c echo.Context) error {
	c.String(http.StatusOK, "not implemented")
	return nil
}

func (h courseOOPHandler) Update(c echo.Context) error {
	c.String(http.StatusOK, "not implemented")
	return nil
}
