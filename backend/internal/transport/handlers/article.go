package handlers

import (
	"github.com/Muxx0002/golang-project/tree/main/backend/internal/database/postgres/actions"
	"github.com/gofiber/fiber/v2"
	"github.com/google/logger"
)

func ListArticlesHandler(c *fiber.Ctx) error {
	articles, err := actions.Articles()
	if err != nil {
		logger.Error("failed getting articles: ", err)
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
