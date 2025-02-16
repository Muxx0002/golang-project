package dto

type AuthData struct {
	Email        string `json:"email"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	RestoreToken string `json:"restore_token"`
}
