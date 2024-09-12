package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func ComparePassword(hashPassowrd, passowrd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassowrd), []byte(passowrd))
	if err != nil {
		// pasword is wrong
		return false
	}
	// password is correct
	return true
}
