package handler

import (
	"database/sql"
	"plan4u/config"

	"github.com/labstack/echo/v4"
)

type AppHandler struct {
	Db  *sql.DB
	Cfg *config.Config
}

func (a *AppHandler) Mount(e *echo.Echo) {
	e.GET("/", a.Home)
	e.GET("/kanban", a.Kanban)
	e.GET("/kanban/detail", a.KanbanDetail)
	e.GET("/kanban/create", a.KanbanCreateForm)
	e.POST("/kanban/create", a.KanbanCreate)
}

func (a *AppHandler) Home(c echo.Context) error {
	return c.Render(200, "index.html", nil)
}
