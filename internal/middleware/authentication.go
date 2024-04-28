package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/ravenocx/EAI-BackendAPI/internal/config"
	"github.com/ravenocx/EAI-BackendAPI/internal/domain"
)

type AuthConfig struct {
	Filter       func(*fiber.Ctx) error
	Unauthorized fiber.Handler
}

func AdminAuthentication(c AuthConfig) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		if ctx.Method() == fiber.MethodGet{
			return ctx.Next()
		}
		header := ctx.GetReqHeaders()

		if _, ok := header["Authorization"]; !ok {
			return c.Unauthorized(ctx)
		}

		var adminToken domain.Token

		headerToken := strings.Split(header["Authorization"][0] , " ") // Check if error

		if len(headerToken) != 2 || headerToken[0] != "Bearer" {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid Authorization header format"})
		}

		if err := config.DB.Where("token = ?", headerToken[1]).First(&adminToken).Error; err != nil {
			return c.Unauthorized(ctx)
		}

		ctx.Locals("admin", adminToken.Token)

		return ctx.Next()

	}
}
