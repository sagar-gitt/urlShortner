package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"urlShortner/internal/services"
)

type ShortenRequest struct {
	URL string `json:"url"`
}

type ShortenResponse struct {
	Code string `json:"code"`
}

func (h *URLHandler) ShortenURL(w http.ResponseWriter, r *http.Request) {
	var req ShortenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	code := services.GenerateShortCode(6)
	expires := time.Now().Add(24 * time.Hour)
	err := h.repo.SaveURL(req.URL, code, expires)
	if err != nil {
		http.Error(w, "Failed to save", http.StatusInternalServerError)
		return
	}

	ip := r.RemoteAddr
	log.Printf("[Shorten] [%s] New URL shortened: %s -> %s", time.Now().Format(time.RFC3339), req.URL, ip)

	err = json.NewEncoder(w).Encode(ShortenResponse{Code: code})
	if err != nil {
		return
	}
}

func (h *URLHandler) ResolveURL(w http.ResponseWriter, r *http.Request) {
	code := chi.URLParam(r, "code")
	url, err := h.repo.GetURLByCode(code)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	h.repo.IncrementHitCount(code)
	ip := r.RemoteAddr
	log.Printf("[Redirect] [%s] Code: %s -> %s (IP: %s)", time.Now().Format(time.RFC3339), code, url.OriginalURL, ip)

	http.Redirect(w, r, url.OriginalURL, http.StatusFound)
}

func (h *URLHandler) GetStats(w http.ResponseWriter, r *http.Request) {
	code := chi.URLParam(r, "code")

	url, err := h.repo.GetURLByCode(code)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	err = json.NewEncoder(w).Encode(url)
	if err != nil {
		return
	}
}
