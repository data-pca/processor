package http

import (
	"github.com/gofiber/fiber/v3"
	"processor/internal/authorization"
	"processor/internal/models/dto"
)

type handler struct {
	uc authorization.UseCase
}

func New(uc authorization.UseCase) authorization.Handler {
	return &handler{uc: uc}
}

// @Summary Sign In
// @Description Sing In request
// @Tags auth
// @Accept json
// @Produce json
// @Param request body dto.SignInRequest true "Sign In Request"
// @Success 200 {object} jwt.Tokens
// @Failure 400 {object} map[string]string "Malformed Request"
// @Failure 502 {object} map[string]string "Internal Server Error"
// @Router /api/sign-in [post]
func (h *handler) SignIn() fiber.Handler {
	return func(c fiber.Ctx) error {

		var params dto.SignInRequest
		if err := c.Bind().Body(&params); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "malformed request",
			})
		}

		tokens, err := h.uc.SignIn(c.Context(), params)
		if err != nil {
			return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
				"error": "cannot sign in due to internal server error",
			})
		}

		return c.Status(fiber.StatusOK).JSON(tokens)
	}
}

// @Summary Sign Up
// @Description sign up request
// @Tags auth
// @Accept json
// @Produce json
// @Param request body dto.SignUpRequest true "Sign In Request"
// @Success 200 {object} jwt.Tokens
// @Failure 400 {object} map[string]string "Malformed Request"
// @Failure 502 {object} map[string]string "Internal Server Error"
// @Router /api/sign-up [post]
func (h *handler) SignUp() fiber.Handler {
	return func(c fiber.Ctx) error {

		var params dto.SignUpRequest
		if err := c.Bind().Body(&params); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "malformed request",
			})
		}

		tokens, err := h.uc.SignUp(c.Context(), params)
		if err != nil {
			return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
				"error": "cannot sign up due to internal server error",
			})
		}

		return c.Status(fiber.StatusOK).JSON(tokens)
	}
}
