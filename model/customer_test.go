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
			userName := "user-name"
			firstName := "some-first-name"
			lastName := "some-last-name"
			email := "user@server.com"
			password := "password"
			goCustomer = Customer{
				ID:        uint(12345),
				UserName:  &userName,
				FirstName: firstName,
				LastName:  lastName,
				Email:     email,
				Password:  password,
				VetOrgID:  uint(12345),
			}
			jsonCustomer = `{"id":12345,"user_name":"user-name","first_name":"some-first-name","last_name":"some-last-name","email":"user@server.com","password":"password","vet_org_id":12345}`
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
