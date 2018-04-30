package service

import "github.com/jmelchio/vetlab/model"

//go:generate counterfeiter . UserRepo

// UserRepo declares the persistence interface for the User struct
type UserRepo interface {
	Create(model.User) (*model.User, error)
	Update(model.User) (*model.User, error)
	Delete(model.User) error
	GetByID(string) (*model.User, error)
	GetByOrgID(string) ([]model.User, error)
	GetByUserNamePassword(string, string) (model.User, error)
}

//go:generate counterfeiter . DiagnosticReportRepo

// DiagnosticReportRepo describes the persistence interface for a veterinary lab report
type DiagnosticReportRepo interface {
	Create(model.DiagnosticReport) (*model.DiagnosticReport, error)
	Update(model.DiagnosticReport) (*model.DiagnosticReport, error)
	Delete(string) error
	GetByID(string) (*model.DiagnosticReport, error)
	GetByOrgID(string) ([]model.DiagnosticReport, error)
	GetByUserID(string) ([]model.DiagnosticRequest, error)
}

//go:generate counterfeiter . DiagnosticRequestRepo

// DiagnosticRequestRepo describes the persistence interface for diagnostic requests
type DiagnosticRequestRepo interface {
	Create(model.DiagnosticRequest) (*model.DiagnosticRequest, error)
	Update(model.DiagnosticRequest) (*model.DiagnosticRequest, error)
	Delete(model.DiagnosticRequest) error
	GetByID(string) (*model.DiagnosticRequest, error)
	GetByOrgID(string) ([]model.DiagnosticRequest, error)
	GetByUserID(string) ([]model.DiagnosticRequest, error)
}

//go:generate counterfeiter . VetOrgRepo

// VetOrgRepo declares the persistence interface for the model.VetOrg struct
type VetOrgRepo interface {
	Create(model.VetOrg) (*model.VetOrg, error)
	Update(model.VetOrg) (*model.VetOrg, error)
	Delete(model.VetOrg) error
	GetByID(string) (*model.VetOrg, error)
	GetByName(string) ([]model.VetOrg, error)
}
