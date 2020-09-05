package api_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"

	"github.com/tedsuo/rata"

	"github.com/jmelchio/vetlab/api/apifakes"
	"github.com/jmelchio/vetlab/model"

	"github.com/jmelchio/vetlab/api"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("UserHandler", func() {
	var (
		handler          http.Handler
		recorder         *httptest.ResponseRecorder
		err              error
		userService      *apifakes.FakeUserService
		requestGenerator *rata.RequestGenerator
		userName         string
		firstName        string
		lastName         string
		email            string
		password         string
	)

	BeforeEach(func() {
		userService = new(apifakes.FakeUserService)
		handler, err = api.NewUserHandler(userService)
		requestGenerator = rata.NewRequestGenerator("", api.UserRoutes)
		Expect(err).NotTo(HaveOccurred())

		userName = "user_name"
		firstName = "first_name"
		lastName = "last_name"
		email = "some_email"
		password = "some_password"
	})

	Describe("Create a user", func() {

		var (
			createUser model.User
		)

		BeforeEach(func() {
			createUser = model.User{
				UserName:  &userName,
				FirstName: firstName,
				LastName:  lastName,
				Email:     email,
				Password:  password,
				AdminUser: false,
			}
		})

		Context("Valid user information is passed", func() {

			BeforeEach(func() {
				userService.CreateUserReturns(&createUser, nil)
				userBytes, err := json.Marshal(createUser)
				Expect(err).NotTo(HaveOccurred())
				recorder = httptest.NewRecorder()
				request, _ := requestGenerator.CreateRequest(api.CreateUser, nil, bytes.NewReader(userBytes))
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
				userService.CreateUserReturns(nil, errors.New("whoot"))
				userBytes, err := json.Marshal(createUser)
				Expect(err).NotTo(HaveOccurred())
				recorder = httptest.NewRecorder()
				request, _ := requestGenerator.CreateRequest(api.CreateUser, nil, bytes.NewReader(userBytes))
				handler.ServeHTTP(recorder, request)
			})

			It("Fails to return a users and returns 500 status code", func() {
				Expect(recorder.Result().StatusCode).To(Equal(http.StatusInternalServerError))
				respBody, err := ioutil.ReadAll(recorder.Result().Body)
				Expect(err).NotTo(HaveOccurred())
				Expect(err).NotTo(HaveOccurred())
				Expect(string(respBody[0 : len(respBody)-1])).To(Equal(api.UnableToCreateUser))
				Expect(userService.CreateUserCallCount()).To(Equal(1))
			})
		})

		Context("Body of the request is empty", func() {

			BeforeEach(func() {
				recorder = httptest.NewRecorder()
				request, _ := requestGenerator.CreateRequest(api.CreateUser, nil, nil)
				handler.ServeHTTP(recorder, request)
			})

			It("Fails to create a user and returns a 400 status code", func() {
				Expect(recorder.Result().StatusCode).To(Equal(http.StatusBadRequest))
				respBody, err := ioutil.ReadAll(recorder.Result().Body)
				Expect(err).NotTo(HaveOccurred())
				Expect(string(respBody[0 : len(respBody)-1])).To(Equal(api.EmptyBody))
				Expect(userService.CreateUserCallCount()).To(Equal(0))
			})
		})

		Context("Body of the request contains invalid data", func() {

			BeforeEach(func() {
				userBytes, err := json.Marshal("createUser")
				Expect(err).NotTo(HaveOccurred())
				recorder = httptest.NewRecorder()
				request, _ := requestGenerator.CreateRequest(api.CreateUser, nil, bytes.NewReader(userBytes))
				handler.ServeHTTP(recorder, request)
			})

			It("Fails to create a user and returns a 400 status code", func() {
				Expect(recorder.Result().StatusCode).To(Equal(http.StatusBadRequest))
				respBody, err := ioutil.ReadAll(recorder.Result().Body)
				Expect(err).NotTo(HaveOccurred())
				Expect(string(respBody[0 : len(respBody)-1])).To(Equal(api.InvalidBody))
				Expect(userService.CreateUserCallCount()).To(Equal(0))
			})
		})
	})

	Describe("Update a user", func() {
		var (
			updateUser model.User
		)

		BeforeEach(func() {
			updateUser = model.User{
				ID:        uint(12345),
				UserName:  &userName,
				FirstName: firstName,
				LastName:  lastName,
				Email:     email,
				Password:  password,
				AdminUser: false,
			}
		})

		Context("Valid user information is passed", func() {

			BeforeEach(func() {
				userService.UpdateUserReturns(&updateUser, nil)
				userBytes, err := json.Marshal(updateUser)
				Expect(err).NotTo(HaveOccurred())
				recorder = httptest.NewRecorder()
				request, _ := requestGenerator.CreateRequest(api.UpdateUser, nil, bytes.NewReader(userBytes))
				handler.ServeHTTP(recorder, request)
			})

			It("Updates a user and returns 200 status code", func() {
				Expect(recorder.Result().StatusCode).To(Equal(http.StatusOK))
				respBody, err := ioutil.ReadAll(recorder.Result().Body)
				Expect(err).NotTo(HaveOccurred())

				var newUser model.User
				err = json.Unmarshal(respBody, &newUser)
				Expect(err).NotTo(HaveOccurred())
				Expect(newUser).NotTo(BeNil())
				Expect(newUser.FirstName).To(Equal(updateUser.FirstName))
				Expect(newUser.LastName).To(Equal(updateUser.LastName))
				Expect(userService.UpdateUserCallCount()).To(Equal(1))
			})
		})

		Context("Valid user information is passed but downstream call fails", func() {

			BeforeEach(func() {
				userService.UpdateUserReturns(nil, errors.New("whoot"))
				userBytes, err := json.Marshal(updateUser)
				Expect(err).NotTo(HaveOccurred())
				recorder = httptest.NewRecorder()
				request, _ := requestGenerator.CreateRequest(api.UpdateUser, nil, bytes.NewReader(userBytes))
				handler.ServeHTTP(recorder, request)
			})

			It("Fails to return a users and returns 500 status code", func() {
				Expect(recorder.Result().StatusCode).To(Equal(http.StatusInternalServerError))
				respBody, err := ioutil.ReadAll(recorder.Result().Body)
				Expect(err).NotTo(HaveOccurred())
				Expect(err).NotTo(HaveOccurred())
				Expect(string(respBody[0 : len(respBody)-1])).To(Equal(api.UnableToUpdateUser))
				Expect(userService.UpdateUserCallCount()).To(Equal(1))
			})
		})

		Context("Body of the request is empty", func() {

			BeforeEach(func() {
				recorder = httptest.NewRecorder()
				request, _ := requestGenerator.CreateRequest(api.UpdateUser, nil, nil)
				handler.ServeHTTP(recorder, request)
			})

			It("Fails to update a user and returns a 400 status code", func() {
				Expect(recorder.Result().StatusCode).To(Equal(http.StatusBadRequest))
				respBody, err := ioutil.ReadAll(recorder.Result().Body)
				Expect(err).NotTo(HaveOccurred())
				Expect(string(respBody[0 : len(respBody)-1])).To(Equal(api.EmptyBody))
				Expect(userService.UpdateUserCallCount()).To(Equal(0))
			})
		})

		Context("Body of the request contains invalid data", func() {

			BeforeEach(func() {
				userBytes, err := json.Marshal("updateUser")
				Expect(err).NotTo(HaveOccurred())
				recorder = httptest.NewRecorder()
				request, _ := requestGenerator.CreateRequest(api.UpdateUser, nil, bytes.NewReader(userBytes))
				handler.ServeHTTP(recorder, request)
			})

			It("Fails to update a user and returns a 400 status code", func() {
				Expect(recorder.Result().StatusCode).To(Equal(http.StatusBadRequest))
				respBody, err := ioutil.ReadAll(recorder.Result().Body)
				Expect(err).NotTo(HaveOccurred())
				Expect(string(respBody[0 : len(respBody)-1])).To(Equal(api.InvalidBody))
				Expect(userService.UpdateUserCallCount()).To(Equal(0))
			})
		})
	})

	Describe("Delete a user", func() {
		var (
			deleteUser model.User
		)

		BeforeEach(func() {
			deleteUser = model.User{
				ID:        uint(12345),
				UserName:  &userName,
				FirstName: firstName,
				LastName:  lastName,
				Email:     email,
				Password:  password,
				AdminUser: false,
			}
		})

		Context("Valid user information is passed", func() {

			BeforeEach(func() {
				userService.DeleteUserReturns(nil)
				userBytes, err := json.Marshal(deleteUser)
				Expect(err).NotTo(HaveOccurred())
				recorder = httptest.NewRecorder()
				request, _ := requestGenerator.CreateRequest(api.DeleteUser, nil, bytes.NewReader(userBytes))
				handler.ServeHTTP(recorder, request)
			})

			It("Deletes a user and returns 204 status code", func() {
				Expect(recorder.Result().StatusCode).To(Equal(http.StatusNoContent))
				Expect(userService.DeleteUserCallCount()).To(Equal(1))
			})
		})

		Context("Valid user information is passed but downstream call fails", func() {

			BeforeEach(func() {
				userService.DeleteUserReturns(errors.New("whoot"))
				userBytes, err := json.Marshal(deleteUser)
				Expect(err).NotTo(HaveOccurred())
				recorder = httptest.NewRecorder()
				request, _ := requestGenerator.CreateRequest(api.DeleteUser, nil, bytes.NewReader(userBytes))
				handler.ServeHTTP(recorder, request)
			})

			It("Fails to delete a user and returns a 500 status code", func() {
				Expect(recorder.Result().StatusCode).To(Equal(http.StatusInternalServerError))
				respBody, err := ioutil.ReadAll(recorder.Result().Body)
				Expect(err).NotTo(HaveOccurred())
				Expect(string(respBody[0 : len(respBody)-1])).To(Equal(api.UnableToDeleteUser))
				Expect(userService.DeleteUserCallCount()).To(Equal(1))
			})
		})

		Context("Body of the request is empty", func() {

			BeforeEach(func() {
				recorder = httptest.NewRecorder()
				request, _ := requestGenerator.CreateRequest(api.DeleteUser, nil, nil)
				handler.ServeHTTP(recorder, request)
			})

			It("Fails to delete a user and returns a 400 status code", func() {
				Expect(recorder.Result().StatusCode).To(Equal(http.StatusBadRequest))
				respBody, err := ioutil.ReadAll(recorder.Result().Body)
				Expect(err).NotTo(HaveOccurred())
				Expect(string(respBody[0 : len(respBody)-1])).To(Equal(api.EmptyBody))
				Expect(userService.DeleteUserCallCount()).To(Equal(0))
			})
		})

		Context("Body of the request contains invalid data", func() {

			BeforeEach(func() {
				userBytes, err := json.Marshal("deleteUser")
				Expect(err).NotTo(HaveOccurred())
				recorder = httptest.NewRecorder()
				request, _ := requestGenerator.CreateRequest(api.DeleteUser, nil, bytes.NewReader(userBytes))
				handler.ServeHTTP(recorder, request)
			})

			It("Fails to delete a user and returns a 400 status code", func() {
				Expect(recorder.Result().StatusCode).To(Equal(http.StatusBadRequest))
				respBody, err := ioutil.ReadAll(recorder.Result().Body)
				Expect(err).NotTo(HaveOccurred())
				Expect(string(respBody[0 : len(respBody)-1])).To(Equal(api.InvalidBody))
				Expect(userService.DeleteUserCallCount()).To(Equal(0))
			})
		})
	})

	Describe("Login a user", func() {

		var (
			loginRequest model.LoginRequest
			loginUser    model.User
		)

		BeforeEach(func() {
			loginRequest = model.LoginRequest{
				UserName: "user_name",
				Password: "some_password",
			}

			loginUser = model.User{
				ID:        uint(12345),
				UserName:  &userName,
				FirstName: firstName,
				LastName:  lastName,
				Email:     email,
				Password:  password,
				AdminUser: false,
			}
		})

		Context("Valid user information is passed", func() {

			BeforeEach(func() {
				userService.LoginReturns(&loginUser, nil)
				userBytes, err := json.Marshal(loginRequest)
				Expect(err).NotTo(HaveOccurred())
				recorder = httptest.NewRecorder()
				request, _ := requestGenerator.CreateRequest(api.Login, nil, bytes.NewReader(userBytes))
				handler.ServeHTTP(recorder, request)
			})

			It("Logs a user in and returns 200 status code", func() {
				Expect(recorder.Result().StatusCode).To(Equal(http.StatusOK))
				respBody, err := ioutil.ReadAll(recorder.Result().Body)
				Expect(err).NotTo(HaveOccurred())

				var newUser model.User
				err = json.Unmarshal(respBody, &newUser)
				Expect(err).NotTo(HaveOccurred())
				Expect(newUser).NotTo(BeNil())
				Expect(newUser.FirstName).To(Equal(loginUser.FirstName))
				Expect(newUser.LastName).To(Equal(loginUser.LastName))
				Expect(userService.LoginCallCount()).To(Equal(1))
			})
		})

		Context("Valid user information is passed but downstream call fails", func() {

			BeforeEach(func() {
				userService.LoginReturns(nil, errors.New("whoot"))
				userBytes, err := json.Marshal(loginRequest)
				Expect(err).NotTo(HaveOccurred())
				recorder = httptest.NewRecorder()
				request, _ := requestGenerator.CreateRequest(api.Login, nil, bytes.NewReader(userBytes))
				handler.ServeHTTP(recorder, request)
			})

			It("Fails to login a user and returns a 500 status code", func() {
				Expect(recorder.Result().StatusCode).To(Equal(http.StatusInternalServerError))
				respBody, err := ioutil.ReadAll(recorder.Result().Body)
				Expect(err).NotTo(HaveOccurred())
				Expect(string(respBody[0 : len(respBody)-1])).To(Equal(api.UnableToLoginUser))
				Expect(userService.LoginCallCount()).To(Equal(1))
			})
		})

		Context("Body of the request is empty", func() {

			BeforeEach(func() {
				recorder = httptest.NewRecorder()
				request, _ := requestGenerator.CreateRequest(api.Login, nil, nil)
				handler.ServeHTTP(recorder, request)
			})

			It("Fails to login a user and returns a 400 status code", func() {
				Expect(recorder.Result().StatusCode).To(Equal(http.StatusBadRequest))
				respBody, err := ioutil.ReadAll(recorder.Result().Body)
				Expect(err).NotTo(HaveOccurred())
				Expect(string(respBody[0 : len(respBody)-1])).To(Equal(api.EmptyBody))
				Expect(userService.DeleteUserCallCount()).To(Equal(0))
			})
		})

		Context("Body of the request contains invalid data", func() {

			BeforeEach(func() {
				userBytes, err := json.Marshal("loginRequest")
				Expect(err).NotTo(HaveOccurred())
				recorder = httptest.NewRecorder()
				request, _ := requestGenerator.CreateRequest(api.Login, nil, bytes.NewReader(userBytes))
				handler.ServeHTTP(recorder, request)
			})

			It("Fails to login a user and returns a 400 status code", func() {
				Expect(recorder.Result().StatusCode).To(Equal(http.StatusBadRequest))
				respBody, err := ioutil.ReadAll(recorder.Result().Body)
				Expect(err).NotTo(HaveOccurred())
				Expect(string(respBody[0 : len(respBody)-1])).To(Equal(api.InvalidBody))
				Expect(userService.DeleteUserCallCount()).To(Equal(0))
			})
		})
	})

	Describe("Find a user", func() {

		var (
			userID     uint
			sampleUser model.User
		)

		BeforeEach(func() {
			sampleUser = model.User{
				ID:        uint(12345),
				UserName:  &userName,
				FirstName: firstName,
				LastName:  lastName,
				Email:     email,
				Password:  password,
				AdminUser: false,
			}
		})

		Context("Valid userName is passed", func() {

			Context("User with userName exists", func() {

				BeforeEach(func() {
					userService.FindUserByUserNameReturns(&sampleUser, nil)
					recorder = httptest.NewRecorder()
					request, _ := requestGenerator.CreateRequest(
						api.FindUserByUserName,
						rata.Params{"user_name": userName},
						nil)
					handler.ServeHTTP(recorder, request)
				})

				It("Finds and returns a user and returns 200 status code", func() {
					Expect(recorder.Result().StatusCode).To(Equal(http.StatusOK))
					respBody, err := ioutil.ReadAll(recorder.Result().Body)
					Expect(err).NotTo(HaveOccurred())

					var foundUser model.User
					err = json.Unmarshal(respBody, &foundUser)
					Expect(err).NotTo(HaveOccurred())
					Expect(foundUser).NotTo(BeNil())
					Expect(foundUser.FirstName).To(Equal(sampleUser.FirstName))
					Expect(foundUser.LastName).To(Equal(sampleUser.LastName))
					Expect(userService.FindUserByUserNameCallCount()).To(Equal(1))
				})
			})

			Context("User with userName does not exist", func() {

				BeforeEach(func() {
					userService.FindUserByUserNameReturns(nil, errors.New("whoot"))
					recorder = httptest.NewRecorder()
					request, _ := requestGenerator.CreateRequest(
						api.FindUserByUserName,
						rata.Params{"user_name": userName},
						nil)
					handler.ServeHTTP(recorder, request)
				})

				It("Doesn't find a user and returns 404 status code", func() {
					Expect(recorder.Result().StatusCode).To(Equal(http.StatusNotFound))
					respBody, err := ioutil.ReadAll(recorder.Result().Body)
					Expect(err).NotTo(HaveOccurred())

					Expect(string(respBody[0 : len(respBody)-1])).To(Equal(api.UnableToFindUser))
					Expect(userService.FindUserByUserNameCallCount()).To(Equal(1))
				})
			})
		})

		Context("Valid userID is passed", func() {

			Context("User with userID exists", func() {

				BeforeEach(func() {
					userService.FindUserByIDReturns(&sampleUser, nil)
					userID = uint(12345)
					recorder = httptest.NewRecorder()
					request, _ := requestGenerator.CreateRequest(
						api.FindUser,
						rata.Params{"user_id": strconv.FormatUint(uint64(userID), 10)},
						nil)
					handler.ServeHTTP(recorder, request)
				})

				It("Finds and returns a user and returns 200 status code", func() {
					Expect(recorder.Result().StatusCode).To(Equal(http.StatusOK))
					respBody, err := ioutil.ReadAll(recorder.Result().Body)
					Expect(err).NotTo(HaveOccurred())

					var foundUser model.User
					err = json.Unmarshal(respBody, &foundUser)
					Expect(err).NotTo(HaveOccurred())
					Expect(foundUser).NotTo(BeNil())
					Expect(foundUser.FirstName).To(Equal(sampleUser.FirstName))
					Expect(foundUser.LastName).To(Equal(sampleUser.LastName))
					Expect(userService.FindUserByIDCallCount()).To(Equal(1))
				})
			})
		})

		Context("Valid search parameter information is passed but downstream call fails", func() {

			BeforeEach(func() {
				userService.FindUserByUserNameReturns(nil, errors.New("whoot"))
				recorder = httptest.NewRecorder()
				request, _ := requestGenerator.CreateRequest(
					api.FindUserByUserName,
					rata.Params{"user_name": userName},
					nil)
				handler.ServeHTTP(recorder, request)
			})

			It("Fails to find a user and returns a 404 status code", func() {
				Expect(recorder.Result().StatusCode).To(Equal(http.StatusNotFound))
				respBody, err := ioutil.ReadAll(recorder.Result().Body)
				Expect(err).NotTo(HaveOccurred())
				Expect(string(respBody[0 : len(respBody)-1])).To(Equal(api.UnableToFindUser))
				Expect(userService.FindUserByUserNameCallCount()).To(Equal(1))
			})
		})
	})
})
