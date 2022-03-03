package sql

import (
	"errors"
	"fmt"

	"github.com/jmelchio/vetlab/model"
	"gorm.io/gorm"
)

type DiagnosticRequestRepo struct {
	Database *gorm.DB
}

// Create creates a persistent DiagnosticRequest row in the sql datastore
// When a given DiagnosticRequest already has a non-zero ID an error will be returned
func (diagnosticRequestRepo *DiagnosticRequestRepo) Create(diagnosticRequest *model.DiagnosticRequest) error {
	if diagnosticRequest.ID == 0 {
		if err := diagnosticRequestRepo.Database.Create(diagnosticRequest).Error; err != nil {
			return err
		}
		return nil
	}
	return errors.New("record already in database")
}

func (diagnosticRequestRepo *DiagnosticRequestRepo) Update(diagnosticRequest *model.DiagnosticRequest) error {
	if diagnosticRequest.ID != 0 {
		if err := diagnosticRequestRepo.Database.Save(diagnosticRequest).Error; err != nil {
			return err
		}
		return nil
	}
	return errors.New("record does not exist in database")
}

func (diagnosticRequestRepo *DiagnosticRequestRepo) Delete(diagnosticRequest *model.DiagnosticRequest) error {
	return diagnosticRequestRepo.Database.Delete(diagnosticRequest).Error
}

func (diagnosticRequestRepo *DiagnosticRequestRepo) GetByID(diagnosticRequestID uint) (*model.DiagnosticRequest, error) {
	var diagnosticRequest model.DiagnosticRequest

	if err := diagnosticRequestRepo.Database.First(&diagnosticRequest, diagnosticRequestID).Error; err != nil {
		return nil, err
	}
	return &diagnosticRequest, nil
}

func (diagnosticRequestRepo *DiagnosticRequestRepo) GetByVetOrgID(vetOrgID uint) ([]model.DiagnosticRequest, error) {
	var diagnosticRequests []model.DiagnosticRequest

	result := diagnosticRequestRepo.Database.Where("vet_org_id = ?", vetOrgID).Find(&diagnosticRequests)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("diagnosticRequests with vetOrgId '%d' not found", vetOrgID)
	}
	return diagnosticRequests, nil
}

func (diagnosticRequestRepo *DiagnosticRequestRepo) GetByUserID(userID uint) ([]model.DiagnosticRequest, error) {
	var diagnosticRequests []model.DiagnosticRequest

	result := diagnosticRequestRepo.Database.Where("user_id = ?", userID).Find(&diagnosticRequests)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("diagnosticRequests with userId '%d' not found", userID)
	}
	return diagnosticRequests, nil
}

func (diagnosticRequestRepo *DiagnosticRequestRepo) GetByCustomerID(customerID uint) ([]model.DiagnosticRequest, error) {
	return nil, nil
}
