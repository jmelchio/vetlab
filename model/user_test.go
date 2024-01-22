package model_test

import (
	"encoding/json"
	"time"

	. "github.com/jmelchio/vetlab/model"
	"gorm.io/gorm"

	. "github.com/onsi/ginkgo/v2"
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
			password := "some-password"
			goUser = User{
				Model: gorm.Model{
					ID:        uint(12345),
					CreatedAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
					UpdatedAt: time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
					DeletedAt: gorm.DeletedAt{},
				},
				UserName:  &userName,
				FirstName: firstName,
				LastName:  lastName,
				Email:     email,
				Password:  password,
				AdminUser: false,
			}
			jsonUser = `{"ID":12345,"CreatedAt":"2020-01-01T00:00:00Z","UpdatedAt":"2020-01-01T00:00:00Z","DeletedAt":null,"user_name":"some-user-name","first_name":"some-first-name","last_name":"some-last-name","email":"user@server.com","password":"some-password","admin_user":false}`
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
