package model

import "time"

// DiagnosticRequest describes the structure of a request for diagnostics
type DiagnosticRequest struct {
	RequestID   string
	OrgID       string
	CustomerID  string
	UserID      string
	Date        time.Time
	Description string
}

//go:generate counterfeiter . DiagnosticRequestRepo

// DiagnosticRequestRepo describes the persistence interface for diagnostic requests
type DiagnosticRequestRepo interface {
	Create(DiagnosticRequest) (*DiagnosticRequest, error)
	Update(DiagnosticRequest) (*DiagnosticRequest, error)
	Delete(DiagnosticRequest) error
	GetByID(string) (*DiagnosticRequest, error)
	GetByOrgID(string) ([]DiagnosticRequest, error)
}
