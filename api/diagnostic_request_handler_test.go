package api_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"

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
	})
})
