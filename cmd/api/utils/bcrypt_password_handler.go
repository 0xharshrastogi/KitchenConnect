package utils

import "golang.org/x/crypto/bcrypt"

// bcryptPasswordHandler implements the PasswordHandler interface using bcrypt.
type bcryptPasswordHandler struct{}

// HashPassword generates a hashed password from a plain text password using bcrypt.
func (b *bcryptPasswordHandler) HashPassword(plainText string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(plainText), bcrypt.MinCost)
	return string(hashedPassword), err
}

// ValidatePassword compares a hashed password with a plain text password
// and returns true if they match, false otherwise.
func (b *bcryptPasswordHandler) ValidatePassword(hashedPassword, plainText string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainText))
	return err == nil
}
