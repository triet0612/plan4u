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
}

func (a *AppHandler) Home(c echo.Context) error {
	return c.Render(200, "index.html", nil)
}

func (a *AppHandler) Kanban(c echo.Context) error {
	// rows, err := a.Db.Query("SELECT colname, color, note, col_limit FROM kanban_column;")
	return nil
}
