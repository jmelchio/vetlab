package repository_test

import (
	"github.com/jinzhu/gorm"
	"github.com/jmelchio/vetlab/model"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestRepository(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Repository Suite")
}

var (
	database *gorm.DB
	err      error
)

var _ = BeforeSuite(func() {
	database, err = gorm.Open("postgres", "host=localhost user=postgres dbname=vetlab sslmode=disable")
	Expect(err).NotTo(HaveOccurred())

	database.AutoMigrate(&model.User{})
})

var _ = AfterSuite(func() {
	database.DropTable(&model.User{})

	err = database.Close()
	Expect(err).NotTo(HaveOccurred())
})
