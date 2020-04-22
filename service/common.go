package service

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

const (
	MissingContext     = "context is required"
	PasswordTooShort   = "password should be at least 8 characters"
	HashingFailed      = "failed to salt and hash password: %s"
	UserOrPasswordFail = "user or password mismatch"
	VetOrgRequired     = "vetOrg ID is required"
)

func hashAndSalt(pwd string) (*string, error) {
	if len(pwd) < 8 {
		return nil, errors.New(PasswordTooShort)
	}
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
