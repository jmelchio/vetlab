package model_test

import (
	"encoding/json"

	. "github.com/jmelchio/vetlab/model"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Customer", func() {

	Describe("Customer object can be transformed from and to Json", func() {

		var (
			goCustomer   Customer
			jsonCustomer string
		)

		BeforeEach(func() {
			firstName := "some-first-name"
			lastName := "some-last-name"
			email := "user@server.com"
			goCustomer = Customer{
				ID:        uint(12345),
				FirstName: &firstName,
				LastName:  &lastName,
				Email:     &email,
				VetOrgID:  uint(12345),
			}
			jsonCustomer = `{"id":12345,"first_name":"some-first-name","last_name":"some-last-name","email":"user@server.com","vet_org_id":12345}`
		})

		Context("From Golang to Json", func() {

			It("transforms without errors", func() {
				userBytes, err := json.Marshal(goCustomer)
				Expect(err).NotTo(HaveOccurred())

				jsonResult := string(userBytes)
				Expect(jsonResult).To(Equal(jsonCustomer))
			})
		})

		Context("From Json to Golang", func() {

			It("Transforms without errors", func() {
				var unmarshalCustomer Customer
				err := json.Unmarshal([]byte(jsonCustomer), &unmarshalCustomer)
				Expect(err).NotTo(HaveOccurred())

				Expect(unmarshalCustomer).To(Equal(goCustomer))
			})
		})
	})
})
