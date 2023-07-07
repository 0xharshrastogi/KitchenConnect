package handler

import (
	"errors"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/harshrastogiexe/KitchenConnect/cmd/api/dto"
	"github.com/harshrastogiexe/KitchenConnect/cmd/api/utils"
	"github.com/harshrastogiexe/KitchenConnect/lib/go/common/interfaces"
	"github.com/harshrastogiexe/KitchenConnect/lib/go/db/models"
	"github.com/stroiman/go-automapper"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var (
	ErrInvalidToken = errors.New("invalid jwt token")
)

var (
	tokenExpirationTimeInterval = time.Hour
	JwtCookieName               = "token"
)

type AuthHandler struct {
	validate      *validator.Validate
	user          interfaces.IUserRepository
	logger        *zap.Logger
	signingMethod jwt.SigningMethod
	secret        []byte
}

func NewAuthHandler(l fx.Lifecycle, v *validator.Validate, u interfaces.IUserRepository, logger *zap.Logger) *AuthHandler {
	return &AuthHandler{
		validate:      v,
		user:          u,
		logger:        logger,
		signingMethod: jwt.SigningMethodHS256,
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

		t, err := a.makeJwtToken(jwt.MapClaims{
			"user": map[string]any{
				"id":    u.ID,
				"email": u.Email,
			},
		})

		if err != nil {
			return err
		}
		c.Cookie(createJWTCookie(t))
		if err := c.JSON(fiber.Map{"token": t, "user": a.convertUserOut(u)}); err != nil {
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
		t, err := a.makeJwtTokenFromUser(usr)
		if err != nil {
			return err
		}
		c.Cookie(createJWTCookie(t))
		if err := c.JSON(fiber.Map{"token": t, "user": a.convertUserOut(usr)}); err != nil {
			return err
		}
		return nil
	}

}

func (a *AuthHandler) makeJwtToken(claims jwt.MapClaims) (string, error) {
	claims["exp"] = time.Now().Add(tokenExpirationTimeInterval)
	claims["iat"] = time.Now()

	return jwt.NewWithClaims(a.signingMethod, claims).SignedString(a.secret)
}

func (a *AuthHandler) verifyToken(token string) error {
	t, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvalidToken
		}
		return a.secret, nil
	})

	if err != nil {
		return nil
	}
	if !t.Valid {
		return ErrInvalidToken
	}
	return nil
}

func (a *AuthHandler) convertUserOut(u *models.User) *dto.UserOutDTO {
	uOut := &dto.UserOutDTO{}
	automapper.MapLoose(u, uOut)
	return uOut
}

func (a *AuthHandler) makeJwtTokenFromUser(u *models.User) (string, error) {
	return a.makeJwtToken(jwt.MapClaims{
		"user": map[string]interface{}{
			"id":    u.ID,
			"email": u.Email,
		},
	})
}

func createJWTCookie(token string) *fiber.Cookie {
	return &fiber.Cookie{
		Name:     JwtCookieName,
		Value:    token,
		Expires:  time.Now().Add(tokenExpirationTimeInterval),
		HTTPOnly: true,
	}
}
