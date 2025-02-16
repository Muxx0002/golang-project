package handlers

import (
	"time"

	"github.com/Muxx0002/golang-project/tree/main/backend/internal/database/postgres/actions"
	"github.com/Muxx0002/golang-project/tree/main/backend/internal/dto"
	"github.com/Muxx0002/golang-project/tree/main/backend/pkg/tools"
	"github.com/gofiber/fiber/v2"
	"github.com/google/logger"
)

func Registration(c *fiber.Ctx) error {
	var authData dto.AuthData
	if err := c.BodyParser(authData); err != nil {
		logger.Error("failed parse body: ", err)
		return c.JSON(fiber.Map{"error": err})
	}

	if err := tools.ValidateRegistration(&authData.Email, &authData.Password, &authData.Username); err != nil {
		return c.JSON(fiber.Map{"error": err})
	}

	token := tools.GenerateDoubleID()

	if err := actions.CreateUser(authData.Email, authData.Password, authData.Username, token); err != nil {
		logger.Error(err)
		return c.JSON(fiber.Map{"error": err})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "auth",
		Value:    token,
		HTTPOnly: true,
	})
	return nil
}

func Login(c *fiber.Ctx) error {
	var authData dto.AuthData
	if err := c.BodyParser(authData); err != nil {
		logger.Error("failed parse body: ", err)
		return c.JSON(fiber.Map{"error": err})
	}

	if err := tools.ValidateRegistration(&authData.Email, &authData.Password, &authData.Username); err != nil {
		return c.JSON(fiber.Map{"error": err})
	}

	user, err := actions.CheckUser(authData.Password, authData.Email)
	if err != nil {
		return c.JSON(fiber.Map{"error": err})
	}
	c.Cookie(&fiber.Cookie{
		Name:     "auth",
		Value:    user.Token,
		HTTPOnly: true,
	})
	return nil
}

func LogOut(c *fiber.Ctx) error {
	c.Locals("user", nil)
	c.Cookie(&fiber.Cookie{
		Name:     "auth",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	})
	return c.Redirect("/sign-in", fiber.StatusFound)
}
