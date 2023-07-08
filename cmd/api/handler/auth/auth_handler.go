package auth

import (
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/harshrastogiexe/KitchenConnect/cmd/api/helpers"
	"github.com/harshrastogiexe/KitchenConnect/cmd/api/utils"
	"github.com/harshrastogiexe/KitchenConnect/lib/go/common/interfaces"
	"github.com/harshrastogiexe/KitchenConnect/lib/go/db/models"
	"github.com/stroiman/go-automapper"
	"go.uber.org/zap"
)

var (
	jwtExtTimeInterval = time.Hour
	signingMethod      = jwt.SigningMethodHS256
)

type JwtConfig struct {
	Secret []byte
	Method jwt.SigningMethod
}

type AuthHandler struct {
	Log      *zap.Logger
	Validate *validator.Validate
	Users    interfaces.UserRepository
	Jwt      *JwtConfig
	utils.PasswordHandler
}

func parseRequestBody[T any](l *zap.Logger, c *fiber.Ctx, f []zap.Field) (*T, error) {
	v := new(T)
	if err := c.BodyParser(&v); err != nil {
		msg := "failed to parse json body"
		l.Error(msg, append(f, zap.Error(err))...)
		return nil, fiber.NewError(http.StatusBadRequest, msg)
	}
	return v, nil
}

func (h *AuthHandler) validate(s any, f []zap.Field) error {
	if err := h.Validate.Struct(s); err != nil {
		msg := "validation failed"
		h.Log.Error(msg, append(f, zap.Error(err))...)
		return &fiber.Error{Code: http.StatusBadRequest, Message: msg}
	}
	return nil
}

func (h *AuthHandler) makeJwtFromUser(u *models.User) (string, error) {
	return jwt.NewWithClaims(h.Jwt.Method, newJwtClaimsFromUser(u)).SignedString(h.Jwt.Secret)
}

func newJwtClaimsFromUser(u *models.User) jwt.MapClaims {
	c := newDefaultClaim()
	c["user"] = map[string]any{
		"id":    u.ID,
		"email": u.Email,
	}
	return c
}

func getSecret() []byte {
	return []byte(helpers.GetJwtSecret())
}

func newDefaultClaim() jwt.MapClaims {
	return jwt.MapClaims{
		"exp": time.Now().Add(jwtExtTimeInterval),
		"iat": time.Now(),
	}
}

func convertTo[T any](src any) *T {
	v := new(T)
	automapper.MapLoose(src, v)
	return v
}
