package model

// User declares the structure for a user in the system
type User struct {
	UserID       string `json:"user_id"`
	UserName     string `json:"user_name"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
	OrgID        string `json:"org_id"`
	AdminUser    bool   `json:"admin_user"`
}
