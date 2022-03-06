package sql

import (
	"errors"
	"fmt"

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
		return diagnosticReportRepo.Database.Create(diagnosticReport).Error
	}
	return errors.New("record already in database")
}

func (diagnosticReportRepo *DiagnosticReportRepo) Update(diagnosticReport *model.DiagnosticReport) error {
	if diagnosticReport.ID != 0 {
		if err := diagnosticReportRepo.Database.Save(diagnosticReport).Error; err != nil {
			return err
		}
		return nil
	}
	return errors.New("record does not exist in database")
}

func (diagnosticReportRepo *DiagnosticReportRepo) Delete(diagnosticReport *model.DiagnosticReport) error {
	return diagnosticReportRepo.Database.Delete(diagnosticReport).Error
}

func (diagnosticReportRepo *DiagnosticReportRepo) GetByID(diagnosticReportID uint) (*model.DiagnosticReport, error) {
	var diagnosticReport model.DiagnosticReport

	if err := diagnosticReportRepo.Database.First(&diagnosticReport, diagnosticReportID).Error; err != nil {
		return nil, err
	}
	return &diagnosticReport, nil
}

func (diagnosticReportRepo *DiagnosticReportRepo) GetByVetOrgID(vetOrgID uint) ([]model.DiagnosticReport, error) {
	var diagnosticReports []model.DiagnosticReport

	result := diagnosticReportRepo.Database.Where("vet_org_id = ?", vetOrgID).Find(&diagnosticReports)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("diagnosticReports with vetOrgId '%d' not found", vetOrgID)
	}
	return diagnosticReports, nil
}

func (diagnosticReportRepo *DiagnosticReportRepo) GetByUserID(userID uint) ([]model.DiagnosticReport, error) {
	var diagnosticReports []model.DiagnosticReport

	result := diagnosticReportRepo.Database.Where("user_id = ?", userID).Find(&diagnosticReports)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("diagnosticReports with userId '%d' not found", userID)
	}
	return diagnosticReports, nil
}

func (diagnosticReportRepo *DiagnosticReportRepo) GetByCustomerID(customerID uint) ([]model.DiagnosticReport, error) {
	var diagnosticReports []model.DiagnosticReport

	result := diagnosticReportRepo.Database.Where("customer_id = ?", customerID).Find(&diagnosticReports)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("diagnosticReports with customerId '%d' not found", customerID)
	}
	return diagnosticReports, nil
}
