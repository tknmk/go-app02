package main

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// map for getting base.html
var templates map[string]*template.Template

// Interface For Template
type Template struct {
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	return templates[name].ExecuteTemplate(w, name, data)
}

func main() {

	// create echo instance
	e := echo.New()

	// set Renderer
	t := &Template{}
	e.Renderer = t

	// set middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// set Static Files
	e.Static("/public/css/", "./public/css")
	e.Static("/public/js/", "./public/js/")
	e.Static("/public/img/", "./public/img/")

	e.GET("/", HandleTopGet)

	e.Logger.Fatal(e.Start(":8080"))
}

func init() {
	loadTemplates()
}

func loadTemplates() {
	var baseTemplate = "templates/main/base.html"
	templates = make(map[string]*template.Template)

	templates["Topmain"] = template.Must(template.ParseFiles(
		baseTemplate, "templates/top/main.html", "templates/main/header.html", "templates/main/footer.html"))
}

func HandleTopGet(c echo.Context) error {
	return c.Render(http.StatusOK, "Topmain", "OOOOOOKKKKKKKK")
}
