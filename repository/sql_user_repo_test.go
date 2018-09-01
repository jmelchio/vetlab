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

	Describe("Create a user", func() {

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
				Expect(userOne.ID).To(Equal(uint(0)))
			})

			It("Creates a new user record", func() {
				err := userRepo.Create(&userOne)
				Expect(err).NotTo(HaveOccurred())
				Expect(userOne.ID).NotTo(Equal(uint(0)))
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
				err := userRepo.Create(&userOne)
				Expect(err).NotTo(HaveOccurred())
				err = userRepo.Create(&userTwo)
				Expect(err).To(HaveOccurred())
			})
		})
	})

	Describe("Update a user", func() {

		Context("When a user is found", func() {

			BeforeEach(func() {
				userOne = model.User{
					UserName:     "user_name",
					FirstName:    "first_name",
					LastName:     "last_name",
					Email:        "first.last@gmail.com",
					PasswordHash: "want_some_hash?",
					AdminUser:    false,
				}
				err = userRepo.Create(&userOne)
				Expect(err).NotTo(HaveOccurred())
				Expect(userOne.ID).NotTo(Equal(uint(0)))
			})

			It("It updates the user record and returns updated user", func() {
				userOne.UserName = "new_user_name"
				err = userRepo.Update(&userOne)
				Expect(err).NotTo(HaveOccurred())
			})
		})

		Context("When the user does not exist", func() {

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

			It("Returns an error and nil for the user", func() {
				err = userRepo.Update(&userOne)
				Expect(err).To(HaveOccurred())
			})
		})
	})

	Describe("Delete a user", func() {

		Context("When the use is found", func() {

			BeforeEach(func() {
				userOne = model.User{
					UserName:     "user_name",
					FirstName:    "first_name",
					LastName:     "last_name",
					Email:        "first.last@gmail.com",
					PasswordHash: "want_some_hash?",
					AdminUser:    false,
				}
				err = userRepo.Create(&userOne)
				Expect(err).NotTo(HaveOccurred())
				Expect(userOne.ID).NotTo(Equal(uint(0)))
			})

			It("Deletes the record and returns no error", func() {
				err = userRepo.Delete(&userOne)
				Expect(err).NotTo(HaveOccurred())
				// Todo: need to fetch and see if it finds in DB
			})
		})
	})

	Describe("Get a user by ID", func() {

		Context("When the user is found", func() {

			var foundUser *model.User

			BeforeEach(func() {
				userOne = model.User{
					UserName:     "user_name",
					FirstName:    "first_name",
					LastName:     "last_name",
					Email:        "first.last@gmail.com",
					PasswordHash: "want_some_hash?",
					AdminUser:    false,
				}
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

	Describe("Get a user by OrgID", func() {

		Context("When the user is found", func() {

			var foundUsers []model.User

			BeforeEach(func() {
				userOne = model.User{
					UserName:     "user_name",
					FirstName:    "first_name",
					LastName:     "last_name",
					Email:        "first.last@gmail.com",
					PasswordHash: "want_some_hash?",
					OrgID:        uint(10),
					AdminUser:    false,
				}
				err = userRepo.Create(&userOne)
				Expect(err).NotTo(HaveOccurred())
				Expect(userOne.ID).NotTo(Equal(uint(0)))
			})

			It("It returns the user and nil for error", func() {
				foundUsers, err = userRepo.GetByOrgID(uint(10))
				Expect(err).NotTo(HaveOccurred())
				Expect(foundUsers).NotTo(BeNil())
				Expect(len(foundUsers)).To(Equal(1))
				Expect(foundUsers[0].OrgID).To(Equal(uint(10)))
				Expect(foundUsers[0].UserName).To(Equal(userOne.UserName))
			})
		})

		Context("When the user is not found", func() {

			var foundUsers []model.User

			BeforeEach(func() {
			})

			It("It returns nil for result and an error", func() {
				foundUsers, err = userRepo.GetByOrgID(uint(10))
				Expect(err).To(HaveOccurred())
				Expect(foundUsers).To(BeNil())
			})
		})

		Describe("Get a user by UserName", func() {

			Context("When the user is found", func() {

				var foundUser *model.User

				BeforeEach(func() {
					userOne = model.User{
						UserName:     "user_name",
						FirstName:    "first_name",
						LastName:     "last_name",
						Email:        "first.last@gmail.com",
						PasswordHash: "want_some_hash?",
						AdminUser:    false,
					}
					err = userRepo.Create(&userOne)
					Expect(err).NotTo(HaveOccurred())
					Expect(userOne.ID).NotTo(Equal(uint(0)))
				})

				It("It returns the user and nil for error", func() {
					foundUser, err = userRepo.GetByUserName(userOne.UserName)
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
})
