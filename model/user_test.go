package model_test

import (
	"encoding/json"

	. "github.com/jmelchio/vetlab/model"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("User", func() {
	Describe("User object can be transformed from and to Json", func() {
		Context("From Golang to Json", func() {
			var (
				goUser       User
				expectedJson string
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
				}
				expectedJson = `{"user_id": "some-user-id"}`
			})
			It("transforms without errors", func() {
				userBytes, err := json.Marshal(goUser)
				Expect(err).NotTo(HaveOccurred())

				jsonResult := string(userBytes)
				Expect(jsonResult).To(Equal(expectedJson))
			})
		})
		Context("From Json to Golang", func() {
			BeforeEach(func() {

			})
		})
	})
})
