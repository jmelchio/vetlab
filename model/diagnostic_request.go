package model

import (
	"time"

	"gorm.io/gorm"
)

// DiagnosticRequest describes the structure of a request for diagnostics
type DiagnosticRequest struct {
	gorm.Model
	VetOrgID    uint       `json:"vet_org_id"`
	CustomerID  uint       `json:"customer_id"`
	UserID      uint       `json:"user_id"`
	Date        *time.Time `json:"date,omitempty"`
	Description string     `json:"description,omitempty"`
}
