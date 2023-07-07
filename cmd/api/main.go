package main

import (
	"context"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/harshrastogiexe/KitchenConnect/cmd/api/config"
	"github.com/harshrastogiexe/KitchenConnect/cmd/api/db"
	"github.com/harshrastogiexe/KitchenConnect/cmd/api/handler"
	"github.com/harshrastogiexe/KitchenConnect/cmd/api/utils/logger"
	"github.com/harshrastogiexe/KitchenConnect/lib/go/common/repository"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

const CONNECTION_STRING_KEY string = "KitchenConnectDB"

func main() {
	fx.New(
		fx.Provide(config.BuildAppSetting("appsetting.json")),
		fx.Provide(logger.NewZapLogger),
		fx.Provide(NewValidator),
		fx.Provide(repository.NewUserRepository),
		fx.Provide(handler.NewAuthHandler),
		fx.Provide(NewHttpServer),
		fx.Provide(NewDBConnection),
		fx.Invoke(func(*gorm.DB) {}),
		fx.Invoke(func(*fiber.App) {}),
	).Run()
}

func NewHttpServer(lc fx.Lifecycle, ah handler.AuthHandler) *fiber.App {
	app := fiber.New()

	lc.Append(fx.Hook{
		OnStart: func(c context.Context) error {
			api := app.Group("/api")

			auth := api.Group("/auth")
			{
				auth.Post("/login", ah.LoginRouteHandler())
				auth.Get("/register", ah.RegisterRouteHandler())
			}

			go app.Listen(":8000")
			return nil
		},
		OnStop: func(context.Context) error {
			return app.Shutdown()
		},
	})
	return app
}

func NewDBConnection(lc fx.Lifecycle, setting config.AppSetting) (*gorm.DB, error) {
	dsn := setting.GetConnectionString(CONNECTION_STRING_KEY)
	d, err := db.ConnectSqlServerDB(dsn)
	if err != nil {
		return nil, err
	}
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			s, err := d.DB()
			if err != nil {
				return err
			}
			if err := s.Ping(); err != nil {
				return err
			}
			fmt.Println("Successfully connected to database:", d.Name())
			return nil
		},
		OnStop: func(context.Context) error {
			s, err := d.DB()
			if err != nil {
				return err
			}
			if err := s.Close(); err != nil {
				return err
			}
			return nil
		},
	})
	return d, err
}

func NewValidator(lc fx.Lifecycle) *validator.Validate {
	v := validator.New()
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return nil
		},
	})
	return v
}
