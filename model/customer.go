package model

// Customer declares the structure for a customer in the system
type Customer struct {
	ID        uint    `json:"id" gorm:"primary_key"`
	UserName  *string `json:"user_name" gorm:"not null;unique_index"`
	FirstName *string `json:"first_name"`
	LastName  *string `json:"last_name"`
	Email     *string `json:"email"`
	Password  *string `json:"password"`
	VetOrgID  uint    `json:"vet_org_id"`
}
