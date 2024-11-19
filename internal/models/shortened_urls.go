package models

type ShortenedURL struct {
	ID        int64  `gorm:"primaryKey" json:"id"`
	LongURLID int64  `gorm:"not null" json:"long_url_id"`
	ShortURL  string `gorm:"type:varchar(255);not null" json:"short_url"`
	CreatedAt string `gorm:"type:timestamp;default:now()" json:"created_at"`
	UpdatedAt string `gorm:"type:timestamp;default:now() ON UPDATE now()" json:"updated_at"`
}

func (ShortenedURL) TableName() string {
	return "shortened_urls"
}
