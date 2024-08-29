package security

import "golang.org/x/crypto/bcrypt"

func Hash(pass string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
}

func ValidatePass(passHash, passString string ) error {
	return bcrypt.CompareHashAndPassword([]byte(passHash), []byte(passString))
}