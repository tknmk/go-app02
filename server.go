package main

import (
	"html/templaate"
	"net/http"

	"github.com/labstack/echo"
)

// Interface For Template
type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {

	// create echo instance
	e := echo.New()

	// set Renderer
	t := &Template{
		templates: template.Must(template.ParseGlob("public/views/*.html")),
	}

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World!")
	})
	e.Logger.Fatal(e.Start(":8080"))
}
