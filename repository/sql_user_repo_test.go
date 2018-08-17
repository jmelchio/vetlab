package repository_test

import (
	"github.com/jmelchio/vetlab/model"
	. "github.com/jmelchio/vetlab/repository"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("SqlUserRepo", func() {

	var (
		userRepo *SQLUserRepo
		userOne  model.User
		userTwo  model.User
	)

	BeforeEach(func() {
		userRepo = new(SQLUserRepo)
		userRepo.Database = database
	})

	AfterEach(func() {
		err := userRepo.Database.Where("1 = 1").Delete(&model.User{}).Error
		Expect(err).NotTo(HaveOccurred())
	})

	Describe("User table", func() {

		Context("User table has been created during in BeforeSuite", func() {

			It("Has a user table", func() {
				hasUserTable := userRepo.Database.HasTable(&model.User{})
				Expect(hasUserTable).To(BeTrue())
			})
		})
	})

	FDescribe("Create a user", func() {

		Context("When a username is not taken yet", func() {
			BeforeEach(func() {
				userOne = model.User{
					UserName:     "user_name",
					FirstName:    "first_name",
					LastName:     "last_name",
					Email:        "first.last@gmail.com",
					PasswordHash: "want_some_hash?",
					AdminUser:    false,
				}
			})

			It("Creates a new user record", func() {
				newUser, err := userRepo.Create(userOne)
				Expect(err).NotTo(HaveOccurred())
				Expect(newUser).NotTo(BeNil())
				Expect(newUser.ID).NotTo(Equal(0))
				Expect(newUser.Email).To(Equal(userOne.Email))
			})
		})

		Context("When a username is taken already", func() {

			BeforeEach(func() {
				userOne = model.User{
					UserName:     "user_name",
					FirstName:    "first_name",
					LastName:     "last_name",
					Email:        "first.last@gmail.com",
					PasswordHash: "want_some_hash?",
					AdminUser:    false,
				}
				userTwo = userOne
			})

			It("It returns an error", func() {
				_, err := userRepo.Create(userOne)
				Expect(err).NotTo(HaveOccurred())
				_, err = userRepo.Create(userTwo)
				Expect(err).To(HaveOccurred())
			})
		})
	})

	Describe("Update a user", func() {

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

	Describe("Delete a user", func() {

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

	Describe("Get a user by ID", func() {

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

	Describe("Get a user by OrgID", func() {

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
