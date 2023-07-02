// Package interfaces provides a set of interfaces related to logging.
package interfaces

import "github.com/harshrastogiexe/KitchenConnect/lib/go/db/models"

// ILogger represents a logger interface.
type ILogger interface {
	// Warn logs a warning-level message.
	Warn(message string)

	// Info logs an informational message.
	Info(message string)

	// Error logs an error-level message.
	Error(message string)

	// Debug logs a debug-level message.
	Debug(message string)
}

type IUserRepository interface {
	// Save saves a user to the repository.
	// It checks if a user with the same email already exists in the database.
	// If a user with the same email exists, it returns an error.
	// If the user does not exist, it creates a new user record in the database.
	// If an error occurs during the save operation, the error is returned.
	Save(u *models.User) error

	// FindByEmail retrieves a user from the repository based on the provided email address.
	// If a user with the given email is found, the corresponding user object is returned.
	// If no user is found, nil is returned without an error.
	// If an error occurs during the query execution, the error is returned.
	FindByEmail(email string) (*models.User, error)

	// IsEmailExist checks if a user with the specified email exists in the repository.
	// It returns a boolean value indicating whether the user exists or not.
	// If an error occurs during the check, the error is returned.
	IsEmailExist(email string) (bool, error)
}
