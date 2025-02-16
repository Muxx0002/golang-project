package midleware

import (
	"github.com/Muxx0002/golang-project/tree/main/backend/internal/database/postgres/actions"
	"github.com/gofiber/fiber/v2"
)

func AuthMidleware(c *fiber.Ctx) error {
	token := c.Cookies("auth")
	if token == "" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	user, err := actions.CheckUserByToken(token)
	if err != nil || user.ID == "" {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	c.Locals("user", user)
	return c.Next()
}
