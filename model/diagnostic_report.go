package model

import (
	"time"

	"gorm.io/gorm"
)

// DiagnosticReport describes the structure of a veterinary lab report
type DiagnosticReport struct {
	gorm.Model
	RequestID  uint       `json:"request_id"`
	VetOrgID   uint       `json:"vet_org_id"`
	CustomerID uint       `json:"customer_id"`
	UserID     uint       `json:"user_id"`
	Date       *time.Time `json:"date,omitempty"`
	ReportBody string     `json:"report_body,omitempty"`
	ReportFile string     `json:"report_file,omitempty"`
}
