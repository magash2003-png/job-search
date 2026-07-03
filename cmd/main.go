package main

import (
	"log"
	"net/http"

	"job-search/internal/database"
	"job-search/internal/routes"
)

func main() {
	err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	routes.SetupRoutes()

	log.Println("Server started on :8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}