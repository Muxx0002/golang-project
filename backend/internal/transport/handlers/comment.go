package handlers

import (
	"github.com/Muxx0002/golang-project/tree/main/backend/internal/database/postgres/actions"
	"github.com/Muxx0002/golang-project/tree/main/backend/internal/dto"
	"github.com/Muxx0002/golang-project/tree/main/backend/internal/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/logger"
)

func CreateComment(c *fiber.Ctx) error {
	user, _ := c.Locals("user").(models.Users)
	id := c.Params("id")
	article, err := actions.ArticleByID(id)
	if err != nil {
		if err == fiber.ErrNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Article not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err})
	}
	var comment dto.CommentData
	if err := c.BodyParser(comment); err != nil {
		logger.Error("failed parse body: ", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}
	if comment.Content == "" || len(comment.Content) > 500 {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "the article is empty or contains a valid value"})
	}
	if err := actions.CreateCommentByID(article.ID, comment.Content, user.ID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Comment created successfully"})
}

func ListCommentsHandler(c *fiber.Ctx) error {
	comments, err := actions.GetAllComments()
	if err != nil {
		logger.Error("Failed to fetch comments: ", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch comments",
		})
	}
	return c.JSON(comments)
}
