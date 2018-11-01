package sql_test

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

	err = database.AutoMigrate(&model.DiagnosticReport{}).Error
	Expect(err).NotTo(HaveOccurred())
	err = database.AutoMigrate(&model.DiagnosticRequest{}).Error
	Expect(err).NotTo(HaveOccurred())
	err = database.AutoMigrate(&model.Customer{}).Error
	Expect(err).NotTo(HaveOccurred())
	err = database.AutoMigrate(&model.VetOrg{}).Error
	Expect(err).NotTo(HaveOccurred())
	err = database.AutoMigrate(&model.User{}).Error
	Expect(err).NotTo(HaveOccurred())
})

var _ = AfterSuite(func() {
	database.DropTable(&model.DiagnosticReport{})
	database.DropTable(&model.DiagnosticRequest{})
	database.DropTable(&model.Customer{})
	database.DropTable(&model.VetOrg{})
	database.DropTable(&model.User{})

	err = database.Close()
	Expect(err).NotTo(HaveOccurred())
})
