package main

import (
	"html/template"
	"io"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Template struct {
	templates *template.Template
}

type Config struct {
	t *Template
}

func main() {

	app := &Config{
		t: &Template{templates: template.Must(template.ParseGlob("templates/*.html"))},
	}

	room := NewRoom()
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Renderer = app.t

	e.Static("/css", "css")
	e.GET("/", app.HomePage)
	e.GET("/ws", room.connectOnRequest)
	go room.run()

	log.Println("Spinning up the server...")

	e.Logger.Fatal(e.Start(":3000"))

}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func (app *Config) HomePage(c echo.Context) error {

	return c.Render(http.StatusOK, "Home", nil)
}
