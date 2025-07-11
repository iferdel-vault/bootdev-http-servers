package main

import (
	"encoding/json"
	"log"
	"net/http"
)

const ChirpMaxLength = 140

// personal challenge: add middleware that logs whatever. maybe current time

func handlerValidateChirp(w http.ResponseWriter, r *http.Request) {

	type parameters struct {
		Body string `json:"body"`
	}
	type returnVals struct {
		Valid bool `json:"valid"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		log.Printf("Error decoding parameters: %s", err)
		w.WriteHeader(500)
		return
	}

	if len(params.Body) > ChirpMaxLength {
		respondWithError(w, 400, "Chirp is too long")
		return
	}

	respBody := returnVals{
		Valid: true,
	}
	respondWithJSON(w, 200, respBody)
	return
}
