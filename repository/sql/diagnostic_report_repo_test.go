package sql_test

import (
	"time"

	"github.com/jmelchio/vetlab/model"
	"github.com/jmelchio/vetlab/repository/sql"
	"github.com/jmelchio/vetlab/service"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("DiagnosticReportRepo", func() {

	var (
		diagnosticReportRepo service.DiagnosticReportRepo
		diagnosticReportOne  model.DiagnosticReport
		userID               uint
		vetOrgID             uint
		customerID           uint
		requestID            uint
		date                 time.Time
		reportBody           string
		reportFile           string
	)

	BeforeEach(func() {
		diagnosticReportRepoImpl := sql.DiagnosticReportRepo{Database: database}
		diagnosticReportRepo = &diagnosticReportRepoImpl

		userID = 12345
		vetOrgID = 12345
		customerID = 12345
		requestID = 12345
		date = time.Now()
		reportBody = "the body of the report"
		reportFile = "filename.txt"

		diagnosticReportOne = model.DiagnosticReport{
			RequestID:  requestID,
			VetOrgID:   vetOrgID,
			CustomerID: customerID,
			UserID:     userID,
			Date:       &date,
			ReportBody: reportBody,
			ReportFile: reportFile,
		}
	})

	AfterEach(func() {
		err = database.Where("1 = 1").Delete(&model.DiagnosticReport{}).Error
		Expect(err).NotTo(HaveOccurred())
	})

	Describe("Diagnostic Report table", func() {

		Context("Diagnostic report table has been created during in BeforeSuite", func() {

			It("Has a diagnostic report table", func() {
				hasDiagnosticReportTable := database.Migrator().HasTable(&model.DiagnosticReport{})
				Expect(hasDiagnosticReportTable).To(BeTrue())
			})
		})
	})

	Describe("Create a diagnosticReport", func() {

		BeforeEach(func() {
			Expect(diagnosticReportOne.ID).To(Equal(uint(0)))
		})

		It("Creates a new diagnosticReport record", func() {
			err = diagnosticReportRepo.Create(&diagnosticReportOne)
			Expect(err).NotTo(HaveOccurred())
			Expect(diagnosticReportOne.ID).NotTo(Equal(uint(0)))
		})
	})

	Describe("Update a diagnosticReport", func() {

		Context("When a diagnosticReport is found", func() {

			BeforeEach(func() {
				err = diagnosticReportRepo.Create(&diagnosticReportOne)
				Expect(err).NotTo(HaveOccurred())
				Expect(diagnosticReportOne.ID).NotTo(Equal(uint(0)))
			})

			Context("When the diagnosticReport exists", func() {

				It("It updates the diagnosticReport record and returns updated diagnosticReport", func() {
					diagnosticReportOne.ReportFile = "new_diagnosticReport_filename.txt"
					diagnosticReportOne.ReportBody = "look at this fancy new body of mine"
					err = diagnosticReportRepo.Update(&diagnosticReportOne)
					Expect(err).NotTo(HaveOccurred())
					diagnosticReportFound, err := diagnosticReportRepo.GetByID(diagnosticReportOne.ID)
					Expect(err).NotTo(HaveOccurred())
					Expect(diagnosticReportFound.ReportFile).To(Equal(diagnosticReportOne.ReportFile))
					Expect(diagnosticReportFound.ReportBody).To(Equal(diagnosticReportOne.ReportBody))
				})
			})
		})

		Context("When the diagnosticReport does not exist", func() {

			BeforeEach(func() {
			})

			It("Returns an error and nil for the diagnosticReport", func() {
				err = diagnosticReportRepo.Update(&diagnosticReportOne)
				Expect(err).To(HaveOccurred())
			})
		})
	})

	Describe("Delete a diagnosticReport", func() {

		Context("When the diagnosticReport exists", func() {

			BeforeEach(func() {
				err = diagnosticReportRepo.Create(&diagnosticReportOne)
				Expect(err).NotTo(HaveOccurred())
				Expect(diagnosticReportOne.ID).NotTo(Equal(uint(0)))
			})

			It("Deletes the record and returns no error", func() {
				err = diagnosticReportRepo.Delete(&diagnosticReportOne)
				Expect(err).NotTo(HaveOccurred())
				var foundDiagnosticReport *model.DiagnosticReport
				foundDiagnosticReport, err = diagnosticReportRepo.GetByID(diagnosticReportOne.ID)
				Expect(err).To(HaveOccurred())
				Expect(foundDiagnosticReport).To(BeNil())
			})
		})
	})

	Describe("Get a diagnosticReport by ID", func() {

		Context("When the diagnosticReport is found", func() {

			var foundDiagnosticReport *model.DiagnosticReport

			BeforeEach(func() {
				err = diagnosticReportRepo.Create(&diagnosticReportOne)
				Expect(err).NotTo(HaveOccurred())
				Expect(diagnosticReportOne.ID).NotTo(Equal(uint(0)))
			})

			It("It returns the diagnosticReport and nil for error", func() {
				foundDiagnosticReport, err = diagnosticReportRepo.GetByID(diagnosticReportOne.ID)
				Expect(err).NotTo(HaveOccurred())
				Expect(foundDiagnosticReport).NotTo(BeNil())
				Expect(foundDiagnosticReport.ReportBody).To(Equal(diagnosticReportOne.ReportBody))
			})
		})

		Context("When the diagnosticReport is not found", func() {

			var foundDiagnosticReport *model.DiagnosticReport

			It("It returns and error and nil for the diagnosticReport", func() {
				foundDiagnosticReport, err = diagnosticReportRepo.GetByID(uint(10))
				Expect(err).To(HaveOccurred())
				Expect(foundDiagnosticReport).To(BeNil())
			})
		})
	})

	Describe("Get diagnosticReports by VetOrgID", func() {

		Context("When diagnosticReport(s) are found", func() {

			var foundDiagnosticReports []model.DiagnosticReport

			BeforeEach(func() {
				err = diagnosticReportRepo.Create(&diagnosticReportOne)
				Expect(err).NotTo(HaveOccurred())
				Expect(diagnosticReportOne.ID).NotTo(Equal(uint(0)))
			})

			It("It returns an array with one diagnosticReport and nil for error", func() {
				foundDiagnosticReports, err = diagnosticReportRepo.GetByVetOrgID(diagnosticReportOne.VetOrgID)
				Expect(err).NotTo(HaveOccurred())
				Expect(foundDiagnosticReports).NotTo(BeNil())
				Expect(foundDiagnosticReports).To(HaveLen(1))
				Expect(foundDiagnosticReports[0].ID).To(Equal(diagnosticReportOne.ID))
				Expect(foundDiagnosticReports[0].ReportBody).To(Equal(diagnosticReportOne.ReportBody))
			})
		})

		Context("When no diagnosticReports are found", func() {

			var foundDiagnosticReports []model.DiagnosticReport

			BeforeEach(func() {
			})

			It("It returns no diagnosticReports and an error", func() {
				foundDiagnosticReports, err = diagnosticReportRepo.GetByVetOrgID(19)
				Expect(err).To(HaveOccurred())
				Expect(foundDiagnosticReports).To(BeNil())
			})
		})
	})

	Describe("Get diagnosticReports by UserID", func() {

		Context("When diagnosticReport(s) are found", func() {

			var foundDiagnosticReports []model.DiagnosticReport

			BeforeEach(func() {
				err = diagnosticReportRepo.Create(&diagnosticReportOne)
				Expect(err).NotTo(HaveOccurred())
				Expect(diagnosticReportOne.ID).NotTo(Equal(uint(0)))
			})

			It("It returns an array with one diagnosticReport and nil for error", func() {
				foundDiagnosticReports, err = diagnosticReportRepo.GetByUserID(diagnosticReportOne.UserID)
				Expect(err).NotTo(HaveOccurred())
				Expect(foundDiagnosticReports).NotTo(BeNil())
				Expect(foundDiagnosticReports).To(HaveLen(1))
				Expect(foundDiagnosticReports[0].ID).To(Equal(diagnosticReportOne.ID))
				Expect(foundDiagnosticReports[0].ReportBody).To(Equal(diagnosticReportOne.ReportBody))
			})
		})

		Context("When no diagnosticReports are found", func() {

			var foundDiagnosticReports []model.DiagnosticReport

			BeforeEach(func() {
			})

			It("It returns no diagnosticReports and an error", func() {
				foundDiagnosticReports, err = diagnosticReportRepo.GetByUserID(19)
				Expect(err).To(HaveOccurred())
				Expect(foundDiagnosticReports).To(BeNil())
			})
		})
	})

	Describe("Get diagnosticReports by CustomerID", func() {

		Context("When diagnosticReport(s) are found", func() {

			var foundDiagnosticReports []model.DiagnosticReport

			BeforeEach(func() {
				err = diagnosticReportRepo.Create(&diagnosticReportOne)
				Expect(err).NotTo(HaveOccurred())
				Expect(diagnosticReportOne.ID).NotTo(Equal(uint(0)))
			})

			It("It returns an array with one diagnosticReport and nil for error", func() {
				foundDiagnosticReports, err = diagnosticReportRepo.GetByCustomerID(diagnosticReportOne.CustomerID)
				Expect(err).NotTo(HaveOccurred())
				Expect(foundDiagnosticReports).NotTo(BeNil())
				Expect(foundDiagnosticReports).To(HaveLen(1))
				Expect(foundDiagnosticReports[0].ID).To(Equal(diagnosticReportOne.ID))
				Expect(foundDiagnosticReports[0].ReportBody).To(Equal(diagnosticReportOne.ReportBody))
			})
		})

		Context("When no diagnosticReports are found", func() {

			var foundDiagnosticReports []model.DiagnosticReport

			BeforeEach(func() {
			})

			It("It returns no diagnosticReports and an error", func() {
				foundDiagnosticReports, err = diagnosticReportRepo.GetByCustomerID(19)
				Expect(err).To(HaveOccurred())
				Expect(foundDiagnosticReports).To(BeNil())
			})
		})
	})
})
