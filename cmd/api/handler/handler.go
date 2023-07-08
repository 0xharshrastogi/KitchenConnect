package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/harshrastogiexe/KitchenConnect/cmd/api/handler/auth"
	"github.com/harshrastogiexe/KitchenConnect/cmd/api/helpers"
	"github.com/harshrastogiexe/KitchenConnect/cmd/api/utils"
	"github.com/harshrastogiexe/KitchenConnect/lib/go/common/interfaces"
	"go.uber.org/zap"
)

type (
	AuthHandler interface {
		LoginRouteHandler() fiber.Handler
		RegisterRouteHandler() fiber.Handler
	}
)

func NewAuthHandler(v *validator.Validate, ur interfaces.UserRepository, log *zap.Logger) AuthHandler {
	h := &auth.AuthHandler{
		Log:             log,
		Validate:        v,
		Users:           ur,
		PasswordHandler: utils.NewPasswordHandler(),
		Jwt: &auth.JwtConfig{
			Secret: []byte(helpers.GetJwtSecret()),
			Method: jwt.GetSigningMethod(helpers.GetJwtAlgorithm()),
		},
	}
	return h
}
