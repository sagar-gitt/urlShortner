package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"

	"urlShortner/internal/config"
	"urlShortner/internal/handlers"
	"urlShortner/internal/repository"
	"urlShortner/internal/utils"
	"urlShortner/internal/worker"
)

func main() {
	//Loading .env file for credentials
	config.LoadEnv()

	//Connecting to PostgreSQL
	db := InitDB()

	//Initialize table
	repository.InitSchema(db)

	//Initialize Repo
	repo := repository.NewURLRepository(db)

	r := chi.NewRouter()
	h := handlers.NewURLHandler(repo)
	h.RegisterRoutes(r)

	go worker.StartCleaner(repo)
	utils.InitLogger()

	log.Println("Server started on :8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
		return
	}
}

func InitDB() *sql.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=require",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"))

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Failed to connect to DB:", err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal("DB unreachable:", err)
	}
	return db
}
