package delivery

import (
	"github.com/gofiber/fiber/v3"
	"processor/internal/authorization/delivery/handlers"
)

func MatchRoutes(apiRouter fiber.Router) {
	r := apiRouter.Group("authorization")
	h := handlers.New(nil)

	r.Get("/", h.CheckAuth())
}
