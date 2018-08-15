package repository_test

import (
	"github.com/jinzhu/gorm"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestRepository(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Repository Suite")
}

var (
	Database *gorm.DB
	err      error
)

var _ = BeforeSuite(func() {
	// TODO: initialize database and verify it's up and running
	Database, err = gorm.Open("postgres", "host=localhost user=postgres dbname=vetlab sslmode=disable")
	Expect(err).NotTo(HaveOccurred())
})

var _ = AfterSuite(func() {
	// TODO: cleanup database
	err = Database.Close()
	Expect(err).NotTo(HaveOccurred())
})
