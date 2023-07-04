package handler

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/harshrastogiexe/KitchenConnect/cmd/api/utils"
	"github.com/harshrastogiexe/KitchenConnect/lib/go/common/interfaces"
	"github.com/harshrastogiexe/KitchenConnect/lib/go/db/models"
	"github.com/stroiman/go-automapper"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type AuthHandler struct {
	validate *validator.Validate
	user     interfaces.IUserRepository
	logger   *zap.Logger
}

func NewAuthHandler(l fx.Lifecycle, v *validator.Validate, u interfaces.IUserRepository, logger *zap.Logger) *AuthHandler {
	return &AuthHandler{
		validate: v,
		user:     u,
		logger:   logger,
	}
}

func (a *AuthHandler) LoginHandler() fiber.Handler {
	type credential struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,gte=6,lte=50"`
	}
	var (
		password = utils.NewPasswordHandler()
	)
	return func(c *fiber.Ctx) error {
		var (
			cred credential
			zf   = make([]zap.Field, 0, 5)
		)

		if err := c.BodyParser(&cred); err != nil {
			zf = append(zf, zap.Error(err))
			a.logger.Error("failed to parse json body", zf...)
			return err
		}

		if err := a.validate.Struct(&cred); err != nil {
			zf = append(zf, zap.Error(err))
			a.logger.Error("validation failed", zf...)
			return &fiber.Error{Code: http.StatusBadRequest, Message: err.Error()}
		}

		u, err := a.user.FindByEmail(cred.Email)
		if err != nil {
			return err
		}
		zf = append(zf, zap.String("email", cred.Email))
		if u == nil || !password.ValidatePassword(u.Password, cred.Password) {
			a.logger.Error("user validation failed, either email not exist or invalid password", zf...)
			return fiber.NewError(http.StatusUnauthorized)
		}

		if err := c.JSON(fiber.Map{"message": "user login success"}); err != nil {
			return err
		}

		a.logger.Info("user logged in successful", zf...)
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
	var (
		password = utils.NewPasswordHandler()
	)
	return func(c *fiber.Ctx) error {
		var (
			u  UserInfo
			zf = make([]zap.Field, 0, 5)
		)

		if err := c.BodyParser(&u); err != nil {
			zf = append(zf, zap.Error(err))
			a.logger.Error("failed to parse json body", zf...)
			return &fiber.Error{Code: http.StatusInternalServerError, Message: err.Error()}
		}

		if err := a.validate.Struct(&u); err != nil {
			zf = append(zf, zap.Error(err))
			a.logger.Error("validation failed for user", zf...)
			return &fiber.Error{Code: http.StatusBadRequest, Message: err.Error()}
		}

		usr := &models.User{}
		automapper.MapLoose(u, usr)
		zf = append(zf, zap.String("email", u.Email))

		hash, err := password.HashPassword(usr.Password)
		if err != nil {
			zf = append(zf, zap.Error(err))
			a.logger.Error("password hashing failed", zf...)
			return err
		}

		usr.Password = hash

		if err := a.user.Save(usr); err != nil {
			zf = append(zf, zap.Error(err))
			a.logger.Error("failed to save info", zf...)
			return err
		}
		a.logger.Info("user info registered", zf...)
		return nil
	}

}
