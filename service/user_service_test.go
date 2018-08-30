package service_test

import (
	"context"
	"errors"
	"fmt"

	"github.com/jmelchio/vetlab/model"
	. "github.com/jmelchio/vetlab/service"
	"github.com/jmelchio/vetlab/service/servicefakes"
	"golang.org/x/crypto/bcrypt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("UserService", func() {

	var (
		userService User
		userRepo    *servicefakes.FakeUserRepo
		user        model.User
	)

	BeforeEach(func() {
		userRepo = new(servicefakes.FakeUserRepo)
		userService = User{UserRepo: userRepo}

		user = model.User{
			UserName:     "some-name",
			FirstName:    "first-name",
			LastName:     "last-name",
			Email:        "email@domain.com",
			PasswordHash: "password-hash",
			OrgID:        12345,
			AdminUser:    false,
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
				user.PasswordHash = "uhSeven"
			})

			It("Returns an error and does not call UserRepo.Create", func() {
				zeUser, err := userService.CreateUser(context.TODO(), user)
				Expect(err).To(HaveOccurred())
				Expect(zeUser).To(BeNil())
				Expect(err.Error()).To(Equal(PasswordTooShort))
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
				userRepo.CreateReturns(errors.New("Unable to create user"))
			})

			It("Returns an error after calling UserRepo.Create", func() {
				zeUser, err := userService.CreateUser(context.TODO(), user)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("Unable to create user"))
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
				userRepo.UpdateReturns(errors.New("Unable to update user"))
			})

			It("Returns an error after calling UserRepo.Create", func() {
				zeUser, err := userService.UpdateUser(context.TODO(), user)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("Unable to update user"))
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
				userRepo.DeleteReturns(errors.New("Unable to delete user"))
			})

			It("Returns an error after calling UserRepo.Create", func() {
				err := userService.DeleteUser(context.TODO(), user)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("Unable to delete user"))
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
				err = bcrypt.CompareHashAndPassword([]byte(updatedUser.PasswordHash), []byte(newPwd))
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
				userRepo.UpdateReturns(errors.New("Unable to update user"))
			})

			It("returns an error", func() {
				updatedUser, err := userService.UpdatePassword(context.TODO(), user, newPwd)
				Expect(err).To(HaveOccurred())
				Expect(updatedUser).To(BeNil())
				Expect(err.Error()).To(Equal("Unable to update user"))
				Expect(userRepo.UpdateCallCount()).To(Equal(1))
			})
		})

		Context("We have a valid user, a short password and context", func() {

			BeforeEach(func() {
				newPwd = "uhseven"
				userRepo.UpdateReturns(errors.New("Unable to update user"))
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
			userName   string
			password   string
			sampleUser model.User
		)

		BeforeEach(func() {
			userName = "user"
			password = "somepwdd"
		})

		Context("Username, password and context are all correct", func() {

			BeforeEach(func() {
				passwordHash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
				sampleUser = model.User{
					ID:           12345,
					UserName:     userName,
					PasswordHash: string(passwordHash),
				}
				userRepo.GetByUserNameReturns(&sampleUser, nil)
			})

			It("Returns an authenticated user", func() {
				user, err := userService.Login(context.TODO(), userName, password)
				Expect(err).NotTo(HaveOccurred())
				Expect(user).NotTo(BeNil())
				Expect(user.ID).To(Equal(sampleUser.ID))
				Expect(userRepo.GetByUserNameCallCount()).To(Equal(1))
			})
		})

		Context("No context provided", func() {

			It("Returns an error", func() {
				user, err := userService.Login(nil, userName, password)
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
				passwordHash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
				sampleUser = model.User{
					ID:           12345,
					UserName:     userName,
					PasswordHash: string(passwordHash),
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

	Describe("Find a user by veterinary practice", func() {

		var (
			vetOrg model.VetOrg
			user   model.User
		)

		BeforeEach(func() {
			vetOrg = model.VetOrg{
				ID: 12345,
			}

			user = model.User{
				ID:        12345,
				UserName:  "some-user-name",
				FirstName: "john",
				LastName:  "doe",
				OrgID:     12345,
			}

			userRepo.GetByOrgIDReturns([]model.User{user}, nil)
		})

		Context("Context and VetOrg are correct", func() {

			It("Returns a list of users that are part of the given VetOrg", func() {
				users, err := userService.FindUsersByVetOrg(context.TODO(), vetOrg)
				Expect(err).NotTo(HaveOccurred())
				Expect(users).NotTo(BeNil())
				Expect(users).To(HaveLen(1))
				Expect(userRepo.GetByOrgIDCallCount()).To(Equal(1))
				Expect(users[0].OrgID).To(Equal(vetOrg.ID))
			})
		})

		Context("VetOrg provided but Context missing", func() {

			It("Returns and error", func() {
				users, err := userService.FindUsersByVetOrg(nil, vetOrg)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal(MissingContext))
				Expect(users).To(BeNil())
			})
		})

		Context("Context provided but VetOrg lacks ID", func() {

			It("Returns an error", func() {
				users, err := userService.FindUsersByVetOrg(context.TODO(), model.VetOrg{})
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal(VetOrgRequired))
				Expect(users).To(BeNil())
			})
		})

		Context("Context and Vetorg are correct but Repo returns error", func() {

			BeforeEach(func() {
				userRepo.GetByOrgIDReturns(nil, errors.New("BAM!"))
			})

			It("Returns an error", func() {
				users, err := userService.FindUsersByVetOrg(context.TODO(), vetOrg)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("BAM!"))
				Expect(users).To(BeNil())
			})
		})
	})

	Describe("Find a user by username", func() {

		var (
			user     model.User
			userName string
		)

		BeforeEach(func() {
			user = model.User{
				ID:        12345,
				UserName:  "some-user-name",
				FirstName: "john",
				LastName:  "doe",
				OrgID:     12345,
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
				Expect(foundUser.UserName).To(Equal(userName))
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
				userRepo.GetByUserNameReturns(nil, errors.New("BAM!"))
			})

			It("Returns an error", func() {
				foundUser, err := userService.FindUserByUserName(context.TODO(), userName)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("BAM!"))
				Expect(foundUser).To(BeNil())
			})
		})
	})

	Describe("Find a user by ID", func() {

		var (
			user   model.User
			userID uint
		)

		BeforeEach(func() {
			user = model.User{
				ID:        12345,
				UserName:  "some-user-name",
				FirstName: "john",
				LastName:  "doe",
				OrgID:     12345,
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
				userRepo.GetByIDReturns(nil, errors.New("BAM!"))
			})

			It("Returns an error", func() {
				foundUser, err := userService.FindUserByID(context.TODO(), userID)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("BAM!"))
				Expect(foundUser).To(BeNil())
			})
		})
	})
})
