package main

import (
	"net/http"

	"github.com/iferdel-vault/bootdev-http-servers/internal/auth"
)

func (cfg *apiConfig) handlerRevoke(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't get bearer token from Header", err)
		return
	}
	dbRefreshToken, err := cfg.db.GetRefreshToken(r.Context(), token)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "token from header does not exist in database", err)
		return
	}
	err = cfg.db.RevokeRefreshToken(r.Context(), dbRefreshToken.Token)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "cannot revoke token", err)
		return
	}
	respondWithJSON(w, http.StatusNoContent, nil)
}
