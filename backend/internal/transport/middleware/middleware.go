package middleware

import (
	"github.com/Muxx0002/golang-project/tree/main/backend/internal/database/postgres/actions"
	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(c *fiber.Ctx) error {
	token := c.Cookies("auth")
	if token == "" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	user, err := actions.CheckUserByToken(token)
	if err != nil || user.ID == "" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	c.Locals("user", user)
	path := c.Path()
	if path == "/auth/sign-in" || path == "/auth/sign-up" {
		return c.Redirect("/articles", fiber.StatusSeeOther)
	}
	return c.Next()
}

func IsAdmin(c *fiber.Ctx) error {
	token := c.Cookies("auth")
	if token == "" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	user, err := actions.CheckUserByToken(token)
	if err != nil || user.Role == "" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	if user.ID != "admin" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	return c.Next()
}
