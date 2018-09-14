package api_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"

	"github.com/jmelchio/vetlab/api/apifakes"
	"github.com/jmelchio/vetlab/model"
	"github.com/tedsuo/rata"

	. "github.com/jmelchio/vetlab/api"

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
	)

	BeforeEach(func() {
		userService = new(apifakes.FakeUserService)
		handler, err = NewUserHandler(userService)
		requestGenerator = rata.NewRequestGenerator("", UserRoutes)
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
				OrgID:        uint(12345),
				AdminUser:    false,
			}
		})

		Context("Valid user information is passed", func() {

			BeforeEach(func() {
				userService.CreateUserReturns(&createUser, nil)
				userBytes, err := json.Marshal(createUser)
				Expect(err).NotTo(HaveOccurred())
				recorder = httptest.NewRecorder()
				request, _ := requestGenerator.CreateRequest(CreateUser, nil, bytes.NewReader(userBytes))
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
				request, _ := requestGenerator.CreateRequest(CreateUser, nil, bytes.NewReader(userBytes))
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
				request, _ := requestGenerator.CreateRequest(CreateUser, nil, nil)
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
				request, _ := requestGenerator.CreateRequest(CreateUser, nil, bytes.NewReader(userBytes))
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
		var (
			updateUser model.User
		)

		BeforeEach(func() {
			updateUser = model.User{
				ID:           uint(12345),
				UserName:     "user_name",
				FirstName:    "first_name",
				LastName:     "last_name",
				Email:        "some_email",
				PasswordHash: "some_hash",
				OrgID:        uint(12345),
				AdminUser:    false,
			}
		})

		Context("Valid user information is passed", func() {

			BeforeEach(func() {
				userService.UpdateUserReturns(&updateUser, nil)
				userBytes, err := json.Marshal(updateUser)
				Expect(err).NotTo(HaveOccurred())
				recorder = httptest.NewRecorder()
				request, _ := requestGenerator.CreateRequest(UpdateUser, nil, bytes.NewReader(userBytes))
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
				userService.UpdateUserReturns(nil, errors.New("Whoot?"))
				userBytes, err := json.Marshal(updateUser)
				Expect(err).NotTo(HaveOccurred())
				recorder = httptest.NewRecorder()
				request, _ := requestGenerator.CreateRequest(UpdateUser, nil, bytes.NewReader(userBytes))
				handler.ServeHTTP(recorder, request)
			})

			It("Fails to return a users and returns 500 status code", func() {
				Expect(recorder.Result().StatusCode).To(Equal(http.StatusInternalServerError))
				respBody, err := ioutil.ReadAll(recorder.Result().Body)
				Expect(err).NotTo(HaveOccurred())
				Expect(err).NotTo(HaveOccurred())
				Expect(string(respBody[0 : len(respBody)-1])).To(Equal(UnableToUpdateUser))
				Expect(userService.UpdateUserCallCount()).To(Equal(1))
			})
		})

		Context("Body of the request is empty", func() {

			BeforeEach(func() {
				recorder = httptest.NewRecorder()
				request, _ := requestGenerator.CreateRequest(UpdateUser, nil, nil)
				handler.ServeHTTP(recorder, request)
			})

			It("Fails to update a user and returns a 400 status code", func() {
				Expect(recorder.Result().StatusCode).To(Equal(http.StatusBadRequest))
				respBody, err := ioutil.ReadAll(recorder.Result().Body)
				Expect(err).NotTo(HaveOccurred())
				Expect(string(respBody[0 : len(respBody)-1])).To(Equal(EmptyBody))
				Expect(userService.UpdateUserCallCount()).To(Equal(0))
			})
		})

		Context("Body of the request contains invalid data", func() {

			BeforeEach(func() {
				userBytes, err := json.Marshal("updateUser")
				Expect(err).NotTo(HaveOccurred())
				recorder = httptest.NewRecorder()
				request, _ := requestGenerator.CreateRequest(UpdateUser, nil, bytes.NewReader(userBytes))
				handler.ServeHTTP(recorder, request)
			})

			It("Fails to update a user and returns a 400 status code", func() {
				Expect(recorder.Result().StatusCode).To(Equal(http.StatusBadRequest))
				respBody, err := ioutil.ReadAll(recorder.Result().Body)
				Expect(err).NotTo(HaveOccurred())
				Expect(string(respBody[0 : len(respBody)-1])).To(Equal(InvalidBody))
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
				ID:           uint(12345),
				UserName:     "user_name",
				FirstName:    "first_name",
				LastName:     "last_name",
				Email:        "some_email",
				PasswordHash: "some_hash",
				OrgID:        uint(12345),
				AdminUser:    false,
			}
		})

		Context("Valid user information is passed", func() {

			BeforeEach(func() {
				userService.DeleteUserReturns(nil)
				userBytes, err := json.Marshal(deleteUser)
				Expect(err).NotTo(HaveOccurred())
				recorder = httptest.NewRecorder()
				request, _ := requestGenerator.CreateRequest(DeleteUser, nil, bytes.NewReader(userBytes))
				handler.ServeHTTP(recorder, request)
			})

			It("Deletes a user and returns 204 status code", func() {
				Expect(recorder.Result().StatusCode).To(Equal(http.StatusNoContent))
				Expect(userService.DeleteUserCallCount()).To(Equal(1))
			})
		})

		Context("Valid user information is passed but downstream call fails", func() {

			BeforeEach(func() {
				userService.DeleteUserReturns(errors.New("Whoot?"))
				userBytes, err := json.Marshal(deleteUser)
				Expect(err).NotTo(HaveOccurred())
				recorder = httptest.NewRecorder()
				request, _ := requestGenerator.CreateRequest(DeleteUser, nil, bytes.NewReader(userBytes))
				handler.ServeHTTP(recorder, request)
			})

			It("Fails to delete a user and returns a 500 status code", func() {
				Expect(recorder.Result().StatusCode).To(Equal(http.StatusInternalServerError))
				respBody, err := ioutil.ReadAll(recorder.Result().Body)
				Expect(err).NotTo(HaveOccurred())
				Expect(string(respBody[0 : len(respBody)-1])).To(Equal(UnableToDeleteUser))
				Expect(userService.DeleteUserCallCount()).To(Equal(1))
			})
		})

		Context("Body of the request is empty", func() {

			BeforeEach(func() {
				recorder = httptest.NewRecorder()
				request, _ := requestGenerator.CreateRequest(DeleteUser, nil, nil)
				handler.ServeHTTP(recorder, request)
			})

			It("Fails to delete a user and returns a 400 status code", func() {
				Expect(recorder.Result().StatusCode).To(Equal(http.StatusBadRequest))
				respBody, err := ioutil.ReadAll(recorder.Result().Body)
				Expect(err).NotTo(HaveOccurred())
				Expect(string(respBody[0 : len(respBody)-1])).To(Equal(EmptyBody))
				Expect(userService.DeleteUserCallCount()).To(Equal(0))
			})
		})

		Context("Body of the request contains invalid data", func() {

			BeforeEach(func() {
				userBytes, err := json.Marshal("deleteUser")
				Expect(err).NotTo(HaveOccurred())
				recorder = httptest.NewRecorder()
				request, _ := requestGenerator.CreateRequest(DeleteUser, nil, bytes.NewReader(userBytes))
				handler.ServeHTTP(recorder, request)
			})

			It("Fails to delete a user and returns a 400 status code", func() {
				Expect(recorder.Result().StatusCode).To(Equal(http.StatusBadRequest))
				respBody, err := ioutil.ReadAll(recorder.Result().Body)
				Expect(err).NotTo(HaveOccurred())
				Expect(string(respBody[0 : len(respBody)-1])).To(Equal(InvalidBody))
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
				ID:           uint(12345),
				UserName:     "user_name",
				FirstName:    "first_name",
				LastName:     "lastName",
				Email:        "some_email",
				PasswordHash: "some_hash",
				OrgID:        uint(12345),
				AdminUser:    false,
			}
		})

		Context("Valid user information is passed", func() {

			BeforeEach(func() {
				userService.LoginReturns(&loginUser, nil)
				userBytes, err := json.Marshal(loginRequest)
				Expect(err).NotTo(HaveOccurred())
				recorder = httptest.NewRecorder()
				request, _ := requestGenerator.CreateRequest(Login, nil, bytes.NewReader(userBytes))
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
				userService.LoginReturns(nil, errors.New("Whoot?"))
				userBytes, err := json.Marshal(loginRequest)
				Expect(err).NotTo(HaveOccurred())
				recorder = httptest.NewRecorder()
				request, _ := requestGenerator.CreateRequest(Login, nil, bytes.NewReader(userBytes))
				handler.ServeHTTP(recorder, request)
			})

			It("Fails to login a user and returns a 500 status code", func() {
				Expect(recorder.Result().StatusCode).To(Equal(http.StatusInternalServerError))
				respBody, err := ioutil.ReadAll(recorder.Result().Body)
				Expect(err).NotTo(HaveOccurred())
				Expect(string(respBody[0 : len(respBody)-1])).To(Equal(UnableToLoginUser))
				Expect(userService.LoginCallCount()).To(Equal(1))
			})
		})

		Context("Body of the request is empty", func() {

			BeforeEach(func() {
				recorder = httptest.NewRecorder()
				request, _ := requestGenerator.CreateRequest(Login, nil, nil)
				handler.ServeHTTP(recorder, request)
			})

			It("Fails to login a user and returns a 400 status code", func() {
				Expect(recorder.Result().StatusCode).To(Equal(http.StatusBadRequest))
				respBody, err := ioutil.ReadAll(recorder.Result().Body)
				Expect(err).NotTo(HaveOccurred())
				Expect(string(respBody[0 : len(respBody)-1])).To(Equal(EmptyBody))
				Expect(userService.DeleteUserCallCount()).To(Equal(0))
			})
		})

		Context("Body of the request contains invalid data", func() {

			BeforeEach(func() {
				userBytes, err := json.Marshal("loginRequest")
				Expect(err).NotTo(HaveOccurred())
				recorder = httptest.NewRecorder()
				request, _ := requestGenerator.CreateRequest(Login, nil, bytes.NewReader(userBytes))
				handler.ServeHTTP(recorder, request)
			})

			It("Fails to login a user and returns a 400 status code", func() {
				Expect(recorder.Result().StatusCode).To(Equal(http.StatusBadRequest))
				respBody, err := ioutil.ReadAll(recorder.Result().Body)
				Expect(err).NotTo(HaveOccurred())
				Expect(string(respBody[0 : len(respBody)-1])).To(Equal(InvalidBody))
				Expect(userService.DeleteUserCallCount()).To(Equal(0))
			})
		})
	})

	Describe("Find a user", func() {

		var (
			userName   string
			userID     uint
			vetOrgID   uint
			sampleUser model.User
		)

		BeforeEach(func() {
			sampleUser = model.User{
				ID:           uint(12345),
				UserName:     "user_name",
				FirstName:    "first_name",
				LastName:     "lastName",
				Email:        "some_email",
				PasswordHash: "some_hash",
				OrgID:        uint(12345),
				AdminUser:    false,
			}
		})

		Context("Valid userName is passed", func() {

			Context("User with userName exists", func() {

				BeforeEach(func() {
					userService.FindUserByUserNameReturns(&sampleUser, nil)
					userName = "user_name"
					recorder = httptest.NewRecorder()
					request, _ := requestGenerator.CreateRequest(FindUser, nil, nil)
					// request, _ := http.NewRequest("GET", "/user/find", nil)
					q := url.Values{}
					q.Add("user_name", userName)
					request.URL.RawQuery = q.Encode()
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
					userService.FindUserByUserNameReturns(nil, errors.New("Whoot?"))
					userName = "user_name"
					recorder = httptest.NewRecorder()
					request, _ := requestGenerator.CreateRequest(FindUser, nil, nil)
					q := url.Values{}
					q.Add("user_name", userName)
					request.URL.RawQuery = q.Encode()
					handler.ServeHTTP(recorder, request)
				})

				It("Doesn't find a user and returns 404 status code", func() {
					Expect(recorder.Result().StatusCode).To(Equal(http.StatusNotFound))
					respBody, err := ioutil.ReadAll(recorder.Result().Body)
					Expect(err).NotTo(HaveOccurred())

					Expect(string(respBody[0 : len(respBody)-1])).To(Equal(UnableToFindUser))
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
					request, _ := requestGenerator.CreateRequest(FindUser, nil, nil)
					q := url.Values{}
					q.Add("user_id", fmt.Sprint(userID))
					request.URL.RawQuery = q.Encode()
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

		Context("Valid vetOrgID is passed", func() {

			Context("User with vetOrgID exists", func() {

				BeforeEach(func() {
					userSlice := []model.User{sampleUser}
					userService.FindUsersByVetOrgIDReturns(userSlice, nil)
					vetOrgID = uint(12345)
					recorder = httptest.NewRecorder()
					request, _ := requestGenerator.CreateRequest(FindUser, nil, nil)
					q := url.Values{}
					q.Add("vet_org_id", fmt.Sprint(vetOrgID))
					request.URL.RawQuery = q.Encode()
					handler.ServeHTTP(recorder, request)
				})

				It("Finds and returns a slice of user and returns 200 status code", func() {
					Expect(recorder.Result().StatusCode).To(Equal(http.StatusOK))
					respBody, err := ioutil.ReadAll(recorder.Result().Body)
					Expect(err).NotTo(HaveOccurred())

					var foundUsers []model.User
					err = json.Unmarshal(respBody, &foundUsers)
					Expect(err).NotTo(HaveOccurred())
					Expect(foundUsers).NotTo(BeNil())
					Expect(foundUsers[0].FirstName).To(Equal(sampleUser.FirstName))
					Expect(foundUsers[0].LastName).To(Equal(sampleUser.LastName))
					Expect(userService.FindUsersByVetOrgIDCallCount()).To(Equal(1))
				})
			})
		})

		Context("Valid search parameter information is passed but downstream call fails", func() {

			BeforeEach(func() {
				userService.FindUserByUserNameReturns(nil, errors.New("Whoot?"))
				userName = "user_name"
				recorder = httptest.NewRecorder()
				request, _ := requestGenerator.CreateRequest(FindUser, nil, nil)
				q := url.Values{}
				q.Add("user_name", userName)
				request.URL.RawQuery = q.Encode()
				handler.ServeHTTP(recorder, request)
			})

			It("Fails to find a user and returns a 404 status code", func() {
				Expect(recorder.Result().StatusCode).To(Equal(http.StatusNotFound))
				respBody, err := ioutil.ReadAll(recorder.Result().Body)
				Expect(err).NotTo(HaveOccurred())
				Expect(string(respBody[0 : len(respBody)-1])).To(Equal(UnableToFindUser))
				Expect(userService.FindUserByUserNameCallCount()).To(Equal(1))
			})
		})

		Context("No request parameters are passed", func() {

			BeforeEach(func() {
				recorder = httptest.NewRecorder()
				request, _ := requestGenerator.CreateRequest(FindUser, nil, nil)
				handler.ServeHTTP(recorder, request)
			})

			It("Fails to find a user and returns a 400 status code", func() {
				Expect(recorder.Result().StatusCode).To(Equal(http.StatusBadRequest))
				respBody, err := ioutil.ReadAll(recorder.Result().Body)
				Expect(err).NotTo(HaveOccurred())
				Expect(string(respBody[0 : len(respBody)-1])).To(Equal(NoParamsFound))
				Expect(userService.FindUserByUserNameCallCount()).To(Equal(0))
				Expect(userService.FindUserByIDCallCount()).To(Equal(0))
				Expect(userService.FindUsersByVetOrgIDCallCount()).To(Equal(0))
			})
		})

		Context("Bad request parameters are passed", func() {

			BeforeEach(func() {
				recorder = httptest.NewRecorder()
				request, _ := requestGenerator.CreateRequest(FindUser, nil, nil)
				q := url.Values{}
				q.Add("first_name", "nobody")
				request.URL.RawQuery = q.Encode()
				handler.ServeHTTP(recorder, request)
			})

			It("Fails to find a user and returns a 400 status code", func() {
				Expect(recorder.Result().StatusCode).To(Equal(http.StatusBadRequest))
				respBody, err := ioutil.ReadAll(recorder.Result().Body)
				Expect(err).NotTo(HaveOccurred())
				Expect(string(respBody[0 : len(respBody)-1])).To(Equal(NoParamsFound))
				Expect(userService.FindUserByUserNameCallCount()).To(Equal(0))
				Expect(userService.FindUserByIDCallCount()).To(Equal(0))
				Expect(userService.FindUsersByVetOrgIDCallCount()).To(Equal(0))
			})
		})
	})
})
