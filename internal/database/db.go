package database

import (
	"errors"
	"fmt"

	"github.com/pratikpjain/go-url-shortner-app/internal/models"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	// Construct DSN (Data Source Name)
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8mb4,utf8",
		viper.GetString("DB_USER"),
		viper.GetString("DB_PASSWORD"),
		viper.GetString("DB_HOST"),
		viper.GetString("DB_PORT"),
		viper.GetString("DB_NAME"),
	)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	err = DB.AutoMigrate(&models.LongURL{}, &models.ShortenedURL{}, &models.ShortUrlMetadata{})
	if err != nil {
		panic(errors.New("failed to migrate database: " + err.Error()))
	}
	fmt.Println("Successfully connected to MySQL database")

}

func CloseDB() {
	dbSQL, err := DB.DB()
	if err != nil {
		panic(err)
	}
	dbSQL.Close()
}
