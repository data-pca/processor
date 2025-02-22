package authorization

import "github.com/gofiber/fiber/v3"

type Handler interface {
	CheckAuth() fiber.Handler
}
