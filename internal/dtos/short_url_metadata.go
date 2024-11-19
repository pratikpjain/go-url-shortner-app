package dtos

type ShortUrlMetadata struct {
	ID              int64  `json:"id"`
	UserIPAddress   string `json:"user_ip_address"`
	Longitude       string `json:"longitude"`
	Latitude        string `json:"latitude"`
	RequestedAtTime string `json:"requested_at_time"`
	ShortenedURLID  int64  `json:"shortened_url_id"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
}
