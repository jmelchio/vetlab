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
			goDiagnosticReport = DiagnosticReport{
				ReportID:   "some-report-id",
				RequestID:  "some-request-id",
				OrgID:      "some-org-id",
				CustomerID: "some-customer-id",
				UserID:     "some-user-id",
				Date:       time.Time{},
				ReportBody: "some-report-body",
				ReportFile: "some-report-file",
			}
			jsonDiagnosticReport = `{"report_id":"some-report-id","request_id":"some-request-id","org_id":"some-org-id","customer_id":"some-customer-id","user_id":"some-user-id","date":"0001-01-01T00:00:00Z","report_body":"some-report-body","report_file":"some-report-file"}`
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
