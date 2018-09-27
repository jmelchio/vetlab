package service

import "github.com/jmelchio/vetlab/model"

//go:generate counterfeiter . UserRepo

// UserRepo declares the persistence interface for the User struct
type UserRepo interface {
	Create(*model.User) error
	Update(*model.User) error
	Delete(*model.User) error
	GetByID(uint) (*model.User, error)
	GetByUserName(string) (*model.User, error)
}

//go:generate counterfeiter . DiagnosticReportRepo

// DiagnosticReportRepo describes the persistence interface for a veterinary lab report
type DiagnosticReportRepo interface {
	Create(*model.DiagnosticReport) error
	Update(*model.DiagnosticReport) error
	Delete(string) error
	GetByID(uint) (*model.DiagnosticReport, error)
	GetByVetOrgID(uint) ([]model.DiagnosticReport, error)
	GetByUserID(uint) ([]model.DiagnosticRequest, error)
}

//go:generate counterfeiter . DiagnosticRequestRepo

// DiagnosticRequestRepo describes the persistence interface for diagnostic requests
type DiagnosticRequestRepo interface {
	Create(*model.DiagnosticRequest) error
	Update(*model.DiagnosticRequest) error
	Delete(*model.DiagnosticRequest) error
	GetByID(uint) (*model.DiagnosticRequest, error)
	GetByVetOrgID(uint) ([]model.DiagnosticRequest, error)
	GetByUserID(uint) ([]model.DiagnosticRequest, error)
}

//go:generate counterfeiter . VetOrgRepo

// VetOrgRepo declares the persistence interface for the model.VetOrg struct
type VetOrgRepo interface {
	Create(*model.VetOrg) error
	Update(*model.VetOrg) error
	Delete(*model.VetOrg) error
	GetByID(uint) (*model.VetOrg, error)
	GetByName(string) ([]model.VetOrg, error)
}
