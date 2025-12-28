package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func Connect() *sql.DB {
	host := "localhost"
	port := 5432
	user := "postgres"
	password := "admin"
	dbname := "backend_golang"

	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("FAILED CONNECT DB: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("FAILED PING DB: %v", err)
	}

	log.Println("POSTGRES DATABASE CONNECTED")
	return db
}
