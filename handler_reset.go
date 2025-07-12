package main

import (
	"log"
	"net/http"
)

func (apiCfg *apiConfig) handlerResetUsers(w http.ResponseWriter, r *http.Request) {

	if apiCfg.platform != "dev" {
		respondWithError(w, 403, "Forbidden")
		return
	}

	err := apiCfg.db.DeleteUsers(r.Context())
	if err != nil {
		log.Printf("Error deleting users: %s", err)
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(200)
	return
}
