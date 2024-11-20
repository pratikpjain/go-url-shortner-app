package repositories

import (
	"github.com/pratikpjain/go-url-shortner-app/internal/database"
	"github.com/pratikpjain/go-url-shortner-app/internal/dtos"
	"github.com/pratikpjain/go-url-shortner-app/internal/models"
	"gorm.io/gorm"
)

type ShortenedURLRepository struct {
	db *gorm.DB
}

func NewShortenedURLRepository() *ShortenedURLRepository {
	return &ShortenedURLRepository{
		db: database.DB,
	}
}

func (s *ShortenedURLRepository) SaveShortenedUrl(shortenedUrl string, longUrlID int64) error {

	shortenedUrlData := models.ShortenedURL{
		LongURLID: longUrlID,
		ShortURL:  shortenedUrl,
	}
	err := s.db.Create(&shortenedUrlData).Error
	if err != nil {
		return err
	}
	return nil

}

func (s *ShortenedURLRepository) CheckIfShortenedUrlExists(shortenedUrl string) (dtos.ShortenedURL, error) {
	var shortUrlData dtos.ShortenedURL
	err := s.db.Model(&models.ShortenedURL{}).Where("short_url = ?", shortenedUrl).First(&shortUrlData).Error
	if err != nil {
		return dtos.ShortenedURL{}, err
	}
	return shortUrlData, nil
}
