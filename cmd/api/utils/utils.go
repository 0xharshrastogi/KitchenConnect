// Package utils provides utility functions for handling passwords.
package utils

type PasswordHandler interface {
	// HashPassword generates a hashed password from a plain text password.
	HashPassword(plainText string) (hashedPassword string, err error)

	// ValidatePassword compares a hashed password with a plain text password
	// and returns true if they match, false otherwise.
	ValidatePassword(hashedPassword, plainText string) bool
}

// NewPasswordHandler creates a new instance of PasswordHandler using bcrypt.
func NewPasswordHandler() PasswordHandler {
	return &bcryptPasswordHandler{}
}
