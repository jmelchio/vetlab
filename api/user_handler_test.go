package api_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"

	"github.com/jmelchio/vetlab/api/apifakes"
	"github.com/jmelchio/vetlab/model"

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

		var (
			createUser model.User
		)

		BeforeEach(func() {
			createUser = model.User{
				UserName:     "user_name",
				FirstName:    "first_name",
				LastName:     "last_name",
				Email:        "some_email",
				PasswordHash: "some_hash",
				OrgID:        "some_org_id",
				AdminUser:    false,
			}
		})

		Context("Valid user information is passed", func() {

			BeforeEach(func() {
				userService.CreateUserReturns(&createUser, nil)
				userBytes, err := json.Marshal(createUser)
				Expect(err).NotTo(HaveOccurred())
				recorder = httptest.NewRecorder()
				request, _ := http.NewRequest("POST", "/user/create", bytes.NewReader(userBytes))
				handler.ServeHTTP(recorder, request)
			})

			It("Creates a user and returns 201 status code", func() {
				Expect(recorder.Result().StatusCode).To(Equal(http.StatusCreated))
				respBody, err := ioutil.ReadAll(recorder.Result().Body)
				Expect(err).NotTo(HaveOccurred())

				var newUser model.User
				err = json.Unmarshal(respBody, &newUser)
				Expect(err).NotTo(HaveOccurred())
				Expect(newUser).NotTo(BeNil())
				Expect(newUser.FirstName).To(Equal(createUser.FirstName))
				Expect(newUser.LastName).To(Equal(createUser.LastName))
				Expect(userService.CreateUserCallCount()).To(Equal(1))
			})

		})

		Context("Valid user information is passed but downstream call fails", func() {

			BeforeEach(func() {
				userService.CreateUserReturns(nil, errors.New("Whoot?"))
				userBytes, err := json.Marshal(createUser)
				Expect(err).NotTo(HaveOccurred())
				recorder = httptest.NewRecorder()
				request, _ := http.NewRequest("POST", "/user/create", bytes.NewReader(userBytes))
				handler.ServeHTTP(recorder, request)
			})

			It("Fails to return a users and returns 500 status code", func() {
				Expect(recorder.Result().StatusCode).To(Equal(http.StatusInternalServerError))
				respBody, err := ioutil.ReadAll(recorder.Result().Body)
				Expect(err).NotTo(HaveOccurred())
				Expect(err).NotTo(HaveOccurred())
				Expect(string(respBody[0 : len(respBody)-1])).To(Equal(UnableToCreateUser))
				Expect(userService.CreateUserCallCount()).To(Equal(1))
			})
		})

		Context("Body of the request is empty", func() {

			BeforeEach(func() {
				recorder = httptest.NewRecorder()
				request, _ := http.NewRequest("POST", "/user/create", nil)
				handler.ServeHTTP(recorder, request)
			})

			It("Fails to create a user and returns a 400 status code", func() {
				Expect(recorder.Result().StatusCode).To(Equal(http.StatusBadRequest))
				respBody, err := ioutil.ReadAll(recorder.Result().Body)
				Expect(err).NotTo(HaveOccurred())
				Expect(string(respBody[0 : len(respBody)-1])).To(Equal(EmptyBody))
				Expect(userService.CreateUserCallCount()).To(Equal(0))
			})
		})

		Context("Body of the request contains invalid data", func() {

			BeforeEach(func() {
				userBytes, err := json.Marshal("createUser")
				Expect(err).NotTo(HaveOccurred())
				recorder = httptest.NewRecorder()
				request, _ := http.NewRequest("POST", "/user/create", bytes.NewReader(userBytes))
				handler.ServeHTTP(recorder, request)
			})

			It("Fails to create a user and returns a 400 status code", func() {
				Expect(recorder.Result().StatusCode).To(Equal(http.StatusBadRequest))
				respBody, err := ioutil.ReadAll(recorder.Result().Body)
				Expect(err).NotTo(HaveOccurred())
				Expect(string(respBody[0 : len(respBody)-1])).To(Equal(InvalidBody))
				Expect(userService.CreateUserCallCount()).To(Equal(0))
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
