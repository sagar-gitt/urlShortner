package handlers

import (
	"github.com/go-chi/chi/v5"
	"urlShortner/internal/repository"
)

type URLHandler struct {
	repo *repository.URLRepository
}

func NewURLHandler(repo *repository.URLRepository) *URLHandler {
	return &URLHandler{repo: repo}
}

func (h *URLHandler) RegisterRoutes(r chi.Router) {
	r.Post("/shorten", h.ShortenURL)
	r.Get("/{code}", h.ResolveURL)
	r.Get("/stats/{code}", h.GetStats)
}
