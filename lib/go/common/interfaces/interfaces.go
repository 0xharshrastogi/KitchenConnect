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
	Save(u *models.User) error

	// FindByEmail retrieves a user from the repository based on the provided email address.
	// If a user with the given email is found, the corresponding user object is returned.
	// If no user is found, nil is returned without an error.
	// If an error occurs during the query execution, the error is returned.
	FindByEmail(email string) (*models.User, error)
}
