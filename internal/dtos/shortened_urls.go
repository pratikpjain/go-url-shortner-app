package dtos

type ShortenedURL struct {
	ID        int64  `json:"id"`
	LongURLID int64  `json:"long_url_id"`
	ShortURL  string `json:"short_url"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
