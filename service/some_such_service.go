package service

import (
	"errors"

	"github.com/jmelchio/vetlab/model"
	"github.com/jmelchio/vetlab/repository"
)

// SomeSuchService is just such a service
type SomeSuchService struct {
	userRepo *repository.SQLUserRepo
}

// Login will log you in to the system or not (maybe?)
func (someService SomeSuchService) Login(userID uint, password string) (*model.User, error) {
	user, err := someService.userRepo.GetByID(userID)
	if err != nil {
		return nil, err
	}

	if user.PasswordHash == password {
		return user, nil
	}
	return nil, errors.New("Unable to log into the system")
}
