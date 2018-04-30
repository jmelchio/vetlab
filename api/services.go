package api

import (
	"context"
	"time"

	"github.com/jmelchio/vetlab/model"
)

//go:generate counterfeiter . UserService

// UserService provides the business operations for maintaining users within the application
type UserService interface {
	CreateUser(model.User, context.Context) (model.User, error)
	UpdateUser(model.User, context.Context) (model.User, error)
	DeleteUser(model.User, context.Context) error

	Login(string, string, context.Context) (model.User, error)

	FindUsersByVetOrg(model.VetOrg, context.Context) ([]model.User, error)
	FindUsersByName(string, context.Context) ([]model.User, error)
	FindUserByID(string, context.Context) (model.User, error)
}

//go:generate counterfeiter . ReportService

// ReportService provides the business operations for requesting,finding and retrieving
// diagnostic reports
type ReportService interface {
	SubmitDiagnosticRequest(model.DiagnosticRequest, context.Context) (model.DiagnosticRequest, error)
	FindReportByDateRange(time.Time, time.Time, model.VetOrg, context.Context) ([]model.DiagnosticReport, error)
	FindReportByID(string, context.Context) (model.DiagnosticReport, error)
	FindReportByVetOrg(model.VetOrg, context.Context) ([]model.DiagnosticReport, error)
	FindReportByUser(model.User, context.Context) ([]model.DiagnosticReport, error)

	FindRequestByDateRange(time.Time, time.Time, model.VetOrg, context.Context) ([]model.DiagnosticRequest, error)
	FindRequestByID(string, context.Context) (model.DiagnosticRequest, error)
	FindRequestByVetOrg(model.VetOrg, context.Context) ([]model.DiagnosticRequest, error)
	FindRequestByUser(model.User, context.Context) ([]model.DiagnosticRequest, error)
}

//go:generate counterfeiter . VetOrgService

// VetOrgService provides the business operations for maintianing veterinary practices
// within the application
type VetOrgService interface {
	CreateVetOrg(model.VetOrg, context.Context) (model.VetOrg, error)
	UpdateVetOrg(model.VetOrg, context.Context) (model.VetOrg, error)
	DeteleVetOrg(model.VetOrg, context.Context) error

	AddUserToVetOrg(model.User, model.VetOrg, context.Context) (model.User, error)

	FindVetOrgByName(string, context.Context) (model.VetOrg, error)
	FindVetOrgByID(string, context.Context) (model.VetOrg, error)
}
