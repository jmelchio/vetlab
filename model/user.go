package model

// User declares the structure for a user in the system
type User struct {
	ID        uint    `json:"id" gorm:"primary_key"`
	UserName  *string `json:"user_name" gorm:"unique_index;not null"`
	FirstName *string `json:"first_name"`
	LastName  *string `json:"last_name"`
	Email     *string `json:"email"`
	Password  *string `json:"password"`
	AdminUser bool    `json:"admin_user"`
}
