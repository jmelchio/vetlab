package service_test

import (
	"context"
	"errors"
	"fmt"

	"github.com/jmelchio/vetlab/api"
	"github.com/jmelchio/vetlab/model"
	. "github.com/jmelchio/vetlab/service"
	"github.com/jmelchio/vetlab/service/servicefakes"
	"golang.org/x/crypto/bcrypt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("CustomerService", func() {

	var (
		customerService api.CustomerService
		customerRepo    *servicefakes.FakeCustomerRepo
		customer        model.Customer
		userName        string
		firstName       string
		lastName        string
		email           string
		password        string
	)

	BeforeEach(func() {
		customerRepo = new(servicefakes.FakeCustomerRepo)
		customerServiceImpl := Customer{CustomerRepo: customerRepo}
		customerService = customerServiceImpl

		userName = "some-name"
		firstName = "first-name"
		lastName = "last-name"
		email = "email@domain.com"
		password = "password-hash"

		customer = model.Customer{
			UserName:  &userName,
			FirstName: firstName,
			LastName:  lastName,
			Email:     email,
			Password:  password,
		}
	})

	Describe("Create a customer", func() {

		BeforeEach(func() {
			customerRepo.CreateReturns(nil)
		})

		Context("We have a valid customer and 'todo' context", func() {

			It("Returns a customer with a new customer ID and calls CustomerRepo.Create", func() {
				zeCustomer, err := customerService.CreateCustomer(context.TODO(), customer)
				Expect(err).ToNot(HaveOccurred())
				Expect(zeCustomer).NotTo(BeNil())
				Expect(customerRepo.CreateCallCount()).To(Equal(1))
			})
		})

		Context("Customer password is too short but we have a valid context", func() {

			BeforeEach(func() {
				customer.Password = "uhSeven"
			})

			It("Returns an error and does not call CustomerRepo.Create", func() {
				zeCustomer, err := customerService.CreateCustomer(context.TODO(), customer)
				Expect(err).To(HaveOccurred())
				Expect(zeCustomer).To(BeNil())
				Expect(err.Error()).To(Equal(PasswordTooShort))
				Expect(customerRepo.CreateCallCount()).To(Equal(0))
			})
		})

		Context("We have a valid customer but no context", func() {

			It("Returns an error and does not call CustomerRepo.Create", func() {
				zeCustomer, err := customerService.CreateCustomer(nil, customer)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal(MissingContext))
				Expect(zeCustomer).To(BeNil())
				Expect(customerRepo.CreateCallCount()).To(Equal(0))
			})
		})

		Context("We have a customer and Context but repo cannot create customer", func() {

			BeforeEach(func() {
				customerRepo.CreateReturns(errors.New("Unable to create customer"))
			})

			It("Returns an error after calling CustomerRepo.Create", func() {
				zeCustomer, err := customerService.CreateCustomer(context.TODO(), customer)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("Unable to create customer"))
				Expect(zeCustomer).To(BeNil())
				Expect(customerRepo.CreateCallCount()).To(Equal(1))
			})
		})
	})

	Describe("Update a customer", func() {

		BeforeEach(func() {
			customerRepo.UpdateReturns(nil)
		})

		Context("We have a valid customer and context", func() {

			It("Returns the updated customer and no error", func() {
				zeCustomer, err := customerService.UpdateCustomer(context.TODO(), customer)
				Expect(err).NotTo(HaveOccurred())
				Expect(customerRepo.UpdateCallCount()).To(Equal(1))
				Expect(zeCustomer).NotTo(BeNil())
			})
		})

		Context("We have a valid customer but no context", func() {

			It("Returns and error and no updated customer", func() {
				zeCustomer, err := customerService.UpdateCustomer(nil, customer)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal(MissingContext))
				Expect(zeCustomer).To(BeNil())
				Expect(customerRepo.UpdateCallCount()).To(Equal(0))
			})
		})

		Context("We have a customer and Context but repo cannot create customer", func() {

			BeforeEach(func() {
				customerRepo.UpdateReturns(errors.New("Unable to update customer"))
			})

			It("Returns an error after calling CustomerRepo.Create", func() {
				zeCustomer, err := customerService.UpdateCustomer(context.TODO(), customer)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("Unable to update customer"))
				Expect(zeCustomer).To(BeNil())
				Expect(customerRepo.UpdateCallCount()).To(Equal(1))
			})
		})
	})

	Describe("Delete a customer", func() {

		BeforeEach(func() {
			customerRepo.DeleteReturns(nil)
		})

		Context("We have a valid customer and context", func() {

			It("Returns no error", func() {
				err := customerService.DeleteCustomer(context.TODO(), customer)
				Expect(err).NotTo(HaveOccurred())
				Expect(customerRepo.DeleteCallCount()).To(Equal(1))
			})
		})

		Context("We have a valid customer but no context", func() {

			It("Returns and error and no updated customer", func() {
				err := customerService.DeleteCustomer(nil, customer)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal(MissingContext))
				Expect(customerRepo.DeleteCallCount()).To(Equal(0))
			})
		})

		Context("We have a customer and Context but repo cannot delete customer", func() {

			BeforeEach(func() {
				customerRepo.DeleteReturns(errors.New("Unable to delete customer"))
			})

			It("Returns an error after calling CustomerRepo.Create", func() {
				err := customerService.DeleteCustomer(context.TODO(), customer)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("Unable to delete customer"))
				Expect(customerRepo.DeleteCallCount()).To(Equal(1))
			})
		})
	})

	Describe("UpdatePassword", func() {

		var (
			newPwd string
		)

		BeforeEach(func() {
			newPwd = "some-magic-password"
			customerRepo.UpdateReturns(nil)
		})

		Context("We have a valid customer, password and context", func() {

			It("It returns no error", func() {
				updatedCustomer, err := customerService.UpdatePassword(context.TODO(), customer, newPwd)
				Expect(err).ToNot(HaveOccurred())
				Expect(customerRepo.UpdateCallCount()).To(Equal(1))
				err = bcrypt.CompareHashAndPassword([]byte(updatedCustomer.Password), []byte(newPwd))
				Expect(err).ToNot(HaveOccurred())
			})
		})

		Context("We have a valid customer, password but no context", func() {

			It("Returns an error", func() {
				updatedCustomer, err := customerService.UpdatePassword(nil, customer, newPwd)
				Expect(err).To(HaveOccurred())
				Expect(updatedCustomer).To(BeNil())
				Expect(err.Error()).To(Equal(MissingContext))
				Expect(customerRepo.UpdateCallCount()).To(Equal(0))
			})
		})

		Context("We have a valid customer, password and context but update fails", func() {

			BeforeEach(func() {
				customerRepo.UpdateReturns(errors.New("Unable to update customer"))
			})

			It("returns an error", func() {
				updatedCustomer, err := customerService.UpdatePassword(context.TODO(), customer, newPwd)
				Expect(err).To(HaveOccurred())
				Expect(updatedCustomer).To(BeNil())
				Expect(err.Error()).To(Equal("Unable to update customer"))
				Expect(customerRepo.UpdateCallCount()).To(Equal(1))
			})
		})

		Context("We have a valid customer, a short password and context", func() {

			BeforeEach(func() {
				newPwd = "uhseven"
				customerRepo.UpdateReturns(errors.New("Unable to update customer"))
			})

			It("returns an error", func() {
				updatedCustomer, err := customerService.UpdatePassword(context.TODO(), customer, newPwd)
				Expect(err).To(HaveOccurred())
				Expect(updatedCustomer).To(BeNil())
				Expect(err.Error()).To(Equal(fmt.Sprintf(HashingFailed, PasswordTooShort)))
				Expect(customerRepo.UpdateCallCount()).To(Equal(0))
			})
		})
	})

	Describe("Login", func() {

		var (
			luserName      string
			password       string
			sampleCustomer model.Customer
		)

		BeforeEach(func() {
			luserName = "user"
			password = "somepwdd"
		})

		Context("Customername, password and context are all correct", func() {

			BeforeEach(func() {
				password, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
				stringPassword := string(password)
				sampleCustomer = model.Customer{
					ID:       uint(12345),
					UserName: &luserName,
					Password: stringPassword,
				}
				customerRepo.GetByUserNameReturns(&sampleCustomer, nil)
			})

			It("Returns an authenticated customer", func() {
				customer, err := customerService.Login(context.TODO(), luserName, password)
				Expect(err).NotTo(HaveOccurred())
				Expect(customer).NotTo(BeNil())
				Expect(customer.ID).To(Equal(sampleCustomer.ID))
				Expect(customerRepo.GetByUserNameCallCount()).To(Equal(1))
			})
		})

		Context("No context provided", func() {

			It("Returns an error", func() {
				customer, err := customerService.Login(nil, luserName, password)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal(MissingContext))
				Expect(customer).To(BeNil())
				Expect(customerRepo.GetByUserNameCallCount()).To(Equal(0))
			})
		})

		Context("Customername or password are not provided", func() {

			It("Returns an error", func() {
				customer, err := customerService.Login(context.TODO(), "", "")
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal(UserOrPasswordFail))
				Expect(customer).To(BeNil())
				Expect(customerRepo.GetByUserNameCallCount()).To(Equal(0))
			})
		})

		Context("Wrong password is provided", func() {

			BeforeEach(func() {
				password, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
				stringPassword := string(password)
				sampleCustomer = model.Customer{
					ID:       uint(12345),
					UserName: &luserName,
					Password: stringPassword,
				}
				customerRepo.GetByUserNameReturns(&sampleCustomer, nil)
			})

			It("Returns an error", func() {
				customer, err := customerService.Login(context.TODO(), userName, "some-other-pwd")
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal(UserOrPasswordFail))
				Expect(customer).To(BeNil())
				Expect(customerRepo.GetByUserNameCallCount()).To(Equal(1))
			})
		})

		Context("Underlying repository returns an error", func() {

			BeforeEach(func() {
				customerRepo.GetByUserNameReturns(nil, errors.New("some-error"))
			})

			It("Returns an error", func() {
				customer, err := customerService.Login(context.TODO(), userName, password)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal(UserOrPasswordFail))
				Expect(customer).To(BeNil())
				Expect(customerRepo.GetByUserNameCallCount()).To(Equal(1))
			})
		})
	})

	Describe("Find a customer by username", func() {

		var (
			customer   model.Customer
			userName   string
			fUserName  string
			fFirstName string
			fLastName  string
		)

		BeforeEach(func() {
			fUserName = "some-user-name"
			fFirstName = "john"
			fLastName = "doe"

			customer = model.Customer{
				UserName:  &fUserName,
				FirstName: fFirstName,
				LastName:  fLastName,
			}

			userName = "some-user-name"

			customerRepo.GetByUserNameReturns(&customer, nil)
		})

		Context("Context is correct and username exists", func() {

			It("Returns a customer with the given username", func() {
				foundCustomer, err := customerService.FindCustomerByUserName(context.TODO(), userName)
				Expect(err).NotTo(HaveOccurred())
				Expect(customer).NotTo(BeNil())
				Expect(customerRepo.GetByUserNameCallCount()).To(Equal(1))
				Expect(*foundCustomer.UserName).To(Equal(userName))
			})
		})

		Context("UserName provided but Context missing", func() {

			It("Returns and error", func() {
				foundCustomer, err := customerService.FindCustomerByUserName(nil, userName)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal(MissingContext))
				Expect(foundCustomer).To(BeNil())
			})
		})

		Context("Context and customerName provided but customer is not found", func() {

			BeforeEach(func() {
				customerRepo.GetByUserNameReturns(nil, nil)
			})

			It("Returns no error and no customer", func() {
				foundCustomer, err := customerService.FindCustomerByUserName(context.TODO(), "john doe")
				Expect(err).ToNot(HaveOccurred())
				Expect(foundCustomer).To(BeNil())
			})
		})

		Context("Context and userName are correct but Repo returns error", func() {

			BeforeEach(func() {
				customerRepo.GetByUserNameReturns(nil, errors.New("BAM!"))
			})

			It("Returns an error", func() {
				foundCustomer, err := customerService.FindCustomerByUserName(context.TODO(), userName)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("BAM!"))
				Expect(foundCustomer).To(BeNil())
			})
		})
	})

	Describe("Find a customer by ID", func() {

		var (
			customer   model.Customer
			customerID uint
			fUserName  string
			fFirstName string
			fLastName  string
		)

		BeforeEach(func() {
			fUserName = "some-user-name"
			fFirstName = "john"
			fLastName = "doe"

			customer = model.Customer{
				ID:        uint(12345),
				UserName:  &fUserName,
				FirstName: fFirstName,
				LastName:  fLastName,
			}

			customerID = 12345

			customerRepo.GetByIDReturns(&customer, nil)
		})

		Context("Context is correct and customerID exists", func() {

			It("Returns a customer with the given customerID", func() {
				foundCustomer, err := customerService.FindCustomerByID(context.TODO(), customerID)
				Expect(err).NotTo(HaveOccurred())
				Expect(customer).NotTo(BeNil())
				Expect(customerRepo.GetByIDCallCount()).To(Equal(1))
				Expect(foundCustomer.ID).To(Equal(customerID))
			})
		})

		Context("customerID provided but Context missing", func() {

			It("Returns and error", func() {
				foundCustomer, err := customerService.FindCustomerByID(nil, customerID)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal(MissingContext))
				Expect(foundCustomer).To(BeNil())
			})

		})

		Context("Context and customerID provided but customer is not found", func() {

			BeforeEach(func() {
				customerRepo.GetByIDReturns(nil, nil)
			})

			It("Returns no error and no customer", func() {
				foundCustomer, err := customerService.FindCustomerByID(context.TODO(), 12345)
				Expect(err).ToNot(HaveOccurred())
				Expect(foundCustomer).To(BeNil())
			})
		})

		Context("Context and customerID are correct but Repo returns error", func() {

			BeforeEach(func() {
				customerRepo.GetByIDReturns(nil, errors.New("BAM!"))
			})

			It("Returns an error", func() {
				foundCustomer, err := customerService.FindCustomerByID(context.TODO(), customerID)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("BAM!"))
				Expect(foundCustomer).To(BeNil())
			})
		})
	})

	Describe("Find customers by VetOrg", func() {

		var (
			customer   model.Customer
			vetOrg     model.VetOrg
			customerID uint
			fUserName  string
			fFirstName string
			fLastName  string
		)

		BeforeEach(func() {
			fUserName = "some-user-name"
			fFirstName = "john"
			fLastName = "doe"

			customer = model.Customer{
				ID:        uint(12345),
				UserName:  &fUserName,
				FirstName: fFirstName,
				LastName:  fLastName,
			}

			vetOrg = model.VetOrg{
				ID: uint(12345),
			}

			customerID = 12345

			customerRepo.GetByVetOrgIDReturns([]model.Customer{customer}, nil)
		})

		Context("Context is correct and vetOrg exists", func() {

			It("Returns customers with the given vetOrgID", func() {
				foundCustomer, err := customerService.FindCustomerByVetOrg(context.TODO(), vetOrg)
				Expect(err).NotTo(HaveOccurred())
				Expect(customer).NotTo(BeNil())
				Expect(customerRepo.GetByVetOrgIDCallCount()).To(Equal(1))
				Expect(foundCustomer[0].ID).To(Equal(customerID))
			})
		})

		Context("vetOrg provided but Context missing", func() {

			It("Returns an error", func() {
				foundCustomer, err := customerService.FindCustomerByVetOrg(nil, vetOrg)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal(MissingContext))
				Expect(foundCustomer).To(BeNil())
			})

		})

		Context("vetOrg empty and Context provided", func() {

			It("Returns an error", func() {
				foundCustomer, err := customerService.FindCustomerByVetOrg(context.TODO(), model.VetOrg{})
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal(VetOrgRequired))
				Expect(foundCustomer).To(BeNil())
			})

		})

		Context("Context and vetOrg provided but no customers found", func() {

			BeforeEach(func() {
				customerRepo.GetByVetOrgIDReturns(nil, nil)
			})

			It("Returns no error and no customer", func() {
				foundCustomer, err := customerService.FindCustomerByVetOrg(context.TODO(), vetOrg)
				Expect(err).ToNot(HaveOccurred())
				Expect(foundCustomer).To(BeNil())
			})
		})

		Context("Context and vetOrg are correct but Repo returns error", func() {

			BeforeEach(func() {
				customerRepo.GetByVetOrgIDReturns(nil, errors.New("BAM!"))
			})

			It("Returns an error", func() {
				foundCustomer, err := customerService.FindCustomerByVetOrg(context.TODO(), vetOrg)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("BAM!"))
				Expect(foundCustomer).To(BeNil())
			})
		})
	})
})
