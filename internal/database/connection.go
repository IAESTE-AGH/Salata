package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Init(connString string) {
	db, err := sql.Open("postgres", connString)
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal("Cannot connect to database:", err)
	}

	DB = db
	log.Println("Connected to database")
}
