package middleware

import (
	"strings"
	"topupservice/domain"

	"github.com/gofiber/fiber/v2"
)

func Authenticate(authService domain.AuthService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		authHeader := ctx.Get("Authorization")
		token := strings.ReplaceAll(authHeader, "Bearer ", "")
		if token == "" {
			return ctx.SendStatus(401)
		}

		// Memparsing token (memisahkan Bearer dan token)
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return ctx.SendStatus(401)
		}

		user, err := authService.Refresh(ctx.Context(), token)
		if err != nil {
			return ctx.SendStatus(401)
		}

		ctx.Locals("x-user", user)
		return ctx.Next()
	}
}
