package sql_test

import (
	"time"

	"github.com/jmelchio/vetlab/model"
	"github.com/jmelchio/vetlab/repository/sql"
	"github.com/jmelchio/vetlab/service"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("DiagnosticRequestRepo", func() {

	var (
		diagnosticRequestRepo service.DiagnosticRequestRepo
		diagnosticRequestOne  model.DiagnosticRequest
		userID                uint
		vetOrgID              uint
		customerID            uint
		description           string
		date                  time.Time
	)

	BeforeEach(func() {
		diagnosticRequestRepoImpl := sql.DiagnosticRequestRepo{Database: database}
		diagnosticRequestRepo = &diagnosticRequestRepoImpl

		userID = 12345
		vetOrgID = 12345
		customerID = 12345
		description = "request description"
		date = time.Now()

		diagnosticRequestOne = model.DiagnosticRequest{
			VetOrgID:    vetOrgID,
			CustomerID:  customerID,
			UserID:      userID,
			Date:        &date,
			Description: description,
		}
	})

	AfterEach(func() {
		err = database.Where("1 = 1").Delete(&model.DiagnosticRequest{}).Error
		Expect(err).NotTo(HaveOccurred())
	})

	Describe("Diagnostic Request table", func() {

		Context("Diagnostic request table has been created during in BeforeSuite", func() {

			It("Has a diagnostic request table", func() {
				hasDiagnosticRequestTable := database.Migrator().HasTable(&model.DiagnosticRequest{})
				Expect(hasDiagnosticRequestTable).To(BeTrue())
			})
		})
	})

	Describe("Create a diagnosticRequest", func() {

		BeforeEach(func() {
			Expect(diagnosticRequestOne.ID).To(Equal(uint(0)))
		})

		It("Creates a new diagnosticRequest record", func() {
			err = diagnosticRequestRepo.Create(&diagnosticRequestOne)
			Expect(err).NotTo(HaveOccurred())
			Expect(diagnosticRequestOne.ID).NotTo(Equal(uint(0)))
		})
	})

	Describe("Update a diagnosticRequest", func() {

		Context("When a diagnosticRequest is found", func() {

			BeforeEach(func() {
				err = diagnosticRequestRepo.Create(&diagnosticRequestOne)
				Expect(err).NotTo(HaveOccurred())
				Expect(diagnosticRequestOne.ID).NotTo(Equal(uint(0)))
			})

			Context("When the diagnosticRequest exists", func() {

				It("It updates the diagnosticRequest record and returns updated diagnosticRequest", func() {
					diagnosticRequestOne.Description = "new request description"
					err = diagnosticRequestRepo.Update(&diagnosticRequestOne)
					Expect(err).NotTo(HaveOccurred())
					diagnosticRequestFound, err := diagnosticRequestRepo.GetByID(diagnosticRequestOne.ID)
					Expect(err).NotTo(HaveOccurred())
					Expect(diagnosticRequestFound.Description).To(Equal(diagnosticRequestOne.Description))
				})
			})
		})

		Context("When the diagnosticRequest does not exist", func() {

			BeforeEach(func() {
			})

			It("Returns an error and nil for the diagnosticRequest", func() {
				err = diagnosticRequestRepo.Update(&diagnosticRequestOne)
				Expect(err).To(HaveOccurred())
			})
		})
	})
})
