package main

import (
	"net/http"
)

func (cfg *apiConfig) handlerChirpsGet(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	chirpID := r.PathValue("ChirpID")
	if chirpID == "" {
		respondWithError(w, http.StatusBadRequest, "The chirp ID should be in the path", nil)
	}

}

func (cfg *apiConfig) handlerChirpsRetrieve(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	chirps, err := cfg.db.ListChirps(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't retrieve chirps", err)
	}

	responseBody := []Chirp{}
	for _, chirp := range chirps {
		responseBody = append(responseBody, Chirp{
			ID:        chirp.ID,
			CreatedAt: chirp.CreatedAt,
			UpdatedAt: chirp.UpdatedAt,
			Body:      chirp.Body,
			UserID:    chirp.UserID,
		})
	}

	respondWithJSON(w, http.StatusOK, responseBody)
}
