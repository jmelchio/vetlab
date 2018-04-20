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
		Context("From Golang to Json", func() {
			var (
				goDiagnosticReport DiagnosticReport
				expectedJson       string
			)
			BeforeEach(func() {
				goDiagnosticReport = DiagnosticReport{
					ReportID:   "some-report-id",
					OrgID:      "some-org-id",
					CustomerID: "some-customer-id",
					UserID:     "some-user-id",
					Date:       time.Time{},
					ReportBody: "some-report-body",
					ReportFile: "some-report-file",
				}
				expectedJson = `{"report_id": "some-report-id"}`
			})
			It("Transforms without errors", func() {
				diagnosticReportBytes, err := json.Marshal(goDiagnosticReport)
				Expect(err).NotTo(HaveOccurred())

				jsonResult := string(diagnosticReportBytes)
				Expect(jsonResult).To(Equal(expectedJson))
			})
		})
		Context("From Json to Golang", func() {

		})
	})
})
