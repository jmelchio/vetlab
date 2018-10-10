package repository

import (
	"errors"

	"github.com/jinzhu/gorm"
	// import the proper dialect for Gorm
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/jmelchio/vetlab/model"
)

// SQLCustomerRepo describes the sql database that persists the Customer
type SQLCustomerRepo struct {
	Database *gorm.DB
}

// Create creates a persistent Customer row in the sql datastore
func (sqlCustomerRepo SQLCustomerRepo) Create(customer *model.Customer) error {
	if sqlCustomerRepo.Database.NewRecord(customer) {
		if err := sqlCustomerRepo.Database.Create(customer).Error; err != nil {
			return err
		}
		return nil
	}
	return errors.New("Record already in database")
}

// Update modifies a Customer row in the sql datastore
// If the password is less than 50 characters long it's probably not hashed and
// should therefore not be saved to the database
func (sqlCustomerRepo SQLCustomerRepo) Update(customer *model.Customer) error {
	if !sqlCustomerRepo.Database.NewRecord(customer) {
		if len(customer.Password) < 50 {
			if err := sqlCustomerRepo.Database.Model(customer).Updates(
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
			if err := sqlCustomerRepo.Database.Save(customer).Error; err != nil {
				return err
			}
		}
		return nil
	}
	return errors.New("Record does not exist in database")
}

// Delete removes a Customer row in the sql datastore
func (sqlCustomerRepo SQLCustomerRepo) Delete(customer *model.Customer) error {
	if err := sqlCustomerRepo.Database.Delete(customer).Error; err != nil {
		return err
	}
	return nil
}

// GetByID fetches a Customer from the sql datastore
func (sqlCustomerRepo SQLCustomerRepo) GetByID(customerID uint) (*model.Customer, error) {
	var customer model.Customer

	if err := sqlCustomerRepo.Database.First(&customer, customerID).Error; err != nil {
		return nil, err
	}
	return &customer, nil
}

// GetByVetOrgID fetches all customers by VetOrg from the sql datastore
func (sqlCustomerRepo SQLCustomerRepo) GetByVetOrgID(vetOrgID uint) ([]model.Customer, error) {
	return nil, nil
}

// GetByUserName fetches all customers by UserName from the sql datastore
func (sqlCustomerRepo SQLCustomerRepo) GetByUserName(userName string) (*model.Customer, error) {
	var customer model.Customer

	if err := sqlCustomerRepo.Database.Where("user_name = ?", userName).Find(&customer).Error; err != nil {
		return nil, err
	}
	return &customer, nil
}
