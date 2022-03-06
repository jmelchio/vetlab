package sql

import (
	"errors"

	"github.com/jmelchio/vetlab/model"
	"gorm.io/gorm"
)

type VetOrgRepo struct {
	Database *gorm.DB
}

func (vetOrgRepo VetOrgRepo) Create(vetOrg *model.VetOrg) error {
	if vetOrg.ID == 0 {
		return vetOrgRepo.Database.Create(vetOrg).Error
	}
	return errors.New("record already in database")
}

func (vetOrgRepo VetOrgRepo) Update(vetOrg *model.VetOrg) error {
	return errors.New("not yet implemented")
}

func (vetOrgRepo VetOrgRepo) Delete(vetOrg *model.VetOrg) error {
	return errors.New("not yet implemented")
}

func (vetOrgRepo VetOrgRepo) GetByID(vetOrgID uint) (*model.VetOrg, error) {
	return nil, errors.New("not yet implemented")
}

func (vetOrgRepo VetOrgRepo) GetByName(vetOrgName string) ([]model.VetOrg, error) {
	return nil, errors.New("not yet implemented")
}
