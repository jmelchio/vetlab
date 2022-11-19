package sql_test

import (
	"github.com/jmelchio/vetlab/model"
	"github.com/jmelchio/vetlab/repository/sql"
	"github.com/jmelchio/vetlab/service"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("VetOrgRepo", func() {

	var (
		vetOrgRepo         service.VetOrgRepo
		vetOrgOne          model.VetOrg
		vetOrgTwo          model.VetOrg
		orgName            string
		street             string
		houseNumber        string
		city               string
		province           string
		country            string
		postalCode         string
		email              string
		phone              string
		fax                string
		customers          []model.Customer
		diagnosticReports  []model.DiagnosticReport
		diagnosticRequests []model.DiagnosticRequest
	)

	BeforeEach(func() {
		vetOrgRepoImpl := sql.VetOrgRepo{Database: database}
		vetOrgRepo = &vetOrgRepoImpl

		orgName = "zeVetOrg"
		street = "zeStreet"
		houseNumber = "909-1"
		city = "pleasantville"
		province = "province"
		country = "neverland"
		postalCode = "90210"
		email = "someone@someplace.com"
		phone = "101-101-1010"
		fax = "202-202-2020"
		customers = nil
		diagnosticReports = nil
		diagnosticRequests = nil

		vetOrgOne = model.VetOrg{
			OrgName:            &orgName,
			Street:             street,
			HouseNumber:        houseNumber,
			City:               city,
			Province:           province,
			Country:            country,
			PostalCode:         postalCode,
			Email:              email,
			Phone:              phone,
			Fax:                fax,
			Customers:          customers,
			DiagnosticReports:  diagnosticReports,
			DiagnosticRequests: diagnosticRequests,
		}
	})

	AfterEach(func() {
		err = database.Where("1 = 1").Delete(&model.VetOrg{}).Error
		Expect(err).NotTo(HaveOccurred())
	})

	Describe("VetOrg table", func() {

		Context("VetOrg table has been created during in BeforeSuite", func() {

			It("Has a customer table", func() {
				hasVetOrgTable := database.Migrator().HasTable(&model.VetOrg{})
				Expect(hasVetOrgTable).To(BeTrue())
			})
		})
	})

	Describe("Create a vetOrg", func() {

		Context("When a vetOrgname is not taken yet", func() {
			BeforeEach(func() {
				Expect(vetOrgOne.ID).To(Equal(uint(0)))
			})

			It("Creates a new vetOrg record", func() {
				err = vetOrgRepo.Create(&vetOrgOne)
				Expect(err).NotTo(HaveOccurred())
				Expect(vetOrgOne.ID).NotTo(Equal(uint(0)))
			})
		})

		Context("When a vetOrgname is taken already", func() {

			BeforeEach(func() {
				vetOrgTwo = vetOrgOne
			})

			It("It returns an error", func() {
				err = vetOrgRepo.Create(&vetOrgOne)
				Expect(err).NotTo(HaveOccurred())
				err = vetOrgRepo.Create(&vetOrgTwo)
				Expect(err).To(HaveOccurred())
			})
		})
	})

	Describe("Update a vetOrg", func() {

		Context("When a vetOrg is found", func() {

			BeforeEach(func() {
				err = vetOrgRepo.Create(&vetOrgOne)
				Expect(err).NotTo(HaveOccurred())
				Expect(vetOrgOne.ID).NotTo(Equal(uint(0)))
			})

			Context("When the vetOrg exists", func() {

				It("It updates the vetOrg record and returns updated vetOrg", func() {
					vetOrgOne.City = "weupdatedthiscity"
					vetOrgOne.Country = "weupdatedthiscountry"
					err = vetOrgRepo.Update(&vetOrgOne)
					Expect(err).NotTo(HaveOccurred())
					vetOrgFound, err := vetOrgRepo.GetByID(vetOrgOne.ID)
					Expect(err).NotTo(HaveOccurred())
					Expect(vetOrgFound.City).To(Equal(vetOrgOne.City))
					Expect(vetOrgFound.Country).To(Equal(vetOrgOne.Country))
				})
			})
		})

		Context("When the vetOrg does not exist", func() {

			BeforeEach(func() {
			})

			It("Returns an error and nil for the vetOrg", func() {
				err = vetOrgRepo.Update(&vetOrgOne)
				Expect(err).To(HaveOccurred())
			})
		})
	})

	Describe("Delete a vetOrg", func() {

		Context("When the vetOrg exists", func() {

			BeforeEach(func() {
				err = vetOrgRepo.Create(&vetOrgOne)
				Expect(err).NotTo(HaveOccurred())
				Expect(vetOrgOne.ID).NotTo(Equal(uint(0)))
			})

			It("Deletes the record and returns no error", func() {
				err = vetOrgRepo.Delete(&vetOrgOne)
				Expect(err).NotTo(HaveOccurred())
				var foundDiagnosticReport *model.VetOrg
				foundDiagnosticReport, err = vetOrgRepo.GetByID(vetOrgOne.ID)
				Expect(err).To(HaveOccurred())
				Expect(foundDiagnosticReport).To(BeNil())
			})
		})
	})

	Describe("Get a vetOrg by ID", func() {

		Context("When the vetOrg is found", func() {

			var foundDiagnosticReport *model.VetOrg

			BeforeEach(func() {
				err = vetOrgRepo.Create(&vetOrgOne)
				Expect(err).NotTo(HaveOccurred())
				Expect(vetOrgOne.ID).NotTo(Equal(uint(0)))
			})

			It("It returns the vetOrg and nil for error", func() {
				foundDiagnosticReport, err = vetOrgRepo.GetByID(vetOrgOne.ID)
				Expect(err).NotTo(HaveOccurred())
				Expect(foundDiagnosticReport).NotTo(BeNil())
				Expect(foundDiagnosticReport.City).To(Equal(vetOrgOne.City))
			})
		})

		Context("When the vetOrg is not found", func() {

			var foundDiagnosticReport *model.VetOrg

			It("It returns and error and nil for the vetOrg", func() {
				foundDiagnosticReport, err = vetOrgRepo.GetByID(uint(10))
				Expect(err).To(HaveOccurred())
				Expect(foundDiagnosticReport).To(BeNil())
			})
		})
	})

	Describe("Get a vetOrg by Name", func() {

		Context("When the vetOrg is found", func() {

			var foundVetOrg []model.VetOrg

			BeforeEach(func() {
				err = vetOrgRepo.Create(&vetOrgOne)
				Expect(err).NotTo(HaveOccurred())
				Expect(vetOrgOne.ID).NotTo(Equal(uint(0)))
			})

			It("It returns the vetOrg and nil for error", func() {
				foundVetOrg, err = vetOrgRepo.GetByName(*vetOrgOne.OrgName)
				Expect(err).NotTo(HaveOccurred())
				Expect(foundVetOrg).Should(HaveLen(1))
				Expect(foundVetOrg[0].ID).To(Equal(vetOrgOne.ID))
				Expect(foundVetOrg[0].OrgName).To(Equal(vetOrgOne.OrgName))
			})
		})

		Context("When the vetOrg is not found", func() {

			var foundCustomer []model.VetOrg

			BeforeEach(func() {
			})

			It("It returns the vetOrg and nil for error", func() {
				foundCustomer, err = vetOrgRepo.GetByName("some_user_name")
				Expect(err).To(HaveOccurred())
				Expect(foundCustomer).To(BeNil())
			})
		})
	})
})
