package repository_test

import (
	. "github.com/jmelchio/vetlab/repository"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("SqlUserRepo", func() {

	var (
		userRepo *SQLUserRepo
	)

	BeforeEach(func() {
		userRepo = new(SQLUserRepo)
	})

	Describe("Create", func() {

		Context("When a username is not taken yet", func() {

			It("Creates a new user record", func() {
				Expect("this").NotTo(Equal("this"))
			})
		})

		Context("When a username is taken already", func() {

			It("It updates the user record", func() {
				Expect("this").NotTo(Equal("this"))
			})
		})
	})

	Describe("Update", func() {

		Context("When a user is found", func() {

			It("It updates the user record and returns updated user", func() {
				Expect("this").NotTo(Equal("this"))
			})
		})

		Context("When the user does not exist", func() {

			It("Returns an error and nil for the user", func() {
				Expect("this").NotTo(Equal("this"))
			})
		})
	})

	Describe("Delete", func() {

		Context("When the use is found", func() {

			It("Deletes the record and returns no error", func() {
				Expect("this").NotTo(Equal("this"))
			})
		})

		Context("When the user does not exist", func() {

			It("Returns an error", func() {
				Expect("this").NotTo(Equal("this"))
			})
		})
	})

	Describe("GetByID", func() {

		Context("When the user is found", func() {

			It("It returns the user and nil for error", func() {
				Expect("this").NotTo(Equal("this"))
			})
		})

		Context("When the user is not found", func() {

			It("It returns nil for the user and nil for an error", func() {
				Expect("this").NotTo(Equal("this"))
			})
		})
	})

	Describe("GetByOrgID", func() {

		Context("When the user is found", func() {

			It("It returns the user and nil for the error", func() {
				Expect("this").NotTo(Equal("this"))
			})
		})

		Context("When the user is not found", func() {

			It("It returns nil for user and nil for error", func() {
				Expect("this").NotTo(Equal("this"))
			})
		})
	})
})
