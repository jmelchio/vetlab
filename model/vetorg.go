package model

import "gorm.io/gorm"

// VetOrg declares the structure for veterinary practice data
type VetOrg struct {
	gorm.Model
	OrgName            *string             `json:"org_name" gorm:"uniqueIndex;not null"`
	Street             string              `json:"street,omitempty"`
	HouseNumber        string              `json:"house_number,omitempty"`
	City               string              `json:"city,omitempty"`
	Province           string              `json:"province,omitempty"`
	Country            string              `json:"country,omitempty"`
	PostalCode         string              `json:"postal_code,omitempty"`
	Email              string              `json:"email,omitempty"`
	Phone              string              `json:"phone,omitempty"`
	Fax                string              `json:"fax,omitempty"`
	Customers          []Customer          `json:"customers,omitempty"`
	DiagnosticReports  []DiagnosticReport  `json:"diagnostic_reports,omitempty"`
	DiagnosticRequests []DiagnosticRequest `json:"diagnostic_requests,omitempty"`
}
