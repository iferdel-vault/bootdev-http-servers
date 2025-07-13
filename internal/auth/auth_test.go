package auth

import (
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestValidateJWT(t *testing.T) {
	const tokenSecret = "a"
	userUUID := uuid.New()

	tests := map[string]struct {
		UserID      uuid.UUID
		expiresIn   time.Duration
		tokenSecret string
		WantErr     bool
	}{
		"right token": {
			UserID:      userUUID,
			tokenSecret: tokenSecret,
			expiresIn:   time.Hour * 1,
			WantErr:     false,
		},
		"wrong token": {
			UserID:      userUUID,
			tokenSecret: "badtoken",
			expiresIn:   time.Hour * 1,
			WantErr:     true,
		},
		"expired token": {
			UserID:      userUUID,
			tokenSecret: tokenSecret,
			expiresIn:   -time.Hour * 1,
			WantErr:     true,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			token, _ := MakeJWT(tc.UserID, tc.tokenSecret, tc.expiresIn)
			userID, _ := ValidateJWT(token, tokenSecret)
			if (userID != tc.UserID) != tc.WantErr {
				t.Errorf("WantErr is %v --> got %v, want %v", tc.WantErr, userID, tc.UserID)
			}
		})
	}
}

func TestCheckPasswordHash(t *testing.T) {
	password1 := "correctPassword123!"
	password2 := "anotherPassword456!"
	hash1, _ := HashPassword(password1)
	hash2, _ := HashPassword(password2)

	tests := []struct {
		name     string
		password string
		hash     string
		wantErr  bool
	}{
		{
			name:     "Correct password",
			password: password1,
			hash:     hash1,
			wantErr:  false,
		},
		{
			name:     "Incorrect password",
			password: "wrongPassword",
			hash:     hash1,
			wantErr:  true,
		},
		{
			name:     "Password doesn't match different hash",
			password: password1,
			hash:     hash2,
			wantErr:  true,
		},
		{
			name:     "Empty password",
			password: "",
			hash:     hash1,
			wantErr:  true,
		},
		{
			name:     "Invalid hash",
			password: password1,
			hash:     "invalidhash",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := CheckPasswordHash(tt.password, tt.hash)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckPasswordHash() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
