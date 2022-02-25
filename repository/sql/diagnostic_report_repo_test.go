package sql_test

import (
	"github.com/jmelchio/vetlab/model"
	"github.com/jmelchio/vetlab/repository/sql"
	"github.com/jmelchio/vetlab/service"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("DiagnosticReportRepo", func() {

	var (
		diagnosticReportRepo service.DiagnosticReportRepo
		diagnosticReportOne  model.DiagnosticReport
		diagnosticReportTwo  model.DiagnosticReport
		diagnosticReportID   uint
		userID               uint
		vetOrgID             uint
		customerID           uint
	)

	BeforeEach(func() {
		diagnosticReportRepoImpl := sql.DiagnosticReportRepo{Database: database}
		diagnosticReportRepo = &diagnosticReportRepoImpl

		diagnosticReportID = 12345
		userID = 12345
		vetOrgID = 12345
		customerID = 12345

		diagnosticReportOne = model.DiagnosticReport{}
		diagnosticReportTwo = model.DiagnosticReport{}
	})

	AfterEach(func() {
		err = database.Where("1 = 1").Delete(&model.DiagnosticReport{}).Error
		Expect(err).NotTo(HaveOccurred())
	})

})
