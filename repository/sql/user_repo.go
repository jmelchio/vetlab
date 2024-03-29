package sql

import (
	"errors"
	"fmt"

	"gorm.io/gorm"

	"github.com/jmelchio/vetlab/model"
)

// UserRepo describes the sql database that persists the User
type UserRepo struct {
	Database *gorm.DB
}

// Create creates a persistent User row in the sql datastore
func (userRepo *UserRepo) Create(user *model.User) error {
	if user.ID == 0 {
		return userRepo.Database.Create(user).Error
	}
	return errors.New("record already in database")
}

// Update modifies a User row in the sql datastore
// If the password is less than 50 characters long it's probably not hashed and
// should therefore not be saved to the database (yes, it's janky)
func (userRepo *UserRepo) Update(user *model.User) error {
	if user.ID != 0 {
		if len(user.Password) < 50 {
			return userRepo.Database.Model(user).Updates(
				model.User{
					UserName:  user.UserName,
					FirstName: user.FirstName,
					LastName:  user.LastName,
					Email:     user.Email,
					AdminUser: user.AdminUser,
				}).Error
		} else {
			return userRepo.Database.Save(user).Error
		}
	}
	return errors.New("record does not exist in database")
}

// Delete removes a User row in the sql datastore
func (userRepo *UserRepo) Delete(user *model.User) error {
	return userRepo.Database.Delete(user).Error
}

// GetByID fetches a User from the sql datastore
func (userRepo *UserRepo) GetByID(userID uint) (*model.User, error) {
	var user model.User

	if err := userRepo.Database.First(&user, userID).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// GetByUserName fetches all users by UserName from the sql datastore
// The username search string should be at least 3 characters long
func (userRepo *UserRepo) GetByUserName(userName string) (*model.User, error) {
	var user model.User

	result := userRepo.Database.Where("user_name LIKE ?", "%"+userName+"%").Find(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("user(s) with username like '%s' not found", userName)
	}
	return &user, nil
}
