package repository

import (
	"errors"

	"github.com/jinzhu/gorm"
	// import the proper dialect for Gorm
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/jmelchio/vetlab/model"
)

// SQLUserRepo describes the sql database that persists the User
type SQLUserRepo struct {
	Database *gorm.DB
}

// Create creates a persistent User row in the sql datastore
func (sqlUserRepo SQLUserRepo) Create(user *model.User) error {
	if sqlUserRepo.Database.NewRecord(user) {
		if err := sqlUserRepo.Database.Create(user).Error; err != nil {
			return err
		}
		return nil
	}
	return errors.New("Record already in database")
}

// Update modifies a User row in the sql datastore
// If the password is less than 50 characters long it's probably not hashed and
// should therefore not be saved to the database
func (sqlUserRepo SQLUserRepo) Update(user *model.User) error {
	if !sqlUserRepo.Database.NewRecord(user) {
		if len(*user.PasswordHash) < 25 {
			if err := sqlUserRepo.Database.Model(user).Updates(
				model.User{
					UserName:  user.UserName,
					FirstName: user.FirstName,
					LastName:  user.LastName,
					Email:     user.Email,
					OrgID:     user.OrgID,
					AdminUser: user.AdminUser,
				}).Error; err != nil {
				return err
			}
		} else {
			if err := sqlUserRepo.Database.Save(user).Error; err != nil {
				return err
			}
		}
		return nil
	}
	return errors.New("Record does not exist in database")
}

// Delete removes a User row in the sql datastore
func (sqlUserRepo SQLUserRepo) Delete(user *model.User) error {
	if err := sqlUserRepo.Database.Delete(user).Error; err != nil {
		return err
	}
	return nil
}

// GetByID fetches a User from the sql datastore
func (sqlUserRepo SQLUserRepo) GetByID(userID uint) (*model.User, error) {
	var user model.User

	if err := sqlUserRepo.Database.First(&user, userID).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// GetByOrgID fetches all users for a vet org from the sql datastore
func (sqlUserRepo SQLUserRepo) GetByOrgID(orgID uint) ([]model.User, error) {
	var users []model.User

	if err := sqlUserRepo.Database.Where("org_id = ?", orgID).Find(&users).Error; err != nil {
		return nil, err
	}
	if len(users) > 0 {
		return users, nil
	}
	return nil, errors.New("Zero records found that match given org_id")
}

// GetByUserName fetches all users by UserName from the sql datastore
func (sqlUserRepo SQLUserRepo) GetByUserName(userName string) (*model.User, error) {
	var user model.User

	if err := sqlUserRepo.Database.Where("user_name = ?", userName).Find(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
