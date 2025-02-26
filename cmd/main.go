package main

import (
	"context"
	"github.com/gofiber/fiber/v3"
	"github.com/jmoiron/sqlx"
	"processor/config"
	"processor/internal/authorization"
	authRoutes "processor/internal/authorization/delivery"
	authHandler "processor/internal/authorization/delivery/handlers"
	authRepo "processor/internal/authorization/storage/postgres"
	authUsecase "processor/internal/authorization/usecase"
	"processor/pkg/postgres"
)

var (
	pgDB *sqlx.DB
)

func main() {
	ctx := context.Background()
	app := fiber.New()

	if err := setup(ctx); err != nil {
		return
	}

	defer shutdown()

	authR := authRepo.New(pgDB)
	authU := authUsecase.New(authR)
	authH := authHandler.New(authU)

	if err := multiplex(app, authH); err != nil {
		return
	}
}

// setup - initializes the applications databases and servicies
func setup(ctx context.Context) (err error) {
	pgDB, err = postgres.NewConnector(ctx, config.Cfg.Postgres)
	if err != nil {
		return err
	}

	return
}

// shutdown - closes the applications databases and servicies
func shutdown() {
	_ = pgDB.Close()
}

// multiplex - wires fiber app to existing repository handlers
func multiplex(app *fiber.App, h authorization.Handler) error {
	apiGroup := app.Group("api")

	authRoutes.MatchRoutes(apiGroup, h)

	apiGroup.Get("/heartbeat", func(c fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	})

	return app.Listen(config.Cfg.GetMultiplexURL())
}
