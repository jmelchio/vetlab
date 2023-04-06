package api

import (
	"context"
	"time"

	"github.com/jmelchio/vetlab/model"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . UserService

// UserService provides the business operations for maintaining users within the application
type UserService interface {
	CreateUser(ctx context.Context, user model.User) (*model.User, error)
	UpdateUser(ctx context.Context, user model.User) (*model.User, error)
	DeleteUser(ctx context.Context, user model.User) error
	UpdatePassword(ctx context.Context, user model.User, password string) (*model.User, error)

	Login(ctx context.Context, userName string, password string) (*model.User, error)

	FindUserByUserName(ctx context.Context, userName string) (*model.User, error)
	FindUserByID(ctx context.Context, userID uint) (*model.User, error)
}

//go:generate counterfeiter . CustomerService

// CustomerService provides the business operations for maintaining users within the application
type CustomerService interface {
	CreateCustomer(ctx context.Context, user model.Customer) (*model.Customer, error)
	UpdateCustomer(ctx context.Context, user model.Customer) (*model.Customer, error)
	DeleteCustomer(ctx context.Context, user model.Customer) error
	UpdatePassword(ctx context.Context, user model.Customer, password string) (*model.Customer, error)

	Login(ctx context.Context, userName string, password string) (*model.Customer, error)

	FindCustomerByUserName(ctx context.Context, userName string) (*model.Customer, error)
	FindCustomerByID(ctx context.Context, userID uint) (*model.Customer, error)
	FindCustomerByVetOrg(ctx context.Context, vetOrg model.VetOrg) ([]model.Customer, error)
}

//go:generate counterfeiter . DiagnosticRequestService

// DiagnosticRequestService provides the business operations for requesting diagnostic reports
type DiagnosticRequestService interface {
	SubmitDiagnosticRequest(ctx context.Context, diagReq model.DiagnosticRequest) (*model.DiagnosticRequest, error)

	FindRequestByDateRange(ctx context.Context, start time.Time, end time.Time, vetOrg model.VetOrg) ([]model.DiagnosticRequest, error)
	FindRequestByID(ctx context.Context, requestID uint) (*model.DiagnosticRequest, error)
	FindRequestByVetOrg(ctx context.Context, vetOrg model.VetOrg) ([]model.DiagnosticRequest, error)
	FindRequestByUser(ctx context.Context, user model.User) ([]model.DiagnosticRequest, error)
	FindRequestByCustomer(ctx context.Context, customer model.Customer) ([]model.DiagnosticRequest, error)
}

//go:generate counterfeiter . DiagnosticReportService

// DiagnosticReportService provides the business operations for providing diagnostic reports
type DiagnosticReportService interface {
	SumbmitDiagnosticReport(ctx context.Context, diagReport model.DiagnosticReport) (*model.DiagnosticReport, error)

	FindReportByDateRange(ctx context.Context, start time.Time, end time.Time, vetOrg model.VetOrg) ([]model.DiagnosticReport, error)
	FindReportByID(ctx context.Context, reportID uint) (*model.DiagnosticReport, error)
	FindReportByVetOrg(ctx context.Context, vetOrg model.VetOrg) ([]model.DiagnosticReport, error)
	FindReportByUser(ctx context.Context, user model.User) ([]model.DiagnosticReport, error)
	FindReportByCustomer(ctx context.Context, customer model.Customer) ([]model.DiagnosticReport, error)
}

//go:generate counterfeiter . VetOrgService

// VetOrgService provides the business operations for maintaining veterinary practices
// within the application
type VetOrgService interface {
	CreateVetOrg(ctx context.Context, vetOrg model.VetOrg) (*model.VetOrg, error)
	UpdateVetOrg(ctx context.Context, vetOrg model.VetOrg) (*model.VetOrg, error)
	DeleteVetOrg(ctx context.Context, vetOrg model.VetOrg) error

	AddUserToVetOrg(ctx context.Context, user model.User, vetOrg model.VetOrg) (*model.User, error)

	FindVetOrgByName(ctx context.Context, orgName string) ([]model.VetOrg, error)
	FindVetOrgByID(ctx context.Context, orgID uint) (*model.VetOrg, error)
}
