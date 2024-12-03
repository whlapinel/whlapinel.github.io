package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handler interface {
	Create(c echo.Context) error
	List(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
}

type CourseHandler interface {
	Handler
}
type courseHandler struct {
	service CourseService
}

func NewCourseHandler(service CourseService) CourseHandler {
	return courseHandler{service: service}
}

func (h courseHandler) List(c echo.Context) error {
	_, err := h.service.All()
	if err != nil {
		return err
	}
	c.String(http.StatusOK, "got your courses right here!")
	return nil
}

func (h courseHandler) Delete(c echo.Context) error {
	c.String(http.StatusOK, "not implemented")
	return nil
}

func (h courseHandler) Create(c echo.Context) error {
	c.String(http.StatusOK, "not implemented")
	return nil
}

func (h courseHandler) Update(c echo.Context) error {
	c.String(http.StatusOK, "not implemented")
	return nil
}
