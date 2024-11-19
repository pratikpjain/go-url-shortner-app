package models

type LongURL struct {
	ID          int64  `gorm:"primaryKey" json:"id"`
	OriginalURL string `gorm:"type:varchar(255);not null" json:"original_url"`
	CreatedAt   string `gorm:"type:timestamp;default:now();" json:"created_at"`
	UpdatedAt   string `gorm:"type:timestamp;default:now(); ON UPDATE now();" json:"updated_at"`
}

func (LongURL) TableName() string {
	return "long_urls"
}
