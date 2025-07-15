package repository

import (
	"database/sql"
	"time"

	"urlShortner/internal/models"
)

type URLRepository struct {
	db *sql.DB
}

func NewURLRepository(db *sql.DB) *URLRepository {
	return &URLRepository{db: db}
}

func (r *URLRepository) SaveURL(originalURL, shortCode string, expiresAt time.Time) error {
	_, err := r.db.Exec(`INSERT INTO urls (original_url, short_code, expires_at) VALUES ($1, $2, $3)`, originalURL, shortCode, expiresAt)
	return err
}

func (r *URLRepository) GetURLByCode(code string) (*models.URL, error) {
	row := r.db.QueryRow(`SELECT id, original_url, short_code, created_at, expires_at, hit_count FROM urls WHERE short_code=$1`, code)
	var u models.URL
	err := row.Scan(&u.ID, &u.OriginalURL, &u.ShortCode, &u.CreatedAt, &u.ExpiresAt, &u.HitCount)
	return &u, err
}

func (r *URLRepository) IncrementHitCount(code string) {
	r.db.Exec(`UPDATE urls SET hit_count = hit_count + 1 WHERE short_code=$1`, code)
}

func (r *URLRepository) GetExpiredURLs() ([]string, error) {
	rows, err := r.db.Query(`SELECT short_code FROM urls WHERE expires_at < NOW()`)
	if err != nil {
		return nil, err
	}
	var codes []string
	for rows.Next() {
		var c string
		rows.Scan(&c)
		codes = append(codes, c)
	}
	return codes, nil
}

func (r *URLRepository) DeleteURL(code string) {
	r.db.Exec(`DELETE FROM urls WHERE short_code=$1`, code)
}
