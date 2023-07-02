package repository

import (
	"fmt"

	"github.com/harshrastogiexe/KitchenConnect/lib/go/common/interfaces"
	"github.com/harshrastogiexe/KitchenConnect/lib/go/db/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) interfaces.IUserRepository {
	if db == nil {
		panic("nil instance found")
	}
	return &UserRepository{db: db}
}

// Save saves a user to the repository.
// It checks if a user with the same email already exists in the database.
// If a user with the same email exists, it returns an error.
// If the user does not exist, it creates a new user record in the database.
// If an error occurs during the save operation, the error is returned.
func (r *UserRepository) Save(u *models.User) error {
	su, err := r.FindByEmail(u.Email)
	if err != nil {
		return err
	}
	if su != nil {
		return fmt.Errorf("user already exist with email (%s)", u.Email)
	}
	if err := r.db.Create(u).Error; err != nil {
		return nil
	}
	return nil
}

// FindByEmail retrieves a user from the repository based on the provided email address.
// If a user with the given email is found, the corresponding user object is returned.
// If no user is found, nil is returned without an error.
// If an error occurs during the query execution, the error is returned.
func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	u := &models.User{}
	q := r.db.Where("email=?", email).First(u)
	if err := q.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return u, nil
}

func (r *UserRepository) IsEmailExist(email string) (bool, error) {
	u, err := r.FindByEmail(email)
	return u != nil, err
}
