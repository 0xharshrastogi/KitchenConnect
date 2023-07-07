package auth

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/harshrastogiexe/KitchenConnect/cmd/api/dto"
	"go.uber.org/zap"
)

func (h *AuthHandler) LoginRouteHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var (
			fields = make([]zap.Field, 0, 5)
		)

		cred, err := parseRequestBody[dto.CredentialInDTO](h.Log, c, fields)
		if err != nil {
			return err
		}

		if err := h.validate(cred, fields); err != nil {
			return err
		}

		user, err := h.Users.FindByEmail(cred.Email)
		if err != nil {
			return err
		}

		fields = append(fields, zap.String("email", user.Email))
		if user == nil || !h.ValidatePassword(user.Password, cred.Password) {
			msg := "invalid credentials, either email not exist or invalid password"
			h.Log.Error(msg, fields...)
			return fiber.NewError(http.StatusUnauthorized, msg)
		}

		token, err := makeJwtFromUser(user)
		if err != nil {
			h.Log.Error(err.Error())
			return err
		}

		data := fiber.Map{
			"token": token,
			"user":  convertTo[dto.UserOutDTO](user),
		}

		if err = c.JSON(data); err != nil {
			return err
		}

		h.Log.Info("user logged-in successful", fields...)
		return nil
	}
}
