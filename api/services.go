package api

import (
	"context"
	"time"

	"github.com/jmelchio/vetlab/model"
)

//go:generate counterfeiter . UserService

// UserService provides the business operations for maintaining users within the application
type UserService interface {
	CreateUser(ctx context.Context, user model.User) (*model.User, error)
	UpdateUser(ctx context.Context, user model.User) (*model.User, error)
	DeleteUser(ctx context.Context, user model.User) error
	UpdatePassword(ctx context.Context, user model.User, password string) (*model.User, error)

	Login(ctx context.Context, userName string, password string) (*model.User, error)

	FindUsersByVetOrg(ctx context.Context, vetOrg model.VetOrg) ([]model.User, error)
	FindUserByUserName(ctx context.Context, userName string) (*model.User, error)
	FindUserByID(ctx context.Context, userID string) (*model.User, error)
}

//go:generate counterfeiter . ReportService

// ReportService provides the business operations for requesting,finding and retrieving
// diagnostic reports
type ReportService interface {
	SubmitDiagnosticRequest(ctx context.Context, diagReq model.DiagnosticRequest) (*model.DiagnosticRequest, error)
	FindReportByDateRange(ctx context.Context, start time.Time, end time.Time, vetOrg model.VetOrg) ([]model.DiagnosticReport, error)
	FindReportByID(ctx context.Context, reportID string) (model.DiagnosticReport, error)
	FindReportByVetOrg(ctx context.Context, vetOrg model.VetOrg) ([]model.DiagnosticReport, error)
	FindReportByUser(ctx context.Context, user model.User) ([]model.DiagnosticReport, error)

	FindRequestByDateRange(ctx context.Context, start time.Time, end time.Time, vetOrg model.VetOrg) ([]model.DiagnosticRequest, error)
	FindRequestByID(ctx context.Context, requestID string) (*model.DiagnosticRequest, error)
	FindRequestByVetOrg(ctx context.Context, vetOrg model.VetOrg) ([]model.DiagnosticRequest, error)
	FindRequestByUser(ctx context.Context, user model.User) ([]model.DiagnosticRequest, error)
}

//go:generate counterfeiter . VetOrgService

// VetOrgService provides the business operations for maintianing veterinary practices
// within the application
type VetOrgService interface {
	CreateVetOrg(ctx context.Context, vetOrg model.VetOrg) (*model.VetOrg, error)
	UpdateVetOrg(ctx context.Context, vetOrg model.VetOrg) (*model.VetOrg, error)
	DeteleVetOrg(ctx context.Context, vetOrg model.VetOrg) error

	AddUserToVetOrg(ctx context.Context, user model.User, vetOrg model.VetOrg) (*model.User, error)

	FindVetOrgByName(ctx context.Context, orgName string) (*model.VetOrg, error)
	FindVetOrgByID(ctx context.Context, orgID string) (*model.VetOrg, error)
}