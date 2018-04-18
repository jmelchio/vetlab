package model

// VetOrg declares the structure for veterinary practice data
type VetOrg struct {
	OrgID       string
	OrgName     string
	Street      string
	HouseNumber string
	City        string
	Province    string
	Country     string
	PostalCode  string
	Email       string
	Phone       string
	Fax         string
}

//go:generate counterfeiter . VetOrgRepo

// VetOrgRepo declares the persistence interface for the VetOrg struct
type VetOrgRepo interface {
	Create(VetOrg) (*VetOrg, error)
	Update(VetOrg) (*VetOrg, error)
	Delete(VetOrg) error
	GetByID(string) (*VetOrg, error)
	GetByName(string) ([]VetOrg, error)
}
