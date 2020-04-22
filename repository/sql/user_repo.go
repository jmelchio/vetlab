package sql

import (
	"errors"

	"github.com/jinzhu/gorm"
	// import the proper dialect for Gorm
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/jmelchio/vetlab/model"
)

// UserRepo describes the sql database that persists the User
type UserRepo struct {
	Database *gorm.DB
}

// Create creates a persistent User row in the sql datastore
func (userRepo *UserRepo) Create(user *model.User) error {
	if userRepo.Database.NewRecord(user) {
		if err := userRepo.Database.Create(user).Error; err != nil {
			return err
		}
		return nil
	}
	return errors.New("record already in database")
}

// Update modifies a User row in the sql datastore
// If the password is less than 50 characters long it's probably not hashed and
// should therefore not be saved to the database (yes, it's janky)
func (userRepo *UserRepo) Update(user *model.User) error {
	if !userRepo.Database.NewRecord(user) {
		if len(user.Password) < 50 {
			if err := userRepo.Database.Model(user).Updates(
				model.User{
					UserName:  user.UserName,
					FirstName: user.FirstName,
					LastName:  user.LastName,
					Email:     user.Email,
					AdminUser: user.AdminUser,
				}).Error; err != nil {
				return err
			}
		} else {
			if err := userRepo.Database.Save(user).Error; err != nil {
				return err
			}
		}
		return nil
	}
	return errors.New("record does not exist in database")
}

// Delete removes a User row in the sql datastore
func (userRepo *UserRepo) Delete(user *model.User) error {
	if err := userRepo.Database.Delete(user).Error; err != nil {
		return err
	}
	return nil
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
func (userRepo *UserRepo) GetByUserName(userName string) (*model.User, error) {
	var user model.User

	if err := userRepo.Database.Where("user_name = ?", userName).Find(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
