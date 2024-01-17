package service

import (
	"io"
	"text/template"

	"github.com/labstack/echo/v4"
)

type TmplMan struct {
	templates *template.Template
}

func (t *TmplMan) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func InitTmplMan() *TmplMan {
	return &TmplMan{
		templates: template.Must(template.ParseGlob("./view/*.html")),
	}
}
