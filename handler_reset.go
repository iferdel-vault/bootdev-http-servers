package main

import (
	"fmt"
	"net/http"
)

func (cfg *apiConfig) handlerReset(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if cfg.platform != "dev" {
		respondWithError(w, http.StatusForbidden, "Reset is only allowed in dev environment.", nil)
		return
	}

	cfg.fileserverHits.Store(0)

	err := cfg.db.DeleteUsers(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Failed to reset the database: %v\n", err), err)
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{
		"message": "Hits reset to 0 and database reset to initial state.",
	})
	return
}
