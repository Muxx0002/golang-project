package models

import "time"

type Users struct {
	ID           string    `json:"id"`
	Token        string    `json:"token"`
	Email        string    `json:"email"`
	Username     string    `json:"username"`
	PasswordHash string    `json:"password_hash"`
	Role         string    `json:"role"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type Article struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	AuthorID  string    `json:"author_id"` // Меняем на string для соответствия типу UUID
	CreatedAt time.Time `json:"created_at"`
}

type Comment struct {
	ID        int       `json:"id"`
	Content   string    `json:"content"`
	ArticleID int       `json:"article_id"`
	UserID    string    `json:"user_id"` // Меняем на string для соответствия типу UUID
	CreatedAt time.Time `json:"created_at"`
}
