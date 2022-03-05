package sql_test

import (
	"github.com/jmelchio/vetlab/model"
	"github.com/jmelchio/vetlab/repository/sql"
	"github.com/jmelchio/vetlab/service"
	. "github.com/onsi/ginkgo"
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

		Context("When a username is not taken yet", func() {
			BeforeEach(func() {
				Expect(vetOrgOne.ID).To(Equal(uint(0)))
			})

			It("Creates a new vetOrg record", func() {
				err = vetOrgRepo.Create(&vetOrgOne)
				Expect(err).NotTo(HaveOccurred())
				Expect(vetOrgOne.ID).NotTo(Equal(uint(0)))
			})
		})

		Context("When a username is taken already", func() {

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
})
