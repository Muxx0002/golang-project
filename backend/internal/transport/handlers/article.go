package handlers

import (
	"github.com/Muxx0002/golang-project/tree/main/backend/internal/database/postgres/actions"
	"github.com/Muxx0002/golang-project/tree/main/backend/internal/dto"
	"github.com/Muxx0002/golang-project/tree/main/backend/internal/models"
	"github.com/gofiber/fiber/v2"
)

func ListArticlesHandler(c *fiber.Ctx) error {
	articles, err := actions.Articles()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	c.Set("Content-Type", "application/json")
	return c.JSON(articles)
}

func ArticleByIDHandler(c *fiber.Ctx) error {
	id := c.Params("id")
	article, err := actions.ArticleByID(id)
	if err != nil {
		if err == fiber.ErrNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Article not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	comments, err := actions.CommentsByArticleID(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to get comments"})
	}
	response := fiber.Map{
		"article":  article,
		"comments": comments,
	}
	return c.JSON(response)
}

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
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}
	if err := actions.CreateCommentByID(article.ID, comment.Content, user.ID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Comment created successfully"})
}
