package endpoints

import (
	"io"
	"text/template"

	"github.com/labstack/echo/v4"
)

type TemplateRenderer struct {
	templates *template.Template
}

func NewTemplateRenderer(tmpls *template.Template) TemplateRenderer {
	return TemplateRenderer{tmpls}
}

func (t TemplateRenderer) Render(w io.Writer, name string, data any, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
