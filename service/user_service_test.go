package service_test

import (
	. "github.com/jmelchio/vetlab/service"
	"github.com/jmelchio/vetlab/service/servicefakes"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("UserService", func() {
	var (
		userService User
		userRepo    UserRepo
	)
	BeforeEach(func() {
		userRepo = new(servicefakes.FakeUserRepo)
		userService = User{UserRepo: userRepo}
	})
	Describe("Create a user", func() {
		It("proper test cases and implementation of code", func() {
			Expect("this").To(Equal("something"))
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
