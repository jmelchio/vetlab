package api_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/tedsuo/rata"

	. "github.com/jmelchio/vetlab/api"
	"github.com/jmelchio/vetlab/api/apifakes"
	"github.com/jmelchio/vetlab/model"
)

var _ = Describe("DiagnosticRequestHandler", func() {
	var (
		handler                  http.Handler
		recorder                 *httptest.ResponseRecorder
		err                      error
		diagnosticRequestService *apifakes.FakeDiagnosticRequestService
		requestGenerator         *rata.RequestGenerator
	)

	BeforeEach(func() {
		diagnosticRequestService = new(apifakes.FakeDiagnosticRequestService)
		handler, err = NewDiagnosticRequestHandler(diagnosticRequestService)
		requestGenerator = rata.NewRequestGenerator("", DiagnosticRequestRoutes)
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
				request, _ := requestGenerator.CreateRequest(SubmitDiagnosticRequest, nil, bytes.NewReader(requestBytes))
				handler.ServeHTTP(recorder, request)
			})

			It("Creates a diagnosticRequest and returns 201 status code", func() {
				Expect(recorder.Result().StatusCode).To(Equal(http.StatusCreated))
				respBody, err := ioutil.ReadAll(recorder.Result().Body)
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
				diagnosticRequestService.SubmitDiagnosticRequestReturns(nil, errors.New("Whoot?"))
				userBytes, err := json.Marshal(diagnosticRequest)
				Expect(err).NotTo(HaveOccurred())
				recorder = httptest.NewRecorder()
				request, _ := requestGenerator.CreateRequest(SubmitDiagnosticRequest, nil, bytes.NewReader(userBytes))
				handler.ServeHTTP(recorder, request)
			})

			It("Fails to return a users and returns 500 status code", func() {
				Expect(recorder.Result().StatusCode).To(Equal(http.StatusInternalServerError))
				respBody, err := ioutil.ReadAll(recorder.Result().Body)
				Expect(err).NotTo(HaveOccurred())
				Expect(err).NotTo(HaveOccurred())
				Expect(string(respBody[0 : len(respBody)-1])).To(Equal(UnableToSubmitDiagnosticRequest))
				Expect(diagnosticRequestService.SubmitDiagnosticRequestCallCount()).To(Equal(1))
			})
		})

		Context("Body of the request is empty", func() {

			BeforeEach(func() {
				recorder = httptest.NewRecorder()
				request, _ := requestGenerator.CreateRequest(SubmitDiagnosticRequest, nil, nil)
				handler.ServeHTTP(recorder, request)
			})

			It("Fails to create a user and returns a 400 status code", func() {
				Expect(recorder.Result().StatusCode).To(Equal(http.StatusBadRequest))
				respBody, err := ioutil.ReadAll(recorder.Result().Body)
				Expect(err).NotTo(HaveOccurred())
				Expect(string(respBody[0 : len(respBody)-1])).To(Equal(EmptyBody))
				Expect(diagnosticRequestService.SubmitDiagnosticRequestCallCount()).To(Equal(0))
			})
		})

		Context("Body of the request contains invalid data", func() {

			BeforeEach(func() {
				userBytes, err := json.Marshal("diagnosticRequest")
				Expect(err).NotTo(HaveOccurred())
				recorder = httptest.NewRecorder()
				request, _ := requestGenerator.CreateRequest(SubmitDiagnosticRequest, nil, bytes.NewReader(userBytes))
				handler.ServeHTTP(recorder, request)
			})

			It("Fails to create a user and returns a 400 status code", func() {
				Expect(recorder.Result().StatusCode).To(Equal(http.StatusBadRequest))
				respBody, err := ioutil.ReadAll(recorder.Result().Body)
				Expect(err).NotTo(HaveOccurred())
				Expect(string(respBody[0 : len(respBody)-1])).To(Equal(InvalidBody))
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
				ID:          uint(98765),
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
					request, _ := requestGenerator.CreateRequest(DiagnosticRequestByID, params, nil)
					handler.ServeHTTP(recorder, request)
				})

				It("Returns the requested diagnostic request information", func() {
					Expect(recorder.Result().StatusCode).To(Equal(http.StatusOK))
					respBody, err := ioutil.ReadAll(recorder.Result().Body)
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
					notFoundError := errors.New("Not found")
					diagnosticRequestService.FindRequestByIDReturns(nil, notFoundError)
					recorder = httptest.NewRecorder()
					params := rata.Params{
						"request_id": "12345",
					}
					request, _ := requestGenerator.CreateRequest(DiagnosticRequestByID, params, nil)
					handler.ServeHTTP(recorder, request)
				})

				It("Returns an error indicating it is unable to find diagnostic request", func() {
					Expect(recorder.Result().StatusCode).To(Equal(http.StatusNotFound))
					respBody, err := ioutil.ReadAll(recorder.Result().Body)
					Expect(err).NotTo(HaveOccurred())
					Expect(strings.TrimSpace(string(respBody))).To(Equal(ErrorFetchingDiagnosticRequests))
					Expect(diagnosticRequestService.FindRequestByIDCallCount()).To(Equal(1))
				})
			})

			Context("No request id provided in the reqest", func() {

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

			Context("Invalid request id provided in the reqest", func() {

				BeforeEach(func() {
					recorder = httptest.NewRecorder()
					request, _ := http.NewRequest("GET", "/diagnosticrequest/one", nil)
					handler.ServeHTTP(recorder, request)
				})

				It("Returns an error indicating it cannot parse the request", func() {
					Expect(recorder.Result().StatusCode).To(Equal(http.StatusBadRequest))
					respBody, err := ioutil.ReadAll(recorder.Result().Body)
					Expect(err).NotTo(HaveOccurred())
					Expect(strings.TrimSpace(string(respBody))).To(Equal(UnableToParseParams))
					Expect(diagnosticRequestService.FindRequestByIDCallCount()).To(Equal(0))
				})
			})
		})
	})
})
