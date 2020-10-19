package config

import (
	"database/sql"
	"log"
)

func Connect() *sql.DB {
	db, err := sql.Open("mysql", "root:P@ssw0rd@tcp(localhost:3306)/webeai")

	if err != nil {
		log.Fatalf("Database error: %v", err)
	} else {
		log.Println("Database Successs")
	}

	return db
}
