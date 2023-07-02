package handler

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/harshrastogiexe/KitchenConnect/lib/go/common/interfaces"
	"github.com/harshrastogiexe/KitchenConnect/lib/go/db/models"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

type AuthHandler struct {
	validate *validator.Validate
	user     interfaces.IUserRepository
}

func NewAuthHandler(l fx.Lifecycle, v *validator.Validate, u interfaces.IUserRepository) *AuthHandler {
	return &AuthHandler{
		validate: v,
		user:     u,
	}
}

func (*AuthHandler) LoginHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return nil
	}
}

func (a *AuthHandler) RegisterHandler() fiber.Handler {
	type (
		UserAddress struct {
			Street      string `json:"street" validate:"required"`
			City        string `json:"city" validate:"required"`
			State       string `json:"state" validate:"required"`
			ZipCode     string `json:"zipCode" validate:"required"`
			CountryCode string `json:"countryCode" validate:"required"`
		}

		UserInfo struct {
			FirstName string       `json:"firstName" validate:"required"`
			LastName  string       `json:"lastName" validate:"required"`
			Email     string       `json:"email" validate:"required"`
			Password  string       `json:"password" validate:"required"`
			Address   *UserAddress `json:"address" validate:"required"`
		}
	)

	return func(c *fiber.Ctx) error {
		var u UserInfo
		if err := c.BodyParser(&u); err != nil {
			return &fiber.Error{Code: http.StatusInternalServerError, Message: err.Error()}
		}

		if err := a.validate.Struct(&u); err != nil {
			return &fiber.Error{Code: http.StatusBadRequest, Message: err.Error()}
		}
		err := a.user.Save(&models.User{
			FirstName:    u.FirstName,
			LastName:     u.LastName,
			Email:        u.Email,
			Password:     u.Password,
			PasswordSalt: u.Password + u.FirstName,
			Address: &models.Address{
				City:        u.Address.City,
				Street:      u.Address.Street,
				State:       u.Address.State,
				CountryCode: "IN",
				ZipCode:     "102101",
			},
		})
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return &fiber.Error{Code: http.StatusNotFound, Message: err.Error()}
			}
			return err
		}
		return nil
	}

}
