package repository

import (
	"database/sql"
	"log"
)

func InitSchema(db *sql.DB) {
	schema := `CREATE TABLE IF NOT EXISTS urls (
		id SERIAL PRIMARY KEY,
		original_url TEXT NOT NULL,
		short_code VARCHAR(10) UNIQUE NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		expires_at TIMESTAMP,
		hit_count INTEGER DEFAULT 0
	)`
	_, err := db.Exec(schema)
	if err != nil {
		log.Fatal("Failed to create table:", err)
	}
}
