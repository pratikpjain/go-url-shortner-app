package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pratikpjain/go-url-shortner-app/internal/controller"
	"gorm.io/gorm"
)

func RegisterRoutes(router *gin.Engine, db *gorm.DB) {

	router.GET("/health-check", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "OK",
		})
	})
	apiV1 := router.Group("/api/v1")

	shortenedUrlHandler := controller.NewUrlShortnerController()

	apiV1.POST("/shorten-url", shortenedUrlHandler.GenerateShortURL)

	router.GET("/url/:short_url", shortenedUrlHandler.RedirectToLongUrl)

}
