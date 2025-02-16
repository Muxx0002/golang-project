package actions

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/Muxx0002/golang-project/tree/main/backend/internal/database/postgres"
	"github.com/Muxx0002/golang-project/tree/main/backend/internal/models"
	"github.com/Muxx0002/golang-project/tree/main/backend/pkg/tools"
	"github.com/gofiber/fiber"
)

func CreateArticle(article models.Article) error {
	id := tools.GenerateID()
	query := `INSERT INTO articles (id, title, content, created_at) VALUES ($1, $2, $3, $4)`
	_, err := postgres.DB.Exec(context.Background(), query, id, article.Title, article.Content, time.Now())
	if err != nil {
		return fmt.Errorf("failed to create article: %w", err)
	}
	return nil
}

func UpdateArticleByID(id string, article models.Article) error {
	query := `UPDATE articles SET title = $1, content = $2, WHERE id = $3`
	result, err := postgres.DB.Exec(context.Background(), query, article.Title, article.Content, id)
	if err != nil {
		return fmt.Errorf("failed to update article: %w", err)
	}
	if rowsAffected := result.RowsAffected(); rowsAffected == 0 {
		return fiber.ErrNotFound
	}
	return nil
}

func Articles() ([]models.Article, error) {
	var articles []models.Article
	query := `SELECT * FROM articles`
	rows, err := postgres.DB.Query(context.Background(), query)
	if err != nil {
		return nil, fmt.Errorf("failed getting articles: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var article models.Article
		if err := rows.Scan(&article.ID, &article.Title, &article.Content, article.CreatedAt); err != nil {
			return nil, fmt.Errorf("failed scanning article: %w", err)
		}
		articles = append(articles, article)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error occurred during row iteration: %w", err)
	}
	return articles, nil
}

func ArticleByID(id string) (models.Article, error) {
	var article models.Article
	query := `SELECT id, title, content, created_at FROM articles WHERE id = $1`
	err := postgres.DB.QueryRow(context.Background(), query, id).Scan(
		&article.ID,
		&article.Title,
		&article.Content,
		&article.CreatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return article, fmt.Errorf("article with id %s not found", id)
		}
		return article, fmt.Errorf("error occurred during row iteration: %w", err)
	}
	return article, nil
}

func DeleteArticleByID(id string) error {
	query := `DELETE FROM articles WHERE id = $1`
	_, err := postgres.DB.Exec(context.Background(), query, id)
	if err != nil {
		return fmt.Errorf("failed to delete article with id %s: %w", id, err)
	}
	return nil
}
