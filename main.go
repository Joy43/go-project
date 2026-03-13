package main

import (
	"go-jwt-auth/config"
	"go-jwt-auth/routes"
	"log"
	"net/http"

	_ "go-jwt-auth/docs" // swagger docs
)

// @title Go JWT Auth API
// @version 1.0
// @description JWT Authentication API using Go, Gorilla Mux, and MySQL
// @host localhost:4000
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer <token>" to authenticate.
func main() {

	// Connect to database
	config.ConnectDatabase()
	defer config.CloseDatabase()

	// Setup routes
	router := routes.SetRoutes()

	// Start server
	log.Println("Server starting on :4000")
	log.Println("Server URL: http://localhost:4000")
	log.Println("Swagger Docs: http://localhost:4000/docs")

	if err := http.ListenAndServe(":4000", router); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}