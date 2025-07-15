package main

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/iferdel-vault/bootdev-http-servers/internal/database"
)

func (cfg *apiConfig) handlerPolkaWebhooks(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	type requestBody struct {
		Event string `json:"event"`
		Data  struct {
			UserID uuid.UUID `json:"user_id"`
		} `json:"data"`
	}

	params := requestBody{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Something went wrong decoding request", err)
		return
	}

	if params.Event != "user.upgraded" {
		respondWithJSON(w, http.StatusNoContent, nil)
		return
	}

	_, err = cfg.db.UpdateUserIsChirpyRed(r.Context(), database.UpdateUserIsChirpyRedParams{
		IsChirpyRed: true,
		ID:          params.Data.UserID,
	})
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Couldn't update chirpy red on user since user is not found", err)
		return
	}

	respondWithJSON(w, http.StatusNoContent, nil)
}
