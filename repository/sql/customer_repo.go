package sql

import (
	"errors"
	"fmt"

	"gorm.io/gorm"

	"github.com/jmelchio/vetlab/model"
)

// CustomerRepo describes the sql database that persists the Customer
type CustomerRepo struct {
	Database *gorm.DB
}

// Create creates a persistent Customer row in the sql datastore
func (customerRepo *CustomerRepo) Create(customer *model.Customer) error {
	if customer.ID == 0 {
		if err := customerRepo.Database.Create(customer).Error; err != nil {
			return err
		}
		return nil
	}
	return errors.New("record already in database")
}

// Update modifies a Customer row in the sql datastore
// If the password is less than 50 characters long it's probably not hashed and
// should therefore not be saved to the database
func (customerRepo *CustomerRepo) Update(customer *model.Customer) error {
	if customer.ID != 0 {
		if len(customer.Password) < 50 {
			if err := customerRepo.Database.Model(customer).Updates(
				model.Customer{
					UserName:  customer.UserName,
					FirstName: customer.FirstName,
					LastName:  customer.LastName,
					Email:     customer.Email,
					VetOrgID:  customer.VetOrgID,
				}).Error; err != nil {
				return err
			}
		} else {
			if err := customerRepo.Database.Save(customer).Error; err != nil {
				return err
			}
		}
		return nil
	}
	return errors.New("record does not exist in database")
}

// Delete removes a Customer row in the sql datastore
func (customerRepo *CustomerRepo) Delete(customer *model.Customer) error {
	return customerRepo.Database.Delete(customer).Error
}

// GetByID fetches a Customer from the sql datastore
func (customerRepo *CustomerRepo) GetByID(customerID uint) (*model.Customer, error) {
	var customer model.Customer

	if err := customerRepo.Database.First(&customer, customerID).Error; err != nil {
		return nil, err
	}
	return &customer, nil
}

// GetByVetOrgID fetches all customers by VetOrg from the sql datastore
func (customerRepo *CustomerRepo) GetByVetOrgID(vetOrgID uint) ([]model.Customer, error) {
	var customers []model.Customer

	result := customerRepo.Database.Where("vet_org_id = ?", vetOrgID).Find(&customers)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("customers with vetOrgId '%d' not found", vetOrgID)
	}
	return customers, nil
}

// GetByUserName fetches all customers by UserName from the sql datastore
func (customerRepo *CustomerRepo) GetByUserName(userName string) (*model.Customer, error) {
	var customer model.Customer

	result := customerRepo.Database.Where("user_name = ?", userName).Find(&customer)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("customer with username '%s' not found", userName)
	}
	return &customer, nil
}
