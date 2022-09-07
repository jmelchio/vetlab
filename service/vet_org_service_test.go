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
				vetOrgRepo.CreateReturns(errors.New("failed to save record"))
			})

			It("Fails to create a vetOrg and returns a 'missing context' error", func() {
				zeCustomer, err := vetOrgService.CreateVetOrg(context.TODO(), vetOrg)
				Expect(err).To(HaveOccurred())
				Expect(zeCustomer).To(BeNil())
				Expect(vetOrgRepo.CreateCallCount()).To(Equal(1))
			})
		})
	})

	Describe("Update a vetOrg", func() {

		BeforeEach(func() {
			vetOrgRepo.UpdateReturns(nil)
		})

		Context("We have a valid vetOrg and context", func() {

			It("Returns the updated vetOrg and no error", func() {
				zeVetOrg, err := vetOrgService.UpdateVetOrg(context.TODO(), vetOrg)
				Expect(err).NotTo(HaveOccurred())
				Expect(vetOrgRepo.UpdateCallCount()).To(Equal(1))
				Expect(zeVetOrg).NotTo(BeNil())
			})
		})

		Context("We have a valid vetOrg but no context", func() {

			It("Returns and error and no updated vetOrg", func() {
				zeVetOrg, err := vetOrgService.UpdateVetOrg(nil, vetOrg)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal(MissingContext))
				Expect(zeVetOrg).To(BeNil())
				Expect(vetOrgRepo.UpdateCallCount()).To(Equal(0))
			})
		})

		Context("We have a vetOrg and Context but repo cannot update the vetOrg", func() {

			BeforeEach(func() {
				vetOrgRepo.UpdateReturns(errors.New("unable to update the vetOrg"))
			})

			It("Returns an error after calling VetOrgRepo.Create", func() {
				zeVetOrg, err := vetOrgService.UpdateVetOrg(context.TODO(), vetOrg)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("unable to update the vetOrg"))
				Expect(zeVetOrg).To(BeNil())
				Expect(vetOrgRepo.UpdateCallCount()).To(Equal(1))
			})
		})
	})

	Describe("Delete a vetOrg", func() {

		BeforeEach(func() {
			vetOrgRepo.DeleteReturns(nil)
		})

		Context("We have a valid vetOrg and context", func() {

			It("Returns no error", func() {
				err := vetOrgService.DeleteVetOrg(context.TODO(), vetOrg)
				Expect(err).NotTo(HaveOccurred())
				Expect(vetOrgRepo.DeleteCallCount()).To(Equal(1))
			})
		})

		Context("We have a valid vetOrg but no context", func() {

			It("Returns and error and no updated vetOrg", func() {
				err := vetOrgService.DeleteVetOrg(nil, vetOrg)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal(MissingContext))
				Expect(vetOrgRepo.DeleteCallCount()).To(Equal(0))
			})
		})

		Context("We have a vetOrg and Context but repo cannot delete vetOrg", func() {

			BeforeEach(func() {
				vetOrgRepo.DeleteReturns(errors.New("unable to delete the vetOrg"))
			})

			It("Returns an error after calling VetOrgRepo.Create", func() {
				err := vetOrgService.DeleteVetOrg(context.TODO(), vetOrg)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("unable to delete the vetOrg"))
				Expect(vetOrgRepo.DeleteCallCount()).To(Equal(1))
			})
		})
	})

	Describe("Find VetOrg by name", func() {
		Context("We have a valid VetOrg name and Context", func() {

			BeforeEach(func() {
				resultVetOrgList := []model.VetOrg{vetOrg}
				vetOrgRepo.GetByNameReturns(resultVetOrgList, nil)
			})

			It("Returns a list of VetOrg(s) and no error from the repo", func() {
				result, err := vetOrgService.FindVetOrgByName(context.TODO(), "orgName")
				Expect(err).ToNot(HaveOccurred())
				Expect(len(result)).To(Equal(1))
				Expect(result[0]).To(Equal(vetOrg))
			})
		})

		Context("We have a valid VetOrg name but no Context", func() {

			BeforeEach(func() {
				vetOrgRepo.GetByNameReturns(nil, errors.New(MissingContext))
			})

			It("Returns an error indicating the absence of the context", func() {
				result, err := vetOrgService.FindVetOrgByName(context.TODO(), "orgName")
				Expect(err).To(HaveOccurred())
				Expect(result).To(BeNil())
				Expect(err.Error()).To(Equal(MissingContext))
			})
		})

		Context("We have a valid Context but no valid VetOrg name", func() {

			BeforeEach(func() {
				vetOrgRepo.GetByNameReturns(nil, nil)
			})

			It("Returns no results list and no error", func() {
				result, err := vetOrgService.FindVetOrgByName(context.TODO(), "orgName")
				Expect(err).NotTo(HaveOccurred())
				Expect(result).To(BeNil())
			})
		})

		Context("We have a valid Context and VetOrg name but repo errors out", func() {

			BeforeEach(func() {
				vetOrgRepo.GetByNameReturns(nil, errors.New("BAM"))
			})

			It("Returns no results list and an error", func() {
				result, err := vetOrgService.FindVetOrgByName(context.TODO(), "orgName")
				Expect(err).To(HaveOccurred())
				Expect(result).To(BeNil())
				Expect(err.Error()).To(Equal("BAM"))
			})
		})
	})
})
