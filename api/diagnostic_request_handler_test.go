package api_test

import (
	"bytes"
	"encoding/json"
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
			diagnosticRequest = model.DiagnosticRequest{}
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
		})
	})
})
