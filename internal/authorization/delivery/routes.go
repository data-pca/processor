package delivery

import (
	"github.com/gofiber/fiber/v3"
	"processor/internal/authorization"
)

func MatchRoutes(apiRouter fiber.Router, h authorization.Handler) {
	r := apiRouter.Group("authorization")

	r.Get("/", h.CheckAuth())
}
