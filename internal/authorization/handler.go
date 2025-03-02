package authorization

import "github.com/gofiber/fiber/v3"

type Handler interface {
	SignIn() fiber.Handler
	SignUp() fiber.Handler
}
