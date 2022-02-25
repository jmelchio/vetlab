package sql

import (
	"github.com/jmelchio/vetlab/model"
	"gorm.io/gorm"
)

type DiagnosticReportRepo struct {
	Database *gorm.DB
}

func (diagnosticReportRepo *DiagnosticReportRepo) Create(diagnosticReport *model.DiagnosticReport) error {
	//TODO implement me
	panic("implement me")
}

func (diagnosticReportRepo *DiagnosticReportRepo) Update(diagnosticReport *model.DiagnosticReport) error {
	//TODO implement me
	panic("implement me")
}

func (diagnosticReportRepo *DiagnosticReportRepo) Delete(name string) error {
	//TODO implement me
	panic("implement me")
}

func (diagnosticReportRepo *DiagnosticReportRepo) GetByID(diagnosticReportID uint) (*model.DiagnosticReport, error) {
	//TODO implement me
	panic("implement me")
}

func (diagnosticReportRepo *DiagnosticReportRepo) GetByVetOrgID(vetOrgID uint) ([]model.DiagnosticReport, error) {
	//TODO implement me
	panic("implement me")
}

func (diagnosticReportRepo *DiagnosticReportRepo) GetByUserID(userID uint) ([]model.DiagnosticRequest, error) {
	//TODO implement me
	panic("implement me")
}

func (diagnosticReportRepo *DiagnosticReportRepo) GetByCustomerID(customerID uint) ([]model.DiagnosticRequest, error) {
	//TODO implement me
	panic("implement me")
}
