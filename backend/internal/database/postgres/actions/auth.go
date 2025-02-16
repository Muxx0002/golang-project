package actions

import (
	"context"
	"errors"
	"fmt"

	"time"

	"github.com/Muxx0002/golang-project/tree/main/backend/internal/database/postgres"
	"github.com/Muxx0002/golang-project/tree/main/backend/internal/models"
	"github.com/Muxx0002/golang-project/tree/main/backend/pkg/crypto"
	"github.com/Muxx0002/golang-project/tree/main/backend/pkg/tools"
	"github.com/jackc/pgx"
)

func CreateUser(email, password, username, token string) error {
	id := tools.GenerateID()
	query := `INSERT INTO users (id, token, email, username, password_hash, role, created_at, update_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8);`
	_, err := postgres.DB.Exec(context.Background(), query, id, token, email, username, crypto.GetMD5Hash(password), "user", time.Now(), time.Now())
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}
	return nil
}

func CheckUser(password, email string) (models.Users, error) {
	var user models.Users
	query := "SELECT * FROM users WHERE email = $1 AND password_hash = $2;"
	err := postgres.DB.QueryRow(context.Background(), query, email, crypto.GetMD5Hash(password)).Scan(
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
		if errors.Is(err, pgx.ErrNoRows) {
			return user, errors.New("invalid email or password")
		}
		return user, fmt.Errorf("failed to query user: %w", err)
	}
	return user, nil
}

func CheckUserByToken(token string) (models.Users, error) {
	var user models.Users
	query := "SELECT * FROM users WHERE token = $1;"
	err := postgres.DB.QueryRow(context.Background(), query, token).Scan(
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
		return user, errors.New("user not found")
	}
	return user, nil
}

func DeleteUser(token string) error {
	query := `DELETE FROM users WHERE token = $1`
	_, err := postgres.DB.Exec(context.Background(), query, token)
	if err != nil {
		return errors.New("failed to delete user")
	}
	return nil
}
