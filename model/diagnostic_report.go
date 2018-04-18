package model

import "time"

// DiagnosticReport describes the structure of a veterinary lab report
type DiagnosticReport struct {
	ReportID   string
	OrgID      string
	CustomerID string
	UserID     string
	Date       time.Time
	ReportBody string
	ReportFile string
}

//go:generate counterfeiter . DiagnosticReportRepo

// DiagnosticReportRepo describes the persistence interface for a veterinary lab report
type DiagnosticReportRepo interface {
	Create(DiagnosticReport) (*DiagnosticReport, error)
	Update(DiagnosticReport) (*DiagnosticReport, error)
	Delete(string) error
	GetByID(string) (*DiagnosticReport, error)
	GetByOrgID(string) ([]DiagnosticReport, error)
}
