package model

// VetOrg declares the structure for veterinary practice data
type VetOrg struct {
	OrgID       string `json:"org_id"`
	OrgName     string `json:"org_name"`
	Street      string `json:"street"`
	HouseNumber string `json:"house_number"`
	City        string `json:"city"`
	Province    string `json:"province"`
	Country     string `json:"country"`
	PostalCode  string `json:"postal_code"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	Fax         string `json:"fax"`
}
