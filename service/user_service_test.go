package service_test

import (
	"context"
	"errors"

	"github.com/jmelchio/vetlab/model"
	. "github.com/jmelchio/vetlab/service"
	"github.com/jmelchio/vetlab/service/servicefakes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("UserService", func() {
	var (
		userService User
		userRepo    *servicefakes.FakeUserRepo
	)
	BeforeEach(func() {
		userRepo = new(servicefakes.FakeUserRepo)
		userService = User{UserRepo: userRepo}
	})
	Describe("Create a user", func() {
		var (
			newUser     model.User
			createdUser model.User
		)
		BeforeEach(func() {

			newUser = model.User{
				UserName:     "some-name",
				FirstName:    "first-name",
				LastName:     "last-name",
				Email:        "email@domain.com",
				PasswordHash: "passord-hash",
				OrgID:        "org-id",
				AdminUser:    false,
			}
			createdUser = newUser
			createdUser.UserID = "created-user-id"

			userRepo.CreateReturns(&createdUser, nil)
		})
		Context("We have a valid user and 'todo' context", func() {
			It("Returns a user with a new user ID and calls UserRepo.Create", func() {
				zeUser, err := userService.CreateUser(context.TODO(), newUser)
				Expect(err).ToNot(HaveOccurred())
				Expect(zeUser.UserID).To(Equal("created-user-id"))
				Expect(userRepo.CreateCallCount()).To(Equal(1))
			})
		})
		Context("We have a valid user but no context", func() {
			It("Returns an error and does not call UserRepo.Create", func() {
				zeUser, err := userService.CreateUser(nil, newUser)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal(MissingContext))
				Expect(zeUser).To(BeNil())
				Expect(userRepo.CreateCallCount()).To(Equal(0))
			})
		})
		Context("We have a user and Context but repo cannot create user", func() {
			BeforeEach(func() {
				userRepo.CreateReturns(nil, errors.New("Unable to create user"))
			})
			It("Returns an error after calling UserRepo.Create", func() {
				zeUser, err := userService.CreateUser(context.TODO(), newUser)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("Unable to create user"))
				Expect(zeUser).To(BeNil())
				Expect(userRepo.CreateCallCount()).To(Equal(1))
			})
		})
	})
	Describe("Update a user", func() {
		var (
			user        model.User
			updatedUser model.User
		)
		BeforeEach(func() {

			user = model.User{
				UserID:       "some-user-id",
				UserName:     "some-name",
				FirstName:    "first-name",
				LastName:     "last-name",
				Email:        "email@domain.com",
				PasswordHash: "passord-hash",
				OrgID:        "org-id",
				AdminUser:    false,
			}
			updatedUser = user
			updatedUser.UserName = "updated-username"
			updatedUser.FirstName = "second-first-name"

			userRepo.UpdateReturns(&updatedUser, nil)
		})
		Context("We have a valid user and context", func() {
			It("Returns the updated user and no error", func() {
				zeUser, err := userService.UpdateUser(context.TODO(), user)
				Expect(err).NotTo(HaveOccurred())
				Expect(userRepo.UpdateCallCount()).To(Equal(1))
				Expect(zeUser.FirstName).To(Equal("second-first-name"))
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
				userRepo.UpdateReturns(nil, errors.New("Unable to update user"))
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
		var (
			user model.User
		)
		BeforeEach(func() {

			user = model.User{
				UserID:       "some-user-id",
				UserName:     "some-name",
				FirstName:    "first-name",
				LastName:     "last-name",
				Email:        "email@domain.com",
				PasswordHash: "passord-hash",
				OrgID:        "org-id",
				AdminUser:    false,
			}

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
	Describe("Login", func() {
		It("proper test cases and implementation of code", func() {
			Expect("this").To(Equal("something"))
		})
	})
	Describe("Find a user by veterinary practice", func() {
		It("proper test cases and implementation of code", func() {
			Expect("this").To(Equal("something"))
		})
	})
	Describe("Find a user by name", func() {
		It("proper test cases and implementation of code", func() {
			Expect("this").To(Equal("something"))
		})
	})
	Describe("Find a user by ID", func() {
		It("proper test cases and implementation of code", func() {
			Expect("this").To(Equal("something"))
		})
	})
})
