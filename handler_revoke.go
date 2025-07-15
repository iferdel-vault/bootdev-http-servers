package main

import (
	"net/http"

	"github.com/iferdel-vault/bootdev-http-servers/internal/auth"
)

func (cfg *apiConfig) handlerRevoke(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	refreshToken, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't get bearer token from Header", err)
		return
	}
	err = cfg.db.RevokeRefreshToken(r.Context(), refreshToken)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't revoke session", err)
		return
	}
	respondWithJSON(w, http.StatusNoContent, nil)
}
