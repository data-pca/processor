package main

import (
	"context"
	"github.com/Flussen/swagger-fiber-v3"
	"github.com/gofiber/fiber/v3"
	"github.com/jmoiron/sqlx"
	"processor/config"
	_ "processor/docs"
	"processor/internal/authorization"
	authRoutes "processor/internal/authorization/delivery"
	authHandler "processor/internal/authorization/delivery/http"
	authRepo "processor/internal/authorization/storage/postgres"
	authUsecase "processor/internal/authorization/usecase"
	"processor/pkg/postgres"
)

var (
	pgDB *sqlx.DB
)

// @title Fiber Example API
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:53247
// @BasePath /

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

	apiGroup.Get("/swagger/*", swagger.HandlerDefault)

	authRoutes.MatchRoutes(apiGroup, h)

	apiGroup.Get("/heartbeat", func(c fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	})

	return app.Listen(":53247")
}
