package sql_test

import (
	"github.com/jmelchio/vetlab/model"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

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
	database, err = gorm.Open(postgres.Open( "host=localhost port=5432 user=postgres password=password dbname=vetlab sslmode=disable"), &gorm.Config{})
	Expect(err).NotTo(HaveOccurred())

	err = database.Migrator().AutoMigrate(&model.DiagnosticReport{})
	Expect(err).NotTo(HaveOccurred())
	err = database.Migrator().AutoMigrate(&model.DiagnosticRequest{})
	Expect(err).NotTo(HaveOccurred())
	err = database.Migrator().AutoMigrate(&model.Customer{})
	Expect(err).NotTo(HaveOccurred())
	err = database.Migrator().AutoMigrate(&model.VetOrg{})
	Expect(err).NotTo(HaveOccurred())
	err = database.Migrator().AutoMigrate(&model.User{})
	Expect(err).NotTo(HaveOccurred())
})

var _ = AfterSuite(func() {
	_ = database.Migrator().DropTable(&model.DiagnosticReport{})
	_ = database.Migrator().DropTable(&model.DiagnosticRequest{})
	_ = database.Migrator().DropTable(&model.Customer{})
	_ = database.Migrator().DropTable(&model.VetOrg{})
	_ = database.Migrator().DropTable(&model.User{})

})
