package sql

import (
	"errors"

	"github.com/jmelchio/vetlab/model"
	"gorm.io/gorm"
)

type VetOrgRepo struct {
	Database *gorm.DB
}

func (v VetOrgRepo) Create(vetOrg *model.VetOrg) error {
	return errors.New("not yet implemented")
}

func (v VetOrgRepo) Update(vetOrg *model.VetOrg) error {
	return errors.New("not yet implemented")
}

func (v VetOrgRepo) Delete(vetOrg *model.VetOrg) error {
	return errors.New("not yet implemented")
}

func (v VetOrgRepo) GetByID(vetOrgID uint) (*model.VetOrg, error) {
	return nil, errors.New("not yet implemented")
}

func (v VetOrgRepo) GetByName(vetOrgName string) ([]model.VetOrg, error) {
	return nil, errors.New("not yet implemented")
}
