package service

import (
	"context"
	"errors"

	"github.com/jmelchio/vetlab/model"
)

// User implements the api.UserService interface
type User struct {
	UserRepo UserRepo
}

const (
	MissingContext = "Context is required"
)

// CreateUser creates a new model.User in the vetlab system
func (userService User) CreateUser(ctx context.Context, user model.User) (*model.User, error) {
	if ctx == nil {
		return nil, errors.New(MissingContext)
	}

	newUser, err := userService.UserRepo.Create(user)
	if err != nil {
		return nil, err
	}
	return newUser, nil
}

// UpdateUser updates a model.User in the vetlab system
func (userService User) UpdateUser(ctx context.Context, user model.User) (*model.User, error) {
	if ctx == nil {
		return nil, errors.New(MissingContext)
	}

	newUser, err := userService.UserRepo.Update(user)
	if err != nil {
		return nil, err
	}
	return newUser, nil
}

// DeleteUser delets a model.User from the vetlab system
func (userService User) DeleteUser(ctx context.Context, user model.User) error {
	if ctx == nil {
		return errors.New(MissingContext)
	}
	err := userService.UserRepo.Delete(user)
	return err
}

// Login tries to login a user into the vetlab system
func (userService User) Login(ctx context.Context, userName string, password string) (*model.User, error) {
	return nil, nil
}

// FindUsersByVetOrg attempts to find users by the veterinary organization
func (userService User) FindUsersByVetOrg(ctx context.Context, vetOrg model.VetOrg) ([]model.User, error) {
	return nil, nil
}

// FindUsersByName attempts to find users by their name
func (userService User) FindUsersByName(ctx context.Context, userName string) ([]model.User, error) {
	return nil, nil
}

// FindUserByID finds users by their unique ID
func (userService User) FindUserByID(ctx context.Context, userID string) (*model.User, error) {
	return nil, nil
}
