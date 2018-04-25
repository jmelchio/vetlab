package api

import "github.com/jmelchio/vetlab/model"

//go:generate counterfeiter . UserService

// UserService provides the business operations for maintaining users within the application
type UserService interface {
}

//go:generate counterfeiter . ReportService

// ReportService provides the business operations for requesting,finding and retrieving
// diagnostic reports
type ReportService interface {
	SubmitDiagnosticRequest(model.DiagnosticRequest) (model.DiagnosticRequest, error)
}

//go:generate counterfeiter . VetOrgService

// VetOrgService provides the business operations for maintianing veterinary practices
// within the application
type VetOrgService interface {
}

//go:generate counterfeiter . AdminService

// AdminService provides the business operations for all administrative functions
// within the application
type AdminService interface {
}
