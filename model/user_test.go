package model_test

import (
	"encoding/json"

	. "github.com/jmelchio/vetlab/model"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("User", func() {

	Describe("User object can be transformed from and to Json", func() {

		var (
			goUser   User
			jsonUser string
		)

		BeforeEach(func() {
			userName := "some-user-name"
			firstName := "some-first-name"
			lastName := "some-last-name"
			email := "user@server.com"
			passwordHash := "some-password-hashed"
			goUser = User{
				ID:           uint(12345),
				UserName:     &userName,
				FirstName:    &firstName,
				LastName:     &lastName,
				Email:        &email,
				PasswordHash: &passwordHash,
				VetOrgID:     uint(12345),
				AdminUser:    false,
			}
			jsonUser = `{"id":12345,"user_name":"some-user-name","first_name":"some-first-name","last_name":"some-last-name","email":"user@server.com","password_hash":"some-password-hashed","vet_org_id":12345,"admin_user":false}`
		})

		Context("From Golang to Json", func() {

			It("transforms without errors", func() {
				userBytes, err := json.Marshal(goUser)
				Expect(err).NotTo(HaveOccurred())

				jsonResult := string(userBytes)
				Expect(jsonResult).To(Equal(jsonUser))
			})
		})

		Context("From Json to Golang", func() {

			It("Transforms without errors", func() {
				var unmarshalUser User
				err := json.Unmarshal([]byte(jsonUser), &unmarshalUser)
				Expect(err).NotTo(HaveOccurred())

				Expect(unmarshalUser).To(Equal(goUser))
			})
		})
	})
})
