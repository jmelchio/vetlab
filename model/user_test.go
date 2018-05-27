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
			goUser = User{
				UserID:       "some-user-id",
				UserName:     "some-user-name",
				FirstName:    "some-first-name",
				LastName:     "some-last-name",
				Email:        "user@server.com",
				PasswordHash: "some-password-hashed",
				OrgID:        "some-org-id",
				AdminUser:    false,
			}
			jsonUser = `{"user_id":"some-user-id","user_name":"some-user-name","first_name":"some-first-name","last_name":"some-last-name","email":"user@server.com","password_hash":"some-password-hashed","org_id":"some-org-id","admin_user":false}`
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
