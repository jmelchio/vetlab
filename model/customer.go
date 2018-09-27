package model

// Customer declares the structure for a customer in the system
type Customer struct {
	ID        uint    `json:"id" gorm:"primary_key"`
	FirstName *string `json:"first_name"`
	LastName  *string `json:"last_name"`
	Email     *string `json:"email"`
	VetOrgID  uint    `json:"vet_org_id"`
}
