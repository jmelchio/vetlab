package repository_test

import (
	"github.com/jmelchio/vetlab/model"
	. "github.com/jmelchio/vetlab/repository"
	"github.com/jmelchio/vetlab/service"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("SqlCustomerRepo", func() {

	var (
		customerRepo service.CustomerRepo
		customerOne  model.Customer
		customerTwo  model.Customer
		userName     string
		firstName    string
		lastName     string
		email        string
		password     string
		vetOrgID     uint
	)

	BeforeEach(func() {
		customerRepoImpl := SQLCustomerRepo{Database: database}
		customerRepo = customerRepoImpl

		userName = "user_name"
		firstName = "first_name"
		lastName = "last_name"
		email = "first.last@gmail.com"
		password = "want_some_hash?"
		vetOrgID = uint(12345)

		customerOne = model.Customer{
			UserName:  &userName,
			FirstName: firstName,
			LastName:  lastName,
			Email:     email,
			Password:  password,
			VetOrgID:  vetOrgID,
		}
	})

	AfterEach(func() {
		err = database.Where("1 = 1").Delete(&model.Customer{}).Error
		Expect(err).NotTo(HaveOccurred())
	})

	Describe("Customer table", func() {

		Context("Customer table has been created during in BeforeSuite", func() {

			It("Has a customer table", func() {
				hasCustomerTable := database.HasTable(&model.Customer{})
				Expect(hasCustomerTable).To(BeTrue())
			})
		})
	})

	Describe("Create a customer", func() {

		Context("When a username is not taken yet", func() {
			BeforeEach(func() {
				Expect(customerOne.ID).To(Equal(uint(0)))
			})

			It("Creates a new customer record", func() {
				err = customerRepo.Create(&customerOne)
				Expect(err).NotTo(HaveOccurred())
				Expect(customerOne.ID).NotTo(Equal(uint(0)))
			})
		})

		Context("When a username is taken already", func() {

			BeforeEach(func() {
				customerTwo = customerOne
			})

			It("It returns an error", func() {
				err = customerRepo.Create(&customerOne)
				Expect(err).NotTo(HaveOccurred())
				err = customerRepo.Create(&customerTwo)
				Expect(err).To(HaveOccurred())
			})
		})
	})

	Describe("Update a user", func() {

		Context("When a user is found", func() {

			BeforeEach(func() {
				err = customerRepo.Create(&customerOne)
				Expect(err).NotTo(HaveOccurred())
				Expect(customerOne.ID).NotTo(Equal(uint(0)))
			})

			Context("When the updated password is less than 50 characters", func() {

				It("It updates the customer record and returns updated customer with unchanged password", func() {
					*customerOne.UserName = "new_user_name"
					customerOne.Password = "short_password"
					err = customerRepo.Update(&customerOne)
					Expect(err).NotTo(HaveOccurred())
					customerFound, ferr := customerRepo.GetByID(customerOne.ID)
					Expect(ferr).NotTo(HaveOccurred())
					Expect(customerFound.Password).To(Equal("want_some_hash?"))
					Expect(customerFound.UserName).To(Equal(customerOne.UserName))
				})
			})

			Context("When the updated password is more than 50 characters", func() {

				It("It updates the customer record and returns updated customer with unchanged password", func() {
					*customerOne.UserName = "new_user_name"
					customerOne.Password = "long_password_of_more_than_fifty_characters_so_that_its"
					err = customerRepo.Update(&customerOne)
					Expect(err).NotTo(HaveOccurred())
					customerFound, ferr := customerRepo.GetByID(customerOne.ID)
					Expect(ferr).NotTo(HaveOccurred())
					Expect(customerFound.Password).To(Equal(customerOne.Password))
					Expect(customerFound.UserName).To(Equal(customerOne.UserName))
				})
			})
		})

		Context("When the customer does not exist", func() {

			It("Returns an error and nil for the customer", func() {
				err = customerRepo.Update(&customerOne)
				Expect(err).To(HaveOccurred())
			})
		})
	})

	Describe("Delete a customer", func() {

		Context("When the use is found", func() {

			BeforeEach(func() {
				err = customerRepo.Create(&customerOne)
				Expect(err).NotTo(HaveOccurred())
				Expect(customerOne.ID).NotTo(Equal(uint(0)))
			})

			It("Deletes the record and returns no error", func() {
				err = customerRepo.Delete(&customerOne)
				Expect(err).NotTo(HaveOccurred())
				var foundCustomer *model.Customer
				foundCustomer, err = customerRepo.GetByID(customerOne.ID)
				Expect(err).To(HaveOccurred())
				Expect(foundCustomer).To(BeNil())
			})
		})
	})

	Describe("Get a customer by ID", func() {

		Context("When the customer is found", func() {

			var foundCustomer *model.Customer

			BeforeEach(func() {
				err = customerRepo.Create(&customerOne)
				Expect(err).NotTo(HaveOccurred())
				Expect(customerOne.ID).NotTo(Equal(uint(0)))
			})

			It("It returns the customer and nil for error", func() {
				foundCustomer, err = customerRepo.GetByID(customerOne.ID)
				Expect(err).NotTo(HaveOccurred())
				Expect(foundCustomer).NotTo(BeNil())
				Expect(foundCustomer.UserName).To(Equal(customerOne.UserName))
			})
		})

		Context("When the customer is not found", func() {

			var foundCustomer *model.Customer

			It("It returns and error and nil for the customer", func() {
				foundCustomer, err = customerRepo.GetByID(uint(10))
				Expect(err).To(HaveOccurred())
				Expect(foundCustomer).To(BeNil())
			})
		})
	})

	Describe("Get a customer by UserName", func() {

		Context("When the customer is found", func() {

			var foundCustomer *model.Customer

			BeforeEach(func() {
				err = customerRepo.Create(&customerOne)
				Expect(err).NotTo(HaveOccurred())
				Expect(customerOne.ID).NotTo(Equal(uint(0)))
			})

			It("It returns the customer and nil for error", func() {
				foundCustomer, err = customerRepo.GetByUserName(*customerOne.UserName)
				Expect(err).NotTo(HaveOccurred())
				Expect(foundCustomer).NotTo(BeNil())
				Expect(foundCustomer.ID).To(Equal(customerOne.ID))
				Expect(foundCustomer.UserName).To(Equal(customerOne.UserName))
			})
		})

		Context("When the customer is not found", func() {

			var foundCustomer *model.Customer

			BeforeEach(func() {
			})

			It("It returns the customer and nil for error", func() {
				foundCustomer, err = customerRepo.GetByUserName("some_user_name")
				Expect(err).To(HaveOccurred())
				Expect(foundCustomer).To(BeNil())
			})
		})
	})
})
