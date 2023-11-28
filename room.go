package main

import (
	"fmt"

	"github.com/gorilla/websocket"
)

type room struct {
	forward chan string
	clients map[*client]bool
	join    chan *client
	leave   chan *client
}

func NewRoom() *room {
	return &room{
		forward: make(chan string, 256),
		clients: make(map[*client]bool, 256),
		join:    make(chan *client),
		leave:   make(chan *client),
	}
}

var (
	upgrader = websocket.Upgrader{}
)

func (r *room) run() error {
	for {
		select {
		case client := <-r.join:
			r.clients[client] = true
			fmt.Println("A new client joined")

		case client := <-r.leave:
			delete(r.clients, client)
			close(client.send)
			fmt.Println("A client left the chat room")

		case msg := <-r.forward:
			for client := range r.clients {
				client.send <- msg
				fmt.Println("A new message was sent in the chat. Forwarding it")
			}
		}
	}
}
