package model

import "gorm.io/gorm"

// User declares the structure for a user in the system
type User struct {
	gorm.Model
	UserName           *string             `json:"user_name" gorm:"uniqueIndex;not null"`
	FirstName          string              `json:"first_name,omitempty"`
	LastName           string              `json:"last_name,omitempty"`
	Email              string              `json:"email,omitempty"`
	Password           string              `json:"password,omitempty"`
	AdminUser          bool                `json:"admin_user"`
	DiagnosticReports  []DiagnosticReport  `json:"diagnostic_reports,omitempty"`
	DiagnosticRequests []DiagnosticRequest `json:"diagnostic_requests,omitempty"`
}
