package main

import (
	"net/http"
	"sort"

	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerChirpsGet(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	type responseBody struct {
		Chirp
	}

	chirpIDString := r.PathValue("chirpID")
	chirpID, err := uuid.Parse(chirpIDString)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Invalid chirp ID", err)
		return
	}

	chirp, err := cfg.db.GetChirp(r.Context(), chirpID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Chirp not found", err)
		return
	}

	respondWithJSON(w, http.StatusOK, responseBody{
		Chirp: Chirp{
			ID:        chirp.ID,
			CreatedAt: chirp.CreatedAt,
			UpdatedAt: chirp.UpdatedAt,
			Body:      chirp.Body,
			UserID:    chirp.UserID,
		},
	})

}

func (cfg *apiConfig) handlerChirpsRetrieve(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	authorID := uuid.Nil
	authorIDString := r.URL.Query().Get("author_id")
	if authorIDString != "" {
		var err error
		authorID, err = uuid.Parse(authorIDString) // variable shadowing
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "Invalid author ID", err)
			return
		}
	}

	sortDirection := "asc"
	sortDirectionParam := r.URL.Query().Get("sort")
	if sortDirectionParam == "desc" {
		sortDirection = "desc"
	}

	chirps, err := cfg.db.ListChirps(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't retrieve chirps", err)
		return
	}

	responseBody := []Chirp{}
	for _, chirp := range chirps {
		if authorID != uuid.Nil && chirp.UserID != authorID {
			continue
		}
		responseBody = append(responseBody, Chirp{
			ID:        chirp.ID,
			CreatedAt: chirp.CreatedAt,
			UpdatedAt: chirp.UpdatedAt,
			Body:      chirp.Body,
			UserID:    chirp.UserID,
		})
	}

	sort.Slice(responseBody, func(i, j int) bool {
		if sortDirection == "desc" {
			return chirps[i].CreatedAt.After(chirps[j].CreatedAt)
		}
		return chirps[i].CreatedAt.Before(chirps[j].CreatedAt)
	})

	respondWithJSON(w, http.StatusOK, responseBody)
}
