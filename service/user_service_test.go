package service_test

import (
	"context"
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"

	"github.com/jmelchio/vetlab/api"
	"github.com/jmelchio/vetlab/model"
	. "github.com/jmelchio/vetlab/service"
	"github.com/jmelchio/vetlab/service/servicefakes"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("UserService", func() {

	var (
		userService api.UserService
		userRepo    *servicefakes.FakeUserRepo
		user        model.User
		userName    string
		firstName   string
		lastName    string
		email       string
		password    string
	)

	BeforeEach(func() {
		userRepo = new(servicefakes.FakeUserRepo)
		userService = User{UserRepo: userRepo}

		userName = "some-name"
		firstName = "first-name"
		lastName = "last-name"
		email = "email@domain.com"
		password = "password-hash"

		user = model.User{
			UserName:  &userName,
			FirstName: firstName,
			LastName:  lastName,
			Email:     email,
			Password:  password,
			AdminUser: false,
		}
	})

	Describe("Create a user", func() {

		BeforeEach(func() {
			userRepo.CreateReturns(nil)
		})

		Context("We have a valid user and 'todo' context", func() {

			It("Returns a user with a new user ID and calls UserRepo.Create", func() {
				zeUser, err := userService.CreateUser(context.TODO(), user)
				Expect(err).ToNot(HaveOccurred())
				Expect(zeUser).NotTo(BeNil())
				Expect(userRepo.CreateCallCount()).To(Equal(1))
			})
		})

		Context("User password is too short but we have a valid context", func() {

			BeforeEach(func() {
				user.Password = "uhSeven"
			})

			It("Returns an error and does not call UserRepo.Create", func() {
				zeUser, err := userService.CreateUser(context.TODO(), user)
				Expect(err).To(HaveOccurred())
				Expect(zeUser).To(BeNil())
				Expect(err.Error()).To(ContainSubstring(PasswordTooShort))
				Expect(userRepo.CreateCallCount()).To(Equal(0))
			})
		})

		Context("We have a valid user but no context", func() {

			It("Returns an error and does not call UserRepo.Create", func() {
				zeUser, err := userService.CreateUser(nil, user)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal(MissingContext))
				Expect(zeUser).To(BeNil())
				Expect(userRepo.CreateCallCount()).To(Equal(0))
			})
		})

		Context("We have a user and Context but repo cannot create user", func() {

			BeforeEach(func() {
				userRepo.CreateReturns(errors.New("unable to create user"))
			})

			It("Returns an error after calling UserRepo.Create", func() {
				zeUser, err := userService.CreateUser(context.TODO(), user)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("unable to create user"))
				Expect(zeUser).To(BeNil())
				Expect(userRepo.CreateCallCount()).To(Equal(1))
			})
		})
	})

	Describe("Update a user", func() {

		BeforeEach(func() {
			userRepo.UpdateReturns(nil)
		})

		Context("We have a valid user and context", func() {

			It("Returns the updated user and no error", func() {
				zeUser, err := userService.UpdateUser(context.TODO(), user)
				Expect(err).NotTo(HaveOccurred())
				Expect(userRepo.UpdateCallCount()).To(Equal(1))
				Expect(zeUser).NotTo(BeNil())
			})
		})

		Context("We have a valid user but no context", func() {

			It("Returns and error and no updated user", func() {
				zeUser, err := userService.UpdateUser(nil, user)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal(MissingContext))
				Expect(zeUser).To(BeNil())
				Expect(userRepo.UpdateCallCount()).To(Equal(0))
			})
		})

		Context("We have a user and Context but repo cannot create user", func() {

			BeforeEach(func() {
				userRepo.UpdateReturns(errors.New("unable to update user"))
			})

			It("Returns an error after calling UserRepo.Create", func() {
				zeUser, err := userService.UpdateUser(context.TODO(), user)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("unable to update user"))
				Expect(zeUser).To(BeNil())
				Expect(userRepo.UpdateCallCount()).To(Equal(1))
			})
		})
	})

	Describe("Delete a user", func() {

		BeforeEach(func() {
			userRepo.DeleteReturns(nil)
		})

		Context("We have a valid user and context", func() {

			It("Returns no error", func() {
				err := userService.DeleteUser(context.TODO(), user)
				Expect(err).NotTo(HaveOccurred())
				Expect(userRepo.DeleteCallCount()).To(Equal(1))
			})
		})

		Context("We have a valid user but no context", func() {

			It("Returns and error and no updated user", func() {
				err := userService.DeleteUser(nil, user)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal(MissingContext))
				Expect(userRepo.DeleteCallCount()).To(Equal(0))
			})
		})

		Context("We have a user and Context but repo cannot delete user", func() {

			BeforeEach(func() {
				userRepo.DeleteReturns(errors.New("unable to delete user"))
			})

			It("Returns an error after calling UserRepo.Create", func() {
				err := userService.DeleteUser(context.TODO(), user)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("unable to delete user"))
				Expect(userRepo.DeleteCallCount()).To(Equal(1))
			})
		})
	})

	Describe("UpdatePassword", func() {

		var (
			newPwd string
		)

		BeforeEach(func() {
			newPwd = "some-magic-password"
			userRepo.UpdateReturns(nil)
		})

		Context("We have a valid user, password and context", func() {

			It("It returns no error", func() {
				updatedUser, err := userService.UpdatePassword(context.TODO(), user, newPwd)
				Expect(err).ToNot(HaveOccurred())
				Expect(userRepo.UpdateCallCount()).To(Equal(1))
				err = bcrypt.CompareHashAndPassword([]byte(updatedUser.Password), []byte(newPwd))
				Expect(err).ToNot(HaveOccurred())
			})
		})

		Context("We have a valid user, password but no context", func() {

			It("Returns an error", func() {
				updatedUser, err := userService.UpdatePassword(nil, user, newPwd)
				Expect(err).To(HaveOccurred())
				Expect(updatedUser).To(BeNil())
				Expect(err.Error()).To(Equal(MissingContext))
				Expect(userRepo.UpdateCallCount()).To(Equal(0))
			})
		})

		Context("We have a valid user, password and context but update fails", func() {

			BeforeEach(func() {
				userRepo.UpdateReturns(errors.New("unable to update user"))
			})

			It("returns an error", func() {
				updatedUser, err := userService.UpdatePassword(context.TODO(), user, newPwd)
				Expect(err).To(HaveOccurred())
				Expect(updatedUser).To(BeNil())
				Expect(err.Error()).To(Equal("unable to update user"))
				Expect(userRepo.UpdateCallCount()).To(Equal(1))
			})
		})

		Context("We have a valid user, a short password and context", func() {

			BeforeEach(func() {
				newPwd = "uhseven"
				userRepo.UpdateReturns(errors.New("unable to update user"))
			})

			It("returns an error", func() {
				updatedUser, err := userService.UpdatePassword(context.TODO(), user, newPwd)
				Expect(err).To(HaveOccurred())
				Expect(updatedUser).To(BeNil())
				Expect(err.Error()).To(Equal(fmt.Sprintf(HashingFailed, PasswordTooShort)))
				Expect(userRepo.UpdateCallCount()).To(Equal(0))
			})
		})
	})

	Describe("Login", func() {

		var (
			luserName  string
			password   string
			sampleUser model.User
		)

		BeforeEach(func() {
			luserName = "user"
			password = "somepwdd"
		})

		Context("Username, password and context are all correct", func() {

			BeforeEach(func() {
				password, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
				stringPassword := string(password)
				sampleUser = model.User{
					ID:       uint(12345),
					UserName: &luserName,
					Password: stringPassword,
				}
				userRepo.GetByUserNameReturns(&sampleUser, nil)
			})

			It("Returns an authenticated user", func() {
				user, err := userService.Login(context.TODO(), luserName, password)
				Expect(err).NotTo(HaveOccurred())
				Expect(user).NotTo(BeNil())
				Expect(user.ID).To(Equal(sampleUser.ID))
				Expect(userRepo.GetByUserNameCallCount()).To(Equal(1))
			})
		})

		Context("No context provided", func() {

			It("Returns an error", func() {
				user, err := userService.Login(nil, luserName, password)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal(MissingContext))
				Expect(user).To(BeNil())
				Expect(userRepo.GetByUserNameCallCount()).To(Equal(0))
			})
		})

		Context("Username or password are not provided", func() {

			It("Returns an error", func() {
				user, err := userService.Login(context.TODO(), "", "")
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal(UserOrPasswordFail))
				Expect(user).To(BeNil())
				Expect(userRepo.GetByUserNameCallCount()).To(Equal(0))
			})
		})

		Context("Wrong password is provided", func() {

			BeforeEach(func() {
				password, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
				stringPassword := string(password)
				sampleUser = model.User{
					ID:       uint(12345),
					UserName: &luserName,
					Password: stringPassword,
				}
				userRepo.GetByUserNameReturns(&sampleUser, nil)
			})

			It("Returns an error", func() {
				user, err := userService.Login(context.TODO(), userName, "some-other-pwd")
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal(UserOrPasswordFail))
				Expect(user).To(BeNil())
				Expect(userRepo.GetByUserNameCallCount()).To(Equal(1))
			})
		})

		Context("Underlying repository returns an error", func() {

			BeforeEach(func() {
				userRepo.GetByUserNameReturns(nil, errors.New("some-error"))
			})

			It("Returns an error", func() {
				user, err := userService.Login(context.TODO(), userName, password)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal(UserOrPasswordFail))
				Expect(user).To(BeNil())
				Expect(userRepo.GetByUserNameCallCount()).To(Equal(1))
			})
		})
	})

	Describe("Find a user by username", func() {

		var (
			user       model.User
			userName   string
			fUserName  string
			fFirstName string
			fLastName  string
		)

		BeforeEach(func() {
			fUserName = "some-user-name"
			fFirstName = "john"
			fLastName = "doe"

			user = model.User{
				UserName:  &fUserName,
				FirstName: fFirstName,
				LastName:  fLastName,
			}

			userName = "some-user-name"

			userRepo.GetByUserNameReturns(&user, nil)
		})

		Context("Context is correct and username exists", func() {

			It("Returns a user with the given username", func() {
				foundUser, err := userService.FindUserByUserName(context.TODO(), userName)
				Expect(err).NotTo(HaveOccurred())
				Expect(user).NotTo(BeNil())
				Expect(userRepo.GetByUserNameCallCount()).To(Equal(1))
				Expect(*foundUser.UserName).To(Equal(userName))
			})
		})

		Context("UserName provided but Context missing", func() {

			It("Returns and error", func() {
				foundUser, err := userService.FindUserByUserName(nil, userName)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal(MissingContext))
				Expect(foundUser).To(BeNil())
			})
		})

		Context("Context and userName provided but user is not found", func() {

			BeforeEach(func() {
				userRepo.GetByUserNameReturns(nil, nil)
			})

			It("Returns no error and no user", func() {
				foundUser, err := userService.FindUserByUserName(context.TODO(), "john doe")
				Expect(err).ToNot(HaveOccurred())
				Expect(foundUser).To(BeNil())
			})
		})

		Context("Context and userName are correct but Repo returns error", func() {

			BeforeEach(func() {
				userRepo.GetByUserNameReturns(nil, errors.New("BAM"))
			})

			It("Returns an error", func() {
				foundUser, err := userService.FindUserByUserName(context.TODO(), userName)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("BAM"))
				Expect(foundUser).To(BeNil())
			})
		})
	})

	Describe("Find a user by ID", func() {

		var (
			user       model.User
			userID     uint
			fUserName  string
			fFirstName string
			fLastName  string
		)

		BeforeEach(func() {
			fUserName = "some-user-name"
			fFirstName = "john"
			fLastName = "doe"

			user = model.User{
				ID:        uint(12345),
				UserName:  &fUserName,
				FirstName: fFirstName,
				LastName:  fLastName,
			}

			userID = 12345

			userRepo.GetByIDReturns(&user, nil)
		})

		Context("Context is correct and userID exists", func() {

			It("Returns a user with the given userID", func() {
				foundUser, err := userService.FindUserByID(context.TODO(), userID)
				Expect(err).NotTo(HaveOccurred())
				Expect(user).NotTo(BeNil())
				Expect(userRepo.GetByIDCallCount()).To(Equal(1))
				Expect(foundUser.ID).To(Equal(userID))
			})
		})

		Context("UserName provided but Context missing", func() {

			It("Returns and error", func() {
				foundUser, err := userService.FindUserByID(nil, userID)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal(MissingContext))
				Expect(foundUser).To(BeNil())
			})

		})

		Context("Context and userID provided but user is not found", func() {

			BeforeEach(func() {
				userRepo.GetByIDReturns(nil, nil)
			})

			It("Returns no error and no user", func() {
				foundUser, err := userService.FindUserByID(context.TODO(), 12345)
				Expect(err).ToNot(HaveOccurred())
				Expect(foundUser).To(BeNil())
			})
		})

		Context("Context and userID are correct but Repo returns error", func() {

			BeforeEach(func() {
				userRepo.GetByIDReturns(nil, errors.New("BAM"))
			})

			It("Returns an error", func() {
				foundUser, err := userService.FindUserByID(context.TODO(), userID)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("BAM"))
				Expect(foundUser).To(BeNil())
			})
		})
	})
})
