package model

import "time"

// DiagnosticRequest describes the structure of a request for diagnostics
type DiagnosticRequest struct {
	ID          uint      `json:"id" gorm:"primary_key"`
	OrgID       uint      `json:"org_id"`
	CustomerID  uint      `json:"customer_id"`
	UserID      uint      `json:"user_id"`
	Date        time.Time `json:"date"`
	Description string    `json:"description"`
}
