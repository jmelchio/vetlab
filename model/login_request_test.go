package model_test

import (
	"encoding/json"

	. "github.com/jmelchio/vetlab/model"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("LoginRequest", func() {

	Describe("LoginRequest object can be transformed from and to Json", func() {

		var (
			goLoginRequest   LoginRequest
			jsonLoginRequest string
		)

		BeforeEach(func() {
			goLoginRequest = LoginRequest{
				UserName: "some-user-name",
				Password: "some-password",
			}
			jsonLoginRequest = `{"user_name":"some-user-name","password":"some-password"}`
		})

		Context("From Golang to Json", func() {

			It("transforms without errors", func() {
				userBytes, err := json.Marshal(goLoginRequest)
				Expect(err).NotTo(HaveOccurred())

				jsonResult := string(userBytes)
				Expect(jsonResult).To(Equal(jsonLoginRequest))
			})
		})

		Context("From Json to Golang", func() {

			It("Transforms without errors", func() {
				var unmarshalLoginRequest LoginRequest
				err := json.Unmarshal([]byte(jsonLoginRequest), &unmarshalLoginRequest)
				Expect(err).NotTo(HaveOccurred())

				Expect(unmarshalLoginRequest).To(Equal(goLoginRequest))
			})
		})
	})
})
