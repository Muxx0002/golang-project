package handlers

import (
	"time"

	"github.com/Muxx0002/golang-project/tree/main/backend/internal/database/postgres/actions"
	"github.com/Muxx0002/golang-project/tree/main/backend/internal/dto"
	"github.com/Muxx0002/golang-project/tree/main/backend/internal/models"
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

func UserProfile(c *fiber.Ctx) error {
	userValue := c.Locals("user")
	if userValue == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Authentication required",
		})
	}
	user, ok := userValue.(models.Users)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Invalid user data format",
		})
	}
	response := dto.ProfileResponse{
		Email:     user.Email,
		Username:  user.Username,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
	}
	return c.JSON(response)
}

func UpdateUserData(c *fiber.Ctx) error {
	userValue := c.Locals("user")
	if userValue == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Authentication required",
		})
	}
	user, ok := userValue.(models.Users)
	if !ok {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Invalid user data format",
		})
	}
	var req dto.UpdateRequest
	if err := c.BodyParser(&req); err != nil {
		logger.Error("failed parse body: ", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request format",
		})
	}
	if exists, _ := actions.IsUsernameExists(req.Username); exists {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error": "Username already taken",
		})
	}

	if err := actions.UpdateUsername(user.ID, req.Username); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"error": "Failed to update username",
			})
	}
	return c.JSON(fiber.Map{
		"message":  "Username updated successfully",
		"username": req.Username,
	})
}
