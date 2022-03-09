package service_test

import (
	"github.com/jmelchio/vetlab/api"
	"github.com/jmelchio/vetlab/model"
	. "github.com/jmelchio/vetlab/service"
	"github.com/jmelchio/vetlab/service/servicefakes"
	. "github.com/onsi/ginkgo"
)

var _ = Describe("VetOrgService", func() {

	var (
		vetOrgService api.VetOrgService
		vetOrgRepo    *servicefakes.FakeVetOrgRepo
		vetOrg        model.VetOrg
		orgName       string
		street        string
		houseNumber   string
		city          string
		province      string
		country       string
		postalCode    string
		email         string
		phone         string
		fax           string
	)

	BeforeEach(func() {
		vetOrgRepo = new(servicefakes.FakeVetOrgRepo)
		vetOrgServiceImpl := VetOrg{VetOrgRepo: vetOrgRepo}
		vetOrgService = vetOrgServiceImpl

		orgName = "some-name"
		street = "zestreet"
		houseNumber = "42"
		city = "zecity"
		province = "zeprovince"
		country = "zecountry"
		postalCode = "1185 JR"
		email = "email@domain.com"
		phone = "020 641 6890"
		fax = "020 641 6890"

		vetOrg = model.VetOrg{
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
			Customers:          nil,
			DiagnosticReports:  nil,
			DiagnosticRequests: nil,
		}
	})

})
