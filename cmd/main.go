package main

import (
	"github.com/gofiber/fiber/v3"
	authorization "processor/internal/authorization/delivery"
)

func main() {
	app := fiber.New()

	multiplex(app)
}

// multiplex - wires fiber app to existing repository handlers
func multiplex(app *fiber.App) {
	apiGroup := app.Group("api")

	authorization.MatchRoutes(apiGroup)

	apiGroup.Get("/heartbeat", heartbeat())
}

// heartbeat - serves as a check if server is available
func heartbeat() fiber.Handler {
	return func(c fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	}
}
