package model_test

import (
	"encoding/json"
	"time"

	. "github.com/jmelchio/vetlab/model"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("DiagnosticRequest", func() {
	Describe("Diagnostic request can be transformed from and to Json", func() {
		Context("From Golang to Json", func() {
			var (
				goDiagnosticRequest DiagnosticRequest
				expectedJson        string
			)
			BeforeEach(func() {
				goDiagnosticRequest = DiagnosticRequest{
					RequestID:   "some-request-id",
					OrgID:       "some-org-id",
					CustomerID:  "some-customer-id",
					UserID:      "some-user-id",
					Date:        time.Time{},
					Description: "some-description",
				}
				expectedJson = `{"request_id": "some-request-id"}`
			})
			It("Transforms without errors", func() {
				diagnosticRequestBytes, err := json.Marshal(goDiagnosticRequest)
				Expect(err).NotTo(HaveOccurred())

				jsonResult := string(diagnosticRequestBytes)
				Expect(jsonResult).To(Equal(expectedJson))
			})
		})
		Context("From Json to Golang", func() {

		})
	})
})
