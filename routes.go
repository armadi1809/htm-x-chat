package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (app *Config) routes() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Renderer = app.t

	e.Static("/css", "css")
	e.GET("/", app.HomePage)
	e.GET("/ws", app.connectOnRequest)
	e.GET("/SignUp", app.SignUp)
	e.POST("/SignUp", app.SignUserUp)

	return e
}
