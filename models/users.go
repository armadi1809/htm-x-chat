package models

import (
	"database/sql"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID             int    `json:"id"`
	Username       string `json:"username"`
	HashedPassword string `json:"hashedPassword"`
}

type UserDb struct {
	DB *sql.DB
}

func (db *UserDb) CreateUser(username, password string) (int, error) {

	insertStmnt := `INSERT INTO users (username, hashedPassword)
	VALUES ($1, $2)
	RETURNING id`

	hashedPassowrd, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	if err != nil {
		return 0, err
	}
	var id int
	err = db.DB.QueryRow(insertStmnt, username, string(hashedPassowrd)).Scan(&id)
	if err != nil {
		return 0, err
	}
	return int(id), nil
}
