package service

import (
	"io"
	"text/template"

	"github.com/labstack/echo/v4"
	g "github.com/maragudk/gomponents"
	c "github.com/maragudk/gomponents/components"
	. "github.com/maragudk/gomponents/html"
)

type TmplMan struct {
	templates *template.Template
}

func (t *TmplMan) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func InitTmplMan() *TmplMan {
	t := &TmplMan{
		templates: template.Must(template.ParseGlob("./view/*.html")),
	}
	return t
}

func Page(title, path string, body g.Node) g.Node {
	return c.HTML5(
		c.HTML5Props{
			Title:    title,
			Language: "en",
			Head:     []g.Node{Script(Src(""))},
			Body:     []g.Node{},
		},
	)
}
