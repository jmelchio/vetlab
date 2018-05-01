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
		Context("We have a valid user and 'todo' context", func() {
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
			It("Returns a user with a new user ID and calls UserRepo.Create", func() {
				zeUser, err := userService.CreateUser(context.TODO(), newUser)
				Expect(err).ToNot(HaveOccurred())
				Expect(zeUser.UserID).To(Equal("created-user-id"))
				Expect(userRepo.CreateCallCount()).To(Equal(1))
			})
		})
		Context("We have a valid user but no context", func() {
			It("Returns an error and does not call UserRepo.Create", func() {
				_, err := userService.CreateUser(nil, newUser)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("Context is required"))
				Expect(userRepo.CreateCallCount()).To(Equal(0))
			})
		})
		Context("We have a user and Context but repo cannot create user", func() {
			BeforeEach(func() {
				userRepo.CreateReturns(nil, errors.New("Unable to create user"))
			})
			It("Returns an error after calling UserRepo.Create", func() {
				_, err := userService.CreateUser(context.TODO(), newUser)
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("Unable to create user"))
				Expect(userRepo.CreateCallCount()).To(Equal(1))
			})
		})
	})
	Describe("Update a user", func() {
		It("proper test cases and implementation of code", func() {
			Expect("this").To(Equal("something"))
		})
	})
	Describe("Delete a user", func() {
		It("proper test cases and implementation of code", func() {
			Expect("this").To(Equal("something"))
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
