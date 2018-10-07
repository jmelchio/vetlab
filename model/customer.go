package model

// Customer declares the structure for a customer in the system
// UserName is a pointer because we want to catch if it's null in
// the database.
type Customer struct {
	ID                 uint                `json:"id" gorm:"primary_key"`
	UserName           *string             `json:"user_name" gorm:"not null;unique_index"`
	FirstName          string              `json:"first_name,omitempty"`
	LastName           string              `json:"last_name,omitempty"`
	Email              string              `json:"email,omitempty"`
	Password           string              `json:"password,omitempty"`
	VetOrgID           uint                `json:"vet_org_id"`
	DiagnosticReports  []DiagnosticReport  `json:"diagnostic_reports,omitempty"`
	DiagnosticRequests []DiagnosticRequest `json:"diagnostic_requests,omitempty"`
}
