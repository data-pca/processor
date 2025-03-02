package delivery

import (
	"github.com/gofiber/fiber/v3"
	"processor/internal/authorization"
)

func MatchRoutes(apiRouter fiber.Router, h authorization.Handler) {
	r := apiRouter.Group("authorization")

	r.Post("/sign-in", h.SignIn())
	r.Post("/sign-up", h.SignUp())
}
