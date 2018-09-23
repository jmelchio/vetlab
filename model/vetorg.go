package model

// VetOrg declares the structure for veterinary practice data
type VetOrg struct {
	ID          uint    `json:"id" gorm:"primary_key"`
	OrgName     *string `json:"org_name" gorm:"unique_index;not null"`
	Street      *string `json:"street" gorm:"not null"`
	HouseNumber *string `json:"house_number" gorm:"not null"`
	City        *string `json:"city" gorm:"not null"`
	Province    *string `json:"province" gorm:"not null"`
	Country     *string `json:"country" gorm:"not null"`
	PostalCode  *string `json:"postal_code" gorm:"not null"`
	Email       *string `json:"email" gorm:"not null"`
	Phone       *string `json:"phone" gorm:"not null"`
	Fax         *string `json:"fax"`
}
