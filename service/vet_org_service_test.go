package service_test

import (
	"context"
	"errors"

	"github.com/jmelchio/vetlab/api"
	"github.com/jmelchio/vetlab/model"
	. "github.com/jmelchio/vetlab/service"
	"github.com/jmelchio/vetlab/service/servicefakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
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

	Describe("Create a vetOrg", func() {

		BeforeEach(func() {
			vetOrgRepo.CreateReturns(nil)
		})

		Context("We have a valid vetOrg and 'todo' context", func() {

			It("Returns a vetOrg with a new vetOrg ID and calls VetOrgRepo.Create", func() {
				zeCustomer, err := vetOrgService.CreateVetOrg(context.TODO(), vetOrg)
				Expect(err).ToNot(HaveOccurred())
				Expect(zeCustomer).NotTo(BeNil())
				Expect(vetOrgRepo.CreateCallCount()).To(Equal(1))
			})
		})

		Context("We have a valid vetOrg but context is missing", func() {

			It("Fails to create a vetOrg and returns a 'missing context' error", func() {
				zeCustomer, err := vetOrgService.CreateVetOrg(nil, vetOrg)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal(MissingContext))
				Expect(zeCustomer).To(BeNil())
				Expect(vetOrgRepo.CreateCallCount()).To(Equal(0))
			})
		})

		Context("We have a valid vetOrg and context but repo fails to save", func() {

			BeforeEach(func() {
				vetOrgRepo.CreateReturns(errors.New("Failed to save record"))
			})

			It("Fails to create a vetOrg and returns a 'missing context' error", func() {
				zeCustomer, err := vetOrgService.CreateVetOrg(context.TODO(), vetOrg)
				Expect(err).To(HaveOccurred())
				Expect(zeCustomer).To(BeNil())
				Expect(vetOrgRepo.CreateCallCount()).To(Equal(1))
			})
		})
	})
})
