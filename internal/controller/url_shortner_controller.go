package controller

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/pratikpjain/go-url-shortner-app/internal/service"
	"github.com/pratikpjain/go-url-shortner-app/internal/utils"
	"gorm.io/gorm"
)

type UrlShortnerController struct {
	UrlShortnerService *service.UrlShortnerService
}

func NewUrlShortnerController() *UrlShortnerController {
	return &UrlShortnerController{
		UrlShortnerService: service.NewUrlShortnerService(),
	}
}

func (us *UrlShortnerController) GenerateShortURL(c *gin.Context) {
	requestID := c.Request.Header.Get("X-Request-ID")

	if requestID == "" {
		requestID, _ = utils.GenerateRandomString(16)
	}

	var requestBody map[string]string
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
		})
		return
	}

	longUrl := requestBody["long_url"]
	shortUrl, err := us.UrlShortnerService.GenerateShortUrl(c, requestID, longUrl)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error while generating short url",
		})
		return
	}
	c.JSON(http.StatusOK, shortUrl)
}

func (us *UrlShortnerController) RedirectToLongUrl(c *gin.Context) {

	shortUrl := c.Param("short_url")
	shortUrl = strings.Trim(shortUrl, "/")
	longUrl, err := us.UrlShortnerService.GetLongUrlByShortUrl(c, shortUrl)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "no url is mapped to this short url",
			})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error while getting long url",
		})
		return
	}
	c.Redirect(http.StatusPermanentRedirect, longUrl)

}
