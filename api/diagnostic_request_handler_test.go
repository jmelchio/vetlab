package api_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/tedsuo/rata"
	"gorm.io/gorm"

	"github.com/jmelchio/vetlab/api"
	"github.com/jmelchio/vetlab/api/apifakes"
	"github.com/jmelchio/vetlab/model"
)

var _ = Describe("DiagnosticRequestHandler", func() {
	var (
		handler                  http.Handler
		recorder                 *httptest.ResponseRecorder
		err                      error
		diagnosticRequestService *apifakes.FakeDiagnosticRequestService
		vetOrgService            *apifakes.FakeVetOrgService
		userService              *apifakes.FakeUserService
		customerService          *apifakes.FakeCustomerService
		requestGenerator         *rata.RequestGenerator
	)

	BeforeEach(func() {
		diagnosticRequestService = new(apifakes.FakeDiagnosticRequestService)
		vetOrgService = new(apifakes.FakeVetOrgService)
		userService = new(apifakes.FakeUserService)
		customerService = new(apifakes.FakeCustomerService)
		handler, err = api.NewDiagnosticRequestHandler(
			diagnosticRequestService,
			vetOrgService,
			userService,
			customerService,
		)
		requestGenerator = rata.NewRequestGenerator("", api.DiagnosticRequestRoutes)
		Expect(err).NotTo(HaveOccurred())
	})

	Describe("Create a DiagnosticRequest", func() {
		var (
			diagnosticRequest model.DiagnosticRequest
		)

		BeforeEach(func() {
			diagnosticRequest = model.DiagnosticRequest{
				VetOrgID:    uint(12345),
				CustomerID:  uint(54321),
				UserID:      uint(23451),
				Description: "this is a good request",
			}
		})

		Context("Valid request information is passed", func() {

			BeforeEach(func() {
				diagnosticRequestService.SubmitDiagnosticRequestReturns(&diagnosticRequest, nil)
				requestBytes, err := json.Marshal(diagnosticRequest)
				Expect(err).NotTo(HaveOccurred())
				recorder = httptest.NewRecorder()
				request, _ := requestGenerator.CreateRequest(api.SubmitDiagnosticRequest, nil, bytes.NewReader(requestBytes))
				handler.ServeHTTP(recorder, request)
			})

			It("Creates a diagnosticRequest and returns 201 status code", func() {
				Expect(recorder.Result().StatusCode).To(Equal(http.StatusCreated))
				respBody, err := io.ReadAll(recorder.Result().Body)
				Expect(err).NotTo(HaveOccurred())

				var newDiagnosticRequest model.DiagnosticRequest
				err = json.Unmarshal(respBody, &newDiagnosticRequest)
				Expect(err).NotTo(HaveOccurred())
				Expect(newDiagnosticRequest).NotTo(BeNil())
				Expect(newDiagnosticRequest.ID).NotTo(BeNil())
				Expect(newDiagnosticRequest.VetOrgID).To(Equal(diagnosticRequest.VetOrgID))
				Expect(newDiagnosticRequest.CustomerID).To(Equal(diagnosticRequest.CustomerID))
				Expect(newDiagnosticRequest.UserID).To(Equal(diagnosticRequest.UserID))
				Expect(newDiagnosticRequest.Description).To(Equal(diagnosticRequest.Description))
				Expect(diagnosticRequestService.SubmitDiagnosticRequestCallCount()).To(Equal(1))
			})
		})

		Context("Valid diagnostic request is passed but downstream call fails", func() {

			BeforeEach(func() {
				diagnosticRequestService.SubmitDiagnosticRequestReturns(nil, errors.New("whoot"))
				userBytes, err := json.Marshal(diagnosticRequest)
				Expect(err).NotTo(HaveOccurred())
				recorder = httptest.NewRecorder()
				request, _ := requestGenerator.CreateRequest(api.SubmitDiagnosticRequest, nil, bytes.NewReader(userBytes))
				handler.ServeHTTP(recorder, request)
			})

			It("Fails to return a users and returns 500 status code", func() {
				Expect(recorder.Result().StatusCode).To(Equal(http.StatusInternalServerError))
				respBody, err := io.ReadAll(recorder.Result().Body)
				Expect(err).NotTo(HaveOccurred())
				Expect(err).NotTo(HaveOccurred())
				Expect(string(respBody[0 : len(respBody)-1])).To(Equal(api.UnableToSubmitDiagnosticRequest))
				Expect(diagnosticRequestService.SubmitDiagnosticRequestCallCount()).To(Equal(1))
			})
		})

		Context("Body of the request is empty", func() {

			BeforeEach(func() {
				recorder = httptest.NewRecorder()
				request, _ := requestGenerator.CreateRequest(api.SubmitDiagnosticRequest, nil, nil)
				handler.ServeHTTP(recorder, request)
			})

			It("Fails to create a user and returns a 400 status code", func() {
				Expect(recorder.Result().StatusCode).To(Equal(http.StatusBadRequest))
				respBody, err := io.ReadAll(recorder.Result().Body)
				Expect(err).NotTo(HaveOccurred())
				Expect(string(respBody[0 : len(respBody)-1])).To(Equal(api.EmptyBody))
				Expect(diagnosticRequestService.SubmitDiagnosticRequestCallCount()).To(Equal(0))
			})
		})

		Context("Body of the request contains invalid data", func() {

			BeforeEach(func() {
				userBytes, err := json.Marshal("diagnosticRequest")
				Expect(err).NotTo(HaveOccurred())
				recorder = httptest.NewRecorder()
				request, _ := requestGenerator.CreateRequest(api.SubmitDiagnosticRequest, nil, bytes.NewReader(userBytes))
				handler.ServeHTTP(recorder, request)
			})

			It("Fails to create a user and returns a 400 status code", func() {
				Expect(recorder.Result().StatusCode).To(Equal(http.StatusBadRequest))
				respBody, err := io.ReadAll(recorder.Result().Body)
				Expect(err).NotTo(HaveOccurred())
				Expect(string(respBody[0 : len(respBody)-1])).To(Equal(api.InvalidBody))
				Expect(diagnosticRequestService.SubmitDiagnosticRequestCallCount()).To(Equal(0))
			})
		})
	})

	Describe("Find a diagnostic request by ID", func() {

		var (
			diagnosticRequest model.DiagnosticRequest
		)

		BeforeEach(func() {
			diagnosticRequest = model.DiagnosticRequest{
				Model:       gorm.Model{ID: uint(98765)},
				VetOrgID:    uint(12345),
				CustomerID:  uint(54321),
				UserID:      uint(23451),
				Description: "this is a good request",
			}
		})

		Context("Valid request information is provided", func() {

			Context("Id is present in backing storage", func() {

				BeforeEach(func() {
					diagnosticRequestService.FindRequestByIDReturns(&diagnosticRequest, nil)
					recorder = httptest.NewRecorder()
					params := rata.Params{
						"request_id": "98765",
					}
					request, _ := requestGenerator.CreateRequest(api.DiagnosticRequestByID, params, nil)
					handler.ServeHTTP(recorder, request)
				})

				It("Returns the requested diagnostic request information", func() {
					Expect(recorder.Result().StatusCode).To(Equal(http.StatusOK))
					respBody, err := io.ReadAll(recorder.Result().Body)
					Expect(err).NotTo(HaveOccurred())

					var findDiagnosticRequest model.DiagnosticRequest
					err = json.Unmarshal(respBody, &findDiagnosticRequest)
					Expect(err).NotTo(HaveOccurred())
					Expect(findDiagnosticRequest).NotTo(BeNil())
					Expect(findDiagnosticRequest.ID).To(Equal(diagnosticRequest.ID))
					Expect(findDiagnosticRequest.Description).To(Equal(diagnosticRequest.Description))
					Expect(diagnosticRequestService.FindRequestByIDCallCount()).To(Equal(1))
				})
			})

			Context("Id is not found in backing storage", func() {

				BeforeEach(func() {
					notFoundError := errors.New("not found")
					diagnosticRequestService.FindRequestByIDReturns(nil, notFoundError)
					recorder = httptest.NewRecorder()
					params := rata.Params{
						"request_id": "12345",
					}
					request, _ := requestGenerator.CreateRequest(api.DiagnosticRequestByID, params, nil)
					handler.ServeHTTP(recorder, request)
				})

				It("Returns an error indicating it is unable to find diagnostic request", func() {
					Expect(recorder.Result().StatusCode).To(Equal(http.StatusNotFound))
					respBody, err := io.ReadAll(recorder.Result().Body)
					Expect(err).NotTo(HaveOccurred())
					Expect(strings.TrimSpace(string(respBody))).To(Equal(api.ErrorFetchingDiagnosticRequests))
					Expect(diagnosticRequestService.FindRequestByIDCallCount()).To(Equal(1))
				})
			})

			Context("No request id provided in the request", func() {

				BeforeEach(func() {
					recorder = httptest.NewRecorder()
					request, _ := http.NewRequest("GET", "/diagnosticrequest/", nil)
					handler.ServeHTTP(recorder, request)
				})

				It("Returns an error indicating it cannot find the page", func() {
					Expect(recorder.Result().StatusCode).To(Equal(http.StatusNotFound))
					Expect(diagnosticRequestService.FindRequestByIDCallCount()).To(Equal(0))
				})
			})

			Context("Invalid request id provided in the request", func() {

				BeforeEach(func() {
					recorder = httptest.NewRecorder()
					params := rata.Params{
						"request_id": "one",
					}
					request, _ := requestGenerator.CreateRequest(api.DiagnosticRequestByID, params, nil)
					handler.ServeHTTP(recorder, request)
				})

				It("Returns an error indicating it cannot parse the request", func() {
					Expect(recorder.Result().StatusCode).To(Equal(http.StatusBadRequest))
					respBody, err := io.ReadAll(recorder.Result().Body)
					Expect(err).NotTo(HaveOccurred())
					Expect(strings.TrimSpace(string(respBody))).To(Equal(api.UnableToParseParams))
					Expect(diagnosticRequestService.FindRequestByIDCallCount()).To(Equal(0))
				})
			})
		})
	})

	Describe("Fetch diagnostic requests by VetOrg", func() {

		var (
			diagnosticRequest     model.DiagnosticRequest
			diagnosticRequestList []model.DiagnosticRequest
			vetOrg                model.VetOrg
		)

		BeforeEach(func() {
			diagnosticRequest = model.DiagnosticRequest{
				Model:       gorm.Model{ID: uint(98765)},
				VetOrgID:    uint(12345),
				CustomerID:  uint(54321),
				UserID:      uint(23451),
				Description: "this is a good request",
			}
			diagnosticRequestList = []model.DiagnosticRequest{diagnosticRequest}
			vetOrgName := "Veterinary Clinic One"
			vetOrg = model.VetOrg{
				Model:   gorm.Model{ID: uint(12345)},
				OrgName: &vetOrgName,
			}
		})

		Context("Valid request information is provided", func() {

			Context("VetOrg and Requests are present in backing storage", func() {

				BeforeEach(func() {
					vetOrgService.FindVetOrgByIDReturns(&vetOrg, nil)
					diagnosticRequestService.FindRequestByVetOrgReturns(diagnosticRequestList, nil)
					recorder = httptest.NewRecorder()
					params := rata.Params{
						"vetorg_id": "12345",
					}
					request, _ := requestGenerator.CreateRequest(api.DiagnosticRequestsByVetOrgID, params, nil)
					handler.ServeHTTP(recorder, request)
				})

				It("Returns the requested diagnostic request information", func() {
					Expect(recorder.Result().StatusCode).To(Equal(http.StatusOK))
					respBody, err := io.ReadAll(recorder.Result().Body)
					Expect(err).NotTo(HaveOccurred())

					var findDiagnosticRequest []model.DiagnosticRequest
					err = json.Unmarshal(respBody, &findDiagnosticRequest)
					Expect(err).NotTo(HaveOccurred())
					Expect(findDiagnosticRequest).NotTo(BeNil())
					Expect(findDiagnosticRequest[0].ID).To(Equal(diagnosticRequest.ID))
					Expect(findDiagnosticRequest[0].Description).To(Equal(diagnosticRequest.Description))
					Expect(vetOrgService.FindVetOrgByIDCallCount()).To(Equal(1))
					Expect(diagnosticRequestService.FindRequestByVetOrgCallCount()).To(Equal(1))
				})
			})

			Context("VetOrg not found in backing storage", func() {

				BeforeEach(func() {
					notFoundError := errors.New("not found")
					vetOrgService.FindVetOrgByIDReturns(nil, notFoundError)
					diagnosticRequestService.FindRequestByVetOrgReturns(nil, notFoundError)
					recorder = httptest.NewRecorder()
					params := rata.Params{
						"vetorg_id": "12345",
					}
					request, _ := requestGenerator.CreateRequest(api.DiagnosticRequestsByVetOrgID, params, nil)
					handler.ServeHTTP(recorder, request)
				})

				It("Returns an error indicating it is unable to find VetOrg", func() {
					Expect(recorder.Result().StatusCode).To(Equal(http.StatusNotFound))
					respBody, err := io.ReadAll(recorder.Result().Body)
					Expect(err).NotTo(HaveOccurred())
					Expect(strings.TrimSpace(string(respBody))).To(Equal(api.ErrorFetchingVetOrg))
					Expect(vetOrgService.FindVetOrgByIDCallCount()).To(Equal(1))
					Expect(diagnosticRequestService.FindRequestByVetOrgCallCount()).To(Equal(0))
				})
			})

			Context("Request(s) not found in backing storage", func() {

				BeforeEach(func() {
					notFoundError := errors.New("not found")
					vetOrgService.FindVetOrgByIDReturns(&vetOrg, nil)
					diagnosticRequestService.FindRequestByVetOrgReturns(nil, notFoundError)
					recorder = httptest.NewRecorder()
					params := rata.Params{
						"vetorg_id": "12345",
					}
					request, _ := requestGenerator.CreateRequest(api.DiagnosticRequestsByVetOrgID, params, nil)
					handler.ServeHTTP(recorder, request)
				})

				It("Returns an error indicating it is unable to find diagnostic request", func() {
					Expect(recorder.Result().StatusCode).To(Equal(http.StatusNotFound))
					respBody, err := io.ReadAll(recorder.Result().Body)
					Expect(err).NotTo(HaveOccurred())
					Expect(strings.TrimSpace(string(respBody))).To(Equal(api.ErrorFetchingDiagnosticRequests))
					Expect(vetOrgService.FindVetOrgByIDCallCount()).To(Equal(1))
					Expect(diagnosticRequestService.FindRequestByVetOrgCallCount()).To(Equal(1))
				})
			})

			Context("No vetorg id provided in the request", func() {

				BeforeEach(func() {
					recorder = httptest.NewRecorder()
					request, _ := http.NewRequest("GET", "/diagnosticrequest/vetorg/", nil)
					handler.ServeHTTP(recorder, request)
				})

				It("Returns an error indicating it cannot find the page", func() {
					Expect(recorder.Result().StatusCode).To(Equal(http.StatusNotFound))
					Expect(vetOrgService.FindVetOrgByIDCallCount()).To(Equal(0))
					Expect(diagnosticRequestService.FindRequestByIDCallCount()).To(Equal(0))
				})
			})

			Context("Invalid vetorg id provided in the request", func() {

				BeforeEach(func() {
					recorder = httptest.NewRecorder()
					params := rata.Params{
						"vetorg_id": "one",
					}
					request, _ := requestGenerator.CreateRequest(api.DiagnosticRequestsByVetOrgID, params, nil)
					handler.ServeHTTP(recorder, request)
				})

				It("Returns an error indicating it cannot parse the request", func() {
					Expect(recorder.Result().StatusCode).To(Equal(http.StatusBadRequest))
					respBody, err := io.ReadAll(recorder.Result().Body)
					Expect(err).NotTo(HaveOccurred())
					Expect(strings.TrimSpace(string(respBody))).To(Equal(api.UnableToParseParams))
					Expect(vetOrgService.FindVetOrgByIDCallCount()).To(Equal(0))
					Expect(diagnosticRequestService.FindRequestByVetOrgCallCount()).To(Equal(0))
				})
			})
		})
	})

	Describe("Fetch diagnostic requests by user", func() {

		var (
			diagnosticRequest     model.DiagnosticRequest
			diagnosticRequestList []model.DiagnosticRequest
			user                  model.User
		)

		BeforeEach(func() {
			diagnosticRequest = model.DiagnosticRequest{
				Model:       gorm.Model{ID: uint(98765)},
				VetOrgID:    uint(12345),
				CustomerID:  uint(54321),
				UserID:      uint(23451),
				Description: "this is a good request",
			}
			diagnosticRequestList = []model.DiagnosticRequest{diagnosticRequest}
			userName := "Some User Name"
			user = model.User{
				Model:    gorm.Model{ID: uint(23451)},
				UserName: &userName,
			}
		})

		Context("Valid request information is provided", func() {

			Context("User and Requests are present in backing storage", func() {

				BeforeEach(func() {
					userService.FindUserByIDReturns(&user, nil)
					diagnosticRequestService.FindRequestByUserReturns(diagnosticRequestList, nil)
					recorder = httptest.NewRecorder()
					params := rata.Params{
						"user_id": "12345",
					}
					request, _ := requestGenerator.CreateRequest(api.DiagnosticRequestsByUserID, params, nil)
					handler.ServeHTTP(recorder, request)
				})

				It("Returns the requested diagnostic request information", func() {
					Expect(recorder.Result().StatusCode).To(Equal(http.StatusOK))
					respBody, err := io.ReadAll(recorder.Result().Body)
					Expect(err).NotTo(HaveOccurred())

					var findDiagnosticRequest []model.DiagnosticRequest
					err = json.Unmarshal(respBody, &findDiagnosticRequest)
					Expect(err).NotTo(HaveOccurred())
					Expect(findDiagnosticRequest).NotTo(BeNil())
					Expect(findDiagnosticRequest[0].ID).To(Equal(diagnosticRequest.ID))
					Expect(findDiagnosticRequest[0].Description).To(Equal(diagnosticRequest.Description))
					Expect(userService.FindUserByIDCallCount()).To(Equal(1))
					Expect(diagnosticRequestService.FindRequestByUserCallCount()).To(Equal(1))
				})
			})

			Context("User not found in backing storage", func() {

				BeforeEach(func() {
					notFoundError := errors.New("not found")
					userService.FindUserByIDReturns(nil, notFoundError)
					diagnosticRequestService.FindRequestByUserReturns(nil, notFoundError)
					recorder = httptest.NewRecorder()
					params := rata.Params{
						"user_id": "12345",
					}
					request, _ := requestGenerator.CreateRequest(api.DiagnosticRequestsByUserID, params, nil)
					handler.ServeHTTP(recorder, request)
				})

				It("Returns an error indicating it is unable to find User", func() {
					Expect(recorder.Result().StatusCode).To(Equal(http.StatusNotFound))
					respBody, err := io.ReadAll(recorder.Result().Body)
					Expect(err).NotTo(HaveOccurred())
					Expect(strings.TrimSpace(string(respBody))).To(Equal(api.ErrorFetchingUser))
					Expect(userService.FindUserByIDCallCount()).To(Equal(1))
					Expect(diagnosticRequestService.FindRequestByUserCallCount()).To(Equal(0))
				})
			})

			Context("Request(s) not found in backing storage", func() {

				BeforeEach(func() {
					notFoundError := errors.New("not found")
					userService.FindUserByIDReturns(&user, nil)
					diagnosticRequestService.FindRequestByUserReturns(nil, notFoundError)
					recorder = httptest.NewRecorder()
					params := rata.Params{
						"user_id": "12345",
					}
					request, _ := requestGenerator.CreateRequest(api.DiagnosticRequestsByUserID, params, nil)
					handler.ServeHTTP(recorder, request)
				})

				It("Returns an error indicating it is unable to find diagnostic request", func() {
					Expect(recorder.Result().StatusCode).To(Equal(http.StatusNotFound))
					respBody, err := io.ReadAll(recorder.Result().Body)
					Expect(err).NotTo(HaveOccurred())
					Expect(strings.TrimSpace(string(respBody))).To(Equal(api.ErrorFetchingDiagnosticRequests))
					Expect(userService.FindUserByIDCallCount()).To(Equal(1))
					Expect(diagnosticRequestService.FindRequestByUserCallCount()).To(Equal(1))
				})
			})

			Context("No user id provided in the request", func() {

				BeforeEach(func() {
					recorder = httptest.NewRecorder()
					request, _ := http.NewRequest("GET", "/diagnosticrequest/user/", nil)
					handler.ServeHTTP(recorder, request)
				})

				It("Returns an error indicating it cannot find the page", func() {
					Expect(recorder.Result().StatusCode).To(Equal(http.StatusNotFound))
					Expect(userService.FindUserByIDCallCount()).To(Equal(0))
					Expect(diagnosticRequestService.FindRequestByIDCallCount()).To(Equal(0))
				})
			})

			Context("Invalid user id provided in the request", func() {

				BeforeEach(func() {
					recorder = httptest.NewRecorder()
					params := rata.Params{
						"user_id": "one",
					}
					request, _ := requestGenerator.CreateRequest(api.DiagnosticRequestsByUserID, params, nil)
					handler.ServeHTTP(recorder, request)
				})

				It("Returns an error indicating it cannot parse the request", func() {
					Expect(recorder.Result().StatusCode).To(Equal(http.StatusBadRequest))
					respBody, err := io.ReadAll(recorder.Result().Body)
					Expect(err).NotTo(HaveOccurred())
					Expect(strings.TrimSpace(string(respBody))).To(Equal(api.UnableToParseParams))
					Expect(userService.FindUserByIDCallCount()).To(Equal(0))
					Expect(diagnosticRequestService.FindRequestByUserCallCount()).To(Equal(0))
				})
			})
		})
	})

	Describe("Find diagnostic requests by customer", func() {

		var (
			diagnosticRequest     model.DiagnosticRequest
			diagnosticRequestList []model.DiagnosticRequest
			customer              model.Customer
		)

		BeforeEach(func() {
			diagnosticRequest = model.DiagnosticRequest{
				Model:       gorm.Model{ID: uint(98765)},
				VetOrgID:    uint(12345),
				CustomerID:  uint(54321),
				UserID:      uint(23451),
				Description: "this is a good request",
			}
			diagnosticRequestList = []model.DiagnosticRequest{diagnosticRequest}
			customerName := "Some Customer Name"
			customer = model.Customer{
				Model:    gorm.Model{ID: uint(23451)},
				UserName: &customerName,
			}
		})

		Context("Valid request information is provided", func() {

			Context("Customer and Requests are present in backing storage", func() {

				BeforeEach(func() {
					customerService.FindCustomerByIDReturns(&customer, nil)
					diagnosticRequestService.FindRequestByCustomerReturns(diagnosticRequestList, nil)
					recorder = httptest.NewRecorder()
					params := rata.Params{
						"customer_id": "12345",
					}
					request, _ := requestGenerator.CreateRequest(api.DiagnosticRequestsByCustomerID, params, nil)
					handler.ServeHTTP(recorder, request)
				})

				It("Returns the requested diagnostic request information", func() {
					Expect(recorder.Result().StatusCode).To(Equal(http.StatusOK))
					respBody, err := io.ReadAll(recorder.Result().Body)
					Expect(err).NotTo(HaveOccurred())

					var findDiagnosticRequest []model.DiagnosticRequest
					err = json.Unmarshal(respBody, &findDiagnosticRequest)
					Expect(err).NotTo(HaveOccurred())
					Expect(findDiagnosticRequest).NotTo(BeNil())
					Expect(findDiagnosticRequest[0].ID).To(Equal(diagnosticRequest.ID))
					Expect(findDiagnosticRequest[0].Description).To(Equal(diagnosticRequest.Description))
					Expect(customerService.FindCustomerByIDCallCount()).To(Equal(1))
					Expect(diagnosticRequestService.FindRequestByCustomerCallCount()).To(Equal(1))
				})
			})

			Context("Customer not found in backing storage", func() {

				BeforeEach(func() {
					notFoundError := errors.New("not found")
					customerService.FindCustomerByIDReturns(nil, notFoundError)
					diagnosticRequestService.FindRequestByCustomerReturns(nil, notFoundError)
					recorder = httptest.NewRecorder()
					params := rata.Params{
						"customer_id": "12345",
					}
					request, _ := requestGenerator.CreateRequest(api.DiagnosticRequestsByCustomerID, params, nil)
					handler.ServeHTTP(recorder, request)
				})

				It("Returns an error indicating it is unable to find Customer", func() {
					Expect(recorder.Result().StatusCode).To(Equal(http.StatusNotFound))
					respBody, err := io.ReadAll(recorder.Result().Body)
					Expect(err).NotTo(HaveOccurred())
					Expect(strings.TrimSpace(string(respBody))).To(Equal(api.ErrorFetchingCustomer))
					Expect(customerService.FindCustomerByIDCallCount()).To(Equal(1))
					Expect(diagnosticRequestService.FindRequestByCustomerCallCount()).To(Equal(0))
				})
			})

			Context("Request(s) not found in backing storage", func() {

				BeforeEach(func() {
					notFoundError := errors.New("not found")
					customerService.FindCustomerByIDReturns(&customer, nil)
					diagnosticRequestService.FindRequestByCustomerReturns(nil, notFoundError)
					recorder = httptest.NewRecorder()
					params := rata.Params{
						"customer_id": "12345",
					}
					request, _ := requestGenerator.CreateRequest(api.DiagnosticRequestsByCustomerID, params, nil)
					handler.ServeHTTP(recorder, request)
				})

				It("Returns an error indicating it is unable to find diagnostic request", func() {
					Expect(recorder.Result().StatusCode).To(Equal(http.StatusNotFound))
					respBody, err := io.ReadAll(recorder.Result().Body)
					Expect(err).NotTo(HaveOccurred())
					Expect(strings.TrimSpace(string(respBody))).To(Equal(api.ErrorFetchingDiagnosticRequests))
					Expect(customerService.FindCustomerByIDCallCount()).To(Equal(1))
					Expect(diagnosticRequestService.FindRequestByCustomerCallCount()).To(Equal(1))
				})
			})

			Context("No customer id provided in the request", func() {

				BeforeEach(func() {
					recorder = httptest.NewRecorder()
					request, _ := http.NewRequest("GET", "/diagnosticrequest/customer/", nil)
					handler.ServeHTTP(recorder, request)
				})

				It("Returns an error indicating it cannot find the page", func() {
					Expect(recorder.Result().StatusCode).To(Equal(http.StatusNotFound))
					Expect(customerService.FindCustomerByIDCallCount()).To(Equal(0))
					Expect(diagnosticRequestService.FindRequestByIDCallCount()).To(Equal(0))
				})
			})

			Context("Invalid customer id provided in the request", func() {

				BeforeEach(func() {
					recorder = httptest.NewRecorder()
					params := rata.Params{
						"customer_id": "one",
					}
					request, _ := requestGenerator.CreateRequest(api.DiagnosticRequestsByCustomerID, params, nil)
					handler.ServeHTTP(recorder, request)
				})

				It("Returns an error indicating it cannot parse the request", func() {
					Expect(recorder.Result().StatusCode).To(Equal(http.StatusBadRequest))
					respBody, err := io.ReadAll(recorder.Result().Body)
					Expect(err).NotTo(HaveOccurred())
					Expect(strings.TrimSpace(string(respBody))).To(Equal(api.UnableToParseParams))
					Expect(customerService.FindCustomerByIDCallCount()).To(Equal(0))
					Expect(diagnosticRequestService.FindRequestByCustomerCallCount()).To(Equal(0))
				})
			})
		})
	})

	Describe("Find diagnostic requests by vetorg and date range", func() {

		var (
			diagnosticRequest     model.DiagnosticRequest
			diagnosticRequestList []model.DiagnosticRequest
			vetOrg                model.VetOrg
			requestDate           time.Time
		)

		BeforeEach(func() {
			requestDate = time.Date(2019, time.April, 10, 23, 0, 0, 0, time.UTC)
			diagnosticRequest = model.DiagnosticRequest{
				Model:       gorm.Model{ID: uint(98765)},
				VetOrgID:    uint(12345),
				CustomerID:  uint(54321),
				UserID:      uint(23451),
				Date:        &requestDate,
				Description: "this is a good request",
			}
			diagnosticRequestList = []model.DiagnosticRequest{diagnosticRequest}
			vetOrgName := "Some VetOrg Name"
			vetOrg = model.VetOrg{
				Model:   gorm.Model{ID: uint(12345)},
				OrgName: &vetOrgName,
			}
		})

		Context("Valid request information is provided", func() {

			Context("VetOrg and Requests are present in backing storage", func() {

				BeforeEach(func() {
					vetOrgService.FindVetOrgByIDReturns(&vetOrg, nil)
					diagnosticRequestService.FindRequestByDateRangeReturns(diagnosticRequestList, nil)
					recorder = httptest.NewRecorder()
					params := rata.Params{
						"vetorg_id":  "12345",
						"start_date": "20190101",
						"end_date":   "20191231",
					}
					request, _ := requestGenerator.CreateRequest(api.DiagnosticRequestsByVetOrgIDAndDateRange, params, nil)
					handler.ServeHTTP(recorder, request)
				})

				It("Returns the requested diagnostic request information", func() {
					Expect(recorder.Result().StatusCode).To(Equal(http.StatusOK))
					respBody, err := io.ReadAll(recorder.Result().Body)
					Expect(err).NotTo(HaveOccurred())

					var findDiagnosticRequest []model.DiagnosticRequest
					err = json.Unmarshal(respBody, &findDiagnosticRequest)
					Expect(err).NotTo(HaveOccurred())
					Expect(findDiagnosticRequest).NotTo(BeNil())
					Expect(findDiagnosticRequest[0].ID).To(Equal(diagnosticRequest.ID))
					Expect(findDiagnosticRequest[0].Description).To(Equal(diagnosticRequest.Description))
					Expect(vetOrgService.FindVetOrgByIDCallCount()).To(Equal(1))
					Expect(diagnosticRequestService.FindRequestByDateRangeCallCount()).To(Equal(1))
				})
			})

			Context("VetOrg not found in backing storage", func() {

				BeforeEach(func() {
					notFoundError := errors.New("not found")
					vetOrgService.FindVetOrgByIDReturns(nil, notFoundError)
					diagnosticRequestService.FindRequestByDateRangeReturns(nil, notFoundError)
					recorder = httptest.NewRecorder()
					params := rata.Params{
						"vetorg_id":  "12345",
						"start_date": "20190101",
						"end_date":   "20191231",
					}
					request, _ := requestGenerator.CreateRequest(api.DiagnosticRequestsByVetOrgIDAndDateRange, params, nil)
					handler.ServeHTTP(recorder, request)
				})

				It("Returns an error indicating it is unable to find VetOrg", func() {
					Expect(recorder.Result().StatusCode).To(Equal(http.StatusNotFound))
					respBody, err := io.ReadAll(recorder.Result().Body)
					Expect(err).NotTo(HaveOccurred())
					Expect(strings.TrimSpace(string(respBody))).To(Equal(api.ErrorFetchingVetOrg))
					Expect(vetOrgService.FindVetOrgByIDCallCount()).To(Equal(1))
					Expect(diagnosticRequestService.FindRequestByDateRangeCallCount()).To(Equal(0))
				})
			})

			Context("Request(s) not found in backing storage", func() {

				BeforeEach(func() {
					notFoundError := errors.New("not found")
					vetOrgService.FindVetOrgByIDReturns(&vetOrg, nil)
					diagnosticRequestService.FindRequestByDateRangeReturns(nil, notFoundError)
					recorder = httptest.NewRecorder()
					params := rata.Params{
						"vetorg_id":  "12345",
						"start_date": "20190101",
						"end_date":   "20191231",
					}
					request, _ := requestGenerator.CreateRequest(api.DiagnosticRequestsByVetOrgIDAndDateRange, params, nil)
					handler.ServeHTTP(recorder, request)
				})

				It("Returns an error indicating it is unable to find diagnostic request", func() {
					Expect(recorder.Result().StatusCode).To(Equal(http.StatusNotFound))
					respBody, err := io.ReadAll(recorder.Result().Body)
					Expect(err).NotTo(HaveOccurred())
					Expect(strings.TrimSpace(string(respBody))).To(Equal(api.ErrorFetchingDiagnosticRequests))
					Expect(vetOrgService.FindVetOrgByIDCallCount()).To(Equal(1))
					Expect(diagnosticRequestService.FindRequestByDateRangeCallCount()).To(Equal(1))
				})
			})

			Context("No vetorg id provided in the request", func() {

				BeforeEach(func() {
					recorder = httptest.NewRecorder()
					request, _ := http.NewRequest("GET", "/diagnosticrequest/customer/", nil)
					handler.ServeHTTP(recorder, request)
				})

				It("Returns an error indicating it cannot find the page", func() {
					Expect(recorder.Result().StatusCode).To(Equal(http.StatusNotFound))
					Expect(vetOrgService.FindVetOrgByIDCallCount()).To(Equal(0))
					Expect(diagnosticRequestService.FindRequestByDateRangeCallCount()).To(Equal(0))
				})
			})

			Context("Invalid vetorg id provided in the request", func() {

				BeforeEach(func() {
					recorder = httptest.NewRecorder()
					params := rata.Params{
						"vetorg_id":  "one",
						"start_date": "20190101",
						"end_date":   "20191231",
					}
					request, _ := requestGenerator.CreateRequest(api.DiagnosticRequestsByVetOrgIDAndDateRange, params, nil)
					handler.ServeHTTP(recorder, request)
				})

				It("Returns an error indicating it cannot parse the request", func() {
					Expect(recorder.Result().StatusCode).To(Equal(http.StatusBadRequest))
					respBody, err := io.ReadAll(recorder.Result().Body)
					Expect(err).NotTo(HaveOccurred())
					Expect(strings.TrimSpace(string(respBody))).To(Equal(api.UnableToParseParams))
					Expect(vetOrgService.FindVetOrgByIDCallCount()).To(Equal(0))
					Expect(diagnosticRequestService.FindRequestByDateRangeCallCount()).To(Equal(0))
				})
			})

			Context("Invalid start date provided in the request", func() {

				BeforeEach(func() {
					recorder = httptest.NewRecorder()
					params := rata.Params{
						"vetorg_id":  "one",
						"start_date": "20190101",
						"end_date":   "20191231",
					}
					request, _ := requestGenerator.CreateRequest(api.DiagnosticRequestsByVetOrgIDAndDateRange, params, nil)
					handler.ServeHTTP(recorder, request)
				})

				It("Returns an error indicating it cannot parse the request", func() {
					Expect(recorder.Result().StatusCode).To(Equal(http.StatusBadRequest))
					respBody, err := io.ReadAll(recorder.Result().Body)
					Expect(err).NotTo(HaveOccurred())
					Expect(strings.TrimSpace(string(respBody))).To(Equal(api.UnableToParseParams))
					Expect(vetOrgService.FindVetOrgByIDCallCount()).To(Equal(0))
					Expect(diagnosticRequestService.FindRequestByDateRangeCallCount()).To(Equal(0))
				})
			})

			Context("Invalid end date provided in the request", func() {

				BeforeEach(func() {
					recorder = httptest.NewRecorder()
					params := rata.Params{
						"vetorg_id":  "one",
						"start_date": "20190101",
						"end_date":   "20191231",
					}
					request, _ := requestGenerator.CreateRequest(api.DiagnosticRequestsByVetOrgIDAndDateRange, params, nil)
					handler.ServeHTTP(recorder, request)
				})

				It("Returns an error indicating it cannot parse the request", func() {
					Expect(recorder.Result().StatusCode).To(Equal(http.StatusBadRequest))
					respBody, err := io.ReadAll(recorder.Result().Body)
					Expect(err).NotTo(HaveOccurred())
					Expect(strings.TrimSpace(string(respBody))).To(Equal(api.UnableToParseParams))
					Expect(vetOrgService.FindVetOrgByIDCallCount()).To(Equal(0))
					Expect(diagnosticRequestService.FindRequestByDateRangeCallCount()).To(Equal(0))
				})
			})
		})
	})
})
