package main

import (
	"encoding/json"
	"net/http"

	"github.com/iferdel-vault/bootdev-http-servers/internal/auth"
	"github.com/iferdel-vault/bootdev-http-servers/internal/database"
)

func (cfg *apiConfig) handlerUsersUpdate(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	type requestBody struct {
		Password string `json:"password"`
		Email    string `json:"email"`
	}

	type responseBody struct {
		User
	}

	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Something went wrong with the token verification from the header content of the request", err)
		return
	}

	userID, err := auth.ValidateJWT(token, cfg.jwtSecret)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Unauthorized...", err)
		return
	}

	decoder := json.NewDecoder(r.Body)
	params := requestBody{}
	err = decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	hashedPassword, err := auth.HashPassword(params.Password)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Password entered not valid", err)
		return
	}

	user, err := cfg.db.UpdateUserEmailAndPassword(r.Context(), database.UpdateUserEmailAndPasswordParams{
		Email:          params.Email,
		HashedPassword: hashedPassword,
		ID:             userID,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't update user email and/or password", err)
		return
	}

	respondWithJSON(w, http.StatusOK, responseBody{
		User: User{
			ID:        user.ID,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
			Email:     user.Email,
		},
	})
}
