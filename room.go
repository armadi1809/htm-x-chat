package main

import (
	"github.com/labstack/echo/v4"
)

type room struct {
	forward chan string
	clients map[*client]bool
}

func (r *room) connect(c echo.Context) error {
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}

	client := &client{
		room:   r,
		socket: ws,
		send:   make(chan string, 256),
	}
	r.clients[client] = true
	return nil
}
