package main

import (
	"database/sql"
	"html/template"
	"log"

	"github.com/armadi1809/htm-x-chat/models"
	_ "github.com/lib/pq"
)

type Config struct {
	t        *Template
	chatRoom *room
	userDb   *models.UserDb
}

const dbSource = "postgresql://root:zizox18099@localhost:5432/Chat?sslmode=disable"

func main() {

	db, err := openDb()

	if err != nil {
		log.Fatal("Can't Connect to the database")
	}
	app := &Config{
		t:        &Template{templates: template.Must(template.ParseGlob("templates/*.html"))},
		chatRoom: NewRoom(),
		userDb:   &models.UserDb{DB: db},
	}
	e := app.routes()
	go app.chatRoom.run()

	log.Println("Spinning up the server...")
	e.Logger.Fatal(e.Start(":3000"))
}

func openDb() (*sql.DB, error) {
	db, err := sql.Open("postgres", dbSource)

	if err != nil {
		return nil, err
	}
	return db, nil
}
