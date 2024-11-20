package models

type ShortUrlMetadata struct {
	ID              int64  `gorm:"primaryKey" json:"id"`
	UserIPAddress   string `gorm:"type:varchar(255);not null" json:"user_ip_address"`
	Longitude       string `gorm:"type:varchar(255);not null" json:"longitude"`
	Latitude        string `gorm:"type:varchar(255);not null" json:"latitude"`
	RequestedAtTime string `gorm:"type:timestamp;not null" json:"requested_at_time"`
	ShortenedURLID  int64  `gorm:"not null" json:"shortened_url_id"`
	CreatedAt       string `gorm:"type:timestamp;default:now()" json:"created_at"`
	UpdatedAt       string `gorm:"type:timestamp;default:now() ON UPDATE now()" json:"updated_at"`
}

func (ShortUrlMetadata) TableName() string {
	return "short_url_metadata"
}
