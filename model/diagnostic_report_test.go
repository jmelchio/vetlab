package model_test

import (
	"encoding/json"
	"time"

	. "github.com/jmelchio/vetlab/model"
	"gorm.io/gorm"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("DiagnosticReport", func() {

	Describe("Diagnostic report can be transformed from and to Json", func() {
		var (
			goDiagnosticReport   DiagnosticReport
			jsonDiagnosticReport string
		)

		BeforeEach(func() {
			reportBody := "some-report-body"
			reportFile := "some-report-file"
			goDiagnosticReport = DiagnosticReport{
				Model: gorm.Model{
					ID:        uint(12345),
					CreatedAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
					UpdatedAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
					DeletedAt: gorm.DeletedAt{},
				},
				RequestID:  uint(12345),
				VetOrgID:   uint(12345),
				CustomerID: uint(12345),
				UserID:     uint(12345),
				Date:       &time.Time{},
				ReportBody: reportBody,
				ReportFile: reportFile,
			}
			jsonDiagnosticReport = `{"ID":12345,"CreatedAt":"2020-01-01T00:00:00Z","UpdatedAt":"2020-01-01T00:00:00Z","DeletedAt":null,"request_id":12345,"vet_org_id":12345,"customer_id":12345,"user_id":12345,"date":"0001-01-01T00:00:00Z","report_body":"some-report-body","report_file":"some-report-file"}`
		})

		Context("From Golang to Json", func() {

			It("Transforms without errors", func() {
				diagnosticReportBytes, err := json.Marshal(goDiagnosticReport)
				Expect(err).NotTo(HaveOccurred())

				jsonResult := string(diagnosticReportBytes)
				Expect(jsonResult).To(Equal(jsonDiagnosticReport))
			})
		})

		Context("From Json to Golang", func() {

			It("Transforms without errors", func() {
				var unmarshalDiagnosticReport DiagnosticReport
				err := json.Unmarshal([]byte(jsonDiagnosticReport), &unmarshalDiagnosticReport)
				Expect(err).NotTo(HaveOccurred())

				Expect(unmarshalDiagnosticReport).To(Equal(goDiagnosticReport))
			})
		})
	})
})
