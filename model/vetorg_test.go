package model_test

import (
	"encoding/json"
	"time"

	. "github.com/jmelchio/vetlab/model"
	"gorm.io/gorm"

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
				Model: gorm.Model{
					ID:        uint(12345),
					CreatedAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
					UpdatedAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
					DeletedAt: gorm.DeletedAt{},
				},
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
			jsonVetorg = `{"ID":12345,"CreatedAt":"2020-01-01T00:00:00Z","UpdatedAt":"2020-01-01T00:00:00Z","DeletedAt":null,"org_name":"some-org-name","street":"some-street","house_number":"some-house-number","city":"some-city","province":"some-province","country":"some-country","postal_code":"some-postal-code","email":"person@domain.com","phone":"some-phone","fax":"shadow-fax"}`
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
