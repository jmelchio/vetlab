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
			description := "some-description"
			goDiagnosticRequest = DiagnosticRequest{
				ID:          12345,
				VetOrgID:    12345,
				CustomerID:  12345,
				UserID:      12345,
				Date:        &time.Time{},
				Description: description,
			}
			jsonDiagnosticRequest = `{"id":12345,"vet_org_id":12345,"customer_id":12345,"user_id":12345,"date":"0001-01-01T00:00:00Z","description":"some-description"}`
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
