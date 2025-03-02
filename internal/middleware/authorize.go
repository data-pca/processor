package middleware

import (
	"github.com/gofiber/fiber/v3"
	"processor/pkg/jwt"
	"strings"
)

const BearerPrefix = "Bearer:"

func AccessCheck(c fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		status := "missing authorization header"
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": status})
	}

	tokenString := strings.TrimPrefix(authHeader, BearerPrefix)
	claims, err := jwt.ValidateAccessToken(strings.TrimSpace(tokenString))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
	}

	c.Locals("userID", claims.UserID)

	return c.Next()
}
