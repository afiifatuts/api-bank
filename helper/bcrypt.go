package helper

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(pw, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pw))
	return err == nil
}
