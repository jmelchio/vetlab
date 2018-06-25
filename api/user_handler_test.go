package api_test

import (
	"net/http"
	"net/http/httptest"

	"github.com/jmelchio/vetlab/api/apifakes"

	. "github.com/jmelchio/vetlab/api"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("UserHandler", func() {
	var (
		handler     http.Handler
		recorder    *httptest.ResponseRecorder
		err         error
		userService *apifakes.FakeUserService
	)

	BeforeEach(func() {
		userService = new(apifakes.FakeUserService)
		handler, err = NewUserHandler(userService)
		Expect(err).NotTo(HaveOccurred())
	})

	Describe("Create a user", func() {

		BeforeEach(func() {
		})

		Context("Valid user information is passed", func() {

			BeforeEach(func() {
				recorder = httptest.NewRecorder()
				request, _ := http.NewRequest("POST", "/user/create", nil)
				handler.ServeHTTP(recorder, request)
			})

			It("Creates a user and returns 201 status code", func() {
				Expect(recorder.Result().StatusCode).To(Equal(http.StatusCreated))
			})
		})
	})

	Describe("Update a user", func() {
	})

	Describe("Delete a user", func() {
	})

	Describe("Login a user", func() {
	})

	Describe("Find a user", func() {
	})
})
