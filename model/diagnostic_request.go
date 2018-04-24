package model

import "time"

// DiagnosticRequest describes the structure of a request for diagnostics
type DiagnosticRequest struct {
	RequestID   string    `json:"request_id"`
	OrgID       string    `json:"org_id"`
	CustomerID  string    `json:"customer_id"`
	UserID      string    `json:"user_id"`
	Date        time.Time `json:"date"`
	Description string    `json:"description"`
}
