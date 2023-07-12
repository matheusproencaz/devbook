package security

import (
	"golang.org/x/crypto/bcrypt"
)

// Hash recebe uma senha e coloca um hash nela
func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// VerifyPassword compara uma senha e um hash e retorna se elas são iguais
func VerifyPassword(passwordHashed string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(passwordHashed), []byte(password))
}
