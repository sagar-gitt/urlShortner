package models

type URL struct {
	ID          int
	OriginalURL string
	ShortCode   string
	CreatedAt   string
	ExpiresAt   string
	HitCount    int
}
