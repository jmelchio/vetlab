package repository_test

import (
	"github.com/jmelchio/vetlab/model"
	. "github.com/jmelchio/vetlab/repository"
	"github.com/jmelchio/vetlab/service"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("SqlUserRepo", func() {

	var (
		userRepo  service.UserRepo
		userOne   model.User
		userTwo   model.User
		userName  string
		firstName string
		lastName  string
		email     string
		password  string
	)

	BeforeEach(func() {
		userRepoImpl := SQLUserRepo{Database: database}
		userRepo = userRepoImpl

		userName = "user_name"
		firstName = "first_name"
		lastName = "last_name"
		email = "first.last@gmail.com"
		password = "want_some_hash?"

		userOne = model.User{
			UserName:  &userName,
			FirstName: firstName,
			LastName:  lastName,
			Email:     email,
			Password:  password,
			AdminUser: false,
		}
	})

	AfterEach(func() {
		err = database.Where("1 = 1").Delete(&model.User{}).Error
		Expect(err).NotTo(HaveOccurred())
	})

	Describe("User table", func() {

		Context("User table has been created during in BeforeSuite", func() {

			It("Has a user table", func() {
				hasUserTable := database.HasTable(&model.User{})
				Expect(hasUserTable).To(BeTrue())
			})
		})
	})

	Describe("Create a user", func() {

		Context("When a username is not taken yet", func() {
			BeforeEach(func() {
				Expect(userOne.ID).To(Equal(uint(0)))
			})

			It("Creates a new user record", func() {
				err = userRepo.Create(&userOne)
				Expect(err).NotTo(HaveOccurred())
				Expect(userOne.ID).NotTo(Equal(uint(0)))
			})
		})

		Context("When a username is taken already", func() {

			BeforeEach(func() {
				userTwo = userOne
			})

			It("It returns an error", func() {
				err = userRepo.Create(&userOne)
				Expect(err).NotTo(HaveOccurred())
				err = userRepo.Create(&userTwo)
				Expect(err).To(HaveOccurred())
			})
		})
	})

	Describe("Update a user", func() {

		Context("When a user is found", func() {

			BeforeEach(func() {
				err = userRepo.Create(&userOne)
				Expect(err).NotTo(HaveOccurred())
				Expect(userOne.ID).NotTo(Equal(uint(0)))
			})

			Context("When the updated password is less than 50 characters", func() {

				It("It updates the user record and returns updated user with unchanged password", func() {
					*userOne.UserName = "new_user_name"
					userOne.Password = "short_password"
					err = userRepo.Update(&userOne)
					Expect(err).NotTo(HaveOccurred())
					userFound, ferr := userRepo.GetByID(userOne.ID)
					Expect(ferr).NotTo(HaveOccurred())
					Expect(userFound.Password).To(Equal("want_some_hash?"))
					Expect(userFound.UserName).To(Equal(userOne.UserName))
				})
			})

			Context("When the updated password is more than 50 characters", func() {

				It("It updates the user record and returns updated user with unchanged password", func() {
					*userOne.UserName = "new_user_name"
					userOne.Password = "long_password_of_more_than_fifty_characters_so_that_its"
					err = userRepo.Update(&userOne)
					Expect(err).NotTo(HaveOccurred())
					userFound, ferr := userRepo.GetByID(userOne.ID)
					Expect(ferr).NotTo(HaveOccurred())
					Expect(userFound.Password).To(Equal(userOne.Password))
					Expect(userFound.UserName).To(Equal(userOne.UserName))
				})
			})
		})

		Context("When the user does not exist", func() {

			BeforeEach(func() {
			})

			It("Returns an error and nil for the user", func() {
				err = userRepo.Update(&userOne)
				Expect(err).To(HaveOccurred())
			})
		})
	})

	Describe("Delete a user", func() {

		Context("When the use is found", func() {

			BeforeEach(func() {
				err = userRepo.Create(&userOne)
				Expect(err).NotTo(HaveOccurred())
				Expect(userOne.ID).NotTo(Equal(uint(0)))
			})

			It("Deletes the record and returns no error", func() {
				err = userRepo.Delete(&userOne)
				Expect(err).NotTo(HaveOccurred())
				var foundUser *model.User
				foundUser, err = userRepo.GetByID(userOne.ID)
				Expect(err).To(HaveOccurred())
				Expect(foundUser).To(BeNil())
			})
		})
	})

	Describe("Get a user by ID", func() {

		Context("When the user is found", func() {

			var foundUser *model.User

			BeforeEach(func() {
				err = userRepo.Create(&userOne)
				Expect(err).NotTo(HaveOccurred())
				Expect(userOne.ID).NotTo(Equal(uint(0)))
			})

			It("It returns the user and nil for error", func() {
				foundUser, err = userRepo.GetByID(userOne.ID)
				Expect(err).NotTo(HaveOccurred())
				Expect(foundUser).NotTo(BeNil())
				Expect(foundUser.UserName).To(Equal(userOne.UserName))
			})
		})

		Context("When the user is not found", func() {

			var foundUser *model.User

			It("It returns and error and nil for the user", func() {
				foundUser, err = userRepo.GetByID(uint(10))
				Expect(err).To(HaveOccurred())
				Expect(foundUser).To(BeNil())
			})
		})
	})

	Describe("Get a user by UserName", func() {

		Context("When the user is found", func() {

			var foundUser *model.User

			BeforeEach(func() {
				err = userRepo.Create(&userOne)
				Expect(err).NotTo(HaveOccurred())
				Expect(userOne.ID).NotTo(Equal(uint(0)))
			})

			It("It returns the user and nil for error", func() {
				foundUser, err = userRepo.GetByUserName(*userOne.UserName)
				Expect(err).NotTo(HaveOccurred())
				Expect(foundUser).NotTo(BeNil())
				Expect(foundUser.ID).To(Equal(userOne.ID))
				Expect(foundUser.UserName).To(Equal(userOne.UserName))
			})
		})

		Context("When the user is not found", func() {

			var foundUser *model.User

			BeforeEach(func() {
			})

			It("It returns the user and nil for error", func() {
				foundUser, err = userRepo.GetByUserName("some_user_name")
				Expect(err).To(HaveOccurred())
				Expect(foundUser).To(BeNil())
			})
		})
	})
})
