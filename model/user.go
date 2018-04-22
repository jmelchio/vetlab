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

//go:generate counterfeiter . UserRepo

// UserRepo declares the persistence interface for the User struct
type UserRepo interface {
	Create(User) (*User, error)
	Update(User) (*User, error)
	Delete(User) error
	GetByID(string) error
	GetByOrgID(string) ([]User, error)
}
