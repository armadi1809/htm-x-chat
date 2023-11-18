package main

import (
	"html/template"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
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
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Renderer = app.t
	e.GET("/", app.HomePage)
	e.GET("/ws", socketConnection)
	log.Println("Spinning up the server...")

	e.Logger.Fatal(e.Start(":3000"))

}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func (app *Config) HomePage(c echo.Context) error {

	return c.Render(http.StatusOK, "Home", nil)
}

var (
	upgrader = websocket.Upgrader{}
)

func socketConnection(c echo.Context) error {
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()

	for {
		// Write
		err := ws.WriteMessage(websocket.TextMessage, []byte("Hello, Client!"))
		if err != nil {
			c.Logger().Error(err)
		}

		// Read
		_, msg, err := ws.ReadMessage()
		if err != nil {
			c.Logger().Error(err)
		}
		log.Printf("%s\n", msg)
	}
}
