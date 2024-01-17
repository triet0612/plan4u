package main

import (
	"database/sql"
	"plan4u/config"
	"plan4u/handler"
	"plan4u/service"

	"fyne.io/systray"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./app.db")
	if err != nil {
		panic(err)
	}
	cfg := config.InitConfig()

	tray := service.SystemTray{Cfg: cfg}

	go runServer(db, cfg)
	systray.Run(tray.OnReady, tray.OnExit)
}

func runServer(db *sql.DB, cfg *config.Config) {
	h := handler.AppHandler{Db: db, Cfg: cfg}
	e := echo.New()
	e.Renderer = service.InitTmplMan()

	// e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("assets", "./view/assets")
	h.Mount(e)

	e.Start(cfg.URL + ":" + cfg.PORT)
}
