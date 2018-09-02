package model

// LoginRequest is the structure passed when logging into the system
type LoginRequest struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}
