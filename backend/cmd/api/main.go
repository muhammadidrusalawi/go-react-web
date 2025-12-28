package main

import (
	"log"
	"net/http"

	"backend/internal/config"
	appHttp "backend/internal/http"
)

func main() {
	log.Println("Starting API service...")

	db := config.Connect()
	defer db.Close()

	router := appHttp.NewRouter()

	log.Println("Server running at :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
