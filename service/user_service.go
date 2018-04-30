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

// CreateUser creates a new model.User in the vetlab system
func (userService User) CreateUser(ctx context.Context, user model.User) (*model.User, error) {
	if ctx == nil {
		return nil, errors.New("Context is required")
	}

	newUser, err := userService.UserRepo.Create(user)
	if err != nil {
		return nil, err
	}
	return newUser, nil
}

// UpdateUser updates a model.User in the vetlab system
func (userService User) UpdateUser(context.Context, model.User) (*model.User, error) {
	return nil, nil
}

// DeleteUser delets a model.User from the vetlab system
func (userService User) DeleteUser(context.Context, model.User) error {
	return nil
}

// Login tries to login a user into the vetlab system
func (userService User) Login(context.Context, string, string) (*model.User, error) {
	return nil, nil
}

// FindUsersByVetOrg attempts to find users by the veterinary organization
func (userService User) FindUsersByVetOrg(context.Context, model.VetOrg) ([]model.User, error) {
	return nil, nil
}

// FindUsersByName attempts to find users by their name
func (userService User) FindUsersByName(context.Context, string) ([]model.User, error) {
	return nil, nil
}

// FindUserByID finds users by their unique ID
func (userService User) FindUserByID(context.Context, string) (*model.User, error) {
	return nil, nil
}
