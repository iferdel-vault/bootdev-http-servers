package main

import (
	"net/http"
)

func (cfg *apiConfig) handlerChirpsGet(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	respondWithJSON(w, http.StatusOK, []Chirp{})

}
