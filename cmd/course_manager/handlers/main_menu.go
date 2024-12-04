package handlers

import (
	"gh_static_portfolio/cmd/course_manager/templates"

	"github.com/labstack/echo/v4"
)

type MainMenuHandler interface {
	Mount(c *echo.Echo)
	ShowMenu(c echo.Context) error
}

type mainMenuHandler struct {
}

func NewMainMenuHandler() MainMenuHandler {
	return mainMenuHandler{}
}

const (
	MainMenuHandlerShowMenu = "GET: /"
)

func (h mainMenuHandler) Mount(e *echo.Echo) {
	e.GET("/", h.ShowMenu).Name = MainMenuHandlerShowMenu
}
func (h mainMenuHandler) ShowMenu(c echo.Context) error {
	component := templates.ManagerMainMenu()
	return RenderTempl(component, c, 200)

}
