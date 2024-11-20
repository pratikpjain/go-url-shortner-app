package service

import (
	"context"
	"errors"
	"log"

	"github.com/pratikpjain/go-url-shortner-app/internal/dtos"
	"github.com/pratikpjain/go-url-shortner-app/internal/repositories"
	"github.com/pratikpjain/go-url-shortner-app/internal/utils"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type UrlShortnerService struct {
	ShortenedURLRepository *repositories.ShortenedURLRepository
	LongURLRepository      *repositories.LongURLRepository
}

func NewUrlShortnerService() *UrlShortnerService {
	return &UrlShortnerService{
		ShortenedURLRepository: repositories.NewShortenedURLRepository(),
		LongURLRepository:      repositories.NewLongURLRepository(),
	}
}

func (us *UrlShortnerService) GetLongUrlByShortUrl(c context.Context, shortUrl string) (string, error) {

	var longUrl string
	shortUrlData, err := us.ShortenedURLRepository.CheckIfShortenedUrlExists(shortUrl)
	if err != nil {
		log.Println("Error while checking if short url exists: ", err)
		return "", err
	}

	longUrl, err = us.LongURLRepository.GetLongUrlByLongUrlID(shortUrlData.LongURLID)
	if err != nil {
		log.Println("Error while getting long url: ", err)
		return "", err
	}
	return longUrl, nil
}

func (us *UrlShortnerService) GenerateShortUrl(c context.Context, requestID, longUrl string) (interface{}, error) {

	if longUrl == "" {
		return nil, errors.New("long url is required")
	}

	var longUrlData dtos.LongURL
	// Check if long url already exists in database
	longUrlData, err := us.LongURLRepository.CheckIfLongUrlExists(longUrl)
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Println("Error while checking if long url exists: ", err)
		return nil, err
	}
	// If long url does not exist in database, save it
	if err == gorm.ErrRecordNotFound {
		longUrlData, err = us.LongURLRepository.SaveLongUrl(longUrl)
		if err != nil {
			log.Println("Error while saving long url: ", err)
			return nil, err
		}
	}
	// generate short url
	shortUrl, err := generateShortUrl()
	if err != nil {
		log.Println("Error while generating short url: ", err)
		return nil, err
	}

	// save short url
	err = us.ShortenedURLRepository.SaveShortenedUrl(shortUrl, longUrlData.ID)
	if err != nil {
		log.Println("Error while saving short url: ", err)
		return nil, err
	}

	shortUrl = viper.GetString("BASE_URL") + "url/" + shortUrl

	return map[string]interface{}{
		"short_url": shortUrl,
	}, nil
}

func generateShortUrl() (string, error) {
	str, err := utils.GenerateRandomString(6)
	if err != nil {
		return "", err
	}
	return str, nil
}
