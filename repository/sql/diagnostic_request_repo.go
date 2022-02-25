package sql

import (
	"github.com/jmelchio/vetlab/model"
	"gorm.io/gorm"
)

type DiagnosticRequestRepo struct {
	Database *gorm.DB
}

func (diagnosticRequestRepo *DiagnosticRequestRepo) Create(diagnosticRequest *model.DiagnosticRequest) error {
	return nil
}

func (diagnosticRequestRepo *DiagnosticRequestRepo) Update(diagnosticRequest *model.DiagnosticRequest) error {
	return nil
}

func (diagnosticRequestRepo *DiagnosticRequestRepo) Delete(diagnosticRequest *model.DiagnosticRequest) error {
	return nil
}

func (diagnosticRequestRepo *DiagnosticRequestRepo) GetByID(diagnosticRequestID uint) (*model.DiagnosticRequest, error) {
	return nil, nil
}

func (diagnosticRequestRepo *DiagnosticRequestRepo) GetByVetOrgID(vetOrgID uint) ([]model.DiagnosticRequest, error) {
	return nil, nil
}

func (diagnosticRequestRepo *DiagnosticRequestRepo) GetByUserID(userID uint) ([]model.DiagnosticRequest, error) {
	return nil, nil
}

func (diagnosticRequestRepo *DiagnosticRequestRepo) GetByCustomerID(customerID uint) ([]model.DiagnosticRequest, error) {
	return nil, nil
}
