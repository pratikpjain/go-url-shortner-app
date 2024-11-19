package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/pratikpjain/go-url-shortner-app/internal/config"
	"github.com/pratikpjain/go-url-shortner-app/internal/database"
	"github.com/pratikpjain/go-url-shortner-app/internal/routes"
)

func main() {

	// Load environment variables
	config := config.LoadConfig()

	// Connect to database
	database.ConnectDB()
	log.Println("Successfully connected to MySQL database")
	// Close database
	defer database.CloseDB()

	// Run server
	router := routes.SetupRoutes(database.DB)

	err := http.ListenAndServe(fmt.Sprintf(":%s", config.AppPort), router)
	if err != nil {
		log.Fatal(errors.New("failed to start server: " + err.Error()))
	}
	log.Println("Server is running on port: ", config.AppPort)

}
