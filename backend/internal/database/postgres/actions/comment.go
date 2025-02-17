package actions

import (
	"context"
	"fmt"
	"time"

	"github.com/Muxx0002/golang-project/tree/main/backend/internal/database/postgres"
	"github.com/Muxx0002/golang-project/tree/main/backend/internal/models"
	"github.com/Muxx0002/golang-project/tree/main/backend/pkg/tools"
)

func CommentsByArticleID(articleID string) ([]models.Comment, error) {
	var comments []models.Comment
	query := `SELECT id, content, user_id, created_at FROM comments WHERE article_id = $1`
	rows, err := postgres.DB.Query(context.Background(), query, articleID)
	if err != nil {
		return nil, fmt.Errorf("failed to query comments: %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		var comment models.Comment
		err := rows.Scan(&comment.ID, &comment.Content, &comment.UserID, &comment.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan comment: %w", err)
		}
		comments = append(comments, comment)
	}
	return comments, nil
}

func CreateCommentByID(articleID string, comment string, userId string) error {
	id := tools.GenerateID()
	query := `
        INSERT INTO comments (id, content, article_id, user_id, created_at, update_at)
        VALUES ($1, $2, $3, $4, $5, $6);`
	_, err := postgres.DB.Exec(context.Background(), query, id, comment, articleID, userId, time.Now(), time.Now())
	if err != nil {
		return fmt.Errorf("failed to create comment: %w", err)
	}
	return nil
}

func UpdateCommentByID(commentID string, newContent string) error {
	query := `
        UPDATE comments 
        SET content = $1, update_at = $2
        WHERE id = $3
    `
	result, err := postgres.DB.Exec(context.Background(), query, newContent, time.Now(), commentID)
	if err != nil {
		return fmt.Errorf("failed to update comment: %w", err)
	}
	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("no comment found with id %s", commentID)
	}
	return nil
}

func DeleteCommentByID(commentID string) error {
	query := `DELETE FROM comments WHERE id = $1`
	result, err := postgres.DB.Exec(context.Background(), query, commentID)
	if err != nil {
		return fmt.Errorf("failed to delete comment: %w", err)
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("no comment found with id %s", commentID)
	}
	return nil
}

func GetAllComments() ([]models.Comment, error) {
	var comments []models.Comment
	query := `SELECT id, text, user_id, article_id, created_at FROM comments`
	rows, err := postgres.DB.Query(context.Background(), query)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch comments: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var comment models.Comment
		err := rows.Scan(
			&comment.ID,
			&comment.Content,
			&comment.UserID,
			&comment.ArticleID,
			&comment.CreatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan comment: %w", err)
		}
		comments = append(comments, comment)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error after iteration: %w", err)
	}
	return comments, nil
}
