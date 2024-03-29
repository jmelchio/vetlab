package model_test

import (
	"encoding/json"
	"time"

	. "github.com/jmelchio/vetlab/model"
	"gorm.io/gorm"

	. "github.com/onsi/ginkgo/v2"
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
				Model: gorm.Model{
					ID:        uint(12345),
					CreatedAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
					UpdatedAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
					DeletedAt: gorm.DeletedAt{},
				},
				VetOrgID:    uint(12345),
				CustomerID:  uint(12345),
				UserID:      uint(12345),
				Date:        &time.Time{},
				Description: description,
			}
			jsonDiagnosticRequest = `{"ID":12345,"CreatedAt":"2020-01-01T00:00:00Z","UpdatedAt":"2020-01-01T00:00:00Z","DeletedAt":null,"vet_org_id":12345,"customer_id":12345,"user_id":12345,"date":"0001-01-01T00:00:00Z","description":"some-description"}`
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
