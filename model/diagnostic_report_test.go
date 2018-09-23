package model_test

import (
	"encoding/json"
	"time"

	. "github.com/jmelchio/vetlab/model"

	. "github.com/onsi/ginkgo"
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
				ID:         uint(12345),
				RequestID:  uint(12345),
				VetOrgID:   uint(12345),
				CustomerID: uint(12345),
				UserID:     uint(12345),
				Date:       &time.Time{},
				ReportBody: &reportBody,
				ReportFile: &reportFile,
			}
			jsonDiagnosticReport = `{"id":12345,"request_id":12345,"vet_org_id":12345,"customer_id":12345,"user_id":12345,"date":"0001-01-01T00:00:00Z","report_body":"some-report-body","report_file":"some-report-file"}`
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
