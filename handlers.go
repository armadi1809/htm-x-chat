package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (app *Config) HomePage(c echo.Context) error {

	return c.Render(http.StatusOK, "Home", nil)
}

func (app *Config) connectOnRequest(c echo.Context) error {
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}

	client := &client{
		room:   app.chatRoom,
		socket: ws,
		send:   make(chan string, 256),
	}

	app.chatRoom.join <- client
	defer func() { app.chatRoom.leave <- client }()
	go client.write()
	client.read()

	return nil
}
