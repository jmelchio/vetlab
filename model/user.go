package model

// User declares the structure for a user in the system
type User struct {
	ID        uint    `json:"id" gorm:"primary_key"`
	UserName  *string `json:"user_name" gorm:"unique_index;not null"`
	FirstName string  `json:"first_name,omitempty"`
	LastName  string  `json:"last_name,omitempty"`
	Email     string  `json:"email,omitempty"`
	Password  string  `json:"password,omitempty"`
	AdminUser bool    `json:"admin_user"`
}
