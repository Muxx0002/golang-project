package actions

import (
	"context"
	"fmt"

	"github.com/Muxx0002/golang-project/tree/main/backend/internal/database/postgres"
	"github.com/Muxx0002/golang-project/tree/main/backend/internal/models"
)

func UpdateUsername(id string, Username string) error {
	query := `
        UPDATE users 
        SET 
            username = $1,
            update_at = CURRENT_TIMESTAMP 
        WHERE id = $2
    `
	_, err := postgres.DB.Exec(context.Background(), query, Username, id)
	if err != nil {
		return fmt.Errorf("error updating username: %w", err)
	}
	return nil
}

func IsUsernameExists(username string) (bool, error) {
	var exists bool
	query := `SELECT EXISTS(SELECT 1 FROM users WHERE username = $1)`
	err := postgres.DB.QueryRow(context.Background(), query, username).Scan(&exists)
	return exists, err
}

func GetAllUsers() ([]models.Users, error) {
	users := make([]models.Users, 0)
	query := `
        SELECT 
            id, 
            token, 
            email, 
            username, 
            password_hash, 
            role, 
            created_at, 
            update_at 
        FROM users
    `

	rows, err := postgres.DB.Query(context.Background(), query)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch users: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var user models.Users
		err := rows.Scan(
			&user.ID,
			&user.Token,
			&user.Email,
			&user.Username,
			&user.PasswordHash,
			&user.Role,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan user: %w", err)
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error after rows iteration: %w", err)
	}
	return users, nil
}

func DeleteUserByID(userID string) error {
	query := `DELETE FROM users WHERE id = $1`
	result, err := postgres.DB.Exec(
		context.Background(),
		query,
		userID,
	)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}
	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("user with id %s not found", userID)
	}
	return nil
}
