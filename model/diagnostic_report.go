package model

import "time"

// DiagnosticReport describes the structure of a veterinary lab report
type DiagnosticReport struct {
	ReportID   string    `json:"report_id"`
	OrgID      string    `json:"org_id"`
	CustomerID string    `json:"customer_id"`
	UserID     string    `json:"user_id"`
	Date       time.Time `json:"date"`
	ReportBody string    `json:"report_body"`
	ReportFile string    `json:"report_file"`
}
