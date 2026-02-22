package main

import (
	"go-jwt-auth/config"
	"go-jwt-auth/routes"
	"log"
	"net/http"
)

func main() {
	// Connect to the database
	config.ConnectDatabase()
	defer config.CloseDatabase()

	// Setup routes
	router := routes.SetRoutes()

	// Start the server
	log.Println("Server starting on :4000")
	// --------- log http url localhost:4000
	log.Println("Visit http://localhost:4000 to access the server")
	if err := http.ListenAndServe(":4000", router); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
