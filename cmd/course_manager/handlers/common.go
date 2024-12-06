package handlers

import (
	"strconv"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

type Handler interface {
	Mount()
}

type CrudderHandler interface {
	Create(c echo.Context) error
	List(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
}

type RouteName string
type ParamName string

func nameRoute(route *echo.Route, name RouteName) {
	route.Name = string(name)
}

func getIntegerParam(c echo.Context, name ParamName) (int, error) {
	num, err := strconv.Atoi(c.Param(string(name)))
	if err != nil {
		return 0, err
	}
	return num, nil
}

func RenderTempl(component templ.Component, c echo.Context, statusCode int) error {
	c.Response().WriteHeader(statusCode)
	return component.Render(c.Request().Context(), c.Response().Writer)
}
