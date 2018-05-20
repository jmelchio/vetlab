package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/jmelchio/vetlab/model"
	"golang.org/x/crypto/bcrypt"
)

// User implements the api.UserService interface
type User struct {
	UserRepo UserRepo
}

const (
	MissingContext     = "Context is required"
	PasswordTooShort   = "Password should be at least 8 characters"
	HashingFailed      = "Failed to salt and hash password: %s"
	UserOrPasswordFail = "User or Password mismatch"
	VetOrgRequired     = "VetOrg ID is required"
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

// UpdatePassword allows for the change of a user password
// We do not check the old password when changing to a new one since we assume the user
// has been authenticated
func (userService User) UpdatePassword(ctx context.Context, user model.User, password string) (*model.User, error) {
	if ctx == nil {
		return nil, errors.New(MissingContext)
	}
	if len(password) < 8 {
		return nil, errors.New(PasswordTooShort)
	}
	pwdHash, err := hashAndSalt(password)
	if err != nil {
		return nil, fmt.Errorf(HashingFailed, err.Error())
	}
	user.PasswordHash = *pwdHash
	_, uerr := userService.UserRepo.Update(user)
	if uerr != nil {
		return nil, uerr
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

	if !equalPasswords(user.PasswordHash, password) {
		return nil, errors.New(UserOrPasswordFail)
	}

	return user, nil
}

// FindUsersByVetOrg attempts to find users by the veterinary organization
func (userService User) FindUsersByVetOrg(ctx context.Context, vetOrg model.VetOrg) ([]model.User, error) {
	if ctx == nil {
		return nil, errors.New(MissingContext)
	}
	if vetOrg.OrgID == "" {
		return nil, errors.New(VetOrgRequired)
	}
	return userService.UserRepo.GetByOrgID(vetOrg.OrgID)
}

// FindUserByUserName attempts to find users by their userName
func (userService User) FindUserByUserName(ctx context.Context, userName string) (*model.User, error) {
	if ctx == nil {
		return nil, errors.New(MissingContext)
	}
	return userService.UserRepo.GetByUserName(userName)
}

// FindUserByID finds users by their unique ID
func (userService User) FindUserByID(ctx context.Context, userID string) (*model.User, error) {
	if ctx == nil {
		return nil, errors.New(MissingContext)
	}
	return userService.UserRepo.GetByID(userID)
}

func hashAndSalt(pwd string) (*string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	pwdHash := string(hash)
	return &pwdHash, nil
}

func equalPasswords(pwdHash string, pwdPlain string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(pwdHash), []byte(pwdPlain))
	if err != nil {
		return false
	}
	return true
}
