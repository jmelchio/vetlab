package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/jmelchio/vetlab/model"
)

// User implements the api.UserService interface
type User struct {
	UserRepo UserRepo
}

// CreateUser creates a new model.User in the vetlab system
// The assumption is that for a new user, the password has not been encrypted
// up until this point so this is done before storing it in the repository.
func (userService User) CreateUser(ctx context.Context, user model.User) (*model.User, error) {
	if ctx == nil {
		return nil, errors.New(MissingContext)
	}

	pwdHash, err := hashAndSalt(user.Password)
	if err != nil {
		return nil, err
	}

	user.Password = *pwdHash
	err = userService.UserRepo.Create(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// UpdateUser updates a model.User in the vetlab system
func (userService User) UpdateUser(ctx context.Context, user model.User) (*model.User, error) {
	if ctx == nil {
		return nil, errors.New(MissingContext)
	}

	err := userService.UserRepo.Update(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// DeleteUser delets a model.User from the vetlab system
func (userService User) DeleteUser(ctx context.Context, user model.User) error {
	if ctx == nil {
		return errors.New(MissingContext)
	}
	err := userService.UserRepo.Delete(&user)
	return err
}

// UpdatePassword allows for the change of a user password
// We do not check the old password when changing to a new one since we assume the user
// has been authenticated
func (userService User) UpdatePassword(ctx context.Context, user model.User, password string) (*model.User, error) {
	if ctx == nil {
		return nil, errors.New(MissingContext)
	}

	pwdHash, err := hashAndSalt(password)
	if err != nil {
		return nil, fmt.Errorf(HashingFailed, err.Error())
	}

	user.Password = *pwdHash
	err = userService.UserRepo.Update(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// Login tries to login a user into the vetlab system
func (userService User) Login(ctx context.Context, userName string, password string) (*model.User, error) {
	if ctx == nil {
		return nil, errors.New(MissingContext)
	}

	if userName == "" || password == "" {
		return nil, errors.New(UserOrPasswordFail)
	}

	user, err := userService.UserRepo.GetByUserName(userName)
	if err != nil {
		return nil, errors.New(UserOrPasswordFail)
	}

	if !equalPasswords(user.Password, password) {
		return nil, errors.New(UserOrPasswordFail)
	}

	return user, nil
}

// FindUserByUserName attempts to find users by their userName
func (userService User) FindUserByUserName(ctx context.Context, userName string) (*model.User, error) {
	if ctx == nil {
		return nil, errors.New(MissingContext)
	}
	return userService.UserRepo.GetByUserName(userName)
}

// FindUserByID finds users by their unique ID
func (userService User) FindUserByID(ctx context.Context, userID uint) (*model.User, error) {
	if ctx == nil {
		return nil, errors.New(MissingContext)
	}
	return userService.UserRepo.GetByID(userID)
}
