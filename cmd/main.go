package main

import (
	"log"
	"net/http"
	"os"

	"job-search/internal/database"
	"job-search/internal/routes"
)

func main() {

	err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	routes.SetupRoutes()

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Server started on :" + port)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
