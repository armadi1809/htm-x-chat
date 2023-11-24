package main

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

type client struct {
	socket *websocket.Conn
	send   chan string
	room   *room
}

type hxJson struct {
	ChatMessage string `json:"chatMessage"`
	Headers     any    `json:"Headers"`
}

func (c *client) read() {
	defer c.socket.Close()

	for {
		hxJSON := &hxJson{ChatMessage: "Hello"}
		err := c.socket.ReadJSON(hxJSON)
		if err != nil {
			log.Println(err.Error())
			log.Println("Can't read the message from socket")
		}
		c.room.forward <- hxJSON.ChatMessage
	}
}

func (c *client) write() {
	defer c.socket.Close()

	for msg := range c.send {
		log.Println("Sending message to client")
		partial := fmt.Sprintf("<p id='chat-box' hx-swap-oob=beforeend>%s <br></p> ", msg)
		err := c.socket.WriteMessage(websocket.TextMessage, []byte(partial))
		if err != nil {
			return
		}
	}
}
