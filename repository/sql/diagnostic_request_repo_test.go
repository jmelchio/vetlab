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

	Describe("Delete a diagnosticRequest", func() {

		Context("When the diagnosticRequest exists", func() {

			BeforeEach(func() {
				err = diagnosticRequestRepo.Create(&diagnosticRequestOne)
				Expect(err).NotTo(HaveOccurred())
				Expect(diagnosticRequestOne.ID).NotTo(Equal(uint(0)))
			})

			It("Deletes the record and returns no error", func() {
				err = diagnosticRequestRepo.Delete(&diagnosticRequestOne)
				Expect(err).NotTo(HaveOccurred())
				var foundDiagnosticRequest *model.DiagnosticRequest
				foundDiagnosticRequest, err = diagnosticRequestRepo.GetByID(diagnosticRequestOne.ID)
				Expect(err).To(HaveOccurred())
				Expect(foundDiagnosticRequest).To(BeNil())
			})
		})
	})

	Describe("Get a diagnosticRequest by ID", func() {

		Context("When the diagnosticRequest is found", func() {

			var foundDiagnosticRequest *model.DiagnosticRequest

			BeforeEach(func() {
				err = diagnosticRequestRepo.Create(&diagnosticRequestOne)
				Expect(err).NotTo(HaveOccurred())
				Expect(diagnosticRequestOne.ID).NotTo(Equal(uint(0)))
			})

			It("It returns the diagnosticRequest and nil for error", func() {
				foundDiagnosticRequest, err = diagnosticRequestRepo.GetByID(diagnosticRequestOne.ID)
				Expect(err).NotTo(HaveOccurred())
				Expect(foundDiagnosticRequest).NotTo(BeNil())
				Expect(foundDiagnosticRequest.Description).To(Equal(diagnosticRequestOne.Description))
			})
		})

		Context("When the diagnosticRequest is not found", func() {

			var foundDiagnosticRequest *model.DiagnosticRequest

			It("It returns and error and nil for the diagnosticRequest", func() {
				foundDiagnosticRequest, err = diagnosticRequestRepo.GetByID(uint(10))
				Expect(err).To(HaveOccurred())
				Expect(foundDiagnosticRequest).To(BeNil())
			})
		})
	})

	Describe("Get diagnosticRequests by VetOrgID", func() {

		Context("When diagnosticRequest(s) are found", func() {

			var foundDiagnosticRequests []model.DiagnosticRequest

			BeforeEach(func() {
				err = diagnosticRequestRepo.Create(&diagnosticRequestOne)
				Expect(err).NotTo(HaveOccurred())
				Expect(diagnosticRequestOne.ID).NotTo(Equal(uint(0)))
			})

			It("It returns an array with one diagnosticRequest and nil for error", func() {
				foundDiagnosticRequests, err = diagnosticRequestRepo.GetByVetOrgID(diagnosticRequestOne.VetOrgID)
				Expect(err).NotTo(HaveOccurred())
				Expect(foundDiagnosticRequests).NotTo(BeNil())
				Expect(foundDiagnosticRequests).To(HaveLen(1))
				Expect(foundDiagnosticRequests[0].ID).To(Equal(diagnosticRequestOne.ID))
				Expect(foundDiagnosticRequests[0].Description).To(Equal(diagnosticRequestOne.Description))
			})
		})

		Context("When no diagnosticRequests are found", func() {

			var foundDiagnosticRequests []model.DiagnosticRequest

			BeforeEach(func() {
			})

			It("It returns no diagnosticRequests and an error", func() {
				foundDiagnosticRequests, err = diagnosticRequestRepo.GetByVetOrgID(19)
				Expect(err).To(HaveOccurred())
				Expect(foundDiagnosticRequests).To(BeNil())
			})
		})
	})

	Describe("Get diagnosticRequests by UserID", func() {

		Context("When diagnosticRequest(s) are found", func() {

			var foundDiagnosticRequests []model.DiagnosticRequest

			BeforeEach(func() {
				err = diagnosticRequestRepo.Create(&diagnosticRequestOne)
				Expect(err).NotTo(HaveOccurred())
				Expect(diagnosticRequestOne.ID).NotTo(Equal(uint(0)))
			})

			It("It returns an array with one diagnosticRequest and nil for error", func() {
				foundDiagnosticRequests, err = diagnosticRequestRepo.GetByUserID(diagnosticRequestOne.UserID)
				Expect(err).NotTo(HaveOccurred())
				Expect(foundDiagnosticRequests).NotTo(BeNil())
				Expect(foundDiagnosticRequests).To(HaveLen(1))
				Expect(foundDiagnosticRequests[0].ID).To(Equal(diagnosticRequestOne.ID))
				Expect(foundDiagnosticRequests[0].Description).To(Equal(diagnosticRequestOne.Description))
			})
		})

		Context("When no diagnosticRequests are found", func() {

			var foundDiagnosticRequests []model.DiagnosticRequest

			BeforeEach(func() {
			})

			It("It returns no diagnosticRequests and an error", func() {
				foundDiagnosticRequests, err = diagnosticRequestRepo.GetByUserID(19)
				Expect(err).To(HaveOccurred())
				Expect(foundDiagnosticRequests).To(BeNil())
			})
		})
	})

	Describe("Get diagnosticRequests by CustomerID", func() {

		Context("When diagnosticRequest(s) are found", func() {

			var foundDiagnosticRequests []model.DiagnosticRequest

			BeforeEach(func() {
				err = diagnosticRequestRepo.Create(&diagnosticRequestOne)
				Expect(err).NotTo(HaveOccurred())
				Expect(diagnosticRequestOne.ID).NotTo(Equal(uint(0)))
			})

			It("It returns an array with one diagnosticRequest and nil for error", func() {
				foundDiagnosticRequests, err = diagnosticRequestRepo.GetByCustomerID(diagnosticRequestOne.CustomerID)
				Expect(err).NotTo(HaveOccurred())
				Expect(foundDiagnosticRequests).NotTo(BeNil())
				Expect(foundDiagnosticRequests).To(HaveLen(1))
				Expect(foundDiagnosticRequests[0].ID).To(Equal(diagnosticRequestOne.ID))
				Expect(foundDiagnosticRequests[0].Description).To(Equal(diagnosticRequestOne.Description))
			})
		})

		Context("When no diagnosticRequests are found", func() {

			var foundDiagnosticRequests []model.DiagnosticRequest

			BeforeEach(func() {
			})

			It("It returns no diagnosticRequests and an error", func() {
				foundDiagnosticRequests, err = diagnosticRequestRepo.GetByCustomerID(19)
				Expect(err).To(HaveOccurred())
				Expect(foundDiagnosticRequests).To(BeNil())
			})
		})
	})
})
