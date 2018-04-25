package api

import "github.com/jmelchio/vetlab/model"

// UserService provides the business operations for maintaining users within the application
type UserService interface {
}

// ReportService provides the business operations for requesting,finding and retrieving
// diagnostic reports
type ReportService interface {
	SubmitDiagnosticRequest(model.DiagnosticRequest) (model.DiagnosticRequest, error)
}

// VetOrgService provides the business operations for maintianing veterinary practices
// within the application
type VetOrgService interface {
}

// AdminService provides the business operations for all administrative functions
// within the application
type AdminService interface {
}
