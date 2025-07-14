package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/iferdel-vault/bootdev-http-servers/internal/auth"
)

const expiresIn = time.Hour

func (cfg *apiConfig) handlerLogin(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	type requestBody struct {
		Email            string         `json:"email"`
		Password         string         `json:"password"`
		ExpiresInSeconds *time.Duration `json:"expires_in_seconds"`
	}
	type responseBody struct {
		User
		Token string `json:"token"`
	}

	params := requestBody{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	user, err := cfg.db.GetUserByEmail(r.Context(), params.Email)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "A user with that email does not exists", err)
		return
	}

	err = auth.CheckPasswordHash(params.Password, user.HashedPassword)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Incorrect email or password", err)
		return
	}

	expiresIn := time.Hour
	if params.ExpiresInSeconds != nil {
		if *params.ExpiresInSeconds < time.Hour {
			expiresIn = *params.ExpiresInSeconds
		}
	}

	jwtToken, err := auth.MakeJWT(user.ID, cfg.jwtSecret, expiresIn)

	respondWithJSON(w, http.StatusOK, responseBody{
		User: User{
			ID:    user.ID,
			Email: user.Email,
		},
		Token: jwtToken,
	})
}
