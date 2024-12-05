package handlers

import (
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

func RenderTempl(component templ.Component, c echo.Context, statusCode int) error {
	c.Response().WriteHeader(statusCode)
	return component.Render(c.Request().Context(), c.Response().Writer)
}
