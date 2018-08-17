package model

import "time"

// DiagnosticReport describes the structure of a veterinary lab report
type DiagnosticReport struct {
	ID         uint      `json:"id" gorm:"primary_key"`
	RequestID  uint      `json:"request_id"`
	OrgID      uint      `json:"org_id"`
	CustomerID uint      `json:"customer_id"`
	UserID     uint      `json:"user_id"`
	Date       time.Time `json:"date"`
	ReportBody string    `json:"report_body"`
	ReportFile string    `json:"report_file"`
}
