package auth

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {

	// bcrypt.GenerateFromPassword()

	return "", nil
}

func CheckPasswordHash(password, hash string) error {

	// bcrypt.CompareHashAndPassword()

	return nil
}
