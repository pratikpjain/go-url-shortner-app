package dtos

type LongURL struct {
	ID          int64  `json:"id"`
	OriginalURL string `json:"original_url"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
