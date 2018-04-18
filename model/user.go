package model

// User declares the structure for a user in the system
type User struct {
	UserID       string
	UserName     string
	FirstName    string
	LastName     string
	Email        string
	PasswordHash string
	OrgID        string
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
