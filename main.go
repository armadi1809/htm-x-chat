package main

import (
	"html/template"
	"log"
)

type Config struct {
	t        *Template
	chatRoom *room
}

func main() {

	app := &Config{
		t:        &Template{templates: template.Must(template.ParseGlob("templates/*.html"))},
		chatRoom: NewRoom(),
	}
	e := app.routes()
	go app.chatRoom.run()

	log.Println("Spinning up the server...")
	e.Logger.Fatal(e.Start(":3000"))
}
