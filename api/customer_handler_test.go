package api_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"

	"github.com/tedsuo/rata"
	"gorm.io/gorm"

	"github.com/jmelchio/vetlab/api/apifakes"
	"github.com/jmelchio/vetlab/model"

	"github.com/jmelchio/vetlab/api"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("CustomerHandler", func() {
	var (
		handler          http.Handler
		recorder         *httptest.ResponseRecorder
		err              error
		customerService  *apifakes.FakeCustomerService
		requestGenerator *rata.RequestGenerator
		userName         string
		firstName        string
		lastName         string
		email            string
		password         string
		vetOrgID         uint
	)

	BeforeEach(func() {
		customerService = new(apifakes.FakeCustomerService)
		handler, err = api.NewCustomerHandler(customerService)
		requestGenerator = rata.NewRequestGenerator("", api.CustomerRoutes)
		Expect(err).NotTo(HaveOccurred())

		userName = "user_name"
		firstName = "first_name"
		lastName = "last_name"
		email = "some_email"
		password = "some_password"
		vetOrgID = uint(1234)
	})

	Describe("Create a customer", func() {

		var (
			createCustomer model.Customer
		)

		BeforeEach(func() {
			createCustomer = model.Customer{
				UserName:  &userName,
				FirstName: firstName,
				LastName:  lastName,
				Email:     email,
				Password:  password,
				VetOrgID:  vetOrgID,
			}
		})

		When("Valid customer information is passed", func() {

			BeforeEach(func() {
				customerService.CreateCustomerReturns(&createCustomer, nil)
				customerBytes, err := json.Marshal(createCustomer)
				Expect(err).NotTo(HaveOccurred())
				recorder = httptest.NewRecorder()
				request, _ := requestGenerator.CreateRequest(api.CreateCustomer, nil, bytes.NewReader(customerBytes))
				handler.ServeHTTP(recorder, request)
			})

			It("Creates a customer and returns 201 status code", func() {
				Expect(recorder.Result().StatusCode).To(Equal(http.StatusCreated))
				respBody, err := io.ReadAll(recorder.Result().Body)
				Expect(err).NotTo(HaveOccurred())

				var newCustomer model.Customer
				err = json.Unmarshal(respBody, &newCustomer)
				Expect(err).NotTo(HaveOccurred())
				Expect(newCustomer).NotTo(BeNil())
				Expect(newCustomer.FirstName).To(Equal(createCustomer.FirstName))
				Expect(newCustomer.LastName).To(Equal(createCustomer.LastName))
				Expect(customerService.CreateCustomerCallCount()).To(Equal(1))
			})

		})

		When("Valid customer information is passed but downstream call fails", func() {

			BeforeEach(func() {
				customerService.CreateCustomerReturns(nil, errors.New("whoot"))
				customerBytes, err := json.Marshal(createCustomer)
				Expect(err).NotTo(HaveOccurred())
				recorder = httptest.NewRecorder()
				request, _ := requestGenerator.CreateRequest(api.CreateCustomer, nil, bytes.NewReader(customerBytes))
				handler.ServeHTTP(recorder, request)
			})

			It("Fails to return a customers and returns 500 status code", func() {
				Expect(recorder.Result().StatusCode).To(Equal(http.StatusInternalServerError))
				respBody, err := io.ReadAll(recorder.Result().Body)
				Expect(err).NotTo(HaveOccurred())
				Expect(err).NotTo(HaveOccurred())
				Expect(string(respBody[0 : len(respBody)-1])).To(Equal(api.UnableToCreateCustomer))
				Expect(customerService.CreateCustomerCallCount()).To(Equal(1))
			})
		})

		Context("Body of the request is empty", func() {

			BeforeEach(func() {
				recorder = httptest.NewRecorder()
				request, _ := requestGenerator.CreateRequest(api.CreateCustomer, nil, nil)
				handler.ServeHTTP(recorder, request)
			})

			It("Fails to create a customer and returns a 400 status code", func() {
				Expect(recorder.Result().StatusCode).To(Equal(http.StatusBadRequest))
				respBody, err := io.ReadAll(recorder.Result().Body)
				Expect(err).NotTo(HaveOccurred())
				Expect(string(respBody[0 : len(respBody)-1])).To(Equal(api.EmptyBody))
				Expect(customerService.CreateCustomerCallCount()).To(Equal(0))
			})
		})

		Context("Body of the request contains invalid data", func() {

			BeforeEach(func() {
				customerBytes, err := json.Marshal("createCustomer")
				Expect(err).NotTo(HaveOccurred())
				recorder = httptest.NewRecorder()
				request, _ := requestGenerator.CreateRequest(api.CreateCustomer, nil, bytes.NewReader(customerBytes))
				handler.ServeHTTP(recorder, request)
			})

			It("Fails to create a customer and returns a 400 status code", func() {
				Expect(recorder.Result().StatusCode).To(Equal(http.StatusBadRequest))
				respBody, err := io.ReadAll(recorder.Result().Body)
				Expect(err).NotTo(HaveOccurred())
				Expect(string(respBody[0 : len(respBody)-1])).To(Equal(api.InvalidBody))
				Expect(customerService.CreateCustomerCallCount()).To(Equal(0))
			})
		})
	})

	Describe("Update a customer", func() {
		var (
			updateCustomer model.Customer
		)

		BeforeEach(func() {
			updateCustomer = model.Customer{
				Model:     gorm.Model{ID: uint(12345)},
				UserName:  &userName,
				FirstName: firstName,
				LastName:  lastName,
				Email:     email,
				Password:  password,
				VetOrgID:  vetOrgID,
			}
		})

		Context("Valid customer information is passed", func() {

			BeforeEach(func() {
				customerService.UpdateCustomerReturns(&updateCustomer, nil)
				customerBytes, err := json.Marshal(updateCustomer)
				Expect(err).NotTo(HaveOccurred())
				recorder = httptest.NewRecorder()
				request, _ := requestGenerator.CreateRequest(api.UpdateCustomer, nil, bytes.NewReader(customerBytes))
				handler.ServeHTTP(recorder, request)
			})

			It("Updates a customer and returns 200 status code", func() {
				Expect(recorder.Result().StatusCode).To(Equal(http.StatusOK))
				respBody, err := io.ReadAll(recorder.Result().Body)
				Expect(err).NotTo(HaveOccurred())

				var newCustomer model.Customer
				err = json.Unmarshal(respBody, &newCustomer)
				Expect(err).NotTo(HaveOccurred())
				Expect(newCustomer).NotTo(BeNil())
				Expect(newCustomer.FirstName).To(Equal(updateCustomer.FirstName))
				Expect(newCustomer.LastName).To(Equal(updateCustomer.LastName))
				Expect(customerService.UpdateCustomerCallCount()).To(Equal(1))
			})
		})

		Context("Valid customer information is passed but downstream call fails", func() {

			BeforeEach(func() {
				customerService.UpdateCustomerReturns(nil, errors.New("whoot"))
				customerBytes, err := json.Marshal(updateCustomer)
				Expect(err).NotTo(HaveOccurred())
				recorder = httptest.NewRecorder()
				request, _ := requestGenerator.CreateRequest(api.UpdateCustomer, nil, bytes.NewReader(customerBytes))
				handler.ServeHTTP(recorder, request)
			})

			It("Fails to return a customers and returns 500 status code", func() {
				Expect(recorder.Result().StatusCode).To(Equal(http.StatusInternalServerError))
				respBody, err := io.ReadAll(recorder.Result().Body)
				Expect(err).NotTo(HaveOccurred())
				Expect(err).NotTo(HaveOccurred())
				Expect(string(respBody[0 : len(respBody)-1])).To(Equal(api.UnableToUpdateCustomer))
				Expect(customerService.UpdateCustomerCallCount()).To(Equal(1))
			})
		})

		Context("Body of the request is empty", func() {

			BeforeEach(func() {
				recorder = httptest.NewRecorder()
				request, _ := requestGenerator.CreateRequest(api.UpdateCustomer, nil, nil)
				handler.ServeHTTP(recorder, request)
			})

			It("Fails to update a customer and returns a 400 status code", func() {
				Expect(recorder.Result().StatusCode).To(Equal(http.StatusBadRequest))
				respBody, err := io.ReadAll(recorder.Result().Body)
				Expect(err).NotTo(HaveOccurred())
				Expect(string(respBody[0 : len(respBody)-1])).To(Equal(api.EmptyBody))
				Expect(customerService.UpdateCustomerCallCount()).To(Equal(0))
			})
		})

		Context("Body of the request contains invalid data", func() {

			BeforeEach(func() {
				customerBytes, err := json.Marshal("updateCustomer")
				Expect(err).NotTo(HaveOccurred())
				recorder = httptest.NewRecorder()
				request, _ := requestGenerator.CreateRequest(api.UpdateCustomer, nil, bytes.NewReader(customerBytes))
				handler.ServeHTTP(recorder, request)
			})

			It("Fails to update a customer and returns a 400 status code", func() {
				Expect(recorder.Result().StatusCode).To(Equal(http.StatusBadRequest))
				respBody, err := io.ReadAll(recorder.Result().Body)
				Expect(err).NotTo(HaveOccurred())
				Expect(string(respBody[0 : len(respBody)-1])).To(Equal(api.InvalidBody))
				Expect(customerService.UpdateCustomerCallCount()).To(Equal(0))
			})
		})
	})

	Describe("Delete a customer", func() {
		var (
			deleteCustomer model.Customer
		)

		BeforeEach(func() {
			deleteCustomer = model.Customer{
				Model:     gorm.Model{ID: uint(12345)},
				UserName:  &userName,
				FirstName: firstName,
				LastName:  lastName,
				Email:     email,
				Password:  password,
				VetOrgID:  vetOrgID,
			}
		})

		Context("Valid customer information is passed", func() {

			BeforeEach(func() {
				customerService.DeleteCustomerReturns(nil)
				customerBytes, err := json.Marshal(deleteCustomer)
				Expect(err).NotTo(HaveOccurred())
				recorder = httptest.NewRecorder()
				request, _ := requestGenerator.CreateRequest(api.DeleteCustomer, nil, bytes.NewReader(customerBytes))
				handler.ServeHTTP(recorder, request)
			})

			It("Deletes a customer and returns 204 status code", func() {
				Expect(recorder.Result().StatusCode).To(Equal(http.StatusNoContent))
				Expect(customerService.DeleteCustomerCallCount()).To(Equal(1))
			})
		})

		Context("Valid customer information is passed but downstream call fails", func() {

			BeforeEach(func() {
				customerService.DeleteCustomerReturns(errors.New("whoot"))
				customerBytes, err := json.Marshal(deleteCustomer)
				Expect(err).NotTo(HaveOccurred())
				recorder = httptest.NewRecorder()
				request, _ := requestGenerator.CreateRequest(api.DeleteCustomer, nil, bytes.NewReader(customerBytes))
				handler.ServeHTTP(recorder, request)
			})

			It("Fails to delete a customer and returns a 500 status code", func() {
				Expect(recorder.Result().StatusCode).To(Equal(http.StatusInternalServerError))
				respBody, err := io.ReadAll(recorder.Result().Body)
				Expect(err).NotTo(HaveOccurred())
				Expect(string(respBody[0 : len(respBody)-1])).To(Equal(api.UnableToDeleteCustomer))
				Expect(customerService.DeleteCustomerCallCount()).To(Equal(1))
			})
		})

		Context("Body of the request is empty", func() {

			BeforeEach(func() {
				recorder = httptest.NewRecorder()
				request, _ := requestGenerator.CreateRequest(api.DeleteCustomer, nil, nil)
				handler.ServeHTTP(recorder, request)
			})

			It("Fails to delete a customer and returns a 400 status code", func() {
				Expect(recorder.Result().StatusCode).To(Equal(http.StatusBadRequest))
				respBody, err := io.ReadAll(recorder.Result().Body)
				Expect(err).NotTo(HaveOccurred())
				Expect(string(respBody[0 : len(respBody)-1])).To(Equal(api.EmptyBody))
				Expect(customerService.DeleteCustomerCallCount()).To(Equal(0))
			})
		})

		Context("Body of the request contains invalid data", func() {

			BeforeEach(func() {
				customerBytes, err := json.Marshal("deleteCustomer")
				Expect(err).NotTo(HaveOccurred())
				recorder = httptest.NewRecorder()
				request, _ := requestGenerator.CreateRequest(api.DeleteCustomer, nil, bytes.NewReader(customerBytes))
				handler.ServeHTTP(recorder, request)
			})

			It("Fails to delete a customer and returns a 400 status code", func() {
				Expect(recorder.Result().StatusCode).To(Equal(http.StatusBadRequest))
				respBody, err := io.ReadAll(recorder.Result().Body)
				Expect(err).NotTo(HaveOccurred())
				Expect(string(respBody[0 : len(respBody)-1])).To(Equal(api.InvalidBody))
				Expect(customerService.DeleteCustomerCallCount()).To(Equal(0))
			})
		})
	})

	Describe("Login a customer", func() {

		var (
			loginRequest  model.LoginRequest
			loginCustomer model.Customer
		)

		BeforeEach(func() {
			loginRequest = model.LoginRequest{
				UserName: "user_name",
				Password: "some_password",
			}

			loginCustomer = model.Customer{
				Model:     gorm.Model{ID: uint(12345)},
				UserName:  &userName,
				FirstName: firstName,
				LastName:  lastName,
				Email:     email,
				Password:  password,
				VetOrgID:  vetOrgID,
			}
		})

		Context("Valid customer information is passed", func() {

			BeforeEach(func() {
				customerService.LoginReturns(&loginCustomer, nil)
				customerBytes, err := json.Marshal(loginRequest)
				Expect(err).NotTo(HaveOccurred())
				recorder = httptest.NewRecorder()
				request, _ := requestGenerator.CreateRequest(api.CustomerLogin, nil, bytes.NewReader(customerBytes))
				handler.ServeHTTP(recorder, request)
			})

			It("Logs a customer in and returns 200 status code", func() {
				Expect(recorder.Result().StatusCode).To(Equal(http.StatusOK))
				respBody, err := io.ReadAll(recorder.Result().Body)
				Expect(err).NotTo(HaveOccurred())

				var newCustomer model.Customer
				err = json.Unmarshal(respBody, &newCustomer)
				Expect(err).NotTo(HaveOccurred())
				Expect(newCustomer).NotTo(BeNil())
				Expect(newCustomer.FirstName).To(Equal(loginCustomer.FirstName))
				Expect(newCustomer.LastName).To(Equal(loginCustomer.LastName))
				Expect(customerService.LoginCallCount()).To(Equal(1))
			})
		})

		Context("Valid customer information is passed but downstream call fails", func() {

			BeforeEach(func() {
				customerService.LoginReturns(nil, errors.New("whoot"))
				customerBytes, err := json.Marshal(loginRequest)
				Expect(err).NotTo(HaveOccurred())
				recorder = httptest.NewRecorder()
				request, _ := requestGenerator.CreateRequest(api.CustomerLogin, nil, bytes.NewReader(customerBytes))
				handler.ServeHTTP(recorder, request)
			})

			It("Fails to login a customer and returns a 500 status code", func() {
				Expect(recorder.Result().StatusCode).To(Equal(http.StatusInternalServerError))
				respBody, err := io.ReadAll(recorder.Result().Body)
				Expect(err).NotTo(HaveOccurred())
				Expect(string(respBody[0 : len(respBody)-1])).To(Equal(api.UnableToLoginCustomer))
				Expect(customerService.LoginCallCount()).To(Equal(1))
			})
		})

		Context("Body of the request is empty", func() {

			BeforeEach(func() {
				recorder = httptest.NewRecorder()
				request, _ := requestGenerator.CreateRequest(api.CustomerLogin, nil, nil)
				handler.ServeHTTP(recorder, request)
			})

			It("Fails to login a customer and returns a 400 status code", func() {
				Expect(recorder.Result().StatusCode).To(Equal(http.StatusBadRequest))
				respBody, err := io.ReadAll(recorder.Result().Body)
				Expect(err).NotTo(HaveOccurred())
				Expect(string(respBody[0 : len(respBody)-1])).To(Equal(api.EmptyBody))
				Expect(customerService.DeleteCustomerCallCount()).To(Equal(0))
			})
		})

		Context("Body of the request contains invalid data", func() {

			BeforeEach(func() {
				customerBytes, err := json.Marshal("loginRequest")
				Expect(err).NotTo(HaveOccurred())
				recorder = httptest.NewRecorder()
				request, _ := requestGenerator.CreateRequest(api.CustomerLogin, nil, bytes.NewReader(customerBytes))
				handler.ServeHTTP(recorder, request)
			})

			It("Fails to login a customer and returns a 400 status code", func() {
				Expect(recorder.Result().StatusCode).To(Equal(http.StatusBadRequest))
				respBody, err := io.ReadAll(recorder.Result().Body)
				Expect(err).NotTo(HaveOccurred())
				Expect(string(respBody[0 : len(respBody)-1])).To(Equal(api.InvalidBody))
				Expect(customerService.DeleteCustomerCallCount()).To(Equal(0))
			})
		})
	})

	Describe("Find a customer", func() {

		var (
			customerID     uint
			sampleCustomer model.Customer
		)

		BeforeEach(func() {
			sampleCustomer = model.Customer{
				Model:     gorm.Model{ID: uint(12345)},
				UserName:  &userName,
				FirstName: firstName,
				LastName:  lastName,
				Email:     email,
				Password:  password,
				VetOrgID:  vetOrgID,
			}
		})

		Context("Valid userName is passed", func() {

			Context("Customer with userName exists", func() {

				BeforeEach(func() {
					customerService.FindCustomerByUserNameReturns(&sampleCustomer, nil)
					recorder = httptest.NewRecorder()
					request, _ := requestGenerator.CreateRequest(
						api.FindCustomerByUserName,
						rata.Params{"user_name": userName},
						nil)
					handler.ServeHTTP(recorder, request)
				})

				It("Finds and returns a customer and returns 200 status code", func() {
					Expect(recorder.Result().StatusCode).To(Equal(http.StatusOK))
					respBody, err := io.ReadAll(recorder.Result().Body)
					Expect(err).NotTo(HaveOccurred())

					var foundCustomer model.Customer
					err = json.Unmarshal(respBody, &foundCustomer)
					Expect(err).NotTo(HaveOccurred())
					Expect(foundCustomer).NotTo(BeNil())
					Expect(foundCustomer.FirstName).To(Equal(sampleCustomer.FirstName))
					Expect(foundCustomer.LastName).To(Equal(sampleCustomer.LastName))
					Expect(customerService.FindCustomerByUserNameCallCount()).To(Equal(1))
				})
			})

			Context("Customer with userName does not exist", func() {

				BeforeEach(func() {
					customerService.FindCustomerByUserNameReturns(nil, errors.New("whoot"))
					recorder = httptest.NewRecorder()
					request, _ := requestGenerator.CreateRequest(
						api.FindCustomerByUserName,
						rata.Params{"user_name": "bad_name"},
						nil)
					handler.ServeHTTP(recorder, request)
				})

				It("Doesn't find a customer and returns 404 status code", func() {
					Expect(recorder.Result().StatusCode).To(Equal(http.StatusNotFound))
					respBody, err := io.ReadAll(recorder.Result().Body)
					Expect(err).NotTo(HaveOccurred())

					Expect(string(respBody[0 : len(respBody)-1])).To(Equal(api.UnableToFindCustomer))
					Expect(customerService.FindCustomerByUserNameCallCount()).To(Equal(1))
				})
			})
		})

		Context("Valid customerID is passed", func() {

			Context("Customer with customerID exists", func() {

				BeforeEach(func() {
					customerService.FindCustomerByIDReturns(&sampleCustomer, nil)
					customerID = uint(12345)
					recorder = httptest.NewRecorder()
					request, _ := requestGenerator.CreateRequest(
						api.FindCustomer,
						rata.Params{"customer_id": fmt.Sprint(customerID)},
						nil)
					handler.ServeHTTP(recorder, request)
				})

				It("Finds and returns a customer and returns 200 status code", func() {
					Expect(recorder.Result().StatusCode).To(Equal(http.StatusOK))
					respBody, err := io.ReadAll(recorder.Result().Body)
					Expect(err).NotTo(HaveOccurred())

					var foundCustomer model.Customer
					err = json.Unmarshal(respBody, &foundCustomer)
					Expect(err).NotTo(HaveOccurred())
					Expect(foundCustomer).NotTo(BeNil())
					Expect(foundCustomer.FirstName).To(Equal(sampleCustomer.FirstName))
					Expect(foundCustomer.LastName).To(Equal(sampleCustomer.LastName))
					Expect(customerService.FindCustomerByIDCallCount()).To(Equal(1))
				})
			})
		})

		Context("Valid search parameter information is passed but downstream call fails", func() {

			BeforeEach(func() {
				customerService.FindCustomerByUserNameReturns(nil, errors.New("whoot"))
				recorder = httptest.NewRecorder()
				request, _ := requestGenerator.CreateRequest(
					api.FindCustomerByUserName,
					rata.Params{"user_name": userName},
					nil)
				handler.ServeHTTP(recorder, request)
			})

			It("Fails to find a customer and returns a 404 status code", func() {
				Expect(recorder.Result().StatusCode).To(Equal(http.StatusNotFound))
				respBody, err := io.ReadAll(recorder.Result().Body)
				Expect(err).NotTo(HaveOccurred())
				Expect(string(respBody[0 : len(respBody)-1])).To(Equal(api.UnableToFindCustomer))
				Expect(customerService.FindCustomerByUserNameCallCount()).To(Equal(1))
			})
		})
	})
})
