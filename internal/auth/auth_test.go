package auth

import (
	"net/http"
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestGetBearerToken(t *testing.T) {
	tests := map[string]struct {
		headers   http.Header
		wantToken string
		wantErr   bool
	}{
		"Valid Bearer token": {
			headers: http.Header{
				"Authorization": []string{"Bearer valid_token"},
			},
			wantToken: "valid_token",
			wantErr:   false,
		},
		"Missing Authorization header": {
			headers:   http.Header{},
			wantToken: "",
			wantErr:   true,
		},
		"Malformed Authorization header": {
			headers: http.Header{
				"Authorization": []string{"InvalidBearer token"},
			},
			wantToken: "",
			wantErr:   true,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			gotToken, err := GetBearerToken(tc.headers)
			if (err != nil) != tc.wantErr {
				t.Errorf("GetBearerToken() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
			if gotToken != tc.wantToken {
				t.Errorf("GetBearerToken() gotToken = %v, want %v", gotToken, tc.wantToken)
			}
		})
	}
}

func TestValidateJWT(t *testing.T) {
	userID := uuid.New()
	validToken, _ := MakeJWT(userID, "secret", time.Hour)
	expiredToken, _ := MakeJWT(userID, "secret", -time.Hour)

	tests := map[string]struct {
		tokenString string
		tokenSecret string
		wantUserID  uuid.UUID
		wantErr     bool
	}{
		"valid token": {
			tokenString: validToken,
			tokenSecret: "secret",
			wantUserID:  userID,
			wantErr:     false,
		},
		"wrong token": {
			tokenString: "invalid.token.string",
			tokenSecret: "secret",
			wantUserID:  uuid.Nil,
			wantErr:     true,
		},
		"wrong secret": {
			tokenString: validToken,
			tokenSecret: "wrong_secret",
			wantUserID:  uuid.Nil,
			wantErr:     true,
		},
		"expired token": {
			tokenString: expiredToken,
			tokenSecret: "secret",
			wantUserID:  uuid.Nil,
			wantErr:     true,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			gotUserID, err := ValidateJWT(tc.tokenString, tc.tokenSecret)
			if (err != nil) != tc.wantErr {
				t.Errorf("ValidateJWT() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
			if gotUserID != tc.wantUserID {
				t.Errorf("ValidateJWT() gotUserID = %v, want %v", gotUserID, tc.wantUserID)
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
