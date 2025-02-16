package dto

type AuthData struct {
	Email        string `json:"email"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	RestoreToken string `json:"restore_token"`
}

type ArticleData struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type CommentData struct {
	Content string `json:"content"`
}
