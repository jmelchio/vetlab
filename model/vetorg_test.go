package model_test

import (
	"encoding/json"

	. "github.com/jmelchio/vetlab/model"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Vetorg", func() {

	Describe("Vetorg can be transformed from and to Json", func() {
		var (
			goVetorg   VetOrg
			jsonVetorg string
		)

		BeforeEach(func() {
			orgName := "some-org-name"
			street := "some-street"
			houseNumber := "some-house-number"
			city := "some-city"
			province := "some-province"
			country := "some-country"
			postalCode := "some-postal-code"
			email := "person@domain.com"
			phone := "some-phone"
			fax := "shadow-fax"
			goVetorg = VetOrg{
				ID:          12345,
				OrgName:     &orgName,
				Street:      street,
				HouseNumber: houseNumber,
				City:        city,
				Province:    province,
				Country:     country,
				PostalCode:  postalCode,
				Email:       email,
				Phone:       phone,
				Fax:         fax,
			}
			jsonVetorg = `{"id":12345,"org_name":"some-org-name","street":"some-street","house_number":"some-house-number","city":"some-city","province":"some-province","country":"some-country","postal_code":"some-postal-code","email":"person@domain.com","phone":"some-phone","fax":"shadow-fax"}`
		})

		Context("From Golang to Json", func() {

			It("Transforms without errors", func() {
				vetOrgBytes, err := json.Marshal(goVetorg)
				Expect(err).NotTo(HaveOccurred())

				jsonResult := string(vetOrgBytes)
				Expect(jsonResult).To(Equal(jsonVetorg))
			})
		})

		Context("From Json to Golang", func() {

			It("Transforms without errors", func() {
				var unmarshalVetOrg VetOrg
				err := json.Unmarshal([]byte(jsonVetorg), &unmarshalVetOrg)
				Expect(err).NotTo(HaveOccurred())

				Expect(unmarshalVetOrg).To(Equal(goVetorg))
			})
		})
	})
})
