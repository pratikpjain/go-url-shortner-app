package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/pratikpjain/go-url-shortner-app/internal/config"
	"github.com/pratikpjain/go-url-shortner-app/internal/database"
	"github.com/pratikpjain/go-url-shortner-app/internal/routes"
)

func main() {

	// Load environment variables
	config := config.LoadConfig()

	// Connect to database
	database.ConnectDB()
	defer database.CloseDB()
	log.Println("Successfully connected to MySQL database")

	// Setup routes
	router := gin.Default()
	routes.RegisterRoutes(router, database.DB)
	log.Println("Successfully registered routes")

	// Start server
	router.Run(fmt.Sprintf(":%s", config.AppPort))
	log.Println("Server is running on port: ", config.AppPort)

}
