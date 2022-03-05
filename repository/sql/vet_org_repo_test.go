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
})
