package main

import (
	"github.com/pratikpjain/go-url-shortner-app/internal/config"
	"github.com/pratikpjain/go-url-shortner-app/internal/database"
)

func main() {

	_ = config.LoadConfig()

	// database connection
	db, err := database.ConnectDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

}
