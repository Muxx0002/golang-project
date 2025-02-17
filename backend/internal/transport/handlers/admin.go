package handlers

import (
	"github.com/Muxx0002/golang-project/tree/main/backend/internal/database/postgres/actions"
	"github.com/Muxx0002/golang-project/tree/main/backend/internal/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/logger"
)

func CreateArticle(c *fiber.Ctx) error {
	var article models.Article
	if err := c.BodyParser(&article); err != nil {
		logger.Error("failed parse body: ", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}
	err := actions.CreateArticle(article)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create article"})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Article created successfully"})
}

func UpdateArticle(c *fiber.Ctx) error {
	id := c.Params("id")
	var article models.Article
	if err := c.BodyParser(&article); err != nil {
		logger.Error("failed parse body: ", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}
	err := actions.UpdateArticleByID(id, article)
	if err != nil {
		logger.Error("failed updated article: ", err)
		if err == fiber.ErrNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Article not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update article"})
	}
	return c.JSON(fiber.Map{
		"message": "Article updated successfully",
	})
}

func DeleteArticle(c *fiber.Ctx) error {
	id := c.Params("id")
	err := actions.DeleteArticleByID(id)
	if err != nil {
		logger.Error("failed delet user: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete article"})
	}
	return c.JSON(fiber.Map{"message": "Article delete successfully"})
}

func ListUsersHandler(c *fiber.Ctx) error {
	users, err := actions.GetAllUsers()
	if err != nil {
		logger.Error("failed scan user: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to scan users"})
	}
	var response []fiber.Map
	for _, user := range users {
		response = append(response, fiber.Map{
			"id":         user.ID,
			"email":      user.Email,
			"username":   user.Username,
			"role":       user.Role,
			"created_at": user.CreatedAt,
			"updated_at": user.UpdatedAt,
		})
	}
	return c.JSON(fiber.Map{
		"users": response,
	})
}

func DeleteCommentByIDHandler(c *fiber.Ctx) error {
	commentID := c.Params("id")
	err := actions.DeleteCommentByID(commentID)
	if err != nil {
		logger.Error("error delete comment")
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"error": "Failed to delete comment",
			})
	}
	return c.JSON(fiber.Map{"message": "comment delete successfully"})
}

func DeleteUserHandler(c *fiber.Ctx) error {
	userID := c.Params("id")
	err := actions.DeleteUserByID(userID)
	if err != nil {
		logger.Error("error delete user")
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"error": "Failed to delete users",
			})
	}
	return c.JSON(fiber.Map{"message": "user delete successfully"})
}
