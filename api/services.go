package api

import (
	"context"
	"time"

	"github.com/jmelchio/vetlab/model"
)

//go:generate counterfeiter . UserService

// UserService provides the business operations for maintaining users within the application
type UserService interface {
	CreateUser(context.Context, model.User) (*model.User, error)
	UpdateUser(context.Context, model.User) (*model.User, error)
	DeleteUser(context.Context, model.User) error

	Login(context.Context, string, string) (*model.User, error)

	FindUsersByVetOrg(context.Context, model.VetOrg) ([]model.User, error)
	FindUsersByName(context.Context, string) ([]model.User, error)
	FindUserByID(context.Context, string) (*model.User, error)
}

//go:generate counterfeiter . ReportService

// ReportService provides the business operations for requesting,finding and retrieving
// diagnostic reports
type ReportService interface {
	SubmitDiagnosticRequest(context.Context, model.DiagnosticRequest) (*model.DiagnosticRequest, error)
	FindReportByDateRange(context.Context, time.Time, time.Time, model.VetOrg) ([]model.DiagnosticReport, error)
	FindReportByID(context.Context, string) (model.DiagnosticReport, error)
	FindReportByVetOrg(context.Context, model.VetOrg) ([]model.DiagnosticReport, error)
	FindReportByUser(context.Context, model.User) ([]model.DiagnosticReport, error)

	FindRequestByDateRange(context.Context, time.Time, time.Time, model.VetOrg) ([]model.DiagnosticRequest, error)
	FindRequestByID(context.Context, string) (*model.DiagnosticRequest, error)
	FindRequestByVetOrg(context.Context, model.VetOrg) ([]model.DiagnosticRequest, error)
	FindRequestByUser(context.Context, model.User) ([]model.DiagnosticRequest, error)
}

//go:generate counterfeiter . VetOrgService

// VetOrgService provides the business operations for maintianing veterinary practices
// within the application
type VetOrgService interface {
	CreateVetOrg(context.Context, model.VetOrg) (*model.VetOrg, error)
	UpdateVetOrg(context.Context, model.VetOrg) (*model.VetOrg, error)
	DeteleVetOrg(context.Context, model.VetOrg) error

	AddUserToVetOrg(context.Context, model.User, model.VetOrg) (*model.User, error)

	FindVetOrgByName(context.Context, string) (*model.VetOrg, error)
	FindVetOrgByID(context.Context, string) (*model.VetOrg, error)
}
