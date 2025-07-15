package worker

import (
	"log"
	"time"

	"urlShortner/internal/repository"
)

func StartCleaner(repo *repository.URLRepository) {
	log.Println("Cleaner started: ")
	for {
		codes, err := repo.GetExpiredURLs()
		if err == nil {
			for _, code := range codes {
				repo.DeleteURL(code)
				log.Println("Deleted expired URL:", code)
			}
		}
		log.Println("Cleaner finished, Next cleaning after 1 hour.")
		time.Sleep(1 * time.Hour)
	}
}
