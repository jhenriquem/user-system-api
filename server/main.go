package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"github.com/jhenriquem/user-system-api/database"
	"github.com/jhenriquem/user-system-api/middleware"
	"github.com/jhenriquem/user-system-api/routes"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Middleares
	routes.Router.Use(middleware.JsonContentTypeMiddleware)
	routes.Router.Use(middleware.LoggingMiddleware)

	// Routes List
	routes.InitRoutes()

	// Database
	database.InitDatabase()
	defer database.DB.Close()

	// Server Configuration
	server := &http.Server{
		Addr:    os.Getenv("PORT"),
		Handler: routes.Router,
	}

	fmt.Printf("Online Server -> http://localhost%s \n", os.Getenv("PORT"))

	server.ListenAndServe()
}
