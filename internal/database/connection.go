package database

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Init(connString string) {
	var db *sql.DB
	var err error

	for i := 0; i < 5; i++ {
		db, err = sql.Open("postgres", connString)
		if err != nil {
			log.Fatal("Error connecting to database:", err)
		}
		if err = db.Ping(); err == nil {
			DB = db
			log.Println("Successfully connected to database")
			break
		} else {
			log.Printf("Database not ready... attempt %d/5. Waiting 2 seconds...\n", i+1)
			time.Sleep(2 * time.Second)
		}
	}
	if err != nil {
		log.Fatal("Cannot connect to database:", err)
	}
}
