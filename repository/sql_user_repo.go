package repository

import (
	"errors"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/jmelchio/vetlab/model"
)

// SQLUserRepo describes the sql database that persists the User
type SQLUserRepo struct {
	Database *gorm.DB
}

// Create creates a persistent User row in the sql datastore
func (sqlUserRepo SQLUserRepo) Create(user model.User) (*model.User, error) {
	if sqlUserRepo.Database.NewRecord(user) {
		if err := sqlUserRepo.Database.Create(&user).Error; err != nil {
			return nil, err
		} else {
			return &user, nil
		}
	}
	return nil, errors.New("Record already in database")
}

// Update modifies a User row in the sql datastore
func (sqlUserRepo SQLUserRepo) Update(user model.User) (*model.User, error) {
	if err := sqlUserRepo.Database.Save(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// Delete removes a User row in the sql datastore
func (sqlUserRepo SQLUserRepo) Delete(user model.User) error {
	return nil
}

// GetByID fetches a User from the sql datastore
func (sqlUserRepo SQLUserRepo) GetByID(userID string) (*model.User, error) {
	return nil, nil
}

// GetByOrgID fetches all users for a vet org from the sql datastore
func (sqlUserRepo SQLUserRepo) GetByOrgID(orgID string) ([]model.User, error) {
	return nil, nil
}
