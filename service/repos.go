package service

import "github.com/jmelchio/vetlab/model"

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

// UserRepo declares the persistence interface for the User struct
//
//counterfeiter:generate . UserRepo
type UserRepo interface {
	Create(*model.User) error
	Update(*model.User) error
	Delete(*model.User) error
	GetByID(uint) (*model.User, error)
	GetByUserName(string) (*model.User, error)
}

// CustomerRepo declares the persistence interface for the Customer struct
//
//counterfeiter:generate . CustomerRepo
type CustomerRepo interface {
	Create(*model.Customer) error
	Update(*model.Customer) error
	Delete(*model.Customer) error
	GetByID(uint) (*model.Customer, error)
	GetByVetOrgID(uint) ([]model.Customer, error)
	GetByUserName(string) (*model.Customer, error)
}

// DiagnosticReportRepo describes the persistence interface for a veterinary lab report
//
//counterfeiter:generate . DiagnosticReportRepo
type DiagnosticReportRepo interface {
	Create(*model.DiagnosticReport) error
	Update(*model.DiagnosticReport) error
	Delete(report *model.DiagnosticReport) error
	GetByID(uint) (*model.DiagnosticReport, error)
	GetByVetOrgID(uint) ([]model.DiagnosticReport, error)
	GetByUserID(uint) ([]model.DiagnosticReport, error)
	GetByCustomerID(uint) ([]model.DiagnosticReport, error)
}

// DiagnosticRequestRepo describes the persistence interface for diagnostic requests
//
//counterfeiter:generate . DiagnosticRequestRepo
type DiagnosticRequestRepo interface {
	Create(*model.DiagnosticRequest) error
	Update(*model.DiagnosticRequest) error
	Delete(*model.DiagnosticRequest) error
	GetByID(uint) (*model.DiagnosticRequest, error)
	GetByVetOrgID(uint) ([]model.DiagnosticRequest, error)
	GetByUserID(uint) ([]model.DiagnosticRequest, error)
	GetByCustomerID(uint) ([]model.DiagnosticRequest, error)
}

// VetOrgRepo declares the persistence interface for the model.VetOrg struct
//
//counterfeiter:generate . VetOrgRepo
type VetOrgRepo interface {
	Create(*model.VetOrg) error
	Update(*model.VetOrg) error
	Delete(*model.VetOrg) error
	GetByID(uint) (*model.VetOrg, error)
	GetByName(string) ([]model.VetOrg, error)
}
