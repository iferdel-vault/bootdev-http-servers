package main

import (
	"net/http"

	"github.com/iferdel-vault/bootdev-http-servers/internal/auth"
)

func (cfg *apiConfig) handlerRefresh(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	type responseBody struct {
		Token string `json:"token"`
	}

	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't get bearer token from Header", err)
		return
	}
	user, err := cfg.db.GetUserFromRefreshToken(r.Context(), token)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "token does not belong to any user or token expired", err)
		return
	}
	accessToken, err := auth.MakeJWT(user.ID, cfg.jwtSecret, expirationTimeAccessToken)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create access JWT", err)
		return
	}
	respondWithJSON(w, http.StatusOK, responseBody{
		Token: accessToken,
	})
}
