package sql

import (
	"errors"
	"fmt"

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
	if vetOrg.ID != 0 {
		return vetOrgRepo.Database.Save(vetOrg).Error
	}
	return errors.New("record does exist in database")
}

func (vetOrgRepo VetOrgRepo) Delete(vetOrg *model.VetOrg) error {
	return vetOrgRepo.Database.Delete(vetOrg).Error
}

func (vetOrgRepo VetOrgRepo) GetByID(vetOrgID uint) (*model.VetOrg, error) {
	var vetOrg model.VetOrg

	if err := vetOrgRepo.Database.First(&vetOrg, vetOrgID).Error; err != nil {
		return nil, err
	}
	return &vetOrg, nil
}

func (vetOrgRepo VetOrgRepo) GetByName(vetOrgName string) ([]model.VetOrg, error) {
	var vetOrgSlice []model.VetOrg

	result := vetOrgRepo.Database.Where("org_name LIKE ?", "%"+vetOrgName+"%").Find(&vetOrgSlice)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("vetOrg(s) with orgName like '%s' not found", vetOrgName)
	}
	return vetOrgSlice, nil
}
