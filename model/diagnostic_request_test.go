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
		var (
			goDiagnosticRequest   DiagnosticRequest
			jsonDiagnosticRequest string
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
			jsonDiagnosticRequest = `{"request_id":"some-request-id","org_id":"some-org-id","customer_id":"some-customer-id","user_id":"some-user-id","date":"0001-01-01T00:00:00Z","description":"some-description"}`
		})
		Context("From Golang to Json", func() {
			It("Transforms without errors", func() {
				diagnosticRequestBytes, err := json.Marshal(goDiagnosticRequest)
				Expect(err).NotTo(HaveOccurred())

				jsonResult := string(diagnosticRequestBytes)
				Expect(jsonResult).To(Equal(jsonDiagnosticRequest))
			})
		})
		Context("From Json to Golang", func() {
			It("Transforms without errors", func() {
				var unmarshalDiagnosticRequest DiagnosticRequest
				err := json.Unmarshal([]byte(jsonDiagnosticRequest), &unmarshalDiagnosticRequest)
				Expect(err).NotTo(HaveOccurred())

				Expect(unmarshalDiagnosticRequest).To(Equal(goDiagnosticRequest))
			})
		})
	})
})
