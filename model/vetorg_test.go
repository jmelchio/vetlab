package model_test

import (
	"encoding/json"

	. "github.com/jmelchio/vetlab/model"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Vetorg", func() {
	Describe("Vetorg can be transformed from and to Json", func() {
		Context("From Golang to Json", func() {
			var (
				goVetorg     VetOrg
				expectedJson string
			)
			BeforeEach(func() {
				goVetorg = VetOrg{
					OrgID:       "some-org-id",
					OrgName:     "some-org-name",
					Street:      "some-street",
					HouseNumber: "some-house-number",
					City:        "some-city",
					Province:    "some-province",
					Country:     "some-country",
					PostalCode:  "some-postal-code",
					Email:       "some-email",
					Phone:       "some-phone",
					Fax:         "some-fax",
				}
				expectedJson = `{"org_id": "some-org-id"}`
			})
			It("Transforms without errors", func() {
				vetOrgBytes, err := json.Marshal(goVetorg)
				Expect(err).NotTo(HaveOccurred())

				jsonResult := string(vetOrgBytes)
				Expect(jsonResult).To(Equal(expectedJson))
			})
		})
		Context("From Json to Golang", func() {

		})
	})
})
