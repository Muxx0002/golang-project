package dto

import "time"

type AuthData struct {
	Email        string `json:"email"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	RestoreToken string `json:"restore_token"`
}

type ProfileResponse struct {
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}

type UpdateRequest struct {
	Username string `json:"username"`
}

type ArticleData struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type CommentData struct {
	Content string `json:"content"`
}
