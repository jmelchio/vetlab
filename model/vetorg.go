package model

// VetOrg declares the structure for veterinary practice data
type VetOrg struct {
	ID          uint    `json:"id" gorm:"primary_key"`
	OrgName     *string `json:"org_name" gorm:"unique_index;not null"`
	Street      string  `json:"street,omitempty" gorm:"not null"`
	HouseNumber string  `json:"house_number,omitempty" gorm:"not null"`
	City        string  `json:"city,omitempty" gorm:"not null"`
	Province    string  `json:"province,omitempty" gorm:"not null"`
	Country     string  `json:"country,omitempty" gorm:"not null"`
	PostalCode  string  `json:"postal_code,omitempty" gorm:"not null"`
	Email       string  `json:"email,omitempty" gorm:"not null"`
	Phone       string  `json:"phone,omitempty" gorm:"not null"`
	Fax         string  `json:"fax,omitempty"`
}
