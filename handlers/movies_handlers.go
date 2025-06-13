package handlers

import (
	"encoding/json"
	"net/http"

	"frontendmasters.com/movies/data"
	"frontendmasters.com/movies/logger"
)

type MovieHandler struct {
	storage data.MovieStorage
	logger  *logger.Logger
}

// Function writeJSONResponse
func (h *MovieHandler) writeJSONResponse(w http.ResponseWriter, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		h.logger.Error("Failed to encode response", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return err
	}
	return nil
}

func (h *MovieHandler) GetTopMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := h.storage.GetTopMovies()
	if err != nil {
		http.Error(w, "Failed to get top movies", 500)
		h.logger.Error("Failed to get top movies", err)
		return
	}
	h.writeJSONResponse(w, movies)
}

func (h *MovieHandler) GetRandomMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := h.storage.GetRandomMovies()
	if err != nil {
		http.Error(w, "Failed to get random movies", 500)
		h.logger.Error("Failed to get random movies", err)
		return
	}
	h.writeJSONResponse(w, movies)
}
