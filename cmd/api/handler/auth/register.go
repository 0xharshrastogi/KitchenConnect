package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/harshrastogiexe/KitchenConnect/cmd/api/dto"
	"github.com/harshrastogiexe/KitchenConnect/lib/go/db/models"
	"github.com/stroiman/go-automapper"
	"go.uber.org/zap"
)

func (h *AuthHandler) RegisterRouteHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var (
			fields = make([]zap.Field, 0, 5)
		)

		u, err := parseRequestBody[dto.UserInDTO](h.Log, c, fields)
		if err != nil {
			return err
		}

		if err := h.validate(u, fields); err != nil {
			return err
		}

		user := &models.User{}
		automapper.MapLoose(u, user)
		fields = append(fields, zap.String("email", u.Email))

		hash, err := h.HashPassword(user.Password)
		if err != nil {
			fields = append(fields, zap.Error(err))
			h.Log.Error("failed to hash password while signing up", fields...)
			return err
		}

		user.Password = hash

		if err := h.Users.Save(user); err != nil {
			fields = append(fields, zap.Error(err))
			h.Log.Error("failed to save info", fields...)
			return err
		}

		h.Log.Info("user info saved", fields...)

		t, err := h.makeJwtFromUser(user)
		if err != nil {
			return err
		}

		data := fiber.Map{
			"token": t,
			"user":  convertTo[dto.UserOutDTO](user),
		}
		if err := c.JSON(data); err != nil {
			return err
		}
		return nil
	}
}
