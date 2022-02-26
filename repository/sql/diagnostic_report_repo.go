package sql

import (
	"errors"

	"github.com/jmelchio/vetlab/model"
	"gorm.io/gorm"
)

type DiagnosticReportRepo struct {
	Database *gorm.DB
}

// Create creates a persistent DiagnosticReport row in the sql datastore
// When a given DiagnosticReport already has a non-zero ID an error will be returned
func (diagnosticReportRepo *DiagnosticReportRepo) Create(diagnosticReport *model.DiagnosticReport) error {
	if diagnosticReport.ID == 0 {
		if err := diagnosticReportRepo.Database.Create(diagnosticReport).Error; err != nil {
			return err
		}
		return nil
	}
	return errors.New("record already in database")
}

func (diagnosticReportRepo *DiagnosticReportRepo) Update(diagnosticReport *model.DiagnosticReport) error {
	return errors.New("not implemented")
}

func (diagnosticReportRepo *DiagnosticReportRepo) Delete(name string) error {
	return errors.New("not implemented")
}

func (diagnosticReportRepo *DiagnosticReportRepo) GetByID(diagnosticReportID uint) (*model.DiagnosticReport, error) {
	return nil, errors.New("not implemented")
}

func (diagnosticReportRepo *DiagnosticReportRepo) GetByVetOrgID(vetOrgID uint) ([]model.DiagnosticReport, error) {
	return nil, errors.New("not implemented")
}

func (diagnosticReportRepo *DiagnosticReportRepo) GetByUserID(userID uint) ([]model.DiagnosticRequest, error) {
	return nil, errors.New("not implemented")
}

func (diagnosticReportRepo *DiagnosticReportRepo) GetByCustomerID(customerID uint) ([]model.DiagnosticRequest, error) {
	return nil, errors.New("not implemented")
}
