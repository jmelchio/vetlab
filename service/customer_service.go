package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/jmelchio/vetlab/model"
)

// Customer implements the api.CustomerService interface
type Customer struct {
	CustomerRepo CustomerRepo
}

// CreateCustomer creates a new model.Customer in the vetlab system
// The assumption is that for a new customer, the password has not been encrypted
// up until this point so this is done before storing it in the repository.
func (customerService Customer) CreateCustomer(ctx context.Context, customer model.Customer) (*model.Customer, error) {
	if ctx == nil {
		return nil, errors.New(MissingContext)
	}

	pwdHash, err := hashAndSalt(customer.Password)
	if err != nil {
		return nil, fmt.Errorf(HashingFailed, err.Error())
	}

	customer.Password = *pwdHash
	err = customerService.CustomerRepo.Create(&customer)
	if err != nil {
		return nil, err
	}
	return &customer, nil
}

// UpdateCustomer updates a model.Customer in the vetlab system
func (customerService Customer) UpdateCustomer(ctx context.Context, customer model.Customer) (*model.Customer, error) {
	if ctx == nil {
		return nil, errors.New(MissingContext)
	}

	err := customerService.CustomerRepo.Update(&customer)
	if err != nil {
		return nil, err
	}
	return &customer, nil
}

// DeleteCustomer deletes a model.Customer from the vetlab system
func (customerService Customer) DeleteCustomer(ctx context.Context, customer model.Customer) error {
	if ctx == nil {
		return errors.New(MissingContext)
	}
	return customerService.CustomerRepo.Delete(&customer)
}

// UpdatePassword allows for the change of a customer password
// We do not check the old password when changing to a new one since we assume the customer
// has been authenticated
func (customerService Customer) UpdatePassword(ctx context.Context, customer model.Customer, password string) (*model.Customer, error) {
	if ctx == nil {
		return nil, errors.New(MissingContext)
	}

	pwdHash, err := hashAndSalt(password)
	if err != nil {
		return nil, fmt.Errorf(HashingFailed, err.Error())
	}

	customer.Password = *pwdHash
	err = customerService.CustomerRepo.Update(&customer)
	if err != nil {
		return nil, err
	}

	return &customer, nil
}

// Login tries to log in a customer into the vetlab system
func (customerService Customer) Login(ctx context.Context, userName string, password string) (*model.Customer, error) {
	if ctx == nil {
		return nil, errors.New(MissingContext)
	}

	if userName == "" || password == "" {
		return nil, errors.New(UserOrPasswordFail)
	}

	customer, err := customerService.CustomerRepo.GetByUserName(userName)
	if err != nil {
		return nil, errors.New(UserOrPasswordFail)
	}

	if !equalPasswords(customer.Password, password) {
		return nil, errors.New(UserOrPasswordFail)
	}

	return customer, nil
}

// FindCustomerByUserName attempts to find customers by their userName
func (customerService Customer) FindCustomerByUserName(ctx context.Context, userName string) (*model.Customer, error) {
	if ctx == nil {
		return nil, errors.New(MissingContext)
	}
	return customerService.CustomerRepo.GetByUserName(userName)
}

// FindCustomerByID finds customers by their unique ID
func (customerService Customer) FindCustomerByID(ctx context.Context, customerID uint) (*model.Customer, error) {
	if ctx == nil {
		return nil, errors.New(MissingContext)
	}
	return customerService.CustomerRepo.GetByID(customerID)
}

// FindCustomerByVetOrg attempts to find customers by the VetOrg they belong to
func (customerService Customer) FindCustomerByVetOrg(ctx context.Context, vetOrg model.VetOrg) ([]model.Customer, error) {
	if ctx == nil {
		return nil, errors.New(MissingContext)
	}
	if vetOrg.ID == uint(0) {
		return nil, errors.New(VetOrgRequired)
	}
	return customerService.CustomerRepo.GetByVetOrgID(vetOrg.ID)
}
