package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

const ChirpMaxLength = 140
const ProfaneReplacement = "****"

var ProfaneWords = map[string]bool{
	"kerfuffle": true,
	"sharbert":  true,
	"fornax":    true,
}

// personal challenge: add middleware that logs whatever. maybe current time

func handlerValidateChirp(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	type requestBody struct {
		Body string `json:"body"`
	}
	type responseBody struct {
		CleanedBody string `json:"cleaned_body"`
	}

	decoder := json.NewDecoder(r.Body)
	params := requestBody{}
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

	cleanedBody := BadWordReplacement(params.Body)
	respondWithJSON(w, 200, responseBody{
		CleanedBody: cleanedBody,
	})
	return
}

func BadWordReplacement(chirp string) string {
	words := strings.Split(chirp, " ")
	for i, word := range words {
		if ok := ProfaneWords[strings.ToLower(word)]; ok {
			words[i] = ProfaneReplacement
		}
	}
	return strings.Join(words, " ")
}
