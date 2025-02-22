package handlers

import (
	"github.com/gofiber/fiber/v3"
	"processor/internal/authorization"
)

type handler struct {
	uc authorization.UseCase
}

func New(uc authorization.UseCase) authorization.Handler {
	return &handler{uc: uc}
}

func (h *handler) CheckAuth() fiber.Handler {
	return func(c fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	}
}
