package repositories

import (
	"github.com/pratikpjain/go-url-shortner-app/internal/database"
	"github.com/pratikpjain/go-url-shortner-app/internal/dtos"
	"github.com/pratikpjain/go-url-shortner-app/internal/models"
	"gorm.io/gorm"
)

type LongURLRepository struct {
	db *gorm.DB
}

func NewLongURLRepository() *LongURLRepository {
	return &LongURLRepository{
		db: database.DB,
	}
}

func (lr *LongURLRepository) CheckIfLongUrlExists(longURL string) (dtos.LongURL, error) {

	var longURLData dtos.LongURL
	err := lr.db.Model(&models.LongURL{}).Where("original_url = ?", longURL).First(&longURLData).Error
	if err != nil {
		return dtos.LongURL{}, err
	}
	return longURLData, nil

}

func (lr *LongURLRepository) SaveLongUrl(longURL string) (dtos.LongURL, error) {

	var newLongURL dtos.LongURL
	longURLData := models.LongURL{
		OriginalURL: longURL,
	}
	err := lr.db.Model(&models.LongURL{}).Create(&longURLData).Error
	if err != nil {
		return newLongURL, err
	}
	lr.db.Model(&models.LongURL{}).Where("original_url = ?", longURL).First(&newLongURL)
	return newLongURL, nil

}

func (lr *LongURLRepository) GetLongUrlByLongUrlID(longUrlID int64) (string, error) {

	var longUrl string
	err := lr.db.Model(&models.LongURL{}).Select("original_url").Where("id = ?", longUrlID).First(&longUrl).Error
	if err != nil {
		return "", err
	}
	return longUrl, nil

}
